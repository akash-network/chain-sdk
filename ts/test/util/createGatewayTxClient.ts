/**
 * Gateway-based TxClient for Node.js environments.
 * 
 * Uses HTTP Gateway REST endpoints (/cosmos/tx/v1beta1/*) instead of Tendermint RPC.
 * This is the same protocol browsers use when they can't access Tendermint RPC directly.
 */
import type { EncodeObject, OfflineSigner } from "@cosmjs/proto-signing";
import { makeSignDoc } from "@cosmjs/proto-signing";
import type { DeliverTxResponse, StdFee } from "@cosmjs/stargate";
import { calculateFee, GasPrice } from "@cosmjs/stargate";
import { BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";

import type { TxClient, TxRaw } from "../../src/sdk/transport/tx/TxClient.ts";
import { TxRaw as TxRawType, TxBody, AuthInfo, SignerInfo, Fee } from "../../src/generated/protos/cosmos/tx/v1beta1/tx.ts";
import { SignMode } from "../../src/generated/protos/cosmos/tx/signing/v1beta1/signing.ts";
import { Any } from "../../src/generated/protos/google/protobuf/any.ts";
import { Coin } from "../../src/generated/protos/cosmos/base/v1beta1/coin.ts";

const DEFAULT_AVERAGE_GAS_PRICE = "0.025uakt";
const DEFAULT_GAS_MULTIPLIER = 1.3;

export interface GatewayTxClientOptions {
  gatewayUrl: string;
  signer: OfflineSigner;
  gasMultiplier?: number;
  defaultGasPrice?: string;
  getMessageType: (typeUrl: string) => any;
}

export function createGatewayTxClient(options: GatewayTxClientOptions): TxClient {
  const gasMultiplier = options.gasMultiplier ?? DEFAULT_GAS_MULTIPLIER;
  const gasPrice = GasPrice.fromString(options.defaultGasPrice ?? DEFAULT_AVERAGE_GAS_PRICE);

  return {
    async estimateFee(messages: EncodeObject[], memo?: string): Promise<StdFee> {
      const messageAnys: Any[] = messages.map(msg => {
        const MessageType = options.getMessageType(msg.typeUrl);
        if (!MessageType) {
          throw new Error(`Message type ${msg.typeUrl} not found in registry`);
        }
        const value = MessageType.encode(msg.value, new BinaryWriter()).finish();
        return {
          typeUrl: msg.typeUrl,
          value: value,
        };
      });

      const txBody: TxBody = {
        messages: messageAnys,
        memo: memo || "",
        timeoutHeight: Long.UZERO,
        timeoutTimestamp: undefined,
        extensionOptions: [],
        nonCriticalExtensionOptions: [],
        unordered: false,
      };

      const bodyBytes = TxBody.encode(txBody, new BinaryWriter()).finish();

      const authInfo: AuthInfo = {
        signerInfos: [{
          publicKey: undefined,
          modeInfo: {
            single: {
              mode: SignMode.SIGN_MODE_DIRECT,
            },
            multi: undefined,
          },
          sequence: Long.UZERO,
        }],
        fee: {
          amount: [{
            denom: "uakt",
            amount: "1",
          }],
          gasLimit: Long.fromNumber(200000),
          payer: "",
          granter: "",
        },
        tip: undefined,
      };

      const authInfoBytes = AuthInfo.encode(authInfo, new BinaryWriter()).finish();

      const txRaw: TxRaw = {
        bodyBytes: bodyBytes,
        authInfoBytes: authInfoBytes,
        signatures: [new Uint8Array(0)],
      };

      const writer = new BinaryWriter();
      TxRawType.encode(txRaw, writer);
      const txBytes = writer.finish();

      const simulateUrl = `${options.gatewayUrl}/cosmos/tx/v1beta1/simulate`;
      const simulateResponse = await fetch(simulateUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
          tx_bytes: Buffer.from(txBytes).toString("base64") 
        }),
      });
      if (!simulateResponse.ok) {
        let errorText = await simulateResponse.text();
        try {
          const errorJson = JSON.parse(errorText);
          errorText = errorJson.message || errorJson.error || errorText;
        } catch {
          // Not JSON, use raw text
        }
        throw new Error(`Simulate failed (${simulateResponse.status}): ${errorText}`);
      }
      const simulateData = await simulateResponse.json();
      const gasWanted = simulateData.gas_info?.gas_wanted ?? 300000;
      const minGas = Math.floor(gasMultiplier * gasWanted);
      return calculateFee(minGas, gasPrice);
    },

    async sign(messages: EncodeObject[], fee: StdFee, memo: string): Promise<TxRaw> {
      const [account] = await options.signer.getAccounts();

      if (!account) {
          throw new Error("No accounts available from signer");
      }
      
      const messageAnys: Any[] = messages.map(msg => {
        const MessageType = options.getMessageType(msg.typeUrl);
        if (!MessageType) {
          throw new Error(`Message type ${msg.typeUrl} not found in registry`);
        }
        const value = MessageType.encode(msg.value, new BinaryWriter()).finish();
        return {
          typeUrl: msg.typeUrl,
          value: value,
        };
      });

      const txBody: TxBody = {
        messages: messageAnys,
        memo: memo,
        timeoutHeight: Long.UZERO,
        timeoutTimestamp: undefined,
        extensionOptions: [],
        nonCriticalExtensionOptions: [],
        unordered: false,
      };

      const bodyBytes = TxBody.encode(txBody, new BinaryWriter()).finish();

      const feeCoins: Coin[] = fee.amount.map(coin => ({
        denom: coin.denom,
        amount: coin.amount,
      }));

      const feeProto: Fee = {
        amount: feeCoins,
        gasLimit: Long.fromString(fee.gas.toString()),
        payer: "",
        granter: "",
      };

      const signerInfo: SignerInfo = {
        publicKey: undefined,
        modeInfo: {
          single: {
            mode: SignMode.SIGN_MODE_DIRECT,
          },
          multi: undefined,
        },
        sequence: Long.UZERO,
      };

      const authInfo: AuthInfo = {
        signerInfos: [signerInfo],
        fee: feeProto,
        tip: undefined,
      };

      const authInfoBytes = AuthInfo.encode(authInfo, new BinaryWriter()).finish();

      const signDoc = makeSignDoc(
        bodyBytes,
        authInfoBytes,
        "akashnet-2",
        0,
      );

      if (!("signDirect" in options.signer) || typeof options.signer.signDirect !== "function") {
        throw new Error("signer must support signDirect method");
      }
      
      const { signed, signature } = await options.signer.signDirect(account.address, signDoc);
      
      const signatureBytes = typeof signature.signature === "string" 
        ? Uint8Array.from(Buffer.from(signature.signature, "base64"))
        : signature.signature;
      
      return {
        bodyBytes: signed.bodyBytes,
        authInfoBytes: signed.authInfoBytes,
        signatures: [signatureBytes],
      };
    },

    async broadcast(signedMessages: TxRaw): Promise<DeliverTxResponse> {
      const writer = new BinaryWriter();
      TxRawType.encode(signedMessages, writer);
      const txBytes = writer.finish();
      
      const broadcastUrl = `${options.gatewayUrl}/cosmos/tx/v1beta1/txs`;
      const broadcastResponse = await fetch(broadcastUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          tx_bytes: Buffer.from(txBytes).toString("base64"),
          mode: "BROADCAST_MODE_SYNC",
        }),
      });

      if (!broadcastResponse.ok) {
        const errorText = await broadcastResponse.text();
        throw new Error(`Broadcast failed: ${broadcastResponse.statusText} - ${errorText}`);
      }

      const broadcastData = await broadcastResponse.json();
      const txResponse = broadcastData.tx_response || broadcastData;

      const code = txResponse.code ?? txResponse.Code ?? 0;
      if (code !== 0) {
        throw new Error(`Transaction failed with code ${code}: ${txResponse.raw_log || txResponse.rawLog || ''}`);
      }

      return {
        height: txResponse.height ?? txResponse.Height ?? 0,
        transactionHash: txResponse.txhash ?? txResponse.txHash ?? txResponse.Txhash ?? '',
        code: code,
        data: txResponse.data ?? txResponse.Data ?? '',
        rawLog: txResponse.raw_log ?? txResponse.rawLog ?? txResponse.RawLog ?? '',
        gasUsed: BigInt(txResponse.gas_used ?? txResponse.gasUsed ?? txResponse.GasUsed ?? 0),
        gasWanted: BigInt(txResponse.gas_wanted ?? txResponse.gasWanted ?? txResponse.GasWanted ?? 0),
        events: txResponse.events ?? txResponse.Events ?? [],
        msgResponses: [],
        txIndex: 0,
      };
    },
  };
}
