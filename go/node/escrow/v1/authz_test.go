package v1_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/stretchr/testify/require"

	dvbeta "pkg.akt.dev/go/node/deployment/v1beta4"
	v1 "pkg.akt.dev/go/node/escrow/v1"
	"pkg.akt.dev/go/node/escrow/module"
	mvbeta "pkg.akt.dev/go/node/market/v1beta5"
	deposit "pkg.akt.dev/go/node/types/deposit/v1"
	"pkg.akt.dev/go/sdkutil"
	"pkg.akt.dev/go/testutil"
)

func TestDepositAuthorization_TryAccept(t *testing.T) {
	denom := sdkutil.DenomUact
	limit := sdk.NewInt64Coin(denom, 1000)
	dep := func(amount sdk.Coin) deposit.Deposit {
		return deposit.Deposit{
			Amount:  amount,
			Sources: deposit.Sources{deposit.SourceGrant},
		}
	}

	tests := []struct {
		name       string
		scopes     v1.DepositAuthorizationScopes
		limits     sdk.Coins
		msg        sdk.Msg
		partial    bool
		expAccept  bool
		expDelete  bool
		expErr     error
		expErrMsg  string
		expUpdated bool
	}{
		{
			name:      "nil_msg",
			scopes:    v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits:    sdk.Coins{limit},
			msg:       nil,
			expErr:    sdkerrors.ErrInvalidType,
			expErrMsg: "msg cannot be nil",
		},
		{
			name:   "unsupported_msg_type",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg:    &authz.MsgGrant{},
			expErr: sdkerrors.ErrInvalidType,
		},
		{
			name:   "account_deposit_deployment_scope",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 500)),
			),
			expAccept:  true,
			expUpdated: true,
		},
		{
			name:   "account_deposit_bid_scope",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeBid},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.BidID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 500)),
			),
			expAccept:  true,
			expUpdated: true,
		},
		{
			name:   "account_deposit_invalid_scope",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: func() sdk.Msg {
				msg := v1.NewMsgAccountDeposit(
					testutil.AccAddress(t).String(),
					testutil.DeploymentID(t).ToEscrowAccountID(),
					dep(sdk.NewInt64Coin(denom, 500)),
				)
				msg.ID.Scope = 99
				return msg
			}(),
			expErr: module.ErrUnauthorizedDepositScope,
		},
		{
			name:   "account_deposit_scope_not_authorized",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeBid},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 500)),
			),
			expErr: module.ErrUnauthorizedDepositScope,
		},
		{
			name:   "create_deployment_msg",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: &dvbeta.MsgCreateDeployment{
				Deposit: dep(sdk.NewInt64Coin(denom, 300)),
			},
			expAccept:  true,
			expUpdated: true,
		},
		{
			name:   "create_deployment_scope_not_authorized",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeBid},
			limits: sdk.Coins{limit},
			msg: &dvbeta.MsgCreateDeployment{
				Deposit: dep(sdk.NewInt64Coin(denom, 300)),
			},
			expErr: module.ErrUnauthorizedDepositScope,
		},
		{
			name:   "create_bid_msg",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeBid},
			limits: sdk.Coins{limit},
			msg: &mvbeta.MsgCreateBid{
				Deposit: dep(sdk.NewInt64Coin(denom, 300)),
			},
			expAccept:  true,
			expUpdated: true,
		},
		{
			name:   "create_bid_scope_not_authorized",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: &mvbeta.MsgCreateBid{
				Deposit: dep(sdk.NewInt64Coin(denom, 300)),
			},
			expErr: module.ErrUnauthorizedDepositScope,
		},
		{
			name:   "denom_not_in_limits",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(sdkutil.DenomUakt, 500)),
			),
			expAccept: false,
		},
		{
			name:   "exact_limit_deletes_authorization",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(limit),
			),
			expAccept:  true,
			expDelete:  true,
			expUpdated: true,
		},
		{
			name:   "exceeds_limit_no_partial",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 1001)),
			),
			partial: false,
			expErr:  sdkerrors.ErrInsufficientFunds,
		},
		{
			name:   "exceeds_limit_with_partial",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{limit},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 1500)),
			),
			partial:    true,
			expAccept:  true,
			expDelete:  true,
			expUpdated: true,
		},
		{
			name:   "legacy_spend_limit_migration",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			limits: sdk.Coins{sdk.NewInt64Coin(denom, 500)},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(denom, 800)),
			),
		},
		{
			name:   "multi_denom_limits",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment, v1.DepositScopeBid},
			limits: sdk.Coins{sdk.NewInt64Coin(denom, 500), sdk.NewInt64Coin(sdkutil.DenomUakt, 300)},
			msg: v1.NewMsgAccountDeposit(
				testutil.AccAddress(t).String(),
				testutil.DeploymentID(t).ToEscrowAccountID(),
				dep(sdk.NewInt64Coin(sdkutil.DenomUakt, 200)),
			),
			expAccept:  true,
			expUpdated: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var da *v1.DepositAuthorization

			if tt.name == "legacy_spend_limit_migration" {
				da = &v1.DepositAuthorization{
					Scopes:      tt.scopes,
					SpendLimit:  sdk.NewInt64Coin(denom, 500),
					SpendLimits: tt.limits,
				}
				resp, err := da.TryAccept(sdk.Context{}, tt.msg, tt.partial)
				require.NoError(t, err)
				require.True(t, resp.Accept)
				return
			}

			da = v1.NewDepositAuthorization(tt.scopes, tt.limits)

			resp, err := da.TryAccept(sdk.Context{}, tt.msg, tt.partial)

			if tt.expErr != nil {
				require.ErrorIs(t, err, tt.expErr)
				if tt.expErrMsg != "" {
					require.Contains(t, err.Error(), tt.expErrMsg)
				}
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expAccept, resp.Accept)

			if !tt.expAccept {
				return
			}

			require.Equal(t, tt.expDelete, resp.Delete)

			if tt.expUpdated {
				updated, ok := resp.Updated.(*v1.DepositAuthorization)
				require.True(t, ok)
				require.Equal(t, tt.scopes, updated.Scopes)
				require.True(t, updated.SpendLimit.Amount.Equal(sdkmath.ZeroInt()))
			}
		})
	}
}

