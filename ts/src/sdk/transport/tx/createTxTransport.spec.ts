import type { DescMessage, DescMethodBiDiStreaming, DescMethodUnary } from "@bufbuild/protobuf";
import type { DeliverTxResponse, StdFee } from "@cosmjs/stargate";
import { describe, expect, it, jest } from "@jest/globals";

import { proto } from "../../../../test/helpers/proto.ts";
import { createAsyncIterable } from "../../client/stream.ts";
import { createTxTransport } from "./createTxTransport.ts";
import type { TxClient, TxRaw } from "./TxClient.ts";
import { TxError } from "./TxError.ts";

describe(createTxTransport.name, () => {
  describe("stream", () => {
    it("throws when `stream` method is called", async () => {
      const { TestServiceSchema } = await setup();
      const transport = createTxTransport({
        client: createMockTxClient(),
        getMessageType,
      });

      await expect(transport.stream(TestServiceSchema.methods.streamMethod, createAsyncIterable([{ test: "input" }]))).rejects.toThrow(/unimplemented/i);
    });
  });

  describe("unary", () => {
    it("calls `estimateFee` if no fee is provided", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({ estimateFee: jest.fn(() => Promise.resolve(fee)) });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" });

      expect(client.estimateFee).toHaveBeenCalled();
      expect(client.sign).toHaveBeenCalledWith(expect.anything(), fee, expect.anything());
    });

    it("calls `estimateFee` with gasMultiplier when provided", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({ estimateFee: jest.fn(() => Promise.resolve(fee)) });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, { gasMultiplier: 1.5 });

      expect(client.estimateFee).toHaveBeenCalledWith(
        expect.anything(),
        "stargate",
        expect.anything(),
        1.5,
      );
    });

    it("calls `estimateFee` if provided only granter", async () => {
      const { TestServiceSchema } = await setup();
      const fee: Partial<StdFee> = {
        granter: "akash1234567890",
      };
      const estimatedFee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(estimatedFee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, { fee });

      expect(client.estimateFee).toHaveBeenCalled();
      expect(client.sign).toHaveBeenCalledWith(expect.anything(), { ...estimatedFee, ...fee }, expect.anything());
    });

    it("does not `estimateFee` if a fee is provided", async () => {
      const { TestServiceSchema } = await setup();
      const client = createMockTxClient();
      const transport = createTxTransport({
        client,
        getMessageType,
      });
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };

      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, {
        fee,
      });

      expect(client.estimateFee).not.toHaveBeenCalled();
      expect(client.sign).toHaveBeenCalledWith(expect.anything(), fee, expect.anything());
    });

    it("signs and broadcasts a transaction with single message", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const txRaw: TxRaw = {
        bodyBytes: new Uint8Array(0),
        authInfoBytes: new Uint8Array(0),
        signatures: [],
      };
      const txResponse: DeliverTxResponse = {
        height: 1,
        txIndex: 0,
        code: 0,
        transactionHash: "123",
        events: [],
        msgResponses: [],
        gasUsed: 1n,
        gasWanted: 1n,
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
        sign: jest.fn(() => Promise.resolve(txRaw)),
        broadcast: jest.fn(() => Promise.resolve(txResponse)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      const afterSign = jest.fn();
      const afterBroadcast = jest.fn();
      const result = await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, {
        memo: "test",
        afterSign,
        afterBroadcast,
        fee,
      });
      const messages = [{
        typeUrl: `/${TestServiceSchema.methods.testMethod.input.$type}`,
        value: { test: "input" },
      }];

      expect(client.estimateFee).not.toHaveBeenCalled();
      expect(client.sign).toHaveBeenCalledWith(messages, fee, "test");
      expect(afterSign).toHaveBeenCalledWith(txRaw);
      expect(client.broadcast).toHaveBeenCalledWith(txRaw);
      expect(afterBroadcast).toHaveBeenCalledWith(txResponse);
      expect(result.message).toEqual({});
    });

    it("decodes response if `msgResponses` has a response", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        broadcast: jest.fn(() => Promise.resolve({
          height: 1,
          txIndex: 0,
          code: 0,
          transactionHash: "123",
          events: [],
          msgResponses: [{
            typeUrl: `/${TestServiceSchema.methods.testMethod.output.$type}`,
            value: new Uint8Array(0),
          }],
          gasUsed: 1n,
          gasWanted: 1n,
        })),
      });
      const transport = createTxTransport({
        client,
        getMessageType: (typeUrl) => ({
          ...getMessageType(),
          decode: jest.fn(() => (typeUrl === `/${TestServiceSchema.methods.testMethod.output.$type}` ? { test: "output", ok: true } : null)),
        }),
      });

      const result = await transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, { fee });
      expect(result.message).toEqual({ test: "output", ok: true });
    });

    it("throws if tx response has a non-zero code", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        broadcast: jest.fn(() => Promise.resolve({
          height: 1,
          txIndex: 0,
          code: 1,
          transactionHash: "123",
          events: [],
          msgResponses: [],
          gasUsed: 1n,
          gasWanted: 1n,
        })),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await expect(transport.unary(TestServiceSchema.methods.testMethod, { test: "input" }, { fee })).rejects.toThrow(TxError);
    });
  });

  describe("batching", () => {
    it("batches multiple calls without explicit fee into a single transaction", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      const [result1, result2] = await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(1);
      expect(client.broadcast).toHaveBeenCalledTimes(1);
      expect(client.sign).toHaveBeenCalledWith(
        [
          { typeUrl: `/${TestServiceSchema.methods.testMethod.input.$type}`, value: { test: "input1" } },
          { typeUrl: `/${TestServiceSchema.methods.testMethod.input.$type}`, value: { test: "input2" } },
        ],
        fee,
        expect.any(String),
      );
      expect(result1.message).toEqual({});
      expect(result2.message).toEqual({});
    });

    it("uses separate queues for calls with different payer", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee: { payer: "payer1" } }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee: { payer: "payer2" } }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(2);
      expect(client.broadcast).toHaveBeenCalledTimes(2);
    });

    it("uses separate queues for calls with different granter", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee: { granter: "granter1" } }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee: { granter: "granter2" } }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(2);
      expect(client.broadcast).toHaveBeenCalledTimes(2);
    });

    it("uses separate queues for calls with same payer but different gasMultiplier", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee: { payer: "payer1" }, gasMultiplier: 1.3 }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee: { payer: "payer1" }, gasMultiplier: 1.5 }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(2);
      expect(client.broadcast).toHaveBeenCalledTimes(2);
    });

    it("batches calls with different gasMultiplier when no payer/granter is set", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { gasMultiplier: 1.3 }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { gasMultiplier: 1.5 }),
      ]);

      // When no payer/granter is set, gasMultiplier doesn't affect batching - uses default queue
      // The first item's gasMultiplier is used for fee estimation
      expect(client.sign).toHaveBeenCalledTimes(1);
      expect(client.broadcast).toHaveBeenCalledTimes(1);
    });

    it("batches calls with the same payer and granter", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee: { payer: "payer1", granter: "granter1" } }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee: { payer: "payer1", granter: "granter1" } }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(1);
      expect(client.broadcast).toHaveBeenCalledTimes(1);
    });

    it("bypasses batching when full fee is provided", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient();
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee }),
      ]);

      expect(client.sign).toHaveBeenCalledTimes(2);
      expect(client.broadcast).toHaveBeenCalledTimes(2);
    });

    it("concatenates memos from batched calls", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { memo: "memo1" }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { memo: "memo2" }),
      ]);

      expect(client.sign).toHaveBeenCalledWith(
        expect.anything(),
        fee,
        "memo1. memo2",
      );
    });

    it("uses default memo when not provided", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" });

      expect(client.sign).toHaveBeenCalledWith(
        expect.anything(),
        fee,
        `akash: ${TestServiceSchema.methods.testMethod.name}`,
      );
    });

    it("truncates memo when it exceeds 256 characters", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      const longMemo = "a".repeat(300);
      await transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { memo: longMemo, fee });

      expect(client.sign).toHaveBeenCalledWith(
        expect.anything(),
        fee,
        expect.stringMatching(/^a{252}\.\.\.$/),
      );
    });

    it("resolves each batched call with the corresponding response", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
        broadcast: jest.fn(() => Promise.resolve({
          height: 1,
          txIndex: 0,
          code: 0,
          transactionHash: "123",
          events: [],
          msgResponses: [
            { typeUrl: `/${TestServiceSchema.methods.testMethod.output.$type}`, value: new Uint8Array(0) },
            { typeUrl: `/${TestServiceSchema.methods.testMethod.output.$type}`, value: new Uint8Array(0) },
          ],
          gasUsed: 1n,
          gasWanted: 1n,
        })),
      });
      let decodeCallCount = 0;
      const transport = createTxTransport({
        client,
        getMessageType: () => ({
          ...getMessageType(),
          decode: jest.fn(() => {
            decodeCallCount++;
            return { result: `output${decodeCallCount}` };
          }),
        }),
      });

      const [result1, result2] = await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }),
      ]);

      expect(result1.message).toEqual({ result: "output1" });
      expect(result2.message).toEqual({ result: "output2" });
    });

    it("rejects all batched calls when transaction fails", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
        broadcast: jest.fn(() => Promise.resolve({
          height: 1,
          txIndex: 0,
          code: 1,
          transactionHash: "123",
          events: [],
          msgResponses: [],
          gasUsed: 1n,
          gasWanted: 1n,
        })),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      const results = await Promise.allSettled([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }),
      ]);

      expect(results[0].status).toBe("rejected");
      expect(results[1].status).toBe("rejected");
      expect((results[0] as PromiseRejectedResult).reason).toBeInstanceOf(TxError);
      expect((results[1] as PromiseRejectedResult).reason).toBeInstanceOf(TxError);
    });
  });

  describe("locking", () => {
    it("processes transactions sequentially to prevent race conditions", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      const executionOrder: string[] = [];
      const client = createMockTxClient({
        estimateFee: jest.fn(() => Promise.resolve(fee)),
        sign: jest.fn(async (messages: { typeUrl: string; value: unknown }[]) => {
          const inputValue = (messages[0].value as { test: string }).test;
          executionOrder.push(`sign-start-${inputValue}`);
          await delay(inputValue === "input1" ? 50 : 10);
          executionOrder.push(`sign-end-${inputValue}`);
          return {
            bodyBytes: new Uint8Array(0),
            authInfoBytes: new Uint8Array(0),
            signatures: [],
          };
        }),
        broadcast: jest.fn(async () => {
          executionOrder.push("broadcast");
          return {
            height: 1,
            txIndex: 0,
            code: 0,
            transactionHash: "123",
            events: [],
            msgResponses: [],
            gasUsed: 1n,
            gasWanted: 1n,
          };
        }),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      // Send two transactions with explicit fees (bypasses batching)
      await Promise.all([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee }),
      ]);

      // First transaction should complete before second starts
      expect(executionOrder).toEqual([
        "sign-start-input1",
        "sign-end-input1",
        "broadcast",
        "sign-start-input2",
        "sign-end-input2",
        "broadcast",
      ]);
    });

    it("continues processing after a failed transaction", async () => {
      const { TestServiceSchema } = await setup();
      const fee: StdFee = {
        amount: [{ denom: "uakt", amount: "100000" }],
        gas: "100000",
      };
      let callCount = 0;
      const client = createMockTxClient({
        sign: jest.fn(async () => {
          callCount++;
          if (callCount === 1) {
            throw new Error("Sign failed");
          }
          return {
            bodyBytes: new Uint8Array(0),
            authInfoBytes: new Uint8Array(0),
            signatures: [],
          };
        }),
      });
      const transport = createTxTransport({
        client,
        getMessageType,
      });

      const results = await Promise.allSettled([
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input1" }, { fee }),
        transport.unary(TestServiceSchema.methods.testMethod, { test: "input2" }, { fee }),
      ]);

      expect(results[0].status).toBe("rejected");
      expect(results[1].status).toBe("fulfilled");
    });
  });

  async function setup() {
    const def = await proto`
      service TestService {
        rpc TestMethod(TestInput) returns (TestOutput);
        rpc StreamMethod(stream TestInput) returns (stream TestOutput);
      }

      message TestInput {
        string test = 1;
      }

      message TestOutput {
        string result = 1;
      }
    `;
    const TestServiceSchema = def.getTsProtoService<{
      testMethod: DescMethodUnary<DescMessage, DescMessage>;
      streamMethod: DescMethodBiDiStreaming<DescMessage, DescMessage>;
    }>("TestService");

    return { TestServiceSchema };
  }

  function createMockTxClient(overrides?: Partial<TxClient>): TxClient {
    return {
      broadcast: jest.fn(() => Promise.resolve({
        height: 1,
        txIndex: 0,
        code: 0,
        transactionHash: "123",
        events: [],
        msgResponses: [],
        gasUsed: 1n,
        gasWanted: 1n,
      })),
      estimateFee: jest.fn(() => Promise.resolve({
        amount: [],
        gas: "100000",
      })),
      sign: jest.fn(() => Promise.resolve({
        bodyBytes: new Uint8Array(0),
        authInfoBytes: new Uint8Array(0),
        signatures: [],
      })),
      ...overrides,
    };
  }

  function getMessageType() {
    return {
      decode: jest.fn(),
      encode: jest.fn(),
      fromPartial: jest.fn(),
    };
  }

  function delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
});
