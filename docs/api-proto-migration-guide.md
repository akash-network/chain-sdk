# Akash Network Protobuf Migration Guide: v0.33.4 to v0.53.4

This document outlines the protobuf definition changes required to migrate your Akash Network protobuf files from version 0.33.4 to 0.53.4. This guide focuses specifically on **protobuf schema changes** and complements the main API migration guide.

## Important Version Note

The Akash Network protobuf definitions follow the same versioning as the main API:

- **v0.33.4**: Earlier Akash Network protobuf schemas
- **v0.53.4**: Current Akash Network protobuf schemas

## Akash-Specific Protobuf Changes

### 1. Deployment Module Changes (v1beta3 → v1beta4)

#### MsgCreateDeployment Proto Definition Updates

**Key Changes:**
- `version` field renamed to `hash`
- `deposit` field type changed from `cosmos.base.v1beta1.Coin` to `akash.base.deposit.v1.Deposit`
- `depositor` field removed
- Import dependencies updated

**Before (v1beta3):**
```protobuf
// akash/deployment/v1beta3/deploymentmsg.proto
syntax = "proto3";
package akash.deployment.v1beta3;

import "gogoproto/gogo.proto";
import "akash/deployment/v1beta3/deployment.proto";
import "akash/deployment/v1beta3/groupspec.proto";
import "cosmos/base/v1beta1/coin.proto";

option go_package = "pkg.akt.dev/go/node/deployment/v1beta3";

message MsgCreateDeployment {
  option (gogoproto.equal) = false;

  DeploymentID id = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customname) = "ID",
    (gogoproto.jsontag)    = "id",
    (gogoproto.moretags)   = "yaml:\"id\""
  ];
  
  repeated GroupSpec groups = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "groups",
    (gogoproto.moretags) = "yaml:\"groups\""
  ];
  
  bytes version = 3 [
    (gogoproto.jsontag)  = "version",
    (gogoproto.moretags) = "yaml:\"version\""
  ];
  
  cosmos.base.v1beta1.Coin deposit = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "deposit",
    (gogoproto.moretags) = "yaml:\"deposit\""
  ];
  
  // Depositor pays for the deposit
  string depositor = 5 [
    (gogoproto.jsontag)  = "depositor",
    (gogoproto.moretags) = "yaml:\"depositor\""
  ];
}
```

**After (v1beta4):**
```protobuf
// akash/deployment/v1beta4/deploymentmsg.proto
syntax = "proto3";
package akash.deployment.v1beta4;

import "gogoproto/gogo.proto";
import "akash/deployment/v1/deployment.proto";
import "akash/deployment/v1beta4/groupspec.proto";
import "akash/base/deposit/v1/deposit.proto";

option go_package = "pkg.akt.dev/go/node/deployment/v1beta4";

message MsgCreateDeployment {
  option (gogoproto.equal) = false;

  // ID is the unique identifier of the deployment.
  akash.deployment.v1.DeploymentID id = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customname) = "ID",
    (gogoproto.jsontag)    = "id",
    (gogoproto.moretags)   = "yaml:\"id\""
  ];

  // GroupSpec is a list of group specifications for the deployment.
  repeated GroupSpec groups = 2 [
    (gogoproto.nullable)     = false,
    (gogoproto.castrepeated) = "GroupSpecs",
    (gogoproto.jsontag)      = "groups",
    (gogoproto.moretags)     = "yaml:\"groups\""
  ];

  // Hash of the deployment (renamed from version).
  bytes hash = 3 [
    (gogoproto.jsontag)  = "hash",
    (gogoproto.moretags) = "yaml:\"hash\""
  ];

  // Deposit specifies the amount and source of deployment deposit.
  akash.base.deposit.v1.Deposit deposit = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "deposit",
    (gogoproto.moretags) = "yaml:\"deposit\""
  ];
}
```

#### New Deposit Proto Definition

The new deposit structure requires a separate proto file:

