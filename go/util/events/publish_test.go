package events

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/cometbft/cometbft/abci/types"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
	"pkg.akt.dev/go/util/pubsub"
)

// errTransientFetch stands in for a flaky-network error. It goes away on its own and
// carries no information the code reads, unlike a node's pruning response.
var errTransientFetch = errors.New("post failed: connection reset by peer")

// undecodableAttr is an attribute value no typed event decoder accepts.
const undecodableAttr = "{not json"

// testNode is a fake node: the HTTP calls the service makes, and its websocket.
type testNode interface {
	queryNode
	dial(ctx context.Context) (headerStream, error)
}

// fakeNode answers the service's HTTP calls and serves its websocket header streams.
type fakeNode struct {
	mu        sync.Mutex
	height    int64
	earliest  int64
	statusErr error
	// statusHang makes Status block until its context ends.
	statusHang bool
	// statusDelay makes Status sleep before answering, ignoring its context, the way a
	// request already on the wire does.
	statusDelay time.Duration
	// statusEntered, when non-nil, is signalled every time Status is called.
	statusEntered chan struct{}
	statusCalls   int
	blocks        map[int64]*coretypes.ResultBlockResults
	blockErr      map[int64]error
	blockCalls    map[int64]int
	// gate, when non-nil, holds every fetch until it is closed or the fetch's context
	// ends.
	gate chan struct{}
	// slow delays the fetch of one height.
	slow map[int64]time.Duration
	// failFirst is how many more times a height fails before it is served, the way a
	// pooled backend that does not have the height yet answers.
	failFirst map[int64]int
	// hangFirst is how many more times a fetch of a height blocks until its context
	// ends, the way a request stuck on a dead connection or queued behind a full
	// connection pool does.
	hangFirst map[int64]int
	// serveOnce is answered, once, to the next fetch of a height instead of its block,
	// the way a misrouted or cached response for another height would be.
	serveOnce map[int64]*coretypes.ResultBlockResults
	// fetchInflight counts fetches in progress; maxFetchInflight is its peak.
	fetchInflight    int
	maxFetchInflight int

	// dialErr fails every dial. dialHang makes a dial block until its context ends,
	// the way a dial to a node that never answers the handshake does.
	dialErr  error
	dialHang bool
	dialed   int
	// streams are the header streams handed out, in order.
	streams []*fakeStream
}

// fakeStream is a header stream the test drives.
type fakeStream struct {
	mu     sync.Mutex
	latest int64
	err    error
	ended  bool
	closed bool
	notify chan struct{}
	done   chan struct{}
}

func newFakeStream() *fakeStream {
	return &fakeStream{notify: make(chan struct{}, 1), done: make(chan struct{})}
}

// push delivers a header. Like the real stream, only a header above the highest seen
// is news.
func (s *fakeStream) push(h int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h <= s.latest {
		return
	}
	s.latest = h
	select {
	case s.notify <- struct{}{}:
	default:
	}
}

// end ends the stream with err, the way a failed connection or subscription does.
func (s *fakeStream) end(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.ended {
		return
	}
	s.ended, s.err = true, err
	close(s.done)
}

func (s *fakeStream) Notify() <-chan struct{} { return s.notify }
func (s *fakeStream) Done() <-chan struct{}   { return s.done }

func (s *fakeStream) Latest() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.latest
}

func (s *fakeStream) Err() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

func (s *fakeStream) Close() {
	s.end(errStreamClosed)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.closed = true
}

func (s *fakeStream) isClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}

// tuning is test shorthand for the time options, so each test states them in one
// place. fetchTimeout keeps its default when left zero.
type tuning struct {
	watchdog         time.Duration
	retryInterval    time.Duration
	prunedCheckAfter time.Duration
	fetchTimeout     time.Duration
}

func (t tuning) options() []Option {
	opts := []Option{
		WithWatchdogTimeout(t.watchdog),
		WithRetryInterval(t.retryInterval),
		withPrunedCheckAfter(t.prunedCheckAfter),
	}
	if t.fetchTimeout > 0 {
		opts = append(opts, WithFetchTimeout(t.fetchTimeout))
	}
	return opts
}

func newFakeNode() *fakeNode {
	return &fakeNode{
		height:     100,
		blocks:     map[int64]*coretypes.ResultBlockResults{},
		blockErr:   map[int64]error{},
		blockCalls: map[int64]int{},
		slow:       map[int64]time.Duration{},
		failFirst:  map[int64]int{},
		hangFirst:  map[int64]int{},
		serveOnce:  map[int64]*coretypes.ResultBlockResults{},
	}
}

