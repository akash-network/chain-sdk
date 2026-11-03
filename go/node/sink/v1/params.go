package v1

import (
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/x/tx/signing/aminojson"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.akt.dev/go/sdkutil"
)

const DefaultChallengePeriodSeconds uint64 = 72 * 60 * 60

// MaxChallengePeriod includes the final fractional second accepted by native Amino signing.
const MaxChallengePeriod = time.Duration(aminojson.MaxDurationSeconds)*time.Second + time.Second - time.Nanosecond

// DefaultParams is a development policy; deposits and refunds start disabled.
func DefaultParams() Params {
	return Params{
		MinAmountPerTransfer:      sdk.NewInt64Coin(sdkutil.DenomUakt, 1_000_000),
		MaxAmountPerTransfer:      sdk.NewInt64Coin(sdkutil.DenomUakt, 1_000_000_000_000),
		MaxTransferAmountPerEpoch: sdk.NewInt64Coin(sdkutil.DenomUakt, 10_000_000_000_000),
		MaxEpochAmountPerAccount:  sdk.NewInt64Coin(sdkutil.DenomUakt, 1_000_000_000_000),
		MaxEpochRefund:            sdk.NewInt64Coin(sdkutil.DenomUakt, 1_000_000_000_000),
		ReportersThreshold:        sdkmath.LegacyMustNewDecFromStr("0.75"),
		ReporterChallengePeriod:   time.Duration(DefaultChallengePeriodSeconds) * time.Second,
	}
}

func (p Params) ValidateBasic() error {
	for _, cap := range []sdk.Coin{p.MinAmountPerTransfer, p.MaxAmountPerTransfer, p.MaxTransferAmountPerEpoch, p.MaxEpochAmountPerAccount, p.MaxEpochRefund} {
		if err := ValidateAmount(cap); err != nil {
			return ErrInvalidParam.Wrap(err.Error())
		}
	}

	if !p.MaxAmountPerTransfer.IsPositive() || !p.MaxTransferAmountPerEpoch.IsPositive() || !p.MaxEpochRefund.IsPositive() {
		return ErrInvalidParam.Wrap("transfer, send and refund caps must be positive")
	}

	if p.Enabled && p.MinAmountPerTransfer.IsZero() {
		return ErrInvalidParam.Wrap("minimum transfer must be positive while deposits are enabled")
	}

	if !p.MaxAmountPerTransfer.Amount.IsUint64() || p.MinAmountPerTransfer.IsGT(p.MaxAmountPerTransfer) ||
		(p.MaxEpochAmountPerAccount.IsPositive() && p.MinAmountPerTransfer.IsGT(p.MaxEpochAmountPerAccount)) ||
		p.MinAmountPerTransfer.IsGT(p.MaxTransferAmountPerEpoch) || p.MaxAmountPerTransfer.IsGT(p.MaxEpochRefund) {
		return ErrInvalidParam.Wrap("minimum must fit transfer and sending caps; maximum must fit uint64 and refund cap")
	}

	// Bound the decimal before multiplying; arbitrary protobuf decimals may overflow.
	if p.ReportersThreshold.IsNil() || p.ReportersThreshold.IsNegative() || p.ReportersThreshold.GT(sdkmath.LegacyOneDec()) {
		return ErrInvalidReporters.Wrap("reporters_threshold must be between zero and one")
	}

	if len(p.Reporters) > MaxReporters {
		return ErrInvalidReporters.Wrap("too many reporters")
	}

	if err := p.Policy(1).Validate(); err != nil {
		return err
	}

	if p.Enabled && len(p.Reporters) == 0 {
		return ErrInvalidParam.Wrap("deposits require reporters")
	}

	if p.ReporterChallengePeriod < time.Duration(MinChallengePeriodSeconds)*time.Second {
		return ErrInvalidParam.Wrap("challenge period below minimum")
	}

	if p.ReporterChallengePeriod > MaxChallengePeriod {
		return ErrInvalidParam.Wrap("challenge period exceeds Amino signing limit")
	}

	if p.PauseGuardian != "" {
		if err := ValidatePauseGuardian(p.PauseGuardian); err != nil {
			return err
		}
	}

	seen := make(map[SolanaAddress]bool, len(p.AllowedDestinationVaults))
	for _, v := range p.AllowedDestinationVaults {
		addr, err := ParseSolanaAddress(v)
		if err != nil {
			return err
		}

		if seen[addr] {
			return ErrInvalidParam.Wrap("duplicate destination vault")
		}

		seen[addr] = true
	}

	return nil
}

func ValidatePauseGuardian(guardian string) error {
	_, err := ParseAccount(guardian, AccountLength, DerivedAccountLength)
	return err
}
