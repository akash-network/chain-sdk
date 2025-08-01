package testutil

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"
	tmbytes "github.com/cometbft/cometbft/libs/bytes"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpcclientmock "github.com/cometbft/cometbft/rpc/client/mock"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"

	arpcclient "pkg.akt.dev/go/node/client"
)

type MockRPC interface {
	client.CometRPC
	Akash(ctx context.Context)
}

var _ arpcclient.RPCClient = (*MockCometRPC)(nil)

type MockCometRPC struct {
	rpcclientmock.Client

	responseQuery abci.ResponseQuery
}

// NewMockCometRPC returns a mock TendermintRPC implementation.
// It is used for CLI testing.
func NewMockCometRPC(respQuery abci.ResponseQuery) MockCometRPC {
	return MockCometRPC{responseQuery: respQuery}
}

func (MockCometRPC) BroadcastTxSync(context.Context, tmtypes.Tx) (*coretypes.ResultBroadcastTx, error) {
	return &coretypes.ResultBroadcastTx{Code: 0}, nil
}

func (m MockCometRPC) ABCIQueryWithOptions(
	_ context.Context,
	_ string,
	_ tmbytes.HexBytes,
	_ rpcclient.ABCIQueryOptions,
) (*coretypes.ResultABCIQuery, error) {
	return &coretypes.ResultABCIQuery{Response: m.responseQuery}, nil
}

func (MockCometRPC) Akash(_ context.Context) *arpcclient.Akash {
	return &arpcclient.Akash{
		ClientInfo: &arpcclient.ClientInfo{
			ApiVersion: arpcclient.VersionV1beta3,
		},
	}
}