func (f *fakeNode) set(mutate func(*fakeNode)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	mutate(f)
}

func (f *fakeNode) dial(ctx context.Context) (headerStream, error) {
	f.mu.Lock()
	err, hang := f.dialErr, f.dialHang
	if hang || err != nil {
		f.dialed++
		f.mu.Unlock()
		if hang {
			<-ctx.Done()
			return nil, ctx.Err()
		}
		return nil, err
	}

	// A successful dial is counted together with its stream, so a test that waits for
	// the nth dial finds the nth stream in place.
	s := newFakeStream()
	f.dialed++
	f.streams = append(f.streams, s)
	f.mu.Unlock()

	return s, nil
}

func (f *fakeNode) dials() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.dialed
}

// waitDials waits until the service has dialed n times.
func (f *fakeNode) waitDials(t *testing.T, n int) {
	t.Helper()
	require.Eventually(t, func() bool { return f.dials() >= n }, 2*time.Second, time.Millisecond, "waiting for dial %d", n)
}

// current is the stream the latest successful dial returned, or nil.
func (f *fakeNode) current() *fakeStream {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.streams) == 0 {
		return nil
	}
	return f.streams[len(f.streams)-1]
}

// pushHeader delivers a header on the latest stream, waiting briefly for the first dial,
// which the service makes as it starts.
func (f *fakeNode) pushHeader(h int64) {
	deadline := time.Now().Add(2 * time.Second)
	for {
		if s := f.current(); s != nil {
			s.push(h)
			return
		}
		if time.Now().After(deadline) {
			return
		}
		time.Sleep(time.Millisecond)
	}
}

// endStream ends the latest stream with err, the way a dead connection does.
func (f *fakeNode) endStream(err error) {
	if s := f.current(); s != nil {
		s.end(err)
	}
}

func (f *fakeNode) allStreamsClosed() bool {
	f.mu.Lock()
	streams := append([]*fakeStream(nil), f.streams...)
	f.mu.Unlock()
	for _, s := range streams {
		if !s.isClosed() {
			return false
		}
	}
	return true
}

func (f *fakeNode) callsAt(h int64) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.blockCalls[h]
}

func (f *fakeNode) statusCallCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.statusCalls
}

func (f *fakeNode) fetchesInflight() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.fetchInflight
}

func (f *fakeNode) peakFetchesInflight() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.maxFetchInflight
}

// requireSignal waits for a signal on ch, failing the test after a generous limit.
func requireSignal(t *testing.T, ch <-chan struct{}, what string) {
	t.Helper()

	select {
	case <-ch:
	case <-time.After(2 * time.Second):
		t.Fatalf("timed out waiting for %s", what)
	}
}

