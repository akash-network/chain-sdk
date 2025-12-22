# Functional Tests

TypeScript tests running against a Go mock gRPC server to validate cross-language protobuf compatibility and transaction validation.

## Purpose

These tests verify:
- TypeScript ↔ Go protobuf encoding/decoding works correctly
- Transaction messages pass Go-side `ValidateBasic()` checks
- Custom types (e.g., `LegacyDec`) serialize consistently between languages
- gRPC and HTTP Gateway endpoints handle requests properly

## Running Tests

```bash
# From project root
make test-functional-ts

# Or from ts/ directory
npm run test:functional

# Run specific test
npm run test:functional -- --testPathPattern=deployments

# Update snapshots
npm run test:functional -- -u
```

## Architecture

**Go Mock Server** (`go/testutil/mock/`):
- Standalone gRPC server with HTTP Gateway
- Uses real Cosmos SDK validation logic (`ValidateBasic()`)
- Serves test fixtures from JSON files
- Built automatically by `make test-functional-ts`

**TypeScript Tests**:
- Spawn the mock server binary
- Test queries, transactions, and message serialization
- Verify encoding matches between TS and Go

## Test Categories

1. **Deployment Tests** (`deployments.spec.ts`)
   - Query deployments from mock server
   - Create/validate deployment transactions
   - Test transaction validation errors

2. **Protoc Plugin Tests**
   - `protoc-gen-customtype-patches.spec.ts` - Custom type handling
   - `protoc-gen-sdk-object.spec.ts` - SDK object generation

## Mock Server

Tests run against a Go mock gRPC server that:
- Starts automatically when tests run
- Uses real Cosmos SDK validation logic
- Serves deterministic test fixtures
- Requires no external dependencies
