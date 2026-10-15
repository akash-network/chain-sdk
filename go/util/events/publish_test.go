package events

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	cmtypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/log"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
	"pkg.akt.dev/go/util/pubsub"
)

// errTransientFetch stands in for a flaky-network error. It goes away on its own and
// carries no information the code reads, unlike a node's pruning response.
var errTransientFetch = errors.New("post failed: connection reset by peer")

// undecodableAttr is an attribute value no typed event decoder accepts.
const undecodableAttr = "{not json"

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
	blockDelay time.Duration
	// failFirst is how many more times a height fails before it is served, the way a
	// pooled backend that does not have the height yet answers.
	failFirst map[int64]int
	headerCh  chan coretypes.ResultEvent

	subscribes   int
	subscribeErr error
	// subscribeHang, when non-nil, makes Subscribe block until it is closed, ignoring
	// the context, the way a client stuck redialing would.
	subscribeHang chan struct{}
	// subscribeUntilCancel makes Subscribe block until its context ends, the way the
	// real client does while it is redialing.
	subscribeUntilCancel bool
	inflight             int
	maxInflight          int
	// calls records subscription calls in completion order.
	calls []string
}

// tuning is test shorthand for the three time options, so each test states all of
// them in one place.
type tuning struct {
	watchdog         time.Duration
	minRetryInterval time.Duration
	prunedCheckAfter time.Duration
}

func (t tuning) options() []Option {
	return []Option{
		WithWatchdogTimeout(t.watchdog),
		withMinRetryInterval(t.minRetryInterval),
		withPrunedCheckAfter(t.prunedCheckAfter),
	}
}

func newFakeNode() *fakeNode {
	return &fakeNode{
		height:     100,
		blocks:     map[int64]*coretypes.ResultBlockResults{},
		blockErr:   map[int64]error{},
		blockCalls: map[int64]int{},
		failFirst:  map[int64]int{},
		headerCh:   make(chan coretypes.ResultEvent, 64),
	}
}

func (f *fakeNode) set(mutate func(*fakeNode)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	mutate(f)
}

// pushHeader delivers a header on the channel the latest Subscribe returned.
func (f *fakeNode) pushHeader(h int64) {
	f.mu.Lock()
	ch := f.headerCh
	f.mu.Unlock()
	ch <- coretypes.ResultEvent{Data: cmtypes.EventDataNewBlockHeader{Header: cmtypes.Header{Height: h}}}
}

// dropSubscription kills the subscription the way a dead connection does: its channel
// stays open and goes silent, and only a fresh Subscribe receives later headers.
func (f *fakeNode) dropSubscription() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.headerCh = make(chan coretypes.ResultEvent, 64)
}

func (f *fakeNode) subscribeCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.subscribes
}

func (f *fakeNode) maxInflightSubscribes() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.maxInflight
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
	delay := f.blockDelay
	f.mu.Unlock()
	if delay > 0 {
		time.Sleep(delay)
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	f.blockCalls[*h]++
	if err := f.blockErr[*h]; err != nil {
		return nil, err
	}
	if f.failFirst[*h] > 0 {
		f.failFirst[*h]--
		return nil, errTransientFetch
	}
	if b, ok := f.blocks[*h]; ok {
		return b, nil
	}
	return &coretypes.ResultBlockResults{Height: *h}, nil
}

func (f *fakeNode) Subscribe(ctx context.Context, _, _ string, _ ...int) (<-chan coretypes.ResultEvent, error) {
	f.mu.Lock()
	f.subscribes++
	f.inflight++
	if f.inflight > f.maxInflight {
		f.maxInflight = f.inflight
	}
	ch, err, hang, untilCancel := f.headerCh, f.subscribeErr, f.subscribeHang, f.subscribeUntilCancel
	f.mu.Unlock()

	defer func() {
		f.mu.Lock()
		f.inflight--
		f.calls = append(f.calls, "subscribe returned")
		f.mu.Unlock()
	}()

	if hang != nil {
		<-hang
	}
	if untilCancel {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (f *fakeNode) Unsubscribe(context.Context, string, string) error { return nil }

func (f *fakeNode) UnsubscribeAll(context.Context, string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, "unsubscribe all")
	return nil
}

func (f *fakeNode) inflightSubscribes() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.inflight
}

