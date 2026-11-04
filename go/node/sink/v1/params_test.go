package v1_test

import (
	"math"
	"math/big"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/sink/v1"
	"pkg.akt.dev/go/sdkutil"
)

func TestRatioAndCapEdges(t *testing.T) {
	require.NoError(t, types.DefaultGenesis().Validate())
	require.False(t, types.DefaultParams().Enabled)
	require.False(t, types.DefaultParams().RefundsEnabled)

	for _, ratio := range []string{"1.000000000000000001", "-1"} {
		p := types.DefaultParams()
		p.Reporters = policy().Members
		p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr(ratio)
		require.Error(t, p.ValidateBasic(), ratio)
	}

	for _, ratio := range []string{"0", "0.5", "0.500000000000000001", "0.6", "0.666666666666666666", "0.666666666666666667", "0.75", "1"} {
		p := types.DefaultParams()
		p.Reporters = policy().Members
		p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr(ratio)
		require.NoError(t, p.ValidateBasic(), ratio)
	}

	p := types.DefaultParams()
	p.Reporters = policy().Members
	p.Reporters = append(p.Reporters, p.Reporters[0])
	require.Error(t, types.ValidateReporters(p.Reporters, p.ReportersThreshold))
	p = types.DefaultParams()
	p.Enabled = true
	require.Error(t, p.ValidateBasic())

	for name, change := range map[string]func(*types.Params){
		"nil coin":           func(p *types.Params) { p.MaxTransferAmountPerEpoch = sdk.Coin{} },
		"nil account cap":    func(p *types.Params) { p.MaxEpochAmountPerAccount = sdk.Coin{} },
		"account cap denom":  func(p *types.Params) { p.MaxEpochAmountPerAccount.Denom = "other" },
		"zero cap":           func(p *types.Params) { p.MaxEpochRefund = types.ZeroAmount() },
		"denom":              func(p *types.Params) { p.MaxAmountPerTransfer.Denom = "other" },
		"min over account":   func(p *types.Params) { p.MaxEpochAmountPerAccount = sdk.NewInt64Coin(sdkutil.DenomUakt, 1) },
		"max over refund":    func(p *types.Params) { p.MaxEpochRefund = p.MinAmountPerTransfer },
		"min over aggregate": func(p *types.Params) { p.MaxTransferAmountPerEpoch = sdk.NewInt64Coin(sdkutil.DenomUakt, 1) },
		"min over max":       func(p *types.Params) { p.MinAmountPerTransfer = p.MaxTransferAmountPerEpoch },
		"uint64 overflow": func(p *types.Params) {
			p.MaxAmountPerTransfer = sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(math.MaxUint64).AddRaw(1))
		},
		"challenge floor": func(p *types.Params) {
			p.ReporterChallengePeriod = time.Duration(types.MinChallengePeriodSeconds)*time.Second - 1
		},
	} {
		t.Run(name, func(t *testing.T) {
			p := types.DefaultParams()
			change(&p)
			require.Error(t, p.ValidateBasic())
		})
	}

	p = types.DefaultParams()
	p.MinAmountPerTransfer = p.MaxAmountPerTransfer
	p.MaxTransferAmountPerEpoch = p.MaxAmountPerTransfer
	p.ReporterChallengePeriod = time.Duration(types.MinChallengePeriodSeconds) * time.Second
	require.NoError(t, p.ValidateBasic())
	p.MaxEpochAmountPerAccount = types.ZeroAmount()
	require.NoError(t, p.ValidateBasic())
}

func TestIndependentSendingCapsAndDisabledMinimum(t *testing.T) {
	p := types.DefaultParams()
	p.MaxEpochAmountPerAccount = p.MinAmountPerTransfer
	p.MaxTransferAmountPerEpoch = p.MinAmountPerTransfer
	require.NoError(t, p.ValidateBasic())

	p.MinAmountPerTransfer = types.ZeroAmount()
	require.NoError(t, p.ValidateBasic())
	p.Reporters = policy().Members
	p.Enabled = true
	require.Error(t, p.ValidateBasic())
	p.MinAmountPerTransfer = sdk.NewInt64Coin(sdkutil.DenomUakt, 1)
	require.NoError(t, p.ValidateBasic())

	p.Reporters = p.Reporters[:3]
	p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr("0.666666666666666666")
	require.NoError(t, p.ValidateBasic())
	p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr("0.666666666666666667")
	require.NoError(t, p.ValidateBasic())
}

