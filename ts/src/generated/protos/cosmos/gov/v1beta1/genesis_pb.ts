// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/gov/v1beta1/genesis.proto (package cosmos.gov.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Deposit, DepositJson, DepositParams, DepositParamsJson, Proposal, ProposalJson, TallyParams, TallyParamsJson, Vote, VoteJson, VotingParams, VotingParamsJson } from "./gov_pb.ts";
import { file_cosmos_gov_v1beta1_gov } from "./gov_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/gov/v1beta1/genesis.proto.
 */
export const file_cosmos_gov_v1beta1_genesis: GenFile = /*@__PURE__*/
  fileDesc("CiBjb3Ntb3MvZ292L3YxYmV0YTEvZ2VuZXNpcy5wcm90bxISY29zbW9zLmdvdi52MWJldGExIsQDCgxHZW5lc2lzU3RhdGUSHAoUc3RhcnRpbmdfcHJvcG9zYWxfaWQYASABKAQSRAoIZGVwb3NpdHMYAiADKAsyGy5jb3Ntb3MuZ292LnYxYmV0YTEuRGVwb3NpdEIVyN4fAKrfHwhEZXBvc2l0c6jnsCoBEjsKBXZvdGVzGAMgAygLMhguY29zbW9zLmdvdi52MWJldGExLlZvdGVCEsjeHwCq3x8FVm90ZXOo57AqARJHCglwcm9wb3NhbHMYBCADKAsyHC5jb3Ntb3MuZ292LnYxYmV0YTEuUHJvcG9zYWxCFsjeHwCq3x8JUHJvcG9zYWxzqOewKgESRAoOZGVwb3NpdF9wYXJhbXMYBSABKAsyIS5jb3Ntb3MuZ292LnYxYmV0YTEuRGVwb3NpdFBhcmFtc0IJyN4fAKjnsCoBEkIKDXZvdGluZ19wYXJhbXMYBiABKAsyIC5jb3Ntb3MuZ292LnYxYmV0YTEuVm90aW5nUGFyYW1zQgnI3h8AqOewKgESQAoMdGFsbHlfcGFyYW1zGAcgASgLMh8uY29zbW9zLmdvdi52MWJldGExLlRhbGx5UGFyYW1zQgnI3h8AqOewKgFCMlowZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay94L2dvdi90eXBlcy92MWJldGExYgZwcm90bzM", [file_gogoproto_gogo, file_cosmos_gov_v1beta1_gov, file_amino_amino]);

/**
 * GenesisState defines the gov module's genesis state.
 *
 * @generated from message cosmos.gov.v1beta1.GenesisState
 */
export type GenesisState = Message<"cosmos.gov.v1beta1.GenesisState"> & {
  /**
   * starting_proposal_id is the ID of the starting proposal.
   *
   * @generated from field: uint64 starting_proposal_id = 1;
   */
  startingProposalId: bigint;

  /**
   * deposits defines all the deposits present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Deposit deposits = 2;
   */
  deposits: Deposit[];

  /**
   * votes defines all the votes present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Vote votes = 3;
   */
  votes: Vote[];

  /**
   * proposals defines all the proposals present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Proposal proposals = 4;
   */
  proposals: Proposal[];

  /**
   * params defines all the parameters of related to deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.DepositParams deposit_params = 5;
   */
  depositParams?: DepositParams;

  /**
   * params defines all the parameters of related to voting.
   *
   * @generated from field: cosmos.gov.v1beta1.VotingParams voting_params = 6;
   */
  votingParams?: VotingParams;

  /**
   * params defines all the parameters of related to tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyParams tally_params = 7;
   */
  tallyParams?: TallyParams;
};

/**
 * GenesisState defines the gov module's genesis state.
 *
 * @generated from message cosmos.gov.v1beta1.GenesisState
 */
export type GenesisStateJson = {
  /**
   * starting_proposal_id is the ID of the starting proposal.
   *
   * @generated from field: uint64 starting_proposal_id = 1;
   */
  startingProposalId?: string;

  /**
   * deposits defines all the deposits present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Deposit deposits = 2;
   */
  deposits?: DepositJson[];

  /**
   * votes defines all the votes present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Vote votes = 3;
   */
  votes?: VoteJson[];

  /**
   * proposals defines all the proposals present at genesis.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Proposal proposals = 4;
   */
  proposals?: ProposalJson[];

  /**
   * params defines all the parameters of related to deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.DepositParams deposit_params = 5;
   */
  depositParams?: DepositParamsJson;

  /**
   * params defines all the parameters of related to voting.
   *
   * @generated from field: cosmos.gov.v1beta1.VotingParams voting_params = 6;
   */
  votingParams?: VotingParamsJson;

  /**
   * params defines all the parameters of related to tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyParams tally_params = 7;
   */
  tallyParams?: TallyParamsJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.GenesisState.
 * Use `create(GenesisStateSchema)` to create a new message.
 */
export const GenesisStateSchema: GenMessage<GenesisState, GenesisStateJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_genesis, 0);

