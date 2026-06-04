import { Duration } from "../../generated/protos/google/protobuf/duration.ts";

const UNIT_SECONDS: Record<string, bigint> = { s: 1n, m: 60n, h: 3600n };

/**
 * Converts a reclamation `min_window` into a proto `Duration`. The SDL schema
 * pattern (`^[1-9][0-9]*(s|m|h)$`) has already guaranteed the format (and `> 0`)
 * before this runs, so on the request hot path we split off the trailing unit
 * and parse the amount directly instead of re-running a regex. The cheap
 * unit/integer/positivity guards stay only to reject obviously-invalid direct
 * calls. Whole-unit windows are an exact second count, so `nanos` is always 0
 * and the BigInt product never loses precision.
 */
export function minWindowToDuration(value: string): Duration {
  const unitSeconds = UNIT_SECONDS[value.at(-1) ?? ""];
  const amount = Number(value.slice(0, -1));
  if (unitSeconds === undefined || !Number.isInteger(amount) || amount <= 0) {
    throw new Error(`invalid reclamation min_window "${value}"`);
  }
  return Duration.fromPartial({ seconds: (BigInt(amount) * unitSeconds).toString(), nanos: 0 });
}
