import type { CustomType } from "./CustomType.ts";
import { Int } from "./Int.ts";
import { LegacyDec } from "./LegacyDec.ts";

export const customTypes: Record<string, CustomType<unknown, unknown>> = {
  [LegacyDec.typeName]: LegacyDec,
  [LegacyDec.shortName]: LegacyDec,
  [Int.typeName]: Int,
  [Int.shortName]: Int,
  // v1beta3 ResourceValue.val declares the legacy sdk.Int customtype name;
  // both map to the same bigint-backed Int handler.
  "pkg.akt.dev/go/node/types/sdk.Int": Int,
};
