package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sustainConfig struct {
	baselineDir string
	currentDir  string
}

func newSustainCmd() *cobra.Command {
	cfg := sustainConfig{}

	cmd := &cobra.Command{
		Use:   "sustain <baseline-artifact-dir> <current-artifact-dir>",
		Short: "Compare current evidence against an attestation baseline",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.baselineDir = args[0]
			cfg.currentDir = args[1]
			return runSustain(cmd, cfg)
		},
	}

	return cmd
}

func runSustain(cmd *cobra.Command, cfg sustainConfig) error {
	baseline, baselineHash, err := loadVerifiedEvidenceArtifactDir(cfg.baselineDir)
	if err != nil {
		return err
	}
	current, currentHash, err := loadVerifiedEvidenceArtifactDir(cfg.currentDir)
	if err != nil {
		return err
	}

	result := compareSustainedEvidence(baseline, current)

	fmt.Fprintf(cmd.OutOrStdout(), "sustained_validation=%s\n", result.status)
	fmt.Fprintf(cmd.OutOrStdout(), "baseline_hash=%s\n", baselineHash)
	fmt.Fprintf(cmd.OutOrStdout(), "current_hash=%s\n", currentHash)
	if result.reason != "" {
		fmt.Fprintf(cmd.OutOrStdout(), "reason=%s\n", result.reason)
	}

	return nil
}

type sustainResult struct {
	status string
	reason string
}

func compareSustainedEvidence(baseline, current EvidenceDocument) sustainResult {
	if baseline.ChainID != current.ChainID {
		return sustainResult{status: "fail", reason: "chain_id_changed"}
	}
	if baseline.Provider != current.Provider {
		return sustainResult{status: "fail", reason: "provider_changed"}
	}
	if baseline.Auditor != current.Auditor {
		return sustainResult{status: "fail", reason: "auditor_changed"}
	}
	if baseline.AttestedTier != current.AttestedTier {
		return sustainResult{status: "fail", reason: "attested_tier_changed"}
	}
	if !stringSlicesEqual(baseline.AttestedCapabilities, current.AttestedCapabilities) {
		return sustainResult{status: "fail", reason: "attested_capabilities_changed"}
	}
	if baseline.Software.BinaryHash != current.Software.BinaryHash {
		return sustainResult{status: "fail", reason: "software_identity_changed"}
	}
	if evidenceHasFailedCheck(current) {
		return sustainResult{status: "fail", reason: "current_check_failed"}
	}

	return sustainResult{status: "pass"}
}

func evidenceHasFailedCheck(evidence EvidenceDocument) bool {
	for _, check := range evidence.Checks {
		if check.Status == "fail" {
			return true
		}
	}

	return false
}
