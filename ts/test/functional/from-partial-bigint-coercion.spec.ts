import { describe, expect, it } from "vitest";

import { DeploymentID } from "../../src/generated/protos/akash/deployment/v1/deployment.ts";
import { BitArray } from "../../src/generated/protos/tendermint/libs/bits/types.ts";

// These assertions also guard the `coerceBigIntFromPartial` pass in
// script/fix-ts-proto-generated-types.ts: if a ts-proto upgrade changes the
// generated `fromPartial` shape and the codemod silently stops matching, the
// coercion disappears and these tests fail loudly instead.
describe("fromPartial bigint coercion", () => {
  describe("scalar bigint field (DeploymentID.dseq)", () => {
    it("coerces string, number and bigint inputs to bigint", () => {
      expect(DeploymentID.fromPartial({ owner: "a", dseq: "42" }).dseq).toBe(42n);
      expect(DeploymentID.fromPartial({ owner: "a", dseq: 42 }).dseq).toBe(42n);
      expect(DeploymentID.fromPartial({ owner: "a", dseq: 42n }).dseq).toBe(42n);
    });

    it("defaults missing/null values to 0n", () => {
      expect(DeploymentID.fromPartial({ owner: "a" }).dseq).toBe(0n);
      expect(DeploymentID.fromPartial({ owner: "a", dseq: undefined }).dseq).toBe(0n);
    });

    it("throws on non-integer / non-numeric input rather than storing it uncoerced", () => {
      expect(() => DeploymentID.fromPartial({ owner: "a", dseq: 1.5 })).toThrow();
      expect(() => DeploymentID.fromPartial({ owner: "a", dseq: "abc" })).toThrow();
    });
  });

  describe("repeated bigint field (BitArray.elems)", () => {
    it("coerces every element to bigint", () => {
      const message = BitArray.fromPartial({ bits: "3", elems: ["1", 2, 3n] });

      expect(message.bits).toBe(3n);
      expect(message.elems).toEqual([1n, 2n, 3n]);
      message.elems.forEach((e) => expect(typeof e).toBe("bigint"));
    });

    it("defaults a missing repeated field to an empty array", () => {
      expect(BitArray.fromPartial({ bits: 0n }).elems).toEqual([]);
    });
  });
});
