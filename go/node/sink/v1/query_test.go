package v1_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestQueryValidation(t *testing.T) {
	for _, p := range []*query.PageRequest{nil, {}, {
		Limit:   500,
		Reverse: true,
	}, {Offset: 1}, {CountTotal: true}, {Limit: 501}} {
		require.NoError(t, (&types.QueryTransfersRequest{Pagination: p}).ValidateBasic())
		require.NoError(t, (&types.QueryDisputesRequest{Pagination: p}).ValidateBasic())
	}

	require.NoError(t, (&types.QueryTransfersRequest{
		Owner:  account(7),
		Status: types.StatusRefunded,
	}).ValidateBasic())
	require.Error(t, (&types.QueryTransfersRequest{Owner: "bad"}).ValidateBasic())
	require.Error(t, (&types.QueryTransfersRequest{Status: 6}).ValidateBasic())
	require.Error(t, (&types.QueryTransferRequest{}).ValidateBasic())
	require.Error(t, (&types.QueryAccountRequest{}).ValidateBasic())
}
