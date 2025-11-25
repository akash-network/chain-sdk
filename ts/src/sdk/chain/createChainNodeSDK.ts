import { createSDK as createCosmosSDK } from "../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK } from "../../generated/createNodeSDK.ts";
import { patches as cosmosPatches } from "../../generated/patches/cosmosCustomTypePatches.ts";
import { patches as nodePatches } from "../../generated/patches/nodeCustomTypePatches.ts";
import { getMessageType } from "../getMessageType.ts";
import { createNoopTransport } from "../transport/createNoopTransport.ts";
import type { GrpcTransportOptions } from "../transport/grpc/createGrpcTransport.ts";
import { createGrpcTransport } from "../transport/grpc/createGrpcTransport.ts";
import { createTxTransport } from "../transport/tx/createTxTransport.ts";
import type { TxClient } from "../transport/tx/TxClient.ts";
import type { Transport, TxCallOptions } from "../transport/types.ts";

export type { PayloadOf, ResponseOf } from "../types.ts";

export function createChainNodeSDK(options: ChainNodeSDKOptions) {
  const queryTransport = createGrpcTransport({
    ...options.query.transportOptions,
    baseUrl: options.query.baseUrl,
  });
  let txTransport: Transport<TxCallOptions>;

  if (options.tx) {
    txTransport = createTxTransport({
      getMessageType,
      client: options.tx.signer,
    });
  } else {
    txTransport = createNoopTransport({
      unaryErrorMessage: `Unable to sign transaction. "tx" option is not provided during chain SDK creation`,
    });
  }
  const nodeSDK = createNodeSDK(queryTransport, txTransport, {
    clientOptions: { typePatches: { ...cosmosPatches, ...nodePatches } },
  });
  const cosmosSDK = createCosmosSDK(queryTransport, txTransport, {
    clientOptions: { typePatches: cosmosPatches },
  });
  return { ...nodeSDK, ...cosmosSDK };
}

export interface ChainNodeSDKOptions {
  query: {
    /**
     * Blockchain gRPC endpoint
     */
    baseUrl: string;

    /**
     * Options for the gRPC transport
     */
    transportOptions?: Pick<GrpcTransportOptions, "pingIdleConnection" | "pingIntervalMs" | "pingTimeoutMs" | "idleConnectionTimeoutMs" | "defaultTimeoutMs">;

  };
  tx?: {
    signer: TxClient;
  };
}
