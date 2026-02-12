import { CoinPrice, CollateralRatio } from "../protos/akash/bme/v1/types.ts";
import { EventMintStatusChange } from "../protos/akash/bme/v1/events.ts";
import { QueryStatusResponse } from "../protos/akash/bme/v1/query.ts";
import { DecCoin } from "../protos/cosmos/base/v1beta1/coin.ts";
import { Balance } from "../protos/akash/escrow/types/v1/balance.ts";
import { AggregatedPrice, PriceDataState } from "../protos/akash/oracle/v1/prices.ts";

import { expect, describe, it } from "@jest/globals";
import { patches } from "./nodeCustomTypePatches.ts";
import { generateMessage, type MessageSchema } from "@test/helpers/generateMessage";
import type { TypePatches } from "../../sdk/client/types.ts";

const messageTypes: Record<string, MessageSchema> = {
  "akash.bme.v1.CollateralRatio": {
    type: CollateralRatio,
    fields: [{name: "ratio",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "referencePrice",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.bme.v1.CoinPrice": {
    type: CoinPrice,
    fields: [{name: "price",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.bme.v1.EventMintStatusChange": {
    type: EventMintStatusChange,
    fields: [{name: "collateralRatio",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.bme.v1.QueryStatusResponse": {
    type: QueryStatusResponse,
    fields: [{name: "collateralRatio",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "warnThreshold",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "haltThreshold",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.base.v1beta1.DecCoin": {
    type: DecCoin,
    fields: [{name: "amount",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.escrow.types.v1.Balance": {
    type: Balance,
    fields: [{name: "amount",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.oracle.v1.PriceDataState": {
    type: PriceDataState,
    fields: [{name: "price",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "akash.oracle.v1.AggregatedPrice": {
    type: AggregatedPrice,
    fields: [{name: "twap",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "medianPrice",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "minPrice",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "maxPrice",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
};
describe("nodeCustomTypePatches.ts", () => {
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
