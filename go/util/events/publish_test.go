package events

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	cmtypes "github.com/cometbft/cometbft/types"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
	"pkg.akt.dev/go/util/pubsub"
)

// fakeNode implements the CometRPC and EventsClient methods the events service uses.
type fakeNode struct {
	sdkclient.CometRPC

	mu       sync.Mutex
	height   int64
	blocks   map[int64]*coretypes.ResultBlockResults
	blockErr map[int64]error
	headerCh chan coretypes.ResultEvent
}

func newFakeNode() *fakeNode {
	return &fakeNode{
		height:   100,
		blocks:   map[int64]*coretypes.ResultBlockResults{},
		blockErr: map[int64]error{},
		headerCh: make(chan coretypes.ResultEvent, 8),
	}
}

func (f *fakeNode) set(mutate func(*fakeNode)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	mutate(f)
}

func (f *fakeNode) pushHeader(h int64) {
	f.headerCh <- coretypes.ResultEvent{Data: cmtypes.EventDataNewBlockHeader{Header: cmtypes.Header{Height: h}}}
}

func (f *fakeNode) Status(context.Context) (*coretypes.ResultStatus, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return &coretypes.ResultStatus{SyncInfo: coretypes.SyncInfo{LatestBlockHeight: f.height}}, nil
}

func (f *fakeNode) BlockResults(_ context.Context, h *int64) (*coretypes.ResultBlockResults, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if err := f.blockErr[*h]; err != nil {
		return nil, err
	}
	if b, ok := f.blocks[*h]; ok {
		return b, nil
	}
	return &coretypes.ResultBlockResults{Height: *h}, nil
}

func (f *fakeNode) Subscribe(context.Context, string, string, ...int) (<-chan coretypes.ResultEvent, error) {
	return f.headerCh, nil
}

func (f *fakeNode) Unsubscribe(context.Context, string, string) error { return nil }
func (f *fakeNode) UnsubscribeAll(context.Context, string) error      { return nil }

func leaseClosedEvent(t *testing.T, dseq uint64) abci.Event {
	t.Helper()
	ev, err := sdk.TypedEventToEvent(&mtypes.EventLeaseClosed{
		ID: mtypes.LeaseID{Owner: "akash1owner", DSeq: dseq, GSeq: 1, OSeq: 1, Provider: "akash1provider"},
	})
	require.NoError(t, err)
	return abci.Event(ev)
}

func txBlock(height int64, evs ...abci.Event) *coretypes.ResultBlockResults {
	return &coretypes.ResultBlockResults{
		Height:     height,
		TxsResults: []*abci.ExecTxResult{{Events: evs}},
	}
}

func newTestService(t *testing.T, node sdkclient.CometRPC, watchdogTimeout time.Duration) pubsub.Subscriber {
	t.Helper()

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := newEvents(ctx, node, "test", bus, watchdogTimeout)
	require.NoError(t, err)
	t.Cleanup(func() {
		svc.Shutdown()
		bus.Close()
	})

	return sub
}

func waitLeaseClosed(t *testing.T, sub pubsub.Subscriber, wantDseq uint64) {
	t.Helper()
	deadline := time.After(2 * time.Second)
	for {
		select {
		case ev := <-sub.Events():
			if closed, ok := ev.(*mtypes.EventLeaseClosed); ok && closed.ID.DSeq == wantDseq {
				return
			}
		case <-deadline:
			t.Fatalf("did not receive EventLeaseClosed for dseq %d", wantDseq)
		}
	}
}

// A block header from the subscription drives processing of that block.
func TestEventsProcessesBlockOnHeader(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, time.Hour) // watchdog idle: header drives processing

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// A header that jumps several heights (as after a dropped and re-established
// subscription) back-fills every block in between, in order.
func TestEventsBackfillsGapBetweenHeaders(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, time.Hour) // watchdog idle: header drives processing

	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
		f.blocks[103] = txBlock(103, leaseClosedEvent(t, 103))
	})
	node.pushHeader(103)

	waitLeaseClosed(t, sub, 101)
	waitLeaseClosed(t, sub, 102)
	waitLeaseClosed(t, sub, 103)
}

// When the subscription goes quiet (dies without closing its channel), the fetch
// fallback catches up so events keep flowing.
func TestEventsFetchFallbackWhenSubscriptionQuiet(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, 10*time.Millisecond)

	// No header pushed: the subscription is silent. A new block appears on chain.
	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.height = 101
	})

	waitLeaseClosed(t, sub, 101)
}

// A height that never becomes fetchable (for example pruned on the node) is retried
// and then skipped, so a single bad block does not wedge the pipeline forever: the
// next block is still processed.
func TestEventsSkipsUnfetchableBlock(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, 5*time.Millisecond)

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errors.New("height 101 is not available, lowest height is 200")
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
		f.height = 102
	})

	waitLeaseClosed(t, sub, 102)
}

// The service starts from the current height, so events in already-committed blocks
// are not replayed.
func TestEventsDoesNotReplayHistory(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) { f.blocks[100] = txBlock(100, leaseClosedEvent(t, 100)) })

	sub := newTestService(t, node, 5*time.Millisecond) // watchdog active, but nothing to replay

	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected replay of historical event: %T", ev)
	case <-time.After(100 * time.Millisecond):
	}
}

func TestProcessEventFiltersUnrecognized(t *testing.T) {
	_, ok := processEvent(leaseClosedEvent(t, 1))
	require.True(t, ok, "a lease-closed event must be recognized")

	other := abci.Event{Type: "coin_received", Attributes: []abci.EventAttribute{{Key: "amount", Value: "1uakt"}}}
	_, ok = processEvent(other)
	require.False(t, ok, "a non-Akash event must be ignored")
}
