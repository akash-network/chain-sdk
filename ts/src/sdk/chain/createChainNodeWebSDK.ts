import { createSDK as createCosmosSDK, serviceLoader as cosmosServiceLoader } from "../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK, serviceLoader as nodeServiceLoader } from "../../generated/createNodeSDK.ts";
import { patches as cosmosPatches } from "../../generated/patches/cosmosCustomTypePatches.ts";
import { patches as nodePatches } from "../../generated/patches/nodeCustomTypePatches.ts";
import { createNoopTransport } from "../transport/createNoopTransport.ts";
import { createGrpcGatewayTransport } from "../transport/grpc-gateway/createGrpcGatewayTransport.ts";
import { createTxTransport } from "../transport/tx/createTxTransport.ts";
import type { TxClient } from "../transport/tx/TxClient.ts";

export type { PayloadOf, ResponseOf } from "../types.ts";

export function createChainNodeWebSDK(options: ChainNodeWebSDKOptions) {
  const queryTransport = createGrpcGatewayTransport({
    baseUrl: options.query.baseUrl,
  });
  const getMessageType = (typeUrl: string) => nodeServiceLoader.getLoadedType(typeUrl) || cosmosServiceLoader.getLoadedType(typeUrl);
  const txTransport = options.tx
    ? createTxTransport({
        getMessageType,
        client: options.tx.signer,
      })
    : createNoopTransport({
        unaryErrorMessage: `Unable to sign transaction. "tx" option is not provided during chain SDK creation`,
      });
  const nodeSDK = createNodeSDK(queryTransport, txTransport, {
    clientOptions: { typePatches: { ...cosmosPatches, ...nodePatches } },
  });
  const cosmosSDK = createCosmosSDK(queryTransport, txTransport, {
    clientOptions: { typePatches: cosmosPatches },
  });
  return { ...nodeSDK, ...cosmosSDK };
}

export interface ChainNodeWebSDKOptions {
  query: {
    /**
     * Blockchain gRPC gateway endpoint (also known as REST endpoint)
     */
    baseUrl: string;
  };
  tx?: {
    signer: TxClient;
  };
}
