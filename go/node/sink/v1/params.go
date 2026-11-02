package v1

import (
	"math"

	errorsmod "cosmossdk.io/errors"
)

// DefaultChallengePeriodSeconds is the 72-hour window AEP-92 assumes.
const DefaultChallengePeriodSeconds uint64 = 72 * 60 * 60

// Defaults are local-development limits, not production policy. Deposits are off.
func DefaultParams() Params {
	return Params{MaxTransfer: 1_000_000_000_000, SolanaDailyMintCap: 10_000_000_000_000,
		MaxDailyPerAccount: 20, MaxPendingPerAccount: 3,
		ChallengePeriodSeconds: DefaultChallengePeriodSeconds, DailyRefundCap: 1_000_000_000_000}
}

func (p Params) Validate() error {
	if err := ValidateReporters(p.Reporters, p.ReporterThreshold); err != nil {
		return err
	}
	if p.Enabled && len(p.Reporters) == 0 {
		return ErrInvalidParam.Wrap("reporters must be configured before enabling deposits")
	}
	// The floor applies even while refunds are off: a void confirmed under a
	// shorter window would keep its deadline after refunds are enabled.
	if p.ChallengePeriodSeconds < MinChallengePeriodSeconds || p.ChallengePeriodSeconds > math.MaxInt64 {
		return ErrInvalidParam.Wrapf("challenge period must be %d to %d seconds", MinChallengePeriodSeconds, uint64(math.MaxInt64))
	}
	if p.DailyRefundCap < p.MaxTransfer {
		return ErrInvalidParam.Wrap("daily refund cap must be at least max transfer")
	}
	if p.MaxTransfer == 0 || p.MaxTransfer > p.SolanaDailyMintCap {
		return ErrInvalidParam.Wrap("max transfer must be positive and no larger than Solana daily mint cap")
	}
	// Zero disables the minimum, preserving parameters stored before this field,
	// but deposits need one: it keeps each worth the Solana rent of settling it.
	// Disabling deposits never fails here, so a guardian pause always works.
	if p.MinTransfer > p.MaxTransfer {
		return ErrInvalidParam.Wrap("min transfer must not exceed max transfer")
	}
	if p.Enabled && p.MinTransfer == 0 {
		return ErrInvalidParam.Wrap("min transfer must be set while deposits are enabled")
	}
	// Zero disables the per-account daily amount cap. A nonzero cap below the
	// minimum would reject every deposit.
	if p.MaxDailyAmountPerAccount != 0 && p.MaxDailyAmountPerAccount < p.MinTransfer {
		return ErrInvalidParam.Wrap("max daily amount per account must be zero or at least min transfer")
	}
	// Empty means no guardian; only governance can then pause.
	if p.PauseGuardian != "" {
		if err := ValidatePauseGuardian(p.PauseGuardian); err != nil {
			return err
		}
	}
	if p.MaxDailyPerAccount == 0 || p.MaxPendingPerAccount == 0 {
		return ErrInvalidParam.Wrap("account limits must be positive")
	}
	if len(p.VaultAllowlist) > 16 {
		return ErrInvalidParam.Wrap("at most 16 vaults")
	}
	seen := make(map[[32]byte]bool)
	for _, v := range p.VaultAllowlist {
		if len(v) != 32 {
			return ErrInvalidParam.Wrap("vault must be 32 bytes")
		}
		key := [32]byte(v)
		if key == [32]byte{} || seen[key] {
			return ErrInvalidParam.Wrap("zero or duplicate vault")
		}
		seen[key] = true
	}
	return nil
}

// ValidatePauseGuardian accepts a canonical, nonzero account of an ordinary
// or derived length, such as a group policy.
func ValidatePauseGuardian(guardian string) error {
	if _, err := ParseAccount(guardian, AccountLength, DerivedAccountLength); err != nil {
		return errorsmod.Wrap(err, "pause guardian")
	}
	return nil
}
