package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type revokeConfig struct {
	artifactDir string
	provider    string
	auditor     string
	reason      string
	chainID     string
}

func newRevokeCmd() *cobra.Command {
	cfg := revokeConfig{}

	cmd := &cobra.Command{
		Use:   "revoke <artifact-dir>",
		Short: "Validate revocation evidence and prepare revoke-attestation tx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.artifactDir = args[0]
			return runRevoke(cmd, cfg)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.provider, "provider", "", "Provider address; defaults to evidence provider")
	flags.StringVar(&cfg.auditor, "auditor", "", "Auditor address for --from; defaults to evidence auditor")
	flags.StringVar(&cfg.reason, "reason", "", "Revocation reason")
	flags.StringVar(&cfg.chainID, "chain-id", "", "Chain id; defaults to evidence chain_id")

	_ = cmd.MarkFlagRequired("reason")

	return cmd
}

func runRevoke(cmd *cobra.Command, cfg revokeConfig) error {
	evidence, evidenceHash, err := loadVerifiedEvidenceArtifactDir(cfg.artifactDir)
	if err != nil {
		return err
	}

	spec, err := buildRevokeCommandSpec(cfg, evidence, evidenceHash)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), "revoke=dry-run")
	fmt.Fprintln(cmd.OutOrStdout(), "broadcast=not_implemented")
	fmt.Fprintf(cmd.OutOrStdout(), "evidence_hash=%s\n", evidenceHash)
	fmt.Fprintf(cmd.OutOrStdout(), "command=%s\n", shellCommand(spec.args))

	return nil
}

func buildRevokeCommandSpec(cfg revokeConfig, evidence EvidenceDocument, evidenceHash string) (txCommandSpec, error) {
	provider := evidence.Provider
	if cfg.provider != "" {
		if cfg.provider != evidence.Provider {
			return txCommandSpec{}, fmt.Errorf("provider flag %q does not match evidence provider %q", cfg.provider, evidence.Provider)
		}
		provider = cfg.provider
	}

	auditor := evidence.Auditor
	if cfg.auditor != "" {
		if cfg.auditor != evidence.Auditor {
			return txCommandSpec{}, fmt.Errorf("auditor flag %q does not match evidence auditor %q", cfg.auditor, evidence.Auditor)
		}
		auditor = cfg.auditor
	}

	reason, err := normalizeRevocationReason(cfg.reason)
	if err != nil {
		return txCommandSpec{}, err
	}
	if err := validateRevocationEvidence(reason, evidence); err != nil {
		return txCommandSpec{}, err
	}

	chainID := evidence.ChainID
	if cfg.chainID != "" {
		if cfg.chainID != evidence.ChainID {
			return txCommandSpec{}, fmt.Errorf("chain-id flag %q does not match evidence chain_id %q", cfg.chainID, evidence.ChainID)
		}
		chainID = cfg.chainID
	}

	args := []string{
		"akash", "tx", "verification", "revoke-attestation",
		"--provider", provider,
		"--reason", reason,
		"--evidence-hash", evidenceHash,
		"--from", auditor,
		"--chain-id", chainID,
	}

	return txCommandSpec{args: args}, nil
}

func validateRevocationEvidence(reason string, evidence EvidenceDocument) error {
	if evidence.FaultContext.Reason == "unspecified" {
		return fmt.Errorf("revocation evidence must set fault_context.reason")
	}
	if evidence.FaultContext.Reason != reason {
		return fmt.Errorf("reason flag %q does not match evidence fault_context.reason %q", reason, evidence.FaultContext.Reason)
	}
	if !validRevocationFaultAttribution(reason, evidence.FaultContext.FaultAttribution) {
		return fmt.Errorf("fault_context.fault_attribution %q is not valid for revocation reason %q", evidence.FaultContext.FaultAttribution, reason)
	}
	if providerMonitoringRevocationReason(reason) && evidence.SustainedValidation.Status != "fail" {
		return fmt.Errorf("provider-side revocation evidence must have sustained_validation.status \"fail\"")
	}

	return nil
}

func validRevocationFaultAttribution(reason, fault string) bool {
	switch reason {
	case "provider_no_longer_qualifies", "snapshot_mismatch", "software_identity_changed", "capability_misrepresented", "provider_non_responsive":
		return fault == "provider_fault" || fault == "no_fault"
	case "auditor_evidence_error":
		return fault == "auditor_fault"
	case "auditor_operational_exit":
		return fault == "no_fault"
	default:
		return false
	}
}

func providerMonitoringRevocationReason(reason string) bool {
	switch reason {
	case "provider_no_longer_qualifies", "snapshot_mismatch", "software_identity_changed", "capability_misrepresented", "provider_non_responsive":
		return true
	default:
		return false
	}
}

func normalizeRevocationReason(val string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(strings.ReplaceAll(val, "-", "_"))) {
	case "provider_no_longer_qualifies", "attestation_revocation_reason_provider_no_longer_qualifies":
		return "provider_no_longer_qualifies", nil
	case "snapshot_mismatch", "attestation_revocation_reason_snapshot_mismatch":
		return "snapshot_mismatch", nil
	case "software_identity_changed", "attestation_revocation_reason_software_identity_changed":
		return "software_identity_changed", nil
	case "capability_misrepresented", "attestation_revocation_reason_capability_misrepresented":
		return "capability_misrepresented", nil
	case "provider_non_responsive", "attestation_revocation_reason_provider_non_responsive":
		return "provider_non_responsive", nil
	case "auditor_evidence_error", "attestation_revocation_reason_auditor_evidence_error":
		return "auditor_evidence_error", nil
	case "auditor_operational_exit", "attestation_revocation_reason_auditor_operational_exit":
		return "auditor_operational_exit", nil
	default:
		return "", fmt.Errorf("invalid revocation reason %q", val)
	}
}
