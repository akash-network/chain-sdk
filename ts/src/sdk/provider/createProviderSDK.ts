import { createSDK } from "../../generated/createProviderSDK.ts";
import type { PickByPath } from "../../utils/types.ts";
import type { GrpcTransportOptions } from "../transport/grpc/createGrpcTransport.ts";
import { createGrpcTransport } from "../transport/grpc/createGrpcTransport.ts";

export type { PayloadOf, ResponseOf } from "../types.ts";

type ProviderSDK = PickByPath<ReturnType<typeof createSDK>, "akash.provider.v1">;

export function createProviderSDK(options: ProviderSDKOptions): ProviderSDK {
  const certificateOptions = options.authentication?.type === "mtls"
    ? {
        cert: options.authentication?.cert,
        key: options.authentication?.key,
      }
    : null;

  return createSDK(
    createGrpcTransport({
      ...options.transportOptions,
      baseUrl: options.baseUrl,
      nodeOptions: {
        ...certificateOptions,
        rejectUnauthorized: false,
      },
    }),
  );
}

export interface ProviderSDKOptions {
  /**
   * Provider gRPC endpoint
   */
  baseUrl: string;

  /**
   * Authentication options
   */
  authentication?: {
    type: "mtls";
    cert: string;
    key: string;
  };

  /**
   * Options for the gRPC transport
   */
  transportOptions?: Pick<GrpcTransportOptions, "pingIdleConnection" | "pingIntervalMs" | "pingTimeoutMs" | "idleConnectionTimeoutMs" | "defaultTimeoutMs">;
}
