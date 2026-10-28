package events

import (
	"context"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	atypes "pkg.akt.dev/go/node/audit/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1"
	mtypes "pkg.akt.dev/go/node/market/v1"
)

// blockEvents fetches height, giving up after the fetch timeout, and returns the events
// worth publishing, in the order the chain executed them. It publishes nothing, so a
// retry of the same height cannot re-publish what already went out. A recognized event
// that cannot be decoded is skipped and logged: retrying cannot change the outcome, and
// it must not hold the rest of the height or the heights after it. An answer for another
// height is a failed fetch, so a misrouted or stale response is retried rather than
// published as the requested height.
func (e *events) blockEvents(ctx context.Context, height int64) ([]interface{}, error) {
	fctx, cancel := context.WithTimeout(ctx, e.opts.fetchTimeout)
	defer cancel()

	blkResults, err := e.client.BlockResults(fctx, &height)
	if err != nil {
		return nil, err
	}
	if blkResults.Height != height {
		return nil, fmt.Errorf("node returned block results for height %d, asked for %d", blkResults.Height, height)
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
	collect := func(raw []abci.Event) {
		for _, ev := range raw {
			mev, ok, err := processEvent(ev)
			if err != nil {
				e.log.Error("skipping undecodable event", "height", height, "type", ev.Type, "err", err)
				continue
			}
			if ok {
				evs = append(evs, mev)
			}
		}
	}

	collect(begin)
	for _, tx := range blkResults.TxsResults {
		if tx == nil {
			continue
		}
		collect(tx.Events)
	}
	collect(end)

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
// not be read, such as one carrying an enum value from a newer chain.
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
