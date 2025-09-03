# Akash Network API Migration Guide: v33.4 to v53.4 and v46.*

This document outlines the necessary steps to migrate your Akash Network application from version 33.4 to 53.4, incorporating relevant changes from Cosmos SDK version 46.*. The guide focuses on **breaking changes only** - modules without significant changes are omitted.

## Important Version Note

The Akash Network uses a custom fork of Cosmos SDK v0.53.4-akash.0, which incorporates Akash-specific modifications on top of the upstream Cosmos SDK. The version numbers referenced in this guide correspond to:

- **v33.4**: Earlier Akash Network version
- **v46.***: Intermediate Cosmos SDK version changes  
- **v53.4**: Current Akash Network version using Cosmos SDK v0.53.4-akash.0 fork

## Akash-Specific Changes

### 1. Deployment Module Changes (v1beta3 → v1beta4)

#### Message Structure Updates

**MsgCreateDeployment Changes:**
- `version` field renamed to `hash`
- `deposit` field type changed from `cosmos.base.v1beta1.Coin` to `deposit.Deposit` (from `pkg.akt.dev/go/node/types/deposit/v1`)
- `depositor` field removed (deposit now includes `Amount` and `Sources` fields)
- `groups` field now uses `GroupSpecs` type casting

**Migration Required:**

**Before (v1beta3):**
```go
msg := &v1beta3.MsgCreateDeployment{
    ID:        deploymentID,
    Groups:    groups,
    Version:   versionBytes,
    Deposit:   coin,
    Depositor: depositorAddr,
}
```

**After (v1beta4):**
```go
msg := &v1beta4.MsgCreateDeployment{
    ID:     deploymentID,
    Groups: groups, // Now uses GroupSpecs type
    Hash:   versionBytes, // Renamed from version
    Deposit: deposit.Deposit{
        Amount:  coin,
        Sources: []deposit.Source{deposit.SourceBalance},
    },
}
```

#### Import Path Changes

**Before (v1beta3):**
```go
import (
    "pkg.akt.dev/go/node/deployment/v1beta3"
    "github.com/cosmos/cosmos-sdk/types"
)
```

**After (v1beta4):**
```go
import (
    "pkg.akt.dev/go/node/deployment/v1beta4"
    v1 "pkg.akt.dev/go/node/deployment/v1"
    deposit "pkg.akt.dev/go/node/types/deposit/v1"
)
```

### 2. Market Module Changes (v1beta4 → v1beta5)

#### Bid Structure Updates
- `BidID` field renamed to `ID` 
- `Bid` message now includes `ResourcesOffer` field for enhanced resource specification
- Updated message handling for bid processing
- `ID` field now uses `v1.BidID` type (from `pkg.akt.dev/go/node/market/v1`) instead of local `BidID`

**Migration Required:**

**Before (v1beta4):**
```go
import "pkg.akt.dev/go/node/market/v1beta4"

bid := &v1beta4.Bid{
    BidID:     bidID,
    State:     state,
    Price:     price,
    CreatedAt: createdAt,
}
```

**After (v1beta5):**
```go
import (
    "pkg.akt.dev/go/node/market/v1beta5"
    v1 "pkg.akt.dev/go/node/market/v1"
)

bid := &v1beta5.Bid{
    ID:             bidID, // Field renamed from BidID to ID
    State:          state,
    Price:          price,
    CreatedAt:      createdAt,
    ResourcesOffer: resourcesOffer, // New field
}
```

### 3. Provider Module Changes (v1beta3 → v1beta4)

#### Info Structure Update
- `ProviderInfo` renamed to `Info`
- Simplified provider information structure

**Migration Required:**

**Before (v1beta3):**
```go
import "pkg.akt.dev/go/node/provider/v1beta3"

provider := &v1beta3.Provider{
    Owner:      owner,
    HostURI:    hostURI,
    Attributes: attributes,
    Info:       v1beta3.ProviderInfo{
        EMail:   email,
        Website: website,
    },
}
```

**After (v1beta4):**
```go
import "pkg.akt.dev/go/node/provider/v1beta4"

provider := &v1beta4.Provider{
    Owner:      owner,
    HostURI:    hostURI,
    Attributes: attributes,
    Info:       v1beta4.Info{ // Renamed from ProviderInfo
        EMail:   email,
        Website: website,
    },
}
```

### 4. Resource Type System Updates

