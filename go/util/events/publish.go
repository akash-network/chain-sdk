package events

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/boz/go-lifecycle"

	abci "github.com/cometbft/cometbft/abci/types"
	cmclient "github.com/cometbft/cometbft/rpc/client"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	cmtypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/log"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	atypes "pkg.akt.dev/go/node/audit/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1"
	mtypes "pkg.akt.dev/go/node/market/v1"

	"pkg.akt.dev/go/util/ctxlog"
	"pkg.akt.dev/go/util/pubsub"
)

const (
	queueSize          = 1000
	unsubscribeTimeout = 5 * time.Second
)

// subscribeResult carries the outcome of a re-subscribe back to the run loop.
type subscribeResult struct {
	ch      <-chan ctypes.ResultEvent
	err     error
	attempt int
}

// stall is the height the cursor cannot get past. The zero value means healthy.
type stall struct {
	height  int64
	nextTry time.Time
	// pruneChecked is set once the node has been asked whether height is pruned, so a
	// cursor parked for days asks once rather than on every retry.
	pruneChecked bool
}

type events struct {
	ctx        context.Context
	ebus       cmclient.EventsClient
	client     sdkclient.CometRPC
	bus        pubsub.Bus
	lc         lifecycle.Lifecycle
	log        log.Logger
	opts       options
	lastHeight int64
	// lastProgress is when a block was last fetched and published, the clock a parked
	// cursor's stall is measured against.
	lastProgress time.Time
	stall        stall
}

// Service represents an event monitoring service that watches the chain and
// publishes the transaction events it recognises to a message bus.
type Service interface {
	// Shutdown gracefully stops the event monitoring service and waits for it to finish.
	Shutdown()
}

// NewEvents creates and starts a blockchain event monitoring service.
//
// It subscribes to block headers and processes each block in order, publishing the
// events it recognises exactly once: none is skipped and none is published twice. It
// tracks the last processed height, so a gap between headers (for example after a
// dropped and re-established subscription) is back-filled rather than lost.
//
// A watchdog resets on every block and fires if none arrives within the timeout
// (WithWatchdogTimeout, DefaultWatchdogTimeout by default). The subscription channel is
// never closed when the connection under it dies, so the watchdog is what notices: it
// fetches the missed blocks to keep events flowing and re-subscribes. The re-subscribe
// matters because the CometBFT websocket client redials only after a websocket close
// or a failed write. After a reset or a keepalive timeout it goes silent for good
// until something writes to it again. Every loss of the subscription is logged at warn
// level and every attempt to re-establish it at info level (see WithLogger).
//
// A height that cannot be fetched or decoded parks the cursor and is retried
// indefinitely, at most once per minRetryInterval as headers and watchdog ticks arrive.
// The cursor never moves past it, even once the node has pruned it, because a skipped
// height is a lost event: a missed EventLeaseClosed leaves a provider running a closed
// lease. A parked cursor holds back every later event until the height is served.
//
// It starts from the node's latest height, so historical events are not replayed,
// unless WithStartHeight says where the caller left off.
func NewEvents(pctx context.Context, node sdkclient.CometRPC, name string, bus pubsub.Bus, opts ...Option) (Service, error) {
	o, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	lg := o.log
	if lg == nil {
		lg = ctxlog.Logger(pctx)
	}

	ev := &events{
		ctx:    pctx,
		ebus:   node.(cmclient.EventsClient),
		client: node,
		bus:    bus,
		lc:     lifecycle.New(),
		log:    lg.With("cmp", "events"),
		opts:   o,
	}

	if o.startHeightSet {
		ev.lastHeight = o.startHeight
	} else {
		status, err := node.Status(pctx)
		if err != nil {
			return nil, err
		}
		ev.lastHeight = status.SyncInfo.LatestBlockHeight
	}
	ev.lastProgress = time.Now()

	blkHeaderName := name + "-blk-hdr"
	blkch, err := ev.ebus.Subscribe(pctx, blkHeaderName, blkHeaderQuery().String(), queueSize)
	if err != nil {
		return nil, err
	}

	go ev.lc.WatchContext(pctx)
	go ev.run(blkHeaderName, blkch)

	return ev, nil
}

func (e *events) Shutdown() {
	select {
	case <-e.lc.Done():
		return
	default:
		e.lc.Shutdown(nil)
	}
}

