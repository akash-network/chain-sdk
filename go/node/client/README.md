# Blockchain API Version Discovery

The `client` package provides multi-version discovery and negotiation between
Akash clients (akt CLI, SDKs) and nodes. A node advertises the API versions it
supports; the client picks the best mutual match and instantiates the
corresponding versioned client.

## Discovery Response

Every discovery endpoint returns the same `Akash` protobuf message:

```json
{
  "client_info": { "api_version": "v1beta3" },
  "supported_versions": [
    {
      "api_version": "v1beta4",
      "modules": [
        { "module": "deployment", "version": "v1beta4" },
        { "module": "market",     "version": "v1beta5" },
        { "module": "oracle",     "version": "v2" }
      ],
      "features": []
    },
    {
      "api_version": "v1beta3",
      "modules": [ ... ]
    }
  ],
  "chain_id": "akashnet-2",
  "node_version": "v0.36.0",
  "min_client_version": "v0.30.0"
}
```

| Field | Description |
|-------|-------------|
| `client_info.api_version` | Oldest supported version. Kept for backward compatibility with old clients that only read this field. |
| `supported_versions` | All API versions the node supports, newest first. Each entry lists the per-module versions it bundles. |
| `supported_versions[].features` | Reserved for future capability flags (e.g. `"authz"`, `"feegrant"`). Currently empty. Clients should ignore unknown feature strings. |
| `chain_id` | Network identifier (e.g. `akashnet-2`). |
| `node_version` | Node software version. |
| `min_client_version` | Minimum client version the node accepts. This is advisory data only — the node does not enforce it at the transport level. Clients are expected to read this field and self-check: refuse to proceed if their own version is below the minimum, or warn the user. |

### Backward Compatibility

- **Old clients connecting to new nodes** read only `client_info.api_version`
  (`v1beta3`), ignore the new fields, and continue working.
- **New clients connecting to old nodes** see an empty `supported_versions`
  list and fall back to `client_info.api_version`.

## Endpoints

Discovery is available on three transports. All return the same data.

| Transport         | Endpoint                               | Notes                                                                 |
|-------------------|----------------------------------------|-----------------------------------------------------------------------|
| CometBFT JSON-RPC | `POST / {"method":"akash"}`            | Registered automatically via `init()` when the package is imported.   |
| gRPC              | `akash.discovery.v1.Discovery/GetInfo` | Requires explicit `RegisterDiscoveryService` call on the gRPC server. |
| REST              | `GET /akash/discovery/v1/info`         | gRPC-Gateway route; requires `RegisterDiscoveryHandlerServer`.        |

## Server Side

### 1. Configure the Registry

Create a `VersionRegistry` during node startup. `DefaultRegistry` includes
`v1beta3` and `v1beta4` with their module version mappings.

```go
import aclient "pkg.akt.dev/go/node/client"

registry := aclient.DefaultRegistry(
    aclient.WithChainID(chainID),
    aclient.WithNodeVersion(version.Version),
    aclient.WithMinClientVersion("v0.30.0"),
)

// Set the global registry. This configures the CometBFT JSON-RPC
// "akash" route — no further action needed for JSON-RPC.
aclient.SetRegistry(registry)
```

### 2. Register gRPC Discovery Service

Call `RegisterDiscoveryService` on the gRPC server. The function accepts
`gogogrpc.Server`, so it works both inside Cosmos SDK's
`RegisterGRPCServerWithSkipCheckHeader` and with a bare `*grpc.Server`.

```go
func (app *App) RegisterGRPCServerWithSkipCheckHeader(server gogogrpc.Server, skip bool) {
    app.BaseApp.RegisterGRPCServerWithSkipCheckHeader(server, skip)

    // Register Akash Discovery alongside Cosmos module query services.
    aclient.RegisterDiscoveryService(server, aclient.GetRegistry())
}
```

### 3. Register REST Gateway Route

Register the gRPC-Gateway handler in `RegisterAPIRoutes` so the REST
endpoint is available at `GET /akash/discovery/v1/info`.

```go
func (app *App) RegisterAPIRoutes(apiSvr *api.Server, cfg config.APIConfig) {
    // ... existing route registrations ...

    aclient.RegisterDiscoveryHandlerServer(
        context.Background(),
        apiSvr.GRPCGatewayRouter,
        aclient.NewDiscoveryServer(aclient.GetRegistry()),
    )
}
```

## Client Side

### Using the Discovery Wrapper (Recommended)

The `discovery` sub-package provides typed convenience functions that handle
negotiation and return concrete client interfaces.

```go
import aclient "pkg.akt.dev/go/node/client/discovery"

// For queries (read-only):
lightCl, err := aclient.DiscoverLightClient(ctx, cctx)

// For transactions (read + write):
cl, err := aclient.DiscoverClient(ctx, cctx, opts...)

// For raw queries without tx support:
qcl, err := aclient.DiscoverQueryClient(ctx, cctx)
```

These functions call the CometBFT JSON-RPC `"akash"` route, negotiate the
best version, and return a client for that version.

