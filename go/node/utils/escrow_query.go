package utils

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func LeaseCalcBalanceRemain(balance sdkmath.LegacyDec, currBlock, settledAt int64, leasePrice sdk.DecCoin) sdk.DecCoin {
	res, _ := sdkmath.LegacyNewDecFromStr(balance.String())
	diff := sdkmath.LegacyZeroDec()

	diff = diff.Add(leasePrice.Amount)
	diff = diff.MulInt64(currBlock - settledAt)

	res = res.Sub(diff)

	return sdk.NewDecCoinFromDec(leasePrice.Denom, res)
}

func LeaseCalcBlocksRemain(balance sdkmath.LegacyDec, leasePrice sdkmath.LegacyDec) int64 {
	return balance.Quo(leasePrice).TruncateInt64()
}
