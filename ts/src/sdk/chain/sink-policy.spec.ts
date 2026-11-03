import { describe, expect, it } from "vitest";

import { GenesisState } from "../../generated/protos/akash/sink/v1/genesis.ts";
import { QueryReportersResponse } from "../../generated/protos/akash/sink/v1/query.ts";
import { ReporterPolicy, SettlementData } from "../../generated/protos/akash/sink/v1/types.ts";

describe("sink native reporter policy", () => {
  it.each([
    ["0.750000000000000000", "750000000000000000", "0.75"],
    ["0.000000000000000000", "0", "0"],
    ["1.000000000000000000", "1000000000000000000", "1"],
    ["0.000000000000000001", "1", "0.000000000000000001"],
    ["0.999999999999999999", "999999999999999999", "0.999999999999999999"],
    ["0.666666666666666667", "666666666666666667", "0.666666666666666667"],
  ])("preserves threshold precision and atomic wire bytes for %s", (decimal, atoms, normalized) => {
    const policy = ReporterPolicy.fromJSON({ threshold: decimal });
    const wire = Buffer.concat([Buffer.from([26, atoms.length]), Buffer.from(atoms)]);

    expect(ReporterPolicy.encode(policy).finish()).toEqual(new Uint8Array(wire));
    expect(ReporterPolicy.encode(ReporterPolicy.fromPartial({ threshold: decimal })).finish()).toEqual(new Uint8Array(wire));
    expect(ReporterPolicy.toJSON(policy)).toEqual({ threshold: decimal });
    const decoded = ReporterPolicy.decode(wire);
    expect(decoded.threshold).toBe(normalized);
    expect(ReporterPolicy.toJSON(decoded)).toEqual({ threshold: normalized });
    expect(ReporterPolicy.encode(decoded).finish()).toEqual(new Uint8Array(wire));
  });

  it("rejects more than 18 decimal places before policy encoding", () => {
    const policy = ReporterPolicy.fromPartial({ threshold: "0.0000000000000000001" });

    expect(() => ReporterPolicy.encode(policy).finish()).toThrow();
  });

  it("preserves policy fractions in queries, settlement evidence and genesis", () => {
    const version = 18446744073709551615n;
    const policy = ReporterPolicy.fromPartial({
      version,
      members: ["akash1reporter"],
      threshold: "0.666666666666666667",
    });
    const settlement = SettlementData.fromPartial({ policy });
    const reporters = QueryReportersResponse.fromPartial({ reporters: [policy] });
    const genesis = GenesisState.fromPartial({
      policySequence: version,
      policies: [policy],
      records: [{ settlement }],
    });

    expect(SettlementData.decode(SettlementData.encode(settlement).finish())).toEqual(settlement);
    expect(QueryReportersResponse.decode(QueryReportersResponse.encode(reporters).finish())).toEqual(reporters);
    expect(GenesisState.decode(GenesisState.encode(genesis).finish())).toEqual(genesis);
    expect(GenesisState.fromJSON(GenesisState.toJSON(genesis))).toEqual(genesis);
    expect(GenesisState.toJSON(genesis)).toMatchObject({ policy_sequence: version.toString() });
  });

  it("uses Artur's eight genesis field numbers and removes the old version fields", () => {
    const genesis = GenesisState.fromPartial({
      params: {},
      totals: {},
      policySequence: 10n,
      policies: [{}],
      sendUsage: {},
      refundUsage: {},
      records: [{}],
      accounts: [{}],
    });
    const wire = Uint8Array.of(10, 0, 18, 0, 24, 10, 34, 0, 42, 0, 50, 0, 58, 0, 66, 0);

    expect(GenesisState.encode(genesis).finish()).toEqual(wire);
    expect(GenesisState.decode(wire)).toEqual(genesis);
    expect(genesis).not.toHaveProperty("schemaVersion");
    expect(genesis).not.toHaveProperty("currentPolicyVersion");
    expect(GenesisState.toJSON(genesis)).not.toHaveProperty("schema_version");
    expect(GenesisState.toJSON(genesis)).not.toHaveProperty("current_policy_version");
  });
});