func (f *fakeNode) callLog() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.calls...)
}

// requireNoMoreEvents fails if anything else is published within wait, which is how a
// duplicate would show up.
func requireNoMoreEvents(t *testing.T, sub pubsub.Subscriber, wait time.Duration) {
	t.Helper()
	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected extra event %T %+v", ev, ev)
	case <-time.After(wait):
	}
}

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

// dseqAt is the DSeq the tests give the lease-closed event of block h.
func dseqAt(h int64) uint64 {
	if h < 0 {
		panic(fmt.Sprintf("negative block height %d", h))
	}
	return uint64(h)
}

// withAttrs returns ev with attrs appended, the way the SDK tags the events it records.
func withAttrs(ev abci.Event, attrs ...abci.EventAttribute) abci.Event {
	ev.Attributes = append(append([]abci.EventAttribute{}, ev.Attributes...), attrs...)
	return ev
}

// collectLeaseClosed returns the DSeqs of the next n lease-closed events, in the order
// they were published.
func collectLeaseClosed(t *testing.T, sub pubsub.Subscriber, n int) []uint64 {
	t.Helper()
	return collectLeaseClosedWithin(t, sub, n, 2*time.Second)
}

func collectLeaseClosedWithin(t *testing.T, sub pubsub.Subscriber, n int, wait time.Duration) []uint64 {
	t.Helper()
	var got []uint64
	deadline := time.After(wait)
	for len(got) < n {
		select {
		case ev := <-sub.Events():
			if closed, ok := ev.(*mtypes.EventLeaseClosed); ok {
				got = append(got, closed.ID.DSeq)
			}
		case <-deadline:
			t.Fatalf("observed %d of %d events: %v", len(got), n, got)
		}
	}
	return got
}

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

type logEntry struct {
	level string
	msg   string
}

// recLogger records what the service logs, so tests can check that disconnects and
// reconnects are reported.
type recLogger struct {
	mu      *sync.Mutex
	entries *[]logEntry
}

var _ log.Logger = recLogger{}

func newRecLogger() recLogger {
	return recLogger{mu: &sync.Mutex{}, entries: &[]logEntry{}}
}

func (l recLogger) add(level, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	*l.entries = append(*l.entries, logEntry{level: level, msg: msg})
}

func (l recLogger) Debug(msg string, _ ...any) { l.add(levelDebug, msg) }
func (l recLogger) Info(msg string, _ ...any)  { l.add(levelInfo, msg) }
func (l recLogger) Warn(msg string, _ ...any)  { l.add(levelWarn, msg) }
func (l recLogger) Error(msg string, _ ...any) { l.add(levelError, msg) }
func (l recLogger) With(...any) log.Logger     { return l }
func (l recLogger) Impl() any                  { return l }

// count returns how many entries at level start with prefix.
func (l recLogger) count(level, prefix string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	n := 0
	for _, e := range *l.entries {
		if e.level == level && strings.HasPrefix(e.msg, prefix) {
			n++
		}
	}
	return n
}

func (l recLogger) has(level, prefix string) bool { return l.count(level, prefix) > 0 }

