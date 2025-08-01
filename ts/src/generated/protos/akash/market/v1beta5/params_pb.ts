// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/market/v1beta5/params.proto (package akash.market.v1beta5, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Coin, CoinJson } from "../../../cosmos/base/v1beta1/coin_pb.ts";
import { file_cosmos_base_v1beta1_coin } from "../../../cosmos/base/v1beta1/coin_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/market/v1beta5/params.proto.
 */
export const file_akash_market_v1beta5_params: GenFile = /*@__PURE__*/
  fileDesc("CiFha2FzaC9tYXJrZXQvdjFiZXRhNS9wYXJhbXMucHJvdG8SFGFrYXNoLm1hcmtldC52MWJldGE1ItUBCgZQYXJhbXMSdgoPYmlkX21pbl9kZXBvc2l0GAEgASgLMhkuY29zbW9zLmJhc2UudjFiZXRhMS5Db2luQkLI3h8A4t4fDUJpZE1pbkRlcG9zaXTq3h8PYmlkX21pbl9kZXBvc2l08t4fFnlhbWw6ImJpZF9taW5fZGVwb3NpdCISUwoOb3JkZXJfbWF4X2JpZHMYAiABKA1CO+LeHwxPcmRlck1heEJpZHPq3h8Ob3JkZXJfbWF4X2JpZHPy3h8VeWFtbDoib3JkZXJfbWF4X2JpZHMiQiRaInBrZy5ha3QuZGV2L2dvL25vZGUvbWFya2V0L3YxYmV0YTViBnByb3RvMw", [file_gogoproto_gogo, file_cosmos_base_v1beta1_coin]);

/**
 * Params is the params for the x/market module.
 *
 * @generated from message akash.market.v1beta5.Params
 */
export type Params = Message<"akash.market.v1beta5.Params"> & {
  /**
   * BidMinDeposit is a parameter for the minimum deposit on a Bid.
   *
   * @generated from field: cosmos.base.v1beta1.Coin bid_min_deposit = 1;
   */
  bidMinDeposit?: Coin;

  /**
   * OrderMaxBids is a parameter for the maximum number of bids in an order.
   *
   * @generated from field: uint32 order_max_bids = 2;
   */
  orderMaxBids: number;
};

/**
 * Params is the params for the x/market module.
 *
 * @generated from message akash.market.v1beta5.Params
 */
export type ParamsJson = {
  /**
   * BidMinDeposit is a parameter for the minimum deposit on a Bid.
   *
   * @generated from field: cosmos.base.v1beta1.Coin bid_min_deposit = 1;
   */
  bidMinDeposit?: CoinJson;

  /**
   * OrderMaxBids is a parameter for the maximum number of bids in an order.
   *
   * @generated from field: uint32 order_max_bids = 2;
   */
  orderMaxBids?: number;
};

/**
 * Describes the message akash.market.v1beta5.Params.
 * Use `create(ParamsSchema)` to create a new message.
 */
export const ParamsSchema: GenMessage<Params, ParamsJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_params, 0);

