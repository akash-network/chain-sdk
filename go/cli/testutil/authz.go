package testutil

import (
	"context"

	"cosmossdk.io/core/address"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/testutil"

	"pkg.akt.dev/go/cli"
)

func ExecCreateGrant(ctx context.Context, cctx client.Context, ac address.Codec, args ...string) (testutil.BufferWriter, error) {
	cmd := cli.GetTxAuthzGrantAuthorizationCmd(ac)
	return ExecTestCLICmd(ctx, cctx, cmd, args...)
}

func ExecRevokeAuthz(ctx context.Context, cctx client.Context, ac address.Codec, args ...string) (testutil.BufferWriter, error) {
	cmd := cli.GetTxAuthzRevokeAuthorizationCmd(ac)
	return ExecTestCLICmd(ctx, cctx, cmd, args...)
}

func ExecAuthorization(ctx context.Context, cctx client.Context, args ...string) (testutil.BufferWriter, error) {
	cmd := cli.GetTxAuthzExecAuthorizationCmd()
	return ExecTestCLICmd(ctx, cctx, cmd, args...)
}
