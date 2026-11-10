import type { Duration } from "../generated/protos/google/protobuf/duration.ts";

const MAX_SECONDS = 315576000000n;

export const durationJson = {
  fromJSON(value: unknown): Duration {
    if (typeof value !== "string") {
      throw new TypeError("Duration JSON must be a seconds string");
    }

    const match = /^(-?)(\d+)(?:\.(\d{1,9}))?s$/.exec(value);
    if (!match || match[0] !== value) {
      throw new TypeError("Invalid Duration JSON seconds string");
    }

    const seconds = BigInt(match[2]);
    const nanos = Number((match[3] ?? "").padEnd(9, "0"));
    const negative = match[1] === "-";
    const duration = {
      seconds: negative ? -seconds : seconds,
      nanos: negative && nanos !== 0 ? -nanos : nanos,
    };
    validateDuration(duration);
    return duration;
  },

  toJSON(duration: Duration): string {
    validateDuration(duration);
    const negative = duration.seconds < 0n || duration.nanos < 0;
    const seconds = duration.seconds < 0n ? -duration.seconds : duration.seconds;
    const nanos = Math.abs(duration.nanos);
    const fraction = nanos === 0 ? "" : `.${nanos.toString().padStart(9, "0").replace(/(?:000){1,2}$/, "")}`;
    return `${negative ? "-" : ""}${seconds}${fraction}s`;
  },
};

function validateDuration({ seconds, nanos }: Duration): void {
  if (seconds < -MAX_SECONDS || seconds > MAX_SECONDS) {
    throw new RangeError("Duration seconds out of range");
  }
  if (!Number.isInteger(nanos) || nanos < -999999999 || nanos > 999999999) {
    throw new RangeError("Duration nanos out of range");
  }
  if ((seconds < 0n && nanos > 0) || (seconds > 0n && nanos < 0)) {
    throw new RangeError("Duration seconds and nanos must have the same sign");
  }
}