func (f *fakeNode) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	f.mu.Lock()
	f.statusCalls++
	hang, delay, entered := f.statusHang, f.statusDelay, f.statusEntered
	f.mu.Unlock()

	if entered != nil {
		select {
		case entered <- struct{}{}:
		default:
		}
	}
	if hang {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	if delay > 0 {
		time.Sleep(delay)
	}

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

func (f *fakeNode) BlockResults(ctx context.Context, h *int64) (*coretypes.ResultBlockResults, error) {
	f.mu.Lock()
	f.fetchInflight++
	if f.fetchInflight > f.maxFetchInflight {
		f.maxFetchInflight = f.fetchInflight
	}
	delay, gate := f.slow[*h], f.gate
	hang := f.hangFirst[*h] > 0
	if hang {
		f.hangFirst[*h]--
		f.blockCalls[*h]++
	}
	f.mu.Unlock()
	defer func() {
		f.mu.Lock()
		f.fetchInflight--
		f.mu.Unlock()
	}()

	if hang {
		<-ctx.Done()
		return nil, ctx.Err()
	}
	if gate != nil {
		select {
		case <-gate:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
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
	if b, ok := f.serveOnce[*h]; ok {
		delete(f.serveOnce, *h)
		return b, nil
	}
	if b, ok := f.blocks[*h]; ok {
		return b, nil
	}
	return &coretypes.ResultBlockResults{Height: *h}, nil
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

// dseqsThrough returns the DSeqs the tests give blocks 101, the first after the start
// height, through last.
func dseqsThrough(last int64) []uint64 {
	var out []uint64
	for h := int64(101); h <= last; h++ {
		out = append(out, dseqAt(h))
	}
	return out
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

// newEventsOn starts the service on node, applying opts the way NewEvents does.
func newEventsOn(ctx context.Context, node testNode, bus pubsub.Bus, opts ...Option) (*events, error) {
	o, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	return newEvents(ctx, node, node.dial, "test", bus, o)
}

func newTestService(t *testing.T, node testNode, tune tuning, extra ...Option) pubsub.Subscriber {
	t.Helper()

	_, sub := startTestService(t, node, tune, extra...)
	return sub
}

// startTestService is newTestService for tests that also watch the service itself.
func startTestService(t *testing.T, node testNode, tune tuning, extra ...Option) (*events, pubsub.Subscriber) {
	t.Helper()

	bus := pubsub.NewBus()
	sub, err := bus.Subscribe()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	svc, err := newEventsOn(ctx, node, bus, append(tune.options(), extra...)...)
	require.NoError(t, err)
	t.Cleanup(func() {
		svc.Shutdown()
		bus.Close()
	})

	return svc, sub
}

// requireStopped waits for svc to stop on its own and returns the error it stopped on.
func requireStopped(t *testing.T, svc Service, within time.Duration) error {
	t.Helper()

	select {
	case <-svc.Done():
		return svc.Err()
	case <-time.After(within):
		t.Fatalf("service still running after %s", within)
		return nil
	}
}

// withAttr returns ev with the value of attribute key replaced, or the attribute
// added when ev has none.
func withAttr(ev abci.Event, key, value string) abci.Event {
	out := abci.Event{Type: ev.Type, Attributes: make([]abci.EventAttribute, 0, len(ev.Attributes)+1)}
	found := false
	for _, attr := range ev.Attributes {
		if attr.Key == key {
			attr.Value = value
			found = true
		}
		out.Attributes = append(out.Attributes, attr)
	}
	if !found {
		out.Attributes = append(out.Attributes, abci.EventAttribute{Key: key, Value: value})
	}
	return out
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
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog idle, header drives processing

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// A header that jumps several heights (as after a dropped and re-established
// subscription) back-fills every block in between, in order.
func TestEventsBackfillsGapBetweenHeaders(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog idle, header drives processing

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

// When no header arrives, the watchdog asks the node how far the chain has got and
// catches up, so events keep flowing.
func TestEventsFetchFallbackWhenSubscriptionQuiet(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond})

	// No header pushed, the subscription is silent. A new block appears on chain.
	node.set(func(f *fakeNode) {
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.height = 101
	})

	waitLeaseClosed(t, sub, 101)
}

// A stream that ends is redialed at once; the service follows headers on the new one.
func TestEventsRedialsWhenStreamEnds(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: time.Hour, prunedCheckAfter: time.Hour})

	node.waitDials(t, 1)
	node.endStream(io.EOF)
	node.waitDials(t, 2) // the retry interval is an hour: only an immediate redial gets here

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// After the immediate redial, redials wait the retry interval: a node that keeps
// refusing is asked about once per interval, forever, never in a busy loop.
func TestEventsPacesRedialsAfterTheFirst(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 50 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.waitDials(t, 1)
	node.set(func(f *fakeNode) { f.dialErr = errors.New("connection refused") })
	node.endStream(io.EOF)

	time.Sleep(500 * time.Millisecond)
	dials := node.dials()
	require.GreaterOrEqual(t, dials, 5, "redials go on")
	require.LessOrEqual(t, dials, 13, "about one redial per retry interval, not a busy loop")
	require.GreaterOrEqual(t, rec.count(levelWarn, "block subscription dial failed"), 4)
}

// A live stream that delivers no block is a slow chain, not a lost subscription: the
// service polls the node but neither redials nor reports an outage.
func TestEventsSlowChainDoesNotRedial(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	newTestService(t, node, tuning{watchdog: 20 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.waitDials(t, 1)
	before := node.statusCallCount()
	time.Sleep(200 * time.Millisecond)

	require.Equal(t, 1, node.dials())
	require.Zero(t, rec.count(levelWarn, "block subscription lost"))
	require.GreaterOrEqual(t, node.statusCallCount()-before, 3, "the watchdog polls the node instead")
	require.True(t, rec.has(levelDebug, "no new block within watchdog window"))
}

// Losing the subscription is logged at warn level once per outage, however many
// redials it takes, and every dial at info level; the first header after it reports
// the subscription restored, once.
func TestEventsLogsSubscriptionLossOncePerOutage(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	for i := 1; i <= 3; i++ {
		node.waitDials(t, i)
		node.endStream(io.EOF)
	}
	node.waitDials(t, 4)

	require.Equal(t, 1, rec.count(levelWarn, "block subscription lost"))
	require.GreaterOrEqual(t, rec.count(levelInfo, "dialing block subscription"), 4)
	require.Zero(t, rec.count(levelInfo, "block subscription restored"), "a dial is not proof the subscription is back")

	node.set(func(f *fakeNode) { f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101)) })
	node.pushHeader(101)
	waitLeaseClosed(t, sub, 101)

	require.Equal(t, 1, rec.count(levelInfo, "block subscription restored"))
}

// Each outage is its own: a loss after the subscription was restored is reported again,
// redialed at once again, and reported restored again by its first new header.
func TestEventsLogsEachOutage(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 102; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
	})
	// An hour's retry interval: only an outage's immediate redial can reconnect within
	// the test.
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: time.Hour, prunedCheckAfter: time.Hour}, WithLogger(rec))

	for i, h := range []int64{101, 102} {
		outage := i + 1
		node.waitDials(t, outage)
		node.endStream(io.EOF)
		node.waitDials(t, outage+1)
		node.pushHeader(h)
		waitLeaseClosed(t, sub, dseqAt(h))

		require.Equal(t, outage, rec.count(levelWarn, "block subscription lost"), "outage %d", outage)
		require.Equal(t, outage, rec.count(levelInfo, "block subscription restored"), "outage %d", outage)
	}
}

