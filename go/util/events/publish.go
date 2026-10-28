package events

import (
	"context"
	"sync"
	"time"

	"github.com/boz/go-lifecycle"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"

	"cosmossdk.io/log"

	"pkg.akt.dev/go/util/ctxlog"
	"pkg.akt.dev/go/util/pubsub"
)

// Service represents an event monitoring service that watches the chain and
// publishes the transaction events it recognises to a message bus.
type Service interface {
	// Shutdown stops the event monitoring service and waits for it to finish.
	Shutdown()
	// Done is closed once the service has stopped, whether through Shutdown, its parent
	// context, or on its own (see Err).
	Done() <-chan struct{}
	// Err returns why the service stopped on its own: the bus's error when the bus
	// stopped accepting events. It returns nil while the service runs and once it was
	// stopped through Shutdown or its parent context.
	Err() error
}

type events struct {
	ctx    context.Context
	client queryNode
	dial   dialFunc
	bus    pubsub.Bus
	lc     lifecycle.Lifecycle
	log    log.Logger
	opts   options
	// start is the last height the caller had accounted for when the service started.
	start int64
	// failure is why the run loop stopped on its own. It is written before the
	// lifecycle completes and read only after Done is closed.
	failure error
}

// NewEvents creates and starts a blockchain event monitoring service for the node at
// nodeURI: tcp://, http://, https://, ws://, wss:// or unix://.
//
// The service opens its own connection to the node, a websocket for block headers and
// an HTTP client for block results, so its lifecycle is its own, separate from any
// client the caller holds. It processes each block in order, publishing the events it
// recognises exactly once: none is published twice and no height is skipped (see below
// for the one kind of event that is). It tracks the last published height, so a gap
// between headers, for example across a reconnect, is back-filled rather than lost.
//
// The websocket is watched with ping and pong. A connection that stops answering,
// whether reset, closed or silently dropped somewhere along the way, is noticed within
// the pong wait and redialed: the first time at once, then once per retry interval
// (WithRetryInterval, DefaultRetryInterval by default) for as long as it takes. The
// service never subscribes twice on one connection; a node that rejects the subscription
// is asked again on a fresh connection, at the same pace. Every loss of the
// subscription is logged at warn level and every dial at info level (see WithLogger).
// While no header arrives, a watchdog (WithWatchdogTimeout, DefaultWatchdogTimeout by
// default) asks the node how far the chain has got, so events keep flowing while the
// websocket is down or the chain is slow.
//
// Blocks are fetched off the service's loop, up to WithPrefetch (DefaultPrefetch by
// default) at a time while it catches up, and published strictly in order. Every call to
// the node is bounded by the fetch timeout (WithFetchTimeout, DefaultFetchTimeout by
// default) and cancelled when the service stops, so a stuck request never holds it and
// Shutdown returns promptly.
//
// A height that cannot be fetched parks the cursor and is retried once per retry
// interval on its own timer, so a transient failure (a pooled backend that does not have
// the height yet, a rate-limited request) costs one retry interval rather than waiting
// for the next header or watchdog tick. The cursor never moves past it, because a
// skipped height is a lost event: a missed EventLeaseClosed leaves a provider running a
// closed lease. Later events wait behind it until it is served, however long that
// takes: the service never gives up on a height. A long stall is logged at error level,
// once. A node that answers with another height's results is treated as failing.
//
// A recognized event that cannot be decoded, for example one carrying an enum value
// from a newer chain, is skipped and logged at error level. Retrying its height cannot
// change the outcome, so the rest of the height and later heights go out as usual.
//
// The service stops on its own only when the bus refuses an event; Err then says why.
//
// It starts from the node's latest height, so historical events are not replayed,
// unless WithStartHeight says where the caller left off.
func NewEvents(ctx context.Context, nodeURI string, name string, bus pubsub.Bus, opts ...Option) (Service, error) {
	o, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	u, err := parseNodeURI(nodeURI)
	if err != nil {
		return nil, err
	}

	client, err := newQueryNode(ctx, u)
	if err != nil {
		return nil, err
	}

	cfg := wsConfig{
		target:     newWSTarget(u),
		query:      blkHeaderQuery().String(),
		timeout:    o.fetchTimeout,
		pingPeriod: o.pingPeriod,
		pongWait:   o.pongWait,
	}
	dial := func(ctx context.Context) (headerStream, error) { return dialHeaders(ctx, cfg) }

	e, err := newEvents(ctx, client, dial, name, bus, o)
	if err != nil {
		closeIdle(client)
		return nil, err
	}

	return e, nil
}

