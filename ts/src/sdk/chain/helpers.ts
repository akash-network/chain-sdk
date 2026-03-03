import { getMetadata, type SDKMethod } from "../client/sdkMetadata.ts";
import type { TxClient, TxSignAndBroadcastOptions } from "../transport/tx/TxClient.ts";

export function msg<T extends SDKMethod>(tx: T, ...args: [Parameters<T>[0]]): TxMessage<T>;
export function msg<T extends SDKMethod>(tx: T, data?: Parameters<T>[0]): TxMessage<T> {
  return {
    tx,
    data,
  };
}

/**
 * Representation of a transaction message, containing the SDK transaction method and its input.
 */
export type TxMessage<T extends SDKMethod = SDKMethod> = {
  tx: T;
  data: Parameters<T>[0];
};

export const SIGNER_KEY = Symbol("signer");

export async function transaction(sdk: { [SIGNER_KEY]?: TxClient }, messages: TxMessage[], options?: TxSignAndBroadcastOptions) {
  const signer = sdk[SIGNER_KEY];

  if (!signer) {
    throw new Error("Transaction signer is not available in the SDK instance");
  }

  const encodedObjects = await Promise.all(messages.map(async (msg) => {
    const meta = getMetadata(msg.tx);
    if (!meta) {
      throw new Error(`No metadata found for method SDK method ${msg.tx.name}`);
    }

    const serviceDesc = await meta.serviceLoader.loadAt(meta.path[0]);
    const method = serviceDesc.methods[meta.path[1]];

    if (!method) {
      throw new Error(`No method found at path [${meta.path[0]}, ${meta.path[1]}] in "${serviceDesc.typeName}"`);
    }

    return {
      typeUrl: `/${method.input.$type}`,
      value: method.input.fromPartial(msg.data ?? {}),
    };
  }));

  return await signer.signAndBroadcast(encodedObjects, options);
}
