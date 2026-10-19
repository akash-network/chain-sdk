package v1

import (
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
		return sdkerrors.ErrInvalidRequest.Wrap("sink destination type must be 1 (Solana)")
	}
	if _, err := DecodeDestination(m.Dest); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}
	if m.Amount.Amount.IsNil() || m.Amount.Denom != Denom || !m.Amount.Amount.IsPositive() || !m.Amount.Amount.IsUint64() {
		return sdkerrors.ErrInvalidCoins.Wrap("sink accepts positive uint64 uakt only")
	}
	if sender, err := sdk.AccAddressFromBech32(m.Sender); err != nil || len(sender) != 20 {
		return sdkerrors.ErrInvalidAddress.Wrap("invalid sink sender")
	}
	return nil
}

// ValidateBasic checks the guardian address and that at least one switch is
// requested. The keeper checks that the signer is params.pause_guardian.
func (m *MsgGuardianPause) ValidateBasic() error {
	if err := ValidatePauseGuardian(m.Guardian); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	if !m.Deposits && !m.Refunds {
		return sdkerrors.ErrInvalidRequest.Wrap("guardian pause must disable deposits, refunds or both")
	}
	return nil
}
