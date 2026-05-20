package cli

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/verification/v1"
)

func TestParseVerificationTier(t *testing.T) {
	testCases := []struct {
		input string
		want  types.VerificationTier
	}{
		{"L1", types.TierIdentified},
		{"identified", types.TierIdentified},
		{"verification_tier_verified", types.TierVerified},
		{"established", types.TierEstablished},
		{"trusted", types.TierTrusted},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseVerificationTier(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseVerificationTier("unknown")
	require.Error(t, err)
}

func TestParseCapabilityFlag(t *testing.T) {
	testCases := []struct {
		input string
		want  types.CapabilityFlag
	}{
		{"tee", types.CapabilityTEEHardwareAttestation},
		{"confidential-computing", types.CapabilityConfidentialComputing},
		{"persistent_storage", types.CapabilityPersistentStorage},
		{"capability_bare_metal", types.CapabilityBareMetal},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseCapabilityFlag(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseCapabilityFlag("unknown")
	require.Error(t, err)
}

func TestParseHexHash(t *testing.T) {
	raw := []byte("12345678901234567890123456789012")
	hexVal := hex.EncodeToString(raw)

	testCases := []string{
		hexVal,
		"0x" + hexVal,
		"sha256:" + hexVal,
	}

	for _, tc := range testCases {
		t.Run(tc, func(t *testing.T) {
			got, err := parseHexHash(tc)
			require.NoError(t, err)
			require.Equal(t, raw, got)
		})
	}

	_, err := parseHexHash("not-hex")
	require.Error(t, err)
}

func TestReadResourceSummaryFlag(t *testing.T) {
	inline := `{"total_vcpus":24,"total_memory_mb":"1024","total_storage_mb":"2048","software_version":"test"}`

	cmd := newResourceSummaryTestCmd(inline)
	summary, err := readResourceSummaryFlag(cmd)
	require.NoError(t, err)
	require.Equal(t, uint32(24), summary.TotalVCPUs)
	require.Equal(t, uint64(1024), summary.TotalMemoryMB)
	require.Equal(t, uint64(2048), summary.TotalStorageMB)
	require.Equal(t, "test", summary.SoftwareVersion)

	path := filepath.Join(t.TempDir(), "summary.json")
	require.NoError(t, os.WriteFile(path, []byte(inline), 0o600))

	cmd = newResourceSummaryTestCmd(path)
	summary, err = readResourceSummaryFlag(cmd)
	require.NoError(t, err)
	require.Equal(t, uint32(24), summary.TotalVCPUs)
}

func newResourceSummaryTestCmd(val string) *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().String(flagResourceSummary, "", "")
	_ = cmd.Flags().Set(flagResourceSummary, val)
	return cmd
}
