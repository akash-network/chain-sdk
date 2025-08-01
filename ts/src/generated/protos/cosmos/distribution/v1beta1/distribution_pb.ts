// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/distribution/v1beta1/distribution.proto (package cosmos.distribution.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Coin, CoinJson, DecCoin, DecCoinJson } from "../../base/v1beta1/coin_pb.ts";
import { file_cosmos_base_v1beta1_coin } from "../../base/v1beta1/coin_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/distribution/v1beta1/distribution.proto.
 */
export const file_cosmos_distribution_v1beta1_distribution: GenFile = /*@__PURE__*/
  fileDesc("Ci5jb3Ntb3MvZGlzdHJpYnV0aW9uL3YxYmV0YTEvZGlzdHJpYnV0aW9uLnByb3RvEhtjb3Ntb3MuZGlzdHJpYnV0aW9uLnYxYmV0YTEi5AIKBlBhcmFtcxJTCg1jb21tdW5pdHlfdGF4GAEgASgJQjzI3h8A2t4fJmdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVj0rQtCmNvc21vcy5EZWMSXAoUYmFzZV9wcm9wb3Nlcl9yZXdhcmQYAiABKAlCPhgByN4fANreHyZnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkRlY9K0LQpjb3Ntb3MuRGVjEl0KFWJvbnVzX3Byb3Bvc2VyX3Jld2FyZBgDIAEoCUI+GAHI3h8A2t4fJmdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVj0rQtCmNvc21vcy5EZWMSHQoVd2l0aGRyYXdfYWRkcl9lbmFibGVkGAQgASgIOimYoB8AiuewKiBjb3Ntb3Mtc2RrL3gvZGlzdHJpYnV0aW9uL1BhcmFtcyKuAQoaVmFsaWRhdG9ySGlzdG9yaWNhbFJld2FyZHMSdwoXY3VtdWxhdGl2ZV9yZXdhcmRfcmF0aW8YASADKAsyHC5jb3Ntb3MuYmFzZS52MWJldGExLkRlY0NvaW5COMjeHwCq3x8rZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay90eXBlcy5EZWNDb2luc6jnsCoBEhcKD3JlZmVyZW5jZV9jb3VudBgCIAEoDSKSAQoXVmFsaWRhdG9yQ3VycmVudFJld2FyZHMSZwoHcmV3YXJkcxgBIAMoCzIcLmNvc21vcy5iYXNlLnYxYmV0YTEuRGVjQ29pbkI4yN4fAKrfHytnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkRlY0NvaW5zqOewKgESDgoGcGVyaW9kGAIgASgEIowBCh5WYWxpZGF0b3JBY2N1bXVsYXRlZENvbW1pc3Npb24SagoKY29tbWlzc2lvbhgBIAMoCzIcLmNvc21vcy5iYXNlLnYxYmV0YTEuRGVjQ29pbkI4yN4fAKrfHytnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkRlY0NvaW5zqOewKgEihgEKG1ZhbGlkYXRvck91dHN0YW5kaW5nUmV3YXJkcxJnCgdyZXdhcmRzGAEgAygLMhwuY29zbW9zLmJhc2UudjFiZXRhMS5EZWNDb2luQjjI3h8Aqt8fK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjQ29pbnOo57AqASJ/ChNWYWxpZGF0b3JTbGFzaEV2ZW50EhgKEHZhbGlkYXRvcl9wZXJpb2QYASABKAQSTgoIZnJhY3Rpb24YAiABKAlCPMjeHwDa3h8mZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay90eXBlcy5EZWPStC0KY29zbW9zLkRlYyJ5ChRWYWxpZGF0b3JTbGFzaEV2ZW50cxJbChZ2YWxpZGF0b3Jfc2xhc2hfZXZlbnRzGAEgAygLMjAuY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLlZhbGlkYXRvclNsYXNoRXZlbnRCCcjeHwCo57AqAToEmKAfACJ5CgdGZWVQb29sEm4KDmNvbW11bml0eV9wb29sGAEgAygLMhwuY29zbW9zLmJhc2UudjFiZXRhMS5EZWNDb2luQjjI3h8Aqt8fK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjQ29pbnOo57AqASKbAQoZVG9rZW5pemVTaGFyZVJlY29yZFJld2FyZBIRCglyZWNvcmRfaWQYASABKAQSYQoGcmV3YXJkGAIgAygLMhwuY29zbW9zLmJhc2UudjFiZXRhMS5EZWNDb2luQjPI3h8Aqt8fK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjQ29pbnM6CIigHwCYoB8BIuMBChpDb21tdW5pdHlQb29sU3BlbmRQcm9wb3NhbBINCgV0aXRsZRgBIAEoCRITCgtkZXNjcmlwdGlvbhgCIAEoCRIRCglyZWNpcGllbnQYAyABKAkSYAoGYW1vdW50GAQgAygLMhkuY29zbW9zLmJhc2UudjFiZXRhMS5Db2luQjXI3h8Aqt8fKGdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuQ29pbnOo57AqATosGAGIoB8AmKAfAOigHwDKtC0aY29zbW9zLmdvdi52MWJldGExLkNvbnRlbnQiuwEKFURlbGVnYXRvclN0YXJ0aW5nSW5mbxIXCg9wcmV2aW91c19wZXJpb2QYASABKAQSSwoFc3Rha2UYAiABKAlCPMjeHwDa3h8mZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay90eXBlcy5EZWPStC0KY29zbW9zLkRlYxI8CgZoZWlnaHQYAyABKARCLOreHw9jcmVhdGlvbl9oZWlnaHSi57AqD2NyZWF0aW9uX2hlaWdodKjnsCoBIsIBChlEZWxlZ2F0aW9uRGVsZWdhdG9yUmV3YXJkEjMKEXZhbGlkYXRvcl9hZGRyZXNzGAEgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSZgoGcmV3YXJkGAIgAygLMhwuY29zbW9zLmJhc2UudjFiZXRhMS5EZWNDb2luQjjI3h8Aqt8fK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjQ29pbnOo57AqAToIiKAfAJigHwEipwEKJUNvbW11bml0eVBvb2xTcGVuZFByb3Bvc2FsV2l0aERlcG9zaXQSDQoFdGl0bGUYASABKAkSEwoLZGVzY3JpcHRpb24YAiABKAkSEQoJcmVjaXBpZW50GAMgASgJEg4KBmFtb3VudBgEIAEoCRIPCgdkZXBvc2l0GAUgASgJOiaIoB8AmKAfAcq0LRpjb3Ntb3MuZ292LnYxYmV0YTEuQ29udGVudEI3WjFnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvZGlzdHJpYnV0aW9uL3R5cGVzqOIeAWIGcHJvdG8z", [file_gogoproto_gogo, file_cosmos_base_v1beta1_coin, file_cosmos_proto_cosmos, file_amino_amino]);