func newTestService(t *testing.T, node sdkclient.CometRPC, tune tuning, extra ...Option) pubsub.Subscriber {
	t.Helper()

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := NewEvents(ctx, node, "test", bus, append(tune.options(), extra...)...)
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

// A subscription whose connection died keeps its channel open and silent. The
// watchdog must re-subscribe and the service must follow headers on the new channel.
// Status stays at the start height, so polling alone cannot publish block 101.
func TestEventsResubscribesWhenSubscriptionGoesSilent(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 20 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.dropSubscription()
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
	require.GreaterOrEqual(t, node.subscribeCount(), 2, "the initial subscribe alone cannot see a header on the new channel")
}

// A re-subscribe that fails, as it does while the node is unreachable, is retried on
// the next watchdog expiry rather than given up on.
func TestEventsRetriesFailedResubscribe(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		f.subscribeErr = errors.New("client is not running")
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.dropSubscription()

	require.Eventually(t, func() bool { return node.subscribeCount() >= 3 }, 2*time.Second, time.Millisecond,
		"a failed re-subscribe must be retried on later expiries")
	require.Eventually(t, func() bool { return rec.count(levelWarn, "re-subscribe failed") >= 2 }, 2*time.Second, time.Millisecond,
		"every failed attempt is logged as a disconnect")

	node.set(func(f *fakeNode) { f.subscribeErr = nil })
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// Subscribe blocks for as long as the websocket client is redialing. The watchdog's
// catch-up must keep publishing meanwhile, and re-subscribes must not pile up.
func TestEventsHungResubscribeDoesNotStallCatchUp(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	hang := make(chan struct{})
	t.Cleanup(func() { close(hang) }) // runs before the service shuts down
	node.set(func(f *fakeNode) { f.subscribeHang = hang })
	node.dropSubscription()

	for h := int64(101); h <= 105; h++ {
		node.set(func(f *fakeNode) {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
			f.height = h
		})
		waitLeaseClosed(t, sub, dseqAt(h))
	}

	require.Equal(t, 1, node.maxInflightSubscribes(), "at most one re-subscribe may be in flight")
}

// Losing the subscription is logged at warn level once per outage, and every step of
// re-establishing it at info level, ending with the first header that proves it is
// back.
func TestEventsLogsSubscriptionLossAndRestore(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.dropSubscription()

	require.Eventually(t, func() bool { return rec.count(levelInfo, "re-subscribe sent") >= 2 }, 2*time.Second, time.Millisecond)
	require.Equal(t, 1, rec.count(levelWarn, "block subscription presumed lost"), "an outage is reported as lost once, not on every expiry")
	require.GreaterOrEqual(t, rec.count(levelInfo, "re-subscribing to block headers"), 2, "every attempt is logged")
	require.False(t, rec.has(levelInfo, "block subscription restored"), "a sent re-subscribe is not proof the subscription is back")

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)
	waitLeaseClosed(t, sub, 101)

	require.Equal(t, 1, rec.count(levelInfo, "block subscription restored"))
}

// A client that redials on its own within the watchdog window leaves no outage, only
// heights it never delivered. That gap is logged as a disconnect and back-filled.
func TestEventsLogsMissedHeaders(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 104; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
	})

	node.pushHeader(101)
	waitLeaseClosed(t, sub, 101)
	require.False(t, rec.has(levelWarn, "block subscription missed blocks"), "the first header after start is not a gap")

	node.pushHeader(104)
	require.Equal(t, []uint64{102, 103, 104}, collectLeaseClosed(t, sub, 3))
	require.Equal(t, 1, rec.count(levelWarn, "block subscription missed blocks"))
}

// A closed subscription channel is logged as a disconnect, once, and the watchdog's
// re-subscribe takes the service back to following headers.
func TestEventsLogsClosedSubscription(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 50 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
	})

	// A header first, so the watchdog is freshly armed when the channel closes.
	node.pushHeader(101)
	waitLeaseClosed(t, sub, 101)

	node.set(func(f *fakeNode) {
		close(f.headerCh)
		f.headerCh = make(chan coretypes.ResultEvent, 64)
	})
	require.Eventually(t, func() bool { return rec.count(levelInfo, "re-subscribe sent") >= 1 }, 2*time.Second, time.Millisecond)

	node.pushHeader(102)
	waitLeaseClosed(t, sub, 102)

	require.Equal(t, 1, rec.count(levelWarn, "block subscription channel closed"))
	require.Zero(t, rec.count(levelWarn, "block subscription presumed lost"), "the close already reported the outage")
	require.Equal(t, 1, rec.count(levelInfo, "block subscription restored"))
}

