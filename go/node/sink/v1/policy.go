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

// slices expected to be sorted
func (p *ReporterPolicy) Equal(other ReporterPolicy) bool {
	if p.Threshold != other.Threshold {
		return false
	}

	return slices.Equal(p.Members, other.Members)
}
