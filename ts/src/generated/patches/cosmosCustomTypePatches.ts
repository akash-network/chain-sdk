import { LegacyDec } from "../../encoding/customTypes/LegacyDec.ts";
import type * as _protos_cosmos_base_v1beta1_coin from "../protos/cosmos/base/v1beta1/coin.ts";
import type * as _protos_cosmos_distribution_v1beta1_distribution from "../protos/cosmos/distribution/v1beta1/distribution.ts";
import type * as _protos_cosmos_gov_v1beta1_gov from "../protos/cosmos/gov/v1beta1/gov.ts";
import { encodeBinary, decodeBinary } from "../../encoding/binaryEncoding.ts";
import type * as _protos_cosmos_mint_v1beta1_mint from "../protos/cosmos/mint/v1beta1/mint.ts";
import type * as _protos_cosmos_mint_v1beta1_query from "../protos/cosmos/mint/v1beta1/query.ts";
import type * as _protos_cosmos_protocolpool_v1_types from "../protos/cosmos/protocolpool/v1/types.ts";
import type * as _protos_cosmos_protocolpool_v1_tx from "../protos/cosmos/protocolpool/v1/tx.ts";
import type * as _protos_cosmos_slashing_v1beta1_slashing from "../protos/cosmos/slashing/v1beta1/slashing.ts";
import type * as _protos_cosmos_staking_v1beta1_staking from "../protos/cosmos/staking/v1beta1/staking.ts";
import type * as _protos_cosmos_staking_v1beta1_tx from "../protos/cosmos/staking/v1beta1/tx.ts";

