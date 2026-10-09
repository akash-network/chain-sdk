package events

import (
	"context"
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

	atypes "pkg.akt.dev/go/node/audit/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1"
	mtypes "pkg.akt.dev/go/node/market/v1"

	"pkg.akt.dev/go/util/ctxlog"
	"pkg.akt.dev/go/util/pubsub"
)

const (
	queueSize          = 1000
	unsubscribeTimeout = 5 * time.Second
	// defaultWatchdogTimeout is how long we wait for a new block before assuming the
	// subscription has silently stalled and fetching the missed blocks ourselves.
	defaultWatchdogTimeout  = 10 * time.Second
	defaultMinRetryInterval = 2 * time.Second
	// defaultPrunedCheckAfter is how long a height stays stalled before we ask the node
	// whether it has been pruned. It gates the question, never the decision to skip.
	defaultPrunedCheckAfter = 30 * time.Second
)

// tuning holds the time constants the service runs on, so tests can drive it fast.
type tuning struct {
	watchdog         time.Duration
	minRetryInterval time.Duration
	prunedCheckAfter time.Duration
}

// stall is the height the cursor cannot get past. The zero value means healthy.
type stall struct {
	height  int64
	nextTry time.Time
}

type events struct {
	ctx        context.Context
	ebus       cmclient.EventsClient
	client     sdkclient.CometRPC
	bus        pubsub.Bus
	lc         lifecycle.Lifecycle
	log        log.Logger
	tune       tuning
	lastHeight int64
	// lastProgress is when a block was last fetched and published. A jump over pruned
	// heights is not progress, so chasing a floor that keeps moving does not keep
	// rearming the pruned check and pay for it again on every landing.
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
// events it recognises. It tracks the last processed height, so a gap between headers
// (for example after a dropped and re-established subscription) is back-filled rather
// than lost. A watchdog resets on every block and fires if none arrives within the
// timeout. The subscription channel is never closed when it silently stalls (a load
// balancer holding the socket open, say), so the watchdog is what notices and fetches
// the missed blocks to keep events flowing. A height that cannot be fetched is retried
// indefinitely, at most once per minRetryInterval as headers and watchdog ticks arrive,
// and is skipped only once the node itself reports
// it below its retained history (SyncInfo.EarliestBlockHeight), never on a guess from
// an attempt count. A height that is unfetchable but not pruned parks the cursor
// rather than silently dropping its events. It starts from the current height, so
// historical events are not replayed.
func NewEvents(pctx context.Context, node sdkclient.CometRPC, name string, bus pubsub.Bus) (Service, error) {
	return newEvents(pctx, node, name, bus, tuning{
		watchdog:         defaultWatchdogTimeout,
		minRetryInterval: defaultMinRetryInterval,
		prunedCheckAfter: defaultPrunedCheckAfter,
	})
}

func newEvents(pctx context.Context, node sdkclient.CometRPC, name string, bus pubsub.Bus, tune tuning) (Service, error) {
	ev := &events{
		ctx:    pctx,
		ebus:   node.(cmclient.EventsClient),
		client: node,
		bus:    bus,
		lc:     lifecycle.New(),
		log:    ctxlog.Logger(pctx).With("cmp", "events"),
		tune:   tune,
	}

	status, err := node.Status(pctx)
	if err != nil {
		return nil, err
	}
	ev.lastHeight = status.SyncInfo.LatestBlockHeight
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

	defer func() {
		initiate(exitErr)

		unsubCtx, ucancel := context.WithTimeout(context.Background(), unsubscribeTimeout)
		_ = e.ebus.UnsubscribeAll(unsubCtx, subs)
		ucancel()
		e.lc.ShutdownCompleted()
	}()

	watchdog := time.NewTimer(e.tune.watchdog)
	defer watchdog.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case ev, ok := <-ch:
			if !ok {
				// The subscription channel should never close; if it does, drop it and
				// let the watchdog carry the service rather than hot-spinning.
				e.log.Error("block subscription channel closed, relying on watchdog fetch")
				ch = nil
				continue
			}
			if evt, ok := ev.Data.(cmtypes.EventDataNewBlockHeader); ok {
				if err := e.catchUp(ctx, evt.Header.Height); err != nil {
					e.log.Error("catch up, stopping", "err", err)
					exitErr = err
					return
				}
			}
			resetTimer(watchdog, e.tune.watchdog)

		case <-watchdog.C:
			if status, err := e.client.Status(ctx); err != nil {
				e.log.Error("watchdog fetch chain status failed, will retry", "err", err)
			} else if err := e.catchUp(ctx, status.SyncInfo.LatestBlockHeight); err != nil {
				e.log.Error("catch up, stopping", "err", err)
				exitErr = err
				return
			}
			resetTimer(watchdog, e.tune.watchdog)
		}
	}
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
// A height that fails is retried indefinitely, paced by e.tune.minRetryInterval, and is
// skipped only once parkCursor sees the node itself report the height below its
// retained history. It returns non-nil only on a terminal publish error, which means
// the caller must stop rather than keep retrying.
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
			if e.parkCursor(ctx, height, err) {
				continue
			}
			return nil
		}

		for _, ev := range evs {
			if err := e.bus.Publish(ev); err != nil {
				return err
			}
		}

		e.stall = stall{}
		e.lastProgress = time.Now()
		e.lastHeight = height
	}

	return nil
}

