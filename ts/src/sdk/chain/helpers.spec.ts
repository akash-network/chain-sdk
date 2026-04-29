import { describe, expect, it, vi } from "vitest";

import type { ServiceLoader } from "../client/createServiceLoader.ts";
import { type SDKMethod, withMetadata } from "../client/sdkMetadata.ts";
import type { MessageDesc, ServiceDesc } from "../client/types.ts";
import type { TxClient } from "../transport/tx/TxClient.ts";
import { msg, SIGNER_KEY, transaction, type TxMessage } from "./helpers.ts";

function createMessageDesc(typeName: string): MessageDesc {
  return {
    $type: typeName,
    encode: vi.fn() as MessageDesc["encode"],
    decode: vi.fn() as MessageDesc["decode"],
    fromPartial: vi.fn((v: unknown) => v) as MessageDesc["fromPartial"],
    toJSON: vi.fn() as MessageDesc["toJSON"],
    fromJSON: vi.fn() as MessageDesc["fromJSON"],
  };
}

function createMethod(data?: unknown) {
  const serviceDesc: ServiceDesc = {
    typeName: "test.v1.MsgService",
    methods: {
      testMethod: {
        kind: "unary" as const,
        name: "testMethod",
        parent: null!,
        input: createMessageDesc("test.v1.MsgTest"),
        output: createMessageDesc("test.v1.MsgTestResponse"),
      },
    },
  };

  return withMetadata(
    vi.fn<SDKMethod>().mockResolvedValue(data),
    {
      path: [0, "testMethod"],
      serviceLoader: {
        loadAt: vi.fn<() => Promise<ServiceDesc>>().mockResolvedValue(serviceDesc),
        getLoadedType: vi.fn(),
      } as unknown as ServiceLoader<readonly (() => unknown)[]>,
    },
  );
}

function createSigner(): TxClient {
  return {
    signAndBroadcast: vi.fn<TxClient["signAndBroadcast"]>().mockResolvedValue({
      code: 0,
      transactionHash: "abc123",
    } as never),
  };
}

describe(msg.name, () => {
  it("creates a TxMessage from a method and data", () => {
    const tx = createMethod();
    const data = { sender: "akash1abc" };
    const result = msg(tx, data);

    expect(result).toEqual({
      tx,
      data,
    });
  });

  it("creates a TxMessage without data", () => {
    const tx = createMethod();
    const result = msg(tx, undefined);

    expect(result).toEqual({
      tx,
      data: undefined,
    });
  });
});

describe(transaction.name, () => {
  it("encodes messages and calls signAndBroadcast", async () => {
    const method = createMethod();
    const messages: TxMessage[] = [msg(method, { sender: "akash1abc" })];
    const signer = createSigner();
    const sdk = { [SIGNER_KEY]: signer };

    const result = await transaction(sdk, messages);

    expect(signer.signAndBroadcast).toHaveBeenCalledWith(
      [{ typeUrl: "/test.v1.MsgTest", value: { sender: "akash1abc" } }],
      undefined,
    );
    expect(result).toEqual(expect.objectContaining({ code: 0 }));
  });

  it("passes options to signAndBroadcast", async () => {
    const method = createMethod();
    const messages: TxMessage[] = [msg(method, { sender: "akash1abc" })];
    const signer = createSigner();
    const sdk = { [SIGNER_KEY]: signer };
    const options = { memo: "test memo" };

    await transaction(sdk, messages, options);

    expect(signer.signAndBroadcast).toHaveBeenCalledWith(
      expect.any(Array),
      options,
    );
  });

  it("encodes multiple messages", async () => {
    const method = createMethod();
    const messages: TxMessage[] = [
      msg(method, { sender: "akash1abc" }),
      msg(method, { sender: "akash1def" }),
    ];
    const signer = createSigner();
    const sdk = { [SIGNER_KEY]: signer };

    await transaction(sdk, messages);

    expect(signer.signAndBroadcast).toHaveBeenCalledWith(
      [
        { typeUrl: "/test.v1.MsgTest", value: { sender: "akash1abc" } },
        { typeUrl: "/test.v1.MsgTest", value: { sender: "akash1def" } },
      ],
      undefined,
    );
  });

  it("throws when signer is not available", async () => {
    const method = createMethod();
    const messages: TxMessage[] = [msg(method, { sender: "akash1abc" })];

    await expect(transaction({}, messages)).rejects.toThrow(
      "Transaction signer is not available in the SDK instance",
    );
  });

  it("throws when method metadata is not found", async () => {
    const method = vi.fn<SDKMethod>().mockResolvedValue(undefined);
    const messages: TxMessage[] = [msg(method, { sender: "akash1abc" })];
    const signer = createSigner();
    const sdk = { [SIGNER_KEY]: signer };

    await expect(transaction(sdk, messages)).rejects.toThrow(
      "No metadata found for method SDK method",
    );
  });

  it("throws when method is not found in service descriptor", async () => {
    const method = withMetadata(
      vi.fn<SDKMethod>().mockResolvedValue(undefined),
      {
        path: [0, "nonExistent"],
        serviceLoader: {
          loadAt: vi.fn<() => Promise<ServiceDesc>>().mockResolvedValue({
            typeName: "test.v1.MsgService",
            methods: {},
          }),
          getLoadedType: vi.fn(),
        } as unknown as ServiceLoader<readonly (() => unknown)[]>,
      },
    );
    const messages: TxMessage[] = [msg(method, { sender: "akash1abc" })];
    const signer = createSigner();
    const sdk = { [SIGNER_KEY]: signer };

    await expect(transaction(sdk, messages)).rejects.toThrow(
      "No method found at path [0, nonExistent] in \"test.v1.MsgService\"",
    );
  });
});
