package client

import (
	"context"
	"errors"
	"fmt"
	"sync"

	cmtrpc "github.com/cometbft/cometbft/rpc/core"
	cmjclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
	cmtrpcsrv "github.com/cometbft/cometbft/rpc/jsonrpc/server"
	cmtrpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cltypes "pkg.akt.dev/go/node/client/types"
	"pkg.akt.dev/go/node/client/v1beta3"
	"pkg.akt.dev/go/node/client/v1beta4"
)

var (
	ErrDetectClientVersion  = errors.New("chain-sdk: unable detect client version")
	ErrUnknownClientVersion = errors.New("chain-sdk: unknown client version")
)

const (
	VersionV1beta3 = "v1beta3"
	VersionV1beta4 = "v1beta4"
)

func init() {
	// register akash api routes
	cmtrpc.Routes["akash"] = cmtrpcsrv.NewRPCFunc(RPCAkash, "")
}

// SetupFn defines a function that takes a parameter, ideally a Client or QueryClient.
// These functions must validate the client and make it accessible.
type SetupFn func(interface{}) error

// clientPreferenceOrder defines the version negotiation preference, newest first.
// When both client and server support multiple versions, the client picks the first
// version from this list that the server also supports.
var clientPreferenceOrder = []string{VersionV1beta4, VersionV1beta3}

// negotiateVersion picks the best API version given a server's Akash discovery response.
// It checks SupportedVersions first (new-style multi-version discovery), then falls back
// to ClientInfo.ApiVersion (legacy single-version discovery from old nodes).
func negotiateVersion(result *Akash) string {
	if len(result.SupportedVersions) > 0 {
		serverVersions := make(map[string]struct{}, len(result.SupportedVersions))
		for _, v := range result.SupportedVersions {
			serverVersions[v.ApiVersion] = struct{}{}
		}
		for _, cv := range clientPreferenceOrder {
			if _, ok := serverVersions[cv]; ok {
				return cv
			}
		}
		return "" // no compatible version
	}

	// Legacy: old node returns only ClientInfo.ApiVersion
	return result.ClientInfo.ApiVersion
}

func queryClientInfo(ctx context.Context, cctx sdkclient.Context) (*Akash, error) {
	result := new(Akash)
	var err error

	if !cctx.Offline {
		if cctx.Client != nil {
			switch rpc := cctx.Client.(type) {
			case RPCClient:
				if result, err = rpc.Akash(ctx); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("unsupported RPC client [%T]", rpc)
			}
		} else {
			rpc, err := cmjclient.New(NormalizeEndpoint(cctx.NodeURI))
			if err != nil {
				return nil, err
			}

			params := make(map[string]interface{})
			if _, err = rpc.Call(ctx, "akash", params, result); err != nil {
				return nil, err
			}
		}

		if len(result.SupportedVersions) == 0 && result.ClientInfo.ApiVersion == "" {
			return nil, ErrDetectClientVersion
		}
	} else {
		result.ClientInfo = ClientInfo{ApiVersion: VersionV1beta3}
	}

	return result, nil
}

func newClientForVersion(ctx context.Context, cctx sdkclient.Context, version string, opts ...cltypes.ClientOption) (interface{}, error) {
	switch version {
	case VersionV1beta4:
		return v1beta4.NewClient(ctx, cctx, opts...)
	case VersionV1beta3:
		return v1beta3.NewClient(ctx, cctx, opts...)
	default:
		return nil, ErrUnknownClientVersion
	}
}

func newLightClientForVersion(cctx sdkclient.Context, version string) (interface{}, error) {
	switch version {
	case VersionV1beta4:
		return v1beta4.NewLightClient(cctx)
	case VersionV1beta3:
		return v1beta3.NewLightClient(cctx)
	default:
		return nil, ErrUnknownClientVersion
	}
}

func newQueryClientForVersion(cctx sdkclient.Context, version string) (interface{}, error) {
	switch version {
	case VersionV1beta4:
		return v1beta4.NewQueryClient(cctx), nil
	case VersionV1beta3:
		return v1beta3.NewQueryClient(cctx), nil
	default:
		return nil, ErrUnknownClientVersion
	}
}

// DiscoverClient queries an RPC node to get the version of the client and executes a SetupFn function
// passing a new versioned Client instance as parameter.
// If any error occurs when calling the RPC node or the Cosmos SDK client Context is set to offline the default value of
// DefaultClientAPIVersion will be used.
// An error is returned if client discovery is not successful.
func DiscoverClient(ctx context.Context, cctx sdkclient.Context, setup SetupFn, opts ...cltypes.ClientOption) error {
	result, err := queryClientInfo(ctx, cctx)
	if err != nil {
		return err
	}

	version := negotiateVersion(result)
	cl, err := newClientForVersion(ctx, cctx, version, opts...)
	if err != nil {
		return err
	}

	return setup(cl)
}

func DiscoverLightClient(ctx context.Context, cctx sdkclient.Context, setup SetupFn) error {
	result, err := queryClientInfo(ctx, cctx)
	if err != nil {
		return err
	}

	version := negotiateVersion(result)
	cl, err := newLightClientForVersion(cctx, version)
	if err != nil {
		return err
	}

	return setup(cl)
}

// DiscoverQueryClient queries an RPC node to get the version of the client and executes a SetupFn function
// passing a new versioned QueryClient instance as parameter.
// If any error occurs when calling the RPC node or the Cosmos SDK client Context is set to offline the default value of
// DefaultClientAPIVersion will be used.
// An error is returned if client discovery is not successful.
func DiscoverQueryClient(ctx context.Context, cctx sdkclient.Context, setup SetupFn) error {
	result, err := queryClientInfo(ctx, cctx)
	if err != nil {
		return err
	}

	version := negotiateVersion(result)
	cl, err := newQueryClientForVersion(cctx, version)
	if err != nil {
		return err
	}

	return setup(cl)
}

var (
	// rLock guards concurrent access to defaultRegistry.
	rLock sync.RWMutex
	// defaultRegistry is the global version registry used by the CometBFT
	// JSON-RPC handler and the gRPC Discovery service.
	defaultRegistry = DefaultRegistry()
)

// SetRegistry replaces the global version registry used by RPCAkash and
// GetRegistry. It must be called during node startup before the RPC server
// begins accepting requests. Passing nil panics.
func SetRegistry(r *VersionRegistry) {
	if r == nil {
		panic("client: SetRegistry called with nil registry")
	}
	rLock.Lock()
	defaultRegistry = r
	rLock.Unlock()
}

// GetRegistry returns the current global version registry.
func GetRegistry() *VersionRegistry {
	rLock.RLock()
	r := defaultRegistry
	rLock.RUnlock()
	return r
}

func RPCAkash(_ *cmtrpctypes.Context) (*Akash, error) {
	rLock.RLock()
	r := defaultRegistry
	rLock.RUnlock()
	return r.ToAkash(), nil
}
