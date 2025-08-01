// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/escrow/v1/account.proto (package akash.escrow.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import type { DecCoin, DecCoinJson } from "../../../cosmos/base/v1beta1/coin_pb.ts";
import { file_cosmos_base_v1beta1_coin } from "../../../cosmos/base/v1beta1/coin_pb.ts";
import type { AccountID, AccountIDJson } from "./accountid_pb.ts";
import { file_akash_escrow_v1_accountid } from "./accountid_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/escrow/v1/account.proto.
 */
export const file_akash_escrow_v1_account: GenFile = /*@__PURE__*/
  fileDesc("Ch1ha2FzaC9lc2Nyb3cvdjEvYWNjb3VudC5wcm90bxIPYWthc2guZXNjcm93LnYxIv8FCgdBY2NvdW50EkUKAmlkGAEgASgLMhouYWthc2guZXNjcm93LnYxLkFjY291bnRJREIdyN4fAOLeHwJJROreHwJpZPLeHwl5YW1sOiJpZCISQAoFb3duZXIYAiABKAlCMereHwVvd25lcvLeHwx5YW1sOiJvd25lciLStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSSAoFc3RhdGUYAyABKA4yHi5ha2FzaC5lc2Nyb3cudjEuQWNjb3VudC5TdGF0ZUIZ6t4fBXN0YXRl8t4fDHlhbWw6InN0YXRlIhJQCgdiYWxhbmNlGAQgASgLMhwuY29zbW9zLmJhc2UudjFiZXRhMS5EZWNDb2luQiHI3h8A6t4fB2JhbGFuY2Xy3h8OeWFtbDoiYmFsYW5jZSISXAoLdHJhbnNmZXJyZWQYBSABKAsyHC5jb3Ntb3MuYmFzZS52MWJldGExLkRlY0NvaW5CKcjeHwDq3h8LdHJhbnNmZXJyZWTy3h8SeWFtbDoidHJhbnNmZXJyZWQiEkIKCnNldHRsZWRfYXQYBiABKANCLuLeHwlTZXR0bGVkQXTq3h8Jc2V0dGxlZEF08t4fEHlhbWw6InNldHRsZWRBdCISTAoJZGVwb3NpdG9yGAcgASgJQjnq3h8JZGVwb3NpdG9y8t4fEHlhbWw6ImRlcG9zaXRvciLStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSSgoFZnVuZHMYCCABKAsyHC5jb3Ntb3MuYmFzZS52MWJldGExLkRlY0NvaW5CHcjeHwDq3h8FZnVuZHPy3h8MeWFtbDoiZnVuZHMiIpIBCgVTdGF0ZRIkCgdpbnZhbGlkEAAaF4qdIBNBY2NvdW50U3RhdGVJbnZhbGlkEhkKBG9wZW4QARoPip0gC0FjY291bnRPcGVuEh0KBmNsb3NlZBACGhGKnSANQWNjb3VudENsb3NlZBIjCglvdmVyZHJhd24QAxoUip0gEEFjY291bnRPdmVyZHJhd24aBIijHgBCH1odcGtnLmFrdC5kZXYvZ28vbm9kZS9lc2Nyb3cvdjFiBnByb3RvMw", [file_gogoproto_gogo, file_cosmos_proto_cosmos, file_cosmos_base_v1beta1_coin, file_akash_escrow_v1_accountid]);

/**
 * Account stores state for an escrow account.
 *
 * @generated from message akash.escrow.v1.Account
 */
