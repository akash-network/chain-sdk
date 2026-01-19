import type { Interceptor } from "@connectrpc/connect";
import { createContextValues } from "@connectrpc/connect";
import type { EncodeObject, GeneratedType } from "@cosmjs/proto-signing";

import type { MessageDesc, MessageInitShape, MessageShape, MethodDesc } from "../../client/types.ts";
import { runUnaryCall } from "../runCall.ts";
import { TransportError } from "../TransportError.ts";
import type { Transport, TxCallOptions, UnaryRequest, UnaryResponse } from "../types.ts";
import { BatchQueue } from "./BatchQueue/BatchQueue.ts";
import type { StdFee, TxClient } from "./TxClient.ts";
import { TxError } from "./TxError.ts";

const DEFAULT_QUEUE_ID = "default";
const MAX_MEMO_LENGTH = 256;

export function createTxTransport(options: TransactionTransportOptions): Transport<TxCallOptions> {
  const maxMessagesInBatchedTx = Number(options.maxMessagesInBatchedTx ?? 1);
  return new TxTransport({
    ...options,
    maxMessagesInBatchedTx: Number.isNaN(maxMessagesInBatchedTx) || maxMessagesInBatchedTx < 1 ? 1 : Math.min(maxMessagesInBatchedTx, 10),
  });
}

type TransportOptionsWithDefaults = Omit<TransactionTransportOptions, "maxMessagesInBatchedTx"> & Required<Pick<TransactionTransportOptions, "maxMessagesInBatchedTx">>;
class TxTransport implements Transport<TxCallOptions> {
  readonly requiresTypePatching = true;
  readonly #options: TransportOptionsWithDefaults;
  readonly #txQueues = new Map<string, BatchQueue<TxItem<MessageDesc, MessageDesc>, void>>();
  #sendTxTail = Promise.resolve();

  constructor(options: TransportOptionsWithDefaults) {
    this.#options = options;
  }

  async unary<I extends MessageDesc, O extends MessageDesc>(
    method: MethodDesc<"unary", I, O>,
    input: MessageInitShape<I>,
    callOptions?: TxCallOptions,
  ): Promise<UnaryResponse<I, O>> {
    return runUnaryCall({
      interceptors: this.#options.interceptors,
      req: {
        stream: false,
        service: method.parent,
        method,
        contextValues: callOptions?.contextValues ?? createContextValues(),
        message: input,
        requestMethod: "POST",
        url: "",
      },
      next: (req) => {
        return new Promise((resolve, reject) => {
          const txItem = {
            req,
            callOptions,
            resolve,
            reject,
          };

          if (callOptions?.fee?.amount?.length || callOptions?.fee?.gas) {
            return this.#enqueueTxBatchSign([txItem]);
          }

          const batchId = callOptions?.fee?.payer || callOptions?.fee?.granter
            ? `${callOptions?.fee?.payer || "<default>"}:${callOptions?.fee?.granter || "<default>"}:${callOptions?.gasMultiplier ?? "<default>"}`
            : DEFAULT_QUEUE_ID;
          this.#getBatchQueue(batchId).add(txItem as unknown as TxItem<MessageDesc, MessageDesc>);
        });
      },
    });
  }

  async stream(): Promise<never> {
    throw new TransportError(`Transaction transport doesn't support streaming`, TransportError.Code.Unimplemented);
  }

