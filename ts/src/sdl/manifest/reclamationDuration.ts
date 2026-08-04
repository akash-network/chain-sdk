import { Duration } from "../../generated/protos/google/protobuf/duration.ts";
import { parseDuration } from "../duration.ts";

/**
 * Converts a reclamation `min_window` into a proto `Duration`. Reclamation
 * windows require a positive whole-number value using seconds, minutes, or
 * hours. The shared duration parser performs the unit conversion exactly.
 */
export function minWindowToDuration(value: string): Duration {
  const duration = parseDuration(value);
  if (
    !duration
    || duration.unit === undefined
    || duration.unit === "ms"
    || duration.milliseconds === 0n
  ) {
    throw new Error(`invalid reclamation min_window "${value}"`);
  }

  return Duration.fromPartial({ seconds: duration.milliseconds / 1_000n, nanos: 0 });
}
