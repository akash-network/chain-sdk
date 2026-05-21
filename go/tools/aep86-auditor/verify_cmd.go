package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type verifyConfig struct {
	artifactDir string
}

func newVerifyCmd() *cobra.Command {
	cfg := verifyConfig{}

	cmd := &cobra.Command{
		Use:   "verify <artifact-dir>",
		Short: "Verify collected evidence artifacts offline",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.artifactDir = args[0]
			return runVerify(cmd, cfg)
		},
	}

	return cmd
}

func runVerify(cmd *cobra.Command, cfg verifyConfig) error {
	evidence, hash, err := loadVerifiedEvidenceArtifactDir(cfg.artifactDir)
	if err != nil {
		return err
	}

	if err := validateCollectedArtifactLayout(cfg.artifactDir, evidence); err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), "verify=pass")
	fmt.Fprintf(cmd.OutOrStdout(), "evidence_hash=%s\n", hash)

	return nil
}

func validateEvidenceHashArtifact(artifactDir, expectedHash string) error {
	if artifactDir == "" {
		return fmt.Errorf("artifact-dir is required")
	}

	raw, err := os.ReadFile(filepath.Join(artifactDir, evidenceHashFile))
	if err != nil {
		return fmt.Errorf("read %s: %w", evidenceHashFile, err)
	}

	actualHash := strings.TrimSpace(string(raw))
	if !isSHA256Ref(actualHash) {
		return fmt.Errorf("%s must contain sha256:<64 hex>", evidenceHashFile)
	}
	if actualHash != expectedHash {
		return fmt.Errorf("%s mismatch: got %s, expected %s", evidenceHashFile, actualHash, expectedHash)
	}

	return nil
}
