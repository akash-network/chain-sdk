import { bytesFromBase64 } from "../base64.ts";
import type { CustomType } from "./CustomType.ts";

const decoder = new TextDecoder("utf-8");
const encoder = new TextEncoder();

/**
 * `cosmossdk.io/math.Int` (and the legacy `pkg.akt.dev/go/node/types/sdk.Int`)
 * are stored on the wire as a `bytes` field whose content is the decimal ASCII
 * representation of the integer — exactly what Go's `Int.MarshalText()` emits
 * (the integer 1000 -> the 4 bytes "1000"). We surface it as a `bigint`.
 */
export const Int = {
  typeName: "cosmossdk.io/math.Int",
  shortName: "Int",
  jsType: "bigint",
  encode(value: bigint): string {
    return value.toString();
  },
  decode(value: string): bigint {
    return value.length ? BigInt(value) : 0n;
  },
} as const satisfies CustomType<bigint, string>;

export function bigIntFromBytes(bytes: Uint8Array): bigint {
  return Int.decode(decoder.decode(bytes));
}

export function bytesFromBigInt(value: bigint): Uint8Array {
  return encoder.encode(Int.encode(value));
}

export function bigIntFromJSON(value: unknown): bigint {
  if (typeof value === "bigint") return value;
  if (typeof value === "number") return BigInt(value);
  if (typeof value === "string") {
    if (/^-?\d+$/.test(value)) return BigInt(value);
    return bigIntFromBytes(bytesFromBase64(value));
  }
  return 0n;
}

export function bigIntFromPartial(value: unknown): bigint {
  if (typeof value === "bigint") return value;
  if (typeof value === "number") return BigInt(value);
  if (typeof value === "string") return value.length ? BigInt(value) : 0n;
  return 0n;
}
