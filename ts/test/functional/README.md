# Functional Tests

Clean, working tests for the Akash Chain SDK.

## Environment Variables

For deployment transaction tests, you need to set up a test mnemonic:

```bash
# Set a funded testnet account mnemonic for deployment tests
export TEST_MNEMONIC="word1 word2 word3 word4 word5 word6 word7 word8 word9 word10 word11 word12"
```

**Important Security Notes:**
- Only use testnet accounts with test tokens
- Never use production mnemonics in tests

## Configuration

The tests use local testnet endpoints by default:

```typescript
const sdk = createChainNodeSDK({
  query: {
    baseUrl: "http://localhost:9090",  // gRPC endpoint
  },
  tx: {
    baseUrl: "http://localhost:26657",  // RPC endpoint
    signer: wallet,
  },
});
```

Override with environment variables:
```bash
# Local testnet (default)
export QUERY_RPC_URL="http://localhost:9090"
export TX_RPC_URL="http://localhost:26657"
export REST_API_URL="http://localhost:1317"  # For balance checks

# Or use remote testnet
export QUERY_RPC_URL="http://grpc.sandbox-2.aksh.pw:9090"
export TX_RPC_URL="https://rpc.sandbox-2.aksh.pw:443"
```

### Local Testnet Setup

To run tests against a local testnet, chain-sdk is self-sufficient:

**Single command (recommended):**
```bash
TEST_MNEMONIC="word1 word2 ... word12" make test-functional-ts-local
```

This will:
1. Initialize local node (if not already done)
2. Start the local node (if not running)
3. Fund your test account from genesis account
4. Run the functional tests

**Manual steps (if needed):**

1. Initialize local node:
   ```bash
   make local-node-init
   ```

2. Start the local node:
   ```bash
   make local-node-run
   ```

3. Check node status:
   ```bash
   make local-node-status
   ```

4. Stop the node:
   ```bash
   make local-node-stop
   ```

5. Clean node data:
   ```bash
   make local-node-clean
   ```

The tests will automatically check if your account has sufficient balance (100 AKT minimum) and fund it from the genesis account if needed.

## Running Tests

```bash
# Run all functional tests
npm run test:functional

# Run specific test file
npm run test:functional -- --testPathPattern=deployments

# Run with environment variable for deployment tests
TEST_MNEMONIC="your testnet mnemonic here" npm run test:functional

# Run specific deployment transaction test
TEST_MNEMONIC="your testnet mnemonic here" npm test -- --testPathPattern=deployments --testNamePattern="should create a deployment transaction"
```

## Test Types

- **Query Tests**: Test deployment querying functionality (no mnemonic needed)
- **Serialization Tests**: Test protobuf message serialization consistency (no mnemonic needed)  
- **Transaction Tests**: Test actual deployment creation (requires TEST_MNEMONIC)
