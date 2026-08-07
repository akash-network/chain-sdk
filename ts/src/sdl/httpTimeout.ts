import { parseDuration } from "./duration.ts";

export const MAX_HTTP_TIMEOUT_MILLISECONDS = 4_294_967_295;

const MAX_HTTP_TIMEOUT_MILLISECONDS_BIGINT = BigInt(MAX_HTTP_TIMEOUT_MILLISECONDS);

type HTTPTimeoutParseResult =
  | { milliseconds: number; ok: true }
  | { message: string; ok: false };

export function parseHTTPTimeout(value: number | string): HTTPTimeoutParseResult {
  if (typeof value === "number") {
    return validateMilliseconds(value);
  }

  const duration = parseDuration(value);
  if (!duration) {
    return {
      message: "must be milliseconds or a whole-number duration using ms, s, m, or h",
      ok: false,
    };
  }

  if (duration.milliseconds > MAX_HTTP_TIMEOUT_MILLISECONDS_BIGINT) {
    return {
      message: `cannot be greater than ${MAX_HTTP_TIMEOUT_MILLISECONDS} ms`,
      ok: false,
    };
  }

  return { milliseconds: Number(duration.milliseconds), ok: true };
}

function validateMilliseconds(milliseconds: number): HTTPTimeoutParseResult {
  if (!Number.isSafeInteger(milliseconds) || milliseconds < 0) {
    return { message: "must resolve to a non-negative whole number of milliseconds", ok: false };
  }

  if (milliseconds > MAX_HTTP_TIMEOUT_MILLISECONDS) {
    return {
      message: `cannot be greater than ${MAX_HTTP_TIMEOUT_MILLISECONDS} ms`,
      ok: false,
    };
  }

  return { milliseconds, ok: true };
}
