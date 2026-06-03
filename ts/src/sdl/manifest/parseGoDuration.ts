import { Duration } from "../../generated/protos/google/protobuf/duration.ts";

/**
 * Hand-rolled to match Go's `time.ParseDuration` exactly so the TS SDL pipeline
 * accepts/rejects reclamation `min_window` values identically to the chain's Go
 * parser. See `go/sdl/reclamation.go` (`toDeploymentReclamation`), which calls
 * `time.ParseDuration` then rejects `<= 0`. Positivity is NOT enforced here (Go's
 * `ParseDuration` happily parses "0s" and "-1h"); the caller checks `nanos <= 0`.
 *
 * Accumulation is done in BigInt because a year-scale window (e.g. "8760h" =
 * 3.1536e16 ns) exceeds `Number.MAX_SAFE_INTEGER` (~9.0e15) and would lose
 * precision as a JS `number`. Go's int64-nanosecond overflow ceiling is mirrored.
 */

const NS_PER_SECOND = 1_000_000_000n;

// Go's time.Duration is an int64 nanosecond count.
const INT64_MAX = (1n << 63n) - 1n; // 9_223_372_036_854_775_807
const TWO_POW_63 = 1n << 63n; //        9_223_372_036_854_775_808

// unit -> nanoseconds, matching Go's `time.unitMap`.
const UNIT_MAP: Record<string, bigint> = {
  ns: 1n,
  us: 1_000n,
  µs: 1_000n, // U+00B5 MICRO SIGN
  μs: 1_000n, // U+03BC GREEK SMALL LETTER MU
  ms: 1_000_000n,
  s: 1_000_000_000n,
  m: 60_000_000_000n,
  h: 3_600_000_000_000n,
};

export type ParseGoDurationResult =
  | { ok: true; nanos: bigint }
  | { ok: false; error: string };

function isDigit(c: string): boolean {
  return c >= "0" && c <= "9";
}

/**
 * Consumes the leading run of digits from `s`. Returns null on overflow past
 * 2^63 (Go's `errLeadingInt`).
 */
function leadingInt(s: string): { x: bigint; rem: string } | null {
  let x = 0n;
  let i = 0;
  for (; i < s.length; i++) {
    const c = s[i];
    if (!isDigit(c)) break;
    if (x > TWO_POW_63 / 10n) return null; // overflow
    x = x * 10n + BigInt(c.charCodeAt(0) - 48);
    if (x > TWO_POW_63) return null; // overflow
  }
  return { x, rem: s.slice(i) };
}

/**
 * Consumes the leading run of digits from `s` as a fraction, returning the
 * integer value of the digits and the scale (10^len). Digits that would
 * overflow are dropped, matching Go's `leadingFraction`.
 */
function leadingFraction(s: string): { x: bigint; scale: number; rem: string } {
  let x = 0n;
  let scale = 1;
  let overflow = false;
  let i = 0;
  for (; i < s.length; i++) {
    const c = s[i];
    if (!isDigit(c)) break;
    if (overflow) continue;
    if (x > INT64_MAX / 10n) {
      overflow = true;
      continue;
    }
    const y = x * 10n + BigInt(c.charCodeAt(0) - 48);
    if (y > TWO_POW_63) {
      overflow = true;
      continue;
    }
    x = y;
    scale *= 10;
  }
  return { x, scale, rem: s.slice(i) };
}

/**
 * Parses a Go duration string into a signed nanosecond count. Grammar (per Go):
 * `[-+]?([0-9]*(\.[0-9]*)?[unit]+)+` with units ns/us/µs/μs/ms/s/m/h, plus the
 * special case `"0"`.
 */
export function parseGoDuration(input: string): ParseGoDurationResult {
  const invalid = (): ParseGoDurationResult => ({ ok: false, error: `invalid duration "${input}"` });

  let s = input;
  let d = 0n;
  let neg = false;

  // Consume [-+]?
  if (s.length > 0 && (s[0] === "-" || s[0] === "+")) {
    neg = s[0] === "-";
    s = s.slice(1);
  }

  // Special case: "0" (with the optional sign already stripped).
  if (s === "0") return { ok: true, nanos: 0n };
  if (s === "") return invalid();

  while (s !== "") {
    let v = 0n; // integer part, in the current unit
    let f = 0n; // fractional digits as an integer
    let scale = 1; // 10^(number of fractional digits)

    // The next character must be [0-9.]
    if (!(s[0] === "." || isDigit(s[0]))) return invalid();

    // Consume [0-9]*
    const beforeInt = s.length;
    const li = leadingInt(s);
    if (li === null) return invalid();
    v = li.x;
    s = li.rem;
    const pre = beforeInt !== s.length;

    // Consume (\.[0-9]*)?
    let post = false;
    if (s !== "" && s[0] === ".") {
      s = s.slice(1);
      const beforeFrac = s.length;
      const lf = leadingFraction(s);
      f = lf.x;
      scale = lf.scale;
      s = lf.rem;
      post = beforeFrac !== s.length;
    }

    if (!pre && !post) return invalid(); // no digits (e.g. ".s")

    // Consume the unit.
    let i = 0;
    for (; i < s.length; i++) {
      const c = s[i];
      if (c === "." || isDigit(c)) break;
    }
    if (i === 0) return { ok: false, error: `missing unit in duration "${input}"` };
    const unitName = s.slice(0, i);
    s = s.slice(i);
    const unit = UNIT_MAP[unitName];
    if (unit === undefined) return { ok: false, error: `unknown unit "${unitName}" in duration "${input}"` };

    if (v > TWO_POW_63 / unit) return invalid(); // overflow
    v *= unit;

    if (f > 0n) {
      // float64(f) * (float64(unit) / scale), truncated toward zero — matches
      // Go's `uint64(float64(f) * (float64(unit)/scale))`, preserving its rounding.
      v += BigInt(Math.trunc(Number(f) * (Number(unit) / scale)));
      if (v > TWO_POW_63) return invalid(); // overflow
    }

    d += v;
    if (d > TWO_POW_63) return invalid(); // overflow
  }

  if (neg) return { ok: true, nanos: -d };
  if (d > INT64_MAX) return invalid(); // positive overflow
  return { ok: true, nanos: d };
}

/**
 * Splits a signed nanosecond count into a proto `Duration` ({ seconds, nanos }),
 * keeping the two fields sign-consistent as the proto convention requires. This
 * is the split Go's proto marshaling performs on a `time.Duration`.
 */
export function protoDurationFromNanos(nanos: bigint): Duration {
  const seconds = nanos / NS_PER_SECOND; // truncates toward zero
  const remainder = nanos % NS_PER_SECOND; // same sign as `nanos`
  return Duration.fromPartial({ seconds: seconds.toString(), nanos: Number(remainder) });
}
