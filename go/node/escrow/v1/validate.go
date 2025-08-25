package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (obj *AccountID) ValidateBasic() error {
	if len(obj.Scope) == 0 {
		return ErrInvalidAccountID.Wrap("empty scope")
	}

	if len(obj.XID) == 0 {
		return ErrInvalidAccountID.Wrap("empty XID")
	}

	return nil
}

func (m *Account) ValidateBasic() error {
	if err := m.ID.ValidateBasic(); err != nil {
		return ErrInvalidAccount.Wrap(err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return ErrInvalidAccount.Wrap(err.Error())
	}
	if m.State == StateInvalid {
		return ErrInvalidAccount.Wrap("invalid state")
	}

	for _, deposit := range m.Deposits {
		if _, err := sdk.AccAddressFromBech32(deposit.Depositor); err != nil {
			return ErrInvalidAccount.Wrapf("invalid depositor")
		}
	}

	return nil
}

func (obj *FractionalPayment) ValidateBasic() error {
	if err := obj.AccountID.ValidateBasic(); err != nil {
		return ErrInvalidPayment.Wrap(err.Error())
	}
	if len(obj.PaymentID) == 0 {
		return ErrInvalidPayment.Wrap("empty payment id")
	}
	if obj.Rate.IsZero() {
		return ErrInvalidPayment.Wrap("payment rate zero")
	}
	if obj.State == StateInvalid {
		return ErrInvalidPayment.Wrap("invalid state")
	}
	return nil
}