const p = {
  "cosmos.base.v1beta1.DecCoin"(value: _protos_cosmos_base_v1beta1_coin.DecCoin | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.amount != null) newValue.amount = LegacyDec[transformType](value.amount);
    return newValue;
  },
  "cosmos.base.v1beta1.DecProto"(value: _protos_cosmos_base_v1beta1_coin.DecProto | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.dec != null) newValue.dec = LegacyDec[transformType](value.dec);
    return newValue;
  },
  "cosmos.distribution.v1beta1.Params"(value: _protos_cosmos_distribution_v1beta1_distribution.Params | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.communityTax != null) newValue.communityTax = LegacyDec[transformType](value.communityTax);
    if (value.baseProposerReward != null) newValue.baseProposerReward = LegacyDec[transformType](value.baseProposerReward);
    if (value.bonusProposerReward != null) newValue.bonusProposerReward = LegacyDec[transformType](value.bonusProposerReward);
    return newValue;
  },
  "cosmos.distribution.v1beta1.ValidatorSlashEvent"(value: _protos_cosmos_distribution_v1beta1_distribution.ValidatorSlashEvent | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.fraction != null) newValue.fraction = LegacyDec[transformType](value.fraction);
    return newValue;
  },
  "cosmos.distribution.v1beta1.DelegatorStartingInfo"(value: _protos_cosmos_distribution_v1beta1_distribution.DelegatorStartingInfo | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.stake != null) newValue.stake = LegacyDec[transformType](value.stake);
    return newValue;
  },
  "cosmos.gov.v1beta1.WeightedVoteOption"(value: _protos_cosmos_gov_v1beta1_gov.WeightedVoteOption | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.weight != null) newValue.weight = LegacyDec[transformType](value.weight);
    return newValue;
  },
  "cosmos.gov.v1beta1.TallyParams"(value: _protos_cosmos_gov_v1beta1_gov.TallyParams | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.quorum != null) newValue.quorum = encodeBinary(LegacyDec[transformType](decodeBinary(value.quorum)));
    if (value.threshold != null) newValue.threshold = encodeBinary(LegacyDec[transformType](decodeBinary(value.threshold)));
    if (value.vetoThreshold != null) newValue.vetoThreshold = encodeBinary(LegacyDec[transformType](decodeBinary(value.vetoThreshold)));
    return newValue;
  },
  "cosmos.mint.v1beta1.Minter"(value: _protos_cosmos_mint_v1beta1_mint.Minter | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.inflation != null) newValue.inflation = LegacyDec[transformType](value.inflation);
    if (value.annualProvisions != null) newValue.annualProvisions = LegacyDec[transformType](value.annualProvisions);
    return newValue;
  },
  "cosmos.mint.v1beta1.Params"(value: _protos_cosmos_mint_v1beta1_mint.Params | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.inflationRateChange != null) newValue.inflationRateChange = LegacyDec[transformType](value.inflationRateChange);
    if (value.inflationMax != null) newValue.inflationMax = LegacyDec[transformType](value.inflationMax);
    if (value.inflationMin != null) newValue.inflationMin = LegacyDec[transformType](value.inflationMin);
    if (value.goalBonded != null) newValue.goalBonded = LegacyDec[transformType](value.goalBonded);
    return newValue;
  },
  "cosmos.mint.v1beta1.QueryInflationResponse"(value: _protos_cosmos_mint_v1beta1_query.QueryInflationResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.inflation != null) newValue.inflation = encodeBinary(LegacyDec[transformType](decodeBinary(value.inflation)));
    return newValue;
  },
  "cosmos.mint.v1beta1.QueryAnnualProvisionsResponse"(value: _protos_cosmos_mint_v1beta1_query.QueryAnnualProvisionsResponse | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.annualProvisions != null) newValue.annualProvisions = encodeBinary(LegacyDec[transformType](decodeBinary(value.annualProvisions)));
    return newValue;
  },
  "cosmos.protocolpool.v1.ContinuousFund"(value: _protos_cosmos_protocolpool_v1_types.ContinuousFund | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.percentage != null) newValue.percentage = LegacyDec[transformType](value.percentage);
    return newValue;
  },
  "cosmos.protocolpool.v1.MsgCreateContinuousFund"(value: _protos_cosmos_protocolpool_v1_tx.MsgCreateContinuousFund | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.percentage != null) newValue.percentage = LegacyDec[transformType](value.percentage);
    return newValue;
  },
  "cosmos.slashing.v1beta1.Params"(value: _protos_cosmos_slashing_v1beta1_slashing.Params | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.minSignedPerWindow != null) newValue.minSignedPerWindow = encodeBinary(LegacyDec[transformType](decodeBinary(value.minSignedPerWindow)));
    if (value.slashFractionDoubleSign != null) newValue.slashFractionDoubleSign = encodeBinary(LegacyDec[transformType](decodeBinary(value.slashFractionDoubleSign)));
    if (value.slashFractionDowntime != null) newValue.slashFractionDowntime = encodeBinary(LegacyDec[transformType](decodeBinary(value.slashFractionDowntime)));
    return newValue;
  },
  "cosmos.staking.v1beta1.Validator"(value: _protos_cosmos_staking_v1beta1_staking.Validator | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.delegatorShares != null) newValue.delegatorShares = LegacyDec[transformType](value.delegatorShares);
    return newValue;
  },
  "cosmos.staking.v1beta1.CommissionRates"(value: _protos_cosmos_staking_v1beta1_staking.CommissionRates | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.rate != null) newValue.rate = LegacyDec[transformType](value.rate);
    if (value.maxRate != null) newValue.maxRate = LegacyDec[transformType](value.maxRate);
    if (value.maxChangeRate != null) newValue.maxChangeRate = LegacyDec[transformType](value.maxChangeRate);
    return newValue;
  },
  "cosmos.staking.v1beta1.Delegation"(value: _protos_cosmos_staking_v1beta1_staking.Delegation | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.shares != null) newValue.shares = LegacyDec[transformType](value.shares);
    return newValue;
  },
  "cosmos.staking.v1beta1.RedelegationEntry"(value: _protos_cosmos_staking_v1beta1_staking.RedelegationEntry | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.sharesDst != null) newValue.sharesDst = LegacyDec[transformType](value.sharesDst);
    return newValue;
  },
  "cosmos.staking.v1beta1.Params"(value: _protos_cosmos_staking_v1beta1_staking.Params | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.minCommissionRate != null) newValue.minCommissionRate = LegacyDec[transformType](value.minCommissionRate);
    return newValue;
  },
  "cosmos.staking.v1beta1.MsgEditValidator"(value: _protos_cosmos_staking_v1beta1_tx.MsgEditValidator | undefined | null, transformType: 'encode' | 'decode') {
    if (value == null) return;
    const newValue = { ...value };
    if (value.commissionRate != null) newValue.commissionRate = LegacyDec[transformType](value.commissionRate);
    return newValue;
  }
};

export const patches = p;
