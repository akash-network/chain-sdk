package v1

import (
	"context"
	"fmt"
	"reflect"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"

	v1 "pkg.akt.dev/go/node/deployment/v1"
	dv1beta "pkg.akt.dev/go/node/deployment/v1beta4"
	ev1 "pkg.akt.dev/go/node/escrow/v1"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
	deposit "pkg.akt.dev/go/node/types/deposit/v1"
)

type AuthzKeeper interface {
	DeleteGrant(ctx context.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error
	GetAuthorization(ctx context.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) (authz.Authorization, *time.Time)
	SaveGrant(ctx context.Context, grantee sdk.AccAddress, granter sdk.AccAddress, authorization authz.Authorization, expiration *time.Time) error
	GetGranteeGrantsByMsgType(ctx context.Context, grantee sdk.AccAddress, msgType string, onGrant func(context.Context, sdk.AccAddress, authz.Authorization, *time.Time) bool)
}

type BankKeeper interface {
	SpendableCoin(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
}

func AuthorizeDeposit(sctx sdk.Context, authZ AuthzKeeper, bank BankKeeper, msg sdk.Msg) ([]ev1.Deposit, error) {
	// find the DepositDeploymentAuthorization given to the owner by the depositor and check
	// acceptance

	depositors := make([]ev1.Deposit, 0, 1)

	hasDeposit, valid := msg.(deposit.HasDeposit)
	if !valid {
		return nil, fmt.Errorf("%w: message [%s] does not implement deposit.HasDeposit", v1.ErrInvalidDeposit, reflect.TypeOf(msg).String())
	}

	lMsg, valid := msg.(sdk.LegacyMsg)
	if !valid {
		return nil, fmt.Errorf("%w: message [%s] does not implement sdk.LegacyMsg", v1.ErrInvalidDeposit, reflect.TypeOf(msg).String())
	}

	signers := lMsg.GetSigners()
	if len(signers) != 1 {
		return nil, fmt.Errorf("%w: invalid signers", v1.ErrInvalidDeposit)
	}

	owner := signers[0]

	dep := hasDeposit.GetDeposit()
	denom := dep.Amount.Denom

	remainder := sdkmath.NewInt(dep.Amount.Amount.Int64())

	for _, source := range dep.Sources {
		switch source {
		case deposit.SourceBalance:
			spendableAmount := bank.SpendableCoin(sctx, owner, denom)

			if spendableAmount.Amount.IsPositive() {
				requestedSpend := sdk.NewCoin(denom, remainder)

				if spendableAmount.IsLT(requestedSpend) {
					requestedSpend = spendableAmount
				}
				depositors = append(depositors, ev1.Deposit{
					Depositor: owner.String(),
					Height:    sctx.BlockHeight(),
					Amount:    requestedSpend,
					Balance:   sdk.NewDecCoinFromCoin(requestedSpend),
				})

				remainder = remainder.Sub(requestedSpend.Amount)
			}
		case deposit.SourceGrant:
			msgTypeUrl := (&DepositAuthorization{}).MsgTypeURL()

			authZ.GetGranteeGrantsByMsgType(sctx, owner, msgTypeUrl, func(ctx context.Context, granter sdk.AccAddress, authorization authz.Authorization, expiration *time.Time) bool {
				deplAuthz, valid := authorization.(*DepositAuthorization)
				if !valid {
					return false
				}

				authorizedSpend := sdk.Coin{
					Denom:  denom,
					Amount: sdkmath.NewInt(deplAuthz.SpendLimit.Amount.Int64()),
				}

				nDeposit := deposit.Deposit{
					Amount:  sdk.NewCoin(denom, remainder),
					Sources: nil,
				}

				var authzMsg sdk.Msg
				switch m := msg.(type) {
				case *dv1beta.MsgCreateDeployment:
					authzMsg = dv1beta.NewMsgCreateDeployment(m.ID, m.Groups, m.Hash, nDeposit)
				case *dv1beta.MsgDepositDeployment:
					authzMsg = dv1beta.NewMsgDepositDeployment(m.ID, nDeposit)
				case *mv1beta5.MsgCreateBid:
					authzMsg = mv1beta5.NewMsgCreateBid(m.ID, m.Price, nDeposit, m.ResourcesOffer)
				}

				resp, err := authorization.Accept(ctx, authzMsg)
				if err != nil {
					return false
				}

				if resp.Delete {
					err = authZ.DeleteGrant(ctx, owner, granter, msgTypeUrl)
				} else if resp.Updated != nil {
					err = authZ.SaveGrant(ctx, owner, granter, resp.Updated, expiration)
				}

				if !resp.Accept {
					return false
				}

				deplAuthz = resp.Updated.(*DepositAuthorization)

				authorizedSpend = authorizedSpend.Sub(deplAuthz.SpendLimit)

				depositors = append(depositors, ev1.Deposit{
					Depositor: owner.String(),
					Height:    sctx.BlockHeight(),
					Amount:    authorizedSpend,
					Balance:   sdk.NewDecCoinFromCoin(authorizedSpend),
				})
				remainder = remainder.Sub(authorizedSpend.Amount)

				return remainder.IsZero()
			})
		}

		if remainder.IsZero() {
			break
		}
	}

	if !remainder.IsZero() {
		// following check is for sanity. if value is negative, math above went horribly wrong
		if remainder.IsNegative() {
			return nil, fmt.Errorf("%w: deposit overflow", v1.ErrInvalidDeposit)
		} else {
			return nil, fmt.Errorf("%w: insufficient balance", v1.ErrInvalidDeposit)
		}
	}

	return depositors, nil
}