**New File: akash/base/deposit/v1/deposit.proto**
```protobuf
syntax = "proto3";
package akash.base.deposit.v1;

import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";

option go_package = "pkg.akt.dev/go/node/types/deposit/v1";

// Source is an enum which lists source of funds for deployment deposit.
enum Source {
  option (gogoproto.goproto_enum_prefix) = false;

  // Prefix should start with 0 in enum. So declaring dummy state.
  invalid = 0 [(gogoproto.enumvalue_customname) = "SourceInvalid"];
  // DepositSourceBalance denotes account balance as source of funds
  balance = 1 [(gogoproto.enumvalue_customname)  = "SourceBalance"];
  // DepositSourceGrant denotes authz grants as source of funds
  grant = 2 [(gogoproto.enumvalue_customname)  = "SourceGrant"];
}

// Deposit is a data type used by MsgCreateDeployment, MsgDepositDeployment and MsgCreateBid
message Deposit {
  // Amount specifies the coins to include in the deposit.
  cosmos.base.v1beta1.Coin amount = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.moretags) = "yaml:\"amount\""
  ];

  // Sources list of deposit sources, each entry must be unique
  repeated Source sources = 5 [
    (gogoproto.castrepeated) = "Sources",
    (gogoproto.jsontag)      = "deposit_sources",
    (gogoproto.moretags)     = "yaml:\"deposit_sources\""
  ];
}
```

### 2. Market Module Changes (v1beta4 → v1beta5)

#### Bid Proto Definition Updates

**Key Changes:**
- `BidID` field renamed to `ID`
- ID field type changed to `akash.market.v1.BidID`
- Added `ResourcesOffer` field
- Enhanced resource specification

**Before (v1beta4):**
```protobuf
// akash/market/v1beta4/bid.proto (partial)
message Bid {
  option (gogoproto.equal)            = false;
  option (gogoproto.goproto_stringer) = false;

  BidID bid_id = 1 [
    (gogoproto.customname) = "BidID",
    (gogoproto.nullable)   = false,
    (gogoproto.jsontag)    = "id",
    (gogoproto.moretags)   = 'yaml:"id"'
  ];

  State state = 2 [
    (gogoproto.jsontag)  = "state",
    (gogoproto.moretags) = 'yaml:"state"'
  ];

  cosmos.base.v1beta1.DecCoin price = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "price",
    (gogoproto.moretags) = 'yaml:"price"'
  ];
  
  int64 created_at = 4;

  repeated ResourceOffer resources_offer = 5 [
    (gogoproto.nullable)     = false,
    (gogoproto.castrepeated) = "ResourcesOffer",
    (gogoproto.customname)   = "ResourcesOffer",
    (gogoproto.jsontag)      = "resources_offer",
    (gogoproto.moretags)     = 'yaml:"resources_offer"'
  ];
}
```

**After (v1beta5):**
```protobuf
// akash/market/v1beta5/bid.proto
syntax = "proto3";
package akash.market.v1beta5;

import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";
import "akash/market/v1beta5/resourcesoffer.proto";
import "akash/market/v1/bid.proto";

option go_package = "pkg.akt.dev/go/node/market/v1beta5";

message Bid {
  option (gogoproto.equal)            = true;
  option (gogoproto.goproto_stringer) = false;

  enum State {
    option (gogoproto.goproto_enum_prefix) = false;

    invalid = 0 [(gogoproto.enumvalue_customname) = "BidStateInvalid"];
    open = 1 [(gogoproto.enumvalue_customname) = "BidOpen"];
    active = 2 [(gogoproto.enumvalue_customname) = "BidActive"];
    lost = 3 [(gogoproto.enumvalue_customname) = "BidLost"];
    closed = 4 [(gogoproto.enumvalue_customname) = "BidClosed"];
  }

  // BidID stores owner and all other seq numbers (renamed from bid_id to id)
  akash.market.v1.BidID id = 1 [
    (gogoproto.customname) = "ID",
    (gogoproto.nullable)   = false,
    (gogoproto.jsontag)    = "id",
    (gogoproto.moretags)   = "yaml:\"id\""
  ];

  State state = 2 [
    (gogoproto.jsontag)  = "state",
    (gogoproto.moretags) = "yaml:\"state\""
  ];

  cosmos.base.v1beta1.DecCoin price = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "price",
    (gogoproto.moretags) = "yaml:\"price\""
  ];

  int64 created_at = 4 [
    (gogoproto.jsontag)  = "createdAt",
    (gogoproto.moretags) = "yaml:\"created_at\""
  ];

  // ResourceOffer is a list of offers (enhanced resource specification)
  ResourcesOffer resources_offer = 5 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "resources_offer",
    (gogoproto.moretags) = "yaml:\"resources_offer\""
  ];
}
```

