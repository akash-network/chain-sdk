package events

import (
	"context"
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
	defaultWatchdogTimeout = 10 * time.Second
	// maxBlockRetries is how many times a single height is retried before it is
	// skipped, so an unfetchable height (for example pruned on the node) cannot
	// stall the pipeline forever.
	maxBlockRetries = 10
)

type events struct {
	ctx             context.Context
	ebus            cmclient.EventsClient
	client          sdkclient.CometRPC
	bus             pubsub.Bus
	lc              lifecycle.Lifecycle
	log             log.Logger
	watchdogTimeout time.Duration
	lastHeight      int64
	failures        int
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
// timeout: the subscription channel is never closed when it silently stalls (a load
// balancer holding the socket open, say), so the watchdog is what notices and fetches
// the missed blocks to keep events flowing. A height that cannot be fetched is retried
// and eventually skipped so the pipeline never wedges. It starts from the current
// height, so historical events are not replayed.
func NewEvents(pctx context.Context, node sdkclient.CometRPC, name string, bus pubsub.Bus) (Service, error) {
	return newEvents(pctx, node, name, bus, defaultWatchdogTimeout)
}

func newEvents(pctx context.Context, node sdkclient.CometRPC, name string, bus pubsub.Bus, watchdogTimeout time.Duration) (Service, error) {
	ev := &events{
		ctx:             pctx,
		ebus:            node.(cmclient.EventsClient),
		client:          node,
		bus:             bus,
		lc:              lifecycle.New(),
		log:             ctxlog.Logger(pctx).With("cmp", "events"),
		watchdogTimeout: watchdogTimeout,
	}

	status, err := node.Status(pctx)
	if err != nil {
		return nil, err
	}
	ev.lastHeight = status.SyncInfo.LatestBlockHeight

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

	// Bridge a shutdown request (explicit Shutdown or parent-context cancellation
	// via WatchContext) to context cancellation, so an in-progress catch-up is
	// interrupted promptly instead of blocking shutdown until the gap is drained.
	go func() {
		e.lc.ShutdownInitiated(<-e.lc.ShutdownRequest())
		cancel()
	}()

	defer func() {
		unsubCtx, ucancel := context.WithTimeout(context.Background(), unsubscribeTimeout)
		_ = e.ebus.UnsubscribeAll(unsubCtx, subs)
		ucancel()
		e.lc.ShutdownCompleted()
	}()

	// Watchdog: reset on every block, fires after watchdogTimeout of silence.
	watchdog := time.NewTimer(e.watchdogTimeout)
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
				e.catchUp(ctx, evt.Header.Height)
			}
			resetTimer(watchdog, e.watchdogTimeout)

		case <-watchdog.C:
			// No block for watchdogTimeout: the subscription has gone quiet. Fetch the
			// latest height and catch up so events keep flowing until it recovers.
			if status, err := e.client.Status(ctx); err != nil {
				e.log.Error("watchdog: fetch chain status, will retry", "err", err)
			} else {
				e.catchUp(ctx, status.SyncInfo.LatestBlockHeight)
			}
			resetTimer(watchdog, e.watchdogTimeout)
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
// A height that fails is retried on subsequent calls and, after maxBlockRetries, skipped
// with a loud log so an unfetchable (e.g. pruned) height cannot stall the pipeline.
func (e *events) catchUp(ctx context.Context, targetHeight int64) {
	for e.lastHeight < targetHeight {
		select {
		case <-ctx.Done():
			return
		default:
		}

		height := e.lastHeight + 1
		if err := e.processBlock(ctx, height); err != nil {
			e.failures++
			e.log.Error("process block, will retry", "height", height, "attempt", e.failures, "err", err)
			if e.failures < maxBlockRetries {
				return
			}
			e.log.Error("skipping unfetchable block after repeated failures", "height", height)
		}

		e.failures = 0
		e.lastHeight = height
	}
}

func (e *events) processBlock(ctx context.Context, height int64) error {
	blkResults, err := e.client.BlockResults(ctx, &height)
	if err != nil {
		return err
	}

	for _, tx := range blkResults.TxsResults {
		if tx == nil {
			continue
		}

		for _, ev := range tx.Events {
			if mev, ok := processEvent(ev); ok {
				if err := e.bus.Publish(mev); err != nil {
					return err
				}
			}
		}
	}

	return nil
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