// A height that fails once is retried on its own timer. Nothing else arrives to
// trigger it: no further header, and a watchdog too long to matter. Waiting for an
// outside trigger would hold the height, and every event after it, indefinitely.
func TestEventsRetriesParkedHeightOnItsOwn(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 20 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		f.failFirst[101] = 1
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	require.Equal(t, []uint64{101}, collectLeaseClosed(t, sub, 1))
	require.Equal(t, 2, node.callsAt(101), "one failed attempt, then one retry")
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
}

// While the subscription is down the watchdog is the only outside trigger, a full
// window apart. A back-fill whose heights each fail once, as behind a lagging or
// rate-limiting proxy, must still go at the retry pace, not one height per window.
func TestEventsBackfillRetriesAtRetryPaceWhileSubscriptionDown(t *testing.T) {
	const (
		watchdog = 200 * time.Millisecond
		retry    = 10 * time.Millisecond
		heights  = 10
	)

	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: watchdog, minRetryInterval: retry, prunedCheckAfter: time.Hour})

	// The chain moves on after the service started from height 100.
	node.set(func(f *fakeNode) {
		for h := int64(101); h < 101+heights; h++ {
			f.failFirst[h] = 1
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
		f.height = 100 + heights
	})

	// No header ever arrives. The first watchdog tick starts the back-fill; at one
	// height per watchdog window it would take heights*watchdog = 2s.
	start := time.Now()
	want := make([]uint64, 0, heights)
	for h := int64(101); h < 101+heights; h++ {
		want = append(want, dseqAt(h))
	}
	require.Equal(t, want, collectLeaseClosedWithin(t, sub, heights, 3*time.Second))
	require.Less(t, time.Since(start), watchdog+heights*watchdog/2, "the back-fill went at the watchdog pace, not the retry pace")
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
}

// A fetch failure is a disconnect, logged at warn level the first time; the retries
// that follow log at debug level. Only a stall that outlasts prunedCheckAfter is
// logged at error level, and only once.
func TestEventsLogsFetchFailureWarnThenErrorOnce(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	require.Eventually(t, func() bool { return rec.has(levelWarn, "block fetch failed, cursor parked") }, time.Second, time.Millisecond)
	require.False(t, rec.has(levelError, "block fetch"), "a first failure is routine behind pooled backends, not an error")

	require.Eventually(t, func() bool { return rec.has(levelError, "block fetch keeps failing, cursor stalled") }, 2*time.Second, time.Millisecond)
	time.Sleep(100 * time.Millisecond) // many more retries
	require.Equal(t, 1, rec.count(levelError, "block fetch keeps failing, cursor stalled"), "the stall is escalated once")
	require.Equal(t, 1, rec.count(levelWarn, "block fetch failed, cursor parked"), "the first failure is reported once")
	require.Greater(t, rec.count(levelDebug, "block fetch still failing"), 1, "retries in between are debug")

	node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	waitLeaseClosed(t, sub, 101)
	require.True(t, rec.has(levelInfo, "block fetch recovered, cursor resumed"))
}

// A height parked behind a transient failure is retried on every watchdog tick, even
// while Status keeps failing and no further header arrives, because the height a
// header already announced is remembered.
func TestEventsRetriesKnownHeightWhileStatusFails(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		for h := int64(101); h <= 103; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
		f.statusErr = errors.New("rpc unavailable")
	})
	node.pushHeader(103)

	require.Eventually(t, func() bool { return node.callsAt(101) >= 1 }, 2*time.Second, time.Millisecond)
	node.set(func(f *fakeNode) { delete(f.blockErr, 101) })

	require.Equal(t, []uint64{101, 102, 103}, collectLeaseClosed(t, sub, 3))
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
}

