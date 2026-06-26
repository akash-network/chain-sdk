package sdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	verificationv1 "pkg.akt.dev/go/node/verification/v1"
)

func TestV2_1_ParseVerification(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple-verification.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	v := groups[0].Requirements.Verification
	require.NotNil(t, v, "verification block must propagate to Requirements.Verification")

	assert.Equal(t, verificationv1.TierEstablished, v.MinTier)
	assert.Equal(t, verificationv1.AuditorSelectionModeAny, v.AuditorMode)
	require.Len(t, v.RequiredCapabilities, 1)
	assert.Equal(t, verificationv1.CapabilityTEEHardwareAttestation, v.RequiredCapabilities[0])
	require.Len(t, v.RequiredAuditors, 1)
	assert.Equal(t, "akash1auditor1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", v.RequiredAuditors[0])
	assert.Equal(t, uint32(0), v.MinAuditorCount)
}

func TestV2_1_NoVerificationBlockBackwardCompat(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	assert.Nil(t, groups[0].Requirements.Verification,
		"SDLs without a verification block must produce a nil Verification field")
}

func TestV2_NoVerificationBlockBackwardCompat(t *testing.T) {
	sdl, err := ReadFile("./_testdata/simple.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	assert.Nil(t, groups[0].Requirements.Verification,
		"v2.0 SDLs without a verification block must produce a nil Verification field")
}

func TestV2Verification_YAML_InvalidMinTier(t *testing.T) {
	in := []byte(`
min_tier: 5
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "min_tier")
}

func TestV2Verification_YAML_NegativeMinTier(t *testing.T) {
	in := []byte(`
min_tier: -1
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "min_tier")
}

func TestV2Verification_YAML_UnknownCapability(t *testing.T) {
	in := []byte(`
min_tier: 2
capabilities:
  - tee_hardware_attestation
  - not_a_real_capability
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not_a_real_capability")
}

func TestV2Verification_YAML_InvalidAuditorMode(t *testing.T) {
	in := []byte(`
min_tier: 2
auditor_mode: maybe
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "auditor_mode")
}

func TestV2Verification_YAML_EmptyAuditor(t *testing.T) {
	in := []byte(`
min_tier: 2
auditors:
  - ""
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "auditors")
}

func TestV2Verification_YAML_InvalidAuditorPrefix(t *testing.T) {
	in := []byte(`
min_tier: 2
auditors:
  - cosmos1notakashaddress
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "prefix")
}

func TestV2Verification_YAML_AuditorSurroundingWhitespace(t *testing.T) {
	in := []byte(`
min_tier: 2
auditors:
  - " akash1auditor1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx "
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.ErrorIs(t, err, errSDLInvalid)
	assert.Contains(t, err.Error(), "whitespace")
}

func TestV2Verification_YAML_Valid_Defaults(t *testing.T) {
	in := []byte(`
min_tier: 0
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.NoError(t, err)
	assert.Equal(t, int32(0), v.MinTier)
	assert.Equal(t, "", v.AuditorMode)
	assert.Empty(t, v.Capabilities)
	assert.Empty(t, v.Auditors)
}

func TestV2Verification_YAML_TierUnspecifiedRejectsDependentFilters(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
	}{
		{
			name: "capability without tier",
			in: []byte(`
capabilities:
  - tee_hardware_attestation
`),
		},
		{
			name: "capability with explicit L0",
			in: []byte(`
min_tier: 0
capabilities:
  - tee_hardware_attestation
`),
		},
		{
			name: "auditor with explicit L0",
			in: []byte(`
min_tier: 0
auditors:
  - akash1auditor1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
`),
		},
		{
			name: "minimum auditor count with explicit L0",
			in: []byte(`
min_tier: 0
min_auditor_count: 1
`),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var v v2Verification
			err := yaml.Unmarshal(tc.in, &v)
			require.Error(t, err)
			assert.ErrorIs(t, err, errSDLInvalid)
			assert.Contains(t, err.Error(), "min_tier")
		})
	}
}

func TestV2Verification_toProto_NilReceiver(t *testing.T) {
	var v *v2Verification
	assert.Nil(t, v.toProto())
}

func TestV2Verification_toProto_MapsAll(t *testing.T) {
	v := &v2Verification{
		MinTier:         2,
		Capabilities:    []string{"tee_hardware_attestation", "confidential_computing"},
		Auditors:        []string{"akash1abc", "akash1def"},
		AuditorMode:     "all",
		MinAuditorCount: 3,
	}

	got := v.toProto()
	require.NotNil(t, got)
	assert.Equal(t, verificationv1.TierVerified, got.MinTier)
	assert.Equal(t, verificationv1.AuditorSelectionModeAll, got.AuditorMode)
	assert.Equal(t, uint32(3), got.MinAuditorCount)
	assert.Equal(t, []verificationv1.CapabilityFlag{
		verificationv1.CapabilityTEEHardwareAttestation,
		verificationv1.CapabilityConfidentialComputing,
	}, got.RequiredCapabilities)
	assert.Equal(t, []string{"akash1abc", "akash1def"}, got.RequiredAuditors)
}

func TestV2Verification_toProto_DefaultAuditorMode(t *testing.T) {
	v := &v2Verification{MinTier: 1}
	got := v.toProto()
	require.NotNil(t, got)
	assert.Equal(t, verificationv1.AuditorSelectionModeAny, got.AuditorMode)
}

func TestV2Verification_toProto_CollapseVacuous(t *testing.T) {
	cases := []struct {
		name string
		v    *v2Verification
	}{
		{name: "zero struct", v: &v2Verification{}},
		{name: "explicit min_tier 0", v: &v2Verification{MinTier: 0}},
		{name: "auditor_mode alone is no filter", v: &v2Verification{AuditorMode: "all"}},
		{name: "auditor_mode any is no filter", v: &v2Verification{AuditorMode: "any"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Nil(t, tc.v.toProto())
		})
	}
}

func TestV2Verification_toProto_DoesNotCollapseFiltering(t *testing.T) {
	cases := []struct {
		name string
		v    *v2Verification
	}{
		{name: "min_tier only", v: &v2Verification{MinTier: 1}},
		{name: "tier with capability", v: &v2Verification{MinTier: 1, Capabilities: []string{"tee_hardware_attestation"}}},
		{name: "tier with auditor", v: &v2Verification{MinTier: 1, Auditors: []string{"akash1abc"}}},
		{name: "tier with min_auditor_count", v: &v2Verification{MinTier: 1, MinAuditorCount: 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.NotNil(t, tc.v.toProto())
		})
	}
}