// A client that redials on its own within the watchdog window leaves no outage, only
// heights it never delivered. That gap is logged as a disconnect and back-filled.
func TestEventsLogsMissedHeaders(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

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

// A height that fails once is retried on its own timer. Nothing else arrives to
// trigger it: no further header, and a watchdog too long to matter. Waiting for an
// outside trigger would hold the height, and every event after it, indefinitely.
func TestEventsRetriesParkedHeightOnItsOwn(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 20 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		f.failFirst[101] = 1
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	require.Equal(t, []uint64{101}, collectLeaseClosed(t, sub, 1))
	require.Equal(t, 2, node.callsAt(101), "one failed attempt, then one retry")
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
}

// Without headers the watchdog is the only outside trigger, a full window apart. A
// back-fill whose heights each fail once, as behind a lagging or rate-limiting proxy,
// must still go at the retry pace, not one height per window.
func TestEventsBackfillRetriesAtRetryPaceWhileSubscriptionDown(t *testing.T) {
	const (
		watchdog = 200 * time.Millisecond
		retry    = 10 * time.Millisecond
		heights  = 10
	)

	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: watchdog, retryInterval: retry, prunedCheckAfter: time.Hour})

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
	require.Equal(t, dseqsThrough(100+heights), collectLeaseClosedWithin(t, sub, heights, 3*time.Second))
	require.Less(t, time.Since(start), watchdog+heights*watchdog/2, "the back-fill went at the watchdog pace, not the retry pace")
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
}

// A fetch failure is a disconnect, logged at warn level the first time; the retries
// that follow log at debug level. Only a stall that outlasts prunedCheckAfter is
// logged at error level, and only once.
func TestEventsLogsFetchFailureWarnThenErrorOnce(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}, WithLogger(rec))

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

// A height parked behind a transient failure is retried even while Status keeps
// failing and no further header arrives, because the height a header already announced
// is remembered.
func TestEventsRetriesKnownHeightWhileStatusFails(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

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

// Prefetching never has more than the window's worth of fetches in flight, and
// publishes strictly in order however the fetches complete.
func TestEventsPrefetchIsBoundedByWindow(t *testing.T) {
	node := newFakeNode()
	gate := make(chan struct{})
	node.set(func(f *fakeNode) {
		f.gate = gate
		for h := int64(101); h <= 130; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
		f.height = 130
	})
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour},
		WithPrefetch(3), WithStartHeight(100))

	// While the node holds every fetch, the service fills the window and goes no
	// further, however often the watchdog has it poll the node meanwhile.
	require.Eventually(t, func() bool { return node.fetchesInflight() == 3 }, 2*time.Second, time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	require.Equal(t, 3, node.peakFetchesInflight())

	// Released together, the fetches complete in any order.
	close(gate)
	require.Equal(t, dseqsThrough(130), collectLeaseClosedWithin(t, sub, 30, 3*time.Second))
	require.Equal(t, 3, node.peakFetchesInflight(), "never more than the window in flight")
}

