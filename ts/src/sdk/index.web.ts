import type { createChainNodeWebSDK } from "./chain/createChainNodeWebSDK.ts";

export * from "./index.shared.ts";
export { createChainNodeWebSDK, type ChainNodeWebSDKOptions } from "./chain/createChainNodeWebSDK.ts";
export type ChainNodeWebSDK = ReturnType<typeof createChainNodeWebSDK>;
