// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/market/v1beta5/leasemsg.proto (package akash.market.v1beta5, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_cosmos_msg_v1_msg } from "../../../cosmos/msg/v1/msg_pb.ts";
import type { BidID, BidIDJson } from "../v1/bid_pb.ts";
import { file_akash_market_v1_bid } from "../v1/bid_pb.ts";
import type { LeaseID, LeaseIDJson } from "../v1/lease_pb.ts";
import { file_akash_market_v1_lease } from "../v1/lease_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/market/v1beta5/leasemsg.proto.
 */
export const file_akash_market_v1beta5_leasemsg: GenFile = /*@__PURE__*/
  fileDesc("CiNha2FzaC9tYXJrZXQvdjFiZXRhNS9sZWFzZW1zZy5wcm90bxIUYWthc2gubWFya2V0LnYxYmV0YTUicQoOTXNnQ3JlYXRlTGVhc2USSAoGYmlkX2lkGAEgASgLMhYuYWthc2gubWFya2V0LnYxLkJpZElEQiDI3h8A4t4fBUJpZElE6t4fAmlk8t4fCXlhbWw6ImlkIjoV6KAfAILnsCoMYmlkX2lkLm93bmVyIhgKFk1zZ0NyZWF0ZUxlYXNlUmVzcG9uc2UidQoQTXNnV2l0aGRyYXdMZWFzZRJHCgZiaWRfaWQYASABKAsyGC5ha2FzaC5tYXJrZXQudjEuTGVhc2VJREIdyN4fAOLeHwJJROreHwJpZPLeHwl5YW1sOiJpZCI6GOigHwCC57AqD2JpZF9pZC5wcm92aWRlciIaChhNc2dXaXRoZHJhd0xlYXNlUmVzcG9uc2UicwoNTXNnQ2xvc2VMZWFzZRJJCghsZWFzZV9pZBgBIAEoCzIYLmFrYXNoLm1hcmtldC52MS5MZWFzZUlEQh3I3h8A4t4fAklE6t4fAmlk8t4fCXlhbWw6ImlkIjoX6KAfAILnsCoObGVhc2VfaWQub3duZXIiFwoVTXNnQ2xvc2VMZWFzZVJlc3BvbnNlQiRaInBrZy5ha3QuZGV2L2dvL25vZGUvbWFya2V0L3YxYmV0YTViBnByb3RvMw", [file_gogoproto_gogo, file_cosmos_msg_v1_msg, file_akash_market_v1_bid, file_akash_market_v1_lease]);

/**
 * MsgCreateLease is sent to create a lease.
 *
 * @generated from message akash.market.v1beta5.MsgCreateLease
 */
export type MsgCreateLease = Message<"akash.market.v1beta5.MsgCreateLease"> & {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.BidID bid_id = 1;
   */
  bidId?: BidID;
};

/**
 * MsgCreateLease is sent to create a lease.
 *
 * @generated from message akash.market.v1beta5.MsgCreateLease
 */
export type MsgCreateLeaseJson = {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.BidID bid_id = 1;
   */
  bidId?: BidIDJson;
};

/**
 * Describes the message akash.market.v1beta5.MsgCreateLease.
 * Use `create(MsgCreateLeaseSchema)` to create a new message.
 */
export const MsgCreateLeaseSchema: GenMessage<MsgCreateLease, MsgCreateLeaseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 0);

/**
 * MsgCreateLeaseResponse is the response from creating a lease.
 *
 * @generated from message akash.market.v1beta5.MsgCreateLeaseResponse
 */
export type MsgCreateLeaseResponse = Message<"akash.market.v1beta5.MsgCreateLeaseResponse"> & {
};

/**
 * MsgCreateLeaseResponse is the response from creating a lease.
 *
 * @generated from message akash.market.v1beta5.MsgCreateLeaseResponse
 */
export type MsgCreateLeaseResponseJson = {
};

/**
 * Describes the message akash.market.v1beta5.MsgCreateLeaseResponse.
 * Use `create(MsgCreateLeaseResponseSchema)` to create a new message.
 */
