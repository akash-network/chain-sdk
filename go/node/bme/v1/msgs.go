package v1

import (
	"reflect"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"pkg.akt.dev/go/sdkutil"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgFundVault{}
	_ sdk.Msg = &MsgBurnMint{}
	_ sdk.Msg = &MsgMintACT{}
	_ sdk.Msg = &MsgBurnACT{}
)

var (
	msgTypeUpdateParams = ""
	msgTypeBurnMint     = ""
	msgTypeMintACT      = ""
	msgTypeBurnACT      = ""
	msgTypeFundVault    = ""
)

func init() {
	msgTypeUpdateParams = reflect.TypeOf(&MsgUpdateParams{}).Elem().Name()
	msgTypeBurnMint = reflect.TypeOf(&MsgBurnMint{}).Elem().Name()
	msgTypeMintACT = reflect.TypeOf(&MsgMintACT{}).Elem().Name()
	msgTypeBurnACT = reflect.TypeOf(&MsgBurnACT{}).Elem().Name()
	msgTypeFundVault = reflect.TypeOf(&MsgFundVault{}).Elem().Name()
}

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized.Wrapf("invalid authority address: %s", err)
	}
	return msg.Params.Validate()
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{authority}
}

// Type implements the sdk.Msg interface
func (msg *MsgUpdateParams) Type() string { return msgTypeUpdateParams }

// Type implements the sdk.Msg interface
func (msg *MsgBurnMint) Type() string { return msgTypeBurnMint }

// Type implements the sdk.Msg interface
func (msg *MsgMintACT) Type() string { return msgTypeMintACT }

// Type implements the sdk.Msg interface
func (msg *MsgBurnACT) Type() string { return msgTypeBurnACT }

// Type implements the sdk.Msg interface
func (msg *MsgFundVault) Type() string { return msgTypeFundVault }

func (msg MsgFundVault) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrUnauthorized.Wrapf("invalid authority address: %s", err)
	}
	if !msg.Amount.IsValid() || !msg.Amount.IsPositive() {
		return ErrInvalidAmount.Wrap("amount must be positive")
	}
	if msg.Amount.Denom != sdkutil.DenomUakt {
		return ErrInvalidDenom.Wrapf("expected uakt, got %s", msg.Amount.Denom)
	}

	return nil
}

func (msg MsgFundVault) GetSigners() []sdk.AccAddress {
	authority, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{authority}
}

func (msg MsgBurnMint) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{owner}
}

func (msg MsgBurnMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address: %s", err)
	}

	if msg.Owner != msg.To {
		return errors.Wrapf(ErrUnauthorized, "owner and to addresses must match")
	}

	err = msg.CoinsToBurn.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins: %s", err)
	}

	if msg.CoinsToBurn.Equal(sdk.NewInt64Coin(msg.CoinsToBurn.Denom, 0)) {
		return errors.Wrapf(ErrInvalidAmount, "amount must not be 0")
	}

	if err = sdk.ValidateDenom(msg.DenomToMint); err != nil {
		return errors.Wrapf(ErrInvalidDenom, "invalid denom to mint: %s", err)
	}

	return nil
}

func (msg MsgMintACT) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{owner}
}

func (msg MsgMintACT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address: %s", err)
	}

	if msg.Owner != msg.To {
		return errors.Wrapf(ErrUnauthorized, "owner and to addresses must match")
	}

	err = msg.CoinsToBurn.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins: %s", err)
	}

	if msg.CoinsToBurn.Equal(sdk.NewInt64Coin(msg.CoinsToBurn.Denom, 0)) {
		return errors.Wrapf(ErrInvalidAmount, "amount must not be 0")
	}

	return nil
}

func (msg MsgBurnACT) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{owner}
}

func (msg MsgBurnACT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %s", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address: %s", err)
	}

	if msg.Owner != msg.To {
		return errors.Wrapf(ErrUnauthorized, "owner and to addresses must match")
	}

	err = msg.CoinsToBurn.Validate()
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins: %s", err)
	}

	if msg.CoinsToBurn.Equal(sdk.NewInt64Coin(msg.CoinsToBurn.Denom, 0)) {
		return errors.Wrapf(ErrInvalidAmount, "amount must not be 0")
	}

	return nil
}