func TestDepositAuthorization_ValidateBasic(t *testing.T) {
	denom := sdkutil.DenomUact

	tests := []struct {
		name      string
		auth      *v1.DepositAuthorization
		expErr    error
		expErrMsg string
	}{
		{
			name: "valid_single_scope",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
		},
		{
			name: "valid_multiple_scopes",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{v1.DepositScopeDeployment, v1.DepositScopeBid},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
		},
		{
			name: "empty_scopes",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
			expErr:    module.ErrInvalidAuthzScope,
			expErrMsg: "empty scope",
		},
		{
			name: "invalid_scope_value",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{v1.DepositScopeInvalid},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
			expErr:    module.ErrInvalidAuthzScope,
			expErrMsg: "invalid scope",
		},
		{
			name: "unknown_scope_value",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{99},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
			expErr:    module.ErrInvalidAuthzScope,
			expErrMsg: "invalid scope",
		},
		{
			name: "duplicate_scopes",
			auth: v1.NewDepositAuthorization(
				v1.DepositAuthorizationScopes{v1.DepositScopeDeployment, v1.DepositScopeDeployment},
				sdk.Coins{sdk.NewInt64Coin(denom, 100)},
			),
			expErr:    module.ErrInvalidAuthzScope,
			expErrMsg: "duplicate scope",
		},
		{
			name: "invalid_spend_limit",
			auth: &v1.DepositAuthorization{
				Scopes:     v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
				SpendLimit: sdk.Coin{Denom: denom, Amount: sdkmath.NewInt(-1)},
			},
			expErr:    sdkerrors.ErrInvalidCoins,
			expErrMsg: "spend limit is not valid",
		},
		{
			name: "invalid_spend_limits",
			auth: &v1.DepositAuthorization{
				Scopes:      v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
				SpendLimit:  sdk.NewCoin(sdkutil.DenomUakt, sdkmath.ZeroInt()),
				SpendLimits: sdk.Coins{sdk.Coin{Denom: denom, Amount: sdkmath.NewInt(-1)}},
			},
			expErr:    sdkerrors.ErrInvalidCoins,
			expErrMsg: "spend limits are not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.auth.ValidateBasic()

			if tt.expErr != nil {
				require.ErrorIs(t, err, tt.expErr)
				if tt.expErrMsg != "" {
					require.Contains(t, err.Error(), tt.expErrMsg)
				}
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestDepositAuthorizationScopes_Has(t *testing.T) {
	tests := []struct {
		name   string
		scopes v1.DepositAuthorizationScopes
		val    v1.DepositAuthorization_Scope
		exp    bool
	}{
		{
			name:   "has_deployment",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment, v1.DepositScopeBid},
			val:    v1.DepositScopeDeployment,
			exp:    true,
		},
		{
			name:   "has_bid",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment, v1.DepositScopeBid},
			val:    v1.DepositScopeBid,
			exp:    true,
		},
		{
			name:   "missing_scope",
			scopes: v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
			val:    v1.DepositScopeBid,
			exp:    false,
		},
		{
			name:   "empty_scopes",
			scopes: v1.DepositAuthorizationScopes{},
			val:    v1.DepositScopeDeployment,
			exp:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, tt.scopes.Has(tt.val))
		})
	}
}

func TestDepositAuthorization_MsgTypeURL(t *testing.T) {
	da := v1.NewDepositAuthorization(
		v1.DepositAuthorizationScopes{v1.DepositScopeDeployment},
		sdk.Coins{sdk.NewInt64Coin(sdkutil.DenomUact, 100)},
	)
	url := da.MsgTypeURL()
	require.NotEmpty(t, url)
	require.Contains(t, url, "MsgAccountDeposit")
}
