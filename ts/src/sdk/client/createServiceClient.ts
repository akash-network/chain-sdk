import type { CallOptions, Transport } from "../transport/types.ts";
import { createAsyncIterable, handleStreamResponse } from "./stream.ts";
import type { MessageDesc, MessageInitShape, MessageShape, MethodDesc, ServiceDesc } from "./types.ts";

export type Client<Desc extends ServiceDesc, TCallOptions> = {
  [P in keyof Desc["methods"]]:
  Desc["methods"][P] extends MethodDesc<"server_streaming"> ? (input: MessageInitShape<Desc["methods"][P]["input"]>, options?: TCallOptions) => AsyncIterable<MessageShape<Desc["methods"][P]["output"]>>
    : Desc["methods"][P] extends MethodDesc<"client_streaming"> ? (input: AsyncIterable<MessageInitShape<Desc["methods"][P]["input"]>>, options?: TCallOptions) => Promise<MessageShape<Desc["methods"][P]["output"]>>
      : Desc["methods"][P] extends MethodDesc<"bidi_streaming"> ? (input: AsyncIterable<MessageInitShape<Desc["methods"][P]["input"]>>, options?: TCallOptions) => AsyncIterable<MessageShape<Desc["methods"][P]["output"]>>
        : Desc["methods"][P] extends MethodDesc<"unary"> | Omit<MethodDesc<"unary">, "kind"> ? (input: MessageInitShape<Desc["methods"][P]["input"]>, options?: TCallOptions) => Promise<MessageShape<Desc["methods"][P]["output"]>>
          : never;
};

export function createServiceClient<TSchema extends ServiceDesc, TCallOptions>(
  service: TSchema,
  transport: Transport<TCallOptions>,
): Client<TSchema, TCallOptions> {
  const client: Record<string, ReturnType<typeof createMethod>> = {};
  const methodNames = Object.keys(service.methods);
  for (let i = 0; i < methodNames.length; i++) {
    const methodDesc = service.methods[methodNames[i]];
    client[methodNames[i]] = createMethod(methodDesc as MethodDesc, transport);
  }

  return client as Client<TSchema, TCallOptions>;
}

function createMethod(methodDesc: MethodDesc, transport: Transport) {
  switch (methodDesc.kind) {
    case "server_streaming":
      return createServerStreamingFn(transport, methodDesc as MethodDesc<"server_streaming", MessageDesc, MessageDesc>);
    case "client_streaming":
      return createClientStreamingFn(transport, methodDesc as MethodDesc<"client_streaming", MessageDesc, MessageDesc>);
    case "bidi_streaming":
      return createBiDiStreamingFn(transport, methodDesc as MethodDesc<"bidi_streaming", MessageDesc, MessageDesc>);
    default:
      return createUnaryFn(transport, methodDesc as MethodDesc<"unary", MessageDesc, MessageDesc>);
  }
}

type UnaryFn<I extends MessageDesc<unknown>, O extends MessageDesc<unknown>> = (
  input: MessageInitShape<I>,
  options?: CallOptions,
) => Promise<MessageShape<O>>;

function createUnaryFn<I extends MessageDesc<unknown>, O extends MessageDesc<unknown>>(
  transport: Transport,
  method: MethodDesc<"unary", I, O>,
): UnaryFn<I, O> {
  return async (input, options) => {
    const response = await transport.unary(
      method,
      input,
      options,
    );
    options?.onHeader?.(response.header);
    options?.onTrailer?.(response.trailer);

    return response.message;
  };
}

type ServerStreamingFn<I extends MessageDesc, O extends MessageDesc> = (
  input: MessageInitShape<I>,
  options?: CallOptions,
) => AsyncIterable<MessageShape<O>>;

function createServerStreamingFn<
  I extends MessageDesc,
  O extends MessageDesc,
>(
  transport: Transport,
  method: MethodDesc<"server_streaming", I, O>,
): ServerStreamingFn<I, O> {
  return (input, options) => {
    return handleStreamResponse(
      method,
      transport.stream(
        method,
        createAsyncIterable([input]),
        options,
      ),
      options,
    );
  };
}

type ClientStreamingFn<I extends MessageDesc, O extends MessageDesc> = (
  input: AsyncIterable<MessageInitShape<I>>,
  options?: CallOptions,
) => Promise<MessageShape<O>>;

function createClientStreamingFn<
  I extends MessageDesc,
  O extends MessageDesc,
>(
  transport: Transport,
  method: MethodDesc<"client_streaming", I, O>,
): ClientStreamingFn<I, O> {
  return async (input, options) => {
    const response = await transport.stream(
      method,
      input,
      options,
    );
    options?.onHeader?.(response.header);
    let singleMessage: MessageShape<O> | undefined;
    let count = 0;
    for await (const message of response.message) {
      singleMessage = message;
      count++;
    }
    if (!singleMessage) {
      throw new Error("akash sdk protocol error: missing response message");
    }
    if (count > 1) {
      throw new Error("akash sdk protocol error: received extra messages for client streaming method");
    }
    options?.onTrailer?.(response.trailer);
    return singleMessage;
  };
}

type BiDiStreamingFn<I extends MessageDesc, O extends MessageDesc> = (
  input: AsyncIterable<MessageInitShape<I>>,
  options?: CallOptions,
) => AsyncIterable<MessageShape<O>>;

function createBiDiStreamingFn<
  I extends MessageDesc,
  O extends MessageDesc,
>(
  transport: Transport,
  method: MethodDesc<"bidi_streaming", I, O>,
): BiDiStreamingFn<I, O> {
  return (input, options) => {
    return handleStreamResponse(
      method,
      transport.stream(
        method,
        input,
        options,
      ),
      options,
    );
  };
}
