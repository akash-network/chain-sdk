package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

const sustainedArtifactFile = "sustained.json"

type sustainConfig struct {
	baselineDir string
	currentDir  string
	outputDir   string
}

type sustainedArtifact struct {
	SchemaVersion         string            `json:"schema_version"`
	CheckedAt             string            `json:"checked_at"`
	BaselineEvidenceHash  string            `json:"baseline_evidence_hash"`
	CurrentEvidenceHash   string            `json:"current_evidence_hash"`
	SustainedEvidenceHash string            `json:"sustained_evidence_hash"`
	Status                string            `json:"status"`
	Reason                string            `json:"reason,omitempty"`
	RevocationReason      string            `json:"revocation_reason,omitempty"`
	Files                 map[string]string `json:"files"`
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

	cmd.Flags().StringVar(&cfg.outputDir, "output-dir", "", "Write sustained-validation evidence artifacts to this directory")

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
	if cfg.outputDir != "" {
		hash, err := writeSustainedEvidenceArtifacts(cfg.outputDir, cfg.currentDir, current, baselineHash, currentHash, result)
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "sustained_evidence_dir=%s\n", cfg.outputDir)
		fmt.Fprintf(cmd.OutOrStdout(), "sustained_evidence_hash=%s\n", hash)
	}

	return nil
}

type sustainResult struct {
	status      string
	reason      string
	failedCheck string
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
	if failedCheck := firstFailedEvidenceCheck(current); failedCheck != "" {
		return sustainResult{status: "fail", reason: "current_check_failed", failedCheck: failedCheck}
	}

	return sustainResult{status: "pass"}
}

func evidenceHasFailedCheck(evidence EvidenceDocument) bool {
	return firstFailedEvidenceCheck(evidence) != ""
}

func firstFailedEvidenceCheck(evidence EvidenceDocument) string {
	for _, check := range evidence.Checks {
		if check.Status == "fail" {
			return check.Name
		}
	}

	return ""
}

func writeSustainedEvidenceArtifacts(outputDir, currentDir string, current EvidenceDocument, baselineHash, currentHash string, result sustainResult) (string, error) {
	if outputDir == "" {
		return "", fmt.Errorf("output-dir is required")
	}

	evidence := current
	lastCheckedAt := evidence.CollectedAt
	if lastCheckedAt == "" {
		lastCheckedAt = time.Now().UTC().Format(time.RFC3339Nano)
	}
	evidence.SustainedValidation = SustainedValidation{
		BaselineID:    baselineHash,
		Window:        "attestation_ttl",
		LastCheckedAt: lastCheckedAt,
		Status:        result.status,
		ProofRef:      currentHash,
	}
	evidence.Checks = upsertSustainedBaselineCheck(evidence.Checks, result, baselineHash, currentHash, lastCheckedAt)
	if result.status == "fail" {
		evidence.FaultContext = FaultContext{
			FaultAttribution: "provider_fault",
			Reason:           revocationReasonFromSustain(result),
		}
	} else {
		evidence.FaultContext = FaultContext{
			FaultAttribution: "unspecified",
			Reason:           "unspecified",
		}
	}

	raw, hash, err := marshalEvidenceCanonical(evidence)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return "", err
	}
	if err := copyCollectedArtifacts(currentDir, outputDir); err != nil {
		return "", err
	}

	files := map[string]string{
		"evidence_draft":      filepath.Join(outputDir, evidenceDraftFile),
		"evidence_draft_hash": filepath.Join(outputDir, evidenceHashFile),
		"sustained":           filepath.Join(outputDir, sustainedArtifactFile),
	}
	if err := os.WriteFile(files["evidence_draft"], raw, 0o644); err != nil {
		return "", err
	}
	if err := os.WriteFile(files["evidence_draft_hash"], []byte(hash+"\n"), 0o644); err != nil {
		return "", err
	}

	artifact := sustainedArtifact{
		SchemaVersion:         evidenceSchema + ".sustained",
		CheckedAt:             time.Now().UTC().Format(time.RFC3339Nano),
		BaselineEvidenceHash:  baselineHash,
		CurrentEvidenceHash:   currentHash,
		SustainedEvidenceHash: hash,
		Status:                result.status,
		Reason:                result.reason,
		RevocationReason:      evidence.FaultContext.Reason,
		Files:                 files,
	}
	if artifact.Status != "fail" {
		artifact.RevocationReason = ""
	}
	if err := writeJSON(files["sustained"], artifact); err != nil {
		return "", err
	}

	return hash, nil
}

func copyCollectedArtifacts(srcDir, dstDir string) error {
	for _, name := range []string{collectionFile, nonceFile, snapshotPayloadFile, snapshotJSONFile, snapshotSigFile} {
		if err := copyArtifactFile(filepath.Join(srcDir, name), filepath.Join(dstDir, name)); err != nil {
			return err
		}
	}

	return nil
}

func copyArtifactFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("read %s: %w", filepath.Base(src), err)
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("write %s: %w", filepath.Base(dst), err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy %s: %w", filepath.Base(src), err)
	}

	return nil
}

func upsertSustainedBaselineCheck(checks []EvidenceCheck, result sustainResult, baselineHash, currentHash, observedAt string) []EvidenceCheck {
	status := "pass"
	if result.status == "fail" {
		status = "fail"
	}
	check := EvidenceCheck{
		Name:       "sustained_baseline_checked",
		Status:     status,
		ProofRef:   currentHash,
		ObservedAt: observedAt,
		Details: map[string]any{
			"baseline_evidence_hash": baselineHash,
			"current_evidence_hash":  currentHash,
		},
	}
	if result.reason != "" {
		check.Details["reason"] = result.reason
	}
	if result.failedCheck != "" {
		check.Details["failed_check"] = result.failedCheck
	}

	res := append([]EvidenceCheck(nil), checks...)
	for idx := range res {
		if res[idx].Name == check.Name {
			res[idx] = check
			return res
		}
	}

	return append(res, check)
}

func revocationReasonFromSustain(result sustainResult) string {
	switch result.reason {
	case "software_identity_changed":
		return "software_identity_changed"
	case "attested_capabilities_changed":
		return "capability_misrepresented"
	case "current_check_failed":
		switch result.failedCheck {
		case "snapshot_not_suspended", "snapshot_hash_matches_chain", "inventory_nonce_matches", "inventory_signature_valid":
			return "snapshot_mismatch"
		case "software_identity_recorded":
			return "software_identity_changed"
		case "capability_verified":
			return "capability_misrepresented"
		case "contact_responsiveness_checked":
			return "provider_non_responsive"
		default:
			return "provider_no_longer_qualifies"
		}
	default:
		return "provider_no_longer_qualifies"
	}
}
