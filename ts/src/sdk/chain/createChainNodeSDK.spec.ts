import type { OfflineSigner } from "@cosmjs/proto-signing";
import { describe, expect, it } from "@jest/globals";
import { mock } from "jest-mock-extended";

import { createChainNodeSDK } from "./createChainNodeSDK.ts";

describe(createChainNodeSDK.name, () => {
  it("creates ChainNodeSDK with tx transport", () => {
    const sdk = createChainNodeSDK({
      query: { baseUrl: "http://localhost:1317" },
      tx: {
        baseUrl: "http://localhost:26657",
        signer: mock<OfflineSigner>(),
      },
    });

    expect(sdk.akash).toBeDefined();
    expect(sdk.cosmos).toBeDefined();
  });

  it("creates ChainNodeSDK without tx transport", async () => {
    const sdk = createChainNodeSDK({
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
