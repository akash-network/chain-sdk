// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/market/v1beta5/query.proto (package akash.market.v1beta5, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import { file_google_api_annotations } from "../../../google/api/annotations_pb.ts";
import type { PageRequest, PageRequestJson, PageResponse, PageResponseJson } from "../../../cosmos/base/query/v1beta1/pagination_pb.ts";
import { file_cosmos_base_query_v1beta1_pagination } from "../../../cosmos/base/query/v1beta1/pagination_pb.ts";
import type { Account, AccountJson } from "../../escrow/v1/account_pb.ts";
import { file_akash_escrow_v1_account } from "../../escrow/v1/account_pb.ts";
import type { FractionalPayment, FractionalPaymentJson } from "../../escrow/v1/fractional_payment_pb.ts";
import { file_akash_escrow_v1_fractional_payment } from "../../escrow/v1/fractional_payment_pb.ts";
import type { OrderID, OrderIDJson } from "../v1/order_pb.ts";
import { file_akash_market_v1_order } from "../v1/order_pb.ts";
import type { BidID, BidIDJson } from "../v1/bid_pb.ts";
import { file_akash_market_v1_bid } from "../v1/bid_pb.ts";
import type { Lease, LeaseID, LeaseIDJson, LeaseJson } from "../v1/lease_pb.ts";
import { file_akash_market_v1_lease } from "../v1/lease_pb.ts";
import type { LeaseFilters, LeaseFiltersJson } from "../v1/filters_pb.ts";
import { file_akash_market_v1_filters } from "../v1/filters_pb.ts";
import type { Order, OrderJson } from "./order_pb.ts";
import { file_akash_market_v1beta5_order } from "./order_pb.ts";
import type { Bid, BidJson } from "./bid_pb.ts";
import { file_akash_market_v1beta5_bid } from "./bid_pb.ts";
import type { BidFilters, BidFiltersJson, OrderFilters, OrderFiltersJson } from "./filters_pb.ts";
import { file_akash_market_v1beta5_filters } from "./filters_pb.ts";
import type { Params, ParamsJson } from "./params_pb.ts";
import { file_akash_market_v1beta5_params } from "./params_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/market/v1beta5/query.proto.
 */
