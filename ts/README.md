# Akash API TypeScript Bindings

[![npm version](https://badge.fury.io/js/%40akashnetwork%2Fakash-api.svg)](https://badge.fury.io/js/%40akashnetwork%2Fakash-api)
[![License: Apache-2.0](https://img.shields.io/badge/License-apache2.0-yellow.svg)](https://opensource.org/license/apache-2-0)
[![semantic-release: conventionalcommits](https://img.shields.io/badge/semantic--release-conventionalcommits?logo=semantic-release)](https://github.com/semantic-release/semantic-release)

This package provides TypeScript bindings for the Akash API, generated from protobuf definitions.

## Installation

To install the package, run:

```bash
npm install @akashnetwork/chain-sdk
```

## Usage

This package supports commonjs and ESM environments.

### Chain SDK

#### Node.js/Server Environment

```typescript
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { createChainNodeSDK } from "@akashnetwork/chain-sdk/chain";

const mnemonic = "your mnemonic here";
const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, { prefix: "akash" });

// endpoints can be found in https://github.com/akash-network/net
const sdk = createChainNodeSDK({
  query: {
    baseUrl: "http://rpc.dev.akash.pub:31317", // blockchain grpc endpoint url
  },
  tx: {
    baseUrl: 'https://testnetrpc.akashnet.net:443', // blockchain rpc endpoint
    signer: wallet,
  },
});

// Query deployments
const deployments = await sdk.akash.deployment.v1beta4.getDeployments({
  pagination: {
    limit: 1,
  },
});

console.log(deployments);
```

#### Web Environment

```typescript
import { createChainNodeSDK, type TxClient } from "@akashnetwork/chain-sdk/chain/web";

const wallet: TxClient = // kplr or leap wallet object in browser exposed by corresponding extension
const sdk = createChainNodeSDK({
  query: {
    baseUrl: "http://rpc.dev.akash.pub:31317", // grpc gateway api url
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
import { createProviderSDK } from "@akashnetwork/chain-sdk/provider";

const sdk = createProviderSDK({
  baseUrl: "https://provider.provider-02.sandbox-01.aksh.pw:8444",
});

// Get provider status
const status = await sdk.akash.provider.v1.getStatus();
console.log(status);
```

### Stack Definition Language (SDL)

```typescript
import { SDL } from "@akashnetwork/chain-sdk/sdl";

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
