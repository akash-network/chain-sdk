package v1

import (
	"context"
	"reflect"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	dvbeta "pkg.akt.dev/go/node/deployment/v1beta4"
	eid "pkg.akt.dev/go/node/escrow/id/v1"
	"pkg.akt.dev/go/node/escrow/module"
	mvbeta "pkg.akt.dev/go/node/market/v1beta5"
	"pkg.akt.dev/go/sdkutil"
	"pkg.akt.dev/go/util/ctxlog"
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
func NewDepositAuthorization(scopes DepositAuthorizationScopes, spendLimits sdk.Coins) *DepositAuthorization {
	return &DepositAuthorization{
		Scopes:      scopes,
		SpendLimit:  sdk.NewCoin(sdkutil.DenomUakt, sdkmath.ZeroInt()),
		SpendLimits: spendLimits,
	}
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
		amount = mt.Deposit.Amount
	case *mvbeta.MsgCreateBid:
		scope = DepositScopeBid
		amount = mt.Deposit.Amount
	default:
		return authz.AcceptResponse{}, errorsmod.Wrapf(sdkerrors.ErrInvalidType, "\"%s\" is unsupported authorization msg for deposit", reflect.TypeOf(mt).String())
	}

	if !m.Scopes.Has(scope) {
		return authz.AcceptResponse{}, module.ErrUnauthorizedDepositScope
	}

	// migrate spend limit to spend_limits
	if !m.SpendLimit.IsNil() && m.SpendLimit.Amount.GT(sdkmath.ZeroInt()) {
		m.SpendLimits = m.SpendLimits.Add(m.SpendLimit)
	}

	var coin sdk.Coin

	for i := range m.SpendLimits {
		if m.SpendLimits[i].Denom == amount.Denom {
			coin = m.SpendLimits[i]
			break
		}
	}

	if coin.IsNil() {
		return authz.AcceptResponse{Accept: false}, nil
	}

	allowedSpend := amount

	if coin.IsLT(allowedSpend) {
		if partial {
			allowedSpend = coin
		} else {
			return authz.AcceptResponse{}, sdkerrors.ErrInsufficientFunds
		}
	}

	limitsLeft, isNegative := m.SpendLimits.SafeSub(allowedSpend)
	if isNegative {
		return authz.AcceptResponse{}, sdkerrors.ErrInsufficientFunds.Wrapf("requested amount is more than spend limit")
	}

	return authz.AcceptResponse{
			Accept: true,
			Delete: limitsLeft.IsZero(),
			Updated: &DepositAuthorization{
				Scopes:      m.Scopes,
				SpendLimit:  sdk.NewCoin(sdkutil.DenomUakt, sdkmath.ZeroInt()),
				SpendLimits: limitsLeft,
			},
		},
		nil
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

	if !m.SpendLimit.IsNil() {
		ctxlog.Logger(context.Background()).Warn("deposit authorization uses deprecated spend_limit, use spend_limits")

		if !m.SpendLimit.IsValid() {
			return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "spend limit is not valid")
		}
	}

	if !m.SpendLimits.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "spend limits are not valid")
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
