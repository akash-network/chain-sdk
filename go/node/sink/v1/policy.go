package v1

import (
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateReporters(members []string, threshold uint32) error {
	n := int64(len(members))
	required := int64(threshold)

	if n == 0 && required != 0 || n > 0 && (required > n || 3*required <= 2*n) {
		return ErrInvalidReporters.Wrap("threshold must exceed two thirds and not exceed membership")
	}

	seen := make(map[string]bool, len(members))
	for _, member := range members {
		if _, err := sdk.AccAddressFromBech32(member); err != nil {
			return err
		}

		if seen[member] {
			return ErrInvalidReporters.Wrap("duplicate reporter")
		}

		seen[member] = true
	}

	return nil
}

func (p *ReporterPolicy) Validate() error {
	if p.Version == 0 {
		return ErrInvalidReporters.Wrap("policy version must be positive")
	}

	return ValidateReporters(p.Members, p.Threshold)
}

// Policy derives a policy from validated parameters using exact decimal arithmetic.
func (p *Params) Policy(version uint64) ReporterPolicy {
	var threshold uint32
	if len(p.Reporters) > 0 {
		threshold = uint32(p.ReportersThreshold.MulInt64(int64(len(p.Reporters))).Ceil().TruncateInt64()) //nolint:gosec // Validated parameters bound quorum to 32.
	}

	return ReporterPolicy{
		Version:   version,
		Members:   slices.Clone(p.Reporters),
		Threshold: threshold,
	}
}

func (p *ReporterPolicy) SameRule(other ReporterPolicy) bool {
	if p.Threshold != other.Threshold || len(p.Members) != len(other.Members) {
		return false
	}

	for _, member := range p.Members {
		if !slices.Contains(other.Members, member) {
			return false
		}
	}

	return true
}
