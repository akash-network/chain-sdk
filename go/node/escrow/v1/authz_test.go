package v1_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	v1 "pkg.akt.dev/go/node/escrow/v1"
	deposit "pkg.akt.dev/go/node/types/deposit/v1"
	"pkg.akt.dev/go/testutil"
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

	msg = v1.NewMsgAccountDeposit(testutil.DeploymentID(t).ToEscrowAccountID(), deposit.Deposit{
		Amount:  spendReq,
		Sources: deposit.Sources{deposit.SourceGrant},
	})

	response, err = dda.Accept(sctx, msg)
	require.ErrorIs(t, err, sdkerrors.ErrInsufficientFunds)
	require.Zero(t, response)

	// Deposit 1 less than the limit, expect an updated deposit
	msg = v1.NewMsgAccountDeposit(testutil.DeploymentID(t).ToEscrowAccountID(), deposit.Deposit{
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

	// Deposit the limit (now 1), expect that it is not to be deleted
	msg = v1.NewMsgAccountDeposit(testutil.DeploymentID(t).ToEscrowAccountID(), deposit.Deposit{
		Amount:  sdk.NewInt64Coin(testutil.CoinDenom, 1),
		Sources: deposit.Sources{deposit.SourceGrant},
	})
	response, err = dda.Accept(sctx, msg)
	require.NoError(t, err)
	require.True(t, response.Accept)
	require.True(t, response.Delete)
}
