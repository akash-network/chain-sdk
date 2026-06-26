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

func TestVerificationQueryCommandSurface(t *testing.T) {
	cmd := QueryCmd()
	verification := requireSubcommand(t, cmd, "verification")

	for _, name := range []string{
		"params",
		"auditor",
		"auditors",
		"attestation",
		"provider-attestations",
		"auditor-attestations",
		"discrepancy",
		"discrepancies",
		"audit-escrow",
		"provider-audit-escrows",
		"provider-grace",
		"provider-bond",
		"provider-snapshot",
	} {
		requireSubcommand(t, verification, name)
	}
}

func TestVerificationTxCommandSurface(t *testing.T) {
	cmd := TxCmd()
	verification := requireSubcommand(t, cmd, "verification")

	for _, name := range []string{
		"register-auditor",
		"renew-auditor",
		"remove-auditor",
		"post-auditor-bond",
		"resign-auditor",
		"post-provider-bond",
		"withdraw-provider-bond",
		"post-snapshot-hash",
		"open-audit-escrow",
		"cancel-audit-escrow",
		"settle-audit-escrow",
		"submit-attestation",
		"revoke-attestation",
		"remove-attestation",
		"revoke-provider-attestation",
		"revoke-all-provider-attestations",
		"revoke-auditor-attestations",
		"resolve-discrepancy",
		"slash-provider-bond",
		"update-params",
	} {
		requireSubcommand(t, verification, name)
	}
}

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

func TestParseAuditEscrowSettlementReason(t *testing.T) {
	testCases := []struct {
		input string
		want  types.AuditEscrowSettlementReason
	}{
		{"cancelled-unconsumed", types.AuditEscrowSettlementReasonCancelledUnconsumed},
		{"expired_unconsumed", types.AuditEscrowSettlementReasonExpiredUnconsumed},
		{"provider-fault", types.AuditEscrowSettlementReasonProviderFault},
		{"provider_fault", types.AuditEscrowSettlementReasonProviderFault},
		{"audit_escrow_settlement_reason_provider_fault", types.AuditEscrowSettlementReasonProviderFault},
		{"no-fault", types.AuditEscrowSettlementReasonNoFault},
		{"audit_escrow_settlement_reason_no_fault", types.AuditEscrowSettlementReasonNoFault},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseAuditEscrowSettlementReason(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseAuditEscrowSettlementReason("unknown")
	require.Error(t, err)
}

func TestParseGovernanceAttestationReason(t *testing.T) {
	testCases := []struct {
		input string
		want  types.GovernanceAttestationReason
	}{
		{"fraudulent-provider", types.GovernanceAttestationReasonFraudulentProvider},
		{"compromised_provider", types.GovernanceAttestationReasonCompromisedProvider},
		{"governance_attestation_reason_provider_non_cooperation", types.GovernanceAttestationReasonProviderNonCooperation},
		{"faulty_auditor", types.GovernanceAttestationReasonFaultyAuditor},
		{"negligent_auditor", types.GovernanceAttestationReasonNegligentAuditor},
		{"evidence_insufficient", types.GovernanceAttestationReasonEvidenceInsufficient},
		{"emergency_safety_action", types.GovernanceAttestationReasonEmergencySafetyAction},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseGovernanceAttestationReason(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseGovernanceAttestationReason("unknown")
	require.Error(t, err)
}

func TestParseDiscrepancyResolutionReason(t *testing.T) {
	testCases := []struct {
		input string
		want  types.DiscrepancyResolutionReason
	}{
		{"auditor-a-correct", types.DiscrepancyResolutionReasonAuditorACorrect},
		{"auditor_b_correct", types.DiscrepancyResolutionReasonAuditorBCorrect},
		{"discrepancy_resolution_reason_both_auditors_wrong", types.DiscrepancyResolutionReasonBothAuditorsWrong},
		{"provider_fault", types.DiscrepancyResolutionReasonProviderFault},
		{"shared_fault", types.DiscrepancyResolutionReasonSharedFault},
		{"evidence_inconclusive", types.DiscrepancyResolutionReasonEvidenceInconclusive},
		{"governance_timeout_review", types.DiscrepancyResolutionReasonGovernanceTimeoutReview},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseDiscrepancyResolutionReason(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseDiscrepancyResolutionReason("unknown")
	require.Error(t, err)
}

func TestParseProviderBondSlashReason(t *testing.T) {
	testCases := []struct {
		input string
		want  types.ProviderBondSlashReason
	}{
		{"resource-misrepresentation", types.ProviderBondSlashReasonResourceMisrepresentation},
		{"capacity_overstatement", types.ProviderBondSlashReasonCapacityOverstatement},
		{"provider_bond_slash_reason_fraudulent_snapshot", types.ProviderBondSlashReasonFraudulentSnapshot},
		{"provider_compromise", types.ProviderBondSlashReasonProviderCompromise},
		{"sla_breach", types.ProviderBondSlashReasonSLABreach},
		{"non_cooperation_during_audit", types.ProviderBondSlashReasonNonCooperationDuringAudit},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseProviderBondSlashReason(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseProviderBondSlashReason("unknown")
	require.Error(t, err)
}

func TestParseAttestationRevocationReason(t *testing.T) {
	testCases := []struct {
		input string
		want  types.AttestationRevocationReason
	}{
		{"provider-no-longer-qualifies", types.AttestationRevocationReasonProviderNoLongerQualifies},
		{"snapshot_mismatch", types.AttestationRevocationReasonSnapshotMismatch},
		{"attestation_revocation_reason_software_identity_changed", types.AttestationRevocationReasonSoftwareIdentityChanged},
		{"capability_misrepresented", types.AttestationRevocationReasonCapabilityMisrepresented},
		{"provider_non_responsive", types.AttestationRevocationReasonProviderNonResponsive},
		{"auditor_evidence_error", types.AttestationRevocationReasonAuditorEvidenceError},
		{"auditor_operational_exit", types.AttestationRevocationReasonAuditorOperationalExit},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseAttestationRevocationReason(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseAttestationRevocationReason("unknown")
	require.Error(t, err)
}

func TestParseFaultAttribution(t *testing.T) {
	testCases := []struct {
		input string
		want  types.FaultAttribution
	}{
		{"provider-fault", types.FaultAttributionProviderFault},
		{"auditor_fault", types.FaultAttributionAuditorFault},
		{"fault_attribution_shared_fault", types.FaultAttributionSharedFault},
		{"no-fault", types.FaultAttributionNoFault},
		{"inconclusive", types.FaultAttributionInconclusive},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseFaultAttribution(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseFaultAttribution("unknown")
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

	_, err = parseHexHash(hex.EncodeToString([]byte("too short")))
	require.ErrorContains(t, err, "hash must be 32 bytes")
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

func requireSubcommand(t *testing.T, cmd *cobra.Command, name string) *cobra.Command {
	t.Helper()

	for _, subcmd := range cmd.Commands() {
		if subcmd.Name() == name {
			return subcmd
		}
	}

	require.Failf(t, "missing command", "%s under %s", name, cmd.CommandPath())
	return nil
}