#### MsgCreateBid Proto Updates

**Before (v1beta4):**
```protobuf
// Note: v1beta4 used different message structure
message MsgCreateBid {
  // Used local BidID and cosmos.base.v1beta1.Coin for deposit
}
```

**After (v1beta5):**
```protobuf
// akash/market/v1beta5/bidmsg.proto
syntax = "proto3";
package akash.market.v1beta5;

import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";
import "akash/market/v1beta5/resourcesoffer.proto";
import "akash/market/v1/bid.proto";
import "akash/base/deposit/v1/deposit.proto";

option go_package = "pkg.akt.dev/go/node/market/v1beta5";

message MsgCreateBid {
  option (gogoproto.equal) = false;

  akash.market.v1.BidID id = 1 [
    (gogoproto.customname) = "ID",
    (gogoproto.nullable)   = false,
    (gogoproto.jsontag)    = "id",
    (gogoproto.moretags)   = "yaml:\"id\""
  ];

  cosmos.base.v1beta1.DecCoin price = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "price",
    (gogoproto.moretags) = "yaml:\"price\""
  ];

  // Uses new deposit structure
  akash.base.deposit.v1.Deposit deposit = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "deposit",
    (gogoproto.moretags) = "yaml:\"deposit\""
  ];

  // Enhanced resource specification
  repeated ResourceOffer resources_offer = 4 [
    (gogoproto.nullable)     = false,
    (gogoproto.castrepeated) = "ResourcesOffer",
    (gogoproto.customname)   = "ResourcesOffer",
    (gogoproto.jsontag)      = "resources_offer",
    (gogoproto.moretags)     = "yaml:\"resources_offer\""
  ];
}
```

### 3. Provider Module Changes (v1beta3 → v1beta4)

#### Provider Info Structure Update

**Key Changes:**
- `ProviderInfo` message renamed to `Info`
- Simplified structure with same fields
- Updated import paths

**Before (v1beta3):**
```protobuf
// akash/provider/v1beta3/provider.proto
syntax = "proto3";
package akash.provider.v1beta3;

import "gogoproto/gogo.proto";
import "akash/base/v1beta3/attribute.proto";

option go_package = "pkg.akt.dev/go/node/provider/v1beta3";

// ProviderInfo
message ProviderInfo {
  string email = 1 [
    (gogoproto.customname) = "EMail",
    (gogoproto.jsontag)    = "email",
    (gogoproto.moretags)   = "yaml:\"email\""
  ];
  string website = 2 [
    (gogoproto.jsontag)    = "website",
    (gogoproto.moretags)   = "yaml:\"website\""
  ];
}

message MsgCreateProvider {
  option (gogoproto.equal) = false;

  string owner = 1;
  string host_uri = 2;
  repeated akash.base.v1beta3.Attribute attributes = 3;
  
  ProviderInfo info = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "info",
    (gogoproto.moretags) = "yaml:\"info\""
  ];
}
```

