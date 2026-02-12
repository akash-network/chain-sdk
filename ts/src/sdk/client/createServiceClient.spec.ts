import type { DescMethod, DescMethodBiDiStreaming, DescMethodClientStreaming, DescMethodServerStreaming, DescMethodUnary } from "@bufbuild/protobuf";
import { describe, expect, it, jest } from "@jest/globals";

import { proto } from "../../../test/helpers/proto.ts";
import type { StreamResponse, Transport, UnaryResponse } from "../transport/types.ts";
import { createServiceClient } from "./createServiceClient.ts";
import { createAsyncIterable } from "./stream.ts";
import type { MessageDesc } from "./types.ts";

describe(createServiceClient.name, () => {
  describe("unary method", () => {
    it("can create unary method", async () => {
      const { TestServiceSchema } = await setup();
      const transport = createTransport("unary", () => ({
        message: { result: "result" },
        header: new Headers(),
        trailer: new Headers(),
      }));

      const client = createServiceClient(TestServiceSchema, transport);

      const abortSignal = new AbortController().signal;
      const headers = {
        "x-test": "test",
      };
      const onHeader = jest.fn();
      const onTrailer = jest.fn();
      const result = await client.testMethod({ test: "input" }, {
        signal: abortSignal,
        timeoutMs: 1000,
        headers,
        onHeader,
        onTrailer,
      });
      expect(result).toEqual({ result: "result" });
      expect(transport.unary).toHaveBeenCalledWith(
        TestServiceSchema.methods.testMethod,
        { test: "input" },
        {
          signal: abortSignal,
          timeoutMs: 1000,
          headers,
          onHeader,
          onTrailer,
        },
      );

      const transportResult = (await transport.unary.mock.results.at(-1)?.value) as Awaited<ReturnType<typeof transport.unary>>;
      expect(onHeader).toHaveBeenCalledWith(transportResult.header);
      expect(onTrailer).toHaveBeenCalledWith(transportResult.trailer);
    });

    async function setup() {
      const def = await proto`
        service TestService {
          rpc TestMethod(TestInput) returns (TestOutput);
        }

        message TestInput {
          string test = 1;
        }

        message TestOutput {
          string result = 1;
        }
      `;
      const TestInputSchema = def.getMessage<"TestInput", { test: string }>("TestInput");
      const TestOutputSchema = def.getMessage<"TestOutput", { result: string }>("TestOutput");
      const TestServiceSchema = def.getTsProtoService<{
        testMethod: DescMethodUnary<typeof TestInputSchema, typeof TestOutputSchema>;
      }>("TestService");

      return {
        TestInputSchema,
        TestOutputSchema,
        TestServiceSchema,
      };
    }
  });

  describe("server streaming method", () => {
    it("can create server streaming method", async () => {
      const { TestServiceSchema } = await setup();
      const results = Array.from({ length: 3 }, (_, i) => ({ result: `result${i + 1}` }));

      const transport = createTransport("stream", () => ({
        message: createAsyncIterable(results.map((result) => ({
          ...result,
        }))),
        header: new Headers(),
        trailer: new Headers(),
      }));

      const client = createServiceClient(TestServiceSchema, transport);

      const abortSignal = new AbortController().signal;
      const headers = {
        "x-test": "test",
      };
      const onHeader = jest.fn();
      const onTrailer = jest.fn();
      const stream = client.testStreamMethod({ test: "input" }, {
        signal: abortSignal,
        timeoutMs: 1000,
        headers,
        onHeader,
        onTrailer,
      });

      expect(await Array.fromAsync(stream)).toEqual(results);
      expect(transport.stream).toHaveBeenCalledWith(
        TestServiceSchema.methods.testStreamMethod,
        expect.anything(),
        {
          signal: abortSignal,
          timeoutMs: 1000,
          headers,
          onHeader,
          onTrailer,
        },
      );
      expect(await Array.fromAsync(transport.stream.mock.lastCall?.at(1) as AsyncIterable<unknown>)).toEqual([{
        test: "input",
      }]);
      const transportResult = (await transport.stream.mock.results.at(-1)?.value) as Awaited<ReturnType<typeof transport.stream>>;
      expect(onHeader).toHaveBeenCalledWith(transportResult.header);
      expect(onTrailer).toHaveBeenCalledWith(transportResult.trailer);
    });

    async function setup() {
      const def = await proto`
        service TestService {
          rpc TestStreamMethod(TestInput) returns (stream TestOutput);
        }

        message TestInput {
          string test = 1;
        }

        message TestOutput {
          string result = 1;
        }
      `;
      const TestInputSchema = def.getMessage<"TestInput", { test: string }>("TestInput");
      const TestOutputSchema = def.getMessage<"TestOutput", { result: string }>("TestOutput");
      const TestServiceSchema = def.getTsProtoService<{
        testStreamMethod: DescMethodServerStreaming<typeof TestInputSchema, typeof TestOutputSchema>;
      }>("TestService");

      return {
        TestInputSchema,
        TestOutputSchema,
        TestServiceSchema,
      };
    }
  });

  describe("client streaming method", () => {
    it("can create client streaming method", async () => {
      const { TestServiceSchema } = await setup();
      const transport = createTransport("stream", () => ({
        message: createAsyncIterable([
          { result: "result" },
        ]),
        header: new Headers(),
        trailer: new Headers(),
      }));
      const client = createServiceClient(TestServiceSchema, transport);

      const abortSignal = new AbortController().signal;
      const headers = {
        "x-test": "test",
      };
      const onHeader = jest.fn();
      const onTrailer = jest.fn();
      const input = Array.from({ length: 3 }, (_, i) => ({ test: `input${i + 1}` }));
      const result = await client.testClientStreamMethod(createAsyncIterable(input), {
        signal: abortSignal,
        timeoutMs: 1000,
        headers,
        onHeader,
        onTrailer,
      });

      expect(result).toEqual({ result: "result" });
      expect(transport.stream).toHaveBeenCalledWith(
        TestServiceSchema.methods.testClientStreamMethod,
        expect.anything(),
        {
          signal: abortSignal,
          timeoutMs: 1000,
          headers,
          onHeader,
          onTrailer,
        },
      );
      expect(await Array.fromAsync(transport.stream.mock.lastCall?.at(1) as AsyncIterable<unknown>)).toEqual(
        input.map((item) => ({
          ...item,
        })),
      );
      const transportResult = (await transport.stream.mock.results.at(-1)?.value) as Awaited<ReturnType<typeof transport.stream>>;
      expect(onHeader).toHaveBeenCalledWith(transportResult.header);
      expect(onTrailer).toHaveBeenCalledWith(transportResult.trailer);
    });

    async function setup() {
      const def = await proto`
        service TestService {
          rpc TestClientStreamMethod(stream TestInput) returns (TestOutput);
        }

        message TestInput {
          string test = 1;
        }

        message TestOutput {
          string result = 1;
        }
      `;
      const TestInputSchema = def.getMessage<"TestInput", { test: string }>("TestInput");
      const TestOutputSchema = def.getMessage<"TestOutput", { result: string }>("TestOutput");
      const TestServiceSchema = def.getTsProtoService<{
        testClientStreamMethod: DescMethodClientStreaming<typeof TestInputSchema, typeof TestOutputSchema>;
      }>("TestService");

      return {
        TestInputSchema,
        TestOutputSchema,
        TestServiceSchema,
      };
    }
  });

  describe("bidi streaming method", () => {
    it("can create bidi streaming method", async () => {
      const { TestServiceSchema } = await setup();
      const results = Array.from({ length: 3 }, (_, i) => ({
        result: `result${i + 1}`,
      }));
      const transport = createTransport("stream", () => ({
        message: createAsyncIterable(results.map((item) => ({
          ...item,
        }))),
        header: new Headers(),
        trailer: new Headers(),
      }));
      const client = createServiceClient(TestServiceSchema, transport);

      const abortSignal = new AbortController().signal;
      const headers = {
        "x-test": "test",
      };
      const onHeader = jest.fn();
      const onTrailer = jest.fn();
      const input = Array.from({ length: 3 }, (_, i) => ({ test: `input${i + 1}` }));
      const methodsCallResult = client.testBiDiStreamMethod(createAsyncIterable(input), {
        signal: abortSignal,
        timeoutMs: 1000,
        headers,
        onHeader,
        onTrailer,
      });

      expect(await Array.fromAsync(methodsCallResult)).toEqual(results);
      expect(transport.stream).toHaveBeenCalledWith(
        TestServiceSchema.methods.testBiDiStreamMethod,
        expect.anything(),
        {
          signal: abortSignal,
          timeoutMs: 1000,
          headers,
          onHeader,
          onTrailer,
        },
      );

      expect(await Array.fromAsync(transport.stream.mock.lastCall?.at(1) as AsyncIterable<unknown>)).toEqual(
        input.map((item) => ({
          ...item,
        })),
      );
      const transportResult = (await transport.stream.mock.results.at(-1)?.value) as Awaited<ReturnType<typeof transport.stream>>;
      expect(onHeader).toHaveBeenCalledWith(transportResult.header);
      expect(onTrailer).toHaveBeenCalledWith(transportResult.trailer);
    });

    async function setup() {
      const def = await proto`
        service TestService {
          rpc TestBiDiStreamMethod(stream TestInput) returns (stream TestOutput);
        }

        message TestInput {
          string test = 1;
        }

        message TestOutput {
          string result = 1;
        }
      `;
      const TestInputSchema = def.getMessage<"TestInput", { test: string }>("TestInput");
      const TestOutputSchema = def.getMessage<"TestOutput", { result: string }>("TestOutput");
      const TestServiceSchema = def.getTsProtoService<{
        testBiDiStreamMethod: DescMethodBiDiStreaming<typeof TestInputSchema, typeof TestOutputSchema>;
      }>("TestService");

      return {
        TestInputSchema,
        TestOutputSchema,
        TestServiceSchema,
      };
    }
  });

  type Response<T extends "stream" | "unary"> = T extends "stream" ? StreamResponse<MessageDesc, MessageDesc> : UnaryResponse<MessageDesc, MessageDesc>;
  function createTransport<T extends "stream" | "unary">(responseType: T, createResponse: () => Pick<Response<T>, "message"> & Partial<Pick<Response<T>, "header" | "trailer">>) {
    const method = jest.fn(async (method: DescMethod) => {
      return {
        header: new Headers(),
        trailer: new Headers(),
        ...createResponse(),
        stream: responseType === "stream" as const,
        service: method.parent,
        method,
      };
    });

    return {
      requiresTypePatching: true,
      unary: notImplemented,
      stream: notImplemented,
      [responseType]: method,
    } as unknown as Omit<Transport, T> & Record<T, jest.MockedFunction<Transport[T]>>;
  }

  async function notImplemented(): Promise<never> {
    throw new Error("not implemented");
  }
});
