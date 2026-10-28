package cli

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	"pkg.akt.dev/go/util/pubsub"
)

// The events command dials the node itself: the one the client context names, else the
// one the command's --node flag names, and with neither it says so.
func TestGetEventsNodeURI(t *testing.T) {
	run := func(t *testing.T, cctx *sdkclient.Context, node string) error {
		t.Helper()

		cmd := EventsCmd()
		ctx, cancel := context.WithCancel(context.Background())
		t.Cleanup(cancel)
		if cctx != nil {
			ctx = context.WithValue(ctx, sdkclient.ClientContextKey, cctx)
		}
		cmd.SetContext(ctx)
		require.NoError(t, cmd.Flags().Set(cflags.FlagNode, node))

		return getEvents(ctx, cmd, nil)
	}

	t.Run("none", func(t *testing.T) {
		require.EqualError(t, run(t, nil, ""), "events: no node URI")
	})

	t.Run("flag", func(t *testing.T) {
		// Nothing listens on port 1: the service failing to reach it proves the flag's
		// address was the one dialed.
		require.ErrorContains(t, run(t, nil, "tcp://127.0.0.1:1"), "127.0.0.1:1")
	})

	t.Run("context", func(t *testing.T) {
		// Nothing listens on port 2 either. The service dials one node, so failing to
		// reach port 2 proves the client context's node won over the flag's.
		err := run(t, &sdkclient.Context{NodeURI: "tcp://127.0.0.1:2"}, "tcp://127.0.0.1:1")
		require.ErrorContains(t, err, "127.0.0.1:2")
	})
}

// stubEvents is an events.Service whose state the test sets directly.
type stubEvents struct {
	done      chan struct{}
	err       error
	shutdowns int
}

func (s *stubEvents) Shutdown()             { s.shutdowns++ }
func (s *stubEvents) Done() <-chan struct{} { return s.done }
func (s *stubEvents) Err() error            { return s.err }

// The events command fails with the error the events service stopped on, rather than
// waiting on a service that no longer publishes anything.
func TestWaitEventsReturnsServiceError(t *testing.T) {
	svc := &stubEvents{done: make(chan struct{}), err: pubsub.ErrNotRunning}
	close(svc.done)

	require.ErrorIs(t, waitEvents(context.Background(), svc), pubsub.ErrNotRunning)
}

// Interrupting the command shuts the events service down, and is not an error.
func TestWaitEventsShutsServiceDownOnCancel(t *testing.T) {
	svc := &stubEvents{done: make(chan struct{})}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	require.NoError(t, waitEvents(ctx, svc))
	require.Equal(t, 1, svc.shutdowns)
}

// The printer returns once the command's context ends, so a failed events service ends
// the command instead of leaving it waiting on a bus nobody publishes to.
func TestPrintEventsStopsOnCancel(t *testing.T) {
	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)

	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := make(chan error, 1)
	go func() { done <- printEvents(ctx, sdkclient.Context{}, sub) }()

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(time.Second):
		t.Fatal("printer kept running after the context ended")
	}
}
