package sdl

import (
	"testing"

	"github.com/stretchr/testify/require"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
)

func TestV2_1_KBSSelectionProjectsToManifest(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		validate func(*testing.T, *manifest.KBSParams)
	}{
		{
			name: "provider managed",
			path: "./_testdata/v2.1-tee-kbs-provider.yaml",
			validate: func(t *testing.T, params *manifest.KBSParams) {
				require.NotNil(t, params.GetProvider())
				require.Nil(t, params.GetTenant())
			},
		},
		{
			name: "tenant managed",
			path: "./_testdata/v2.1-tee-kbs-tenant.yaml",
			validate: func(t *testing.T, params *manifest.KBSParams) {
				tenant := params.GetTenant()
				require.NotNil(t, tenant)
				require.Equal(t, "https://kbs.tenant.example:8443", tenant.URL)
				require.Contains(t, tenant.Certificate, "tenant-public-ca-fixture")
				require.Equal(t, "kbs:///tenant/security-policy/sha256-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", tenant.ImageSecurityPolicyURI)
				require.Contains(t, tenant.AgentPolicy, "default allow = false")
				require.Nil(t, params.GetProvider())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sdl, err := ReadFile(test.path)
			require.NoError(t, err)
			groups, err := sdl.Manifest()
			require.NoError(t, err)
			params := groups.GetGroups()[0].Services[0].Params.TEE.KBS
			require.NotNil(t, params)
			test.validate(t, params)
		})
	}
}

func TestV2ServiceKBSParamsRejectsInvalidCombinations(t *testing.T) {
	validTenant := v2ServiceKBSParams{
		Mode:                   kbsModeTenant,
		URL:                    "https://kbs.tenant.example",
		Certificate:            "tenant public certificate",
		ImageSecurityPolicyURI: "kbs:///tenant/security-policy/sha256-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		AgentPolicy:            "package agent_policy\n\ndefault allow = false\n",
	}
	tests := []struct {
		name      string
		params    v2ServiceKBSParams
		hasTEE    bool
		wantError string
	}{
		{name: "provider managed", params: v2ServiceKBSParams{Mode: kbsModeProvider}, hasTEE: true},
		{name: "tenant managed", params: validTenant, hasTEE: true},
		{name: "requires tee", params: v2ServiceKBSParams{Mode: kbsModeProvider}, wantError: "requires tee"},
		{name: "requires mode", params: v2ServiceKBSParams{}, hasTEE: true, wantError: "mode"},
		{name: "rejects unknown mode", params: v2ServiceKBSParams{Mode: "other"}, hasTEE: true, wantError: "mode"},
		{
			name: "provider rejects tenant fields",
			params: v2ServiceKBSParams{
				Mode: kbsModeProvider,
				URL:  "https://kbs.tenant.example",
			},
			hasTEE:    true,
			wantError: "provider mode",
		},
		{
			name: "tenant requires complete bundle",
			params: v2ServiceKBSParams{
				Mode: kbsModeTenant,
				URL:  "https://kbs.tenant.example",
			},
			hasTEE:    true,
			wantError: "complete public configuration",
		},
		{
			name: "tenant rejects non-origin URL",
			params: func() v2ServiceKBSParams {
				params := validTenant
				params.URL = "https://kbs.tenant.example/admin"
				return params
			}(),
			hasTEE:    true,
			wantError: "canonical HTTPS origin",
		},
		{
			name: "tenant rejects mutable policy URI",
			params: func() v2ServiceKBSParams {
				params := validTenant
				params.ImageSecurityPolicyURI = "kbs:///tenant/security-policy/latest"
				return params
			}(),
			hasTEE:    true,
			wantError: "content addressed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.params.validate(test.hasTEE)
			if test.wantError == "" {
				require.NoError(t, err)
				return
			}
			require.ErrorContains(t, err, test.wantError)
		})
	}
}

func TestV2ServiceRequiresExplicitKBSSelectionForReferencedData(t *testing.T) {
	tests := []struct {
		name    string
		service v2Service
	}{
		{
			name: "registry credentials",
			service: v2Service{
				Credentials: &v2ServiceCredentials{URI: "kbs:///lease-scope/registry/auth"},
			},
		},
		{
			name:    "sealed environment",
			service: v2Service{Env: []string{"SECRET=sealed.header.payload.signature"}},
		},
		{
			name: "persistent storage key",
			service: v2Service{Params: &v2ServiceParams{Storage: map[string]v2ServiceStorageParams{
				"data": {KeyRef: "sealed.header.payload.signature"},
			}}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.ErrorContains(t, test.service.validateKBS(), "explicit params.kbs selection")

			if test.service.Params == nil {
				test.service.Params = &v2ServiceParams{}
			}
			test.service.Params.TEE = TEETypeCPU
			test.service.Params.KBS = &v2ServiceKBSParams{Mode: kbsModeProvider}
			require.NoError(t, test.service.validateKBS())
		})
	}
}