export type Account = Message<"akash.escrow.v1.Account"> & {
  /**
   * Id is the unique identifier for an escrow account.
   *
   * @generated from field: akash.escrow.v1.AccountID id = 1;
   */
  id?: AccountID;

  /**
   * Owner is the bech32 address of the account.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 2;
   */
  owner: string;

  /**
   * State represents the current state of an Account.
   *
   * @generated from field: akash.escrow.v1.Account.State state = 3;
   */
  state: Account_State;

  /**
   * Balance holds the unspent coins received from the owner's wallet.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin balance = 4;
   */
  balance?: DecCoin;

  /**
   * Transferred total coins spent by this account.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin transferred = 5;
   */
  transferred?: DecCoin;

  /**
   * SettledAt represents the block height at which this account was last settled.
   *
   * @generated from field: int64 settled_at = 6;
   */
  settledAt: bigint;

  /**
   * Depositor is the bech32 address of the depositor.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   * If depositor is same as the owner, then any incoming coins are added to the Balance.
   * If depositor isn't same as the owner, then any incoming coins are added to the Funds.
   *
   * @generated from field: string depositor = 7;
   */
  depositor: string;

  /**
   * Funds are unspent coins received from the (non-Owner) Depositor's wallet.
   * If there are any funds, they should be spent before spending the Balance.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin funds = 8;
   */
  funds?: DecCoin;
};

/**
 * Account stores state for an escrow account.
 *
 * @generated from message akash.escrow.v1.Account
 */
export type AccountJson = {
  /**
   * Id is the unique identifier for an escrow account.
   *
   * @generated from field: akash.escrow.v1.AccountID id = 1;
   */
  id?: AccountIDJson;

  /**
   * Owner is the bech32 address of the account.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 2;
   */
  owner?: string;

  /**
   * State represents the current state of an Account.
   *
   * @generated from field: akash.escrow.v1.Account.State state = 3;
   */
  state?: Account_StateJson;

  /**
   * Balance holds the unspent coins received from the owner's wallet.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin balance = 4;
   */
  balance?: DecCoinJson;

  /**
   * Transferred total coins spent by this account.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin transferred = 5;
   */
  transferred?: DecCoinJson;

  /**
   * SettledAt represents the block height at which this account was last settled.
   *
   * @generated from field: int64 settled_at = 6;
   */
  settledAt?: string;

  /**
   * Depositor is the bech32 address of the depositor.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   * If depositor is same as the owner, then any incoming coins are added to the Balance.
   * If depositor isn't same as the owner, then any incoming coins are added to the Funds.
   *
   * @generated from field: string depositor = 7;
   */
  depositor?: string;

  /**
   * Funds are unspent coins received from the (non-Owner) Depositor's wallet.
   * If there are any funds, they should be spent before spending the Balance.
   *
   * @generated from field: cosmos.base.v1beta1.DecCoin funds = 8;
   */
  funds?: DecCoinJson;
};

/**
 * Describes the message akash.escrow.v1.Account.
 * Use `create(AccountSchema)` to create a new message.
 */
export const AccountSchema: GenMessage<Account, AccountJson> = /*@__PURE__*/
  messageDesc(file_akash_escrow_v1_account, 0);

/**
 * State stores state for an escrow account.
 *
 * @generated from enum akash.escrow.v1.Account.State
 */
export enum Account_State {
  /**
   * AccountStateInvalid is an invalid state.
   *
   * @generated from enum value: invalid = 0;
   */
  invalid = 0,

  /**
   * AccountOpen is the state when an account is open.
   *
   * @generated from enum value: open = 1;
   */
  open = 1,

  /**
   * AccountClosed is the state when an account is closed.
   *
   * @generated from enum value: closed = 2;
   */
  closed = 2,

  /**
   * AccountOverdrawn is the state when an account is overdrawn.
   *
   * @generated from enum value: overdrawn = 3;
   */
  overdrawn = 3,
}

/**
 * State stores state for an escrow account.
 *
 * @generated from enum akash.escrow.v1.Account.State
 */
export type Account_StateJson = "invalid" | "open" | "closed" | "overdrawn";

/**
 * Describes the enum akash.escrow.v1.Account.State.
 */
export const Account_StateSchema: GenEnum<Account_State, Account_StateJson> = /*@__PURE__*/
  enumDesc(file_akash_escrow_v1_account, 0, 0);

