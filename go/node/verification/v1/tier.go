// Package v1 contains hand-written helpers that sit alongside the generated
// protobuf types for the akash.verification.v1 module.
//
// This file (tier.go) is NOT generated. It must not be deleted by the codegen
// pipeline, which only removes files matching *.pb.go and *.pb.gw.go.
package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TierAtLeast returns true iff `have` is at or above `need`. TierUnspecified
// is never sufficient: any comparison whose `have` is TierUnspecified returns
// false, even when `need` is also TierUnspecified. This matches AEP-86
// IMPLEMENTATION.md §3.2 ("Both must be valid tiers (not TierUnspecified)").
func TierAtLeast(have, need VerificationTier) bool {
	return have != TierUnspecified && have >= need
}

// TierBetter returns true iff `a` is strictly more trusted than `b`. As with
// TierAtLeast, TierUnspecified on the left-hand side is never sufficient and
// always returns false. See AEP-86 IMPLEMENTATION.md §3.2.
func TierBetter(a, b VerificationTier) bool {
	return a != TierUnspecified && a > b
}

// TierRequiresSnapshot reports whether tier `t` requires the provider to
// maintain a current snapshot-hash compliance record. Per AEP-86
// IMPLEMENTATION.md (Verification Tiers), L2 (TierVerified) and above require
// snapshot compliance; L0 (TierUnspecified) and L1 (TierIdentified) do not.
func TierRequiresSnapshot(t VerificationTier) bool {
	return t >= TierVerified
}

// TierRequiresProviderBond reports whether tier `t` requires the provider to
// post the resource-scaled provider bond. Per AEP-86, L2 (TierVerified) and
// above require a provider bond; L0 and L1 do not.
func TierRequiresProviderBond(t VerificationTier) bool {
	return t >= TierVerified
}

// MinBondForTier returns the auditor-bond amount required for tier `t`, drawn
// from the per-tier `BondL1`..`BondL4` fields of Params.
//
// TierUnspecified (L0) returns an error rather than a zero coin: L0 is the
// permissionless default with no attestation and therefore no auditor bond is
// associated with it. The spec (§Verification Tiers) treats L1 as the first
// tier with an economic floor; L0 has no bond requirement at all. Returning
// an error makes the "no bond exists for this tier" condition unambiguous at
// call sites — callers must not silently treat a zero coin as "no bond
// required".
//
// An unknown / out-of-range tier value also returns an error.
func MinBondForTier(p Params, t VerificationTier) (sdk.Coin, error) {
	switch t {
	case TierIdentified:
		return p.BondL1, nil
	case TierVerified:
		return p.BondL2, nil
	case TierEstablished:
		return p.BondL3, nil
	case TierTrusted:
		return p.BondL4, nil
	case TierUnspecified:
		return sdk.Coin{}, fmt.Errorf("no bond requirement for tier %v", t)
	default:
		return sdk.Coin{}, fmt.Errorf("unknown verification tier %v", t)
	}
}
