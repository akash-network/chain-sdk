import type {
  EncodeObject,
  GeneratedType,
  OfflineSigner,
} from "@cosmjs/proto-signing";
import {
  Registry,
} from "@cosmjs/proto-signing";
import type {
  HttpEndpoint,
  SigningStargateClientOptions,
} from "@cosmjs/stargate";
import {
  SigningStargateClient,
} from "@cosmjs/stargate";

import type { TxClient } from "../TxClient.ts";

const DEFAULT_FEE_AMOUNT = "2500";
const DEFAULT_GAS_MULTIPLIER = 1.3;

export function createStargateClient(options: StargateClientOptions): TxClient {
  const builtInTypes = options.builtInTypes?.map((type) => [type.typeUrl, type] as [string, GeneratedType]) || [];
  const registry = new Registry(builtInTypes);
  const createStargateClient = options.createClient ?? SigningStargateClient.connectWithSigner;

  let stargateClientPromise: Promise<SigningStargateClient> | undefined;
  const getStargateClient = () => stargateClientPromise ??= createStargateClient(
    options.baseUrl,
    options.signer,
    {
      ...options.stargateOptions,
      registry,
    },
  );

  const getAccount = options.getAccount ?? (() => getDefaultAccount(options.signer));
  const gasMultiplier = options.gasMultiplier ?? DEFAULT_GAS_MULTIPLIER;
  const preloadMessageTypes = (messages: EncodeObject[]) => {
    for (const message of messages) {
      if (registry.lookupType(message.typeUrl)) continue;
      const type = options.getMessageType(message.typeUrl);
      registry.register(message.typeUrl, type);
    }
    return messages;
  };

  return {
    async estimateFee(messages, memo) {
      const account = await getAccount(preloadMessageTypes(messages));
      const client = await getStargateClient();
      const gas = await client.simulate(account, messages, memo);
      return {
        amount: [
          {
            denom: "uakt",
            amount: options.defaultFeeAmount ?? DEFAULT_FEE_AMOUNT,
          },
        ],
        gas: Math.floor(gasMultiplier * gas).toString(),
        granter: account,
      };
    },
    async sign(messages, fee, memo) {
      const account = await getAccount(preloadMessageTypes(messages));
      const client = await getStargateClient();
      return client.sign(account, messages, fee, memo);
    },
    async broadcast(txRaw) {
      const TxRawType = registry.lookupType("/cosmos.tx.v1beta1.TxRaw");
      if (!TxRawType) {
        throw new Error("Cannot broadcast transaction: TxRaw type is not registered");
      }
      const client = await getStargateClient();
      return client.broadcastTx(
        TxRawType.encode(txRaw).finish(),
        options.stargateOptions?.broadcastTimeoutMs,
        options.stargateOptions?.broadcastPollIntervalMs,
      );
    },
  };
}

export interface StargateClientOptions {
  /**
   * Blockchain RPC endpoint
   */
  baseUrl: string;

  signer: OfflineSigner;

  /**
   * Gas multiplier
   * @default 1.3
   */
  gasMultiplier?: number;
  /**
   * @default "2500" uakt
   */
  defaultFeeAmount?: string;
  /**
   * @default returns the first account from the signer
   */
  getAccount?(messages: EncodeObject[]): Promise<string>;
  stargateOptions?: SigningStargateClientOptions;
  builtInTypes?: Array<GeneratedType & { typeUrl: string }>;
  getMessageType: (typeUrl: string) => GeneratedType;
  /**
   * @default `SigningStargateClient.connectWithSigner`
   */
  createClient?: (endpoint: string | HttpEndpoint, signer: OfflineSigner, options?: SigningStargateClientOptions) => Promise<SigningStargateClient>;
}

async function getDefaultAccount(signer: OfflineSigner) {
  const account = await signer.getAccounts();
  return account[0].address;
}
