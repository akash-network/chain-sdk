import { describe, expect, it } from "vitest";

import { CPUInfo } from "../../src/generated/protos/index.provider.akash.v1.ts";

// `CPUInfo.arch` is what lets a consumer tell an arm64 node from an amd64 one in
// provider inventory. It is a generated field, so nothing else in this repo would
// notice if a proto edit or a codegen change dropped it — hence a test that reads
// it back off the wire rather than trusting the type.
describe("inventory CPUInfo.arch", () => {
  it("survives an encode/decode round trip", () => {
    const info = CPUInfo.fromPartial({
      id: "0",
      vendor: "ARM",
      model: "Neoverse-N1",
      vcores: 8,
      arch: "arm64",
    });

    expect(CPUInfo.decode(CPUInfo.encode(info).finish())).toEqual(info);
  });

  it("defaults to an empty string for inventory that predates the field", () => {
    const legacy = CPUInfo.fromPartial({ id: "0", vendor: "GenuineIntel", vcores: 8 });

    expect(legacy.arch).toBe("");
    expect(CPUInfo.decode(CPUInfo.encode(legacy).finish()).arch).toBe("");
  });
});
