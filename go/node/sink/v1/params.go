package v1

import (
	"bytes"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Defaults are local-development limits, not production policy. Deposits are off.
func DefaultParams() Params {
	return Params{MaxTransfer: 1_000_000_000, SolanaDailyMintCap: 10_000_000_000,
		MaxDailyPerAccount: 20, MaxPendingPerAccount: 3,
		ChallengePeriodSeconds: 86_400, DailyRefundCap: 1_000_000_000}
}

func (p Params) Validate() error {
	if err := ValidateReporters(p.Reporters, p.ReporterThreshold); err != nil {
		return err
	}
	if p.Enabled && len(p.Reporters) == 0 {
		return fmt.Errorf("reporters must be configured before enabling deposits")
	}
	if p.ChallengePeriodSeconds == 0 || p.ChallengePeriodSeconds > math.MaxInt64 || p.DailyRefundCap < p.MaxTransfer {
		return fmt.Errorf("positive bounded challenge duration and refund cap at least max transfer required")
	}
	if p.MaxTransfer == 0 || p.MaxTransfer > p.SolanaDailyMintCap {
		return fmt.Errorf("max transfer must be positive and no larger than Solana daily mint cap")
	}
	// Zero disables the minimum, preserving parameters stored before this field.
	if p.MinTransfer > p.MaxTransfer {
		return fmt.Errorf("min transfer must not exceed max transfer")
	}
	// Zero disables the per-account daily amount cap. A nonzero cap below the
	// minimum would reject every deposit.
	if p.MaxDailyAmountPerAccount != 0 && p.MaxDailyAmountPerAccount < p.MinTransfer {
		return fmt.Errorf("max daily amount per account must be zero or at least min transfer")
	}
	// Empty means no guardian; only governance can then pause.
	if p.PauseGuardian != "" {
		if err := ValidatePauseGuardian(p.PauseGuardian); err != nil {
			return err
		}
	}
	if p.MaxDailyPerAccount == 0 || p.MaxPendingPerAccount == 0 {
		return fmt.Errorf("account limits must be positive")
	}
	if len(p.VaultAllowlist) > 16 {
		return fmt.Errorf("at most 16 vaults")
	}
	seen := make(map[[32]byte]bool)
	for _, v := range p.VaultAllowlist {
		if len(v) != 32 {
			return fmt.Errorf("vault must be 32 bytes")
		}
		key := [32]byte(v)
		if key == [32]byte{} || seen[key] {
			return fmt.Errorf("zero or duplicate vault")
		}
		seen[key] = true
	}
	return nil
}

// ValidatePauseGuardian accepts a canonical, nonzero bech32 account of 20
// bytes, or 32 bytes for derived accounts such as group policies.
func ValidatePauseGuardian(guardian string) error {
	addr, err := sdk.AccAddressFromBech32(guardian)
	if err != nil || (len(addr) != 20 && len(addr) != 32) || addr.String() != guardian || bytes.Equal(addr, make([]byte, len(addr))) {
		return fmt.Errorf("pause guardian must be a canonical 20- or 32-byte account address")
	}
	return nil
}