export const file_akash_market_v1beta5_query: GenFile = /*@__PURE__*/
  fileDesc("CiBha2FzaC9tYXJrZXQvdjFiZXRhNS9xdWVyeS5wcm90bxIUYWthc2gubWFya2V0LnYxYmV0YTUiiwEKElF1ZXJ5T3JkZXJzUmVxdWVzdBI5CgdmaWx0ZXJzGAEgASgLMiIuYWthc2gubWFya2V0LnYxYmV0YTUuT3JkZXJGaWx0ZXJzQgTI3h8AEjoKCnBhZ2luYXRpb24YAiABKAsyJi5jb3Ntb3MuYmFzZS5xdWVyeS52MWJldGExLlBhZ2VSZXF1ZXN0Io8BChNRdWVyeU9yZGVyc1Jlc3BvbnNlEjsKBm9yZGVycxgBIAMoCzIbLmFrYXNoLm1hcmtldC52MWJldGE1Lk9yZGVyQg7I3h8Aqt8fBk9yZGVycxI7CgpwYWdpbmF0aW9uGAIgASgLMicuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVzcG9uc2UiRQoRUXVlcnlPcmRlclJlcXVlc3QSMAoCaWQYASABKAsyGC5ha2FzaC5tYXJrZXQudjEuT3JkZXJJREIKyN4fAOLeHwJJRCJGChJRdWVyeU9yZGVyUmVzcG9uc2USMAoFb3JkZXIYASABKAsyGy5ha2FzaC5tYXJrZXQudjFiZXRhNS5PcmRlckIEyN4fACKHAQoQUXVlcnlCaWRzUmVxdWVzdBI3CgdmaWx0ZXJzGAEgASgLMiAuYWthc2gubWFya2V0LnYxYmV0YTUuQmlkRmlsdGVyc0IEyN4fABI6CgpwYWdpbmF0aW9uGAIgASgLMiYuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVxdWVzdCKMAQoRUXVlcnlCaWRzUmVzcG9uc2USOgoEYmlkcxgBIAMoCzImLmFrYXNoLm1hcmtldC52MWJldGE1LlF1ZXJ5QmlkUmVzcG9uc2VCBMjeHwASOwoKcGFnaW5hdGlvbhgCIAEoCzInLmNvc21vcy5iYXNlLnF1ZXJ5LnYxYmV0YTEuUGFnZVJlc3BvbnNlIkEKD1F1ZXJ5QmlkUmVxdWVzdBIuCgJpZBgBIAEoCzIWLmFrYXNoLm1hcmtldC52MS5CaWRJREIKyN4fAOLeHwJJRCJ4ChBRdWVyeUJpZFJlc3BvbnNlEiwKA2JpZBgBIAEoCzIZLmFrYXNoLm1hcmtldC52MWJldGE1LkJpZEIEyN4fABI2Cg5lc2Nyb3dfYWNjb3VudBgCIAEoCzIYLmFrYXNoLmVzY3Jvdy52MS5BY2NvdW50QgTI3h8AIoYBChJRdWVyeUxlYXNlc1JlcXVlc3QSNAoHZmlsdGVycxgBIAEoCzIdLmFrYXNoLm1hcmtldC52MS5MZWFzZUZpbHRlcnNCBMjeHwASOgoKcGFnaW5hdGlvbhgCIAEoCzImLmNvc21vcy5iYXNlLnF1ZXJ5LnYxYmV0YTEuUGFnZVJlcXVlc3QikgEKE1F1ZXJ5TGVhc2VzUmVzcG9uc2USPgoGbGVhc2VzGAEgAygLMiguYWthc2gubWFya2V0LnYxYmV0YTUuUXVlcnlMZWFzZVJlc3BvbnNlQgTI3h8AEjsKCnBhZ2luYXRpb24YAiABKAsyJy5jb3Ntb3MuYmFzZS5xdWVyeS52MWJldGExLlBhZ2VSZXNwb25zZSJFChFRdWVyeUxlYXNlUmVxdWVzdBIwCgJpZBgBIAEoCzIYLmFrYXNoLm1hcmtldC52MS5MZWFzZUlEQgrI3h8A4t4fAklEIoMBChJRdWVyeUxlYXNlUmVzcG9uc2USKwoFbGVhc2UYASABKAsyFi5ha2FzaC5tYXJrZXQudjEuTGVhc2VCBMjeHwASQAoOZXNjcm93X3BheW1lbnQYAiABKAsyIi5ha2FzaC5lc2Nyb3cudjEuRnJhY3Rpb25hbFBheW1lbnRCBMjeHwAiFAoSUXVlcnlQYXJhbXNSZXF1ZXN0Ik4KE1F1ZXJ5UGFyYW1zUmVzcG9uc2USNwoGcGFyYW1zGAEgASgLMhwuYWthc2gubWFya2V0LnYxYmV0YTUuUGFyYW1zQgnI3h8AqOewKgEytQcKBVF1ZXJ5EogBCgZPcmRlcnMSKC5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeU9yZGVyc1JlcXVlc3QaKS5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeU9yZGVyc1Jlc3BvbnNlIimC0+STAiMSIS9ha2FzaC9tYXJrZXQvdjFiZXRhNS9vcmRlcnMvbGlzdBKFAQoFT3JkZXISJy5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeU9yZGVyUmVxdWVzdBooLmFrYXNoLm1hcmtldC52MWJldGE1LlF1ZXJ5T3JkZXJSZXNwb25zZSIpgtPkkwIjEiEvYWthc2gvbWFya2V0L3YxYmV0YTUvb3JkZXJzL2luZm8SgAEKBEJpZHMSJi5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeUJpZHNSZXF1ZXN0GicuYWthc2gubWFya2V0LnYxYmV0YTUuUXVlcnlCaWRzUmVzcG9uc2UiJ4LT5JMCIRIfL2FrYXNoL21hcmtldC92MWJldGE1L2JpZHMvbGlzdBJ9CgNCaWQSJS5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeUJpZFJlcXVlc3QaJi5ha2FzaC5tYXJrZXQudjFiZXRhNS5RdWVyeUJpZFJlc3BvbnNlIieC0+STAiESHy9ha2FzaC9tYXJrZXQvdjFiZXRhNS9iaWRzL2luZm8SiAEKBkxlYXNlcxIoLmFrYXNoLm1hcmtldC52MWJldGE1LlF1ZXJ5TGVhc2VzUmVxdWVzdBopLmFrYXNoLm1hcmtldC52MWJldGE1LlF1ZXJ5TGVhc2VzUmVzcG9uc2UiKYLT5JMCIxIhL2FrYXNoL21hcmtldC92MWJldGE1L2xlYXNlcy9saXN0EoUBCgVMZWFzZRInLmFrYXNoLm1hcmtldC52MWJldGE1LlF1ZXJ5TGVhc2VSZXF1ZXN0GiguYWthc2gubWFya2V0LnYxYmV0YTUuUXVlcnlMZWFzZVJlc3BvbnNlIimC0+STAiMSIS9ha2FzaC9tYXJrZXQvdjFiZXRhNS9sZWFzZXMvaW5mbxKDAQoGUGFyYW1zEiguYWthc2gubWFya2V0LnYxYmV0YTUuUXVlcnlQYXJhbXNSZXF1ZXN0GikuYWthc2gubWFya2V0LnYxYmV0YTUuUXVlcnlQYXJhbXNSZXNwb25zZSIkgtPkkwIeEhwvYWthc2gvbWFya2V0L3YxYmV0YTUvcGFyYW1zQiRaInBrZy5ha3QuZGV2L2dvL25vZGUvbWFya2V0L3YxYmV0YTViBnByb3RvMw", [file_gogoproto_gogo, file_amino_amino, file_google_api_annotations, file_cosmos_base_query_v1beta1_pagination, file_akash_escrow_v1_account, file_akash_escrow_v1_fractional_payment, file_akash_market_v1_order, file_akash_market_v1_bid, file_akash_market_v1_lease, file_akash_market_v1_filters, file_akash_market_v1beta5_order, file_akash_market_v1beta5_bid, file_akash_market_v1beta5_filters, file_akash_market_v1beta5_params]);

