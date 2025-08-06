package testutil

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	sdktest "github.com/cosmos/cosmos-sdk/testutil"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.akt.dev/go/cli"
	// dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
)

// ExecDeploymentCreate is used for testing create deployment tx
func ExecDeploymentCreate(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentCreateCmd(), args...)
}

// ExecDeploymentUpdate is used for testing update deployment tx
func ExecDeploymentUpdate(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentUpdateCmd(), args...)
}

// ExecDeploymentClose is used for testing close deployment tx
// requires --dseq, --fees
func ExecDeploymentClose(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentCloseCmd(), args...)
}

// TxDepositDeploymentExec is used for testing deposit deployment tx
//func TxDepositDeploymentExec(ctx context.Context, cctx client.Context, deposit sdk.Coin, extraArgs ...string) (sdktest.BufferWriter, error) {
//	args := []string{
//		deposit.String(),
//	}
//
//	args = append(args, extraArgs...)
//
//	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentDepositCmd(), args...)
//}
//
// // TxCloseGroupExec is used for testing close group tx
// func TxCloseGroupExec(ctx context.Context, cctx client.Context, extraArgs ...string) (sdktest.BufferWriter, error) {
// 	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentGroupCloseCmd(), extraArgs...)
// }
//
// func TxGrantAuthorizationExec(ctx context.Context, cctx client.Context, grantee sdk.AccAddress, extraArgs ...string) (sdktest.BufferWriter, error) {
// 	dmin, _ := dv1beta4.DefaultParams().MinDepositFor("uakt")
//
// 	spendLimit := sdk.NewCoin(dmin.Denom, dmin.Amount.MulRaw(3))
// 	args := []string{
// 		grantee.String(),
// 		spendLimit.String(),
// 	}
// 	args = append(args, extraArgs...)
//
// 	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentGrantAuthorizationCmd(), args...)
// }
//
// func TxRevokeAuthorizationExec(ctx context.Context, cctx client.Context, grantee sdk.AccAddress, extraArgs ...string) (sdktest.BufferWriter, error) {
// 	args := []string{
// 		grantee.String(),
// 	}
// 	args = append(args, extraArgs...)
//
// 	return ExecTestCLICmd(ctx, cctx, cli.GetTxDeploymentRevokeAuthorizationCmd(), args...)
// }

// ExecQueryDeployments is used for testing deployments query
func ExecQueryDeployments(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetQueryDeploymentsCmd(), args...)
}

// ExecQueryDeployment is used for testing deployment query
func ExecQueryDeployment(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetQueryDeploymentCmd(), args...)
}

// ExecQueryGroup is used for testing group query
func ExecQueryGroup(ctx context.Context, cctx client.Context, args ...string) (sdktest.BufferWriter, error) {
	return ExecTestCLICmd(ctx, cctx, cli.GetQueryDeploymentGroupCmd(), args...)
}