/**
 * Params defines the set of params for the distribution module.
 *
 * @generated from message cosmos.distribution.v1beta1.Params
 */
export type Params = Message<"cosmos.distribution.v1beta1.Params"> & {
  /**
   * @generated from field: string community_tax = 1;
   */
  communityTax: string;

  /**
   * Deprecated: The base_proposer_reward field is deprecated and is no longer used
   * in the x/distribution module's reward mechanism.
   *
   * @generated from field: string base_proposer_reward = 2 [deprecated = true];
   * @deprecated
   */
  baseProposerReward: string;

  /**
   * Deprecated: The bonus_proposer_reward field is deprecated and is no longer used
   * in the x/distribution module's reward mechanism.
   *
   * @generated from field: string bonus_proposer_reward = 3 [deprecated = true];
   * @deprecated
   */
  bonusProposerReward: string;

  /**
   * @generated from field: bool withdraw_addr_enabled = 4;
   */
  withdrawAddrEnabled: boolean;
};

/**
 * Params defines the set of params for the distribution module.
 *
 * @generated from message cosmos.distribution.v1beta1.Params
 */
export type ParamsJson = {
  /**
   * @generated from field: string community_tax = 1;
   */
  communityTax?: string;

  /**
   * Deprecated: The base_proposer_reward field is deprecated and is no longer used
   * in the x/distribution module's reward mechanism.
   *
   * @generated from field: string base_proposer_reward = 2 [deprecated = true];
   * @deprecated
   */
  baseProposerReward?: string;

  /**
   * Deprecated: The bonus_proposer_reward field is deprecated and is no longer used
   * in the x/distribution module's reward mechanism.
   *
   * @generated from field: string bonus_proposer_reward = 3 [deprecated = true];
   * @deprecated
   */
  bonusProposerReward?: string;

  /**
   * @generated from field: bool withdraw_addr_enabled = 4;
   */
  withdrawAddrEnabled?: boolean;
};