  #getBatchQueue(key: string) {
    let queue = this.#txQueues.get(key);
    if (!queue) {
      queue = new BatchQueue({
        scheduleFn: queueMicrotask,
        maxBatchSize: this.#options.maxMessagesInBatchedTx,
        onFlush: async (items) => {
          await this.#enqueueTxBatchSign(items);
          if (key !== DEFAULT_QUEUE_ID && this.#txQueues.get(key)?.size === 0) {
            this.#txQueues.delete(key);
          }
        },
      });
      this.#txQueues.set(key, queue);
    }
    return queue;
  }

  #enqueueSend<T>(fn: () => Promise<T>): Promise<T> {
    const next = this.#sendTxTail.then(fn, fn);
    this.#sendTxTail = new Promise((resolve) => next.finally(resolve));
    return next;
  }

  async #enqueueTxBatchSign<I extends MessageDesc, O extends MessageDesc>(items: TxItem<I, O>[]) {
    return this.#enqueueSend(async () => this.#signTxBatch(items));
  }

  async #signTxBatch<I extends MessageDesc, O extends MessageDesc>(items: TxItem<I, O>[]) {
    try {
      let memo = "";
      const messages: EncodeObject[] = [];
      let hasAfterSign = false;
      let hasAfterBroadcast = false;

      items.forEach((item, index) => {
        messages.push({
          typeUrl: `/${item.req.method.input.$type}`,
          value: item.req.message,
        });
        const itemMemo = item.callOptions?.memo ?? `akash: ${item.req.method.name}`;
        memo += (index === 0 ? "" : ". ") + itemMemo;
        hasAfterSign = hasAfterSign || !!item.callOptions?.afterSign;
        hasAfterBroadcast = hasAfterBroadcast || !!item.callOptions?.afterBroadcast;
      });

      if (memo.length > MAX_MEMO_LENGTH) {
        memo = memo.slice(0, (MAX_MEMO_LENGTH - 1 - 3)) + "...";
      }

      let fee: StdFee;
      // payer, granter and gasMultiplier are the same for all items in batch, so pick values from the first item
      const providedFee = items[0].callOptions?.fee;
      const providedGasMultiplier = items[0].callOptions?.gasMultiplier;
      if (!providedFee?.amount?.length || !providedFee?.gas) {
        const estimatedFee = await this.#options.client.estimateFee(messages, "stargate", memo, providedGasMultiplier);
        fee = providedFee ? { ...estimatedFee, ...providedFee } : estimatedFee;
      } else {
        fee = providedFee as StdFee;
      }

      const txRaw = await this.#options.client.sign(messages, fee, memo);
      if (hasAfterSign) {
        items.forEach((item) => {
          try {
            item.callOptions?.afterSign?.(txRaw);
          } catch (error) {
            item.reject(error);
          }
        });
      }
      const txResponse = await this.#options.client.broadcast(txRaw);
      if (hasAfterBroadcast) {
        items.forEach((item) => {
          try {
            item.callOptions?.afterBroadcast?.(txResponse);
          } catch (error) {
            item.reject(error);
          }
        });
      }

      if (txResponse.code !== 0) {
        throw new TxError(`Transaction failed with code ${txResponse.code}`, txResponse);
      }

      items.forEach((item, index) => {
        const response = txResponse.msgResponses?.[index];
        let responseMessage: MessageShape<O>;
        if (response) {
          const MessageType = this.#options.getMessageType(response.typeUrl);
          if (!MessageType) {
            throw new Error(`Cannot find message type ${response.typeUrl} in type registry. `
              + `If you use cosmos.authz.v1beta1.exec(), then provide custom message types to TxClient.`);
          }
          responseMessage = MessageType.decode(response.value);
        } else {
          responseMessage = {} as MessageShape<O>;
        }
        item.resolve({
          stream: false,
          header: new Headers(),
          trailer: new Headers(),
          message: responseMessage,
          method: item.req.method,
        });
      });
    } catch (error) {
      items.forEach((item) => item.reject(error));
    }
  }
}

export interface TransactionTransportOptions {
  client: TxClient;
  getMessageType: (typeUrl: string) => GeneratedType | undefined;
  interceptors?: Interceptor[];
  maxMessagesInBatchedTx?: number;
}

interface TxItem<I extends MessageDesc, O extends MessageDesc> {
  req: UnaryRequest<I, O>;
  callOptions?: TxCallOptions;
  resolve: (value: UnaryResponse<I, O>) => void;
  reject: (error: unknown) => void;
}
