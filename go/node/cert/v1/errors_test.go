package v1_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "pkg.akt.dev/go/node/cert/v1"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{
			name:         "certificate_not_found_returns_not_found",
			err:          v1.ErrCertificateNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "certificate_exists_returns_already_exists",
			err:          v1.ErrCertificateExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "invalid_address_returns_invalid_argument",
			err:          v1.ErrInvalidAddress,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_serial_number_returns_invalid_argument",
			err:          v1.ErrInvalidSerialNumber,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_certificate_value_returns_invalid_argument",
			err:          v1.ErrInvalidCertificateValue,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_pubkey_value_returns_invalid_argument",
			err:          v1.ErrInvalidPubkeyValue,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_state_returns_invalid_argument",
			err:          v1.ErrInvalidState,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_key_size_returns_invalid_argument",
			err:          v1.ErrInvalidKeySize,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "certificate_already_revoked_returns_failed_precondition",
			err:          v1.ErrCertificateAlreadyRevoked,
			expectedCode: codes.FailedPrecondition,
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