// newEvents starts the service on the given node client and websocket dialer, with the
// options NewEvents applied: it applies none itself, so each option is applied once.
func newEvents(ctx context.Context, client queryNode, dial dialFunc, name string, bus pubsub.Bus, o options) (*events, error) {
	lg := o.log
	if lg == nil {
		lg = ctxlog.Logger(ctx)
	}

	e := &events{
		ctx:    ctx,
		client: client,
		dial:   dial,
		bus:    bus,
		lc:     lifecycle.New(),
		log:    lg.With("cmp", "events", "name", name),
		opts:   o,
	}

	if o.startHeightSet {
		e.start = o.startHeight
	} else {
		status, err := e.status(ctx)
		if err != nil {
			return nil, err
		}
		e.start = status.SyncInfo.LatestBlockHeight
	}

	go e.lc.WatchContext(ctx)
	go e.run()

	return e, nil
}

func (e *events) Shutdown() {
	select {
	case <-e.lc.Done():
		return
	default:
		e.lc.Shutdown(nil)
	}
}

func (e *events) Done() <-chan struct{} {
	return e.lc.Done()
}

func (e *events) Err() error {
	select {
	case <-e.lc.Done():
		return e.failure
	default:
		return nil
	}
}

// status asks the node how far the chain has got, giving up after the fetch timeout.
func (e *events) status(ctx context.Context) (*ctypes.ResultStatus, error) {
	sctx, cancel := context.WithTimeout(ctx, e.opts.fetchTimeout)
	defer cancel()

	return e.client.Status(sctx)
}

func (e *events) run() {
	ctx, cancel := context.WithCancel(e.ctx)
	defer cancel()

	// run can leave on its own when the bus is gone, not only on a shutdown request,
	// and lifecycle watchers stay parked until ShuttingDown is closed. So whichever
	// side gets here first initiates the shutdown, and it happens exactly once.
	var once sync.Once
	initiate := func(err error) {
		once.Do(func() {
			e.lc.ShutdownInitiated(err)
			cancel()
		})
	}

	// Bridge a shutdown request (explicit Shutdown or parent-context cancellation via
	// WatchContext) to context cancellation, which every call to the node is bound to.
	go func() {
		select {
		case err := <-e.lc.ShutdownRequest():
			initiate(err)
		case <-ctx.Done():
		}
	}()

	l := newLoop(ctx, e)
	err := l.run()
	if err != nil {
		e.log.Error("catch up, stopping", "err", err)
	}

	initiate(err)
	l.close()
	closeIdle(e.client)

	// Written before ShutdownCompleted closes Done, which is what Err waits for.
	e.failure = err
	e.lc.ShutdownCompleted()
}

// dialResult, fetchResult, statusResult and pruneResult carry the outcome of a call to
// the node back to the run loop.
type dialResult struct {
	stream headerStream
	err    error
}

type fetchResult struct {
	height int64
	events []interface{}
	err    error
}

type statusResult struct {
	latest int64
	err    error
}

type pruneResult struct {
	height   int64
	earliest int64
	stalled  time.Duration
	err      error
}

