package v1beta3

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/boz/go-lifecycle"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
)

// schedulerHarness wires a bare serialBroadcaster and drives its run() loop.
// Tests enqueue broadcastReq values on reqch, observe dispatch order via
// broadcastch, and complete each in-flight broadcast by sending on donech of
// the dispatched struct.
type schedulerHarness struct {
	t      *testing.T
	c      *serialBroadcaster
	cancel context.CancelFunc
	done   chan struct{}
}

func newSchedulerHarness(t *testing.T) *schedulerHarness {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())

	h := &schedulerHarness{
		t: t,
		c: &serialBroadcaster{
			reqch:       make(chan broadcastReq, 16),
			broadcastch: make(chan broadcast, 1),
			lc:          lifecycle.New(),
			log:         log.NewNopLogger(),
		},
		cancel: cancel,
		done:   make(chan struct{}),
	}

	go h.c.lc.WatchContext(ctx)
	go func() {
		h.c.run()
		close(h.done)
	}()

	t.Cleanup(h.shutdown)

	return h
}

func (h *schedulerHarness) shutdown() {
	h.t.Helper()
	h.cancel()
	select {
	case <-h.done:
	case <-time.After(time.Second):
		h.t.Fatal("scheduler did not shut down")
	}
}

func (h *schedulerHarness) enqueue(tag string, priority bool) {
	h.t.Helper()
	req := broadcastReq{ctx: context.Background(), data: tag}
	if priority {
		req.opts = &BroadcastOptions{priority: true}
	}
	select {
	case h.c.reqch <- req:
	case <-time.After(time.Second):
		h.t.Fatalf("timed out enqueueing %q", tag)
	}
}

func (h *schedulerHarness) enqueueNilOpts(tag string) {
	h.t.Helper()
	select {
	case h.c.reqch <- broadcastReq{ctx: context.Background(), data: tag}:
	case <-time.After(time.Second):
		h.t.Fatalf("timed out enqueueing %q", tag)
	}
}

// recvDispatch reads one dispatch and returns its data tag. Caller must
// complete() the returned broadcast before expecting another dispatch.
func (h *schedulerHarness) recvDispatch() (string, broadcast) {
	h.t.Helper()
	select {
	case b := <-h.c.broadcastch:
		tag, _ := b.data.(string)
		return tag, b
	case <-time.After(time.Second):
		h.t.Fatal("timed out waiting for dispatch")
		return "", broadcast{}
	}
}

func (h *schedulerHarness) expectNoDispatch(d time.Duration) {
	h.t.Helper()
	select {
	case b := <-h.c.broadcastch:
		h.t.Fatalf("unexpected dispatch while gated: %v", b.data)
	case <-time.After(d):
	}
}

func (h *schedulerHarness) complete(b broadcast, err error) {
	h.t.Helper()
	select {
	case b.donech <- err:
	case <-time.After(time.Second):
		h.t.Fatal("timed out signalling broadcast done")
	}
}

func TestSerialBroadcaster_DefaultFIFOPreserved(t *testing.T) {
	h := newSchedulerHarness(t)

	for _, tag := range []string{"a", "b", "c"} {
		h.enqueue(tag, false)
	}

	for _, want := range []string{"a", "b", "c"} {
		got, b := h.recvDispatch()
		require.Equal(t, want, got)
		h.complete(b, nil)
	}
}

func TestSerialBroadcaster_PriorityPreemptsPending(t *testing.T) {
	h := newSchedulerHarness(t)

	h.enqueue("n1", false)
	_, inFlight := h.recvDispatch()

	h.enqueue("n2", false)
	h.enqueue("n3", false)
	h.enqueue("p1", true)

	h.expectNoDispatch(50 * time.Millisecond)

	h.complete(inFlight, nil)

	for _, want := range []string{"p1", "n2", "n3"} {
		got, b := h.recvDispatch()
		require.Equal(t, want, got)
		h.complete(b, nil)
	}
}

func TestSerialBroadcaster_PriorityFIFOAmongItself(t *testing.T) {
	h := newSchedulerHarness(t)

	h.enqueue("n1", false)
	_, inFlight := h.recvDispatch()

	h.enqueue("p1", true)
	h.enqueue("p2", true)
	h.enqueue("p3", true)

	h.complete(inFlight, nil)

	for _, want := range []string{"p1", "p2", "p3"} {
		got, b := h.recvDispatch()
		require.Equal(t, want, got)
		h.complete(b, nil)
	}
}

// Arrival [N1, P1, N2, P2, N3] with N1 held in-flight dispatches as
// [N1, P1, P2, N2, N3]: priority queue drains before pending.
func TestSerialBroadcaster_MixedInterleaving(t *testing.T) {
	h := newSchedulerHarness(t)

	h.enqueue("n1", false)
	_, inFlight := h.recvDispatch()

	h.enqueue("p1", true)
	h.enqueue("n2", false)
	h.enqueue("p2", true)
	h.enqueue("n3", false)

	h.complete(inFlight, nil)

	for _, want := range []string{"p1", "p2", "n2", "n3"} {
		got, b := h.recvDispatch()
		require.Equal(t, want, got)
		h.complete(b, nil)
	}
}

func TestSerialBroadcaster_BroadcastErrorDoesNotStall(t *testing.T) {
	h := newSchedulerHarness(t)

	h.enqueue("a", false)
	h.enqueue("b", false)

	_, first := h.recvDispatch()
	h.complete(first, errors.New("boom"))

	got, second := h.recvDispatch()
	require.Equal(t, "b", got)
	h.complete(second, nil)
}

func TestSerialBroadcaster_NilOptsTreatedAsNonPriority(t *testing.T) {
	h := newSchedulerHarness(t)

	h.enqueueNilOpts("n1")
	_, inFlight := h.recvDispatch()

	h.enqueue("p1", true)
	h.enqueueNilOpts("n2")

	h.complete(inFlight, nil)

	for _, want := range []string{"p1", "n2"} {
		got, b := h.recvDispatch()
		require.Equal(t, want, got)
		h.complete(b, nil)
	}
}
