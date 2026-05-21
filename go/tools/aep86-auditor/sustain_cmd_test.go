package main

import (
	"bytes"
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