/**
 * QueryOrdersRequest is request type for the Query/Orders RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrdersRequest
 */
export type QueryOrdersRequest = Message<"akash.market.v1beta5.QueryOrdersRequest"> & {
  /**
   * Filters holds the fields to filter orders.
   *
   * @generated from field: akash.market.v1beta5.OrderFilters filters = 1;
   */
  filters?: OrderFilters;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;
};

/**
 * QueryOrdersRequest is request type for the Query/Orders RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrdersRequest
 */
export type QueryOrdersRequestJson = {
  /**
   * Filters holds the fields to filter orders.
   *
   * @generated from field: akash.market.v1beta5.OrderFilters filters = 1;
   */
  filters?: OrderFiltersJson;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryOrdersRequest.
 * Use `create(QueryOrdersRequestSchema)` to create a new message.
 */
export const QueryOrdersRequestSchema: GenMessage<QueryOrdersRequest, QueryOrdersRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 0);

/**
 * QueryOrdersResponse is response type for the Query/Orders RPC method
 *
 * @generated from message akash.market.v1beta5.QueryOrdersResponse
 */
export type QueryOrdersResponse = Message<"akash.market.v1beta5.QueryOrdersResponse"> & {
  /**
   * Orders is a list of market orders.
   *
   * @generated from field: repeated akash.market.v1beta5.Order orders = 1;
   */
  orders: Order[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryOrdersResponse is response type for the Query/Orders RPC method
 *
 * @generated from message akash.market.v1beta5.QueryOrdersResponse
 */
export type QueryOrdersResponseJson = {
  /**
   * Orders is a list of market orders.
   *
   * @generated from field: repeated akash.market.v1beta5.Order orders = 1;
   */
  orders?: OrderJson[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryOrdersResponse.
 * Use `create(QueryOrdersResponseSchema)` to create a new message.
 */
export const QueryOrdersResponseSchema: GenMessage<QueryOrdersResponse, QueryOrdersResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 1);

/**
 * QueryOrderRequest is request type for the Query/Order RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrderRequest
 */
export type QueryOrderRequest = Message<"akash.market.v1beta5.QueryOrderRequest"> & {
  /**
   * Id is the unique identifier of the Order.
   *
   * @generated from field: akash.market.v1.OrderID id = 1;
   */
  id?: OrderID;
};

/**
 * QueryOrderRequest is request type for the Query/Order RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrderRequest
 */
export type QueryOrderRequestJson = {
  /**
   * Id is the unique identifier of the Order.
   *
   * @generated from field: akash.market.v1.OrderID id = 1;
   */
  id?: OrderIDJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryOrderRequest.
 * Use `create(QueryOrderRequestSchema)` to create a new message.
 */
export const QueryOrderRequestSchema: GenMessage<QueryOrderRequest, QueryOrderRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 2);

/**
 * QueryOrderResponse is response type for the Query/Order RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrderResponse
 */
export type QueryOrderResponse = Message<"akash.market.v1beta5.QueryOrderResponse"> & {
  /**
   * Order represents a market order.
   *
   * @generated from field: akash.market.v1beta5.Order order = 1;
   */
  order?: Order;
};

/**
 * QueryOrderResponse is response type for the Query/Order RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryOrderResponse
 */
export type QueryOrderResponseJson = {
  /**
   * Order represents a market order.
   *
   * @generated from field: akash.market.v1beta5.Order order = 1;
   */
  order?: OrderJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryOrderResponse.
 * Use `create(QueryOrderResponseSchema)` to create a new message.
 */
export const QueryOrderResponseSchema: GenMessage<QueryOrderResponse, QueryOrderResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 3);

/**
 * QueryBidsRequest is request type for the Query/Bids RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidsRequest
 */
export type QueryBidsRequest = Message<"akash.market.v1beta5.QueryBidsRequest"> & {
  /**
   * Filters holds the fields to filter bids.
   *
   * @generated from field: akash.market.v1beta5.BidFilters filters = 1;
   */
  filters?: BidFilters;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;
};

/**
 * QueryBidsRequest is request type for the Query/Bids RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidsRequest
 */
export type QueryBidsRequestJson = {
  /**
   * Filters holds the fields to filter bids.
   *
   * @generated from field: akash.market.v1beta5.BidFilters filters = 1;
   */
  filters?: BidFiltersJson;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryBidsRequest.
 * Use `create(QueryBidsRequestSchema)` to create a new message.
 */
export const QueryBidsRequestSchema: GenMessage<QueryBidsRequest, QueryBidsRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 4);

/**
 * QueryBidsResponse is response type for the Query/Bids RPC method
 *
 * @generated from message akash.market.v1beta5.QueryBidsResponse
 */
export type QueryBidsResponse = Message<"akash.market.v1beta5.QueryBidsResponse"> & {
  /**
   * Bids is a list of deployment bids.
   *
   * @generated from field: repeated akash.market.v1beta5.QueryBidResponse bids = 1;
   */
  bids: QueryBidResponse[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryBidsResponse is response type for the Query/Bids RPC method
 *
 * @generated from message akash.market.v1beta5.QueryBidsResponse
 */
export type QueryBidsResponseJson = {
  /**
   * Bids is a list of deployment bids.
   *
   * @generated from field: repeated akash.market.v1beta5.QueryBidResponse bids = 1;
   */
  bids?: QueryBidResponseJson[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryBidsResponse.
 * Use `create(QueryBidsResponseSchema)` to create a new message.
 */
export const QueryBidsResponseSchema: GenMessage<QueryBidsResponse, QueryBidsResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 5);

/**
 * QueryBidRequest is request type for the Query/Bid RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidRequest
 */
export type QueryBidRequest = Message<"akash.market.v1beta5.QueryBidRequest"> & {
  /**
   * Id is the unique identifier for the Bid.
   *
   * @generated from field: akash.market.v1.BidID id = 1;
   */
  id?: BidID;
};

/**
 * QueryBidRequest is request type for the Query/Bid RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidRequest
 */
export type QueryBidRequestJson = {
  /**
   * Id is the unique identifier for the Bid.
   *
   * @generated from field: akash.market.v1.BidID id = 1;
   */
  id?: BidIDJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryBidRequest.
 * Use `create(QueryBidRequestSchema)` to create a new message.
 */
export const QueryBidRequestSchema: GenMessage<QueryBidRequest, QueryBidRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 6);

/**
 * QueryBidResponse is response type for the Query/Bid RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidResponse
 */
export type QueryBidResponse = Message<"akash.market.v1beta5.QueryBidResponse"> & {
  /**
   * Bid represents a deployment bid.
   *
   * @generated from field: akash.market.v1beta5.Bid bid = 1;
   */
  bid?: Bid;

  /**
   * EscrowAccount represents the escrow account created for the Bid.
   *
   * @generated from field: akash.escrow.v1.Account escrow_account = 2;
   */
  escrowAccount?: Account;
};

/**
 * QueryBidResponse is response type for the Query/Bid RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryBidResponse
 */
export type QueryBidResponseJson = {
  /**
   * Bid represents a deployment bid.
   *
   * @generated from field: akash.market.v1beta5.Bid bid = 1;
   */
  bid?: BidJson;

  /**
   * EscrowAccount represents the escrow account created for the Bid.
   *
   * @generated from field: akash.escrow.v1.Account escrow_account = 2;
   */
  escrowAccount?: AccountJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryBidResponse.
 * Use `create(QueryBidResponseSchema)` to create a new message.
 */
export const QueryBidResponseSchema: GenMessage<QueryBidResponse, QueryBidResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 7);

/**
 * QueryLeasesRequest is request type for the Query/Leases RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeasesRequest
 */
export type QueryLeasesRequest = Message<"akash.market.v1beta5.QueryLeasesRequest"> & {
  /**
   * Filters holds the fields to filter leases.
   *
   * @generated from field: akash.market.v1.LeaseFilters filters = 1;
   */
  filters?: LeaseFilters;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;
};

/**
 * QueryLeasesRequest is request type for the Query/Leases RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeasesRequest
 */
export type QueryLeasesRequestJson = {
  /**
   * Filters holds the fields to filter leases.
   *
   * @generated from field: akash.market.v1.LeaseFilters filters = 1;
   */
  filters?: LeaseFiltersJson;

  /**
   * Pagination is used to paginate request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryLeasesRequest.
 * Use `create(QueryLeasesRequestSchema)` to create a new message.
 */
export const QueryLeasesRequestSchema: GenMessage<QueryLeasesRequest, QueryLeasesRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 8);

/**
 * QueryLeasesResponse is response type for the Query/Leases RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeasesResponse
 */
export type QueryLeasesResponse = Message<"akash.market.v1beta5.QueryLeasesResponse"> & {
  /**
   * Leases is a list of Lease.
   *
   * @generated from field: repeated akash.market.v1beta5.QueryLeaseResponse leases = 1;
   */
  leases: QueryLeaseResponse[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryLeasesResponse is response type for the Query/Leases RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeasesResponse
 */
export type QueryLeasesResponseJson = {
  /**
   * Leases is a list of Lease.
   *
   * @generated from field: repeated akash.market.v1beta5.QueryLeaseResponse leases = 1;
   */
  leases?: QueryLeaseResponseJson[];

  /**
   * Pagination contains the information about response pagination.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryLeasesResponse.
 * Use `create(QueryLeasesResponseSchema)` to create a new message.
 */
export const QueryLeasesResponseSchema: GenMessage<QueryLeasesResponse, QueryLeasesResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 9);

/**
 * QueryLeaseRequest is request type for the Query/Lease RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeaseRequest
 */
export type QueryLeaseRequest = Message<"akash.market.v1beta5.QueryLeaseRequest"> & {
  /**
   * Id is the unique identifier of the Lease.
   *
   * @generated from field: akash.market.v1.LeaseID id = 1;
   */
  id?: LeaseID;
};

/**
 * QueryLeaseRequest is request type for the Query/Lease RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryLeaseRequest
 */
export type QueryLeaseRequestJson = {
  /**
   * Id is the unique identifier of the Lease.
   *
   * @generated from field: akash.market.v1.LeaseID id = 1;
   */
  id?: LeaseIDJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryLeaseRequest.
 * Use `create(QueryLeaseRequestSchema)` to create a new message.
 */
export const QueryLeaseRequestSchema: GenMessage<QueryLeaseRequest, QueryLeaseRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 10);

/**
 * QueryLeaseResponse is response type for the Query/Lease RPC method
 *
 * @generated from message akash.market.v1beta5.QueryLeaseResponse
 */
export type QueryLeaseResponse = Message<"akash.market.v1beta5.QueryLeaseResponse"> & {
  /**
   * Lease holds the lease for a deployment.
   *
   * @generated from field: akash.market.v1.Lease lease = 1;
   */
  lease?: Lease;

  /**
   * EscrowPayment holds information about the Lease's fractional payment.
   *
   * @generated from field: akash.escrow.v1.FractionalPayment escrow_payment = 2;
   */
  escrowPayment?: FractionalPayment;
};

/**
 * QueryLeaseResponse is response type for the Query/Lease RPC method
 *
 * @generated from message akash.market.v1beta5.QueryLeaseResponse
 */
export type QueryLeaseResponseJson = {
  /**
   * Lease holds the lease for a deployment.
   *
   * @generated from field: akash.market.v1.Lease lease = 1;
   */
  lease?: LeaseJson;

  /**
   * EscrowPayment holds information about the Lease's fractional payment.
   *
   * @generated from field: akash.escrow.v1.FractionalPayment escrow_payment = 2;
   */
  escrowPayment?: FractionalPaymentJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryLeaseResponse.
 * Use `create(QueryLeaseResponseSchema)` to create a new message.
 */
export const QueryLeaseResponseSchema: GenMessage<QueryLeaseResponse, QueryLeaseResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 11);

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryParamsRequest
 */
export type QueryParamsRequest = Message<"akash.market.v1beta5.QueryParamsRequest"> & {
};

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryParamsRequest
 */
export type QueryParamsRequestJson = {
};

/**
 * Describes the message akash.market.v1beta5.QueryParamsRequest.
 * Use `create(QueryParamsRequestSchema)` to create a new message.
 */
export const QueryParamsRequestSchema: GenMessage<QueryParamsRequest, QueryParamsRequestJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 12);

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryParamsResponse
 */
export type QueryParamsResponse = Message<"akash.market.v1beta5.QueryParamsResponse"> & {
  /**
   * params defines the parameters of the module.
   *
   * @generated from field: akash.market.v1beta5.Params params = 1;
   */
  params?: Params;
};

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message akash.market.v1beta5.QueryParamsResponse
 */
export type QueryParamsResponseJson = {
  /**
   * params defines the parameters of the module.
   *
   * @generated from field: akash.market.v1beta5.Params params = 1;
   */
  params?: ParamsJson;
};

/**
 * Describes the message akash.market.v1beta5.QueryParamsResponse.
 * Use `create(QueryParamsResponseSchema)` to create a new message.
 */
export const QueryParamsResponseSchema: GenMessage<QueryParamsResponse, QueryParamsResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_market_v1beta5_query, 13);

/**
 * Query defines the gRPC querier service for the market package.
 *
 * @generated from service akash.market.v1beta5.Query
 */
export const Query: GenService<{
  /**
   * Orders queries orders with filters.
   *
   * @generated from rpc akash.market.v1beta5.Query.Orders
   */
  orders: {
    methodKind: "unary";
    input: typeof QueryOrdersRequestSchema;
    output: typeof QueryOrdersResponseSchema;
  },
  /**
   * Order queries order details.
   *
   * @generated from rpc akash.market.v1beta5.Query.Order
   */
  order: {
    methodKind: "unary";
    input: typeof QueryOrderRequestSchema;
    output: typeof QueryOrderResponseSchema;
  },
  /**
   * Bids queries bids with filters.
   *
   * @generated from rpc akash.market.v1beta5.Query.Bids
   */
  bids: {
    methodKind: "unary";
    input: typeof QueryBidsRequestSchema;
    output: typeof QueryBidsResponseSchema;
  },
  /**
   * Bid queries bid details.
   *
   * @generated from rpc akash.market.v1beta5.Query.Bid
   */
  bid: {
    methodKind: "unary";
    input: typeof QueryBidRequestSchema;
    output: typeof QueryBidResponseSchema;
  },
  /**
   * Leases queries leases with filters.
   *
   * @generated from rpc akash.market.v1beta5.Query.Leases
   */
  leases: {
    methodKind: "unary";
    input: typeof QueryLeasesRequestSchema;
    output: typeof QueryLeasesResponseSchema;
  },
  /**
   * Lease queries lease details.
   *
   * @generated from rpc akash.market.v1beta5.Query.Lease
   */
  lease: {
    methodKind: "unary";
    input: typeof QueryLeaseRequestSchema;
    output: typeof QueryLeaseResponseSchema;
  },
  /**
   * Params returns the total set of minting parameters.
   *
   * @generated from rpc akash.market.v1beta5.Query.Params
   */
  params: {
    methodKind: "unary";
    input: typeof QueryParamsRequestSchema;
    output: typeof QueryParamsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_akash_market_v1beta5_query, 0);

