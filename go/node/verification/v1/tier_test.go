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
		name    string
		tier    VerificationTier
		want    sdk.Coin
		wantErr bool
	}{
		{"L1 returns BondL1", TierIdentified, bondL1, false},
		{"L2 returns BondL2", TierVerified, bondL2, false},
		{"L3 returns BondL3", TierEstablished, bondL3, false},
		{"L4 returns BondL4", TierTrusted, bondL4, false},
		{"L0 returns error", TierUnspecified, sdk.Coin{}, true},
		{"unknown tier returns error", VerificationTier(99), sdk.Coin{}, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := MinBondForTier(p, tc.tier)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("MinBondForTier(%v) returned no error; want error", tc.tier)
				}
				return
			}
			if err != nil {
				t.Fatalf("MinBondForTier(%v) returned error %v; want nil", tc.tier, err)
			}
			if !got.IsEqual(tc.want) {
				t.Errorf("MinBondForTier(%v) = %v; want %v", tc.tier, got, tc.want)
			}
		})
	}
}
