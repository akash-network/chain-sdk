import { describe, expect, it } from "vitest";

import { parseDuration } from "./duration.ts";

describe(parseDuration.name, () => {
  it.each([
    ["60000", 60_000n, undefined],
    ["500ms", 500n, "ms"],
    ["60s", 60_000n, "s"],
    ["1m", 60_000n, "m"],
    ["1h", 3_600_000n, "h"],
  ])("parses %s as milliseconds", (value, milliseconds, unit) => {
    expect(parseDuration(value)).toEqual({ milliseconds, unit });
  });

  it.each(["", "-1s", "1.5s", "1h30m", "60seconds"])("rejects %j", (value) => {
    expect(parseDuration(value)).toBeUndefined();
  });
});
