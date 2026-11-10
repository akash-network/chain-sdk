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
		require.Equal(t, &id, key.GetID())
		require.Equal(t, status, key.Status)
	}

	for _, status := range []types.Status{types.StatusInvalid, -1, 6, 256} {
		sid := types.NewSettlementID(status, id)
		err := sid.ValidateBasic()
		require.Error(t, err)
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
	} {
		sid := types.NewSettlementID(types.StatusMinted, bad)
		err := sid.ValidateBasic()

		require.Error(t, err)
	}
}
