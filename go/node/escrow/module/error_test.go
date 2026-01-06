package module_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pkg.akt.dev/go/node/escrow/module"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{
			name:         "account_not_found_returns_not_found",
			err:          module.ErrAccountNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "payment_not_found_returns_not_found",
			err:          module.ErrPaymentNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "account_exists_returns_already_exists",
			err:          module.ErrAccountExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "payment_exists_returns_already_exists",
			err:          module.ErrPaymentExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "invalid_denomination_returns_invalid_argument",
			err:          module.ErrInvalidDenomination,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "payment_rate_zero_returns_invalid_argument",
			err:          module.ErrPaymentRateZero,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_payment_returns_invalid_argument",
			err:          module.ErrInvalidPayment,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_settlement_returns_invalid_argument",
			err:          module.ErrInvalidSettlement,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_id_returns_invalid_argument",
			err:          module.ErrInvalidID,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_account_returns_invalid_argument",
			err:          module.ErrInvalidAccount,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_account_depositor_returns_invalid_argument",
			err:          module.ErrInvalidAccountDepositor,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_deposit_returns_invalid_argument",
			err:          module.ErrInvalidDeposit,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_authz_scope_returns_invalid_argument",
			err:          module.ErrInvalidAuthzScope,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "account_closed_returns_failed_precondition",
			err:          module.ErrAccountClosed,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "account_overdrawn_returns_failed_precondition",
			err:          module.ErrAccountOverdrawn,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "payment_closed_returns_failed_precondition",
			err:          module.ErrPaymentClosed,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "unauthorized_deposit_scope_returns_permission_denied",
			err:          module.ErrUnauthorizedDepositScope,
			expectedCode: codes.PermissionDenied,
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

