package sdl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceCredentialsValidationModes(t *testing.T) {
	tests := []struct {
		name         string
		credentials  v2ServiceCredentials
		confidential bool
		wantError    string
	}{
		{
			name: "ordinary inline",
			credentials: v2ServiceCredentials{
				Host:     "registry.example",
				Username: "tenant",
				Password: "secret-value",
			},
		},
		{
			name:         "confidential reference",
			credentials:  v2ServiceCredentials{URI: "kbs:///lease-scope/registry/auth"},
			confidential: true,
		},
		{
			name:         "confidential inline rejected",
			credentials:  v2ServiceCredentials{Host: "registry.example", Username: "tenant", Password: "secret-value"},
			confidential: true,
			wantError:    "require a KBS resource URI",
		},
		{
			name:        "ordinary reference rejected",
			credentials: v2ServiceCredentials{URI: "kbs:///lease-scope/registry/auth"},
			wantError:   "requires a confidential service",
		},
		{
			name:         "mixed mode rejected",
			credentials:  v2ServiceCredentials{Host: "registry.example", URI: "kbs:///lease-scope/registry/auth"},
			confidential: true,
			wantError:    "cannot mix inline fields with uri",
		},
		{
			name:         "host rejected",
			credentials:  v2ServiceCredentials{URI: "kbs://server/lease-scope/registry/auth"},
			confidential: true,
			wantError:    "canonical kbs:///repo/type/tag",
		},
		{
			name:         "query rejected",
			credentials:  v2ServiceCredentials{URI: "kbs:///lease-scope/registry/auth?version=1"},
			confidential: true,
			wantError:    "canonical kbs:///repo/type/tag",
		},
		{
			name:         "leading dot rejected",
			credentials:  v2ServiceCredentials{URI: "kbs:///.lease-scope/registry/auth"},
			confidential: true,
			wantError:    "canonical kbs:///repo/type/tag",
		},
		{
			name:         "oversized rejected",
			credentials:  v2ServiceCredentials{URI: "kbs:///" + strings.Repeat("a", kbsResourceURIMaxBytes) + "/registry/auth"},
			confidential: true,
			wantError:    "bounded canonical",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.credentials.validate(test.confidential)
			if test.wantError == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, test.wantError)
		})
	}
}
