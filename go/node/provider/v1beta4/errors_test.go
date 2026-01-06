package v1beta4_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1beta4 "pkg.akt.dev/go/node/provider/v1beta4"
)

func TestErrorGRPCStatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedCode codes.Code
	}{
		{
			name:         "provider_not_found_returns_not_found",
			err:          v1beta4.ErrProviderNotFound,
			expectedCode: codes.NotFound,
		},
		{
			name:         "provider_exists_returns_already_exists",
			err:          v1beta4.ErrProviderExists,
			expectedCode: codes.AlreadyExists,
		},
		{
			name:         "invalid_provider_uri_returns_invalid_argument",
			err:          v1beta4.ErrInvalidProviderURI,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "not_abs_provider_uri_returns_invalid_argument",
			err:          v1beta4.ErrNotAbsProviderURI,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_address_returns_invalid_argument",
			err:          v1beta4.ErrInvalidAddress,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "attributes_returns_invalid_argument",
			err:          v1beta4.ErrAttributes,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "invalid_info_website_returns_invalid_argument",
			err:          v1beta4.ErrInvalidInfoWebsite,
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "incompatible_attributes_returns_failed_precondition",
			err:          v1beta4.ErrIncompatibleAttributes,
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