The `opts` parameter accepts `cltypes.ClientOption` functions from
`pkg.akt.dev/go/node/client/types`. These configure the transaction factory
used by the full `Client` (not needed for read-only light/query clients):

```go
import cltypes "pkg.akt.dev/go/node/client/types"

cl, err := aclient.DiscoverClient(ctx, cctx,
    cltypes.WithGasPrices("0.025uakt"),
    cltypes.WithGasAdjustment(1.5),
    cltypes.WithGas(cltypes.GasSetting{Simulate: true}),
    cltypes.WithSkipConfirm(true),
)
```

### Using the Low-Level API

The base `client` package exposes the negotiation machinery directly via
callback-style `SetupFn` functions.

```go
import aclient "pkg.akt.dev/go/node/client"

var myClient aclient.Client

err := aclient.DiscoverClient(ctx, cctx, func(cl interface{}) error {
    var ok bool
    if myClient, ok = cl.(aclient.Client); !ok {
        return fmt.Errorf("unexpected client type %T", cl)
    }
    return nil
}, opts...)
```

### How Negotiation Works

1. The client calls `queryClientInfo`, which invokes the `"akash"` JSON-RPC
   route on the connected node.
2. `negotiateVersion` inspects the response:
   - If `SupportedVersions` is populated (new node), it intersects
     the server's versions with `clientPreferenceOrder` (newest first)
     and returns the first match.
   - If `SupportedVersions` is empty (old node), it falls back to
     `ClientInfo.ApiVersion`.
3. The negotiated version string selects the factory function
   (`v1beta4.NewClient`, `v1beta3.NewClient`, etc.).

```text
Client preference: [v1beta4, v1beta3]  (newest first)
Server supports:   [v1beta4, v1beta3]
Negotiated:        v1beta4
```

### Negotiation Failure

When the server's `SupportedVersions` and the client's
`clientPreferenceOrder` have no overlap, `negotiateVersion` returns an
empty string. The subsequent `newClientForVersion("")` call fails with
`ErrUnknownClientVersion`. This is a hard error — no fallback is attempted.

The same error is returned when an old node reports a `ClientInfo.ApiVersion`
that the client does not recognise (e.g. a future version the client binary
predates).

Callers should handle `ErrUnknownClientVersion` as a signal that the client
and node are incompatible and the user needs to upgrade one or the other.

### Offline Mode

When `cctx.Offline` is true, discovery is skipped and `v1beta3` is assumed.

## Adding a New API Version

To add a new composite API version (e.g. `v1beta5`):

1. **Define module versions** in `registry.go`:
   ```go
   var v1beta5Modules = []ModuleVersion{
       {Module: "deployment", Version: "v1beta5"},
       // ...
   }
   ```

2. **Add the version constant** in `client.go`:
   ```go
   const VersionV1beta5 = "v1beta5"
   ```

3. **Add to `clientPreferenceOrder`** in `client.go` (newest first):
   ```go
   var clientPreferenceOrder = []string{VersionV1beta5, VersionV1beta4, VersionV1beta3}
   ```

4. **Create the client package** `v1beta5/client.go`.

   **If module interfaces are unchanged**, delegate to the previous version:
   ```go
   package v1beta5

   type Client = v1beta4.Client
   // ...
   func NewClient(ctx context.Context, cctx sdkclient.Context, opts ...cltypes.ClientOption) (Client, error) {
       return v1beta4.NewClient(ctx, cctx, opts...)
   }
   ```

   **If a module introduces a new interface** (e.g. `deployment/v1beta5`
   adds or changes query methods), copy `v1beta4/` and update the imports
   and composite interfaces. Follow the pattern in `v1beta3/query.go` — each
   module's `QueryClient` is imported as a type alias and wired into the
   composite `queryClient` struct:
   ```go
   package v1beta5

   import (
       dtypes "pkg.akt.dev/go/node/deployment/v1beta5" // new module version
       mtypes "pkg.akt.dev/go/node/market/v1beta5"     // unchanged — same import
   )

   // QueryClient exposes the new deployment/v1beta5 interface.
   type QueryClient interface {
       Deployment() dtypes.QueryClient  // updated return type
       Market()     mtypes.QueryClient
       // ...
   }
   ```
   The composite `Client`, `LightClient`, and `TxClient` interfaces must
   be updated to match. This is intentionally a breaking change within
   the versioned package — callers that pin to `v1beta4` are unaffected.

5. **Add switch cases** in `newClientForVersion`, `newLightClientForVersion`,
   and `newQueryClientForVersion` in `client.go`.

6. **Register in `DefaultRegistry`** in `registry.go` (newest first):
   ```go
   func DefaultRegistry(opts ...RegistryOption) *VersionRegistry {
       return NewRegistry([]VersionInfo{
           {ApiVersion: VersionV1beta5, Modules: v1beta5Modules},
           {ApiVersion: VersionV1beta4, Modules: v1beta4Modules},
           {ApiVersion: VersionV1beta3, Modules: v1beta3Modules},
       }, opts...)
   }
   ```

7. **Regenerate protos** if any proto messages changed: `make proto-gen-go`.
