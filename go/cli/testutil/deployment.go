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
