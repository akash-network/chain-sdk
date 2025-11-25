# Akash API TypeScript Bindings

[![npm version](https://badge.fury.io/js/%40akashnetwork%2Fchain-sdk.svg)](https://badge.fury.io/js/%40akashnetwork%2Fchain-sdk)
[![License: Apache-2.0](https://img.shields.io/badge/License-apache2.0-yellow.svg)](https://opensource.org/license/apache-2-0)
[![semantic-release: conventionalcommits](https://img.shields.io/badge/semantic--release-conventionalcommits?logo=semantic-release)](https://github.com/semantic-release/semantic-release)

This package provides TypeScript bindings for the Akash API, generated from protobuf definitions.

## Installation

⚠️ **NOTICE:** 

The new Chain SDK for TypeScript is currently in alpha. As such, small breaking changes may occur between versions.
To ensure stability of your own scripts, pin a specific version of the SDK in your package.json (avoid using `^` or `~` in front of version). 

We are actively gathering developer feedback and improving the DX (Developer Experience).
Please report any issues or suggestions via:
* GitHub Issues (preferred)
* [Discord](https://akash.network/docs/getting-started/technical-support/)

To install the package, run:

```bash
npm install @akashnetwork/chain-sdk@alpha
```

## Usage

This package supports commonjs and ESM environments.

### Chain SDK

#### Node.js/Server Environment

This implementation uses gRPC transport to fetch data from blockchain

```typescript
import { createChainNodeSDK, createStargateClient } from "@akashnetwork/chain-sdk";

const mnemonic = "your mnemonic here";
const signer = createStargateClient({
  baseUrl: 'https://rpc.sandbox-2.aksh.pw:443', // blockchain rpc endpoint
  signerMnemonic: mnemonic
});

// endpoints can be found in https://github.com/akash-network/net
const chainSdk = createChainNodeSDK({
  query: {
    baseUrl: "http://grpc.sandbox-2.aksh.pw:9090", // blockchain gRPC endpoint url
  },
  tx: {
    signer,
  },
});

// Query deployments
const deployments = await chainSdk.akash.deployment.v1beta4.getDeployments({
  pagination: {
    limit: 1,
  },
});

console.log(deployments);
```

It's also possible to create `StargateClient` from a `DirectSecp256k1HdWallet` instance:

```ts
import { createChainNodeSDK, createStargateClient } from "@akashnetwork/chain-sdk";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

const mnemonic = "your mnemonic here";
const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, { prefix: "akash" });

const signer = createStargateClient({
  baseUrl: 'https://rpc.sandbox-2.aksh.pw:443', // blockchain rpc endpoint
  signer: wallet
});
```

#### Web Environment

This implementation can be used in both browser and nodejs, since it uses gRPC Gateway transport to fetch data from blockchain

```typescript
import { createChainNodeWebSDK, type TxClient } from "@akashnetwork/chain-sdk/web";

const wallet: TxClient = // kplr or leap wallet object in browser exposed by corresponding extension
const sdk = createChainNodeWebSDK({
  query: {
    baseUrl: "https://api.sandbox-2.aksh.pw:443", // gRPC Gateway api url
  },
  tx: {
    signer: wallet,
  },
});

// Query deployments
const deployments = await sdk.akash.deployment.v1beta4.getDeployments({
  pagination: {
    limit: 1,
  },
});
```

### Provider SDK

Currently provider SDK supports only `getStatus` and `streamStatus` methods over gRPC protocol.

```typescript
import { createProviderSDK } from "@akashnetwork/chain-sdk";

const sdk = createProviderSDK({
  baseUrl: "https://provider.provider-02.sandbox-01.aksh.pw:8444",
});

// Get provider status
const status = await sdk.akash.provider.v1.getStatus();
console.log(status);
```

#### Authentication

The Provider API currently supports two types of authentication:

##### JWT (Recommended)

This is the recommended method for getting authorized access to your resources on different providers. JWT authentication is simpler than mTLS and works even when the blockchain is down.

**Generating a JWT Token**

```ts
import { Secp256k1HdWallet } from "@cosmjs/amino";
import { JwtTokenManager } from "@akashnetwork/chain-sdk"

const wallet = await Secp256k1HdWallet.fromMnemonic(mnemonic, { prefix: "akash" });
const accounts = await wallet.getAccounts();
const tokenManager = new JwtTokenManager(wallet);

// See https://akash.network/roadmap/aep-64/ for details
const token = await tokenManager.generateToken({
  iss: accounts[0].address,
  exp: Math.floor(Date.now() / 1000) + 3600,
  iat: Math.floor(Date.now() / 1000),
  version: "v1",
  leases: { access: "full" },
});

// Use the token to authenticate API requests
const lease = {
  dseq: "...",
  gseq: 1,
  oseq: 1,
};
const leaseDetails = await fetch(`https://some-provider.url:8443/lease/${lease.dseq}/${lease.gseq}/${lease.oseq}/status`, {
  headers: {
    Authorization: `Bearer ${token}`
  },
});
```

If the provider responds with a self-signed certificate, the client must validate it to ensure the provider's identity is correct.

##### mTLS (Legacy)

> **⚠️ Important:** This method of authentication is deprecated and should not be used in new clients. Always prefer JWT authentication as it is simpler and works even when the blockchain is unavailable.

When a client interacts with the Provider API to access lease details, it must establish a TLS connection and present its client certificate pair. If no valid certificate is provided, the API will return an "Unauthenticated" error.

It is essential to store the generated certificate on-chain, as the provider verifies its availability during authentication. See [the documentation](https://akash.network/docs/other-resources/authentication/) for additional details.

**Generating a Client Certificate**

```ts
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { certificateManager } from "@akashnetwork/chain-sdk"
import { fetch, Agent } from 'undici'
import { chainSdk } from "./chainSdk"; // chainSdk created in the example above

const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, { prefix: "akash" });
const accounts = await wallet.getAccounts();

// Generate certificate pair (do this only once, then store and reuse the certificate)
const clientCertPair = await certificateManager.generatePEM(accounts[0].address);

// Store certificate on-chain (one-time operation)
await chainSdk.akash.cert.v1.createCertificate({
  owner: accounts[0].address,
  cert: Buffer.from(clientCertPair.cert, 'utf-8'),
  pubkey: Buffer.from(clientCertPair.publicKey, 'utf-8'),
});

// Use certificate for API requests
const lease = {
  dseq: "...",
  gseq: 1,
  oseq: 1,
};
const leaseDetails = await fetch(`https://some-provider.url:8443/lease/${lease.dseq}/${lease.gseq}/${lease.oseq}/status`, {
  dispatcher: new Agent({
    connect: {
      cert: clientCertPair.cert,
      key: clientCertPair.privateKey
    }
  })
});
```

**Important Notes:**
- Generate the certificate only once and reuse it while it's valid
- Do not create a new certificate for every request
- Verify the provider's identity when it responds with a self-signed certificate


### Stack Definition Language (SDL)

```typescript
import { SDL } from "@akashnetwork/chain-sdk";

const yaml = `
version: "2.0"
services:
  web:
    image: nginx
    expose:
      - port: 80
        as: 80
        to:
          - global: true
`;

const sdl = SDL.fromString(yaml);
const manifest = sdl.manifest();
```

### Contributing

Contributions are welcome. Please submit a pull request or create an issue to discuss the changes you want to make.

### License

This package is licensed under the Apache-2.0.
