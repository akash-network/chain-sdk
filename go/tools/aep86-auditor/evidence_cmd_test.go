package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvidenceCmdValidatesCanonicalEvidenceAndPrintsHash(t *testing.T) {
	dir, expectedHash := writeCanonicalEvidenceArtifact(t)

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"evidence", dir})

	require.NoError(t, root.Execute())
	require.Equal(t, "evidence_hash="+expectedHash+"\n", out.String())
}

func TestEvidenceCmdWritesHashFile(t *testing.T) {
	dir, expectedHash := writeCanonicalEvidenceArtifact(t)

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"evidence", "--write-hash", dir})

	require.NoError(t, root.Execute())

	hashRaw, err := os.ReadFile(filepath.Join(dir, evidenceHashFile))
	require.NoError(t, err)
	require.Equal(t, expectedHash+"\n", string(hashRaw))
	require.Contains(t, out.String(), "wrote evidence hash to "+filepath.Join(dir, evidenceHashFile)+"\n")
	require.Contains(t, out.String(), "evidence_hash="+expectedHash+"\n")
}

func TestEvidenceCmdRejectsNonCanonicalEvidence(t *testing.T) {
	dir := t.TempDir()
	raw, err := json.MarshalIndent(validEvidenceDocument(), "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"evidence", dir})

	err = root.Execute()
	require.ErrorContains(t, err, "evidence is not canonical JSON")
}

func TestEvidenceCmdRejectsUnsortedAttestedCapabilities(t *testing.T) {
	dir := t.TempDir()
	evidence := validEvidenceDocument()
	evidence.AttestedCapabilities = []string{"persistent_storage", "bare_metal"}
	raw, err := json.Marshal(evidence)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"evidence", dir})

	err = root.Execute()
	require.ErrorContains(t, err, "evidence is not canonical JSON")
}

func TestEvidenceCmdRejectsSchemaViolation(t *testing.T) {
	dir := t.TempDir()
	evidence := validEvidenceDocument()
	evidence.Software.BinaryHash = "not-a-sha256-ref"
	raw, err := json.Marshal(evidence)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"evidence", dir})

	err = root.Execute()
	require.ErrorContains(t, err, "evidence schema validation failed")
	require.ErrorContains(t, err, "software.binary_hash")
}

func writeCanonicalEvidenceArtifact(t *testing.T) (string, string) {
	t.Helper()

	dir := t.TempDir()
	raw, hash, err := marshalEvidenceCanonical(validEvidenceDocument())
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceDraftFile), raw, 0o644))

	return dir, hash
}
