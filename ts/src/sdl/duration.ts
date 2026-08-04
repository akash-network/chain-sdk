const DURATION_PATTERN = /^([0-9]+)(ms|s|m|h)?$/;

const UNIT_MILLISECONDS = {
  ms: 1n,
  s: 1_000n,
  m: 60_000n,
  h: 3_600_000n,
};

type DurationUnit = keyof typeof UNIT_MILLISECONDS;

type ParsedDuration = {
  milliseconds: bigint;
  unit: DurationUnit | undefined;
};

export function parseDuration(value: string): ParsedDuration | undefined {
  const match = DURATION_PATTERN.exec(value);
  if (!match) return undefined;

  const unit = match[2];
  if (unit !== undefined && !isDurationUnit(unit)) return undefined;

  return {
    milliseconds: BigInt(match[1]) * (unit === undefined ? 1n : UNIT_MILLISECONDS[unit]),
    unit,
  };
}

function isDurationUnit(value: string): value is DurationUnit {
  return Object.hasOwn(UNIT_MILLISECONDS, value);
}
