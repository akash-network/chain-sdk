package v1_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "pkg.akt.dev/go/node/market/v1"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{
			name:         "unknown_bid_returns_not_found",
			err:          v1.ErrUnknownBid,
			expectedCode: codes.NotFound,
		},
		{
			name:         "unknown_lease_returns_not_found",
			err:          v1.ErrUnknownLease,
			expectedCode: codes.NotFound,
		},
		{
			name:         "unknown_lease_for_bid_returns_not_found",
			err:          v1.ErrUnknownLeaseForBid,
			expectedCode: codes.NotFound,
		},
		{
			name:         "unknown_order_for_bid_returns_not_found",
			err:          v1.ErrUnknownOrderForBid,
			expectedCode: codes.NotFound,
		},
		{
			name:         "no_lease_for_order_returns_not_found",
			err:          v1.ErrNoLeaseForOrder,
			expectedCode: codes.NotFound,
		},
		{
			name:         "order_not_found_returns_not_found",
			err:          v1.ErrOrderNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "group_not_found_returns_not_found",
			err:          v1.ErrGroupNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "bid_not_found_returns_not_found",
			err:          v1.ErrBidNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "lease_not_found_returns_not_found",
			err:          v1.ErrLeaseNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "unknown_provider_returns_not_found",
			err:          v1.ErrUnknownProvider,
			expectedCode: codes.NotFound,
		},
		{
			name:         "bid_exists_returns_already_exists",
			err:          v1.ErrBidExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "order_exists_returns_already_exists",
			err:          v1.ErrOrderExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "empty_provider_returns_invalid_argument",
			err:          v1.ErrEmptyProvider,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "same_account_returns_invalid_argument",
			err:          v1.ErrSameAccount,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "bid_over_order_returns_invalid_argument",
			err:          v1.ErrBidOverOrder,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "attribute_mismatch_returns_invalid_argument",
			err:          v1.ErrAttributeMismatch,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "capabilities_mismatch_returns_invalid_argument",
			err:          v1.ErrCapabilitiesMismatch,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "bid_zero_price_returns_invalid_argument",
			err:          v1.ErrBidZeroPrice,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "bid_invalid_price_returns_invalid_argument",
			err:          v1.ErrBidInvalidPrice,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_deposit_returns_invalid_argument",
			err:          v1.ErrInvalidDeposit,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_param_returns_invalid_argument",
			err:          v1.ErrInvalidParam,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_bid_returns_invalid_argument",
			err:          v1.ErrInvalidBid,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_lease_closed_reason_returns_invalid_argument",
			err:          v1.ErrInvalidLeaseClosedReason,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_escrow_id_returns_invalid_argument",
			err:          v1.ErrInvalidEscrowID,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "lease_not_active_returns_failed_precondition",
			err:          v1.ErrLeaseNotActive,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "bid_not_active_returns_failed_precondition",
			err:          v1.ErrBidNotActive,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "bid_not_open_returns_failed_precondition",
			err:          v1.ErrBidNotOpen,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "group_not_open_returns_failed_precondition",
			err:          v1.ErrGroupNotOpen,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "order_not_open_returns_failed_precondition",
			err:          v1.ErrOrderNotOpen,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "order_active_returns_failed_precondition",
			err:          v1.ErrOrderActive,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "order_closed_returns_failed_precondition",
			err:          v1.ErrOrderClosed,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "order_too_early_returns_failed_precondition",
			err:          v1.ErrOrderTooEarly,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "order_duration_exceeded_returns_failed_precondition",
			err:          v1.ErrOrderDurationExceeded,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "internal_returns_internal",
			err:          v1.ErrInternal,
			expectedCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st, ok := status.FromError(tt.err)
			require.True(t, ok, "error should be convertible to gRPC status")
			require.Equal(t, tt.expectedCode, st.Code(), "gRPC status code mismatch")
		})
	}
}

