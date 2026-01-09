package v1_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "pkg.akt.dev/go/node/market/v1"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name             string
		err              *sdkerrors.Error
		expectedGRPCCode codes.Code
		expectedABCICode uint32
	}{
		{
			name:             "empty_provider",
			err:              v1.ErrEmptyProvider,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 1,
		},
		{
			name:             "same_account",
			err:              v1.ErrSameAccount,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 2,
		},
		{
			name:             "internal",
			err:              v1.ErrInternal,
			expectedGRPCCode: codes.Internal,
			expectedABCICode: 3,
		},
		{
			name:             "bid_over_order",
			err:              v1.ErrBidOverOrder,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 4,
		},
		{
			name:             "attribute_mismatch",
			err:              v1.ErrAttributeMismatch,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 5,
		},
		{
			name:             "unknown_bid",
			err:              v1.ErrUnknownBid,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 6,
		},
		{
			name:             "unknown_lease",
			err:              v1.ErrUnknownLease,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 7,
		},
		{
			name:             "unknown_lease_for_bid",
			err:              v1.ErrUnknownLeaseForBid,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 8,
		},
		{
			name:             "unknown_order_for_bid",
			err:              v1.ErrUnknownOrderForBid,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 9,
		},
		{
			name:             "lease_not_active",
			err:              v1.ErrLeaseNotActive,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 10,
		},
		{
			name:             "bid_not_active",
			err:              v1.ErrBidNotActive,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 11,
		},
		{
			name:             "bid_not_open",
			err:              v1.ErrBidNotOpen,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 12,
		},
		{
			name:             "order_not_open",
			err:              v1.ErrOrderNotOpen,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 13,
		},
		{
			name:             "no_lease_for_order",
			err:              v1.ErrNoLeaseForOrder,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 14,
		},
		{
			name:             "order_not_found",
			err:              v1.ErrOrderNotFound,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 15,
		},
		{
			name:             "group_not_found",
			err:              v1.ErrGroupNotFound,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 16,
		},
		{
			name:             "group_not_open",
			err:              v1.ErrGroupNotOpen,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 17,
		},
		{
			name:             "bid_not_found",
			err:              v1.ErrBidNotFound,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 18,
		},
		{
			name:             "bid_zero_price",
			err:              v1.ErrBidZeroPrice,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 19,
		},
		{
			name:             "lease_not_found",
			err:              v1.ErrLeaseNotFound,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 20,
		},
		{
			name:             "bid_exists",
			err:              v1.ErrBidExists,
			expectedGRPCCode: codes.AlreadyExists,
			expectedABCICode: 21,
		},
		{
			name:             "bid_invalid_price",
			err:              v1.ErrBidInvalidPrice,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 22,
		},
		{
			name:             "order_active",
			err:              v1.ErrOrderActive,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 23,
		},
		{
			name:             "order_closed",
			err:              v1.ErrOrderClosed,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 24,
		},
		{
			name:             "order_exists",
			err:              v1.ErrOrderExists,
			expectedGRPCCode: codes.AlreadyExists,
			expectedABCICode: 25,
		},
		{
			name:             "order_duration_exceeded",
			err:              v1.ErrOrderDurationExceeded,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 26,
		},
		{
			name:             "order_too_early",
			err:              v1.ErrOrderTooEarly,
			expectedGRPCCode: codes.FailedPrecondition,
			expectedABCICode: 27,
		},
		{
			name:             "invalid_deposit",
			err:              v1.ErrInvalidDeposit,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 28,
		},
		{
			name:             "invalid_param",
			err:              v1.ErrInvalidParam,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 29,
		},
		{
			name:             "unknown_provider",
			err:              v1.ErrUnknownProvider,
			expectedGRPCCode: codes.NotFound,
			expectedABCICode: 30,
		},
		{
			name:             "invalid_bid",
			err:              v1.ErrInvalidBid,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 31,
		},
		{
			name:             "capabilities_mismatch",
			err:              v1.ErrCapabilitiesMismatch,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 32,
		},
		{
			name:             "invalid_lease_closed_reason",
			err:              v1.ErrInvalidLeaseClosedReason,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 33,
		},
		{
			name:             "invalid_escrow_id",
			err:              v1.ErrInvalidEscrowID,
			expectedGRPCCode: codes.InvalidArgument,
			expectedABCICode: 34,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st, ok := status.FromError(tt.err)
			require.True(t, ok, "error should be convertible to gRPC status")
			require.Equal(t, tt.expectedGRPCCode, st.Code(), "gRPC status code mismatch")
			require.Equal(t, tt.expectedABCICode, tt.err.ABCICode(), "ABCI error code mismatch")
		})
	}
}

