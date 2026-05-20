package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

type submitConfig struct {
	artifactDir   string
	provider      string
	auditor       string
	auditEscrowID string
	tier          string
	capabilities  []string
	fee           string
	deposit       string
	chainID       string
}

func newSubmitCmd() *cobra.Command {
	cfg := submitConfig{}

	cmd := &cobra.Command{
		Use:   "submit <artifact-dir>",
		Short: "Validate evidence and print submit-attestation command",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.artifactDir = args[0]
			return runSubmit(cmd, cfg)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&cfg.provider, "provider", "", "Provider address; defaults to evidence provider")
	flags.StringVar(&cfg.auditor, "auditor", "", "Auditor address for --from; defaults to evidence auditor")
	flags.StringVar(&cfg.auditEscrowID, "audit-escrow-id", "", "Audit escrow id; defaults to evidence audit_escrow_id")
	flags.StringVar(&cfg.tier, "tier", "", "Attested tier; defaults to evidence attested_tier")
	flags.StringSliceVar(&cfg.capabilities, "capability", nil, "Attested capability; repeat or comma-separate; defaults to evidence attested_capabilities")
	flags.StringVar(&cfg.fee, "fee", "", "Audit fee coin")
	flags.StringVar(&cfg.deposit, "deposit", "", "Auditor deposit coin")
	flags.StringVar(&cfg.chainID, "chain-id", "", "Chain id; defaults to evidence chain_id")

	_ = cmd.MarkFlagRequired("fee")
	_ = cmd.MarkFlagRequired("deposit")

	return cmd
}

func runSubmit(cmd *cobra.Command, cfg submitConfig) error {
	evidence, evidenceHash, err := loadVerifiedEvidenceArtifactDir(cfg.artifactDir)
	if err != nil {
		return err
	}

	spec, err := buildSubmitCommandSpec(cfg, evidence, evidenceHash)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), "submit=dry-run")
	fmt.Fprintln(cmd.OutOrStdout(), "broadcast=not_implemented")
	fmt.Fprintf(cmd.OutOrStdout(), "evidence_hash=%s\n", evidenceHash)
	fmt.Fprintf(cmd.OutOrStdout(), "command=%s\n", shellCommand(spec.args))

	return nil
}

type submitCommandSpec struct {
	args []string
}

func buildSubmitCommandSpec(cfg submitConfig, evidence EvidenceDocument, evidenceHash string) (submitCommandSpec, error) {
	provider := evidence.Provider
	if cfg.provider != "" {
		if cfg.provider != evidence.Provider {
			return submitCommandSpec{}, fmt.Errorf("provider flag %q does not match evidence provider %q", cfg.provider, evidence.Provider)
		}
		provider = cfg.provider
	}

	auditor := evidence.Auditor
	if cfg.auditor != "" {
		if cfg.auditor != evidence.Auditor {
			return submitCommandSpec{}, fmt.Errorf("auditor flag %q does not match evidence auditor %q", cfg.auditor, evidence.Auditor)
		}
		auditor = cfg.auditor
	}

	auditEscrowID := evidence.AuditEscrowID
	if cfg.auditEscrowID != "" {
		if !sameUint64String(cfg.auditEscrowID, evidence.AuditEscrowID) {
			return submitCommandSpec{}, fmt.Errorf("audit-escrow-id flag %q does not match evidence audit_escrow_id %q", cfg.auditEscrowID, evidence.AuditEscrowID)
		}
		auditEscrowID = evidence.AuditEscrowID
	}

	tier := evidence.AttestedTier
	if cfg.tier != "" {
		normalizedTier, err := normalizeSubmitTier(cfg.tier)
		if err != nil {
			return submitCommandSpec{}, err
		}
		if normalizedTier != evidence.AttestedTier {
			return submitCommandSpec{}, fmt.Errorf("tier flag %q does not match evidence attested_tier %q", cfg.tier, evidence.AttestedTier)
		}
		tier = normalizedTier
	}

	capabilities := append([]string(nil), evidence.AttestedCapabilities...)
	if cfg.capabilities != nil {
		normalized, err := normalizeSubmitCapabilities(cfg.capabilities)
		if err != nil {
			return submitCommandSpec{}, err
		}
		if !stringSlicesEqual(normalized, evidence.AttestedCapabilities) {
			return submitCommandSpec{}, fmt.Errorf("capability flags %q do not match evidence attested_capabilities %q", strings.Join(normalized, ","), strings.Join(evidence.AttestedCapabilities, ","))
		}
		capabilities = normalized
	}

	chainID := evidence.ChainID
	if cfg.chainID != "" {
		if cfg.chainID != evidence.ChainID {
			return submitCommandSpec{}, fmt.Errorf("chain-id flag %q does not match evidence chain_id %q", cfg.chainID, evidence.ChainID)
		}
		chainID = cfg.chainID
	}

	fee, err := parseRequiredCoinFlag("fee", cfg.fee)
	if err != nil {
		return submitCommandSpec{}, err
	}
	deposit, err := parseRequiredCoinFlag("deposit", cfg.deposit)
	if err != nil {
		return submitCommandSpec{}, err
	}

	args := []string{
		"akash", "tx", "verification", "submit-attestation",
		"--provider", provider,
		"--audit-escrow-id", auditEscrowID,
		"--tier", tier,
	}
	if len(capabilities) > 0 {
		args = append(args, "--capabilities", strings.Join(capabilities, ","))
	}
	args = append(args,
		"--evidence-hash", evidenceHash,
		"--fee", fee.String(),
		"--deposit", deposit.String(),
		"--from", auditor,
		"--chain-id", chainID,
	)

	return submitCommandSpec{args: args}, nil
}

