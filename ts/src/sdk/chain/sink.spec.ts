import { afterEach, describe, expect, it, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { MsgGuardianPause, MsgGuardianPauseResponse, MsgSink, MsgSinkResponse } from "../../generated/protos/akash/sink/v1/msgs.ts";
import { Params } from "../../generated/protos/akash/sink/v1/params.ts";
import {
  QueryDisputesResponse,
  QueryReportersRequest,
  QueryReportersResponse,
  QueryStateResponse,
  QueryTransferResponse,
  QueryTransfersRequest,
  QueryTransfersResponse,
} from "../../generated/protos/akash/sink/v1/query.ts";
import { AccountUsage, DestType, ID, Status, Totals } from "../../generated/protos/akash/sink/v1/types.ts";
import type { TxClient } from "../transport/tx/TxClient.ts";
import { createChainNodeWebSDK } from "./createChainNodeWebSDK.ts";

const MAX_UINT64 = 18446744073709551615n;
const MAX_INT64 = 9223372036854775807n;

describe("sink SDK", () => {
  afterEach(() => vi.unstubAllGlobals());

  it("signs MsgSink and decodes the included record ID", async () => {
    const signer = mock<TxClient>();
    const id = ID.fromPartial({ owner: "akash1sender", height: 9007199254740993n, sequence: MAX_UINT64 });
    signer.signAndBroadcast.mockResolvedValue({
      code: 0,
      height: 10,
      txIndex: 0,
      transactionHash: "A".repeat(64),
      events: [],
      gasUsed: 1n,
      gasWanted: 1n,
      msgResponses: [{
        typeUrl: "/akash.sink.v1.MsgSinkResponse",
        value: MsgSinkResponse.encode({ id }).finish(),
      }],
    });
    const sdk = createChainNodeWebSDK({
      query: { baseUrl: "http://localhost:1317" },
      tx: { signer },
    });
    const message = MsgSink.fromPartial({
      sender: "akash1sender",
      destType: DestType.DEST_TYPE_SOLANA,
      dest: "SolanaPublicKey",
      amount: { denom: "uakt", amount: "1000000" },
    });

    await expect(sdk.akash.sink.v1.sink(message)).resolves.toEqual({ id });
    expect(signer.signAndBroadcast).toHaveBeenCalledWith(
      [{ typeUrl: "/akash.sink.v1.MsgSink", value: message }],
      expect.any(Object),
    );
  });

  it("signs the guardian pause", async () => {
    const signer = mock<TxClient>();
    signer.signAndBroadcast.mockResolvedValue({
      code: 0,
      height: 10,
      txIndex: 0,
      transactionHash: "B".repeat(64),
      events: [],
      gasUsed: 1n,
      gasWanted: 1n,
      msgResponses: [{
        typeUrl: "/akash.sink.v1.MsgGuardianPauseResponse",
        value: MsgGuardianPauseResponse.encode({}).finish(),
      }],
    });
    const sdk = createChainNodeWebSDK({
      query: { baseUrl: "http://localhost:1317" },
      tx: { signer },
    });
    const message = MsgGuardianPause.fromPartial({ guardian: "akash1guardian", deposits: true });

    await expect(sdk.akash.sink.v1.guardianPause(message)).resolves.toEqual({});
    expect(signer.signAndBroadcast).toHaveBeenCalledWith(
      [{ typeUrl: "/akash.sink.v1.MsgGuardianPause", value: message }],
      expect.any(Object),
    );
    expect(message.refunds).toBe(false);
    expect(MsgGuardianPause.decode(MsgGuardianPause.encode(message).finish())).toEqual(message);
    expect(MsgGuardianPause.toJSON(message)).toEqual({ guardian: "akash1guardian", deposits: true });
    expect(Params.fromJSON({ pause_guardian: "akash1guardian" }).pauseGuardian).toBe("akash1guardian");
  });

  it("preserves the Params enabled JSON name and field number", () => {
    const params = Params.fromJSON({ enabled: true });
    expect(params.enabled).toBe(true);
    expect(Params.toJSON(params)).toMatchObject({ enabled: true });
    expect(Params.toJSON(params)).not.toHaveProperty("deposits_enabled");
    expect(Params.decode(Uint8Array.of(8, 1)).enabled).toBe(true);
    expect(Params.encode(Params.fromPartial({ enabled: true })).finish()).toEqual(Uint8Array.of(8, 1));
    expect(Params.fromJSON({}).enabled).toBe(false);
  });

  it("queries composite IDs without losing precision or paid status", async () => {
    const id = ID.fromPartial({ owner: "akash1owner", height: 9007199254740993n, sequence: MAX_UINT64 });
    const response = QueryTransferResponse.fromPartial({ transfer: { id, status: Status.STATUS_REFUNDED, settlement: { round: 9007199254740993n } } });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify(QueryTransferResponse.toJSON(response))));
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });
    const result = await sdk.akash.sink.v1.getTransfer({ id });
    expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe(`/akash/sink/v1/transfers/${id.owner}/${id.height}/${id.sequence}`);
    expect(result.transfer?.id).toEqual(id);
    expect(result.transfer?.status).toBe(Status.STATUS_REFUNDED);
    expect(result.transfer?.settlement?.round).toBe(9007199254740993n);
    await expect(sdk.akash.sink.v1.getTransfer({ id: { owner: id.owner } })).rejects.toThrow("id.height");
  });

  it("pages transfers at a fixed height with native pagination", async () => {
    const response = QueryTransfersResponse.fromPartial({
      transfers: [{ status: Status.STATUS_REFUND_PENDING }],
      pagination: { nextKey: Uint8Array.of(3, 4) },
    });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify(QueryTransfersResponse.toJSON(response))));
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });
    const request = QueryTransfersRequest.fromPartial({ owner: "akash1owner", status: Status.STATUS_REFUND_PENDING, pagination: { key: Uint8Array.of(1, 2), limit: 500n } });
    const result = await sdk.akash.sink.v1.getTransfers(request, { header: { "x-cosmos-block-height": "9001" } });
    const url = new URL(fetch.mock.calls[0][0].toString());
    expect(url.pathname).toBe("/akash/sink/v1/transfers");
    expect(url.searchParams.get("pagination.key")).toBe("AQI=");
    expect(url.searchParams.get("pagination.limit")).toBe("500");
    expect(url.searchParams.get("status")).toBe("STATUS_REFUND_PENDING");
    expect(new Headers(fetch.mock.calls[0][1]?.headers).get("x-cosmos-block-height")).toBe("9001");
    expect(result.pagination?.nextKey).toEqual(Uint8Array.of(3, 4));
    expect(QueryTransfersRequest.decode(QueryTransfersRequest.encode(request).finish())).toEqual(request);
  });

  it("lists disputes including terminal conflicts", async () => {
    const response = QueryDisputesResponse.fromPartial({ transfers: [{ status: Status.STATUS_MINTED, settlement: { hold: 2 } }] });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify(QueryDisputesResponse.toJSON(response))));
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });
    const result = await sdk.akash.sink.v1.getDisputes({});
    expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe("/akash/sink/v1/disputes");
    expect(result.transfers[0].status).toBe(Status.STATUS_MINTED);
  });

  it("round trips Coin totals above uint64 and the exact LegacyDec ratio", () => {
    const amount = (MAX_UINT64 * 2n).toString();
    const params = Params.fromJSON({ max_transfer_amount_per_epoch: { denom: "uakt", amount }, reporters_threshold: "0.750000000000000000" });
    expect(params.maxTransferAmountPerEpoch?.amount).toBe(amount);
    expect(Params.decode(Params.encode(params).finish()).reportersThreshold).toBe("0.75");
    const totals = Totals.fromPartial({ pendingAmount: { denom: "uakt", amount }, pendingCount: MAX_UINT64 });
    expect(Totals.decode(Totals.encode(totals).finish())).toEqual(totals);
    const usage = AccountUsage.fromPartial({ totalSent: { denom: "uakt", amount }, epochSent: { denom: "uakt", amount: "1" }, lastEpoch: MAX_INT64 });
    expect(AccountUsage.decode(AccountUsage.encode(usage).finish())).toEqual(usage);
  });

  it("queries independent send/refund usage and accounting without Params or Policy", async () => {
    const response = QueryStateResponse.fromPartial({
      totals: { pendingCount: MAX_UINT64 },
      epoch: 12n,
      sendUsage: { epoch: 10n, amount: { denom: "uakt", amount: "4" } },
      refundUsage: { epoch: 11n, amount: { denom: "uakt", amount: "9" } },
      refundHead: { owner: "akash1owner", height: 1n, sequence: MAX_UINT64 },
    });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify(QueryStateResponse.toJSON(response))));
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });
    const result = await sdk.akash.sink.v1.getState();
    expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe("/akash/sink/v1/state");
    expect(result).toEqual(response);
    expect(result).not.toHaveProperty("params");
    expect(result).not.toHaveProperty("policy");
    expect(QueryStateResponse.toJSON(result)).not.toHaveProperty("params");
    expect(QueryStateResponse.toJSON(result)).not.toHaveProperty("policy");
  });

  it("encodes the five State fields with the current protobuf field numbers", () => {
    const response = QueryStateResponse.fromPartial({
      totals: {},
      epoch: 10n,
      sendUsage: {},
      refundUsage: {},
      refundHead: {},
    });
    const wire = Uint8Array.of(10, 0, 16, 10, 26, 0, 34, 0, 42, 0);

    expect(QueryStateResponse.encode(response).finish()).toEqual(wire);
    expect(QueryStateResponse.decode(wire)).toEqual(response);
  });

  it("queries reporter policies with an explicit version without losing precision", async () => {
    const json = {
      reporters: [{ version: MAX_UINT64.toString(), members: ["akash1reporter"], threshold: 1 }],
    };
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify(json)));
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });
    const request = QueryReportersRequest.fromPartial({ version: MAX_UINT64 });

    const result = await sdk.akash.sink.v1.getReporters(request);

    const url = new URL(fetch.mock.calls[0][0].toString());
    expect(url.pathname).toBe("/akash/sink/v1/reporters");
    expect(url.searchParams.get("version")).toBe(MAX_UINT64.toString());
    expect(result.reporters).toEqual([{ version: MAX_UINT64, members: ["akash1reporter"], threshold: 1 }]);
    expect(QueryReportersResponse.toJSON(result)).toEqual(json);
    expect(QueryReportersRequest.decode(QueryReportersRequest.encode(request).finish())).toEqual(request);
  });
});
