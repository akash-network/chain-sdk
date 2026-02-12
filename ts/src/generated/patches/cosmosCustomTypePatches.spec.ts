import { DecCoin, DecProto } from "../protos/cosmos/base/v1beta1/coin.ts";
import { DelegatorStartingInfo, Params, ValidatorSlashEvent } from "../protos/cosmos/distribution/v1beta1/distribution.ts";
import { TallyParams, WeightedVoteOption } from "../protos/cosmos/gov/v1beta1/gov.ts";
import { Minter, Params as Params$1 } from "../protos/cosmos/mint/v1beta1/mint.ts";
import { QueryAnnualProvisionsResponse, QueryInflationResponse } from "../protos/cosmos/mint/v1beta1/query.ts";
import { ContinuousFund } from "../protos/cosmos/protocolpool/v1/types.ts";
import { MsgCreateContinuousFund } from "../protos/cosmos/protocolpool/v1/tx.ts";
import { Params as Params$2 } from "../protos/cosmos/slashing/v1beta1/slashing.ts";
import { CommissionRates, Delegation, Params as Params$3, RedelegationEntry, Validator } from "../protos/cosmos/staking/v1beta1/staking.ts";
import { MsgEditValidator } from "../protos/cosmos/staking/v1beta1/tx.ts";

import { expect, describe, it } from "@jest/globals";
import { patches } from "./cosmosCustomTypePatches.ts";
import { generateMessage, type MessageSchema } from "@test/helpers/generateMessage";
import type { TypePatches } from "../../sdk/client/types.ts";

const messageTypes: Record<string, MessageSchema> = {
  "cosmos.base.v1beta1.DecCoin": {
    type: DecCoin,
    fields: [{name: "amount",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.base.v1beta1.DecProto": {
    type: DecProto,
    fields: [{name: "dec",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.distribution.v1beta1.Params": {
    type: Params,
    fields: [{name: "communityTax",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "baseProposerReward",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "bonusProposerReward",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.distribution.v1beta1.ValidatorSlashEvent": {
    type: ValidatorSlashEvent,
    fields: [{name: "fraction",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.distribution.v1beta1.DelegatorStartingInfo": {
    type: DelegatorStartingInfo,
    fields: [{name: "stake",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.gov.v1beta1.WeightedVoteOption": {
    type: WeightedVoteOption,
    fields: [{name: "weight",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.gov.v1beta1.TallyParams": {
    type: TallyParams,
    fields: [{name: "quorum",kind: "scalar",scalarType: 12,customType: "LegacyDec",},{name: "threshold",kind: "scalar",scalarType: 12,customType: "LegacyDec",},{name: "vetoThreshold",kind: "scalar",scalarType: 12,customType: "LegacyDec",},],
  },
  "cosmos.mint.v1beta1.Minter": {
    type: Minter,
    fields: [{name: "inflation",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "annualProvisions",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.mint.v1beta1.Params": {
    type: Params$1,
    fields: [{name: "inflationRateChange",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "inflationMax",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "inflationMin",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "goalBonded",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.mint.v1beta1.QueryInflationResponse": {
    type: QueryInflationResponse,
    fields: [{name: "inflation",kind: "scalar",scalarType: 12,customType: "LegacyDec",},],
  },
  "cosmos.mint.v1beta1.QueryAnnualProvisionsResponse": {
    type: QueryAnnualProvisionsResponse,
    fields: [{name: "annualProvisions",kind: "scalar",scalarType: 12,customType: "LegacyDec",},],
  },
  "cosmos.protocolpool.v1.ContinuousFund": {
    type: ContinuousFund,
    fields: [{name: "percentage",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.protocolpool.v1.MsgCreateContinuousFund": {
    type: MsgCreateContinuousFund,
    fields: [{name: "percentage",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.slashing.v1beta1.Params": {
    type: Params$2,
    fields: [{name: "minSignedPerWindow",kind: "scalar",scalarType: 12,customType: "LegacyDec",},{name: "slashFractionDoubleSign",kind: "scalar",scalarType: 12,customType: "LegacyDec",},{name: "slashFractionDowntime",kind: "scalar",scalarType: 12,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.Validator": {
    type: Validator,
    fields: [{name: "delegatorShares",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.CommissionRates": {
    type: CommissionRates,
    fields: [{name: "rate",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "maxRate",kind: "scalar",scalarType: 9,customType: "LegacyDec",},{name: "maxChangeRate",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.Delegation": {
    type: Delegation,
    fields: [{name: "shares",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.RedelegationEntry": {
    type: RedelegationEntry,
    fields: [{name: "sharesDst",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.Params": {
    type: Params$3,
    fields: [{name: "minCommissionRate",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
  "cosmos.staking.v1beta1.MsgEditValidator": {
    type: MsgEditValidator,
    fields: [{name: "commissionRate",kind: "scalar",scalarType: 9,customType: "LegacyDec",},],
  },
};
describe("cosmosCustomTypePatches.ts", () => {
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
