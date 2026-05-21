package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/jsonpb"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/verification/v1"
)

const (
	flagAuditEscrowID     = "audit-escrow-id"
	flagAuditor           = "auditor"
	flagAuthority         = "authority"
	flagCapabilities      = "capabilities"
	flagDeposit           = "deposit"
	flagEvidenceHash      = "evidence-hash"
	flagExpiresAt         = "expires-at"
	flagFee               = "fee"
	flagFaultAttribution  = "fault-attribution"
	flagMaxTier           = "max-tier"
	flagMetadataHash      = "metadata-hash"
	flagProvider          = "provider"
	flagProviderDeposit   = "provider-deposit"
	flagReason            = "reason"
	flagRequestedTier     = "requested-tier"
	flagResourceSummary   = "resource-summary"
	flagSnapshotHash      = "snapshot-hash"
	flagSnapshotTimestamp = "snapshot-timestamp"
	flagTier              = "tier"
)

func GetTxVerificationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        verificationModuleName,
		Short:                      "Verification transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetTxVerificationRegisterAuditorCmd(),
		GetTxVerificationPostAuditorBondCmd(),
		GetTxVerificationPostProviderBondCmd(),
		GetTxVerificationPostSnapshotHashCmd(),
		GetTxVerificationOpenAuditEscrowCmd(),
		GetTxVerificationCancelAuditEscrowCmd(),
		GetTxVerificationSettleAuditEscrowCmd(),
		GetTxVerificationSubmitAttestationCmd(),
		GetTxVerificationRevokeAttestationCmd(),
		GetTxVerificationRemoveAttestationCmd(),
	)

	return cmd
}

func GetTxVerificationRegisterAuditorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "register-auditor [auditor]",
		Short:             "Register an auditor",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			auditor, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			authority, err := readAuthority(cmd, cctx)
			if err != nil {
				return err
			}
			tier, err := readTierFlag(cmd, flagMaxTier)
			if err != nil {
				return err
			}
			metadataHash, err := readOptionalHashFlag(cmd, flagMetadataHash)
			if err != nil {
				return err
			}

			msg := &types.MsgRegisterAuditor{
				Authority:          authority,
				Auditor:            auditor.String(),
				MaxAttestationTier: tier,
				MetadataHash:       metadataHash,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagAuthority, "", "Governance authority address; defaults to --from")
	cmd.Flags().String(flagMaxTier, "identified", "Maximum attestation tier")
	cmd.Flags().String(flagMetadataHash, "", "Optional metadata hash in hex or sha256:<hex> form")

	return cmd
}

func GetTxVerificationPostAuditorBondCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "post-auditor-bond [amount]",
		Short:             "Post auditor bond",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgPostAuditorBond{
				Auditor: cctx.GetFromAddress().String(),
				Amount:  amount,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetTxVerificationPostProviderBondCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "post-provider-bond [amount]",
		Short:             "Post provider bond",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgPostProviderBond{
				Provider: cctx.GetFromAddress().String(),
				Amount:   amount,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetTxVerificationPostSnapshotHashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "post-snapshot-hash",
		Short:             "Post provider snapshot hash",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			snapshotHash, err := readRequiredHashFlag(cmd, flagSnapshotHash)
			if err != nil {
				return err
			}
			summary, err := readResourceSummaryFlag(cmd)
			if err != nil {
				return err
			}
			snapshotTimestamp, err := readTimeFlag(cmd, flagSnapshotTimestamp)
			if err != nil {
				return err
			}

			msg := &types.MsgPostSnapshotHash{
				Provider:          cctx.GetFromAddress().String(),
				SnapshotHash:      snapshotHash,
				ResourceSummary:   summary,
				SnapshotTimestamp: snapshotTimestamp,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagSnapshotHash, "", "Snapshot hash in hex or sha256:<hex> form")
	cmd.Flags().String(flagResourceSummary, "", "Resource summary JSON or path to JSON file")
	cmd.Flags().String(flagSnapshotTimestamp, "", "Snapshot timestamp in RFC3339 format")
	mustMarkRequired(cmd, flagSnapshotHash, flagResourceSummary, flagSnapshotTimestamp)

	return cmd
}

func GetTxVerificationOpenAuditEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "open-audit-escrow",
		Short:             "Open provider audit escrow",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			tier, err := readTierFlag(cmd, flagRequestedTier)
			if err != nil {
				return err
			}
			capabilities, err := readCapabilitiesFlag(cmd, flagCapabilities)
			if err != nil {
				return err
			}
			fee, err := readCoinFlag(cmd, flagFee)
			if err != nil {
				return err
			}
			providerDeposit, err := readCoinFlag(cmd, flagProviderDeposit)
			if err != nil {
				return err
			}
			expiresAt, err := readTimeFlag(cmd, flagExpiresAt)
			if err != nil {
				return err
			}
			metadataHash, err := readOptionalHashFlag(cmd, flagMetadataHash)
			if err != nil {
				return err
			}

			msg := &types.MsgOpenAuditEscrow{
				Provider:              cctx.GetFromAddress().String(),
				RequestedTier:         tier,
				RequestedCapabilities: capabilities,
				Fee:                   fee,
				ProviderDeposit:       providerDeposit,
				ExpiresAt:             expiresAt,
				MetadataHash:          metadataHash,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagRequestedTier, "identified", "Requested verification tier")
	cmd.Flags().String(flagCapabilities, "", "Comma-separated requested capabilities")
	cmd.Flags().String(flagFee, "", "Audit fee coin")
	cmd.Flags().String(flagProviderDeposit, "", "Provider deposit coin")
	cmd.Flags().String(flagExpiresAt, "", "Escrow expiry in RFC3339 format")
	cmd.Flags().String(flagMetadataHash, "", "Optional metadata hash in hex or sha256:<hex> form")
	mustMarkRequired(cmd, flagFee, flagProviderDeposit, flagExpiresAt)

	return cmd
}

func GetTxVerificationCancelAuditEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "cancel-audit-escrow [audit-escrow-id]",
		Short:             "Cancel an open provider audit escrow",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			escrowID, err := parseUint64Arg(args[0], "audit escrow id")
			if err != nil {
				return err
			}

			msg := &types.MsgCancelAuditEscrow{
				Provider:      cctx.GetFromAddress().String(),
				AuditEscrowID: escrowID,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetTxVerificationSettleAuditEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "settle-audit-escrow [audit-escrow-id]",
		Short:             "Settle an unconsumed audit escrow",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			escrowID, err := parseUint64Arg(args[0], "audit escrow id")
			if err != nil {
				return err
			}
			authority, err := readAuthority(cmd, cctx)
			if err != nil {
				return err
			}
			reason, err := readAuditEscrowSettlementReasonFlag(cmd)
			if err != nil {
				return err
			}
			fault, err := readFaultAttributionFlag(cmd)
			if err != nil {
				return err
			}
			evidenceHash, err := readRequiredHashFlag(cmd, flagEvidenceHash)
			if err != nil {
				return err
			}

			msg := &types.MsgSettleAuditEscrow{
				Authority:        authority,
				AuditEscrowID:    escrowID,
				Reason:           reason,
				FaultAttribution: fault,
				EvidenceHash:     evidenceHash,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagAuthority, "", "Governance authority address; defaults to --from")
	cmd.Flags().String(flagReason, "", "Settlement reason: provider_fault or no_fault")
	cmd.Flags().String(flagFaultAttribution, "", "Fault attribution: provider_fault or no_fault")
	cmd.Flags().String(flagEvidenceHash, "", "Evidence hash in hex or sha256:<hex> form")
	mustMarkRequired(cmd, flagReason, flagFaultAttribution, flagEvidenceHash)

	return cmd
}

func GetTxVerificationSubmitAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "submit-attestation",
		Short:             "Submit provider attestation",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			provider, err := readAddressFlag(cmd, flagProvider)
			if err != nil {
				return err
			}
			escrowID, err := readUint64Flag(cmd, flagAuditEscrowID)
			if err != nil {
				return err
			}
			tier, err := readTierFlag(cmd, flagTier)
			if err != nil {
				return err
			}
			capabilities, err := readCapabilitiesFlag(cmd, flagCapabilities)
			if err != nil {
				return err
			}
			evidenceHash, err := readRequiredHashFlag(cmd, flagEvidenceHash)
			if err != nil {
				return err
			}
			fee, err := readCoinFlag(cmd, flagFee)
			if err != nil {
				return err
			}
			deposit, err := readCoinFlag(cmd, flagDeposit)
			if err != nil {
				return err
			}

			msg := &types.MsgSubmitAttestation{
				Provider:      provider.String(),
				Auditor:       cctx.GetFromAddress().String(),
				Tier:          tier,
				Capabilities:  capabilities,
				EvidenceHash:  evidenceHash,
				Fee:           fee,
				Deposit:       deposit,
				AuditEscrowID: escrowID,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagProvider, "", "Provider address")
	cmd.Flags().Uint64(flagAuditEscrowID, 0, "Audit escrow id")
	cmd.Flags().String(flagTier, "identified", "Attested verification tier")
	cmd.Flags().String(flagCapabilities, "", "Comma-separated attested capabilities")
	cmd.Flags().String(flagEvidenceHash, "", "Evidence hash in hex or sha256:<hex> form")
	cmd.Flags().String(flagFee, "", "Audit fee coin")
	cmd.Flags().String(flagDeposit, "", "Auditor deposit coin")
	mustMarkRequired(cmd, flagProvider, flagAuditEscrowID, flagEvidenceHash, flagFee, flagDeposit)

	return cmd
}

func GetTxVerificationRevokeAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "revoke-attestation",
		Short:             "Revoke provider attestation",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			provider, err := readAddressFlag(cmd, flagProvider)
			if err != nil {
				return err
			}
			reason, err := readAttestationRevocationReasonFlag(cmd)
			if err != nil {
				return err
			}
			evidenceHash, err := readRequiredHashFlag(cmd, flagEvidenceHash)
			if err != nil {
				return err
			}

			msg := &types.MsgRevokeAttestation{
				Provider:     provider.String(),
				Auditor:      cctx.GetFromAddress().String(),
				Reason:       reason,
				EvidenceHash: evidenceHash,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagProvider, "", "Provider address")
	cmd.Flags().String(flagReason, "", "Revocation reason")
	cmd.Flags().String(flagEvidenceHash, "", "Evidence hash in hex or sha256:<hex> form")
	mustMarkRequired(cmd, flagProvider, flagReason, flagEvidenceHash)

	return cmd
}

func GetTxVerificationRemoveAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "remove-attestation",
		Short:             "Remove provider attestation",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			auditor, err := readAddressFlag(cmd, flagAuditor)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveAttestation{
				Provider: cctx.GetFromAddress().String(),
				Auditor:  auditor.String(),
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagAuditor, "", "Auditor address for the attestation to remove")
	mustMarkRequired(cmd, flagAuditor)

	return cmd
}

func readAuthority(cmd *cobra.Command, cctx sdkclient.Context) (string, error) {
	authority, err := cmd.Flags().GetString(flagAuthority)
	if err != nil {
		return "", err
	}
	if authority == "" {
		return cctx.GetFromAddress().String(), nil
	}
	addr, err := sdk.AccAddressFromBech32(authority)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

func readAddressFlag(cmd *cobra.Command, name string) (sdk.AccAddress, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, err
	}
	return sdk.AccAddressFromBech32(val)
}

func readCoinFlag(cmd *cobra.Command, name string) (sdk.Coin, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return sdk.Coin{}, err
	}
	return sdk.ParseCoinNormalized(val)
}