/**
 * Describes the message cosmos.distribution.v1beta1.Params.
 * Use `create(ParamsSchema)` to create a new message.
 */
export const ParamsSchema: GenMessage<Params, ParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 0);

/**
 * ValidatorHistoricalRewards represents historical rewards for a validator.
 * Height is implicit within the store key.
 * Cumulative reward ratio is the sum from the zeroeth period
 * until this period of rewards / tokens, per the spec.
 * The reference count indicates the number of objects
 * which might need to reference this historical entry at any point.
 * ReferenceCount =
 *    number of outstanding delegations which ended the associated period (and
 *    might need to read that record)
 *  + number of slashes which ended the associated period (and might need to
 *  read that record)
 *  + one per validator for the zeroeth period, set on initialization
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorHistoricalRewards
 */
export type ValidatorHistoricalRewards = Message<"cosmos.distribution.v1beta1.ValidatorHistoricalRewards"> & {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin cumulative_reward_ratio = 1;
   */
  cumulativeRewardRatio: DecCoin[];

  /**
   * @generated from field: uint32 reference_count = 2;
   */
  referenceCount: number;
};

/**
 * ValidatorHistoricalRewards represents historical rewards for a validator.
 * Height is implicit within the store key.
 * Cumulative reward ratio is the sum from the zeroeth period
 * until this period of rewards / tokens, per the spec.
 * The reference count indicates the number of objects
 * which might need to reference this historical entry at any point.
 * ReferenceCount =
 *    number of outstanding delegations which ended the associated period (and
 *    might need to read that record)
 *  + number of slashes which ended the associated period (and might need to
 *  read that record)
 *  + one per validator for the zeroeth period, set on initialization
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorHistoricalRewards
 */
export type ValidatorHistoricalRewardsJson = {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin cumulative_reward_ratio = 1;
   */
  cumulativeRewardRatio?: DecCoinJson[];

  /**
   * @generated from field: uint32 reference_count = 2;
   */
  referenceCount?: number;
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorHistoricalRewards.
 * Use `create(ValidatorHistoricalRewardsSchema)` to create a new message.
 */
export const ValidatorHistoricalRewardsSchema: GenMessage<ValidatorHistoricalRewards, ValidatorHistoricalRewardsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 1);

/**
 * ValidatorCurrentRewards represents current rewards and current
 * period for a validator kept as a running counter and incremented
 * each block as long as the validator's tokens remain constant.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorCurrentRewards
 */
export type ValidatorCurrentRewards = Message<"cosmos.distribution.v1beta1.ValidatorCurrentRewards"> & {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin rewards = 1;
   */
  rewards: DecCoin[];

  /**
   * @generated from field: uint64 period = 2;
   */
  period: bigint;
};

/**
 * ValidatorCurrentRewards represents current rewards and current
 * period for a validator kept as a running counter and incremented
 * each block as long as the validator's tokens remain constant.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorCurrentRewards
 */
export type ValidatorCurrentRewardsJson = {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin rewards = 1;
   */
  rewards?: DecCoinJson[];

  /**
   * @generated from field: uint64 period = 2;
   */
  period?: string;
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorCurrentRewards.
 * Use `create(ValidatorCurrentRewardsSchema)` to create a new message.
 */
export const ValidatorCurrentRewardsSchema: GenMessage<ValidatorCurrentRewards, ValidatorCurrentRewardsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 2);

/**
 * ValidatorAccumulatedCommission represents accumulated commission
 * for a validator kept as a running counter, can be withdrawn at any time.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
 */
export type ValidatorAccumulatedCommission = Message<"cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"> & {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin commission = 1;
   */
  commission: DecCoin[];
};

