package v1_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "pkg.akt.dev/go/node/deployment/v1"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{
			name:         "deployment_not_found_returns_not_found",
			err:          v1.ErrDeploymentNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "group_not_found_returns_not_found",
			err:          v1.ErrGroupNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "name_does_not_exist_returns_not_found",
			err:          v1.ErrNameDoesNotExist,
			expectedCode: codes.NotFound,
		},
		{
			name:         "deployment_exists_returns_already_exists",
			err:          v1.ErrDeploymentExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "invalid_request_returns_invalid_argument",
			err:          v1.ErrInvalidRequest,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_groups_returns_invalid_argument",
			err:          v1.ErrInvalidGroups,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_deployment_id_returns_invalid_argument",
			err:          v1.ErrInvalidDeploymentID,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "empty_hash_returns_invalid_argument",
			err:          v1.ErrEmptyHash,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_hash_returns_invalid_argument",
			err:          v1.ErrInvalidHash,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_deployment_returns_invalid_argument",
			err:          v1.ErrInvalidDeployment,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_group_id_returns_invalid_argument",
			err:          v1.ErrInvalidGroupID,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_deposit_returns_invalid_argument",
			err:          v1.ErrInvalidDeposit,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_id_path_returns_invalid_argument",
			err:          v1.ErrInvalidIDPath,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_param_returns_invalid_argument",
			err:          v1.ErrInvalidParam,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_escrow_id_returns_invalid_argument",
			err:          v1.ErrInvalidEscrowID,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "owner_account_missing_returns_invalid_argument",
			err:          v1.ErrOwnerAcctMissing,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "group_spec_invalid_returns_invalid_argument",
			err:          v1.ErrGroupSpecInvalid,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "deployment_closed_returns_failed_precondition",
			err:          v1.ErrDeploymentClosed,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "group_closed_returns_failed_precondition",
			err:          v1.ErrGroupClosed,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "group_open_returns_failed_precondition",
			err:          v1.ErrGroupOpen,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "group_paused_returns_failed_precondition",
			err:          v1.ErrGroupPaused,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "group_not_open_returns_failed_precondition",
			err:          v1.ErrGroupNotOpen,
			expectedCode: codes.FailedPrecondition,
		},
		{
			name:         "internal_error_returns_internal",
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
