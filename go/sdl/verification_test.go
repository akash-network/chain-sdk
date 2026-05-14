package sdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	verificationv1 "pkg.akt.dev/go/node/verification/v1"
)

// TestV2_1_ParseVerification asserts that a v2.1 SDL containing a
// `verification:` block under a placement profile is decoded end-to-end
// into a populated `PlacementRequirements.Verification` proto on the
// resulting deployment group.
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

// TestV2_1_NoVerificationBlockBackwardCompat asserts that an SDL that omits
// the verification block continues to parse and produces a nil Verification
// pointer on the deployment group, preserving backward compatibility with
// pre-AEP-86 SDLs.
func TestV2_1_NoVerificationBlockBackwardCompat(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	assert.Nil(t, groups[0].Requirements.Verification,
		"SDLs without a verification block must produce a nil Verification field")
}

// TestV2_NoVerificationBlockBackwardCompat verifies the same backward-compat
// guarantee for v2.0 SDLs (which share `v2ProfilePlacement` with v2.1).
func TestV2_NoVerificationBlockBackwardCompat(t *testing.T) {
	sdl, err := ReadFile("./_testdata/simple.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	assert.Nil(t, groups[0].Requirements.Verification,
		"v2.0 SDLs without a verification block must produce a nil Verification field")
}

// TestV2Verification_YAML_InvalidMinTier exercises the UnmarshalYAML path
// directly with an out-of-range min_tier value.
func TestV2Verification_YAML_InvalidMinTier(t *testing.T) {
	in := []byte(`
min_tier: 5
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "min_tier")
}

// TestV2Verification_YAML_NegativeMinTier checks the lower bound of the
// validation range.
func TestV2Verification_YAML_NegativeMinTier(t *testing.T) {
	in := []byte(`
min_tier: -1
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "min_tier")
}

// TestV2Verification_YAML_UnknownCapability ensures an unknown capability
// name is rejected with the offending name surfaced in the error.
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

// TestV2Verification_YAML_InvalidAuditorMode checks the auditor_mode allow-list.
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

// TestV2Verification_YAML_EmptyAuditor rejects a blank auditor entry.
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

// TestV2Verification_YAML_InvalidAuditorPrefix rejects an auditor address
// that doesn't start with the expected bech32 HRP.
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

// TestV2Verification_YAML_Valid_Defaults exercises the happy-path defaults:
// no auditor_mode (defaults to any), no capabilities, no auditors, min_tier=0.
func TestV2Verification_YAML_Valid_Defaults(t *testing.T) {
	in := []byte(`
min_tier: 0
`)
	var v v2Verification
	err := yaml.Unmarshal(in, &v)
	require.NoError(t, err)
	assert.Equal(t, 0, v.MinTier)
	assert.Equal(t, "", v.AuditorMode)
	assert.Empty(t, v.Capabilities)
	assert.Empty(t, v.Auditors)
}

// TestV2Verification_toProto_NilReceiver ensures the helper is safe to call
// on a nil receiver — the contract relied on by groupBuilder_v2*.go.
func TestV2Verification_toProto_NilReceiver(t *testing.T) {
	var v *v2Verification
	assert.Nil(t, v.toProto())
}

// TestV2Verification_toProto_MapsAll verifies that every field of the SDL
// struct maps onto the right proto field, including the empty-mode default
// to Any.
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

// TestV2Verification_toProto_DefaultAuditorMode confirms that an empty
// AuditorMode string is materialized as AuditorSelectionModeAny, which is
// the documented default behavior.
func TestV2Verification_toProto_DefaultAuditorMode(t *testing.T) {
	v := &v2Verification{MinTier: 1}
	got := v.toProto()
	require.NotNil(t, got)
	assert.Equal(t, verificationv1.AuditorSelectionModeAny, got.AuditorMode)
}
