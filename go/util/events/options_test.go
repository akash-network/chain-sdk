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
	require.Equal(t, defaultMinRetryInterval, o.minRetryInterval)
	require.Equal(t, defaultPrunedCheckAfter, o.prunedCheckAfter)
	require.Nil(t, o.log, "no logger means the one carried by the context")
	require.False(t, o.startHeightSet, "no start height means the node's latest")
}

func TestOptionsApply(t *testing.T) {
	rec := newRecLogger()

	o, err := newOptions(WithWatchdogTimeout(3*time.Second), WithLogger(rec), WithStartHeight(0), nil)
	require.NoError(t, err)

	require.Equal(t, 3*time.Second, o.watchdog)
	require.Equal(t, rec, o.log)
	require.True(t, o.startHeightSet, "zero is a valid start height: publish from the first block")
	require.Zero(t, o.startHeight)
}

// An invalid option fails NewEvents before it touches the node.
func TestNewEventsRejectsInvalidOptions(t *testing.T) {
	for name, opt := range map[string]Option{
		"zero watchdog":         WithWatchdogTimeout(0),
		"negative watchdog":     WithWatchdogTimeout(-time.Second),
		"nil logger":            WithLogger(nil),
		"negative start height": WithStartHeight(-1),
	} {
		t.Run(name, func(t *testing.T) {
			node := newFakeNode()
			bus := pubsub.NewBus()
			t.Cleanup(bus.Close)

			svc, err := NewEvents(context.Background(), node, "test", bus, opt)
			require.Error(t, err)
			require.Nil(t, svc)
			require.Zero(t, node.subscribeCount(), "options are validated before the node is touched")
		})
	}
}

// Without WithLogger the service logs to the logger carried by the context, as it did
// before the option existed. With it, the given logger wins.
func TestNewEventsLoggerSelection(t *testing.T) {
	tune := tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}

	start := func(t *testing.T, ctxLogger recLogger, opts ...Option) {
		t.Helper()

		ctx, cancel := context.WithCancel(ctxlog.WithLogger(context.Background(), ctxLogger))
		t.Cleanup(cancel)

		bus := pubsub.NewBus()
		t.Cleanup(bus.Close)

		svc, err := NewEvents(ctx, newFakeNode(), "test", bus, append(tune.options(), opts...)...)
		require.NoError(t, err)
		t.Cleanup(svc.Shutdown)
	}

	t.Run("context", func(t *testing.T) {
		fromCtx := newRecLogger()
		start(t, fromCtx)

		require.Eventually(t, func() bool { return fromCtx.has(levelWarn, "block subscription presumed lost") }, 2*time.Second, time.Millisecond)
	})

	t.Run("option", func(t *testing.T) {
		fromCtx, given := newRecLogger(), newRecLogger()
		start(t, fromCtx, WithLogger(given))

		require.Eventually(t, func() bool { return given.has(levelWarn, "block subscription presumed lost") }, 2*time.Second, time.Millisecond)
		require.False(t, fromCtx.has(levelWarn, "block subscription presumed lost"), "WithLogger must take precedence over the context logger")
	})
}