// loop is the run goroutine's state. Only the run goroutine touches it, and it never
// waits on the node: every call runs in a goroutine of its own, bound to the service
// context, and reports back on one of the channels. Their senders never block, as each
// send also gives way to the context ending.
type loop struct {
	e   *events
	ctx context.Context
	wg  sync.WaitGroup
	cur *cursor

	// stream is the live header subscription, nil while there is none. At most one
	// dial is in flight.
	stream  headerStream
	dialing bool
	dials   chan dialResult
	redial  *time.Timer

	// outage runs from a stream ending until the first new header after it. redialed is
	// set once the outage's immediate redial is spent; later ones wait the retry
	// interval. attempts counts the dials since the last header, for the logs.
	outage   bool
	redialed bool
	attempts int

	fetches chan fetchResult
	retry   *time.Timer

	// watchdog fires when no new header has arrived for a window; the loop then asks
	// the node for its height, one call at a time.
	watchdog       *time.Timer
	statuses       chan statusResult
	statusInflight bool
	pruneChecks    chan pruneResult

	// lastHeader is the highest header seen, and lastHeaderAt when it arrived. Only a
	// header above lastHeader is a new block.
	lastHeader   int64
	lastHeaderAt time.Time
}

func newLoop(ctx context.Context, e *events) *loop {
	return &loop{
		e:            e,
		ctx:          ctx,
		cur:          newCursor(e.start, e.opts.prefetch, e.opts.retryInterval, e.opts.prunedCheckAfter),
		dials:        make(chan dialResult, 1),
		redial:       newStoppedTimer(),
		fetches:      make(chan fetchResult, e.opts.prefetch),
		retry:        newStoppedTimer(),
		watchdog:     time.NewTimer(e.opts.watchdog),
		statuses:     make(chan statusResult, 1),
		pruneChecks:  make(chan pruneResult, 1),
		lastHeaderAt: time.Now(),
	}
}

// run serves the loop until the service context ends, or until the bus refuses an
// event, which it returns.
func (l *loop) run() error {
	l.startDial()

	for {
		var notify, ended <-chan struct{}
		if l.stream != nil {
			notify, ended = l.stream.Notify(), l.stream.Done()
		}

		select {
		case <-l.ctx.Done():
			return nil
		case <-notify:
			l.onHeader(l.stream.Latest())
		case <-ended:
			l.onStreamEnded()
		case r := <-l.dials:
			l.onDial(r)
		case <-l.redial.C:
			l.startDial()
		case r := <-l.fetches:
			if err := l.onFetch(r); err != nil {
				return err
			}
		case <-l.retry.C:
			l.schedule()
		case <-l.watchdog.C:
			l.onWatchdog()
		case r := <-l.statuses:
			l.onStatus(r)
		case r := <-l.pruneChecks:
			l.onPruneCheck(r)
		}
	}
}

// close releases what the loop holds once the service context has ended: the stream,
// every goroutine (each returns promptly, its call being bound to the context), and a
// stream a dial delivered too late to be used.
func (l *loop) close() {
	if l.stream != nil {
		l.stream.Close()
		l.stream = nil
	}

	l.wg.Wait()

	select {
	case r := <-l.dials:
		if r.stream != nil {
			r.stream.Close()
		}
	default:
	}

	l.watchdog.Stop()
	l.retry.Stop()
	l.redial.Stop()
}

// onHeader handles the highest header the stream has seen. Only a header above the last
// one is a new block: a repeated or older one neither holds the watchdog off nor claims
// the subscription is back.
func (l *loop) onHeader(height int64) {
	if height > l.lastHeader {
		switch {
		case l.outage:
			l.e.log.Info("block subscription restored",
				"down_for", time.Since(l.lastHeaderAt), "attempts", l.attempts, "height", height, "last_height", l.cur.last())
			l.outage, l.redialed = false, false
		case l.lastHeader > 0 && height > l.lastHeader+1:
			// The stream went on without delivering some heights, so the only trace of
			// the gap is the heights it never delivered.
			l.e.log.Warn("block subscription missed blocks, back-filling", "from", l.lastHeader+1, "to", height-1)
		}
		l.lastHeader = height
		l.lastHeaderAt = time.Now()
		l.attempts = 0
		resetTimer(l.watchdog, l.e.opts.watchdog)
	}

	l.cur.observe(height)
	l.schedule()
}