// Only a block above the last one seen is new. Repeated or older headers, such as a
// lagging backend behind a load balancer sends, must neither hold the watchdog off nor
// announce the subscription restored.
func TestEventsStaleHeadersDoNotHoldOffWatchdog(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 50 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
	})
	node.pushHeader(101)
	waitLeaseClosed(t, sub, 101)

	stop, done := make(chan struct{}), make(chan struct{})
	go func() {
		defer close(done)
		tick := time.NewTicker(5 * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-stop:
				return
			case <-tick.C:
				node.pushHeader(101)
				node.pushHeader(100)
			}
		}
	}()

	require.Eventually(t, func() bool { return rec.has(levelWarn, "block subscription presumed lost") }, time.Second, time.Millisecond,
		"stale headers must not hold the watchdog off")
	time.Sleep(30 * time.Millisecond) // more stale headers after the outage is declared
	close(stop)
	<-done
	require.Zero(t, rec.count(levelInfo, "block subscription restored"), "a stale header is not proof the subscription is back")

	node.pushHeader(102)
	waitLeaseClosed(t, sub, 102)
	require.Equal(t, 1, rec.count(levelInfo, "block subscription restored"))
}

// A catch-up that outlasts the watchdog window leaves the headers that arrived
// meanwhile queued. They prove the subscription alive, so the expired watchdog must not
// declare an outage. When both are ready the run loop's select picks at random, so one
// round would catch a regression only half the time; several rounds make it near-certain.
func TestEventsQueuedHeadersAreNotAnOutage(t *testing.T) {
	const (
		rounds   = 4
		perRound = 8 // blocks per catch-up, at blockDelay each: 96ms, past the 80ms watchdog
		queued   = 2 // headers that arrive while it runs
	)

	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 80 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	last := int64(100 + rounds*(perRound+queued))
	node.set(func(f *fakeNode) {
		f.blockDelay = 12 * time.Millisecond
		for h := int64(101); h <= last; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
	})

	from := int64(101)
	for r := 0; r < rounds; r++ {
		top := from + perRound - 1
		for h := top; h <= top+queued; h++ {
			node.pushHeader(h)
		}

		var want []uint64
		for h := from; h <= top+queued; h++ {
			want = append(want, dseqAt(h))
		}
		require.Equal(t, want, collectLeaseClosed(t, sub, len(want)), "round %d", r)
		from = top + queued + 1
	}

	require.Zero(t, rec.count(levelWarn, "block subscription presumed lost"), "queued headers are not an outage")
}

// A recognized block-level event that cannot be decoded fails its whole height as a
// decode error, so the cursor parks without publishing any of the height.
func TestBlockEventsReportsUndecodableBlockLevelEvent(t *testing.T) {
	node := newFakeNode()

	bad := withAttrs(abci.Event{Type: leaseClosedEvent(t, 2).Type, Attributes: []abci.EventAttribute{{Key: "id", Value: undecodableAttr}}},
		abci.EventAttribute{Key: modeKey, Value: modeEndBlock})
	blk := txBlock(101, leaseClosedEvent(t, 1))
	blk.FinalizeBlockEvents = []abci.Event{bad}
	node.set(func(f *fakeNode) { f.blocks[101] = blk })

	e := &events{client: node}
	evs, err := e.blockEvents(context.Background(), 101)

	var derr *decodeError
	require.ErrorAs(t, err, &derr)
	require.Equal(t, int64(101), derr.height)
	require.Nil(t, evs, "nothing of a height that fails to decode may be returned")
}

// WithStartHeight makes the service publish from the block after the one the caller
// accounted for, without asking the node where the chain is.
func TestEventsStartsAfterStartHeight(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) {
		f.statusErr = errors.New("rpc unavailable")
		for h := int64(90); h <= 93; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
	})

	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithStartHeight(90))
	node.pushHeader(93)

	require.Equal(t, []uint64{91, 92, 93}, collectLeaseClosed(t, sub, 3), "block 90 is the caller's, everything after it is published")
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
	_, ok, err := processEvent(leaseClosedEvent(t, 1))
	require.NoError(t, err)
	require.True(t, ok, "a lease-closed event must be recognized")

	other := abci.Event{Type: "coin_received", Attributes: []abci.EventAttribute{{Key: "amount", Value: "1uakt"}}}
	_, ok, err = processEvent(other)
	require.NoError(t, err, "an event the service does not publish is not a decode failure")
	require.False(t, ok, "a non-Akash event must be ignored")
}