func readUint64Flag(cmd *cobra.Command, name string) (uint64, error) {
	val, err := cmd.Flags().GetUint64(name)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func parseUint64Arg(val, name string) (uint64, error) {
	res, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: %w", name, val, err)
	}
	return res, nil
}

func readTierFlag(cmd *cobra.Command, name string) (types.VerificationTier, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return types.TierUnspecified, err
	}
	return parseVerificationTier(val)
}

func parseVerificationTier(val string) (types.VerificationTier, error) {
	switch normalizeEnumInput(val) {
	case "l1", "identified", "verification_tier_identified":
		return types.TierIdentified, nil
	case "l2", "verified", "verification_tier_verified":
		return types.TierVerified, nil
	case "l3", "established", "verification_tier_established":
		return types.TierEstablished, nil
	case "l4", "trusted", "verification_tier_trusted":
		return types.TierTrusted, nil
	default:
		return types.TierUnspecified, fmt.Errorf("invalid verification tier %q", val)
	}
}

func readAuditEscrowSettlementReasonFlag(cmd *cobra.Command) (types.AuditEscrowSettlementReason, error) {
	val, err := cmd.Flags().GetString(flagReason)
	if err != nil {
		return types.AuditEscrowSettlementReasonUnspecified, err
	}
	return parseAuditEscrowSettlementReason(val)
}

func readAttestationRevocationReasonFlag(cmd *cobra.Command) (types.AttestationRevocationReason, error) {
	val, err := cmd.Flags().GetString(flagReason)
	if err != nil {
		return types.AttestationRevocationReasonUnspecified, err
	}
	return parseAttestationRevocationReason(val)
}

func parseAttestationRevocationReason(val string) (types.AttestationRevocationReason, error) {
	switch normalizeEnumInput(val) {
	case "provider_no_longer_qualifies", "attestation_revocation_reason_provider_no_longer_qualifies":
		return types.AttestationRevocationReasonProviderNoLongerQualifies, nil
	case "snapshot_mismatch", "attestation_revocation_reason_snapshot_mismatch":
		return types.AttestationRevocationReasonSnapshotMismatch, nil
	case "software_identity_changed", "attestation_revocation_reason_software_identity_changed":
		return types.AttestationRevocationReasonSoftwareIdentityChanged, nil
	case "capability_misrepresented", "attestation_revocation_reason_capability_misrepresented":
		return types.AttestationRevocationReasonCapabilityMisrepresented, nil
	case "provider_non_responsive", "attestation_revocation_reason_provider_non_responsive":
		return types.AttestationRevocationReasonProviderNonResponsive, nil
	case "auditor_evidence_error", "attestation_revocation_reason_auditor_evidence_error":
		return types.AttestationRevocationReasonAuditorEvidenceError, nil
	case "auditor_operational_exit", "attestation_revocation_reason_auditor_operational_exit":
		return types.AttestationRevocationReasonAuditorOperationalExit, nil
	default:
		return types.AttestationRevocationReasonUnspecified, fmt.Errorf("invalid attestation revocation reason %q", val)
	}
}

func parseAuditEscrowSettlementReason(val string) (types.AuditEscrowSettlementReason, error) {
	switch normalizeEnumInput(val) {
	case "provider_fault", "audit_escrow_settlement_reason_provider_fault":
		return types.AuditEscrowSettlementReasonProviderFault, nil
	case "no_fault", "audit_escrow_settlement_reason_no_fault":
		return types.AuditEscrowSettlementReasonNoFault, nil
	default:
		return types.AuditEscrowSettlementReasonUnspecified, fmt.Errorf("invalid audit escrow settlement reason %q", val)
	}
}

func readFaultAttributionFlag(cmd *cobra.Command) (types.FaultAttribution, error) {
	val, err := cmd.Flags().GetString(flagFaultAttribution)
	if err != nil {
		return types.FaultAttributionUnspecified, err
	}
	return parseFaultAttribution(val)
}

