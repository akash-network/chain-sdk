import { createSDK as createCosmosSDK, serviceLoader as cosmosServiceLoader } from "../../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK, serviceLoader as nodeServiceLoader } from "../../../generated/createNodeSDK.ts";
import { createGrpcGatewayTransport } from "../../transport/grpc-gateway/createGrpcGatewayTransport.ts";
import { createTxTransport } from "../../transport/tx/createTxTransport.ts";
import type { TxClient } from "../../transport/tx/TxClient.ts";

export type { PayloadOf, ResponseOf } from "../../types.ts";

export function createChainNodeSDK(options: ChainNodeSDKOptions) {
  const queryTransport = createGrpcGatewayTransport({
    baseUrl: options.query.baseUrl,
  });
  const getMessageType = (typeUrl: string) => nodeServiceLoader.getLoadedType(typeUrl) || cosmosServiceLoader.getLoadedType(typeUrl);
  const txTransport = createTxTransport({
    getMessageType,
    client: options.tx.signer,
  });
  const nodeSDK = createNodeSDK(queryTransport, txTransport);
  const cosmosSDK = createCosmosSDK(queryTransport, txTransport);
  return { ...nodeSDK, ...cosmosSDK };
}

export interface ChainNodeSDKOptions {
  query: {
    baseUrl: string;
  };
  tx: {
    signer: TxClient;
  };
}
