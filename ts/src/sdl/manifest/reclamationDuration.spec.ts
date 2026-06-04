import { describe, expect, it } from "vitest";

import { minWindowToDuration } from "./reclamationDuration.ts";

describe("minWindowToDuration", () => {
  it.each([
    ["1s", "1"],
    ["30m", "1800"],
    ["1h", "3600"],
    ["24h", "86400"],
    ["720h", "2592000"],
    ["8760h", "31536000"], // year-scale window: exact seconds, no precision loss
  ])("converts %j to %s seconds with zero nanos", (input, seconds) => {
    const duration = minWindowToDuration(input);
    expect(duration.seconds.toString()).toBe(seconds);
    expect(duration.nanos).toBe(0);
  });

  // Defensive: the SDL schema pattern guarantees a valid format upstream, but a
  // direct call with a value the schema would have rejected throws clearly.
  it.each(["1h30m", "1.5h", "500ms", "0s", "-1h", "100", "abc", ""])("throws on the schema-rejected value %j", (input) => {
    expect(() => minWindowToDuration(input)).toThrow(/invalid reclamation min_window/);
  });
});
