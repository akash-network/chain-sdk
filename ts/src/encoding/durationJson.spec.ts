import { describe, expect, it } from "vitest";

import { Duration } from "../generated/protos/google/protobuf/duration.ts";

describe("Duration ProtoJSON", () => {
  it.each([
    { json: "0s", seconds: 0n, nanos: 0 },
    { json: "259200s", seconds: 259200n, nanos: 0 },
    { json: "259200.250s", seconds: 259200n, nanos: 250000000 },
    { json: "60.000001s", seconds: 60n, nanos: 1000 },
    { json: "60.000000001s", seconds: 60n, nanos: 1 },
    { json: "-60.250s", seconds: -60n, nanos: -250000000 },
    { json: "-0.000000001s", seconds: 0n, nanos: -1 },
    { json: "315576000000.999999999s", seconds: 315576000000n, nanos: 999999999 },
    { json: "-315576000000.999999999s", seconds: -315576000000n, nanos: -999999999 },
  ])("reads and writes $json", ({ json, seconds, nanos }) => {
    const duration = { seconds, nanos };

    expect(Duration.fromJSON(json)).toEqual(duration);
    expect(Duration.toJSON(duration)).toBe(json);
    expect(Duration.decode(Duration.encode(duration).finish())).toEqual(duration);
  });

  it.each([
    { json: "1.2s", canonical: "1.200s" },
    { json: "1.0001s", canonical: "1.000100s" },
    { json: "1.0000001s", canonical: "1.000000100s" },
    { json: "-0.000000000s", canonical: "0s" },
    { json: "1.000000000s", canonical: "1s" },
  ])("normalizes $json to $canonical", ({ json, canonical }) => {
    expect(Duration.toJSON(Duration.fromJSON(json))).toBe(canonical);
  });

  it.each([
    "", "259200", "1ms", "1.s", ".1s", "1.0000000000s", "1e3s",
    "1s trailing", "1s\n", " 1s", "+1s", "NaNs", "Infinitys",
    "315576000001s", "-315576000001s", null, 1, {},
  ])("rejects invalid JSON %j", (value) => {
    expect(() => Duration.fromJSON(value)).toThrow();
  });

  it.each([
    { seconds: 315576000001n, nanos: 0 },
    { seconds: -315576000001n, nanos: 0 },
    { seconds: 0n, nanos: 1000000000 },
    { seconds: 0n, nanos: -1000000000 },
    { seconds: 1n, nanos: -1 },
    { seconds: -1n, nanos: 1 },
    { seconds: 0n, nanos: 0.5 },
    { seconds: 0n, nanos: NaN },
    { seconds: 0n, nanos: Infinity },
  ])("rejects seconds=$seconds nanos=$nanos when writing JSON", (duration) => {
    expect(() => Duration.toJSON(duration)).toThrow();
  });
});
