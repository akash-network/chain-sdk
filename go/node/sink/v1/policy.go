package v1

import (
	"slices"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateReporters(members []string, threshold math.LegacyDec) error {
	if threshold.IsNil() || threshold.IsNegative() || threshold.GT(math.LegacyOneDec()) {
		return ErrInvalidReporters.Wrap("threshold must be between zero and one")
	}

	n := int64(len(members))
	required := threshold.MulInt64(n).Ceil().TruncateInt64()

	if n > 0 && 3*required <= 2*n {
		return ErrInvalidReporters.Wrap("threshold must exceed two thirds and not exceed membership")
	}

	seen := make(map[string]bool, len(members))
	for _, member := range members {
		addr, err := sdk.AccAddressFromBech32(member)
		if err != nil {
			return err
		}

		// Bech32 also decodes the upper-case spelling of an address, so
		// duplicates are found by address rather than by string.
		key := addr.String()
		if seen[key] {
			return ErrInvalidReporters.Wrap("duplicate reporter")
		}

		seen[key] = true
	}

	return nil
}

func (p *ReporterPolicy) Validate() error {
	if p.Version == 0 {
		return ErrInvalidReporters.Wrap("policy version must be positive")
	}

	return ValidateReporters(p.Members, p.Threshold)
}

// slices expected to be sorted
func (p *ReporterPolicy) Equal(other ReporterPolicy) bool {
	if p.Threshold.IsNil() || other.Threshold.IsNil() || !p.Threshold.Equal(other.Threshold) {
		return false
	}

	return slices.Equal(p.Members, other.Members)
}