**After (v1beta4):**
```protobuf
// akash/provider/v1beta4/provider.proto
syntax = "proto3";
package akash.provider.v1beta4;

import "gogoproto/gogo.proto";
import "cosmos_proto/cosmos.proto";
import "akash/base/attributes/v1/attribute.proto";

option go_package = "pkg.akt.dev/go/node/provider/v1beta4";

// Info (renamed from ProviderInfo)
message Info {
  string email = 1 [
    (gogoproto.customname) = "EMail",
    (gogoproto.jsontag)    = "email",
    (gogoproto.moretags)   = "yaml:\"email\""
  ];
  string website = 2 [
    (gogoproto.jsontag)    = "website",
    (gogoproto.moretags)   = "yaml:\"website\""
  ];
}

// akash/provider/v1beta4/msg.proto - MsgCreateProvider
message MsgCreateProvider {
  option (gogoproto.equal) = false;

  string owner = 1 [
    (cosmos_proto.scalar) = "cosmos.AddressString",
    (gogoproto.jsontag)   = "owner",
    (gogoproto.moretags)  = "yaml:\"owner\""
  ];
  
  string host_uri = 2 [
    (gogoproto.customname) = "HostURI",
    (gogoproto.jsontag)    = "host_uri",
    (gogoproto.moretags)   = "yaml:\"host_uri\""
  ];
  
  repeated akash.base.attributes.v1.Attribute attributes = 3 [
    (gogoproto.castrepeated) = "pkg.akt.dev/go/node/types/attributes/v1.Attributes",
    (gogoproto.nullable)     = false,
    (gogoproto.jsontag)      = "attributes",
    (gogoproto.moretags)     = "yaml:\"attributes\""
  ];
  
  akash.provider.v1beta4.Info info = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.jsontag)  = "info",
    (gogoproto.moretags) = "yaml:\"info\""
  ];
}
```

## Import Path Changes

### 1. Deployment Module Imports

**Before (v1beta3):**
```protobuf
import "akash/deployment/v1beta3/deployment.proto";
import "akash/deployment/v1beta3/groupspec.proto";
import "cosmos/base/v1beta1/coin.proto";
```

**After (v1beta4):**
```protobuf
import "akash/deployment/v1/deployment.proto";
import "akash/deployment/v1beta4/groupspec.proto";
import "akash/base/deposit/v1/deposit.proto";
```

### 2. Market Module Imports

**Before (v1beta4):**
```protobuf
// Basic market imports in v1beta4
import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";
```

**After (v1beta5):**
```protobuf
import "gogoproto/gogo.proto";
import "cosmos/base/v1beta1/coin.proto";
import "akash/market/v1beta5/resourcesoffer.proto";
import "akash/market/v1/bid.proto";
import "akash/base/deposit/v1/deposit.proto";
```

### 3. Provider Module Imports

**Before (v1beta3):**
```protobuf
import "akash/base/v1beta3/attribute.proto";
```

**After (v1beta4):**
```protobuf
import "cosmos_proto/cosmos.proto";
import "akash/base/attributes/v1/attribute.proto";
```

### 4. New Required Proto Files

Add these new proto files for v1beta4+ compatibility:

```protobuf
// akash/base/deposit/v1/deposit.proto - New deposit structure
// akash/deployment/v1/deployment.proto - Core deployment types
// akash/market/v1/bid.proto - Core market types
// akash/market/v1beta5/resourcesoffer.proto - Enhanced resource offers
```

## Service Definition Updates

### 1. Deployment Service (v1beta4)

**New Service Definition:**
```protobuf
// akash/deployment/v1beta4/service.proto
syntax = "proto3";
package akash.deployment.v1beta4;

import "akash/deployment/v1beta4/deploymentmsg.proto";
import "akash/deployment/v1beta4/groupmsg.proto";
import "akash/deployment/v1beta4/paramsmsg.proto";
import "cosmos/msg/v1/msg.proto";

option go_package = "pkg.akt.dev/go/node/deployment/v1beta4";

service Msg {
  option (cosmos.msg.v1.service) = true;

  rpc CreateDeployment(MsgCreateDeployment) returns (MsgCreateDeploymentResponse);
  rpc UpdateDeployment(MsgUpdateDeployment) returns (MsgUpdateDeploymentResponse);
  rpc CloseDeployment(MsgCloseDeployment) returns (MsgCloseDeploymentResponse);
  rpc CloseGroup(MsgCloseGroup) returns (MsgCloseGroupResponse);
  rpc PauseGroup(MsgPauseGroup) returns (MsgPauseGroupResponse);
  rpc StartGroup(MsgStartGroup) returns (MsgStartGroupResponse);
  rpc UpdateParams(MsgUpdateParams) returns (MsgUpdateParamsResponse);
}
```