// onStreamEnded drops a stream that has ended and redials: at once the first time in an
// outage, then once per retry interval.
func (l *loop) onStreamEnded() {
	s := l.stream
	l.stream = nil

	// A header that raced the end is still a header.
	if h := s.Latest(); h > l.lastHeader {
		l.onHeader(h)
	}
	err := s.Err()
	s.Close()

	if !l.outage {
		l.outage, l.redialed, l.attempts = true, false, 0
		l.e.log.Warn("block subscription lost, redialing", "err", err, "last_header", l.lastHeader, "last_height", l.cur.last())
	}

	if !l.redialed {
		l.redialed = true
		l.startDial()
		return
	}
	resetTimer(l.redial, l.e.opts.retryInterval)
}

// startDial dials the node's websocket off the loop, unless a dial is in flight or a
// stream is live.
func (l *loop) startDial() {
	if l.dialing || l.stream != nil {
		return
	}
	l.dialing = true
	l.attempts++
	l.e.log.Info("dialing block subscription", "attempt", l.attempts)

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()

		dctx, cancel := context.WithTimeout(l.ctx, l.e.opts.fetchTimeout)
		defer cancel()

		s, err := l.e.dial(dctx)
		select {
		case l.dials <- dialResult{stream: s, err: err}:
		case <-l.ctx.Done():
			if s != nil {
				s.Close()
			}
		}
	}()
}

func (l *loop) onDial(r dialResult) {
	l.dialing = false

	if l.ctx.Err() != nil {
		if r.stream != nil {
			r.stream.Close()
		}
		return
	}

	if r.err != nil {
		l.e.log.Warn("block subscription dial failed, retrying",
			"attempt", l.attempts, "retry_in", l.e.opts.retryInterval, "err", r.err)
		resetTimer(l.redial, l.e.opts.retryInterval)
		return
	}

	l.stream = r.stream
	l.e.log.Info("block subscription connected", "attempt", l.attempts)
}

// onWatchdog asks the node how far the chain has got when no new header arrived for a
// window. A live stream with no header is a slow chain, not a lost subscription, so
// nothing is redialed.
func (l *loop) onWatchdog() {
	l.e.log.Debug("no new block within watchdog window, polling node",
		"timeout", l.e.opts.watchdog, "last_header", l.lastHeader, "last_height", l.cur.last())

	if !l.statusInflight {
		l.statusInflight = true
		l.wg.Add(1)
		go func() {
			defer l.wg.Done()

			var r statusResult
			if status, err := l.e.status(l.ctx); err != nil {
				r.err = err
			} else {
				r.latest = status.SyncInfo.LatestBlockHeight
			}

			select {
			case l.statuses <- r:
			case <-l.ctx.Done():
			}
		}()
	}

	resetTimer(l.watchdog, l.e.opts.watchdog)
}

func (l *loop) onStatus(r statusResult) {
	l.statusInflight = false

	if r.err != nil {
		if l.ctx.Err() == nil {
			l.e.log.Warn("watchdog fetch chain status failed, will retry", "err", r.err)
		}
		return
	}

	l.cur.observe(r.latest)
	l.schedule()
}

// schedule starts the fetches the cursor wants now, and arms the retry timer for a
// parked height.
func (l *loop) schedule() {
	for _, h := range l.cur.toFetch(time.Now()) {
		l.startFetch(h)
	}

	if at, ok := l.cur.wakeAt(); ok {
		resetTimer(l.retry, time.Until(at))
	} else {
		stopTimer(l.retry)
	}
}