/**
 * ValidatorAccumulatedCommission represents accumulated commission
 * for a validator kept as a running counter, can be withdrawn at any time.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
 */
export type ValidatorAccumulatedCommissionJson = {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin commission = 1;
   */
  commission?: DecCoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorAccumulatedCommission.
 * Use `create(ValidatorAccumulatedCommissionSchema)` to create a new message.
 */
export const ValidatorAccumulatedCommissionSchema: GenMessage<ValidatorAccumulatedCommission, ValidatorAccumulatedCommissionJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 3);

/**
 * ValidatorOutstandingRewards represents outstanding (un-withdrawn) rewards
 * for a validator inexpensive to track, allows simple sanity checks.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorOutstandingRewards
 */
export type ValidatorOutstandingRewards = Message<"cosmos.distribution.v1beta1.ValidatorOutstandingRewards"> & {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin rewards = 1;
   */
  rewards: DecCoin[];
};

/**
 * ValidatorOutstandingRewards represents outstanding (un-withdrawn) rewards
 * for a validator inexpensive to track, allows simple sanity checks.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorOutstandingRewards
 */
export type ValidatorOutstandingRewardsJson = {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin rewards = 1;
   */
  rewards?: DecCoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorOutstandingRewards.
 * Use `create(ValidatorOutstandingRewardsSchema)` to create a new message.
 */
export const ValidatorOutstandingRewardsSchema: GenMessage<ValidatorOutstandingRewards, ValidatorOutstandingRewardsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 4);

/**
 * ValidatorSlashEvent represents a validator slash event.
 * Height is implicit within the store key.
 * This is needed to calculate appropriate amount of staking tokens
 * for delegations which are withdrawn after a slash has occurred.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorSlashEvent
 */
export type ValidatorSlashEvent = Message<"cosmos.distribution.v1beta1.ValidatorSlashEvent"> & {
  /**
   * @generated from field: uint64 validator_period = 1;
   */
  validatorPeriod: bigint;

  /**
   * @generated from field: string fraction = 2;
   */
  fraction: string;
};

/**
 * ValidatorSlashEvent represents a validator slash event.
 * Height is implicit within the store key.
 * This is needed to calculate appropriate amount of staking tokens
 * for delegations which are withdrawn after a slash has occurred.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorSlashEvent
 */
export type ValidatorSlashEventJson = {
  /**
   * @generated from field: uint64 validator_period = 1;
   */
  validatorPeriod?: string;

  /**
   * @generated from field: string fraction = 2;
   */
  fraction?: string;
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorSlashEvent.
 * Use `create(ValidatorSlashEventSchema)` to create a new message.
 */
export const ValidatorSlashEventSchema: GenMessage<ValidatorSlashEvent, ValidatorSlashEventJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 5);

/**
 * ValidatorSlashEvents is a collection of ValidatorSlashEvent messages.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorSlashEvents
 */
export type ValidatorSlashEvents = Message<"cosmos.distribution.v1beta1.ValidatorSlashEvents"> & {
  /**
   * @generated from field: repeated cosmos.distribution.v1beta1.ValidatorSlashEvent validator_slash_events = 1;
   */
  validatorSlashEvents: ValidatorSlashEvent[];
};

/**
 * ValidatorSlashEvents is a collection of ValidatorSlashEvent messages.
 *
 * @generated from message cosmos.distribution.v1beta1.ValidatorSlashEvents
 */
export type ValidatorSlashEventsJson = {
  /**
   * @generated from field: repeated cosmos.distribution.v1beta1.ValidatorSlashEvent validator_slash_events = 1;
   */
  validatorSlashEvents?: ValidatorSlashEventJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.ValidatorSlashEvents.
 * Use `create(ValidatorSlashEventsSchema)` to create a new message.
 */
export const ValidatorSlashEventsSchema: GenMessage<ValidatorSlashEvents, ValidatorSlashEventsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 6);

/**
 * FeePool is the global fee pool for distribution.
 *
 * @generated from message cosmos.distribution.v1beta1.FeePool
 */
export type FeePool = Message<"cosmos.distribution.v1beta1.FeePool"> & {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin community_pool = 1;
   */
  communityPool: DecCoin[];
};

