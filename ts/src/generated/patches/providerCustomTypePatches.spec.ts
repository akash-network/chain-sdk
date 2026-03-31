import { DecCoin } from "../protos/cosmos/base/v1beta1/coin.ts";

import { expect, describe, it } from "@jest/globals";
import { patches } from "./providerCustomTypePatches.ts";
import { generateMessage, type MessageSchema } from "@test/helpers/generateMessage";
import type { TypePatches } from "../../sdk/client/types.ts";

const messageTypes: Record<string, MessageSchema> = {
  "cosmos.base.v1beta1.DecCoin": {
    type: DecCoin,
    fields: [{name: "amount",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
};
describe("providerCustomTypePatches.ts", () => {
  describe.each(Object.entries(patches))('patch %s', (typeName, patch: TypePatches[keyof TypePatches]) => {
    it('returns undefined if receives null or undefined', () => {
      expect(patch(null, 'encode')).toBe(undefined);
      expect(patch(null, 'decode')).toBe(undefined);
      expect(patch(undefined, 'encode')).toBe(undefined);
      expect(patch(undefined, 'decode')).toBe(undefined);
    });

    it.each(generateTestCases(typeName, messageTypes))('patches and returns cloned value: %s', (name, value) => {
      const transformedValue = patch(patch(value, 'encode'), 'decode');
      expect(value).toEqual(transformedValue);
      expect(value).not.toBe(transformedValue);
    });
  });

  function generateTestCases(typeName: string, messageTypes: Record<string, MessageSchema>) {
    const type = messageTypes[typeName];
    const cases = type.fields.map((field) => ["single " + field.name + " field", generateMessage(typeName, {
      ...messageTypes,
      [typeName]: {
        ...type,
        fields: [field],
      }
    })]);
    cases.push(["all fields", generateMessage(typeName, messageTypes)]);
    return cases;
  }
});
