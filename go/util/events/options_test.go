package events

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"pkg.akt.dev/go/util/ctxlog"
	"pkg.akt.dev/go/util/pubsub"
)

func TestOptionsDefaults(t *testing.T) {
	o, err := newOptions()
	require.NoError(t, err)

	require.Equal(t, 8*time.Second, o.watchdog)
	require.Equal(t, DefaultWatchdogTimeout, o.watchdog)
	require.Equal(t, 10*time.Second, o.fetchTimeout)
	require.Equal(t, DefaultFetchTimeout, o.fetchTimeout)
	require.Equal(t, 5*time.Second, o.retryInterval)
	require.Equal(t, DefaultRetryInterval, o.retryInterval)
	require.Equal(t, 4, o.prefetch)
	require.Equal(t, DefaultPrefetch, o.prefetch)
	require.Equal(t, 5*time.Second, o.pingPeriod)
	require.Equal(t, 15*time.Second, o.pongWait)
	require.Equal(t, defaultPrunedCheckAfter, o.prunedCheckAfter)
	require.Nil(t, o.log, "no logger means the one carried by the context")
	require.False(t, o.startHeightSet, "no start height means the node's latest")
}

func TestOptionsApply(t *testing.T) {
	rec := newRecLogger()

	o, err := newOptions(WithWatchdogTimeout(3*time.Second), WithFetchTimeout(4*time.Second), WithRetryInterval(7*time.Second),
		WithPrefetch(2), withPing(time.Second, 3*time.Second), WithLogger(rec), WithStartHeight(0), nil)
	require.NoError(t, err)

	require.Equal(t, 3*time.Second, o.watchdog)
	require.Equal(t, 4*time.Second, o.fetchTimeout)
	require.Equal(t, 7*time.Second, o.retryInterval)
	require.Equal(t, 2, o.prefetch)
	require.Equal(t, time.Second, o.pingPeriod)
	require.Equal(t, 3*time.Second, o.pongWait)
	require.Equal(t, rec, o.log)
	require.True(t, o.startHeightSet, "zero is a valid start height: publish from the first block")
	require.Zero(t, o.startHeight)
}

// An invalid option fails NewEvents before it touches the node.
func TestNewEventsRejectsInvalidOptions(t *testing.T) {
	for name, opt := range map[string]Option{
		"zero watchdog":            WithWatchdogTimeout(0),
		"negative watchdog":        WithWatchdogTimeout(-time.Second),
		"zero fetch timeout":       WithFetchTimeout(0),
		"negative fetch timeout":   WithFetchTimeout(-time.Second),
		"zero retry interval":      WithRetryInterval(0),
		"negative retry interval":  WithRetryInterval(-time.Second),
		"zero prefetch":            WithPrefetch(0),
		"negative prefetch":        WithPrefetch(-1),
		"zero ping period":         withPing(0, time.Second),
		"pong wait not above ping": withPing(time.Second, time.Second),
		"nil logger":               WithLogger(nil),
		"negative start height":    WithStartHeight(-1),
	} {
		t.Run(name, func(t *testing.T) {
			bus := pubsub.NewBus()
			t.Cleanup(bus.Close)

			// Nothing listens on port 1, so any error but the option's own means the
			// node was contacted before the options were checked.
			svc, err := NewEvents(context.Background(), "tcp://127.0.0.1:1", "test", bus, opt)
			require.EqualError(t, err, opt(&options{}).Error())
			require.Nil(t, svc)
		})
	}
}

// NewEvents applies each option once, so an option need not be idempotent.
func TestNewEventsAppliesEachOptionOnce(t *testing.T) {
	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)

	applied := 0
	count := Option(func(*options) error {
		applied++
		return nil
	})

	// Nothing listens on port 1, so NewEvents fails on its first call to the node, after
	// it has taken its options.
	svc, err := NewEvents(context.Background(), "tcp://127.0.0.1:1", "test", bus, count)
	require.Error(t, err)
	require.Nil(t, svc)
	require.Equal(t, 1, applied)
}

// A node URI the service cannot reach a node through fails NewEvents.
func TestNewEventsRejectsInvalidNodeURI(t *testing.T) {
	for _, uri := range []string{"", "localhost:26657", "ftp://node:21"} {
		t.Run(uri, func(t *testing.T) {
			bus := pubsub.NewBus()
			t.Cleanup(bus.Close)

			svc, err := NewEvents(context.Background(), uri, "test", bus)
			require.Error(t, err)
			require.Nil(t, svc)
		})
	}
}

// Without WithLogger the service logs to the logger carried by the context, as it did
// before the option existed. With it, the given logger wins.
func TestNewEventsLoggerSelection(t *testing.T) {
	tune := tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}

	start := func(t *testing.T, ctxLogger recLogger, opts ...Option) {
		t.Helper()

		ctx, cancel := context.WithCancel(ctxlog.WithLogger(context.Background(), ctxLogger))
		t.Cleanup(cancel)

		bus := pubsub.NewBus()
		t.Cleanup(bus.Close)

		node := newFakeNode()
		svc, err := newEventsOn(ctx, node, bus, append(tune.options(), opts...)...)
		require.NoError(t, err)
		t.Cleanup(svc.Shutdown)
	}

	t.Run("context", func(t *testing.T) {
		fromCtx := newRecLogger()
		start(t, fromCtx)

		require.Eventually(t, func() bool { return fromCtx.has(levelInfo, "dialing block subscription") }, 2*time.Second, time.Millisecond)
	})

	t.Run("option", func(t *testing.T) {
		fromCtx, given := newRecLogger(), newRecLogger()
		start(t, fromCtx, WithLogger(given))

		require.Eventually(t, func() bool { return given.has(levelInfo, "dialing block subscription") }, 2*time.Second, time.Millisecond)
		require.False(t, fromCtx.has(levelInfo, "dialing block subscription"), "WithLogger must take precedence over the context logger")
	})
}
