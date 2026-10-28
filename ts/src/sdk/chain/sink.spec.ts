import { afterEach, describe, expect, it, vi } from "vitest";
import { mock } from "vitest-mock-extended";

import { MsgGuardianPause, MsgGuardianPauseResponse, MsgSink, MsgSinkResponse } from "../../generated/protos/akash/sink/v1/msgs.ts";
import { Params } from "../../generated/protos/akash/sink/v1/params.ts";
import {
  QueryOpenTransfersRequest,
  QueryOpenTransfersResponse,
  QuerySinkResponse,
  QuerySinkResponse_SinkStatus,
  QueryStateResponse,
} from "../../generated/protos/akash/sink/v1/query.ts";
import { AccountUsage, Settlement, Status, Totals } from "../../generated/protos/akash/sink/v1/types.ts";
import type { TxClient } from "../transport/tx/TxClient.ts";
import { createChainNodeWebSDK } from "./createChainNodeWebSDK.ts";

const MAX_UINT64 = 18446744073709551615n;

describe("sink SDK", () => {
  afterEach(() => vi.unstubAllGlobals());

  it("signs MsgSink and decodes the included record ID", async () => {
    const signer = mock<TxClient>();
    const id = MAX_UINT64.toString();
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
      destType: 1,
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

  it("queries a record over REST without losing uint64 precision or refund state", async () => {
    const response = QuerySinkResponse.fromPartial({
      status: QuerySinkResponse_SinkStatus.void,
      transfer: {
        record: Uint8Array.of(1, 2, 3),
        settlement: { status: Status.STATUS_REFUNDED, round: 9007199254740993n },
      },
    });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(
      new Response(JSON.stringify(QuerySinkResponse.toJSON(response))),
    );
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });

    const result = await sdk.akash.sink.v1.getSink({ id: MAX_UINT64.toString() });

    expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe(`/akash/sink/v1/records/${MAX_UINT64}`);
    expect(result.status).toBe(QuerySinkResponse_SinkStatus.void);
    expect(result.transfer?.settlement?.status).toBe(Status.STATUS_REFUNDED);
    expect(result.transfer?.settlement?.round).toBe(9007199254740993n);
  });

  it("pages open transfers over REST with unsigned cursors", async () => {
    const response = QueryOpenTransfersResponse.fromPartial({
      transfers: [
        { record: Uint8Array.of(1), settlement: { status: Status.STATUS_PENDING, round: 1n } },
        { record: Uint8Array.of(2), voidOrdered: true, settlement: { status: Status.STATUS_VOIDED, round: 2n } },
      ],
      nextStartAfter: MAX_UINT64,
    });
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(
      new Response(JSON.stringify(QueryOpenTransfersResponse.toJSON(response))),
    );
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });

    const result = await sdk.akash.sink.v1.getOpenTransfers({ startAfter: MAX_UINT64 - 1n, limit: 500 });

    const url = new URL(fetch.mock.calls[0][0].toString());
    expect(url.pathname).toBe("/akash/sink/v1/open_transfers");
    expect(url.searchParams.get("start_after")).toBe((MAX_UINT64 - 1n).toString());
    expect(url.searchParams.get("limit")).toBe("500");
    expect(result.transfers).toHaveLength(2);
    expect(result.transfers[1].voidOrdered).toBe(true);
    expect(result.transfers[1].settlement?.status).toBe(Status.STATUS_VOIDED);
    expect(result.nextStartAfter).toBe(MAX_UINT64);

    const request = QueryOpenTransfersRequest.fromPartial({ startAfter: MAX_UINT64 });
    expect(request.startAfter).toBe(MAX_UINT64);
    expect(request.limit).toBe(0);
    expect(QueryOpenTransfersRequest.decode(QueryOpenTransfersRequest.encode(request).finish())).toEqual(request);
  });

  it("decodes uint64 values above the signed range in fromJSON and fromPartial", () => {
    const max = MAX_UINT64.toString();
    const params = Params.fromJSON({ max_transfer: max, min_transfer: "1000000" });
    expect(params.maxTransfer).toBe(MAX_UINT64);
    expect(params.minTransfer).toBe(1000000n);
    expect(Params.decode(Params.encode(params).finish())).toEqual(params);

    const totals = Totals.fromPartial({ lastTransferId: MAX_UINT64, reporterEpoch: MAX_UINT64 });
    expect(totals.lastTransferId).toBe(MAX_UINT64);
    expect(totals.reporterEpoch).toBe(MAX_UINT64);
    expect(Totals.decode(Totals.encode(totals).finish())).toEqual(totals);

    const capped = Params.fromJSON({ max_daily_amount_per_account: max });
    expect(capped.maxDailyAmountPerAccount).toBe(MAX_UINT64);
    expect(Params.fromPartial({ maxDailyAmountPerAccount: MAX_UINT64 }).maxDailyAmountPerAccount).toBe(MAX_UINT64);
    expect(Params.toJSON(capped)).toMatchObject({ max_daily_amount_per_account: max });
    expect(Params.fromJSON({}).maxDailyAmountPerAccount).toBe(0n);

    const usage = AccountUsage.fromJSON({ daily_amount: max, daily_count: 2 });
    expect(usage.dailyAmount).toBe(MAX_UINT64);
    expect(AccountUsage.decode(AccountUsage.encode(usage).finish())).toEqual(usage);

    // Signed fields keep their sign.
    expect(Settlement.fromPartial({ challengeEnd: -1n }).challengeEnd).toBe(-1n);
    expect(Settlement.fromJSON({ challenge_end: "-1" }).challengeEnd).toBe(-1n);
  });

  it("queries state over REST without turning uint64 limits negative", async () => {
    const max = MAX_UINT64.toString();
    const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(
      new Response(JSON.stringify({ params: { max_transfer: max, daily_refund_cap: max }, totals: { pending_amount: max } })),
    );
    vi.stubGlobal("fetch", fetch);
    const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });

    const result = await sdk.akash.sink.v1.getState();

    expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe("/akash/sink/v1/state");
    expect(result.params?.maxTransfer).toBe(MAX_UINT64);
    expect(result.params?.dailyRefundCap).toBe(MAX_UINT64);
    expect(result.totals?.pendingAmount).toBe(MAX_UINT64);
    expect(QueryStateResponse.toJSON(result)).toMatchObject({ params: { max_transfer: max } });
  });
});
