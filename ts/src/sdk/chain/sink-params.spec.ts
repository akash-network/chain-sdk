import { afterEach, describe, expect, it, vi } from "vitest";

import vectors from "../../../../go/node/sink/v1/testdata/params_vectors.json" with { type: "json" };
import { Params } from "../../generated/protos/akash/sink/v1/params.ts";
import { QueryStateResponse } from "../../generated/protos/akash/sink/v1/query.ts";
import { createChainNodeWebSDK } from "./createChainNodeWebSDK.ts";

describe("sink native Params vectors", () => {
  afterEach(() => vi.unstubAllGlobals());

  for (const vector of vectors) {
    const nanos = BigInt(vector.duration_nanos);
    const duration = {
      seconds: nanos / 1000000000n,
      nanos: Number(nanos % 1000000000n),
    };

    it(`encodes Go JSON as the same protobuf bytes for ${vector.name}`, () => {
      const params = Params.fromJSON(vector.json);

      expect(params.reporterChallengePeriod).toEqual(duration);
      expect(Buffer.from(Params.encode(params).finish()).toString("hex")).toBe(vector.proto_hex);
    });

    it(`preserves Go protobuf and writes native Duration JSON for ${vector.name}`, () => {
      const params = Params.decode(Buffer.from(vector.proto_hex, "hex"));

      expect(params.reporterChallengePeriod).toEqual(duration);
      expect(Buffer.from(Params.encode(params).finish()).toString("hex")).toBe(vector.proto_hex);
      expect(Params.toJSON(params)).toMatchObject({
        reporter_challenge_period: vector.json.reporter_challenge_period,
      });
    });

    it(`reads Go Params through getState for ${vector.name}`, async () => {
      const fetch = vi.fn<typeof globalThis.fetch>().mockResolvedValue(new Response(JSON.stringify({ params: vector.json })));
      vi.stubGlobal("fetch", fetch);
      const sdk = createChainNodeWebSDK({ query: { baseUrl: "http://localhost:1317" } });

      const result = await sdk.akash.sink.v1.getState();

      expect(fetch).toHaveBeenCalledTimes(1);
      expect(new URL(fetch.mock.calls[0][0].toString()).pathname).toBe("/akash/sink/v1/state");
      expect(result.params?.reporterChallengePeriod).toEqual(duration);
      expect(QueryStateResponse.toJSON(result)).toMatchObject({
        params: { reporter_challenge_period: vector.json.reporter_challenge_period },
      });
    });
  }
});
