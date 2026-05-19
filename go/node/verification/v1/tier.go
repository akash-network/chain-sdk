package v1

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TierAtLeast returns true when have satisfies need.
func TierAtLeast(have, need VerificationTier) bool {
	return have != TierUnspecified && have >= need
}

// TierBetter returns true when a is strictly higher than b.
func TierBetter(a, b VerificationTier) bool {
	return a != TierUnspecified && a > b
}

// TierRequiresSnapshot reports whether t requires snapshot compliance.
func TierRequiresSnapshot(t VerificationTier) bool {
	return t >= TierVerified
}

// TierRequiresProviderBond reports whether t requires a provider bond.
func TierRequiresProviderBond(t VerificationTier) bool {
	return t >= TierVerified
}

// MinBondForTier returns the auditor bond amount required for t.
func MinBondForTier(p Params, t VerificationTier) sdk.Coin {
	switch t {
	case TierIdentified:
		return p.BondL1
	case TierVerified:
		return p.BondL2
	case TierEstablished:
		return p.BondL3
	case TierTrusted:
		return p.BondL4
	case TierUnspecified:
		return sdk.NewCoin(p.BondL1.Denom, math.ZeroInt())
	default:
		panic(fmt.Sprintf("verification: unknown tier %v", t))
	}
}
