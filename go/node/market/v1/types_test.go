package v1

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	dtypes "pkg.akt.dev/go/node/deployment/v1"
)

func TestLeaseClosedReasonRanges(t *testing.T) {
	tests := []struct {
		name             string
		reason           LeaseClosedReason
		expectedOwner    bool
		expectedProvider bool
		expectedNetwork  bool
	}{
		{
			name:             "owner range upper boundary (9999)",
			reason:           9999,
			expectedOwner:    true,
			expectedProvider: false,
			expectedNetwork:  false,
		},
		{
			name:             "provider range lower boundary (10000)",
			reason:           10000,
			expectedOwner:    false,
			expectedProvider: true,
			expectedNetwork:  false,
		},
		{
			name:             "provider range upper boundary (19999)",
			reason:           19999,
			expectedOwner:    false,
			expectedProvider: true,
			expectedNetwork:  false,
		},
		{
			name:             "network range lower boundary (20000)",
			reason:           20000,
			expectedOwner:    false,
			expectedProvider: false,
			expectedNetwork:  true,
		},
		{
			name:             "network range upper boundary (29999)",
			reason:           29999,
			expectedOwner:    false,
			expectedProvider: false,
			expectedNetwork:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isOwner := tc.reason.IsRange(LeaseClosedReasonRangeOwner)
			isProvider := tc.reason.IsRange(LeaseClosedReasonRangeProvider)
			isNetwork := tc.reason.IsRange(LeaseClosedReasonRangeNetwork)

			require.Equal(t, tc.expectedOwner, isOwner, "owner range check for %d", tc.reason)
			require.Equal(t, tc.expectedProvider, isProvider, "provider range check for %d", tc.reason)
			require.Equal(t, tc.expectedNetwork, isNetwork, "network range check for %d", tc.reason)
		})
	}
}

func TestBidIDValidateAllowsNonzeroBSeq(t *testing.T) {
	owner := sdk.AccAddress(bytes.Repeat([]byte{1}, 20))
	provider := sdk.AccAddress(bytes.Repeat([]byte{2}, 20))
	orderID := MakeOrderID(dtypes.MakeGroupID(dtypes.DeploymentID{
		Owner: owner.String(),
		DSeq:  1,
	}, 1), 1)
	bidID := MakeBidID(orderID, provider)
	bidID.BSeq = 2

	require.NoError(t, bidID.Validate())
}