func (e *events) run(subs string, ch <-chan ctypes.ResultEvent) {
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

	// Bridge a shutdown request (explicit Shutdown or parent-context cancellation
	// via WatchContext) to context cancellation, so an in-progress catch-up is
	// interrupted promptly instead of blocking shutdown until the gap is drained.
	go func() {
		select {
		case err := <-e.lc.ShutdownRequest():
			initiate(err)
		case <-ctx.Done():
		}
	}()

	// The reason run leaves is the one thing lc.Error() can carry upward, so give it
	// the terminal error rather than a nil that reads as a clean shutdown.
	var exitErr error

	// resub is non-nil while a re-subscribe is in flight. At most one runs at a time,
	// and it runs off the loop, because Subscribe blocks while the websocket client is
	// redialing and the loop must keep catching up meanwhile.
	var resub chan subscribeResult

	defer func() {
		initiate(exitErr)

		// Let an in-flight re-subscribe land before unsubscribing, or it could register
		// the query again after UnsubscribeAll cleared it. ctx is already cancelled, so
		// it returns promptly.
		if resub != nil {
			select {
			case <-resub:
			case <-time.After(unsubscribeTimeout):
			}
		}

		unsubCtx, ucancel := context.WithTimeout(context.Background(), unsubscribeTimeout)
		_ = e.ebus.UnsubscribeAll(unsubCtx, subs)
		ucancel()
		e.lc.ShutdownCompleted()
	}()

	watchdog := time.NewTimer(e.opts.watchdog)
	defer watchdog.Stop()

	// known is the highest height the chain is known to have reached, from headers and
	// from Status. Every header and watchdog tick catches up to it, so a parked height
	// is retried on each tick even while Status is failing and no header arrives.
	known := e.lastHeight

	// Disconnects are logged at warn level and every step of re-establishing the
	// subscription at info level. down is set from the first watchdog expiry or channel
	// close until a new block arrives again, so an outage is reported as lost once and
	// restored once. Only a header above lastHeader is a new block: a repeated or older
	// one must neither hold the watchdog off nor claim the subscription is back.
	var (
		lastHeader   int64
		lastHeaderAt = time.Now()
		down         bool
		attempts     int
	)

	for {
		select {
		case <-ctx.Done():
			return

		case ev, ok := <-ch:
			if !ok {
				// The subscription channel should never close; if it does, drop it and
				// let the watchdog re-subscribe rather than hot-spinning. The close is
				// itself the loss, so the watchdog does not report the outage again.
				e.log.Warn("block subscription channel closed, re-subscribing on watchdog expiry",
					"last_header", lastHeader, "last_height", e.lastHeight)
				down = true
				ch = nil
				continue
			}

			evt, ok := ev.Data.(cmtypes.EventDataNewBlockHeader)
			if !ok {
				continue
			}
			height := evt.Header.Height

			if height > lastHeader {
				switch {
				case down:
					e.log.Info("block subscription restored",
						"down_for", time.Since(lastHeaderAt), "attempts", attempts, "height", height, "last_height", e.lastHeight)
					down, attempts = false, 0
				case lastHeader > 0 && height > lastHeader+1:
					// The client redialed on its own within the watchdog window, so the
					// only trace of the disconnect is the heights it never delivered.
					e.log.Warn("block subscription missed blocks, back-filling",
						"from", lastHeader+1, "to", height-1)
				}
				lastHeader = height
				lastHeaderAt = time.Now()
				resetTimer(watchdog, e.opts.watchdog)
			}
			if height > known {
				known = height
			}

			// A repeated or older header still retries a parked cursor; catchUp paces
			// retries, and the cursor keeps it from publishing anything twice.
			if err := e.catchUp(ctx, known); err != nil {
				e.log.Error("catch up, stopping", "err", err)
				exitErr = err
				return
			}

		case <-watchdog.C:
			if len(ch) > 0 {
				// Headers arrived while a long catch-up held the loop; the subscription
				// is alive and they are next in line, so this is not an outage.
				resetTimer(watchdog, e.opts.watchdog)
				continue
			}
			if !down {
				down = true
				e.log.Warn("block subscription presumed lost: no new block within watchdog window",
					"timeout", e.opts.watchdog, "last_header", lastHeader, "last_height", e.lastHeight)
			}
			if resub == nil {
				attempts++
				e.log.Info("re-subscribing to block headers", "attempt", attempts, "down_for", time.Since(lastHeaderAt))
				resub = e.resubscribe(ctx, subs, attempts)
			}
			if status, err := e.client.Status(ctx); err != nil {
				e.log.Warn("watchdog fetch chain status failed, will retry", "err", err)
			} else if latest := status.SyncInfo.LatestBlockHeight; latest > known {
				known = latest
			}
			if err := e.catchUp(ctx, known); err != nil {
				e.log.Error("catch up, stopping", "err", err)
				exitErr = err
				return
			}
			resetTimer(watchdog, e.opts.watchdog)

		case res := <-resub:
			resub = nil
			if res.err != nil {
				e.log.Warn("re-subscribe failed, retrying on next watchdog expiry", "attempt", res.attempt, "err", res.err)
				continue
			}
			e.log.Info("re-subscribe sent, waiting for block headers", "attempt", res.attempt)
			ch = res.ch
		}
	}
}

