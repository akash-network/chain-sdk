// Package v1beta4 provides the v1beta4 composite client for the Akash blockchain.
// Currently it delegates to v1beta3 as both share the same module versions.
// When module versions diverge, this package will carry its own implementations.
package v1beta4

import (
	"context"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cltypes "pkg.akt.dev/go/node/client/types"
	"pkg.akt.dev/go/node/client/v1beta3"
)

// Re-export v1beta3 interfaces so callers use the same types.
type (
	QueryClient = v1beta3.QueryClient
	TxClient    = v1beta3.TxClient
	NodeClient  = v1beta3.NodeClient
	LightClient = v1beta3.LightClient
	Client      = v1beta3.Client
)

// Re-export broadcast types that callers may need.
type (
	BroadcastOption  = v1beta3.BroadcastOption
	BroadcastOptions = v1beta3.BroadcastOptions
	ConfirmFn        = v1beta3.ConfirmFn
)

// NewClient creates a new v1beta4 Client.
// Currently delegates to v1beta3 since module versions are identical.
func NewClient(ctx context.Context, cctx sdkclient.Context, opts ...cltypes.ClientOption) (Client, error) {
	return v1beta3.NewClient(ctx, cctx, opts...)
}

// NewLightClient creates a new v1beta4 LightClient.
func NewLightClient(cctx sdkclient.Context) (LightClient, error) {
	return v1beta3.NewLightClient(cctx)
}

// NewQueryClient creates a new v1beta4 QueryClient.
func NewQueryClient(cctx sdkclient.Context) QueryClient {
	return v1beta3.NewQueryClient(cctx)
}