### 2. Market Service (v1beta5)

**Enhanced Service Definition:**
```protobuf
// akash/market/v1beta5/service.proto
syntax = "proto3";
package akash.market.v1beta5;

import "akash/market/v1beta5/bidmsg.proto";
import "akash/market/v1beta5/leasemsg.proto";
import "akash/market/v1beta5/paramsmsg.proto";
import "cosmos/msg/v1/msg.proto";

option go_package = "pkg.akt.dev/go/node/market/v1beta5";

service Msg {
  option (cosmos.msg.v1.service) = true;
  
  rpc CreateBid(MsgCreateBid) returns (MsgCreateBidResponse);
  rpc CloseBid(MsgCloseBid) returns (MsgCloseBidResponse);
  rpc WithdrawLease(MsgWithdrawLease) returns (MsgWithdrawLeaseResponse);
  rpc CreateLease(MsgCreateLease) returns (MsgCreateLeaseResponse);
  rpc CloseLease(MsgCloseLease) returns (MsgCloseLeaseResponse);
  rpc UpdateParams(MsgUpdateParams) returns (MsgUpdateParamsResponse);
}
```

## Protobuf Compilation Updates

### 1. Buf Configuration Changes

Update your `buf.yaml` to exclude deprecated versions:

```yaml
version: v2
modules:
  - path: proto/node
    excludes:
      - proto/node/akash/audit/v1beta3
      - proto/node/akash/base/v1beta3
      - proto/node/akash/cert/v1beta3
      - proto/node/akash/deployment/v1beta3  # Exclude old version
      - proto/node/akash/escrow/v1beta3
      - proto/node/akash/market/v1beta4      # Exclude old version
      - proto/node/akash/provider/v1beta3    # Exclude old version
      - proto/node/akash/take/v1beta3
```

### 2. Code Generation Updates

Update your protobuf generation scripts to handle new imports:

```bash
# Ensure new proto files are included in generation
protoc --go_out=. --go-grpc_out=. \
  proto/node/akash/base/deposit/v1/deposit.proto \
  proto/node/akash/deployment/v1beta4/*.proto \
  proto/node/akash/market/v1beta5/*.proto \
  proto/node/akash/provider/v1beta4/*.proto
```

## Migration Checklist

### Protobuf Schema Updates
- [ ] Update deployment module proto files (v1beta3 → v1beta4)
  - [ ] Rename `version` field to `hash` in MsgCreateDeployment
  - [ ] Replace `cosmos.base.v1beta1.Coin deposit` with `akash.base.deposit.v1.Deposit deposit`
  - [ ] Remove `depositor` field
  - [ ] Update import statements
- [ ] Update market module proto files (v1beta4 → v1beta5)
  - [ ] Rename `BidID bid_id` to `akash.market.v1.BidID id` in Bid message
  - [ ] Add `ResourcesOffer resources_offer` field
  - [ ] Update MsgCreateBid to use new deposit structure
- [ ] Update provider module proto files (v1beta3 → v1beta4)
  - [ ] Rename `ProviderInfo` message to `Info`
  - [ ] Update import paths
- [ ] Create new proto files
  - [ ] Add `akash/base/deposit/v1/deposit.proto`
  - [ ] Add `akash/market/v1beta5/resourcesoffer.proto`

### Build Configuration Updates
- [ ] Update `buf.yaml` to exclude deprecated proto versions
- [ ] Update protobuf generation scripts
- [ ] Update import paths in existing proto files
- [ ] Verify protobuf compilation with new schemas

### Service Definition Updates
- [ ] Update gRPC service definitions for new message types
- [ ] Update query service endpoints
- [ ] Verify service registration in application

