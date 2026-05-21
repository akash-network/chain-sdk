package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSustainCmdPassesMatchingEvidence(t *testing.T) {
	baselineDir, baselineHash := writeVerifiedEvidenceArtifact(t)
	currentDir, currentHash := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"sustain", baselineDir, currentDir})

	require.NoError(t, root.Execute())
	require.Equal(t,
		"sustained_validation=pass\n"+
			"baseline_hash="+baselineHash+"\n"+
			"current_hash="+currentHash+"\n",
		out.String(),
	)
}

func TestSustainCmdFailsOnSoftwareChange(t *testing.T) {
	baselineDir, _ := writeVerifiedEvidenceArtifact(t)
	currentDir := writeEvidenceArtifact(t, func(evidence *EvidenceDocument) {
		evidence.Software.BinaryHash = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	})

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"sustain", baselineDir, currentDir})

	require.NoError(t, root.Execute())
	require.Contains(t, out.String(), "sustained_validation=fail\n")
	require.Contains(t, out.String(), "reason=software_identity_changed\n")
}

func TestSustainCmdFailsOnCurrentCheckFailure(t *testing.T) {
	baselineDir, _ := writeVerifiedEvidenceArtifact(t)
	currentDir := writeEvidenceArtifact(t, func(evidence *EvidenceDocument) {
		evidence.Checks[0].Status = "fail"
	})

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"sustain", baselineDir, currentDir})

	require.NoError(t, root.Execute())
	require.Contains(t, out.String(), "sustained_validation=fail\n")
	require.Contains(t, out.String(), "reason=current_check_failed\n")
}

func TestSustainCmdWritesRevocationEvidence(t *testing.T) {
	baselineDir, baselineHash := writeVerifiedEvidenceArtifact(t)
	currentDir, currentHash := writeCollectedEvidenceArtifact(t, func(evidence *EvidenceDocument) {
		evidence.Software.BinaryHash = "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	})
	outputDir := t.TempDir()

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"sustain", "--output-dir", outputDir, baselineDir, currentDir})

	require.NoError(t, root.Execute())
	require.Contains(t, out.String(), "sustained_validation=fail\n")
	require.Contains(t, out.String(), "reason=software_identity_changed\n")
	require.Contains(t, out.String(), "sustained_evidence_dir="+outputDir+"\n")

	evidence := readEvidenceDocument(t, filepath.Join(outputDir, evidenceDraftFile))
	require.Equal(t, "fail", evidence.SustainedValidation.Status)
	require.Equal(t, baselineHash, evidence.SustainedValidation.BaselineID)
	require.Equal(t, currentHash, evidence.SustainedValidation.ProofRef)
	require.Equal(t, "software_identity_changed", evidence.FaultContext.Reason)
	require.Equal(t, "provider_fault", evidence.FaultContext.FaultAttribution)
	require.Equal(t, "fail", evidenceChecksByName(evidence.Checks)["sustained_baseline_checked"].Status)

	revoke := newRootCmd()
	revoke.SetOut(&bytes.Buffer{})
	revoke.SetArgs([]string{"revoke", "--reason", "software_identity_changed", outputDir})
	require.NoError(t, revoke.Execute())

	verify := newRootCmd()
	verify.SetOut(&bytes.Buffer{})
	verify.SetArgs([]string{"verify", outputDir})
	require.NoError(t, verify.Execute())
}

func TestSustainCmdMapsFailedSnapshotCheckToRevocationReason(t *testing.T) {
	baselineDir, _ := writeVerifiedEvidenceArtifact(t)
	currentDir, _ := writeCollectedEvidenceArtifact(t, func(evidence *EvidenceDocument) {
		evidence.Checks = append(evidence.Checks, EvidenceCheck{
			Name:     "snapshot_hash_matches_chain",
			Status:   "fail",
			ProofRef: evidence.SnapshotHash,
		})
	})
	outputDir := t.TempDir()

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"sustain", "--output-dir", outputDir, baselineDir, currentDir})

	require.NoError(t, root.Execute())

	revoke := newRootCmd()
	revoke.SetOut(&bytes.Buffer{})
	revoke.SetArgs([]string{"revoke", "--reason", "snapshot_mismatch", outputDir})
	require.NoError(t, revoke.Execute())
}

func TestSustainCmdWritesPassingEvidence(t *testing.T) {
	baselineDir, baselineHash := writeVerifiedEvidenceArtifact(t)
	currentDir, currentHash := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	outputDir := t.TempDir()

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"sustain", "--output-dir", outputDir, baselineDir, currentDir})

	require.NoError(t, root.Execute())

	evidence := readEvidenceDocument(t, filepath.Join(outputDir, evidenceDraftFile))
	require.Equal(t, "pass", evidence.SustainedValidation.Status)
	require.Equal(t, baselineHash, evidence.SustainedValidation.BaselineID)
	require.Equal(t, currentHash, evidence.SustainedValidation.ProofRef)
	require.Equal(t, "unspecified", evidence.FaultContext.Reason)
	require.Equal(t, "unspecified", evidence.FaultContext.FaultAttribution)
}

func writeEvidenceArtifact(t *testing.T, mutate func(*EvidenceDocument)) string {
	t.Helper()

	dir, _ := writeEvidenceArtifactWithHash(t, mutate)

	return dir
}

func writeEvidenceArtifactWithHash(t *testing.T, mutate func(*EvidenceDocument)) (string, string) {
	t.Helper()

	dir := t.TempDir()
	evidence := validEvidenceDocument()
	mutate(&evidence)
	raw, hash, err := marshalEvidenceCanonical(evidence)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(hash+"\n"), 0o644))

	return dir, hash
}

func readEvidenceDocument(t *testing.T, path string) EvidenceDocument {
	t.Helper()

	raw, err := os.ReadFile(path)
	require.NoError(t, err)

	var evidence EvidenceDocument
	require.NoError(t, json.Unmarshal(raw, &evidence))

	return evidence
}
