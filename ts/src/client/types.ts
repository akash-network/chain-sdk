import type { DescMethod } from "@bufbuild/protobuf";
import type { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

import type { DeepPartial } from "../utils/types.ts";

export interface ServiceDesc {
  typeName: string;
  methods: Record<string, MethodDesc>;
}

export interface MethodDesc<
  TMethodKind extends DescMethod["methodKind"] | undefined = DescMethod["methodKind"] | undefined,
  TInputMessageDesc extends MessageDesc = MessageDesc,
  TOutputMessageDesc extends MessageDesc = MessageDesc,
> {
  kind?: TMethodKind;
  name: string;
  httpPath?: string;
  httpMethod?: string;
  parent: ServiceDesc;
  input: TInputMessageDesc;
  output: TOutputMessageDesc;
}

export interface MessageDesc<TValue = unknown, TTypeName = string> {
  readonly $type: TTypeName;
  encode(message: TValue, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): TValue;
  fromPartial(object: DeepPartial<TValue>): TValue;
}

export type MessageShape<T> = T extends Pick<MessageDesc, "decode"> ? ReturnType<T["decode"]> : never;

export type MessageInitShape<T> = T extends Pick<MessageDesc, "decode"> ? DeepPartial<ReturnType<T["decode"]>> : never;
