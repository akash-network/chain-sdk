package v1

import (
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	"cosmossdk.io/math"
	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/gagliardetto/solana-go"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const GovernanceRefundScope uint64 = 0

// Validate checks evidence syntax only. Observers establish Solana truth and finality.
func (e SolanaEvidence) Validate() error {
	if e.Slot == 0 || len(e.TransactionSignature) != solana.SignatureLength || solana.SignatureFromBytes(e.TransactionSignature).IsZero() {
		return ErrInvalidObservation.Wrap("positive slot and a nonzero 64-byte transaction signature required")
	}

	return nil
}

func (o *Observation) Validate() error {
	if o.Outcome != OutcomeMinted && o.Outcome != OutcomeVoided {
		return ErrInvalidObservation.Wrap("unknown outcome")
	}

	return (SolanaEvidence{
		Slot:                 o.Slot,
		TransactionSignature: o.TransactionSignature,
	}).Validate()
}

func ValidateReason(reason string) error {
	if !utf8.ValidString(reason) || strings.TrimSpace(reason) == "" || len(reason) > 256 {
		return ErrInvalidReason.Wrap("reason must contain 1 to 256 UTF-8 bytes")
	}

	return nil
}

func (s Status) IsValid() bool {
	switch s {
	case StatusTransferPending, StatusMinted, StatusVoided, StatusRefundPending, StatusRefunded:
		return true
	default:
		return false
	}
}

func (s Status) Terminal() bool {
	return s == StatusMinted || s == StatusRefunded
}

// Tallies counts only votes under the embedded authorization policy.
func (s *SettlementData) Tallies() (minted, voided uint32) {
	for _, vote := range s.Votes {
		if vote.PolicyVersion != s.Policy.Version {
			continue
		}

		switch vote.Observation.Outcome {
		case OutcomeMinted:
			minted++
		case OutcomeVoided:
			voided++
		}
	}

	return
}

func (s *SettlementData) HasVoidAuthorization() bool {
	if s.GovernanceObservation != nil {
		return s.GovernanceObservation.Outcome == OutcomeVoided
	}

	if s.Policy.Threshold.IsNil() || !s.Policy.Threshold.IsPositive() || s.Policy.Threshold.GT(math.LegacyOneDec()) {
		return false
	}

	_, voided := s.Tallies()
	required := s.Policy.Threshold.MulInt64(int64(len(s.Policy.Members))).Ceil().TruncateInt64()

	return required > 0 && int64(voided) == required
}

func (s *SettlementData) Payable(status Status, current uint64) bool {
	return status == StatusRefundPending && s.Hold == HoldNone && s.HasVoidAuthorization() &&
		(s.GovernanceObservation != nil || s.Policy.Version == current)
}

// Validate checks key identity and retained evidence. The keeper authenticates governance vetoes.
func (s *SettlementData) Validate(id ID, status Status, policies map[uint64]ReporterPolicy) error {
	if err := id.ValidateBasic(); err != nil {
		return err
	}

	if s.Round == 0 || !status.IsValid() {
		return ErrInvalidSettlement.Wrap("positive round and valid financial status required")
	}

	if err := s.Policy.Validate(); err != nil {
		return err
	}

	policy, ok := policies[s.Policy.Version]
	if !ok || !policy.Equal(s.Policy) {
		return ErrInvalidSettlement.Wrap("missing or inconsistent authorization policy")
	}

	if s.Hold != HoldNone && s.Hold != HoldVeto && s.Hold != HoldConflict {
		return ErrInvalidSettlement.Wrap("unknown hold")
	}

	if (status == StatusRefundPending) != (s.ChallengeEnd > 0) || s.ChallengeEnd < 0 {
		return ErrInvalidSettlement.Wrap("deadline required exactly for refund-pending")
	}

	if s.ChallengeEnd > 0 {
		if _, err := gogotypes.TimestampProto(time.Unix(s.ChallengeEnd, 0)); err != nil {
			return ErrInvalidSettlement.Wrap("invalid challenge deadline")
		}
	}

	if receipt := s.VoidAuthorization; receipt != nil {
		if receipt.RequestedHeight < id.Height || receipt.RequestedRound == 0 || receipt.RequestedRound > s.Round {
			return ErrInvalidSettlement.Wrap("invalid permanent void authorization")
		}
	}

	if status == StatusVoided && s.VoidAuthorization == nil || status == StatusTransferPending && s.VoidAuthorization != nil {
		return ErrInvalidSettlement.Wrap("pending status must match permanent void authorization")
	}

	veto := s.VetoBy != "" || s.VetoReason != "" || s.VetoPolicyVersion != 0
	if veto {
		if s.Hold == HoldNone || status != StatusRefundPending {
			return ErrInvalidSettlement.Wrap("orphaned veto")
		}

		if _, err := sdk.AccAddressFromBech32(s.VetoBy); err != nil {
			return err
		}

		if err := ValidateReason(s.VetoReason); err != nil {
			return err
		}

		if s.VetoPolicyVersion != GovernanceRefundScope {
			p, exists := policies[s.VetoPolicyVersion]
			if !exists || !slices.Contains(p.Members, s.VetoBy) {
				return ErrInvalidSettlement.Wrap("veto outside its reporter policy")
			}
		}
	}

	if s.Hold == HoldVeto && !veto {
		return ErrInvalidSettlement.Wrap("veto hold requires a refund candidate, signer and reason")
	}

	var primary Outcome
	switch status {
	case StatusMinted:
		primary = OutcomeMinted
	case StatusRefundPending, StatusRefunded:
		primary = OutcomeVoided
	}

	if s.GovernanceObservation != nil {
		if primary == OutcomeUnspecified || s.Round == 1 || s.GovernanceObservation.Outcome != primary {
			return ErrInvalidSettlement.Wrap("orphaned or contradictory governance observation")
		}

		if err := s.GovernanceObservation.Validate(); err != nil {
			return err
		}

		if status == StatusRefundPending && s.VoidAuthorization == nil {
			return ErrInvalidSettlement.Wrap("governance refund authorization requires a void order")
		}
	}

	if primary == OutcomeUnspecified && len(s.Votes) > 0 {
		primary = s.Votes[0].Observation.Outcome
	}

	seen := make(map[string]bool, len(s.Votes))
	var certificate, witnesses int64
	for _, vote := range s.Votes {
		p, exists := policies[vote.PolicyVersion]
		if !exists || !slices.Contains(p.Members, vote.Reporter) {
			return ErrInvalidSettlement.Wrap("vote outside its reporter policy")
		}

		if err := vote.Observation.Validate(); err != nil {
			return err
		}

		if vote.Observation.Outcome == primary {
			if s.GovernanceObservation != nil || vote.PolicyVersion != s.Policy.Version || seen[vote.Reporter] {
				return ErrInvalidSettlement.Wrap("invalid or duplicate certificate vote")
			}

			seen[vote.Reporter] = true
			certificate++
		} else {
			witnesses++
		}
	}

	required := s.Policy.Threshold.MulInt64(int64(len(s.Policy.Members))).Ceil().TruncateInt64()
	if witnesses > 1 || certificate > required {
		return ErrInvalidSettlement.Wrap("certificate or opposite witness exceeds authorization bound")
	}

	conflict := witnesses == 1 && (certificate > 0 || s.GovernanceObservation != nil)
	if (s.Hold == HoldConflict) != conflict {
		return ErrInvalidSettlement.Wrap("conflict hold must have opposing evidence")
	}

	authorized := s.GovernanceObservation != nil || required > 0 && certificate == required
	if status == StatusMinted || status == StatusRefundPending || status == StatusRefunded {
		if !authorized {
			return ErrInvalidSettlement.Wrap("missing retained quorum authorization")
		}
	} else if authorized {
		return ErrInvalidSettlement.Wrap("quorum must have an authorized financial status")
	}

	return nil
}

func policyHistory(policies []ReporterPolicy, sequence uint64) (map[uint64]ReporterPolicy, error) {
	history := make(map[uint64]ReporterPolicy, len(policies))

	for _, policy := range policies {
		if err := policy.Validate(); err != nil {
			return nil, err
		}

		if _, exists := history[policy.Version]; exists || policy.Version > sequence {
			return nil, ErrInvalidGenesis.Wrap("duplicate or future reporter policy")
		}

		history[policy.Version] = policy
	}

	return history, nil
}
