import { type DescField, type DescMessage, ScalarType } from "@bufbuild/protobuf";
import { BinaryReader } from "@bufbuild/protobuf/wire";

import type { CustomType } from "./CustomType.ts";
import { customTypes } from "./registry.ts";

/**
 * @see https://github.com/gogo/protobuf/blob/master/gogoproto/gogo.proto#L132
 */
const CUSTOM_TYPE_FIELD_NUMBER = 65003;
const parseCustomTypeValue = (value: Uint8Array) => new BinaryReader(value).string();

const customTypeOfFieldCache = new Map<DescField, CustomType<unknown, unknown> | null>();
export function getCustomType(field: DescField): CustomType<unknown, unknown> | null {
  if (!customTypeOfFieldCache.has(field)) {
    const option = field.proto.options?.$unknown?.find((o) => o.no === CUSTOM_TYPE_FIELD_NUMBER);
    const optionValue = option ? parseCustomTypeValue(option.data) : null;
    let type = optionValue && optionValue in customTypes ? customTypes[optionValue as keyof typeof customTypes] : null;
    // Representation-changing custom types (those with a `jsType`, e.g. Int:
    // bytes<->bigint) are only supported on `bytes` fields. The same customtype
    // name (e.g. `cosmossdk.io/math.Int`) is also declared on plenty of string
    // fields (Coin.amount, staking tokens, …); those must be left as generated
    // strings, so we treat them as having no custom type here.
    if (type?.jsType && field.scalar !== ScalarType.BYTES) {
      type = null;
    }
    customTypeOfFieldCache.set(field, type || null);
    return type;
  }

  return customTypeOfFieldCache.get(field) || null;
}

let pathToFieldWithCustomTypeCache: Record<string, DescField[][]> = {};
export function findPathsToCustomField(
  schema: DescMessage,
  shouldVisitNestedMessage: (typeName: string) => boolean = () => true,
  cache: Record<string, DescField[][]> = pathToFieldWithCustomTypeCache,
): DescField[][] {
  if (cache[schema.typeName]) return cache[schema.typeName];

  const paths: DescField[][] = [];
  cache[schema.typeName] = paths; // to prevent cyclic refs

  for (const field of schema.fields) {
    if (field.message && shouldVisitNestedMessage(field.message.typeName)) {
      const nestedPaths = findPathsToCustomField(field.message, shouldVisitNestedMessage, cache);
      nestedPaths.forEach((path) => {
        if (path.length > 0) {
          paths.push([field, ...path]);
        }
      });
    } else if (getCustomType(field)) {
      paths.push([field]);
    }
  }

  return paths;
}

findPathsToCustomField.clearCache = (): void => {
  pathToFieldWithCustomTypeCache = {};
};
