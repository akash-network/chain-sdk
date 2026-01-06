# Functional Tests

TypeScript tests running against a Go mock gRPC server to validate cross-language protobuf compatibility and transaction validation.

## Purpose

These tests verify:
- TypeScript ↔ Go protobuf encoding/decoding works correctly
- Transaction messages pass Go-side `ValidateBasic()` checks
- Custom types (e.g., `LegacyDec`) serialize consistently between languages
- gRPC and HTTP Gateway endpoints handle requests properly

## Running Tests

**Important:** Functional tests must be run via make from the repository root:

```bash
# From repository root (required)
make test-functional-ts
```

This automatically ensures:
- buf binary is installed to `.cache/bin/buf` (version defined in Makefile)
- Go dependencies are vendored (`make modvendor`)
- Mock server binary is built to `.cache/bin/mock-server`
- Environment variables are properly set

Running `npm run test:functional` directly will fail with an error unless:
1. The environment is configured via direnv (which sets `AKASH_DEVCACHE_BIN`)
2. All dependencies are already installed

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