export const MsgCreateLeaseResponseSchema: GenMessage<MsgCreateLeaseResponse, MsgCreateLeaseResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 1);

/**
 * MsgWithdrawLease defines an SDK message for withdrawing lease funds.
 *
 * @generated from message akash.market.v1beta5.MsgWithdrawLease
 */
export type MsgWithdrawLease = Message<"akash.market.v1beta5.MsgWithdrawLease"> & {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.LeaseID bid_id = 1;
   */
  bidId?: LeaseID;
};

/**
 * MsgWithdrawLease defines an SDK message for withdrawing lease funds.
 *
 * @generated from message akash.market.v1beta5.MsgWithdrawLease
 */
export type MsgWithdrawLeaseJson = {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.LeaseID bid_id = 1;
   */
  bidId?: LeaseIDJson;
};

/**
 * Describes the message akash.market.v1beta5.MsgWithdrawLease.
 * Use `create(MsgWithdrawLeaseSchema)` to create a new message.
 */
export const MsgWithdrawLeaseSchema: GenMessage<MsgWithdrawLease, MsgWithdrawLeaseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 2);

/**
 * MsgWithdrawLeaseResponse defines the Msg/WithdrawLease response type.
 *
 * @generated from message akash.market.v1beta5.MsgWithdrawLeaseResponse
 */
export type MsgWithdrawLeaseResponse = Message<"akash.market.v1beta5.MsgWithdrawLeaseResponse"> & {
};

/**
 * MsgWithdrawLeaseResponse defines the Msg/WithdrawLease response type.
 *
 * @generated from message akash.market.v1beta5.MsgWithdrawLeaseResponse
 */
export type MsgWithdrawLeaseResponseJson = {
};

/**
 * Describes the message akash.market.v1beta5.MsgWithdrawLeaseResponse.
 * Use `create(MsgWithdrawLeaseResponseSchema)` to create a new message.
 */
export const MsgWithdrawLeaseResponseSchema: GenMessage<MsgWithdrawLeaseResponse, MsgWithdrawLeaseResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 3);

/**
 * MsgCloseLease defines an SDK message for closing order.
 *
 * @generated from message akash.market.v1beta5.MsgCloseLease
 */
export type MsgCloseLease = Message<"akash.market.v1beta5.MsgCloseLease"> & {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.LeaseID lease_id = 1;
   */
  leaseId?: LeaseID;
};

/**
 * MsgCloseLease defines an SDK message for closing order.
 *
 * @generated from message akash.market.v1beta5.MsgCloseLease
 */
export type MsgCloseLeaseJson = {
  /**
   * BidId is the unique identifier of the Bid.
   *
   * @generated from field: akash.market.v1.LeaseID lease_id = 1;
   */
  leaseId?: LeaseIDJson;
};

/**
 * Describes the message akash.market.v1beta5.MsgCloseLease.
 * Use `create(MsgCloseLeaseSchema)` to create a new message.
 */
export const MsgCloseLeaseSchema: GenMessage<MsgCloseLease, MsgCloseLeaseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 4);

/**
 * MsgCloseLeaseResponse defines the Msg/CloseLease response type.
 *
 * @generated from message akash.market.v1beta5.MsgCloseLeaseResponse
 */
export type MsgCloseLeaseResponse = Message<"akash.market.v1beta5.MsgCloseLeaseResponse"> & {
};

/**
 * MsgCloseLeaseResponse defines the Msg/CloseLease response type.
 *
 * @generated from message akash.market.v1beta5.MsgCloseLeaseResponse
 */
export type MsgCloseLeaseResponseJson = {
};

/**
 * Describes the message akash.market.v1beta5.MsgCloseLeaseResponse.
 * Use `create(MsgCloseLeaseResponseSchema)` to create a new message.
 */
export const MsgCloseLeaseResponseSchema: GenMessage<MsgCloseLeaseResponse, MsgCloseLeaseResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_leasemsg, 5);

