package events

import (
	"context"
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/stretchr/testify/require"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	mtypes "pkg.akt.dev/go/node/market/v1"
	"pkg.akt.dev/go/util/pubsub"
)

// blockResultsClient is a minimal CometRPC stub that returns canned block
// results so processBlock can be exercised without a live node.
type blockResultsClient struct {
	sdkclient.CometRPC
	results *ctypes.ResultBlockResults
}

func (c *blockResultsClient) BlockResults(_ context.Context, _ *int64) (*ctypes.ResultBlockResults, error) {
	return c.results, nil
}

func leaseClosedABCIEvent(t *testing.T, owner string) abci.Event {
	t.Helper()
	ev, err := sdk.TypedEventToEvent(&mtypes.EventLeaseClosed{
		ID: mtypes.LeaseID{Owner: owner, DSeq: 1, GSeq: 1, OSeq: 1, Provider: "akash1provider"},
	})
	require.NoError(t, err)
	return abci.Event(ev)
}

// Test_processBlock_FinalizeBlockEvents guards the #515 fix: a lease closed in
// an EndBlocker (e.g. escrow depletion) surfaces in FinalizeBlockEvents with no
// associated transaction, and must still reach the bus.
func Test_processBlock_FinalizeBlockEvents(t *testing.T) {
	bus := pubsub.NewBus()
	defer bus.Close()

	sub, err := bus.Subscribe()
	require.NoError(t, err)

	e := &events{
		ctx: context.Background(),
		bus: bus,
		client: &blockResultsClient{
			results: &ctypes.ResultBlockResults{
				Height: 1,
				// a tx-initiated close (must still be published)
				TxsResults: []*abci.ExecTxResult{
					{Events: []abci.Event{leaseClosedABCIEvent(t, "akash1txowner")}},
				},
				// an EndBlock close (escrow depletion) — previously dropped
				FinalizeBlockEvents: []abci.Event{leaseClosedABCIEvent(t, "akash1endblockowner")},
			},
		},
	}

	e.processBlock(1)

	owners := map[string]bool{}
	for i := 0; i < 2; i++ {
		select {
		case ev := <-sub.Events():
			lc, ok := ev.(*mtypes.EventLeaseClosed)
			require.True(t, ok, "expected EventLeaseClosed, got %T", ev)
			owners[lc.ID.Owner] = true
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for published events, got %d", len(owners))
		}
	}

	require.True(t, owners["akash1txowner"], "tx-initiated lease close must be published")
	require.True(t, owners["akash1endblockowner"], "EndBlock (escrow-depletion) lease close must be published")
}
