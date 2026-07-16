import { describe, expect, it } from "vitest";

import { convertCpuResourceString, convertResourceString } from "./sizes.ts";

describe("convertResourceString", () => {
  describe("integer inputs", () => {
    it("should convert kilobytes (decimal)", () => {
      expect(convertResourceString("1k")).toBe(1000n);
    });

    it("should convert kilobytes (binary)", () => {
      expect(convertResourceString("1Ki")).toBe(1024n);
    });

    it("should convert megabytes (decimal)", () => {
      expect(convertResourceString("1m")).toBe(1000000n);
    });

    it("should convert megabytes (binary)", () => {
      expect(convertResourceString("1Mi")).toBe(1048576n);
    });

    it("should convert gigabytes (decimal)", () => {
      expect(convertResourceString("1g")).toBe(1000000000n);
    });

    it("should convert gigabytes (binary)", () => {
      expect(convertResourceString("1Gi")).toBe(1073741824n);
    });

    it("should convert terabytes (decimal)", () => {
      expect(convertResourceString("1t")).toBe(1000000000000n);
    });

    it("should convert terabytes (binary)", () => {
      expect(convertResourceString("1Ti")).toBe(1099511627776n);
    });

    it("should convert petabytes (decimal)", () => {
      expect(convertResourceString("1p")).toBe(1000000000000000n);
    });

    it("should convert petabytes (binary)", () => {
      expect(convertResourceString("1Pi")).toBe(1125899906842624n);
    });

    it("should convert exabytes (decimal)", () => {
      expect(convertResourceString("1e")).toBe(1000000000000000000n);
    });

    it("should convert exabytes (binary)", () => {
      expect(convertResourceString("1Ei")).toBe(1152921504606846976n);
    });
  });

  describe("decimal inputs", () => {
    it("should convert decimal kilobytes (decimal) and return integer", () => {
      expect(convertResourceString("0.5k")).toBe(500n);
    });

    it("should convert decimal kilobytes (binary) and return integer", () => {
      expect(convertResourceString("0.5Ki")).toBe(512n);
    });

    it("should convert decimal megabytes (decimal) and return integer", () => {
      expect(convertResourceString("0.5m")).toBe(500000n);
    });

    it("should convert decimal megabytes (binary) and return integer", () => {
      expect(convertResourceString("0.5Mi")).toBe(524288n);
    });

    it("should convert decimal gigabytes (decimal) and return integer", () => {
      expect(convertResourceString("0.3g")).toBe(300000000n);
    });

    it("should convert decimal gigabytes (binary) and return integer using ceil", () => {
      // 0.3 * 1024^3 = 322122547.2
      expect(convertResourceString("0.3Gi")).toBe(322122548n);
    });

    it("should convert decimal terabytes (decimal) and return integer", () => {
      expect(convertResourceString("0.1t")).toBe(100000000000n);
    });

    it("should convert decimal terabytes (binary) and return integer", () => {
      // 0.1 * 1024^4 = 109951162777.6
      expect(convertResourceString("0.1Ti")).toBe(109951162778n);
    });

    it("should handle very small decimal values and round up", () => {
      // 0.001 * 1024^3 = 1073741.824
      expect(convertResourceString("0.001Gi")).toBe(1073742n);
    });

    it("should handle decimal values with multiple decimal places", () => {
      // 1.234 * 1024^2 = 1293942.784
      expect(convertResourceString("1.234Mi")).toBe(1293943n);
    });

    it("should handle decimal values that result in exact integers", () => {
      // 0.0009765625 * 1024^2 = 1024
      expect(convertResourceString("0.0009765625Mi")).toBe(1024n);
    });
  });

  describe("edge cases", () => {
    it("should handle case insensitivity", () => {
      const result1 = convertResourceString("1GI");
      const result2 = convertResourceString("1gi");
      const result3 = convertResourceString("1Gi");
      expect(result1).toBe(result2);
      expect(result2).toBe(result3);
    });

    it("should handle large values", () => {
      const result = convertResourceString("999.999Gi");
      expect(result).toBeGreaterThan(0n);
    });

    it("should handle zero values", () => {
      expect(convertResourceString("0Gi")).toBe(0n);
    });

    it("should compute large binary quantities without losing precision", () => {
      // 9.5 Ei is well beyond Number.MAX_SAFE_INTEGER and is not exactly
      // representable as an IEEE-754 double, so float arithmetic would round it.
      // Exact bigint arithmetic must reproduce ceil(9.5 * 1024^6) exactly.
      const oneEi = 1024n ** 6n;
      expect(convertResourceString("9.5Ei")).toBe((oneEi * 95n) / 10n);
    });

    it("should always return exact bigints to avoid big.Int unmarshal errors", () => {
      const testCases = ["0.3Gi", "0.7Mi", "1.5k", "2.3m", "0.001Ti", "3.14159g", "0.123Ki"];

      testCases.forEach((testCase) => {
        expect(typeof convertResourceString(testCase)).toBe("bigint");
      });
    });
  });
});

describe("convertCpuResourceString", () => {
  it("should convert whole CPU units to millicpus", () => {
    expect(convertCpuResourceString("1")).toBe(1000n);
  });

  it("should convert decimal CPU units to millicpus", () => {
    expect(convertCpuResourceString("0.5")).toBe(500n);
  });

  it("should keep millicpu values as is", () => {
    expect(convertCpuResourceString("500m")).toBe(500n);
  });

  it("should round fractional millicpu values up to an integer", () => {
    // Millicpu is the smallest on-chain unit, so a fractional millicpu is
    // rounded up rather than passed through as a non-integer.
    expect(convertCpuResourceString("250.5m")).toBe(251n);
  });

  it("should handle case insensitivity", () => {
    const result1 = convertCpuResourceString("500M");
    const result2 = convertCpuResourceString("500m");
    expect(result1).toBe(result2);
  });

  it("should handle zero values", () => {
    expect(convertCpuResourceString("0")).toBe(0n);
  });
});