func loadVerifiedEvidenceArtifactDir(artifactDir string) (EvidenceDocument, string, error) {
	if artifactDir == "" {
		return EvidenceDocument{}, "", fmt.Errorf("artifact-dir is required")
	}

	raw, err := os.ReadFile(filepath.Join(artifactDir, evidenceDraftFile))
	if err != nil {
		return EvidenceDocument{}, "", fmt.Errorf("read %s: %w", evidenceDraftFile, err)
	}

	hash, err := canonicalEvidenceHash(raw)
	if err != nil {
		return EvidenceDocument{}, "", err
	}
	if err := validateEvidenceHashArtifact(artifactDir, hash); err != nil {
		return EvidenceDocument{}, "", err
	}

	var evidence EvidenceDocument
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	if err := decoder.Decode(&evidence); err != nil {
		return EvidenceDocument{}, "", fmt.Errorf("decode evidence: %w", err)
	}

	return evidence, hash, nil
}

func normalizeSubmitTier(val string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(strings.ReplaceAll(val, "-", "_"))) {
	case "l1", "identified", "verification_tier_identified":
		return "L1", nil
	case "l2", "verified", "verification_tier_verified":
		return "L2", nil
	case "l3", "established", "verification_tier_established":
		return "L3", nil
	case "l4", "trusted", "verification_tier_trusted":
		return "L4", nil
	default:
		return "", fmt.Errorf("invalid tier %q", val)
	}
}

func normalizeSubmitCapabilities(capabilities []string) ([]string, error) {
	res := append([]string(nil), capabilities...)
	sort.Strings(res)
	if err := validateCapabilitySet("capability", res); err != nil {
		return nil, err
	}

	return res, nil
}

func parseRequiredCoinFlag(name, val string) (sdk.Coin, error) {
	if strings.TrimSpace(val) == "" {
		return sdk.Coin{}, fmt.Errorf("%s is required", name)
	}
	coin, err := sdk.ParseCoinNormalized(val)
	if err != nil {
		return sdk.Coin{}, fmt.Errorf("invalid %s: %w", name, err)
	}

	return coin, nil
}

func sameUint64String(left, right string) bool {
	leftVal, leftErr := strconv.ParseUint(left, 10, 64)
	rightVal, rightErr := strconv.ParseUint(right, 10, 64)

	return leftErr == nil && rightErr == nil && leftVal == rightVal
}

func stringSlicesEqual(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for idx := range left {
		if left[idx] != right[idx] {
			return false
		}
	}

	return true
}

var safeShellArgPattern = regexp.MustCompile(`^[A-Za-z0-9_./:=@%+,\-]+$`)

func shellCommand(args []string) string {
	quoted := make([]string, 0, len(args))
	for _, arg := range args {
		quoted = append(quoted, shellQuoteArg(arg))
	}

	return strings.Join(quoted, " ")
}

func shellQuoteArg(arg string) string {
	if arg != "" && safeShellArgPattern.MatchString(arg) {
		return arg
	}

	return "'" + strings.ReplaceAll(arg, "'", `'\''`) + "'"
}
