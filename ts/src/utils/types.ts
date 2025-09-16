export type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
    : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
      : { [K in keyof T]?: DeepPartial<T[K]> };

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type Builtin = Date | ((...args: any[]) => any) | Uint8Array | string | number | boolean | undefined;

export type PickByPath<T, Path extends string> = Path extends `${infer First}.${infer Rest}`
  ? First extends keyof T
    ? { [K in First]: PickByPath<T[First], Rest> }
    : never
  : Path extends keyof T
    ? { [K in Path]: T[Path] }
    : never;
