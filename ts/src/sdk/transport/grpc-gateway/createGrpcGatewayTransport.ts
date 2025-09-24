import { Code, ConnectError, createContextValues } from "@connectrpc/connect";
import { requestHeaderWithCompression } from "@connectrpc/connect/protocol-grpc";
import { type GrpcTransportOptions as ConnectGrpcTransportOptions } from "@connectrpc/connect-node";

import type { MessageDesc, MessageInitShape, MessageShape, MethodDesc } from "../../client/types.ts";
import { runUnaryCall } from "../runCall.ts";
import { coerceTimeoutMs } from "../transportUtils.ts";
import type { CallOptions, StreamResponse, Transport, UnaryRequest, UnaryResponse } from "../types.ts";

export type GrpcGatewayCallOptions = Omit<CallOptions, "onHeader" | "onTrailer"> & {
  header?: HeadersInit;
};
export interface GrpcGatewayTransportOptions extends Omit<ConnectGrpcTransportOptions, "useBinaryFormat"> {
  fetch?: typeof globalThis.fetch;
}

export function createGrpcGatewayTransport(options: GrpcGatewayTransportOptions): Transport<GrpcGatewayCallOptions> {
  const useBinaryFormat = false;

  return {
    async unary<I extends MessageDesc, O extends MessageDesc>(
      method: MethodDesc<"unary", I, O>,
      message: MessageInitShape<I>,
      callOptions?: GrpcGatewayCallOptions,
    ): Promise<UnaryResponse<I, O>> {
      const timeoutMs = coerceTimeoutMs(callOptions?.timeoutMs, options.defaultTimeoutMs);

      if (!method.httpPath) {
        throw new ConnectError(`Service ${method.parent.typeName} method "${method.name}" doesn't not support http transport`, Code.InvalidArgument);
      }

      const headers = requestHeaderWithCompression(
        useBinaryFormat,
        timeoutMs,
        callOptions?.header,
        options.acceptCompression ?? [],
        options.sendCompression || null,
      );

      headers.set("Content-type", "application/json");

      return await runUnaryCall<I, O>({
        interceptors: options.interceptors,
        signal: callOptions?.signal,
        timeoutMs,
        req: {
          service: method.parent,
          stream: false,
          method,
          requestMethod: method.httpMethod || "GET",
          url: method.httpPath.replace(/\{[^}]+\}/g, (interpolation) => {
            const data = message as Record<string, unknown> | undefined;
            const key = interpolation.slice(1, -1).trim();
            if (!data || !Object.hasOwn(data, key)) {
              throw new ConnectError(`Cannot construct url for ${method.parent.typeName}.${method.name}: "${key}" is not specified in message`, Code.InvalidArgument);
            }
            return String(data[key]);
          }),
          header: headers,
          contextValues: callOptions?.contextValues ?? createContextValues(),
          message,
        },
        next: async (req: UnaryRequest<I, O>): Promise<UnaryResponse<I, O>> => {
          const fetch = options.fetch ?? globalThis.fetch;
          const url = new URL(options.baseUrl + req.url);

          if (req.requestMethod === "GET") {
            serializeParams(method.input.toJSON(req.message) as Record<string, unknown>, url.searchParams);
          }

          const response = await fetch(url, {
            method: req.requestMethod,
            headers: req.header,
            signal: req.signal,
            body: req.requestMethod === "POST" || req.requestMethod === "PUT" || req.requestMethod === "PATCH"
              ? JSON.stringify(method.input.toJSON(req.message))
              : undefined,
          });
          const body = await response.json();
          return {
            stream: false,
            method,
            header: response.headers,
            trailer: new Headers(),
            message: method.output.fromJSON(body) as MessageShape<O>,
          } satisfies UnaryResponse<I, O>;
        },
      });
    },
    async stream<I extends MessageDesc, O extends MessageDesc>(): Promise<StreamResponse<I, O>> {
      throw new ConnectError(`HTTP transport doesn't support streaming`, Code.Unimplemented);
    },
  };
}

function serializeParams(message: Record<string, unknown>, params: URLSearchParams, prefix = '') {
  Object.keys(message).forEach((key) => {
    const name = prefix + key;
    const value = message[key];
    if (value === null || value === undefined) return;
    if (value.constructor === Object) {
      serializeParams(value as Record<string, unknown>, params, `${name}.`);
      return;
    }

    if (value instanceof Uint8Array) {
      params.append(name, btoa(value as any));
      return;
    }

    params.append(name, String(value));
  });
  return params;
}
