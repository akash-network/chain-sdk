package events

import (
	"context"
	"errors"
	"fmt"
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

// errTransientFetch stands in for a flaky-network error. It goes away on its own and
// carries no information the code reads, unlike a node's pruning response.
var errTransientFetch = errors.New("post failed: connection reset by peer")

// fakeNode implements the CometRPC and EventsClient methods the events service uses.
type fakeNode struct {
	sdkclient.CometRPC

	mu         sync.Mutex
	height     int64
	earliest   int64
	statusErr  error
	blocks     map[int64]*coretypes.ResultBlockResults
	blockErr   map[int64]error
	blockCalls map[int64]int
	headerCh   chan coretypes.ResultEvent
}

func newFakeNode() *fakeNode {
	return &fakeNode{
		height:     100,
		blocks:     map[int64]*coretypes.ResultBlockResults{},
		blockErr:   map[int64]error{},
		blockCalls: map[int64]int{},
		headerCh:   make(chan coretypes.ResultEvent, 64),
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

func (f *fakeNode) callsAt(h int64) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.blockCalls[h]
}

func (f *fakeNode) Status(context.Context) (*coretypes.ResultStatus, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.statusErr != nil {
		return nil, f.statusErr
	}
	return &coretypes.ResultStatus{SyncInfo: coretypes.SyncInfo{
		LatestBlockHeight:   f.height,
		EarliestBlockHeight: f.earliest,
	}}, nil
}

func (f *fakeNode) BlockResults(_ context.Context, h *int64) (*coretypes.ResultBlockResults, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.blockCalls[*h]++
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

func newTestService(t *testing.T, node sdkclient.CometRPC, tune tuning) pubsub.Subscriber {
	t.Helper()

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := newEvents(ctx, node, "test", bus, tune)
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
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog idle, header drives processing

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// A header that jumps several heights (as after a dropped and re-established
// subscription) back-fills every block in between, in order.
func TestEventsBackfillsGapBetweenHeaders(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog idle, header drives processing

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
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond})

	// No header pushed, the subscription is silent. A new block appears on chain.
	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.height = 101
	})

	waitLeaseClosed(t, sub, 101)
}

// A height whose fetch fails only transiently must not be dropped. Its events are
// published once the node recovers, even if headers arrived in a burst while the
// fetch was failing.
func TestEventsRecoversBlockAfterTransientFetchFailure(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 4 * time.Millisecond, prunedCheckAfter: 500 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.height = 115
	})

	for h := int64(101); h <= 115; h++ {
		node.pushHeader(h)
	}

	time.AfterFunc(50*time.Millisecond, func() {
		node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	})

	waitLeaseClosed(t, sub, 101)
}

// The service starts from the current height, so events in already-committed blocks
// are not replayed.
func TestEventsDoesNotReplayHistory(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) { f.blocks[100] = txBlock(100, leaseClosedEvent(t, 100)) })

	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog active, but nothing to replay

	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected replay of historical event %T", ev)
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

// A cursor that has fallen behind a static pruning boundary must jump straight to
// what the node still retains instead of walking the range one height at a time.
func TestEventsJumpsCursorPastPrunedRange(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 3 * time.Millisecond, minRetryInterval: 20 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	node.set(func(f *fakeNode) {
		for h := int64(101); h < 200; h++ {
			f.blockErr[h] = errTransientFetch
		}
		f.blocks[200] = txBlock(200, leaseClosedEvent(t, 200))
		f.earliest = 200
		f.height = 200
	})

	waitLeaseClosed(t, sub, 200)

	calls := 0
	for h := int64(101); h < 200; h++ {
		calls += node.callsAt(h)
	}
	require.LessOrEqual(t, calls, 2, "a one-at-a-time walk would call BlockResults far more than twice below the pruning boundary")
}

// A height exactly equal to the node's earliest retained height is still retained,
// so a fetch failure there must be retried, never treated as pruned.
func TestEventsRetriesHeightEqualToEarliest(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.earliest = 101
		f.height = 101
	})
	node.pushHeader(101)

	time.AfterFunc(100*time.Millisecond, func() {
		node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	})

	waitLeaseClosed(t, sub, 101)
}

// A node reporting earliest 0 (unknown) must never cause a skip. A permanently
// failing height is retried forever and nothing later is ever published.
func TestEventsDoesNotSkipWhenEarliestUnknown(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[105] = txBlock(105, leaseClosedEvent(t, 105))
		f.height = 105
	})
	node.pushHeader(105)

	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected event published while a failing height has unknown earliest, got %T", ev)
	case <-time.After(200 * time.Millisecond):
	}

	require.Greater(t, node.callsAt(101), 1, "height 101 should have been retried, not abandoned")
}

