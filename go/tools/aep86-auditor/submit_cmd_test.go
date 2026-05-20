package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubmitCmdDryRunDerivesAkashTxCommand(t *testing.T) {
	dir, expectedHash := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetArgs([]string{
		"submit",
		"--fee", "100uakt",
		"--deposit", "200uakt",
		dir,
	})

	require.NoError(t, root.Execute())
	require.Equal(t,
		"submit=dry-run\n"+
			"broadcast=not_implemented\n"+
			"evidence_hash="+expectedHash+"\n"+
			"command=akash tx verification submit-attestation --provider akash1provider --audit-escrow-id 7 --tier L1 --capabilities persistent_storage --evidence-hash "+expectedHash+" --fee 100uakt --deposit 200uakt --from akash1auditor --chain-id akash-local\n",
		out.String(),
	)
}

func TestSubmitCmdAcceptsMatchingSubmitFlags(t *testing.T) {
	dir, _ := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{
		"submit",
		"--provider", "akash1provider",
		"--auditor", "akash1auditor",
		"--audit-escrow-id", "7",
		"--tier", "identified",
		"--capability", "persistent_storage",
		"--fee", "100uakt",
		"--deposit", "200uakt",
		"--chain-id", "akash-local",
		dir,
	})

	require.NoError(t, root.Execute())
}

func TestSubmitCmdRejectsMismatchedEvidenceHash(t *testing.T) {
	dir, _ := writeCanonicalEvidenceArtifact(t)
	wrongHash := "sha256:ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(wrongHash+"\n"), 0o644))

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"submit", "--fee", "100uakt", "--deposit", "200uakt", dir})

	err := root.Execute()
	require.ErrorContains(t, err, evidenceHashFile+" mismatch")
}

func TestSubmitCmdRejectsMismatchedProviderFlag(t *testing.T) {
	dir, _ := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{
		"submit",
		"--provider", "akash1other",
		"--fee", "100uakt",
		"--deposit", "200uakt",
		dir,
	})

	err := root.Execute()
	require.ErrorContains(t, err, `provider flag "akash1other" does not match evidence provider "akash1provider"`)
}

func TestSubmitCmdRejectsInvalidFee(t *testing.T) {
	dir, _ := writeVerifiedEvidenceArtifact(t)

	root := newRootCmd()
	root.SetOut(&bytes.Buffer{})
	root.SetArgs([]string{"submit", "--fee", "not-a-coin", "--deposit", "200uakt", dir})

	err := root.Execute()
	require.ErrorContains(t, err, "invalid fee")
}

func writeVerifiedEvidenceArtifact(t *testing.T) (string, string) {
	t.Helper()

	dir, hash := writeCanonicalEvidenceArtifact(t)
	require.NoError(t, os.WriteFile(filepath.Join(dir, evidenceHashFile), []byte(hash+"\n"), 0o644))

	return dir, hash
}