// The SDK appends bookkeeping attributes to the events it records: msg_index to every
// message event, authz_msg_index to events of messages run through authz MsgExec, and
// mode to every BeginBlock and EndBlock event. None of them may stop a recognized
// event from being decoded. mode in particular is a bare word, not JSON.
func TestProcessEventAcceptsSDKBookkeepingAttributes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		attrs []abci.EventAttribute
	}{
		{name: "tx", attrs: []abci.EventAttribute{{Key: "msg_index", Value: "0"}}},
		{name: "authz", attrs: []abci.EventAttribute{{Key: "authz_msg_index", Value: "1"}, {Key: "msg_index", Value: "0"}}},
		{name: "begin block", attrs: []abci.EventAttribute{{Key: modeKey, Value: modeBeginBlock}}},
		{name: "end block", attrs: []abci.EventAttribute{{Key: modeKey, Value: modeEndBlock}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ev := withAttrs(leaseClosedEvent(t, 7), tc.attrs...)

			pev, ok, err := processEvent(ev)
			require.NoError(t, err)
			require.True(t, ok)
			closed, isClosed := pev.(*mtypes.EventLeaseClosed)
			require.True(t, isClosed)
			require.Equal(t, uint64(7), closed.ID.DSeq)
		})
	}
}

// A recognized event that cannot be decoded is reported, never treated as an event
// the service does not publish.
func TestProcessEventReportsUndecodableRecognizedEvent(t *testing.T) {
	ev := leaseClosedEvent(t, 1)
	ev.Attributes = []abci.EventAttribute{{Key: "id", Value: undecodableAttr}}

	_, ok, err := processEvent(ev)
	require.Error(t, err)
	require.False(t, ok)
}

// A height that holds a recognized event the service cannot decode parks the cursor.
// Nothing of that height is published, not even the valid events ahead of the bad
// one, and nothing after it. Once the height can be read, every event goes out once.
func TestEventsParksOnUndecodableEvent(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	bad := leaseClosedEvent(t, 2)
	bad.Attributes = []abci.EventAttribute{{Key: "id", Value: undecodableAttr}}
	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 1), bad)
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
		f.height = 102
	})
	node.pushHeader(102)

	select {
	case ev := <-sub.Events():
		t.Fatalf("published %T from or past a height whose event could not be decoded", ev)
	case <-time.After(200 * time.Millisecond):
	}

	require.Greater(t, node.callsAt(101), 1, "the undecodable height must be retried, not abandoned")
	require.Zero(t, node.callsAt(102), "the cursor must not move past the undecodable height")
	require.True(t, rec.has(levelError, "block decode failed, cursor parked"), "a decode failure must be named as such")
	require.False(t, rec.has(levelError, "block fetch failed"), "a decode failure is not a network problem")

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 1), leaseClosedEvent(t, 2)) })

	require.Equal(t, []uint64{1, 2, 102}, collectLeaseClosed(t, sub, 3))
	requireNoMoreEvents(t, sub, 100*time.Millisecond)
}

// Headers repeat and arrive out of order across re-subscribes, and the watchdog's
// catch-up covers the same heights again. Each height is still published exactly once,
// in order.
func TestEventsPublishesEachHeightOnce(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 104; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
		f.height = 104
	})
	for _, h := range []int64{101, 101, 100, 103, 102, 103, 104, 101} {
		node.pushHeader(h)
	}

	require.Equal(t, []uint64{101, 102, 103, 104}, collectLeaseClosed(t, sub, 4))
	requireNoMoreEvents(t, sub, 100*time.Millisecond)
}