// When Status itself is unavailable, a failing height must be retried rather than
// skipped, since parkCursor cannot confirm pruning and so cannot assume it either.
func TestEventsRetriesWhenStatusUnavailable(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.earliest = 150
		f.statusErr = errors.New("rpc unavailable")
	})

	stop := make(chan struct{})
	defer close(stop)
	go func() {
		tick := time.NewTicker(5 * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tick.C:
				node.pushHeader(101)
			}
		}
	}()

	time.AfterFunc(100*time.Millisecond, func() {
		node.set(func(f *fakeNode) {
			delete(f.blockErr, 101)
			f.statusErr = nil
		})
	})

	waitLeaseClosed(t, sub, 101)
}

// A pruning floor reported above the failing height must not cause a skip until the
// stall has persisted past prunedCheckAfter. A quick recovery still gets published. This
// matters because pooled RPC backends can report very different pruning floors, and
// a single early Status sample must not be trusted over a backend that can still
// serve the height.
func TestEventsDoesNotSkipBeforePrunedThreshold(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 200 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.earliest = 150
		f.height = 101
	})
	node.pushHeader(101)

	time.AfterFunc(20*time.Millisecond, func() {
		node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	})

	waitLeaseClosed(t, sub, 101)
}

// Fifteen headers buffered at once while one height fails must cost exactly one
// fetch attempt for that height within a single minRetryInterval, not one per header.
func TestEventsBurstOfHeadersDoesNotBurnRetries(t *testing.T) {
	node := newFakeNode()
	newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 2 * time.Second, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) { f.blockErr[101] = errTransientFetch })

	for h := int64(101); h <= 115; h++ {
		node.pushHeader(h)
	}

	require.Eventually(t, func() bool { return node.callsAt(101) >= 1 }, time.Second, time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	require.Equal(t, 1, node.callsAt(101), "a burst of buffered headers must not burn through retries for one height")
}

// A height that fails forever, with earliest at or below it, must never be skipped,
// so a later and perfectly fetchable height is never published ahead of it.
func TestEventsDoesNotAdvancePastFailingBlock(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[105] = txBlock(105, leaseClosedEvent(t, 105))
		f.earliest = 50
		f.height = 105
	})
	node.pushHeader(105)

	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected event published past a permanently failing block, got %T", ev)
	case <-time.After(200 * time.Millisecond):
	}

	require.Greater(t, node.callsAt(101), 1, "height 101 should have been retried, not abandoned")
}

// Order is load bearing. bidengine creates order actors with checkForExistingBid
// false on the event path, so a replayed EventOrderCreated after its EventLeaseCreated
// would bid on an already-leased order. A stall must never let a later height publish
// before an earlier one recovers.
func TestEventsOrderPreservedAcrossStallAndRecovery(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 500 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		for h := int64(101); h <= 105; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, uint64(h)))
		}
		f.height = 105
	})
	node.pushHeader(105)

	time.AfterFunc(40*time.Millisecond, func() {
		node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	})

	var got []uint64
	deadline := time.After(2 * time.Second)
	for len(got) < 5 {
		select {
		case ev := <-sub.Events():
			if closed, ok := ev.(*mtypes.EventLeaseClosed); ok {
				got = append(got, closed.ID.DSeq)
			}
		case <-deadline:
			t.Fatalf("did not observe all 5 events, got %v so far", got)
		}
	}

	require.Equal(t, []uint64{101, 102, 103, 104, 105}, got)
}

// A publish failure means the bus is shutting down for good, so the run loop must
// return rather than keep burning retries against a bus that will never accept
// another event.
func TestEventsStopsWhenBusClosed(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })

	bus := pubsub.NewBus()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := newEvents(ctx, node, "test", bus, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})
	require.NoError(t, err)

	ev, ok := svc.(*events)
	require.True(t, ok)

	bus.Close()
	node.pushHeader(101)

	select {
	case <-ev.lc.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("run loop did not stop after the bus closed")
	}

	select {
	case <-ev.lc.ShuttingDown():
	default:
		t.Fatal("lifecycle never entered shutting-down, so WatchContext and the shutdown bridge stay parked forever")
	}
}

type pruningNode struct {
	fakeNode

	pmu      sync.Mutex
	earliest int64
	latest   int64
	ev       abci.Event
	missed   map[int64]struct{}
}