// A height ahead of the cursor that fails is fetched again once the cursor reaches
// it, right away and without parking the cursor on it.
func TestEventsPrefetchFailureAboveCursorWaitsForCursor(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 104; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
		f.slow[101] = 50 * time.Millisecond // 103 fails while the cursor is still at 101
		f.failFirst[103] = 1
	})
	// An hour's retry interval: a parked 103 would never come back within the test.
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: time.Hour, prunedCheckAfter: time.Hour}, WithPrefetch(4), WithLogger(rec))
	node.pushHeader(104)

	require.Equal(t, dseqsThrough(104), collectLeaseClosed(t, sub, 4))
	require.Equal(t, 2, node.callsAt(103))
	require.Zero(t, rec.count(levelWarn, "block fetch failed, cursor parked"), "the cursor never parked on 103")
}

// Headers are handled while fetches hang: the loop never waits on the node.
func TestEventsHeadersHandledDuringBackfill(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 104; h++ {
			f.hangFirst[h] = 1
		}
	})
	newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: time.Hour, prunedCheckAfter: time.Hour, fetchTimeout: time.Minute}, WithLogger(rec))

	node.pushHeader(104)
	require.Eventually(t, func() bool { return node.fetchesInflight() == 4 }, time.Second, time.Millisecond)

	node.pushHeader(150)
	require.Eventually(t, func() bool { return rec.has(levelWarn, "block subscription missed blocks") }, 100*time.Millisecond, time.Millisecond)
}

// Headers repeat across a redial. Each height is still published once, and the
// subscription is reported restored once.
func TestEventsRepeatedHeadersAcrossRedialsPublishOnce(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	node.set(func(f *fakeNode) {
		for h := int64(101); h <= 103; h++ {
			f.blocks[h] = txBlock(h, leaseClosedEvent(t, dseqAt(h)))
		}
	})
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

	node.pushHeader(101)
	node.pushHeader(102)
	require.Equal(t, []uint64{101, 102}, collectLeaseClosed(t, sub, 2))

	node.endStream(io.EOF)
	node.waitDials(t, 2)
	node.pushHeader(101)
	node.pushHeader(102)
	node.pushHeader(103)

	require.Equal(t, []uint64{103}, collectLeaseClosed(t, sub, 1))
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
	require.Equal(t, 1, rec.count(levelInfo, "block subscription restored"))
}

// A recognized event the service cannot decode is skipped and logged once. Retrying its
// height cannot change the outcome, so the rest of the height and every later height
// are published as usual. A newer chain emitting an enum value this build does not
// know is the realistic case.
func TestEventsSkipsUndecodableEvent(t *testing.T) {
	for name, bad := range map[string]abci.Event{
		"malformed attribute":         withAttr(leaseClosedEvent(t, 2), "id", undecodableAttr),
		"enum value from newer chain": withAttr(leaseClosedEvent(t, 2), "reason", `"lease_closed_reason_from_a_newer_chain"`),
	} {
		t.Run(name, func(t *testing.T) {
			node := newFakeNode()
			rec := newRecLogger()
			sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithLogger(rec))

			node.set(func(f *fakeNode) {
				f.blocks[101] = txBlock(101, leaseClosedEvent(t, 1), bad, leaseClosedEvent(t, 3))
				f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
			})
			node.pushHeader(102)

			require.Equal(t, []uint64{1, 3, 102}, collectLeaseClosed(t, sub, 3))
			requireNoMoreEvents(t, sub, 50*time.Millisecond)
			require.Equal(t, 1, node.callsAt(101), "an undecodable event is not a reason to fetch its height again")
			require.Equal(t, 1, rec.count(levelError, "skipping undecodable event"))
		})
	}
}

// A node that answers a fetch with another height's results, as a misrouted or cached
// response would, must not have them published as the requested height: the height is
// fetched again and published from the right answer.
func TestEventsRetriesMismatchedHeight(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		f.serveOnce[101] = txBlock(100, leaseClosedEvent(t, 100))
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	require.Equal(t, []uint64{101}, collectLeaseClosed(t, sub, 1))
	require.Equal(t, 2, node.callsAt(101), "the mismatched answer, then the right one")
	requireNoMoreEvents(t, sub, 50*time.Millisecond)
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

	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}, WithStartHeight(90))
	node.pushHeader(93)

	require.Equal(t, []uint64{91, 92, 93}, collectLeaseClosed(t, sub, 3), "block 90 is the caller's, everything after it is published")
}