// parkCursor records height as the stalled cursor position and, once the cursor has
// gone e.tune.prunedCheckAfter without real progress, asks the node whether height has
// been pruned
// out from under us. Only then, and only if the node's own earliest-retained height
// is strictly past it, does it jump the cursor forward.
//
// It reports whether it jumped, so the caller retries in the same pass. The node keeps
// pruning while we wait, so a jump followed by a wait for the next block lands behind
// the floor again and the cursor never catches it.
func (e *events) parkCursor(ctx context.Context, height int64, err error) bool {
	first := e.stall.height != height
	e.stall = stall{height: height, nextTry: time.Now().Add(e.tune.minRetryInterval)}

	// A parked cursor retries for as long as it takes, so logging every attempt at
	// error level buries the one line that matters under tens of thousands a day.
	stalled := time.Since(e.lastProgress)
	if first {
		e.log.Error("block fetch failed, cursor parked", "height", height, "stalled", stalled, "err", err)
	} else {
		e.log.Debug("block fetch still failing, cursor parked", "height", height, "stalled", stalled, "err", err)
	}

	if stalled < e.tune.prunedCheckAfter {
		return false
	}

	status, serr := e.client.Status(ctx)
	if serr != nil {
		return false
	}

	earliest := status.SyncInfo.EarliestBlockHeight
	if earliest <= height {
		return false
	}

	e.log.Error("node pruned past the cursor, events in range are lost", "from", height, "through", earliest-1, "node_earliest", earliest)
	e.lastHeight = earliest - 1
	e.stall = stall{}

	return true
}

// blockEvents fetches height and returns the events worth publishing. It has no
// side effects, so a retry of the same height cannot re-publish what already went out.
func (e *events) blockEvents(ctx context.Context, height int64) ([]interface{}, error) {
	blkResults, err := e.client.BlockResults(ctx, &height)
	if err != nil {
		return nil, err
	}

	// Only tx events are collected. Every Akash EndBlocker returns no events today, so
	// FinalizeBlockEvents carries nothing we publish. That is an assumption about the
	// node modules, not about this package.
	var evs []interface{}
	for _, tx := range blkResults.TxsResults {
		if tx == nil {
			continue
		}

		for _, ev := range tx.Events {
			if mev, ok := processEvent(ev); ok {
				evs = append(evs, mev)
			}
		}
	}

	return evs, nil
}

func processEvent(bev abci.Event) (interface{}, bool) {
	pev, err := sdk.ParseTypedEvent(bev)
	if err != nil {
		return nil, false
	}

	switch pev.(type) {
	case *atypes.EventTrustedAuditorCreated:
	case *atypes.EventTrustedAuditorDeleted:
	case *dtypes.EventDeploymentCreated:
	case *dtypes.EventDeploymentUpdated:
	case *dtypes.EventDeploymentClosed:
	case *dtypes.EventGroupStarted:
	case *dtypes.EventGroupPaused:
	case *dtypes.EventGroupClosed:
	case *mtypes.EventOrderCreated:
	case *mtypes.EventOrderClosed:
	case *mtypes.EventBidCreated:
	case *mtypes.EventBidClosed:
	case *mtypes.EventLeaseCreated:
	case *mtypes.EventLeaseClosed:
	case *mtypes.EventLeaseReclaimStarted:
	default:
		return nil, false
	}

	return pev, true
}
