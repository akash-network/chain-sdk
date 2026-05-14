// Package v1 contains hand-written helpers that sit alongside the generated
// protobuf types for the akash.verification.v1 module.
//
// This file (tier.go) is NOT generated. It must not be deleted by the codegen
// pipeline, which only removes files matching *.pb.go and *.pb.gw.go.
package v1

import (
	"fmt"

	"cosmossdk.io/math"
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
// from the per-tier `BondL1`..`BondL4` fields of Params. The signature matches
// AEP-86 IMPLEMENTATION.md §3.2 verbatim: a single sdk.Coin return value with
// no error path.
//
// Semantics:
//   - TierIdentified..TierTrusted return the matching Params.BondLn.
//   - TierUnspecified (L0) returns a **zero coin** in the params bond denom.
//     L0 is the permissionless default; no auditor bond is associated with it
//     per the spec. Callers comparing a provider's bond against the result
//     get the natural short-circuit (every coin is >= zero) without an
//     error-handling branch.
//   - An unknown / out-of-range tier value **panics**. Tier values flow only
//     from chain-validated proto fields, so an unrecognized value indicates
//     a programmer bug or a corrupted state read — fail loudly rather than
//     silently authorize the wrong bond amount.
//
// The bond denom is taken from Params.BondL1 when the tier is L0 (so the zero
// coin has the same denom as a real bond requirement). If Params is freshly
// constructed and BondL1.Denom is empty, the returned coin uses the empty
// denom — callers must validate Params at genesis time.
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
		panic(fmt.Sprintf("verification: MinBondForTier: unknown tier %v — programmer bug or corrupted state", t))
	}
}