// A height whose fetch fails only transiently must not be dropped. Its events are
// published once the node recovers, even if headers arrived in a burst while the
// fetch was failing.
func TestEventsRecoversBlockAfterTransientFetchFailure(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 4 * time.Millisecond, prunedCheckAfter: 500 * time.Millisecond})

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

	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 50 * time.Millisecond}) // watchdog active, but nothing to replay

	select {
	case ev := <-sub.Events():
		t.Fatalf("unexpected replay of historical event %T", ev)
	case <-time.After(100 * time.Millisecond):
	}
}

// Headers repeat and arrive out of order, and the watchdog's catch-up covers the same
// heights again. Each height is still published exactly once, in order.
func TestEventsPublishesEachHeightOnce(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

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

// x/gov's EndBlocker runs the messages of passed proposals and re-emits their events at
// block level, so the service must publish block-level events too, in execution order:
// PreBlock, then BeginBlock, then each tx, then EndBlock.
func TestEventsPublishesBlockEventsInExecutionOrder(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})

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
// once in the log, not on every retry. Prefetch is off, so any fetch of a later height
// would mean the cursor had moved.
func TestEventsParksOnPrunedHeight(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: 3 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond}, WithLogger(rec), WithPrefetch(1))

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

// A pruned-check answer that comes back after its height was served explains a stall
// that is over: it is not logged.
func TestEventsIgnoresPrunedCheckAfterRecovery(t *testing.T) {
	node := newFakeNode()
	rec := newRecLogger()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond}, WithLogger(rec))

	entered := make(chan struct{}, 1)
	node.set(func(f *fakeNode) {
		f.statusDelay = 100 * time.Millisecond
		f.statusEntered = entered
		f.earliest = 150
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	requireSignal(t, entered, "the pruned check to ask for Status")
	node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	waitLeaseClosed(t, sub, 101) // served while the pruned check is still waiting

	time.Sleep(150 * time.Millisecond) // the pruned check answers
	require.Zero(t, rec.count(levelWarn, "node has pruned the parked height"))
}

// A height exactly equal to the node's earliest retained height is still retained,
// so a fetch failure there must be retried, never treated as pruned.
func TestEventsRetriesHeightEqualToEarliest(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

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

// A node reporting earliest 0 (unknown) must never cause a skip. A permanently failing
// height is retried, and nothing later is ever published.
func TestEventsDoesNotSkipWhenEarliestUnknown(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

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
// skipped, since the pruned check cannot confirm pruning and so cannot assume it either.
func TestEventsRetriesWhenStatusUnavailable(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

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
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 200 * time.Millisecond})

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

// Fifteen headers at once while one height fails must cost exactly one fetch attempt
// for that height within a single retry interval, not one per header.
func TestEventsBurstOfHeadersDoesNotBurnRetries(t *testing.T) {
	node := newFakeNode()
	newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 2 * time.Second, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) { f.blockErr[101] = errTransientFetch })

	for h := int64(101); h <= 115; h++ {
		node.pushHeader(h)
	}

	require.Eventually(t, func() bool { return node.callsAt(101) >= 1 }, time.Second, time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	require.Equal(t, 1, node.callsAt(101), "a burst of headers must not burn through retries for one height")
}

// A height that fails forever, with earliest at or below it, must never be skipped,
// so a later and perfectly fetchable height is never published ahead of it.
func TestEventsDoesNotAdvancePastFailingBlock(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

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
	sub := newTestService(t, node, tuning{watchdog: 5 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 500 * time.Millisecond})

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

	require.Equal(t, dseqsThrough(105), collectLeaseClosed(t, sub, 5))
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

	svc, err := newEventsOn(ctx, node, bus, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}.options()...)
	require.NoError(t, err)

	bus.Close()
	node.pushHeader(101)

	select {
	case <-svc.lc.Done():
	case <-time.After(2 * time.Second):
		t.Fatal("run loop did not stop after the bus closed")
	}

	select {
	case <-svc.lc.ShuttingDown():
	default:
		t.Fatal("lifecycle never entered shutting-down, so WatchContext and the shutdown bridge stay parked forever")
	}

	require.ErrorIs(t, svc.Err(), pubsub.ErrNotRunning, "the service reports why it stopped")
	require.True(t, node.allStreamsClosed(), "a stopped service leaves no stream open")
}

