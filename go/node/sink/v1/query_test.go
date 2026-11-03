package v1_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestQueryValidation(t *testing.T) {
	for _, p := range []*query.PageRequest{nil, {}, {Limit: 500, Reverse: true}} {
		require.NoError(t, types.ValidatePagination(p))
	}

	for _, p := range []*query.PageRequest{{Offset: 1}, {CountTotal: true}, {Limit: 501}} {
		require.ErrorIs(t, types.ValidatePagination(p), types.ErrInvalidPagination)
	}

	require.NoError(t, (types.QueryTransfersRequest{Owner: account(7), Status: types.StatusRefunded}).ValidateBasic())
	require.Error(t, (types.QueryTransfersRequest{Owner: "bad"}).ValidateBasic())
	require.Error(t, (types.QueryTransfersRequest{Status: 3}).ValidateBasic())
	require.Error(t, (types.QueryTransferRequest{}).ValidateBasic())
	require.Error(t, (types.QueryAccountRequest{}).ValidateBasic())
	require.Error(t, (types.QueryDisputesRequest{Pagination: &query.PageRequest{Offset: 1}}).ValidateBasic())
}
