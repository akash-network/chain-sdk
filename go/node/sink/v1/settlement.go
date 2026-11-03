package v1

import (
	"bytes"
	"slices"
	"strings"
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

	return o.Evidence.Validate()
}

func ValidateReason(reason string) error {
	if !utf8.ValidString(reason) || strings.TrimSpace(reason) == "" || len(reason) > 256 {
		return ErrInvalidReason.Wrap("reason must contain 1 to 256 UTF-8 bytes")
	}

	return nil
}

func (s Status) IsValid() bool {
	switch s {
	case StatusTransferPending, StatusMinted, StatusRefundPending, StatusRefunded:
		return true
	default:
		return false
	}
}

func (s SettlementData) Terminal() bool {
	return s.Status == StatusMinted || s.Status == StatusRefunded
}

// Tallies counts only votes under the settlement's authorization policy.
func (s SettlementData) Tallies() (minted, voided uint32) {
	for _, vote := range s.Votes {
		if vote.PolicyVersion != s.PolicyVersion {
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

func (s SettlementData) HasVoidAuthorization(policy ReporterPolicy) bool {
	if s.Status != StatusRefundPending && s.Status != StatusRefunded {
		return false
	}

	if s.GovernanceEvidence != nil {
		return true
	}

	_, voided := s.Tallies()
	return policy.Version == s.PolicyVersion && policy.Threshold > 0 && voided >= policy.Threshold
}

func (s SettlementData) Payable(policy ReporterPolicy, current uint64) bool {
	return s.Status == StatusRefundPending && s.Hold == HoldNone && (s.GovernanceEvidence != nil || s.PolicyVersion == current && s.HasVoidAuthorization(policy))
}

func (s SettlementData) RefundScope() (scope uint64, ok bool) {
	if s.Status != StatusRefundPending || s.Hold != HoldNone {
		return 0, false
	}

	if s.GovernanceEvidence != nil {
		return GovernanceRefundScope, true
	}

	return s.PolicyVersion, true
}

// Validate checks only the policy versions s references. Policies are validated when stored.
func (s SettlementData) Validate(policies map[uint64]ReporterPolicy) error {
	if s.Round == 0 {
		return ErrInvalidSettlement.Wrap("round must be positive")
	}

	policy, ok := policies[s.PolicyVersion]
	if !ok {
		return ErrInvalidSettlement.Wrap("missing authorization policy")
	}

	if !s.Status.IsValid() {
		return ErrInvalidSettlement.Wrap("unknown financial status")
	}

	if s.Hold != HoldNone && s.Hold != HoldVeto && s.Hold != HoldConflict {
		return ErrInvalidSettlement.Wrap("unknown hold")
	}

	if (s.Status == StatusRefundPending) != (s.ChallengeEnd != nil) {
		return ErrInvalidSettlement.Wrap("deadline required exactly for refund-pending")
	}

	if s.ChallengeEnd != nil {
		if _, err := gogotypes.TimestampProto(*s.ChallengeEnd); err != nil || s.ChallengeEnd.Unix() <= 0 {
			return ErrInvalidSettlement.Wrap("invalid challenge deadline")
		}
	}

	if s.Veto != nil {
		if s.Terminal() || s.Hold == HoldNone {
			return ErrInvalidSettlement.Wrap("orphaned veto")
		}

		if _, err := ParseAccount(s.Veto.Signer, AccountLength, DerivedAccountLength); err != nil {
			return err
		}

		if err := ValidateReason(s.Veto.Reason); err != nil {
			return err
		}
	}

	if s.Hold == HoldVeto && s.Veto == nil {
		return ErrInvalidSettlement.Wrap("veto hold requires signer and reason")
	}

	if s.Hold == HoldVeto && s.Terminal() {
		return ErrInvalidSettlement.Wrap("terminal status cannot have veto hold")
	}

	if s.GovernanceEvidence != nil {
		if s.Status == StatusTransferPending || s.Round == 1 {
			return ErrInvalidSettlement.Wrap("orphaned governance evidence")
		}

		if err := s.GovernanceEvidence.Validate(); err != nil {
			return err
		}
	}

	type voteKey struct {
		reporter string
		outcome  Outcome
		version  uint64
	}

	seen := make(map[voteKey]bool, len(s.Votes))
	var hasMint, hasVoid bool
	for _, vote := range s.Votes {
		p, exists := policies[vote.PolicyVersion]
		if !exists || !slices.Contains(p.Members, vote.Reporter) {
			return ErrInvalidSettlement.Wrap("vote outside its reporter policy")
		}

		if err := vote.Observation.Validate(); err != nil {
			return err
		}

		key := voteKey{vote.Reporter, vote.Observation.Outcome, vote.PolicyVersion}
		if seen[key] {
			return ErrInvalidSettlement.Wrap("duplicate reporter/outcome/policy vote")
		}

		seen[key] = true
		hasMint = hasMint || vote.Observation.Outcome == OutcomeMinted
		hasVoid = hasVoid || vote.Observation.Outcome == OutcomeVoided
	}

	if s.Hold != HoldConflict && (hasMint && hasVoid || s.Status == StatusMinted && hasVoid || (s.Status == StatusRefundPending || s.Status == StatusRefunded) && hasMint) {
		return ErrInvalidSettlement.Wrap("opposing evidence requires conflict hold")
	}

	minted, voided := s.Tallies()
	if s.Status != StatusTransferPending && s.GovernanceEvidence == nil {
		count := voided
		if s.Status == StatusMinted {
			count = minted
		}

		if policy.Threshold == 0 || count < policy.Threshold {
			return ErrInvalidSettlement.Wrap("missing retained quorum authorization")
		}
	}

	if s.Status == StatusTransferPending && s.Hold == HoldNone && policy.Threshold > 0 && (minted >= policy.Threshold || voided >= policy.Threshold) {
		return ErrInvalidSettlement.Wrap("unheld quorum must have a settled financial status")
	}

	return nil
}

func policyHistory(policies []ReporterPolicy) (map[uint64]ReporterPolicy, ReporterPolicy, error) {
	history := make(map[uint64]ReporterPolicy, len(policies))
	for _, policy := range policies {
		if err := policy.Validate(); err != nil {
			return nil, ReporterPolicy{}, err
		}

		if _, exists := history[policy.Version]; exists {
			return nil, ReporterPolicy{}, ErrInvalidGenesis.Wrap("duplicate reporter policy")
		}

		history[policy.Version] = policy
	}

	if len(history) == 0 {
		return nil, ReporterPolicy{}, ErrInvalidGenesis.Wrap("missing reporter policy history")
	}

	for version := uint64(1); version <= uint64(len(history)); version++ {
		if _, exists := history[version]; !exists {
			return nil, ReporterPolicy{}, ErrInvalidGenesis.Wrap("reporter policy history must contain versions 1 through current")
		}
	}

	return history, history[uint64(len(history))], nil
}
