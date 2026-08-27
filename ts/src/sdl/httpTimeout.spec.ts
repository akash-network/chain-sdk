import { describe, expect, it } from "vitest";

import { MAX_HTTP_TIMEOUT_MILLISECONDS, parseHTTPTimeout } from "./httpTimeout.ts";

describe(parseHTTPTimeout.name, () => {
  it.each([
    [60_001, 60_001],
    ["60000", 60_000],
    ["60s", 60_000],
    ["1m", 60_000],
    ["1h", 3_600_000],
    ["1193h", 4_294_800_000],
    [MAX_HTTP_TIMEOUT_MILLISECONDS, MAX_HTTP_TIMEOUT_MILLISECONDS],
    [String(MAX_HTTP_TIMEOUT_MILLISECONDS), MAX_HTTP_TIMEOUT_MILLISECONDS],
  ])("normalizes %s to milliseconds", (input, expected) => {
    expect(parseHTTPTimeout(input)).toEqual({ milliseconds: expected, ok: true });
  });

  it.each(["1.5s", "1h30m", "60seconds", -1, 1.5])("rejects invalid value %s", (input) => {
    expect(parseHTTPTimeout(input)).toEqual(expect.objectContaining({ ok: false }));
  });

  it("rejects values above the manifest maximum", () => {
    expect(parseHTTPTimeout("1194h")).toEqual({
      message: `cannot be greater than ${MAX_HTTP_TIMEOUT_MILLISECONDS} ms`,
      ok: false,
    });
  });
});
