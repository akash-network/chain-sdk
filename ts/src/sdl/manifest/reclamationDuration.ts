import { Duration } from "../../generated/protos/google/protobuf/duration.ts";

const UNIT_SECONDS: Record<string, bigint> = { s: 1n, m: 60n, h: 3600n };

/**
 * Converts a reclamation `min_window` into a proto `Duration`. The SDL schema
 * pattern (`^[1-9][0-9]*(s|m|h)$`) has already guaranteed the format (and `> 0`)
 * before this runs, so we only convert. Whole-unit windows are an exact second
 * count, so `nanos` is always 0 and no BigInt-overflow handling is needed.
 */
export function minWindowToDuration(value: string): Duration {
  const match = /^([1-9][0-9]*)(s|m|h)$/.exec(value);
  if (!match) throw new Error(`invalid reclamation min_window "${value}"`);
  const seconds = BigInt(match[1]) * UNIT_SECONDS[match[2]];
  return Duration.fromPartial({ seconds: seconds.toString(), nanos: 0 });
}
