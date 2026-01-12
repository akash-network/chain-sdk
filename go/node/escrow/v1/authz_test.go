package v1_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "pkg.akt.dev/go/node/escrow/v1"
	deposit "pkg.akt.dev/go/node/types/deposit/v1"
	testutil "pkg.akt.dev/go/testutil/v1beta4"
)

func TestDepositAuthorizationAccept(t *testing.T) {
	limit := sdk.NewInt64Coin(testutil.CoinDenom, 333)
	dda := v1.NewDepositAuthorization(v1.DepositAuthorizationScopes{v1.DepositScopeDeployment}, limit)

	sctx := sdk.Context{}

	// Send the wrong type of message, expect an error
	var msg sdk.Msg
	response, err := dda.Accept(sctx, msg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid type")
	require.Zero(t, response)

	// Try to deposit too much coin, expect an error
	spendReq := limit.Add(sdk.NewInt64Coin(testutil.CoinDenom, 1))

	did := testutil.DeploymentID(t)

	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  spendReq,
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.ErrorIs(t, err, sdkerrors.ErrInsufficientFunds)
	require.Zero(t, response)

	did = testutil.DeploymentID(t)
	// Deposit 1 less than the limit, expect an updated deposit
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  limit.Sub(sdk.NewInt64Coin(testutil.CoinDenom, 1)),
		Sources: deposit.Sources{deposit.SourceGrant},
	})
	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.False(t, response.Delete)

	ok := false
	dda, ok = response.Updated.(*v1.DepositAuthorization)
	require.True(t, ok)

	did = testutil.DeploymentID(t)
	// Deposit the limit (now 1), expect that it is not to be deleted
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin(testutil.CoinDenom, 1),
		Sources: deposit.Sources{deposit.SourceGrant},
	})
	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.True(t, response.Delete)
}

func TestDepositAuthorizationAcceptMultiDenom(t *testing.T) {
	// Create multi-denom authorization
	limits := sdk.NewCoins(
		sdk.NewInt64Coin("uakt", 1000),
		sdk.NewInt64Coin("uact", 2000),
		sdk.NewInt64Coin("usdc", 500),
	)
	dda := v1.NewDepositAuthorizationMultiDenom(v1.DepositAuthorizationScopes{v1.DepositScopeDeployment}, limits)

	sctx := sdk.Context{}

	// Test spending from uakt limit
	did := testutil.DeploymentID(t)
	msg := v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin("uakt", 100),
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err := dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.False(t, response.Delete)

	dda, ok := response.Updated.(*v1.DepositAuthorization)
	require.True(t, ok)
	require.Equal(t, math.NewInt(900), dda.SpendLimits.AmountOf("uakt"))
	require.Equal(t, math.NewInt(2000), dda.SpendLimits.AmountOf("uact"))
	require.Equal(t, math.NewInt(500), dda.SpendLimits.AmountOf("usdc"))

	// Test spending from uact limit
	did = testutil.DeploymentID(t)
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin("uact", 500),
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.False(t, response.Delete)

	dda, ok = response.Updated.(*v1.DepositAuthorization)
	require.True(t, ok)
	require.Equal(t, math.NewInt(900), dda.SpendLimits.AmountOf("uakt"))
	require.Equal(t, math.NewInt(1500), dda.SpendLimits.AmountOf("uact"))
	require.Equal(t, math.NewInt(500), dda.SpendLimits.AmountOf("usdc"))

	// Test spending all of usdc limit - should not delete auth as other denoms remain
	did = testutil.DeploymentID(t)
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin("usdc", 500),
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.False(t, response.Delete) // Other denoms still have balance

	dda, ok = response.Updated.(*v1.DepositAuthorization)
	require.True(t, ok)
	require.Equal(t, math.NewInt(900), dda.SpendLimits.AmountOf("uakt"))
	require.Equal(t, math.NewInt(1500), dda.SpendLimits.AmountOf("uact"))
	require.Equal(t, math.NewInt(0), dda.SpendLimits.AmountOf("usdc"))

	// Test exceeding limit for a denom
	did = testutil.DeploymentID(t)
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin("uakt", 1000), // More than remaining 900
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.ErrorIs(t, err, sdkerrors.ErrInsufficientFunds)

	// Test using unauthorized denom
	did = testutil.DeploymentID(t)
	msg = v1.NewMsgAccountDeposit(did.Owner, did.ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin("atom", 100),
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.False(t, response.Accept) // Should not accept unauthorized denom
}

func TestDepositAuthorizationGetSpendLimits(t *testing.T) {
	// Test GetSpendLimits with single denom (backward compatibility)
	limit := sdk.NewInt64Coin("uakt", 1000)
	authSingle := v1.NewDepositAuthorization(v1.DepositAuthorizationScopes{v1.DepositScopeDeployment}, limit)

	limits := authSingle.GetSpendLimits()
	require.Len(t, limits, 1)
	require.True(t, limits.Equal(sdk.NewCoins(limit)))

	// Test GetSpendLimits with multi-denom
	multiLimits := sdk.NewCoins(
		sdk.NewInt64Coin("uakt", 1000),
		sdk.NewInt64Coin("uact", 2000),
	)
	authMulti := v1.NewDepositAuthorizationMultiDenom(v1.DepositAuthorizationScopes{v1.DepositScopeDeployment}, multiLimits)

	limits = authMulti.GetSpendLimits()
	require.Len(t, limits, 2)
	require.True(t, limits.Equal(multiLimits))
}
