package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyCmdValidatesEvidenceAndHashFile(t *testing.T) {
	dir, expectedHash := writeCanonicalEvidenceArtifact(t)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(expectedHash+"\n"), 0o644))

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{"verify", dir})

	require.NoError(t, root.Execute())
	require.Equal(t, "verify=pass\nevidence_hash="+expectedHash+"\n", out.String())
}

func TestVerifyCmdRejectsMissingHashFile(t *testing.T) {
	dir, _ := writeCanonicalEvidenceArtifact(t)

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "read "+evidenceHashFile)
}

func TestVerifyCmdRejectsMismatchedHashFile(t *testing.T) {
	dir, _ := writeCanonicalEvidenceArtifact(t)
	wrongHash := "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(wrongHash+"\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" mismatch")
}

func TestVerifyCmdRejectsInvalidHashFile(t *testing.T) {
	dir, _ := writeCanonicalEvidenceArtifact(t)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte("not-a-hash\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"verify", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" must contain sha256:<64 hex>")
}
