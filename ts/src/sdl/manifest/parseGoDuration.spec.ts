import { describe, expect, it } from "vitest";

import { parseGoDuration, protoDurationFromNanos } from "./parseGoDuration.ts";

describe("parseGoDuration", () => {
  describe("valid durations (mirrors Go time.ParseDuration)", () => {
    it("parses whole hours", () => {
      expect(parseGoDuration("24h")).toEqual({ ok: true, nanos: 86_400_000_000_000n });
    });

    it("parses compound durations", () => {
      expect(parseGoDuration("1h30m")).toEqual({ ok: true, nanos: 5_400_000_000_000n });
      expect(parseGoDuration("1h30m45s")).toEqual({ ok: true, nanos: 5_445_000_000_000n });
    });

    it("parses fractional durations", () => {
      expect(parseGoDuration("1.5h")).toEqual({ ok: true, nanos: 5_400_000_000_000n });
      expect(parseGoDuration("1.5s")).toEqual({ ok: true, nanos: 1_500_000_000n });
      expect(parseGoDuration(".5s")).toEqual({ ok: true, nanos: 500_000_000n });
    });

    it("parses sub-second units", () => {
      expect(parseGoDuration("500ms")).toEqual({ ok: true, nanos: 500_000_000n });
      expect(parseGoDuration("1us")).toEqual({ ok: true, nanos: 1_000n });
      expect(parseGoDuration("1ns")).toEqual({ ok: true, nanos: 1n });
    });

    it("accepts both micro-sign variants Go accepts", () => {
      expect(parseGoDuration("1µs")).toEqual({ ok: true, nanos: 1_000n }); // µ U+00B5 MICRO SIGN
      expect(parseGoDuration("1μs")).toEqual({ ok: true, nanos: 1_000n }); // μ U+03BC GREEK SMALL MU
    });

    it("accumulates large windows without JS number precision loss (> 2^53 ns)", () => {
      // 8760h == 1 year == 31_536_000 s == 3.1536e16 ns, which exceeds Number.MAX_SAFE_INTEGER (~9.0e15).
      const result = parseGoDuration("8760h");
      expect(result).toEqual({ ok: true, nanos: 31_536_000_000_000_000n });
      expect(31_536_000_000_000_000n > BigInt(Number.MAX_SAFE_INTEGER)).toBe(true);
    });

    it("treats bare \"0\" as zero (Go special case)", () => {
      expect(parseGoDuration("0")).toEqual({ ok: true, nanos: 0n });
    });

    it("parses zero with a unit", () => {
      expect(parseGoDuration("0s")).toEqual({ ok: true, nanos: 0n });
    });

    it("parses negative durations (positivity is enforced by the caller, not the parser)", () => {
      expect(parseGoDuration("-1h")).toEqual({ ok: true, nanos: -3_600_000_000_000n });
      expect(parseGoDuration("+1h")).toEqual({ ok: true, nanos: 3_600_000_000_000n });
    });
  });

  describe("invalid durations", () => {
    it.each(["abc", "", "1.5", "100", "h", "1x", "."])(
      "rejects %j",
      (input) => {
        const result = parseGoDuration(input);
        expect(result.ok).toBe(false);
      },
    );

    it("rejects values overflowing Go's int64 nanosecond ceiling", () => {
      // 10_000_000h == 3.6e19 ns > int64 max (~9.22e18 ns)
      expect(parseGoDuration("10000000h").ok).toBe(false);
    });
  });
});

describe("protoDurationFromNanos", () => {
  it("splits whole seconds", () => {
    const d = protoDurationFromNanos(86_400_000_000_000n);
    expect(d.seconds.toString()).toBe("86400");
    expect(d.nanos).toBe(0);
  });

  it("splits seconds and sub-second nanos", () => {
    const d = protoDurationFromNanos(1_500_000_000n);
    expect(d.seconds.toString()).toBe("1");
    expect(d.nanos).toBe(500_000_000);
  });

  it("represents sub-second durations with zero seconds", () => {
    const d = protoDurationFromNanos(500_000_000n);
    expect(d.seconds.toString()).toBe("0");
    expect(d.nanos).toBe(500_000_000);
  });

  it("keeps seconds and nanos sign-consistent for negatives", () => {
    const d = protoDurationFromNanos(-1_500_000_000n);
    expect(d.seconds.toString()).toBe("-1");
    expect(d.nanos).toBe(-500_000_000);
  });

  it("handles large windows without precision loss", () => {
    const d = protoDurationFromNanos(31_536_000_000_000_000n);
    expect(d.seconds.toString()).toBe("31536000");
    expect(d.nanos).toBe(0);
  });
});