// resubscribe subscribes to block headers again, off the run loop, and delivers the
// outcome on the returned channel. Calling Subscribe is what gets a silently dead
// websocket redialed: the write fails, and the client reconnects and replays its
// subscriptions. On a live connection the node answers that the query is already
// subscribed, which the client ignores, so re-subscribing while the chain is merely
// slow is harmless. The call is bounded by the watchdog window, because Subscribe
// blocks for as long as the client is redialing. Success means the request went out,
// not that headers flow again; the first header after it is what confirms that.
func (e *events) resubscribe(ctx context.Context, subs string, attempt int) chan subscribeResult {
	res := make(chan subscribeResult, 1)

	go func() {
		sctx, cancel := context.WithTimeout(ctx, e.opts.watchdog)
		defer cancel()

		ch, err := e.ebus.Subscribe(sctx, subs, blkHeaderQuery().String(), queueSize)
		res <- subscribeResult{ch: ch, err: err, attempt: attempt}
	}()

	return res
}

// resetTimer safely re-arms t to fire after d. It is only called from the run
// goroutine, so the stop-drain-reset sequence has no concurrent timer access.
func resetTimer(t *time.Timer, d time.Duration) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
		}
	}
	t.Reset(d)
}

// catchUp processes every block from the last one handled up to targetHeight, in order.
// A height that fails is parked and retried indefinitely, paced by
// e.opts.minRetryInterval, and is never skipped. The cursor advances only once every
// event of a height has been published, so no height is published twice either. It
// returns non-nil only on a terminal publish error, which means the caller must stop
// rather than keep retrying.
func (e *events) catchUp(ctx context.Context, targetHeight int64) error {
	for e.lastHeight < targetHeight {
		if ctx.Err() != nil {
			return nil
		}

		height := e.lastHeight + 1
		if e.stall.height == height && time.Now().Before(e.stall.nextTry) {
			return nil
		}

		evs, err := e.blockEvents(ctx, height)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			e.parkCursor(ctx, height, err)
			return nil
		}

		for _, ev := range evs {
			if err := e.bus.Publish(ev); err != nil {
				return err
			}
		}

		if e.stall.height == height {
			e.log.Info("block fetch recovered, cursor resumed", "height", height, "stalled", time.Since(e.lastProgress))
		}

		e.stall = stall{}
		e.lastProgress = time.Now()
		e.lastHeight = height
	}

	return nil
}

// parkCursor records height as the stalled cursor position. The cursor stays there until
// the height is fetched and decoded: it never skips, even once the node has pruned the
// height, because a skipped height is a lost event. Pooled backends report very
// different pruning floors, so the fetch keeps being retried against whichever answers.
//
// Once the stall has lasted e.opts.prunedCheckAfter it asks the node, once per height,
// whether the height is below its retained history. The answer only explains the stall
// in the log, so an operator knows a node that retains the height is needed.
func (e *events) parkCursor(ctx context.Context, height int64, err error) {
	first := e.stall.height != height
	e.stall = stall{
		height:       height,
		nextTry:      time.Now().Add(e.opts.minRetryInterval),
		pruneChecked: !first && e.stall.pruneChecked,
	}

	// A parked cursor retries for as long as it takes, so logging every attempt at
	// error level buries the one line that matters under tens of thousands a day. A
	// decode failure is named as such, so it does not send an operator after the network.
	msgFirst, msgAgain := "block fetch failed, cursor parked", "block fetch still failing, cursor parked"
	var derr *decodeError
	if errors.As(err, &derr) {
		msgFirst, msgAgain = "block decode failed, cursor parked", "block decode still failing, cursor parked"
	}

	stalled := time.Since(e.lastProgress)
	if first {
		e.log.Error(msgFirst, "height", height, "stalled", stalled, "err", err)
	} else {
		e.log.Debug(msgAgain, "height", height, "stalled", stalled, "err", err)
	}

	if e.stall.pruneChecked || stalled < e.opts.prunedCheckAfter {
		return
	}

	status, serr := e.client.Status(ctx)
	if serr != nil {
		return
	}
	e.stall.pruneChecked = true

	if earliest := status.SyncInfo.EarliestBlockHeight; earliest > height {
		e.log.Warn("node has pruned the parked height, cursor stays parked until a node that retains it answers",
			"height", height, "node_earliest", earliest, "stalled", stalled)
	}
}

