const prefixes = "kmgtpe".split("");

/**
 * Converts resource strings like "1k", "5gi", "10m" to their numeric values.
 *
 * The result is a `bigint` computed with exact integer arithmetic so that large
 * quantities (petabyte/exabyte-scale storage and memory) cannot silently lose
 * precision the way intermediate IEEE-754 `number` arithmetic would once the
 * value exceeds `Number.MAX_SAFE_INTEGER`.
 *
 * @example
 * ```ts
 * convertResourceString("1k") // Returns 1000n
 * convertResourceString("5gi") // Returns 5368709120n
 * convertResourceString("10m") // Returns 10000000n
 * ```
 */
export function convertResourceString(resourceStr: string): bigint {
  const [value, prefix, unit] = parseSizeString(resourceStr.toLowerCase());
  const power = prefixes.indexOf(prefix);
  const base = unit === "i" ? 1024n : 1000n;

  const { numerator, denominator } = parseDecimal(value);
  const multiplier = power !== -1 ? base ** BigInt(power + 1) : 1n;

  return ceilDiv(numerator * multiplier, denominator);
}

/**
 * Converts CPU resource strings to their millicpu values as an exact `bigint`.
 *
 * Fractional millicpu inputs (e.g. "250.5m") are rounded up, matching the
 * ceiling behavior used for memory/storage — the on-chain resource value is an
 * integer, so a fractional millicpu cannot be represented.
 *
 * @example
 * ```ts
 * convertCpuResourceString("1") // Returns 1000n
 * convertCpuResourceString("500m") // Returns 500n
 * ```
 */
export function convertCpuResourceString(resourceStr: string): bigint {
  const [value, unit] = parseCpuResourceString(resourceStr.toLowerCase());

  const { numerator, denominator } = parseDecimal(value);
  const multiplier = unit === "m" ? 1n : 1000n;

  return ceilDiv(numerator * multiplier, denominator);
}

/**
 * Parses a size string into value and unit components
 * @internal
 */
function parseSizeString(size: string): [string, string, string] {
  const regex = /^([\d.]+)([a-zA-Z])([a-zA-Z]*)$/;
  const match = size.match(regex);

  if (match) {
    const [, value, unit1, unit2] = match;
    return [value, unit1.toLowerCase(), unit2.toLowerCase()];
  }

  throw new Error(`Invalid size string: ${size}`);
}

/**
 * Parses a CPU resource string into value and unit components
 * @internal
 */
function parseCpuResourceString(size: string): [string, string] {
  const regex = /^([\d.]+)([a-zA-Z]*)$/;
  const match = size.match(regex);

  if (match) {
    const [, value, unit1] = match;
    return [value, unit1.toLowerCase()];
  }

  throw new Error(`Invalid size string: ${size}`);
}

/**
 * Represents a decimal string like "0.3" as an exact fraction
 * (numerator / denominator) using bigint components, so downstream
 * multiplication stays lossless.
 * @internal
 */
function parseDecimal(value: string): { numerator: bigint; denominator: bigint } {
  const [intPart = "", fracPart = ""] = value.split(".");
  const digits = intPart + fracPart;
  return {
    numerator: digits === "" ? 0n : BigInt(digits),
    denominator: 10n ** BigInt(fracPart.length),
  };
}

/**
 * Ceiling division for non-negative bigints — mirrors the `Math.ceil` rounding
 * previously applied to the floating-point result.
 * @internal
 */
function ceilDiv(numerator: bigint, denominator: bigint): bigint {
  return (numerator + denominator - 1n) / denominator;
}
