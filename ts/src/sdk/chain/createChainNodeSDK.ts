import { createSDK as createCosmosSDK, serviceLoader as cosmosServiceLoader } from "../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK, serviceLoader as nodeServiceLoader } from "../../generated/createNodeSDK.ts";
import { TxRaw } from "../../generated/protos/cosmos/tx/v1beta1/tx.ts";
import { createMessageType } from "../client/createServiceLoader.ts";
import { createNoopTransport } from "../transport/createNoopTransport.ts";
import { createGrpcTransport } from "../transport/grpc/createGrpcTransport.ts";
import type { StargateClientOptions } from "../transport/tx/createStargateClient/createStargateClient.ts";
import { createStargateClient } from "../transport/tx/createStargateClient/createStargateClient.ts";
import { createTxTransport } from "../transport/tx/createTxTransport.ts";

export type { PayloadOf, ResponseOf } from "../types.ts";

export function createChainNodeSDK(options: ChainNodeSDKOptions) {
  const queryTransport = createGrpcTransport({
    baseUrl: options.query.baseUrl,
  });
  const getMessageType: StargateClientOptions["getMessageType"] = (typeUrl) => nodeServiceLoader.getLoadedType(typeUrl) || cosmosServiceLoader.getLoadedType(typeUrl);
  const txTransport = options.tx
    ? createTxTransport({
        getMessageType,
        client: createStargateClient({
          ...options.tx,
          getMessageType,
          builtInTypes: [
            createMessageType(TxRaw),
          ],
        }),
      })
    : createNoopTransport({
        unaryErrorMessage: `Unable to sign transaction. "tx" option is not provided during chain SDK creation`,
      });
  const nodeSDK = createNodeSDK(queryTransport, txTransport);
  const cosmosSDK = createCosmosSDK(queryTransport, txTransport);
  return { ...nodeSDK, ...cosmosSDK };
}

export interface ChainNodeSDKOptions {
  query: {
    baseUrl: string;
  };
  tx?: Omit<StargateClientOptions, "getMessageType" | "builtInTypes">;
}