### Testing and Validation
- [ ] Test protobuf compilation with new schemas
- [ ] Validate generated Go code compiles correctly
- [ ] Test gRPC service functionality
- [ ] Verify backward compatibility where needed

---

## Cosmos SDK Specific Changes

### 1. Authz Module Changes

#### Context Migration in Authorization Interface

**Key Changes:**
- Authorization `Accept` method now uses `context.Context` instead of `sdk.Context`
- Enhanced authorization scoping with new enum types
- Updated protobuf annotations for amino compatibility

**Before (v0.33.4):**
```protobuf
// cosmos/authz/v1beta1/authz.proto
message Grant {
  google.protobuf.Any authorization = 1;
  google.protobuf.Timestamp expiration = 2;
}
```

**After (v0.53.4):**
```protobuf
// cosmos/authz/v1beta1/authz.proto  
message Grant {
  google.protobuf.Any authorization = 1;
  google.protobuf.Timestamp expiration = 2;
}

// Enhanced with amino annotations
message DepositAuthorization {
  option (cosmos_proto.message_added_in)     = "chain-sdk v0.1.0";
  option (cosmos_proto.implements_interface) = "cosmos.authz.v1beta1.Authorization";
  option (amino.name)                        = "akash/DepositAuthorization";
  
  cosmos.base.v1beta1.Coin spend_limit = 1;
  repeated Scope scopes = 2;
}
```

#### New Authorization Scoping System

**Enhanced Akash Authorization:**
```protobuf
// akash/escrow/v1/authz.proto
message DepositAuthorization {
  enum Scope {
    option (gogoproto.goproto_enum_prefix) = false;

    invalid = 0 [(gogoproto.enumvalue_customname) = "DepositScopeInvalid"];
    deployment = 1 [(gogoproto.enumvalue_customname) = "DepositScopeDeployment"];
    bid = 2 [(gogoproto.enumvalue_customname) = "DepositScopeBid"];
  }

  cosmos.base.v1beta1.Coin spend_limit = 1;
  repeated Scope scopes = 2;
}
```

### 2. Query Service Updates

#### New Query Methods (v0.33.4 → v0.53.4)

**Added Query Endpoints:**
```protobuf
service Query {
  // Returns list of Authorization, granted to the grantee by the granter.
  rpc Grants(QueryGrantsRequest) returns (QueryGrantsResponse);
  
  // getGranterGrants returns list of GrantAuthorization, granted by granter.
  // Added in v0.53.4
  rpc GranterGrants(QueryGranterGrantsRequest) returns (QueryGranterGrantsResponse);
  
  // getGranteeGrants returns a list of GrantAuthorization by grantee.
  // Added in v0.53.4  
  rpc GranteeGrants(QueryGranteeGrantsRequest) returns (QueryGranteeGrantsResponse);
}
```

### 3. Auth Module Updates

#### UpdateParams Message (v0.33.4 → v0.53.4)

**New Governance Operation:**
```protobuf
// cosmos/auth/v1beta1/tx.proto
message MsgUpdateParams {
  option (cosmos.msg.v1.signer) = "authority";

  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string authority = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  // params defines the x/auth parameters to update.
  Params params = 2 [(gogoproto.nullable) = false, (amino.dont_omitempty) = true];
}
```

### 4. Amino Annotations

#### Enhanced Protobuf Annotations

**New Required Imports:**
```protobuf
import "amino/amino.proto";
import "cosmos_proto/cosmos.proto";
```

**Updated Message Annotations:**
```protobuf
message YourAuthorization {
  option (cosmos_proto.implements_interface) = "cosmos.authz.v1beta1.Authorization";
  option (amino.name) = "your-module/YourAuthorization";
  
  // fields...
}
```

### 5. Staking Authorization Updates

#### Enhanced Staking Authorizations