func newPruningNode(earliest, latest int64, ev abci.Event) *pruningNode {
	n := &pruningNode{earliest: earliest, latest: latest, ev: ev, missed: map[int64]struct{}{}}
	n.blocks = map[int64]*coretypes.ResultBlockResults{}
	n.blockErr = map[int64]error{}
	n.headerCh = make(chan coretypes.ResultEvent, 64)
	return n
}

func (p *pruningNode) window() (int64, int64) {
	p.pmu.Lock()
	defer p.pmu.Unlock()
	return p.earliest, p.latest
}

func (p *pruningNode) advance(by int64) {
	p.pmu.Lock()
	p.earliest += by
	p.latest += by
	p.pmu.Unlock()
}

func (p *pruningNode) Status(context.Context) (*coretypes.ResultStatus, error) {
	e, l := p.window()
	return &coretypes.ResultStatus{SyncInfo: coretypes.SyncInfo{
		LatestBlockHeight:   l,
		EarliestBlockHeight: e,
	}}, nil
}

func (p *pruningNode) BlockResults(_ context.Context, h *int64) (*coretypes.ResultBlockResults, error) {
	e, l := p.window()
	if *h < e {
		p.pmu.Lock()
		p.missed[*h] = struct{}{}
		p.pmu.Unlock()
		return nil, fmt.Errorf("height %d is not available, lowest height is %d", *h, e)
	}
	if *h > l {
		return nil, fmt.Errorf("height %d must be less than or equal to the current blockchain height %d", *h, l)
	}
	return &coretypes.ResultBlockResults{
		Height:     *h,
		TxsResults: []*abci.ExecTxResult{{Events: []abci.Event{p.ev}}},
	}, nil
}

func (p *pruningNode) missedHeights() int {
	p.pmu.Lock()
	defer p.pmu.Unlock()
	return len(p.missed)
}

func (p *pruningNode) Subscribe(context.Context, string, string, ...int) (<-chan coretypes.ResultEvent, error) {
	return p.headerCh, nil
}

// The node keeps pruning while the service waits out prunedCheckAfter, so a jump must
// be retried in the same pass and the pruned check must not rearm on every landing.
// Otherwise each jump lands behind the moving floor and pays a fresh prunedCheckAfter,
// and the cursor converges only by luck. Counting distinct heights that came back
// pruned makes that thrash visible without depending on wall-clock timing.
func TestEventsJumpConvergesInOnePassAgainstMovingFloor(t *testing.T) {
	const blockTime = 6 * time.Millisecond

	node := newPruningNode(100, 100, leaseClosedEvent(t, 1))
	sub := newTestService(t, node, tuning{
		watchdog:         blockTime,
		minRetryInterval: blockTime,
		prunedCheckAfter: 5 * blockTime,
	})

	node.advance(1000)

	stop := make(chan struct{})
	go func() {
		tick := time.NewTicker(blockTime)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tick.C:
				node.advance(1)
			}
		}
	}()
	defer close(stop)

	select {
	case <-sub.Events():
	case <-time.After(3 * time.Second):
		e, l := node.window()
		t.Fatalf("never converged against a floor moving 1 per %s, window now [%d,%d]", blockTime, e, l)
	}

	require.LessOrEqual(t, node.missedHeights(), 2,
		"each distinct pruned height past the first is a jump that landed behind the floor and cost another prunedCheckAfter")
}

// The steady-state driver is the block header, and CometBFT publishes it only after
// that block has already been committed and pruned for. The header path is therefore
// strictly tighter than the watchdog path and must converge on its own.
func TestEventsJumpConvergesWhenDrivenByHeaders(t *testing.T) {
	const blockTime = 6 * time.Millisecond

	node := newPruningNode(100, 100, leaseClosedEvent(t, 1))
	sub := newTestService(t, node, tuning{
		watchdog:         time.Hour,
		minRetryInterval: blockTime,
		prunedCheckAfter: 5 * blockTime,
	})

	node.advance(1000)

	stop := make(chan struct{})
	go func() {
		tick := time.NewTicker(blockTime)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tick.C:
				node.advance(1)
				_, l := node.window()
				node.pushHeader(l)
			}
		}
	}()
	defer close(stop)

	select {
	case <-sub.Events():
	case <-time.After(3 * time.Second):
		e, l := node.window()
		t.Fatalf("header-driven cursor never converged, window now [%d,%d]", e, l)
	}

	require.LessOrEqual(t, node.missedHeights(), 2, "the header path must not thrash against the moving floor either")
}
