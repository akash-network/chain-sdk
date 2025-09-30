import { describe, expect, it } from "@jest/globals";
import { mock } from "jest-mock-extended";

import type { TxClient } from "../transport/tx/TxClient.ts";
import { createChainNodeWebSDK } from "./createChainNodeWebSDK.ts";

describe(createChainNodeWebSDK.name, () => {
  it("creates ChainNodeSDK with tx transport", () => {
    const sdk = createChainNodeWebSDK({
      query: { baseUrl: "http://localhost:1317" },
      tx: {
        signer: mock<TxClient>(),
      },
    });

    expect(sdk.akash).toBeDefined();
    expect(sdk.cosmos).toBeDefined();
  });

  it("creates ChainNodeSDK without tx transport", async () => {
    const sdk = createChainNodeWebSDK({
      query: { baseUrl: "http://localhost:1317" },
    });

    expect(sdk.akash).toBeDefined();
    expect(sdk.cosmos).toBeDefined();
    await expect(sdk.akash.provider.v1beta4.createProvider({
      attributes: [],
      hostUri: "http://localhost:26657",
      info: undefined,
      owner: "akash1...",
    })).rejects.toThrow(/"tx" option is not provided/);
  });
});
