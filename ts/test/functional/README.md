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
- The test will skip gracefully if TEST_MNEMONIC is not set

## Configuration

The tests use these endpoints by default:

```typescript
const sdk = createChainNodeSDK({
  query: {
    baseUrl: "http://rpc.dev.akash.pub:30090",
  },
  tx: {
    baseUrl: "https://testnetrpc.akashnet.net:443",
    signer: wallet,
  },
});
```

Override with environment variables:
```bash
export QUERY_RPC_URL="http://rpc.dev.akash.pub:30090"
export TX_RPC_URL="https://testnetrpc.akashnet.net:443"
```

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
