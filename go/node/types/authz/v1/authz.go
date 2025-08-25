package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	dv1beta "pkg.akt.dev/go/node/deployment/v1beta4"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

type Authorization interface {
	authz.Authorization
	TryAccept(context.Context, sdk.Msg, bool) (authz.AcceptResponse, error)
}

var (
	_ Authorization = &DepositAuthorization{}
)

// NewDepositAuthorization creates a new DepositAuthorization object.
func NewDepositAuthorization(spendLimit sdk.Coin) *DepositAuthorization {
	return &DepositAuthorization{
		SpendLimit: spendLimit,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (m *DepositAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(m)
}

// Accept implements Authorization.Accept.
func (m *DepositAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	return m.TryAccept(ctx, msg, false)
}

func (m *DepositAuthorization) TryAccept(_ context.Context, msg sdk.Msg, partial bool) (authz.AcceptResponse, error) {
	var amount sdk.Coin

	switch m := msg.(type) {
	case *dv1beta.MsgDepositDeployment:
		amount = m.Deposit.Amount
	case *dv1beta.MsgCreateDeployment:
		amount = m.Deposit.Amount
	case *mv1beta5.MsgCreateBid:
		amount = m.Deposit.Amount
	default:
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("unsupported message type")
	}

	if m.SpendLimit.Denom != amount.Denom {
		return authz.AcceptResponse{Accept: false}, nil
	}

	allowedSpend := amount

	if m.SpendLimit.IsLT(allowedSpend) {
		if partial {
			allowedSpend = m.SpendLimit
		} else {
			return authz.AcceptResponse{}, sdkerrors.ErrInsufficientFunds
		}
	}

	limitLeft, err := m.SpendLimit.SafeSub(allowedSpend)
	if err != nil {
		return authz.AcceptResponse{}, err
	}

	return authz.AcceptResponse{Accept: true, Delete: limitLeft.IsZero(), Updated: &DepositAuthorization{SpendLimit: limitLeft}}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (m *DepositAuthorization) ValidateBasic() error {
	if !m.SpendLimit.IsPositive() {
		return sdkerrors.ErrInvalidCoins.Wrapf("spend limit cannot be negative")
	}

	return nil
}
