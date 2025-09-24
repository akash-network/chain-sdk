import { createGrpcGatewayTransport } from "../../transport/grpc-gateway/createGrpcGatewayTransport.ts";
import { serviceLoader as cosmosServiceLoader, createSDK as createCosmosSDK } from "../../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK, serviceLoader as nodeServiceLoader } from "../../../generated/createNodeSDK.ts";
import { createGrpcWebTransport } from "../../transport/grpc-web/createGrpcWebTransport.ts";
import { createTxTransport } from "../../transport/tx/createTxTransport.ts";
import { TxClient } from "../../transport/tx/TxClient.ts";

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
  }
}