func parseFaultAttribution(val string) (types.FaultAttribution, error) {
	switch normalizeEnumInput(val) {
	case "provider_fault", "fault_attribution_provider_fault":
		return types.FaultAttributionProviderFault, nil
	case "auditor_fault", "fault_attribution_auditor_fault":
		return types.FaultAttributionAuditorFault, nil
	case "shared_fault", "fault_attribution_shared_fault":
		return types.FaultAttributionSharedFault, nil
	case "no_fault", "fault_attribution_no_fault":
		return types.FaultAttributionNoFault, nil
	case "inconclusive", "fault_attribution_inconclusive":
		return types.FaultAttributionInconclusive, nil
	default:
		return types.FaultAttributionUnspecified, fmt.Errorf("invalid fault attribution %q", val)
	}
}

func readCapabilitiesFlag(cmd *cobra.Command, name string) ([]types.CapabilityFlag, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(val) == "" {
		return nil, nil
	}

	parts := strings.Split(val, ",")
	res := make([]types.CapabilityFlag, 0, len(parts))
	for _, part := range parts {
		capability, err := parseCapabilityFlag(part)
		if err != nil {
			return nil, err
		}
		res = append(res, capability)
	}

	return res, nil
}

func parseCapabilityFlag(val string) (types.CapabilityFlag, error) {
	switch normalizeEnumInput(val) {
	case "tee", "tee_hardware_attestation", "capability_tee_hardware_attestation":
		return types.CapabilityTEEHardwareAttestation, nil
	case "confidential_computing", "capability_confidential_computing":
		return types.CapabilityConfidentialComputing, nil
	case "persistent_storage", "capability_persistent_storage":
		return types.CapabilityPersistentStorage, nil
	case "bare_metal", "capability_bare_metal":
		return types.CapabilityBareMetal, nil
	default:
		return types.CapabilityUnspecified, fmt.Errorf("invalid capability %q", val)
	}
}

func readRequiredHashFlag(cmd *cobra.Command, name string) ([]byte, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, err
	}
	return parseHexHash(val)
}

func readOptionalHashFlag(cmd *cobra.Command, name string) ([]byte, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, nil
	}
	return parseHexHash(val)
}

func parseHexHash(val string) ([]byte, error) {
	val = strings.TrimSpace(val)
	val = strings.TrimPrefix(val, "sha256:")
	val = strings.TrimPrefix(val, "0x")
	if val == "" {
		return nil, fmt.Errorf("hash is required")
	}
	res, err := hex.DecodeString(val)
	if err != nil {
		return nil, err
	}
	if len(res) != sha256.Size {
		return nil, fmt.Errorf("hash must be %d bytes", sha256.Size)
	}
	return res, nil
}

func readResourceSummaryFlag(cmd *cobra.Command) (types.ResourceSummary, error) {
	val, err := cmd.Flags().GetString(flagResourceSummary)
	if err != nil {
		return types.ResourceSummary{}, err
	}
	if strings.TrimSpace(val) == "" {
		return types.ResourceSummary{}, fmt.Errorf("resource summary is required")
	}

	raw := []byte(val)
	if !strings.HasPrefix(strings.TrimSpace(val), "{") {
		raw, err = os.ReadFile(val) //nolint:gosec
		if err != nil {
			return types.ResourceSummary{}, err
		}
	}

	var summary types.ResourceSummary
	if err := jsonpb.Unmarshal(strings.NewReader(string(raw)), &summary); err != nil {
		return types.ResourceSummary{}, err
	}

	return summary, nil
}

func readTimeFlag(cmd *cobra.Command, name string) (time.Time, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, val)
}

func mustMarkRequired(cmd *cobra.Command, names ...string) {
	for _, name := range names {
		if err := cmd.MarkFlagRequired(name); err != nil {
			panic(err.Error())
		}
	}
}

func normalizeEnumInput(val string) string {
	return strings.ToLower(strings.TrimSpace(strings.ReplaceAll(val, "-", "_")))
}
