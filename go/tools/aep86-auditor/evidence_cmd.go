package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	evidenceDraftFile = "evidence.draft.json"
	evidenceHashFile  = "evidence.draft.sha256"
)

type evidenceConfig struct {
	artifactDir string
	writeHash   bool
}

func newEvidenceCmd() *cobra.Command {
	cfg := evidenceConfig{}

	cmd := &cobra.Command{
		Use:   "evidence <artifact-dir>",
		Short: "Validate collected evidence and print its hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.artifactDir = args[0]
			return runEvidence(cmd, cfg)
		},
	}

	cmd.Flags().BoolVar(&cfg.writeHash, "write-hash", false, "Write the canonical evidence hash to evidence.draft.sha256")

	return cmd
}

func runEvidence(cmd *cobra.Command, cfg evidenceConfig) error {
	hash, err := validateEvidenceArtifactDir(cfg.artifactDir)
	if err != nil {
		return err
	}

	if cfg.writeHash {
		hashPath := filepath.Join(cfg.artifactDir, evidenceHashFile)
		if err := os.WriteFile(hashPath, []byte(hash+"\n"), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "wrote evidence hash to %s\n", hashPath)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "evidence_hash=%s\n", hash)

	return nil
}

func validateEvidenceArtifactDir(artifactDir string) (string, error) {
	if artifactDir == "" {
		return "", fmt.Errorf("artifact-dir is required")
	}

	raw, err := os.ReadFile(filepath.Join(artifactDir, evidenceDraftFile))
	if err != nil {
		return "", fmt.Errorf("read %s: %w", evidenceDraftFile, err)
	}

	hash, err := canonicalEvidenceHash(raw)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func canonicalEvidenceHash(raw []byte) (string, error) {
	if err := validateEvidenceBytes(raw); err != nil {
		return "", err
	}

	var evidence EvidenceDocument
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&evidence); err != nil {
		return "", fmt.Errorf("decode evidence: %w", err)
	}

	canonical, hash, err := marshalEvidenceCanonical(evidence)
	if err != nil {
		return "", err
	}
	if !bytes.Equal(raw, canonical) {
		return "", fmt.Errorf("evidence is not canonical JSON")
	}

	return hash, nil
}
