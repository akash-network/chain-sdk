package v1_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
)

func TestSettlementID(t *testing.T) {
	id := types.ID{
		Owner:    account(17),
		Height:   9001,
		Sequence: 42,
	}
	for status := types.StatusTransferPending; status <= types.StatusRefunded; status++ {
		key := types.NewSettlementID(status, id)
		require.NoError(t, key.ValidateBasic())
		require.Equal(t, id, key.ID())
		require.Equal(t, status, key.Status)
	}
	for _, status := range []types.Status{types.StatusInvalid, -1, 6, 256} {
		require.Error(t, types.NewSettlementID(status, id).ValidateBasic())
	}
	for _, bad := range []types.ID{
		{},
		{
			Owner:    id.Owner,
			Height:   0,
			Sequence: 1,
		},
		{
			Owner:    id.Owner,
			Height:   -1,
			Sequence: 1,
		},
		{
			Owner:    id.Owner,
			Height:   1,
			Sequence: 0,
		},
		{
			Owner:    account(0),
			Height:   1,
			Sequence: 1,
		},
	} {
		require.Error(t, types.NewSettlementID(types.StatusMinted, bad).ValidateBasic())
	}
}