#### Storage to Volumes Migration
- `Storage` type replaced with `Volumes` for better resource management
- Enhanced resource specification with ID-based indexing

**Migration Required:**
```go
// Resource migration helper available in migrate package
import "pkg.akt.dev/go/node/migrate"

// Convert v1beta3 resources to v1beta4
v1beta4Resources := migrate.ResourcesFromV1Beta3(resourceID, v1beta3Resources)
```

### 5. Module Registration Changes

#### ConsensusVersion Implementation
All Akash modules now require `ConsensusVersion()` method implementation:

```go
func (am AppModule) ConsensusVersion() uint64 {
    return 1 // Increment for each consensus-breaking change
}
```

### 6. Akash Fork Dependencies

The Akash Network uses custom forks of key dependencies with specific versions:

```go
// go.mod replacements in Akash fork:
replace (
    // Akash fork of Cosmos SDK
    github.com/cosmos/cosmos-sdk => github.com/akash-network/cosmos-sdk v0.53.4-akash.0
    
    // Akash fork of CometBFT
    github.com/cometbft/cometbft => github.com/akash-network/cometbft v0.38.17-akash.2
    
    // Akash fork of gogoproto
    github.com/cosmos/gogoproto => github.com/akash-network/gogoproto v1.7.0-akash.2
)
```

### 7. Migration Helper Functions

The `migrate` package provides helper functions for seamless migration:

```go
import "pkg.akt.dev/go/node/migrate"

// Available migration functions:
migrate.DeploymentFromV1beta3(cdc, deploymentBytes)
migrate.GroupFromV1Beta3(cdc, groupBytes)
migrate.BidFromV1beta4(cdc, bidBytes)
migrate.ProviderFromV1beta3(cdc, providerBytes)
```

---

## Cosmos SDK Changes

### 1. Context Type Changes

#### sdk.Context → context.Context Migration
Multiple keeper methods now use `context.Context` instead of `sdk.Context`:

**Affected Modules:**
- `x/authz`
- `x/bank` 
- `x/mint`
- `x/crisis`
- `x/distribution`
- `x/evidence`
- `x/gov`
- `x/slashing`
- `x/upgrade`

**Migration Required:**

**Before (v0.46.x):**
```go
func (k Keeper) SomeMethod(ctx sdk.Context, params ...interface{}) error {
    // method implementation
}
```

**After (v0.53.x):**
```go
func (k Keeper) SomeMethod(ctx context.Context, params ...interface{}) error {
    // method implementation
}
```

### 2. Message Interface Changes

#### ValidateBasic Method Removal
- `sdk.Msg` interface no longer requires `ValidateBasic()` method
- Validation now performed directly in message server
- `GetSignBytes()` implementations can be removed
- `LegacyMsg` interface no longer needed

**Migration Required:**

**Before (v0.46.x):**
```go
// Remove these methods from your message types:
func (m *MsgYourMessage) ValidateBasic() error { /* ... */ }
func (m *MsgYourMessage) GetSignBytes() []byte { /* ... */ }
```

**After (v0.53.x):**
```go
// Implement validation in message server instead:
func (ms msgServer) YourMessage(goCtx context.Context, msg *MsgYourMessage) (*MsgYourMessageResponse, error) {
    // Perform validation here
    if err := validateMessage(msg); err != nil {
        return nil, err
    }
    // ... rest of implementation
}
```

### 3. GetSigners Method Changes

#### Automatic Signer Inference
- `GetSigners()` no longer required on `Msg` types
- SDK automatically infers signers from `Signer` field

**Migration Required:**

**Before (v0.46.x):**
```go
// Remove GetSigners method:
func (m *MsgYourMessage) GetSigners() []sdk.AccAddress { /* ... */ }
```

**After (v0.53.x):**
```proto
// Ensure your message has proper Signer field:
message MsgYourMessage {
    string signer = 1; // SDK will automatically use this
    // ... other fields
}
```

### 4. Module Interface Updates

#### AppModuleBasic Simplification
- `GetTxCmd()` and `GetQueryCmd()` no longer required in `AppModuleBasic`
- `RandomizedParams` removed from `AppModuleSimulation`

**Migration Required:**

**Before (v0.46.x):**
```go
// Remove these methods from AppModuleBasic:
func (AppModuleBasic) GetTxCmd() *cobra.Command { /* ... */ }
func (AppModuleBasic) GetQueryCmd() *cobra.Command { /* ... */ }

// Remove from AppModuleSimulation:
func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange { /* ... */ }
```

