package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRevokeCmdDryRunDerivesAkashTxCommand(t *testing.T) {
	dir, expectedHash := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "software_identity_changed"
		evidence.FaultContext.FaultAttribution = "provider_fault"
		evidence.SustainedValidation.Status = "fail"
	})

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{
		"revoke",
		"--reason", "software-identity-changed",
		dir,
	})

	require.NoError(t, root.Execute())
	require.Equal(t,
		"revoke=dry-run\n"+
			"broadcast=not_implemented\n"+
			"evidence_hash="+expectedHash+"\n"+
			"command=akash tx verification revoke-attestation --provider akash1provider --reason software_identity_changed --evidence-hash "+expectedHash+" --from akash1auditor --chain-id akash-local\n",
		out.String(),
	)
}

func TestRevokeCmdAcceptsMatchingFlags(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "provider_no_longer_qualifies"
		evidence.FaultContext.FaultAttribution = "no_fault"
		evidence.SustainedValidation.Status = "fail"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{
		"revoke",
		"--provider", "akash1provider",
		"--auditor", "akash1auditor",
		"--reason", "provider_no_longer_qualifies",
		"--chain-id", "akash-local",
		dir,
	})

	require.NoError(t, root.Execute())
}

func TestRevokeCmdRejectsMismatchedEvidenceHash(t *testing.T) {
	dir, _ := writeCanonicalEvidenceArtifact(t)
	wrongHash := "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(wrongHash+"\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "snapshot_mismatch", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" mismatch")
}

func TestRevokeCmdRejectsInvalidReason(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "snapshot_mismatch"
		evidence.FaultContext.FaultAttribution = "provider_fault"
		evidence.SustainedValidation.Status = "fail"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "provider_fault", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "invalid revocation reason")
}

func TestRevokeCmdRejectsEvidenceWithoutReason(t *testing.T) {
	dir, _ := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "snapshot_mismatch", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "revocation evidence must set fault_context.reason")
}

func TestRevokeCmdRejectsFailedCheckWithoutEvidenceReason(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.Checks[0].Status = "fail"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "snapshot_mismatch", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "revocation evidence must set fault_context.reason")
}

func TestRevokeCmdRejectsMismatchedFaultAttribution(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "auditor_evidence_error"
		evidence.FaultContext.FaultAttribution = "provider_fault"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "auditor_evidence_error", dir})

	err := root.Execute()
	require.ErrorContains(t, err, `fault_context.fault_attribution "provider_fault" is not valid for revocation reason "auditor_evidence_error"`)
}

func TestRevokeCmdRejectsProviderReasonWithoutSustainedFailure(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "snapshot_mismatch"
		evidence.FaultContext.FaultAttribution = "provider_fault"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "snapshot_mismatch", dir})

	err := root.Execute()
	require.ErrorContains(t, err, `provider-side revocation evidence must have sustained_validation.status "fail"`)
}

func TestRevokeCmdAcceptsAuditorOperationalExitWithoutSustainedFailure(t *testing.T) {
	dir, _ := writeEvidenceArtifactWithHash(t, func(evidence *EvidenceDocument) {
		evidence.FaultContext.Reason = "auditor_operational_exit"
		evidence.FaultContext.FaultAttribution = "no_fault"
	})

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "auditor_operational_exit", dir})

	require.NoError(t, root.Execute())
}

func TestRevokeCmdRejectsMismatchedEvidenceReason(t *testing.T) {
	dir := t.TempDir()
	evidence := validEvidenceDocument()
	evidence.FaultContext.Reason = "snapshot_mismatch"
	evidence.FaultContext.FaultAttribution = "provider_fault"
	evidence.SustainedValidation.Status = "fail"
	raw, hash, err := marshalEvidenceCanonical(evidence)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(hash+"\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"revoke", "--reason", "software_identity_changed", dir})

	err = root.Execute()
	require.ErrorContains(t, err, `reason flag "software_identity_changed" does not match evidence fault_context.reason "snapshot_mismatch"`)
}
