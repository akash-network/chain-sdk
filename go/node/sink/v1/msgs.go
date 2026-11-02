package v1

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.HasValidateBasic = &MsgSink{}
	_ sdk.HasValidateBasic = &MsgGuardianPause{}
)

// ValidateBasic performs the stateless MsgSink checks before fees are charged.
// The keeper repeats them and adds the parameter-dependent vault allowlist,
// amount limits and signer account checks.
func (m *MsgSink) ValidateBasic() error {
	if m.DestType != DestTypeSolana {
		return ErrInvalidDestination.Wrapf("destination type must be %d (Solana)", DestTypeSolana)
	}
	if _, err := DecodeDestination(m.Dest); err != nil {
		return err
	}
	if m.Amount.Denom != Denom {
		return sdkerrors.ErrInvalidCoins.Wrapf("expected %s, got %s", Denom, m.Amount.Denom)
	}
	if m.Amount.Amount.IsNil() || !m.Amount.Amount.IsPositive() || !m.Amount.Amount.IsUint64() {
		return sdkerrors.ErrInvalidCoins.Wrap("amount must be a positive 64-bit value")
	}
	if _, err := ParseAccount(m.Sender, AccountLength); err != nil {
		return errorsmod.Wrap(err, "sender")
	}
	return nil
}

// ValidateBasic checks the guardian address and that at least one switch is
// requested. The keeper checks that the signer is params.pause_guardian.
func (m *MsgGuardianPause) ValidateBasic() error {
	if err := ValidatePauseGuardian(m.Guardian); err != nil {
		return err
	}
	if !m.Deposits && !m.Refunds {
		return ErrInvalidGuardianPause
	}
	return nil
}