// blockEvents fetches height and returns the events worth publishing, in the order the
// chain executed them. It has no side effects, so a retry of the same height cannot
// re-publish what already went out. A recognized event that cannot be decoded fails
// the whole height, so the cursor parks on it rather than dropping it.
func (e *events) blockEvents(ctx context.Context, height int64) ([]interface{}, error) {
	blkResults, err := e.client.BlockResults(ctx, &height)
	if err != nil {
		return nil, err
	}

	// Akash's own blockers emit none of the recognized events, but x/gov's EndBlocker
	// executes the messages of passed proposals and re-emits their events there, so
	// block-level events are read too. The SDK keeps them apart from tx events (those
	// appear only in TxsResults) and lists them PreBlock, BeginBlock, EndBlock, so
	// splitting off EndBlock restores execution order around the txs.
	var begin, end []abci.Event
	for _, ev := range blkResults.FinalizeBlockEvents {
		if blockMode(ev) == modeEndBlock {
			end = append(end, ev)
		} else {
			begin = append(begin, ev)
		}
	}

	var evs []interface{}
	collect := func(raw []abci.Event) error {
		for _, ev := range raw {
			mev, ok, err := processEvent(ev)
			if err != nil {
				return &decodeError{height: height, typ: ev.Type, err: err}
			}
			if ok {
				evs = append(evs, mev)
			}
		}
		return nil
	}

	if err := collect(begin); err != nil {
		return nil, err
	}
	for _, tx := range blkResults.TxsResults {
		if tx == nil {
			continue
		}
		if err := collect(tx.Events); err != nil {
			return nil, err
		}
	}
	if err := collect(end); err != nil {
		return nil, err
	}

	return evs, nil
}

const (
	// modeKey is the attribute the SDK appends to every BeginBlock and EndBlock event.
	// Its value is a bare word, not JSON, unlike every attribute of a typed event.
	modeKey        = "mode"
	modeBeginBlock = "BeginBlock"
	modeEndBlock   = "EndBlock"
)

// recognized is the set of typed events the service publishes, keyed by the proto
// message name the SDK writes as the event type.
var recognized = func() map[string]struct{} {
	m := make(map[string]struct{})
	for _, msg := range []proto.Message{
		&atypes.EventTrustedAuditorCreated{},
		&atypes.EventTrustedAuditorDeleted{},
		&dtypes.EventDeploymentCreated{},
		&dtypes.EventDeploymentUpdated{},
		&dtypes.EventDeploymentClosed{},
		&dtypes.EventGroupStarted{},
		&dtypes.EventGroupPaused{},
		&dtypes.EventGroupClosed{},
		&mtypes.EventOrderCreated{},
		&mtypes.EventOrderClosed{},
		&mtypes.EventBidCreated{},
		&mtypes.EventBidClosed{},
		&mtypes.EventLeaseCreated{},
		&mtypes.EventLeaseClosed{},
		&mtypes.EventLeaseReclaimStarted{},
	} {
		m[proto.MessageName(msg)] = struct{}{}
	}
	return m
}()

// blockMode returns the phase the SDK tagged a block-level event with, or "" for a
// PreBlock event, which carries no tag.
func blockMode(ev abci.Event) string {
	for _, attr := range ev.Attributes {
		if attr.Key == modeKey && (attr.Value == modeBeginBlock || attr.Value == modeEndBlock) {
			return attr.Value
		}
	}
	return ""
}

// processEvent decodes bev when it is one of the recognized types; ok is false for any
// other event. The SDK's mode tag is stripped first, because ParseTypedEvent reads
// every attribute value as JSON and the bare word would fail it. A typed event's own
// fields are always JSON, so the tag cannot be mistaken for one. The SDK's other
// bookkeeping attributes (msg_index, authz_msg_index) are JSON numbers the decoder
// already ignores. An error therefore means a recognized event the chain emitted could
// not be read, which must never be dropped.
func processEvent(bev abci.Event) (interface{}, bool, error) {
	if _, ok := recognized[bev.Type]; !ok {
		return nil, false, nil
	}

	attrs := make([]abci.EventAttribute, 0, len(bev.Attributes))
	for _, attr := range bev.Attributes {
		if attr.Key == modeKey && (attr.Value == modeBeginBlock || attr.Value == modeEndBlock) {
			continue
		}
		attrs = append(attrs, attr)
	}

	pev, err := sdk.ParseTypedEvent(abci.Event{Type: bev.Type, Attributes: attrs})
	if err != nil {
		return nil, false, err
	}

	return pev, true, nil
}

// decodeError is a recognized event the chain emitted that could not be decoded. It
// parks the cursor like a fetch failure, since dropping it would lose the event, but it
// is reported as what it is: a format the service cannot read, not a network problem.
type decodeError struct {
	height int64
	typ    string
	err    error
}

func (e *decodeError) Error() string {
	return fmt.Sprintf("height %d: decode %s: %v", e.height, e.typ, e.err)
}

func (e *decodeError) Unwrap() error { return e.err }