func (l *loop) startFetch(height int64) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()

		evs, err := l.e.blockEvents(l.ctx, height)
		select {
		case l.fetches <- fetchResult{height: height, events: evs, err: err}:
		case <-l.ctx.Done():
		}
	}()
}

// onFetch hands a fetch to the cursor and publishes every height that is now next in
// line. It returns an error only when the bus refuses an event.
func (l *loop) onFetch(r fetchResult) error {
	if r.err != nil && l.ctx.Err() != nil {
		return nil // cut short by the service stopping
	}

	if ev, parked := l.cur.done(r.height, r.events, r.err, time.Now()); parked {
		l.parked(ev)
	} else if r.err != nil {
		l.e.log.Debug("block fetch ahead of the cursor failed, refetching when the cursor gets there",
			"height", r.height, "err", r.err)
	}

	for {
		blk, ok := l.cur.pop(time.Now())
		if !ok {
			break
		}
		for _, ev := range blk.events {
			if err := l.e.bus.Publish(ev); err != nil {
				return err
			}
		}
		if blk.recovered {
			l.e.log.Info("block fetch recovered, cursor resumed", "height", blk.height, "stalled", blk.stalled)
		}
	}

	l.schedule()
	return nil
}

// parked logs a failure of the cursor height: at warn level the first time, since
// behind pooled RPC backends a height that one backend does not have yet, or a
// rate-limited request, is routine and recovers on the next attempt; once at error level
// when the stall outlasts the pruned-check delay; at debug level in between, so a long
// stall does not bury the lines that matter. At that delay it also asks the node,
// once, whether the height is pruned.
func (l *loop) parked(ev parkEvent) {
	switch {
	case ev.first:
		l.e.log.Warn("block fetch failed, cursor parked", "height", ev.height, "retry_in", l.e.opts.retryInterval, "err", ev.err)
	case ev.escalate:
		l.e.log.Error("block fetch keeps failing, cursor stalled", "height", ev.height, "stalled", ev.stalled, "err", ev.err)
	default:
		l.e.log.Debug("block fetch still failing, cursor parked", "height", ev.height, "stalled", ev.stalled, "err", ev.err)
	}

	if ev.pruneCheck {
		l.startPruneCheck(ev.height, ev.stalled)
	}
}

// startPruneCheck asks the node, off the loop, whether height is below its retained
// history. The cursor stays parked either way: the answer only explains the stall in
// the log, so an operator knows a node that retains the height is needed.
func (l *loop) startPruneCheck(height int64, stalled time.Duration) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()

		r := pruneResult{height: height, stalled: stalled}
		if status, err := l.e.status(l.ctx); err != nil {
			r.err = err
		} else {
			r.earliest = status.SyncInfo.EarliestBlockHeight
		}

		select {
		case l.pruneChecks <- r:
		case <-l.ctx.Done():
		}
	}()
}

// onPruneCheck explains a stall the node has pruned under. An answer for a height that
// has been served since explains nothing any more.
func (l *loop) onPruneCheck(r pruneResult) {
	if r.err != nil || !l.cur.parkedOn(r.height) {
		return
	}

	if r.earliest > r.height {
		l.e.log.Warn("node has pruned the parked height, cursor stays parked until a node that retains it answers",
			"height", r.height, "node_earliest", r.earliest, "stalled", r.stalled)
	}
}

// newStoppedTimer returns a timer that is not running.
func newStoppedTimer() *time.Timer {
	t := time.NewTimer(time.Hour)
	stopTimer(t)
	return t
}

// resetTimer safely re-arms t to fire after d. It is only called from the run
// goroutine, so the stop-drain-reset sequence has no concurrent timer access.
func resetTimer(t *time.Timer, d time.Duration) {
	stopTimer(t)
	t.Reset(d)
}

// stopTimer stops t and drains a fire that has not been received, so a later select
// does not see it.
func stopTimer(t *time.Timer) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
}