// A block fetch that never returns, as on a dead connection or behind a full
// connection pool, is abandoned after the fetch timeout and retried, so one stuck
// request cannot hold every later event.
func TestEventsRetriesFetchThatHangs(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour, fetchTimeout: 20 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.hangFirst[101] = 1
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
	})
	node.pushHeader(102)

	require.Equal(t, []uint64{101, 102}, collectLeaseClosed(t, sub, 2))
	require.Equal(t, 2, node.callsAt(101), "the hung attempt, then one retry")
}

// The watchdog asks the node how far the chain has got. A node that never answers must
// not hold the run loop, or the headers that arrive meanwhile are never processed.
func TestEventsWatchdogStatusThatHangsDoesNotHoldEvents(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour, fetchTimeout: 20 * time.Millisecond})

	entered := make(chan struct{}, 1)
	node.set(func(f *fakeNode) {
		f.statusHang = true
		f.statusEntered = entered
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	requireSignal(t, entered, "the watchdog to ask for Status")
	node.pushHeader(101)

	waitLeaseClosed(t, sub, 101)
}

// While a height is parked, the pruned check asks the node for its retention floor. A
// node that never answers must not hold the run loop, or the parked height is never
// retried.
func TestEventsPrunedCheckThatHangsDoesNotHoldEvents(t *testing.T) {
	node := newFakeNode()
	sub := newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond, fetchTimeout: 20 * time.Millisecond})

	node.set(func(f *fakeNode) {
		f.statusHang = true
		f.blockErr[101] = errTransientFetch
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})
	node.pushHeader(101)

	time.AfterFunc(100*time.Millisecond, func() {
		node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	})

	waitLeaseClosed(t, sub, 101)
}

// A height that never comes back is retried for as long as the service runs; the
// service neither stops nor publishes past it.
func TestEventsRetriesFailingHeightForever(t *testing.T) {
	node := newFakeNode()
	svc, sub := startTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 2 * time.Millisecond, prunedCheckAfter: time.Hour})

	node.set(func(f *fakeNode) {
		f.blockErr[101] = errTransientFetch
		f.blocks[102] = txBlock(102, leaseClosedEvent(t, 102))
	})
	node.pushHeader(102)

	require.Eventually(t, func() bool { return node.callsAt(101) >= 50 }, 5*time.Second, time.Millisecond)
	select {
	case <-svc.Done():
		t.Fatalf("service stopped: %v", svc.Err())
	default:
	}
	requireNoMoreEvents(t, sub, 20*time.Millisecond)

	node.set(func(f *fakeNode) { delete(f.blockErr, 101) })
	require.Equal(t, []uint64{102}, collectLeaseClosed(t, sub, 1))
}

// The pruned check costs a call to the node, so a parked height asks it once, whether
// or not the node answers.
func TestEventsPrunedCheckRunsOncePerHeight(t *testing.T) {
	node := newFakeNode()
	newTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	before := node.statusCallCount()
	node.set(func(f *fakeNode) {
		f.statusErr = errors.New("rpc unavailable")
		f.blockErr[101] = errTransientFetch
	})
	node.pushHeader(101)

	require.Eventually(t, func() bool { return node.callsAt(101) >= 10 }, 2*time.Second, time.Millisecond,
		"the height keeps being retried well past prunedCheckAfter")
	require.Equal(t, 1, node.statusCallCount()-before)
}

// Shutting the service down while its pruned check is waiting on the node is an owner
// stop, even when the check returns long after. Err must not report a failure the owner
// cut short.
func TestEventsShutdownDuringPrunedCheckIsNotAFailure(t *testing.T) {
	node := newFakeNode()
	svc, _ := startTestService(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: 10 * time.Millisecond})

	entered := make(chan struct{}, 1)
	node.set(func(f *fakeNode) {
		f.statusDelay = 300 * time.Millisecond // returns late, whatever its context
		f.statusEntered = entered
		f.blockErr[101] = errTransientFetch
	})
	node.pushHeader(101)

	requireSignal(t, entered, "the pruned check to ask for Status")
	svc.Shutdown()

	require.NoError(t, requireStopped(t, svc, 2*time.Second))
}