/**
 * FeePool is the global fee pool for distribution.
 *
 * @generated from message cosmos.distribution.v1beta1.FeePool
 */
export type FeePoolJson = {
  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin community_pool = 1;
   */
  communityPool?: DecCoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.FeePool.
 * Use `create(FeePoolSchema)` to create a new message.
 */
export const FeePoolSchema: GenMessage<FeePool, FeePoolJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 7);

/**
 * TokenizeShareRecordReward represents the properties of tokenize share
 *
 * @generated from message cosmos.distribution.v1beta1.TokenizeShareRecordReward
 */
export type TokenizeShareRecordReward = Message<"cosmos.distribution.v1beta1.TokenizeShareRecordReward"> & {
  /**
   * @generated from field: uint64 record_id = 1;
   */
  recordId: bigint;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin reward = 2;
   */
  reward: DecCoin[];
};

/**
 * TokenizeShareRecordReward represents the properties of tokenize share
 *
 * @generated from message cosmos.distribution.v1beta1.TokenizeShareRecordReward
 */
export type TokenizeShareRecordRewardJson = {
  /**
   * @generated from field: uint64 record_id = 1;
   */
  recordId?: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin reward = 2;
   */
  reward?: DecCoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.TokenizeShareRecordReward.
 * Use `create(TokenizeShareRecordRewardSchema)` to create a new message.
 */
export const TokenizeShareRecordRewardSchema: GenMessage<TokenizeShareRecordReward, TokenizeShareRecordRewardJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 8);

/**
 * CommunityPoolSpendProposal details a proposal for use of community funds,
 * together with how many coins are proposed to be spent, and to which
 * recipient account.
 *
 * Deprecated: Do not use. As of the Cosmos SDK release v0.47.x, there is no
 * longer a need for an explicit CommunityPoolSpendProposal. To spend community
 * pool funds, a simple MsgCommunityPoolSpend can be invoked from the x/gov
 * module via a v1 governance proposal.
 *
 * @generated from message cosmos.distribution.v1beta1.CommunityPoolSpendProposal
 * @deprecated
 */
export type CommunityPoolSpendProposal = Message<"cosmos.distribution.v1beta1.CommunityPoolSpendProposal"> & {
  /**
   * @generated from field: string title = 1;
   */
  title: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string recipient = 3;
   */
  recipient: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 4;
   */
  amount: Coin[];
};

/**
 * CommunityPoolSpendProposal details a proposal for use of community funds,
 * together with how many coins are proposed to be spent, and to which
 * recipient account.
 *
 * Deprecated: Do not use. As of the Cosmos SDK release v0.47.x, there is no
 * longer a need for an explicit CommunityPoolSpendProposal. To spend community
 * pool funds, a simple MsgCommunityPoolSpend can be invoked from the x/gov
 * module via a v1 governance proposal.
 *
 * @generated from message cosmos.distribution.v1beta1.CommunityPoolSpendProposal
 * @deprecated
 */
export type CommunityPoolSpendProposalJson = {
  /**
   * @generated from field: string title = 1;
   */
  title?: string;

  /**
   * @generated from field: string description = 2;
   */
  description?: string;

  /**
   * @generated from field: string recipient = 3;
   */
  recipient?: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 4;
   */
  amount?: CoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.CommunityPoolSpendProposal.
 * Use `create(CommunityPoolSpendProposalSchema)` to create a new message.
 * @deprecated
 */
export const CommunityPoolSpendProposalSchema: GenMessage<CommunityPoolSpendProposal, CommunityPoolSpendProposalJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 9);

/**
 * DelegatorStartingInfo represents the starting info for a delegator reward
 * period. It tracks the previous validator period, the delegation's amount of
 * staking token, and the creation height (to check later on if any slashes have
 * occurred). NOTE: Even though validators are slashed to whole staking tokens,
 * the delegators within the validator may be left with less than a full token,
 * thus sdk.Dec is used.
 *
 * @generated from message cosmos.distribution.v1beta1.DelegatorStartingInfo
 */