**Updated Authorization Types:**
```protobuf
// cosmos/staking/v1beta1/authz.proto
enum AuthorizationType {
  // UNSPECIFIED defines an invalid authorization type
  AUTHORIZATION_TYPE_UNSPECIFIED = 0;
  // DELEGATE defines a delegation authorization type
  AUTHORIZATION_TYPE_DELEGATE = 1;
  // UNDELEGATE defines an undelegation authorization type  
  AUTHORIZATION_TYPE_UNDELEGATE = 2;
  // REDELEGATE defines a redelegation authorization type
  AUTHORIZATION_TYPE_REDELEGATE = 3;
  // CANCEL_UNBONDING_DELEGATION defines cancel unbonding delegation authorization type
  AUTHORIZATION_TYPE_CANCEL_UNBONDING_DELEGATION = 4;
}
```

### 6. Bank Authorization Updates

#### Send Authorization Changes

**Enhanced Send Authorization:**
```protobuf
// cosmos/bank/v1beta1/authz.proto
message SendAuthorization {
  option (cosmos_proto.implements_interface) = "cosmos.authz.v1beta1.Authorization";
  option (amino.name) = "cosmos-sdk/SendAuthorization";

  repeated cosmos.base.v1beta1.Coin spend_limit = 1;
  repeated string allow_list = 2;
}
```

### 7. Module Account Permissions

#### Updated Module Account Structure

**Enhanced Permissions System:**
```protobuf
// cosmos/auth/v1beta1/auth.proto
message ModuleAccount {
  option (cosmos_proto.implements_interface) = "cosmos.auth.v1beta1.AccountI";
  option (amino.name) = "cosmos-sdk/ModuleAccount";

  BaseAccount base_account = 1 [(gogoproto.embed) = true];
  string name = 2;
  repeated string permissions = 3;
}
```

## Migration Impact on Authorization

### 1. Interface Changes

**Before (v0.33.4):**
```go
type Authorization interface {
    Accept(ctx sdk.Context, msg sdk.Msg) (AcceptResponse, error)
    ValidateBasic() error
    MsgTypeURL() string
}
```

**After (v0.53.4):**
```go
type Authorization interface {
    Accept(ctx context.Context, msg sdk.Msg) (AcceptResponse, error)
    ValidateBasic() error  
    MsgTypeURL() string
}
```

### 2. Akash-Specific Authorization Updates

**Enhanced Deposit Authorization:**
```go
func (m *DepositAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
    return m.TryAccept(ctx, msg, false)
}

func (m *DepositAuthorization) TryAccept(ctx context.Context, msg sdk.Msg, partial bool) (authz.AcceptResponse, error) {
    // Enhanced logic with scope checking
    switch mt := msg.(type) {
    case *MsgAccountDeposit:
        // Handle escrow deposits
    case *MsgCreateDeployment:
        // Handle deployment deposits  
    case *MsgCreateBid:
        // Handle bid deposits
    }
}
```

## Cosmos SDK Migration Checklist

### Authz Module Updates
- [ ] Update Authorization interface implementations to use `context.Context`
- [ ] Add amino annotations to custom authorization types
- [ ] Implement new query methods (GranterGrants, GranteeGrants)
- [ ] Update authorization scoping logic
- [ ] Add cosmos_proto annotations for interface implementations

### Auth Module Updates  
- [ ] Implement MsgUpdateParams for governance parameter updates
- [ ] Update module account permission structures
- [ ] Add required protobuf imports (amino, cosmos_proto)

### Staking Module Updates
- [ ] Update to new AuthorizationType enum values
- [ ] Handle CANCEL_UNBONDING_DELEGATION authorization type
- [ ] Update staking authorization logic

### Bank Module Updates
- [ ] Update SendAuthorization with new amino annotations
- [ ] Handle enhanced spend limit and allow list logic

---

## Additional Resources

- [Protocol Buffers Documentation](https://developers.google.com/protocol-buffers)
- [Buf CLI Documentation](https://docs.buf.build)
- [Akash Network Documentation](https://docs.akash.network)
- [Cosmos SDK Protobuf Guide](https://docs.cosmos.network/main/build/building-modules/protobuf-annotations)

For questions or issues during protobuf migration, please refer to the [Akash Network Support](https://github.com/akash-network/support) repository.