// Shutdown cancels everything in flight, a dial, a fetch and a Status call, rather
// than waiting any of them out, and leaves no stream open.
func TestEventsShutdownIsPromptWithIOInFlight(t *testing.T) {
	node := newFakeNode()
	svc, _ := startTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: time.Hour, prunedCheckAfter: time.Hour, fetchTimeout: time.Hour})
	node.waitDials(t, 1)

	statusBefore := node.statusCallCount()
	node.set(func(f *fakeNode) {
		f.hangFirst[101] = 1
		f.statusHang = true
		f.dialHang = true
	})
	node.pushHeader(101)
	node.endStream(io.EOF)

	require.Eventually(t, func() bool {
		return node.fetchesInflight() == 1 && node.dials() == 2 && node.statusCallCount() > statusBefore
	}, 2*time.Second, time.Millisecond, "a fetch, a dial and a Status call in flight")

	start := time.Now()
	svc.Shutdown()
	require.Less(t, time.Since(start), 200*time.Millisecond)
	require.True(t, node.allStreamsClosed())
}

// Err reports why the service stopped on its own. A service its owner stopped, through
// Shutdown or by cancelling the parent context, stopped for no error.
func TestEventsErrIsNilWhenStoppedByOwner(t *testing.T) {
	tune := tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour}

	start := func(t *testing.T, ctx context.Context) Service {
		t.Helper()

		bus := pubsub.NewBus()
		t.Cleanup(bus.Close)

		node := newFakeNode()
		svc, err := newEventsOn(ctx, node, bus, tune.options()...)
		require.NoError(t, err)
		return svc
	}

	t.Run("shutdown", func(t *testing.T) {
		svc := start(t, context.Background())
		require.NoError(t, svc.Err(), "a running service has no error")

		svc.Shutdown()
		require.NoError(t, requireStopped(t, svc, time.Second))
	})

	t.Run("parent context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		svc := start(t, ctx)

		cancel()
		require.NoError(t, requireStopped(t, svc, time.Second))
	})
}

// NewEvents bounds its own calls to the node too. A node that never answers Status fails
// NewEvents after the fetch timeout instead of blocking the caller.
func TestNewEventsFailsWhenStatusHangs(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) { f.statusHang = true })

	err := newEventsWithin(t, node, tuning{watchdog: time.Hour, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour, fetchTimeout: 20 * time.Millisecond}, 2*time.Second)
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

// The websocket is dialed in the background: a node that never completes the dial does
// not hold NewEvents, and events still flow through the watchdog.
func TestNewEventsDoesNotWaitForTheWebsocket(t *testing.T) {
	node := newFakeNode()
	node.set(func(f *fakeNode) {
		f.dialHang = true
		f.blocks[101] = txBlock(101, leaseClosedEvent(t, 101))
	})

	start := time.Now()
	sub := newTestService(t, node, tuning{watchdog: 10 * time.Millisecond, retryInterval: 5 * time.Millisecond, prunedCheckAfter: time.Hour})
	require.Less(t, time.Since(start), 100*time.Millisecond)

	node.set(func(f *fakeNode) { f.height = 101 }) // the chain moves on after the start
	waitLeaseClosed(t, sub, 101)
}

// newEventsWithin runs newEvents and returns its error, failing the test if it has not
// returned within the given time. A service it did start is shut down with the test.
func newEventsWithin(t *testing.T, node testNode, tune tuning, within time.Duration) error {
	t.Helper()

	bus := pubsub.NewBus()
	t.Cleanup(bus.Close)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	type result struct {
		svc *events
		err error
	}
	done := make(chan result, 1)
	go func() {
		svc, err := newEventsOn(ctx, node, bus, tune.options()...)
		done <- result{svc: svc, err: err}
	}()

	select {
	case res := <-done:
		if res.svc != nil {
			t.Cleanup(res.svc.Shutdown)
		}
		return res.err
	case <-time.After(within):
		t.Fatalf("newEvents still blocked after %s", within)
		return nil
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

// A node keeps pruning while the service waits, and every block above the floor is
// fetchable. The cursor must still stay on the first pruned height, driven by both
// headers and watchdog ticks, rather than chase the floor. The only height ever asked
// for below the floor is the one the cursor is parked on, prefetch being off.
func TestEventsStaysParkedBehindMovingPruneFloor(t *testing.T) {
	const blockTime = 6 * time.Millisecond

	node := newPruningNode(100, 100, leaseClosedEvent(t, 1))
	sub := newTestService(t, node, tuning{
		watchdog:         blockTime,
		retryInterval:    blockTime,
		prunedCheckAfter: 5 * blockTime,
	}, WithPrefetch(1))

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