export type DelegatorStartingInfo = Message<"cosmos.distribution.v1beta1.DelegatorStartingInfo"> & {
  /**
   * @generated from field: uint64 previous_period = 1;
   */
  previousPeriod: bigint;

  /**
   * @generated from field: string stake = 2;
   */
  stake: string;

  /**
   * @generated from field: uint64 height = 3;
   */
  height: bigint;
};

/**
 * DelegatorStartingInfo represents the starting info for a delegator reward
 * period. It tracks the previous validator period, the delegation's amount of
 * staking token, and the creation height (to check later on if any slashes have
 * occurred). NOTE: Even though validators are slashed to whole staking tokens,
 * the delegators within the validator may be left with less than a full token,
 * thus sdk.Dec is used.
 *
 * @generated from message cosmos.distribution.v1beta1.DelegatorStartingInfo
 */
export type DelegatorStartingInfoJson = {
  /**
   * @generated from field: uint64 previous_period = 1;
   */
  previousPeriod?: string;

  /**
   * @generated from field: string stake = 2;
   */
  stake?: string;

  /**
   * @generated from field: uint64 height = 3;
   */
  height?: string;
};

/**
 * Describes the message cosmos.distribution.v1beta1.DelegatorStartingInfo.
 * Use `create(DelegatorStartingInfoSchema)` to create a new message.
 */
export const DelegatorStartingInfoSchema: GenMessage<DelegatorStartingInfo, DelegatorStartingInfoJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 10);

/**
 * DelegationDelegatorReward represents the properties
 * of a delegator's delegation reward.
 *
 * @generated from message cosmos.distribution.v1beta1.DelegationDelegatorReward
 */
export type DelegationDelegatorReward = Message<"cosmos.distribution.v1beta1.DelegationDelegatorReward"> & {
  /**
   * @generated from field: string validator_address = 1;
   */
  validatorAddress: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin reward = 2;
   */
  reward: DecCoin[];
};

/**
 * DelegationDelegatorReward represents the properties
 * of a delegator's delegation reward.
 *
 * @generated from message cosmos.distribution.v1beta1.DelegationDelegatorReward
 */
export type DelegationDelegatorRewardJson = {
  /**
   * @generated from field: string validator_address = 1;
   */
  validatorAddress?: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.DecCoin reward = 2;
   */
  reward?: DecCoinJson[];
};

/**
 * Describes the message cosmos.distribution.v1beta1.DelegationDelegatorReward.
 * Use `create(DelegationDelegatorRewardSchema)` to create a new message.
 */
export const DelegationDelegatorRewardSchema: GenMessage<DelegationDelegatorReward, DelegationDelegatorRewardJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 11);

/**
 * CommunityPoolSpendProposalWithDeposit defines a CommunityPoolSpendProposal
 * with a deposit
 *
 * @generated from message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit
 */
export type CommunityPoolSpendProposalWithDeposit = Message<"cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"> & {
  /**
   * @generated from field: string title = 1;
   */
  title: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string recipient = 3;
   */
  recipient: string;

  /**
   * @generated from field: string amount = 4;
   */
  amount: string;

  /**
   * @generated from field: string deposit = 5;
   */
  deposit: string;
};

/**
 * CommunityPoolSpendProposalWithDeposit defines a CommunityPoolSpendProposal
 * with a deposit
 *
 * @generated from message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit
 */
export type CommunityPoolSpendProposalWithDepositJson = {
  /**
   * @generated from field: string title = 1;
   */
  title?: string;

  /**
   * @generated from field: string description = 2;
   */
  description?: string;

  /**
   * @generated from field: string recipient = 3;
   */
  recipient?: string;

  /**
   * @generated from field: string amount = 4;
   */
  amount?: string;

  /**
   * @generated from field: string deposit = 5;
   */
  deposit?: string;
};

/**
 * Describes the message cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit.
 * Use `create(CommunityPoolSpendProposalWithDepositSchema)` to create a new message.
 */
export const CommunityPoolSpendProposalWithDepositSchema: GenMessage<CommunityPoolSpendProposalWithDeposit, CommunityPoolSpendProposalWithDepositJson> = /*@__PURE__*/
  messageDesc(file_cosmos_distribution_v1beta1_distribution, 12);

