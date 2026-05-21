package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyCmdValidatesEvidenceAndHashFile(t *testing.T) {
	dir, expectedHash := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"verify", dir})

	require.NoError(t, root.Execute())
	require.Equal(t, "verify=pass\nevidence_hash="+expectedHash+"\n", out.String())
}

func TestVerifyCmdRejectsMissingHashFile(t *testing.T) {
	dir, _ := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	require.NoError(t, os.Remove(filepath.Join(dir, evidenceHashFile)))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "read "+evidenceHashFile)
}

func TestVerifyCmdRejectsMismatchedHashFile(t *testing.T) {
	dir, _ := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	wrongHash := "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(wrongHash+"\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" mismatch")
}

func TestVerifyCmdRejectsInvalidHashFile(t *testing.T) {
	dir, _ := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte("not-a-hash\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" must contain sha256:<64 hex>")
}

func TestVerifyCmdRejectsSnapshotPayloadMismatch(t *testing.T) {
	dir, _ := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	require.NoError(t, os.WriteFile(filepath.Join(dir, snapshotPayloadFile), []byte("different payload"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "snapshot hash mismatch between evidence and "+snapshotPayloadFile)
}

func TestVerifyCmdRejectsNonceMismatch(t *testing.T) {
	dir, _ := writeCollectedEvidenceArtifact(t, func(*EvidenceDocument) {})
	require.NoError(t, os.WriteFile(filepath.Join(dir, nonceFile), []byte("different nonce"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "inventory nonce mismatch between evidence and "+nonceFile)
}