func TestIndependentEpochsAndCoinOverflow(t *testing.T) {
	amount := sdk.NewInt64Coin(sdkutil.DenomUakt, 123)
	send := types.EpochUsage{
		Epoch:  8,
		Amount: amount,
	}
	refund := types.EpochUsage{
		Epoch:  9,
		Amount: amount,
	}
	normalized, err := send.At(9)
	require.NoError(t, err)
	require.True(t, normalized.Amount.IsZero())
	require.Equal(t, uint64(9), normalized.Epoch)
	normalized, err = refund.At(9)
	require.NoError(t, err)
	require.Equal(t, amount, normalized.Amount)
	_, err = refund.At(8)
	require.ErrorIs(t, err, types.ErrInvalidClock)
	require.Equal(t, amount, send.Amount)
	account := types.AccountUsage{
		TotalSent: amount,
		EpochSent: amount,
		LastEpoch: 8,
	}
	usage, err := account.At(9)
	require.NoError(t, err)
	require.True(t, usage.EpochSent.IsZero())
	require.Equal(t, amount, usage.TotalSent)
	_, err = account.At(7)
	require.ErrorIs(t, err, types.ErrInvalidClock)
	_, err = account.At(math.MaxInt64 + 1)
	require.ErrorIs(t, err, types.ErrInvalidClock)
	_, err = send.At(math.MaxInt64 + 1)
	require.ErrorIs(t, err, types.ErrInvalidClock)
	for _, usage := range []types.EpochUsage{
		{
			Epoch:  0,
			Amount: amount,
		},
		{
			Epoch:  math.MaxInt64 + 1,
			Amount: types.ZeroAmount(),
		},
	} {
		require.ErrorIs(t, usage.Validate(), types.ErrInvalidClock)
	}

	account.LastEpoch = 0
	require.ErrorIs(t, account.Validate(), types.ErrInvalidClock)
	zero := types.EpochUsage{Amount: types.ZeroAmount()}
	normalized, err = zero.At(0)
	require.NoError(t, err)
	require.Equal(t, zero, normalized)
	zeroAccount := types.AccountUsage{
		TotalSent: types.ZeroAmount(),
		EpochSent: types.ZeroAmount(),
	}
	usage, err = zeroAccount.At(0)
	require.NoError(t, err)
	require.Equal(t, zeroAccount, usage)
	account.LastEpoch = -1
	_, err = account.At(9)
	require.ErrorIs(t, err, types.ErrInvalidClock)
	a := sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(math.MaxUint64))
	sum, err := types.AddAmount(a, a)
	require.NoError(t, err)
	require.False(t, sum.Amount.IsUint64())
	max := sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromBigInt(new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))))
	_, err = types.AddAmount(max, amount)
	require.ErrorIs(t, err, types.ErrOverflow)
	totals := types.DefaultTotals()
	totals.PendingCount = math.MaxUint64
	require.ErrorIs(t, totals.Add(types.StatusTransferPending, amount), types.ErrOverflow)
	require.True(t, totals.PendingAmount.IsZero())
}

func TestTotalsSubRejectsRecordsOutsideThePopulation(t *testing.T) {
	amount := sdk.NewInt64Coin(sdkutil.DenomUakt, 10)
	totals := types.DefaultTotals()
	require.NoError(t, totals.Add(types.StatusRefundPending, amount))
	before := totals

	for name, tc := range map[string]struct {
		status types.Status
		amount sdk.Coin
	}{
		"empty population":   {types.StatusMinted, amount},
		"larger amount":      {types.StatusRefundPending, amount.Add(amount)},
		"foreign denom":      {types.StatusRefundPending, sdk.NewInt64Coin("other", 10)},
		"uninitialized coin": {types.StatusRefundPending, sdk.Coin{Denom: sdkutil.DenomUakt}},
	} {
		require.Error(t, totals.Sub(tc.status, tc.amount), name)
		require.Equal(t, before, totals, name)
	}

	require.ErrorIs(t, totals.Sub(types.StatusInvalid, amount), types.ErrInvalidSettlement)
	require.Equal(t, before, totals)

	require.NoError(t, totals.Sub(types.StatusRefundPending, amount))
	require.Zero(t, totals.RefundPendingCount)
	require.True(t, totals.RefundPendingAmount.IsZero())

	require.NoError(t, totals.Add(types.StatusVoided, amount))
	require.Equal(t, uint64(1), totals.VoidedCount)
	require.Equal(t, amount, totals.VoidedAmount)
	require.NoError(t, totals.Sub(types.StatusVoided, amount))
	require.True(t, types.DefaultTotals().Equal(totals))
}

func TestReporterPolicyAndMaximumTransferBoundaries(t *testing.T) {
	p := types.DefaultParams()
	for i := 1; i <= 33; i++ {
		p.Reporters = append(p.Reporters, account(byte(i)))
	}

	p.ReportersThreshold = sdkmath.LegacyMustNewDecFromStr("0.75")
	require.NoError(t, p.ValidateBasic())
	reporters := types.ReporterPolicy{
		Version:   1,
		Members:   p.Reporters,
		Threshold: p.ReportersThreshold,
	}
	require.NoError(t, reporters.Validate())

	maximum := sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(math.MaxUint64))
	p.MaxAmountPerTransfer, p.MaxEpochAmountPerAccount, p.MaxTransferAmountPerEpoch, p.MaxEpochRefund = maximum, maximum, maximum, maximum
	require.NoError(t, p.ValidateBasic())
}
