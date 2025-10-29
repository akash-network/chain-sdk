import type {
  DirectSecp256k1HdWalletOptions,
  EncodeObject,
  GeneratedType,
  OfflineSigner,
} from "@cosmjs/proto-signing";
import {
  DirectSecp256k1HdWallet,
  Registry,
} from "@cosmjs/proto-signing";
import type {
  HttpEndpoint,
  SigningStargateClientOptions,
} from "@cosmjs/stargate";
import {
  calculateFee,
  GasPrice,
  SigningStargateClient,
} from "@cosmjs/stargate";

import type { TxClient } from "../TxClient.ts";

const DEFAULT_AVERAGE_GAS_PRICE = "0.025uakt";
const DEFAULT_GAS_MULTIPLIER = 1.3;

export function createGenericStargateClient(options: WithSigner<BaseGenericStargateClientOptions>): StargateTxClient {
  const builtInTypes = options.builtInTypes?.map((type) => [type.typeUrl, type] as [string, GeneratedType]) || [];
  const registry = new Registry(builtInTypes);
  const createStargateClient = options.createClient ?? SigningStargateClient.connectWithSigner;

  let offlineSignerPromise: Promise<OfflineSigner> | undefined;
  const getOfflineSigner = () => offlineSignerPromise ??= createOfflineSigner(options);

  let stargateClientPromise: Promise<SigningStargateClient> | undefined;
  const getStargateClient = () => stargateClientPromise ??= getOfflineSigner().then((signer) => createStargateClient(
    options.baseUrl,
    signer,
    {
      ...options.stargateOptions,
      registry,
    },
  ));

  const getAccount = (messsages: EncodeObject[]) => getOfflineSigner().then((signer) => (options.getAccount ?? getDefaultAccount)(signer, messsages));
  const gasMultiplier = options.gasMultiplier ?? DEFAULT_GAS_MULTIPLIER;
  const preloadMessageTypes = (messages: EncodeObject[]) => {
    for (const message of messages) {
      if (registry.lookupType(message.typeUrl)) continue;
      const type = options.getMessageType(message.typeUrl);
      if (!type) {
        throw new Error(`Cannot find message type ${message.typeUrl} in type registry. Probably it's not loaded yet.`);
      }
      registry.register(message.typeUrl, type);
    }
    return messages;
  };
  const gasPrice = GasPrice.fromString(options.defaultGasPrice ?? DEFAULT_AVERAGE_GAS_PRICE);

  return {
    async estimateFee(messages, memo) {
      const account = await getAccount(preloadMessageTypes(messages));
      const client = await getStargateClient();
      const estimatedGas = await client.simulate(account, messages, memo);
      const minGas = Math.floor(gasMultiplier * estimatedGas);
      const fee = calculateFee(minGas, gasPrice);

      return {
        ...fee,
        granter: account,
      };
    },
    async sign(messages, fee, memo) {
      const account = await getAccount(preloadMessageTypes(messages));
      const client = await getStargateClient();
      return client.sign(account, messages, fee, memo);
    },
    async broadcast(txRaw) {
      const txTypeUrl = "/cosmos.tx.v1beta1.TxRaw";
      const TxRawType = registry.lookupType(txTypeUrl) || options.getMessageType(txTypeUrl);
      if (!TxRawType) {
        throw new Error("Cannot broadcast transaction: TxRaw type is not registered in transaction client");
      }
      const client = await getStargateClient();
      return client.broadcastTx(
        TxRawType.encode(txRaw).finish(),
        options.stargateOptions?.broadcastTimeoutMs,
        options.stargateOptions?.broadcastPollIntervalMs,
      );
    },
    async disconnect() {
      if (!stargateClientPromise) return;

      const client = await stargateClientPromise;
      client.disconnect();
      stargateClientPromise = undefined;
      offlineSignerPromise = undefined;
    },
  };
}

export interface StargateTxClient extends TxClient {
  disconnect(): Promise<void>;
}

export type WithSigner<T> = T & (
  | {
    /**
       * Signer to use for transactions signing
       */
    signer: OfflineSigner;
  }
  | {
    signer?: never;
    /**
       * Uses the mnemonic to create a `DirectSecp256k1HdWallet` to use for transactions signing
       */
    signerMnemonic: string;
    /**
       * Options to pass to the `DirectSecp256k1HdWallet`
       */
    signerOptions?: Partial<Omit<DirectSecp256k1HdWalletOptions, "prefix">>;
  }
);

export interface BaseGenericStargateClientOptions {
  /**
   * Blockchain RPC endpoint
   */
  baseUrl: string;
  /**
   * Gas multiplier
   * @default 1.3
   */
  gasMultiplier?: number;
  /**
   * @default "0.025uakt"
   */
  defaultGasPrice?: string;
  /**
   * Retrieves the account to use for transactions
   * @default returns the first account from the signer
   */
  getAccount?(signer: OfflineSigner, messages: EncodeObject[]): Promise<string>;
  stargateOptions?: Omit<SigningStargateClientOptions, "registry">;
  /**
   * Additional protobuf message types to register with the transaction transport
   */
  builtInTypes?: Array<GeneratedType & { typeUrl: string }>;
  getMessageType: (typeUrl: string) => GeneratedType | undefined;
  /**
   * Allows to use a custom Stargate client implementation.
   * @default `SigningStargateClient.connectWithSigner`
   */
  createClient?: (endpoint: string | HttpEndpoint, signer: OfflineSigner, options?: SigningStargateClientOptions) => Promise<SigningStargateClient>;
}

async function getDefaultAccount(signer: OfflineSigner) {
  const accounts = await signer.getAccounts();
  if (accounts.length === 0) {
    throw new Error("provided offline signer has no accounts");
  }
  return accounts[0].address;
}

function createOfflineSigner(options: WithSigner<BaseGenericStargateClientOptions>) {
  if ("signer" in options && options.signer) return Promise.resolve(options.signer);

  return DirectSecp256k1HdWallet.fromMnemonic(options.signerMnemonic, {
    ...options.signerOptions,
    prefix: "akash",
  });
}
