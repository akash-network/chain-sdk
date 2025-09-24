# Functional Tests

Clean, working tests for the Akash Chain SDK.

## Configuration

Based on the working configuration snippet:

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

## Running Tests

```bash
# Run all functional tests
npm run test:functional

# Run specific test
npm run test:functional -- --testPathPattern=query-deployments
```
