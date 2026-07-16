import type { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

import { Timestamp } from "../generated/protos/google/protobuf/timestamp.ts";

export * from "./base64.ts";

// eslint-disable-next-line @typescript-eslint/no-unsafe-function-type
type Builtin = Date | Function | Uint8Array | string | number | bigint | boolean | undefined | null;

export type DeepPartial<T> = T extends bigint
  ? string | number | bigint
  : T extends Builtin
    ? T
    : T extends globalThis.Array<infer U>
      ? globalThis.Array<DeepPartial<U>>
      : T extends ReadonlyArray<infer U>
        ? ReadonlyArray<DeepPartial<U>>
      // eslint-disable-next-line @typescript-eslint/no-empty-object-type
        : T extends {}
          ? { [K in keyof T]?: DeepPartial<T[K]> }
          : Partial<T>;

export type DeepSimplify<T> = T extends bigint
  ? string | number | bigint
  : T extends Builtin
    ? T
    : T extends globalThis.Array<infer U>
      ? globalThis.Array<DeepSimplify<U>>
      : T extends ReadonlyArray<infer U>
        ? ReadonlyArray<DeepSimplify<U>>
        : { [K in keyof T]: DeepSimplify<T[K]> };

export interface MessageFns<T, V extends string> {
  readonly $type: V;
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  fromPartial(object: DeepPartial<T>): T;
}

export function isSet(value: unknown): boolean {
  return value !== null && value !== undefined;
}

export function toTimestamp(date: Date): Timestamp {
  const millis = date.getTime();
  const seconds = BigInt(Math.floor(millis / 1_000));
  const nanos = (millis - Number(seconds) * 1_000) * 1_000_000;
  return { seconds, nanos };
}

export function fromTimestamp(t: Timestamp): Date {
  let millis = Number(t.seconds ?? 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function fromJsonTimestamp(o: any): Date {
  if (o instanceof globalThis.Date) {
    return o;
  } else if (typeof o === "string") {
    return new globalThis.Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

export function isObject(value: unknown): boolean {
  return typeof value === "object" && value !== null;
}
