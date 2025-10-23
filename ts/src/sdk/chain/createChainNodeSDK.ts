import type { GeneratedType } from "@cosmjs/proto-signing";

import { createSDK as createCosmosSDK, serviceLoader as cosmosServiceLoader } from "../../generated/createCosmosSDK.ts";
import { createSDK as createNodeSDK, serviceLoader as nodeServiceLoader } from "../../generated/createNodeSDK.ts";
import { patches as cosmosPatches } from "../../generated/patches/cosmosCustomTypePatches.ts";
import { patches as nodePatches } from "../../generated/patches/nodeCustomTypePatches.ts";
import { TxRaw } from "../../generated/protos/cosmos/tx/v1beta1/tx.ts";
import { createMessageType } from "../client/createServiceLoader.ts";
import type { MessageDesc } from "../client/types.ts";
import { createNoopTransport } from "../transport/createNoopTransport.ts";
import type { GrpcTransportOptions } from "../transport/grpc/createGrpcTransport.ts";
import { createGrpcTransport } from "../transport/grpc/createGrpcTransport.ts";
import type { StargateClientOptions } from "../transport/tx/createStargateClient/createStargateClient.ts";
import { createStargateClient } from "../transport/tx/createStargateClient/createStargateClient.ts";
import { createTxTransport } from "../transport/tx/createTxTransport.ts";
import type { Transport, TxCallOptions } from "../transport/types.ts";

export type { PayloadOf, ResponseOf } from "../types.ts";

export function createChainNodeSDK(options: ChainNodeSDKOptions) {
  const queryTransport = createGrpcTransport({
    ...options.query.transportOptions,
    baseUrl: options.query.baseUrl,
  });
  let txTransport: Transport<TxCallOptions>;

  if (options.tx) {
    const { builtInTypes, ...txOptions } = options.tx;
    const defaultRegistryTypes = [
      createMessageType(TxRaw),
      ...(builtInTypes ?? []).map(createMessageType),
    ].reduce<Record<string, GeneratedType>>((acc, type) => {
      acc[type.typeUrl] = type;
      return acc;
    }, {});
    const getMessageType: StargateClientOptions["getMessageType"] = (typeUrl) => nodeServiceLoader.getLoadedType(typeUrl) || cosmosServiceLoader.getLoadedType(typeUrl) || defaultRegistryTypes[typeUrl];
    txTransport = createTxTransport({
      getMessageType,
      client: createStargateClient({
        ...txOptions,
        getMessageType,
      }),
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
  tx?: Omit<StargateClientOptions, "getMessageType" | "builtInTypes"> & {
    /**
     * Additional protobuf message types to register with the transaction transport
     */
    builtInTypes?: MessageDesc[];
  };
}