// Shutting down while a re-subscribe is stuck redialing cancels it and waits for it to
// return before unsubscribing, so it cannot register the query again behind
// UnsubscribeAll.
func TestEventsShutdownCancelsInFlightResubscribe(t *testing.T) {
	node := newFakeNode()
	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	// A long watchdog keeps the re-subscribe in flight for a full second, far longer
	// than a cancelled one takes to return.
	svc, err := NewEvents(ctx, node, "test", bus, tuning{watchdog: time.Second, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}.options()...)
	require.NoError(t, err)

	node.set(func(f *fakeNode) { f.subscribeUntilCancel = true })
	require.Eventually(t, func() bool { return node.inflightSubscribes() == 1 }, 3*time.Second, time.Millisecond)

	start := time.Now()
	svc.Shutdown()
	require.Less(t, time.Since(start), 500*time.Millisecond, "shutdown must cancel the in-flight re-subscribe, not wait out its timeout")

	calls := node.callLog()
	require.GreaterOrEqual(t, len(calls), 2)
	require.Equal(t, []string{"subscribe returned", "unsubscribe all"}, calls[len(calls)-2:],
		"the re-subscribe must return before the service unsubscribes")
	require.Zero(t, node.inflightSubscribes())
}

// x/gov's EndBlocker runs the messages of passed proposals and re-emits their events at
// block level, so the service must publish block-level events too, in execution order:
// PreBlock, then BeginBlock, then each tx, then EndBlock.
func TestEventsPublishesBlockEventsInExecutionOrder(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	blk := txBlock(101, leaseClosedEvent(t, 3), leaseClosedEvent(t, 4))
	blk.FinalizeBlockEvents = []abci.Event{
		leaseClosedEvent(t, 1), // PreBlock events carry no mode tag
		withAttrs(leaseClosedEvent(t, 2), abci.EventAttribute{Key: modeKey, Value: modeBeginBlock}),
		withAttrs(leaseClosedEvent(t, 5), abci.EventAttribute{Key: modeKey, Value: modeEndBlock}),
	}
	node.set(func(f *fakeNode) { f.blocks[101] = blk })
	node.pushHeader(101)

	require.Equal(t, []uint64{1, 2, 3, 4, 5}, collectLeaseClosed(t, sub, 5))
}

// The cursor never moves past a height the node has pruned, because publishing what
// follows would mean the pruned height's events were skipped. The stall is explained
// once in the log, not on every retry.
func TestEventsParksOnPrunedHeight(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 3 * time.Millisecond, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond}, WithLogger(rec))

	node.set(func(f *fakeNode) {
		for h := int64(101); h < 200; h++ {
			f.blockErr[h] = errTransientFetch
		}
		f.blocks[200] = txBlock(200, leaseClosedEvent(t, 200))
		f.earliest = 200
		f.height = 200
	})

	select {
	case ev := <-sub.Events():
		t.Fatalf("published %T past a pruned height", ev)
	case <-time.After(200 * time.Millisecond):
	}

	require.Greater(t, node.callsAt(101), 1, "the pruned height must keep being retried")
	require.Zero(t, node.callsAt(102), "the cursor must not move past the pruned height")
	require.Equal(t, 1, rec.count(levelWarn, "node has pruned the parked height"), "the pruned stall is explained once")
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
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
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

	svc, err := NewEvents(ctx, node, "test", bus, tuning{watchdog: time.Hour, minRetryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}.options()...)
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

// A node keeps pruning while the service waits, and every block above the floor is
// fetchable. The cursor must still stay on the first pruned height, driven by both
// headers and watchdog ticks, rather than chase the floor. The only height ever asked
// for below the floor is the one the cursor is parked on.
func TestEventsStaysParkedBehindMovingPruneFloor(t *testing.T) {
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
				_, l := node.window()
				node.pushHeader(l)
			}
		}
	}()
	defer close(stop)

	select {
	case ev := <-sub.Events():
		t.Fatalf("published %T past a pruned height", ev)
	case <-time.After(300 * time.Millisecond):
	}

	require.Equal(t, 1, node.missedHeights(), "the cursor must stay on the first pruned height")
}
