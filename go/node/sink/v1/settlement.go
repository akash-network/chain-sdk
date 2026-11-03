package v1

import (
	"bytes"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	gogotypes "github.com/cosmos/gogoproto/types"
)

const GovernanceRefundScope uint64 = 0

// Validate checks evidence syntax only. Observers establish Solana truth and finality.
func (e SolanaEvidence) Validate() error {
	if e.Slot == 0 || len(e.TransactionSignature) != 64 || bytes.Equal(e.TransactionSignature, make([]byte, 64)) {
		return ErrInvalidObservation.Wrap("positive slot and a nonzero 64-byte transaction signature required")
	}

	return nil
}

func (o Observation) Validate() error {
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
func (s SettlementData) Tallies() (minted, voided uint32) {
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

func (s SettlementData) HasVoidAuthorization() bool {
	if s.GovernanceObservation != nil {
		return s.GovernanceObservation.Outcome == OutcomeVoided
	}

	_, voided := s.Tallies()
	return s.Policy.Threshold > 0 && voided == s.Policy.Threshold
}

func (s SettlementData) Payable(status Status, current uint64) bool {
	return status == StatusRefundPending && s.Hold == HoldNone && s.HasVoidAuthorization() &&
		(s.GovernanceObservation != nil || s.Policy.Version == current)
}

func (s SettlementData) RefundScope(status Status) (scope uint64, ok bool) {
	if status != StatusRefundPending || s.Hold != HoldNone {
		return 0, false
	}

	if s.GovernanceObservation != nil {
		return GovernanceRefundScope, true
	}

	return s.Policy.Version, true
}

// Validate checks key identity and retained evidence. The keeper authenticates governance vetoes.
func (s SettlementData) Validate(id ID, status Status, policies map[uint64]ReporterPolicy) error {
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
	if !ok || !policy.SameRule(s.Policy) {
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

		if _, err := ParseAccount(s.VetoBy, AccountLength); err != nil {
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

	if s.Hold == HoldVeto && (!veto || status != StatusRefundPending) {
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

	if len(s.Votes) > MaxReporters+1 {
		return ErrInvalidSettlement.Wrap("certificate and opposite witness exceed evidence bound")
	}

	if primary == OutcomeUnspecified && len(s.Votes) > 0 {
		primary = s.Votes[0].Observation.Outcome
	}

	seen := make(map[string]bool, len(s.Votes))
	var certificate, witnesses uint32
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

	if witnesses > 1 || certificate > s.Policy.Threshold {
		return ErrInvalidSettlement.Wrap("certificate or opposite witness exceeds authorization bound")
	}

	conflict := witnesses == 1 && (certificate > 0 || s.GovernanceObservation != nil)
	if (s.Hold == HoldConflict) != conflict {
		return ErrInvalidSettlement.Wrap("conflict hold must have opposing evidence")
	}

	authorized := s.GovernanceObservation != nil || s.Policy.Threshold > 0 && certificate == s.Policy.Threshold
	if status == StatusMinted || status == StatusRefundPending || status == StatusRefunded {
		if !authorized {
			return ErrInvalidSettlement.Wrap("missing retained quorum authorization")
		}
	} else if authorized {
		return ErrInvalidSettlement.Wrap("quorum must have an authorized financial status")
	}

	return nil
}

func policyHistory(policies []ReporterPolicy, current uint64) (map[uint64]ReporterPolicy, ReporterPolicy, error) {
	history := make(map[uint64]ReporterPolicy, len(policies))
	for _, policy := range policies {
		if err := policy.Validate(); err != nil {
			return nil, ReporterPolicy{}, err
		}

		if _, exists := history[policy.Version]; exists || policy.Version > current {
			return nil, ReporterPolicy{}, ErrInvalidGenesis.Wrap("duplicate or future reporter policy")
		}

		history[policy.Version] = policy
	}

	policy, exists := history[current]
	if !exists {
		return nil, ReporterPolicy{}, ErrInvalidGenesis.Wrap("missing current reporter policy")
	}

	return history, policy, nil
}
