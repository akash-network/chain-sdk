package v1

import (
	"bytes"
	"math"
	"slices"
	"strings"

	errorsmod "cosmossdk.io/errors"
)

const MaxReporters = 32

func ValidateReporters(members []string, threshold uint32) error {
	if len(members) == 0 {
		if threshold != 0 {
			return ErrInvalidReporters.Wrap("empty reporter set must have zero threshold")
		}
		return nil
	}
	// len(members) <= MaxReporters before either conversion.
	if len(members) > MaxReporters || threshold > uint32(len(members)) || 3*threshold <= 2*uint32(len(members)) { //nolint:gosec
		return ErrInvalidReporters.Wrapf("reporter threshold must exceed two thirds of at most %d members", MaxReporters)
	}
	seen := make(map[string]bool, len(members))
	for _, member := range members {
		if _, err := ParseAccount(member, AccountLength); err != nil {
			return errorsmod.Wrap(err, "reporter")
		}
		if seen[member] {
			return ErrInvalidReporters.Wrapf("duplicate reporter %s", member)
		}
		seen[member] = true
	}
	return nil
}

func (p Params) SameReporterSet(other Params) bool {
	if p.ReporterThreshold != other.ReporterThreshold || len(p.Reporters) != len(other.Reporters) {
		return false
	}
	for _, member := range p.Reporters {
		if !slices.Contains(other.Reporters, member) {
			return false
		}
	}
	return true
}

func (o Observation) Validate() error {
	if o.Outcome != Outcome_OUTCOME_MINTED && o.Outcome != Outcome_OUTCOME_VOIDED {
		return ErrInvalidObservation.Wrapf("unknown outcome %s", o.Outcome)
	}
	if o.Slot == 0 || len(o.TransactionSignature) != 64 || bytes.Equal(o.TransactionSignature, make([]byte, 64)) {
		return ErrInvalidObservation.Wrap("slot must be positive and the transaction signature 64 nonzero bytes")
	}
	return nil
}

func ValidateReason(reason string) error {
	if strings.TrimSpace(reason) == "" || len(reason) > 256 {
		return ErrInvalidReason.Wrap("reason must contain 1 to 256 bytes")
	}
	return nil
}

func (s Settlement) Terminal() bool {
	return s.Status == Status_STATUS_MINTED || s.Status == Status_STATUS_REFUNDED
}

func (s Settlement) Tallies() (minted, voided uint32) {
	// Only these two outcomes count; a future outcome must not reach either quorum.
	for _, vote := range s.Votes {
		switch vote.Observation.Outcome {
		case Outcome_OUTCOME_MINTED:
			minted++
		case Outcome_OUTCOME_VOIDED:
			voided++
		}
	}
	return
}

func (s Settlement) HasVoidAuthorization() bool {
	if s.GovernanceObservation != nil && s.GovernanceObservation.Outcome == Outcome_OUTCOME_VOIDED {
		return true
	}
	_, voided := s.Tallies()
	return s.Policy.Threshold > 0 && voided >= s.Policy.Threshold
}

func (s Settlement) Validate(epoch uint64) error {
	if s.Round == 0 || s.Policy.Epoch > epoch {
		return ErrInvalidSettlement.Wrap("zero round or future reporter epoch")
	}
	if s.Status < Status_STATUS_PENDING || s.Status > Status_STATUS_REFUNDED || s.Hold < Hold_HOLD_NONE || s.Hold > Hold_HOLD_CONFLICT {
		return ErrInvalidSettlement.Wrapf("unknown status %d or hold %d", s.Status, s.Hold)
	}
	if err := ValidateReporters(s.Policy.Members, s.Policy.Threshold); err != nil {
		return err
	}
	if s.Policy.Epoch == 0 && (len(s.Policy.Members) > 0 || len(s.Votes) > 0 || s.GovernanceObservation != nil) {
		return ErrInvalidSettlement.Wrap("missing report epoch")
	}
	if s.GovernanceObservation != nil {
		if err := s.GovernanceObservation.Validate(); err != nil {
			return err
		}
		if s.Status == Status_STATUS_PENDING || (s.Status == Status_STATUS_MINTED) != (s.GovernanceObservation.Outcome == Outcome_OUTCOME_MINTED) {
			return ErrInvalidSettlement.Wrap("governance observation does not match settlement")
		}
	}
	seen := make(map[string]map[Outcome]bool)
	for _, vote := range s.Votes {
		if !slices.Contains(s.Policy.Members, vote.Reporter) {
			return ErrInvalidSettlement.Wrap("vote outside reporter snapshot")
		}
		if err := vote.Observation.Validate(); err != nil {
			return err
		}
		if seen[vote.Reporter] == nil {
			seen[vote.Reporter] = make(map[Outcome]bool)
		}
		if seen[vote.Reporter][vote.Observation.Outcome] {
			return ErrInvalidSettlement.Wrap("duplicate reporter vote")
		}
		seen[vote.Reporter][vote.Observation.Outcome] = true
	}
	minted, voided := s.Tallies()
	if s.Hold != Hold_HOLD_CONFLICT && ((minted > 0 && voided > 0) ||
		(s.Status == Status_STATUS_MINTED && voided > 0) ||
		((s.Status == Status_STATUS_VOIDED || s.Status == Status_STATUS_REFUNDED) && minted > 0)) {
		return ErrInvalidSettlement.Wrap("conflicting reports without a hold")
	}
	if s.Status == Status_STATUS_VOIDED {
		if s.ChallengeEnd <= 0 || !s.HasVoidAuthorization() {
			return ErrInvalidSettlement.Wrap("void requires authorization and a challenge deadline")
		}
	} else if s.ChallengeEnd != 0 {
		return ErrInvalidSettlement.Wrap("challenge deadline on a non-void settlement")
	}
	if s.Hold == Hold_HOLD_VETO {
		if s.VetoBy == "" {
			return ErrInvalidSettlement.Wrap("missing veto signer")
		}
		if err := ValidateReason(s.VetoReason); err != nil {
			return err
		}
	}
	return nil
}

func checkedSum(values ...uint64) (uint64, error) {
	var total uint64
	for _, value := range values {
		if value > math.MaxUint64-total {
			return 0, ErrOverflow.Wrap("uakt total")
		}
		total += value
	}
	return total, nil
}

func (t Totals) DepositedAmount() (uint64, error) {
	return checkedSum(t.PendingAmount, t.MintedAmount, t.VoidedAmount, t.RefundedAmount)
}
func (t Totals) EscrowAmount() (uint64, error) {
	return checkedSum(t.PendingAmount, t.MintedAmount, t.VoidedAmount)
}
