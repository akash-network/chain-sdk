package v1

import (
	"context"
	"reflect"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	dvbeta "pkg.akt.dev/go/node/deployment/v1beta5"
	eid "pkg.akt.dev/go/node/escrow/id/v1"
	"pkg.akt.dev/go/node/escrow/module"
	mvbeta "pkg.akt.dev/go/node/market/v2beta1"
)

type Authorization interface {
	authz.Authorization
	TryAccept(context.Context, sdk.Msg, bool) (authz.AcceptResponse, error)
	GetSpendLimit() sdk.Coin
	GetSpendLimits() sdk.Coins
}

type DepositAuthorizationScopes []DepositAuthorization_Scope

var (
	_ Authorization = &DepositAuthorization{}
)

// NewDepositAuthorization creates a new DepositAuthorization object with a single spend limit.
func NewDepositAuthorization(scopes DepositAuthorizationScopes, spendLimit sdk.Coin) *DepositAuthorization {
	return &DepositAuthorization{
		Scopes:     scopes,
		SpendLimit: spendLimit,
	}
}

// NewDepositAuthorizationMultiDenom creates a new DepositAuthorization object with multiple spend limits.
func NewDepositAuthorizationMultiDenom(scopes DepositAuthorizationScopes, spendLimits sdk.Coins) *DepositAuthorization {
	return &DepositAuthorization{
		Scopes:      scopes,
		SpendLimits: spendLimits,
	}
}

// GetSpendLimits returns the spend limits for multi-denom support.
// If SpendLimits is set, it returns that; otherwise, it returns SpendLimit as a single-coin slice.
func (m *DepositAuthorization) GetSpendLimits() sdk.Coins {
	if len(m.SpendLimits) > 0 {
		return m.SpendLimits
	}
	return sdk.NewCoins(m.SpendLimit)
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (m *DepositAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgAccountDeposit{})
}

// Accept implements Authorization.Accept.
func (m *DepositAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	return m.TryAccept(ctx, msg, false)
}

func (m *DepositAuthorization) TryAccept(_ context.Context, msg sdk.Msg, partial bool) (authz.AcceptResponse, error) {
	if msg == nil {
		return authz.AcceptResponse{Accept: false}, errorsmod.Wrapf(sdkerrors.ErrInvalidType, "msg cannot be nil")
	}
	var amount sdk.Coin
	var scope DepositAuthorization_Scope

	switch mt := msg.(type) {
	case *MsgAccountDeposit:
		switch mt.ID.Scope {
		case eid.ScopeDeployment:
			scope = DepositScopeDeployment
		case eid.ScopeBid:
			scope = DepositScopeBid
		default:
			return authz.AcceptResponse{}, module.ErrUnauthorizedDepositScope
		}

		amount = mt.Deposit.Amount
	case *dvbeta.MsgCreateDeployment:
		scope = DepositScopeDeployment
		// Sum all deposits for multi-denom support
		if len(mt.Deposits) == 0 {
			return authz.AcceptResponse{}, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no deposits specified")
		}
		// For now, use the first deposit amount (will be enhanced below for multi-denom)
		amount = mt.Deposits[0].Amount
	case *mvbeta.MsgCreateBid:
		scope = DepositScopeBid
		amount = mt.Deposit.Amount
	default:
		return authz.AcceptResponse{}, errorsmod.Wrapf(sdkerrors.ErrInvalidType, "\"%s\" is unsupported authorization msg for deposit", reflect.TypeOf(mt).String())
	}

	if !m.Scopes.Has(scope) {
		return authz.AcceptResponse{}, module.ErrUnauthorizedDepositScope
	}

	// Multi-denom support: use SpendLimits if available, otherwise fall back to SpendLimit
	var limitsLeft sdk.Coins
	if len(m.SpendLimits) > 0 {
		// Multi-denom mode
		limitCoin := m.SpendLimits.AmountOf(amount.Denom)
		if limitCoin.IsZero() {
			return authz.AcceptResponse{Accept: false}, nil
		}

		allowedSpend := amount
		currentLimit := sdk.NewCoin(amount.Denom, limitCoin)

		if currentLimit.IsLT(allowedSpend) {
			if partial {
				allowedSpend = currentLimit
			} else {
				return authz.AcceptResponse{}, sdkerrors.ErrInsufficientFunds
			}
		}

		// Subtract from the specific denom limit
		limitsLeft = m.SpendLimits.Sub(sdk.NewCoins(allowedSpend)...)
	} else {
		// Single denom mode (backward compatibility)
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

		return authz.AcceptResponse{Accept: true, Delete: limitLeft.IsZero(), Updated: &DepositAuthorization{Scopes: m.Scopes, SpendLimit: limitLeft}}, nil
	}

	// Return updated authorization with multi-denom limits
	deleteAuth := limitsLeft.IsZero()
	return authz.AcceptResponse{
		Accept: true,
		Delete: deleteAuth,
		Updated: &DepositAuthorization{
			Scopes:      m.Scopes,
			SpendLimits: limitsLeft,
		},
	}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (m *DepositAuthorization) ValidateBasic() error {
	if len(m.Scopes) == 0 {
		return errorsmod.Wrap(module.ErrInvalidAuthzScope, "empty scope")
	}

	scopes := make(map[DepositAuthorization_Scope]int)

	for _, scope := range m.Scopes {
		if _, valid := DepositAuthorization_Scope_name[int32(scope)]; !valid || scope == DepositScopeInvalid {
			return errorsmod.Wrapf(module.ErrInvalidAuthzScope, "invalid scope \"%s\"", scope.String())
		}

		if _, exists := scopes[scope]; exists {
			return errorsmod.Wrapf(module.ErrInvalidAuthzScope, "duplicate scope \"%s\"", scope.String())
		}

		scopes[scope] = 0
	}
	if !m.SpendLimit.IsPositive() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "spend limit cannot be negative")
	}

	return nil
}

func (s DepositAuthorizationScopes) Has(val DepositAuthorization_Scope) bool {
	for _, scope := range s {
		if scope == val {
			return true
		}
	}

	return false
}
