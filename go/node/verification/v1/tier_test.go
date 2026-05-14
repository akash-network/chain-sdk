package v1

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// allTiers lists every VerificationTier in declaration order. The 5x5
// pair-wise tables below iterate this list against itself.
var allTiers = []VerificationTier{
	TierUnspecified,
	TierIdentified,
	TierVerified,
	TierEstablished,
	TierTrusted,
}

func TestTierAtLeast_Matrix(t *testing.T) {
	// Expected[i][j] = TierAtLeast(allTiers[i], allTiers[j]).
	// Definition: have != TierUnspecified && have >= need.
	// Row 0 (have == TierUnspecified) is always false.
	expected := [5][5]bool{
		// need: U     I      V      E      T
		/* have U */ {false, false, false, false, false},
		/* have I */ {true, true, false, false, false},
		/* have V */ {true, true, true, false, false},
		/* have E */ {true, true, true, true, false},
		/* have T */ {true, true, true, true, true},
	}

	for i, have := range allTiers {
		for j, need := range allTiers {
			want := expected[i][j]
			got := TierAtLeast(have, need)
			if got != want {
				t.Errorf("TierAtLeast(%v, %v) = %v; want %v", have, need, got, want)
			}
		}
	}
}

func TestTierBetter_Matrix(t *testing.T) {
	// Expected[i][j] = TierBetter(allTiers[i], allTiers[j]).
	// Definition: a != TierUnspecified && a > b.
	// Row 0 (a == TierUnspecified) is always false.
	expected := [5][5]bool{
		// b:    U     I      V      E      T
		/* a U */ {false, false, false, false, false},
		/* a I */ {true, false, false, false, false},
		/* a V */ {true, true, false, false, false},
		/* a E */ {true, true, true, false, false},
		/* a T */ {true, true, true, true, false},
	}

	for i, a := range allTiers {
		for j, b := range allTiers {
			want := expected[i][j]
			got := TierBetter(a, b)
			if got != want {
				t.Errorf("TierBetter(%v, %v) = %v; want %v", a, b, got, want)
			}
		}
	}
}

// TestTierAtLeast_UnspecifiedSentinel exercises the spec sentinel that
// TierUnspecified on the left is never sufficient, even against
// TierUnspecified on the right. AEP-86 IMPLEMENTATION.md §3.2.
func TestTierAtLeast_UnspecifiedSentinel(t *testing.T) {
	for _, need := range allTiers {
		if TierAtLeast(TierUnspecified, need) {
			t.Errorf("TierAtLeast(TierUnspecified, %v) = true; want false", need)
		}
	}
}

func TestTierRequiresSnapshot(t *testing.T) {
	cases := []struct {
		tier VerificationTier
		want bool
	}{
		{TierUnspecified, false},
		{TierIdentified, false},
		{TierVerified, true},
		{TierEstablished, true},
		{TierTrusted, true},
	}
	for _, tc := range cases {
		if got := TierRequiresSnapshot(tc.tier); got != tc.want {
			t.Errorf("TierRequiresSnapshot(%v) = %v; want %v", tc.tier, got, tc.want)
		}
	}
}

func TestTierRequiresProviderBond(t *testing.T) {
	cases := []struct {
		tier VerificationTier
		want bool
	}{
		{TierUnspecified, false},
		{TierIdentified, false},
		{TierVerified, true},
		{TierEstablished, true},
		{TierTrusted, true},
	}
	for _, tc := range cases {
		if got := TierRequiresProviderBond(tc.tier); got != tc.want {
			t.Errorf("TierRequiresProviderBond(%v) = %v; want %v", tc.tier, got, tc.want)
		}
	}
}

func TestMinBondForTier(t *testing.T) {
	const denom = "uakt"
	bondL1 := sdk.NewCoin(denom, math.NewInt(1))
	bondL2 := sdk.NewCoin(denom, math.NewInt(2))
	bondL3 := sdk.NewCoin(denom, math.NewInt(3))
	bondL4 := sdk.NewCoin(denom, math.NewInt(4))

	p := Params{
		BondL1: bondL1,
		BondL2: bondL2,
		BondL3: bondL3,
		BondL4: bondL4,
	}

	cases := []struct {
		name string
		tier VerificationTier
		want sdk.Coin
	}{
		{"L1 returns BondL1", TierIdentified, bondL1},
		{"L2 returns BondL2", TierVerified, bondL2},
		{"L3 returns BondL3", TierEstablished, bondL3},
		{"L4 returns BondL4", TierTrusted, bondL4},
		{"L0 returns zero coin in BondL1 denom", TierUnspecified, sdk.NewCoin(denom, math.ZeroInt())},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MinBondForTier(p, tc.tier)
			if !got.IsEqual(tc.want) {
				t.Errorf("MinBondForTier(%v) = %v; want %v", tc.tier, got, tc.want)
			}
		})
	}
}

// TestMinBondForTier_UnknownTierPanics asserts that an unrecognized tier value
// panics rather than silently returning a coin. Tier values flow from chain-
// validated proto fields, so an unknown value is a programmer bug or corrupted
// state — fail loud rather than authorize the wrong bond amount.
func TestMinBondForTier_UnknownTierPanics(t *testing.T) {
	p := Params{BondL1: sdk.NewCoin("uakt", math.NewInt(1))}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("MinBondForTier(unknown) did not panic")
		}
	}()
	_ = MinBondForTier(p, VerificationTier(99))
}

// TestMinBondForTier_L0_PreservesDenom proves the zero coin returned for L0
// carries Params.BondL1.Denom so callers comparing against a provider's bond
// in the same denom get a clean comparison.
func TestMinBondForTier_L0_PreservesDenom(t *testing.T) {
	const denom = "uakt"
	p := Params{BondL1: sdk.NewCoin(denom, math.NewInt(42))}
	got := MinBondForTier(p, TierUnspecified)
	if got.Denom != denom {
		t.Errorf("L0 zero coin denom = %q; want %q", got.Denom, denom)
	}
	if !got.Amount.IsZero() {
		t.Errorf("L0 zero coin amount = %v; want zero", got.Amount)
	}
}
