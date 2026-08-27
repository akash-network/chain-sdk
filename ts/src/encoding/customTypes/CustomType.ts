export interface CustomType<SourceType, TargetType> {
  typeName: string;
  shortName: string;
  encode(value: SourceType): TargetType;
  decode(value: TargetType): SourceType;
  /**
   * Set when the JS-side representation differs from the wire representation
   * (e.g. a `bigint` surfaced from a `bytes` field). Its value is the
   * TypeScript type text used to override the generated field type and drive
   * asymmetric encode/decode plus JSON overrides in the patch generator.
   *
   * Representation-preserving custom types (e.g. LegacyDec: string<->string)
   * omit this — their generated field type is left untouched.
   */
  jsType?: string;
}
