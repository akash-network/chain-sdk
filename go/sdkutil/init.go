package sdkutil

import (
	"cosmossdk.io/math"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	DenomAkt  = "akt"  // 1akt
	DenomMakt = "makt" // 10^-3akt
	DenomUakt = "uakt" // 10^-6akt
	BondDenom = DenomUakt

	DenomMaktExponent = 3
	DenomUaktExponent = 6

	Bech32PrefixAccAddr = "akash"
	Bech32PrefixAccPub  = "akashpub"

	Bech32PrefixValAddr = "akashvaloper"
	Bech32PrefixValPub  = "akashvaloperpub"

	Bech32PrefixConsAddr = "akashvalcons"
	Bech32PrefixConsPub  = "akashvalconspub"
)

func init() {
	aktUnit := math.LegacyOneDec()                              // 1 (base denom unit)
	maktUnit := math.LegacyNewDecWithPrec(1, DenomMaktExponent) // 10^-6 (micro)
	uaktUnit := math.LegacyNewDecWithPrec(1, DenomUaktExponent) // 10^-6 (micro)

	err := sdktypes.RegisterDenom(DenomAkt, aktUnit)
	if err != nil {
		panic(err)
	}

	err = sdktypes.RegisterDenom(DenomMakt, maktUnit)
	if err != nil {
		panic(err)
	}

	err = sdktypes.RegisterDenom(DenomUakt, uaktUnit)
	if err != nil {
		panic(err)
	}
}