**After (v0.53.x):**
```go
// Add ProposalMsgs method instead:
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
    return []simtypes.WeightedProposalMsg{
        // Define proposal messages for simulation
    }
}
```

### 5. BeginBlock/EndBlock Signature Changes

#### New Function Signatures

**Before (v0.46.x):**
```go
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
    // implementation
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    // implementation
}
```

**After (v0.53.x):**
```go
func (am AppModule) BeginBlock(ctx context.Context) error {
    // implementation
    return nil
}

func (am AppModule) EndBlock(ctx context.Context) error {
    // implementation
    return nil
}

// For modules needing ValidatorUpdate:
func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
    // implementation
    return validatorUpdates, nil
}
```

### 6. CometBFT Migration

#### Import Path Updates
Replace Tendermint imports with CometBFT:

**Before (v0.46.x):**
```go
import (
    "github.com/tendermint/tendermint/..."
    "github.com/tendermint/tm-db"
)
```

**After (v0.53.x):**
```go
import (
    "github.com/cometbft/cometbft/..."
    "github.com/cometbft/cometbft-db"
)
```

### 7. Protobuf Changes

#### Gogo Protobuf Replacement

**Before (v0.46.x):**
```go
// Update go.mod - remove:
replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

// Update imports:
import "github.com/gogo/protobuf/..."
```

**After (v0.53.x):**
```go
// Updated imports:
import "github.com/cosmos/gogoproto/..."
```

### 8. gRPC Service Registration

#### New Node Query Service
Register the new node query service:

```proto
// Add to your gRPC gateway registration:
import "cosmos/base/node/v1beta1/query.proto"
```

### 9. Module Manager Changes

#### Basic Module Manager Removal
- Remove `module.BasicManager` usage
- Use `module.Manager` directly

**Migration Required:**

**Before (v0.46.x):**
```go
basicManager := module.NewBasicManager(modules...)
manager := module.NewManager(modules...)
```

**After (v0.53.x):**
```go
manager := module.Manager{
    Modules: map[string]interface{}{
        // your modules
    },
}
```

### 10. Collections Migration

#### Error Handling Updates
Some modules now use `collections.ErrNotFound`:

```go
// Update error handling:
import "cosmossdk.io/collections"

if errors.Is(err, collections.ErrNotFound) {
    // handle not found case
}
```

---

## Migration Checklist

### Akash-Specific Updates
- [ ] Update deployment module imports and message structures (v1beta3 → v1beta4)
- [ ] Update deposit structure to use `deposit.Deposit` with `Amount` and `Sources` fields
- [ ] Update market module for ResourcesOffer field and BidID → ID rename (v1beta4 → v1beta5)  
- [ ] Update provider module Info structure (v1beta3 → v1beta4)
- [ ] Implement `ConsensusVersion()` in all modules
- [ ] Update resource type handling (Storage → Volumes)
- [ ] Use migration helper functions from `migrate` package
- [ ] Update go.mod to use Akash forks (cosmos-sdk, cometbft, gogoproto)

### Cosmos SDK Updates
- [ ] Update keeper methods to use `context.Context`
- [ ] Remove `ValidateBasic()` and `GetSignBytes()` from messages
- [ ] Remove `GetSigners()` method from messages
- [ ] Update `BeginBlock`/`EndBlock` signatures
- [ ] Remove deprecated methods from `AppModuleBasic`
- [ ] Replace Tendermint imports with CometBFT
- [ ] Update protobuf imports (gogo → gogoproto)
- [ ] Register new gRPC services
- [ ] Remove basic module manager usage
- [ ] Update error handling for collections

### Testing and Validation
- [ ] Update unit tests for new message structures
- [ ] Test migration functions with existing data
- [ ] Validate gRPC endpoint functionality
- [ ] Test consensus version upgrades
- [ ] Verify CometBFT integration

---

## Additional Resources

- [Cosmos SDK Upgrade Guide](https://docs.cosmos.network/main/build/migrations/upgrading)
- [Akash Network Documentation](https://docs.akash.network)
- [CometBFT Migration Guide](https://docs.cometbft.com)
- [In-Place Store Migrations (ADR-041)](https://docs.cosmos.network/main/build/architecture/adr-041-in-place-store-migrations)

For questions or issues during migration, please refer to the [Akash Network Support](https://github.com/akash-network/support) repository.
