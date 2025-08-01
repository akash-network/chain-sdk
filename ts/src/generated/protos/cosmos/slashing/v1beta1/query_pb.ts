// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/slashing/v1beta1/query.proto (package cosmos.slashing.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { PageRequest, PageRequestJson, PageResponse, PageResponseJson } from "../../base/query/v1beta1/pagination_pb.ts";
import { file_cosmos_base_query_v1beta1_pagination } from "../../base/query/v1beta1/pagination_pb.ts";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_google_api_annotations } from "../../../google/api/annotations_pb.ts";
import type { Params, ParamsJson, ValidatorSigningInfo, ValidatorSigningInfoJson } from "./slashing_pb.ts";
import { file_cosmos_slashing_v1beta1_slashing } from "./slashing_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/slashing/v1beta1/query.proto.
 */
export const file_cosmos_slashing_v1beta1_query: GenFile = /*@__PURE__*/
  fileDesc("CiNjb3Ntb3Mvc2xhc2hpbmcvdjFiZXRhMS9xdWVyeS5wcm90bxIXY29zbW9zLnNsYXNoaW5nLnYxYmV0YTEiFAoSUXVlcnlQYXJhbXNSZXF1ZXN0IlEKE1F1ZXJ5UGFyYW1zUmVzcG9uc2USOgoGcGFyYW1zGAEgASgLMh8uY29zbW9zLnNsYXNoaW5nLnYxYmV0YTEuUGFyYW1zQgnI3h8AqOewKgEiSQoXUXVlcnlTaWduaW5nSW5mb1JlcXVlc3QSLgoMY29uc19hZGRyZXNzGAEgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcibgoYUXVlcnlTaWduaW5nSW5mb1Jlc3BvbnNlElIKEHZhbF9zaWduaW5nX2luZm8YASABKAsyLS5jb3Ntb3Muc2xhc2hpbmcudjFiZXRhMS5WYWxpZGF0b3JTaWduaW5nSW5mb0IJyN4fAKjnsCoBIlYKGFF1ZXJ5U2lnbmluZ0luZm9zUmVxdWVzdBI6CgpwYWdpbmF0aW9uGAEgASgLMiYuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVxdWVzdCKgAQoZUXVlcnlTaWduaW5nSW5mb3NSZXNwb25zZRJGCgRpbmZvGAEgAygLMi0uY29zbW9zLnNsYXNoaW5nLnYxYmV0YTEuVmFsaWRhdG9yU2lnbmluZ0luZm9CCcjeHwCo57AqARI7CgpwYWdpbmF0aW9uGAIgASgLMicuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVzcG9uc2Uy8gMKBVF1ZXJ5EowBCgZQYXJhbXMSKy5jb3Ntb3Muc2xhc2hpbmcudjFiZXRhMS5RdWVyeVBhcmFtc1JlcXVlc3QaLC5jb3Ntb3Muc2xhc2hpbmcudjFiZXRhMS5RdWVyeVBhcmFtc1Jlc3BvbnNlIieC0+STAiESHy9jb3Ntb3Mvc2xhc2hpbmcvdjFiZXRhMS9wYXJhbXMSsQEKC1NpZ25pbmdJbmZvEjAuY29zbW9zLnNsYXNoaW5nLnYxYmV0YTEuUXVlcnlTaWduaW5nSW5mb1JlcXVlc3QaMS5jb3Ntb3Muc2xhc2hpbmcudjFiZXRhMS5RdWVyeVNpZ25pbmdJbmZvUmVzcG9uc2UiPYLT5JMCNxI1L2Nvc21vcy9zbGFzaGluZy92MWJldGExL3NpZ25pbmdfaW5mb3Mve2NvbnNfYWRkcmVzc30SpQEKDFNpZ25pbmdJbmZvcxIxLmNvc21vcy5zbGFzaGluZy52MWJldGExLlF1ZXJ5U2lnbmluZ0luZm9zUmVxdWVzdBoyLmNvc21vcy5zbGFzaGluZy52MWJldGExLlF1ZXJ5U2lnbmluZ0luZm9zUmVzcG9uc2UiLoLT5JMCKBImL2Nvc21vcy9zbGFzaGluZy92MWJldGExL3NpZ25pbmdfaW5mb3NCL1otZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay94L3NsYXNoaW5nL3R5cGVzYgZwcm90bzM", [file_cosmos_base_query_v1beta1_pagination, file_gogoproto_gogo, file_google_api_annotations, file_cosmos_slashing_v1beta1_slashing, file_cosmos_proto_cosmos, file_amino_amino]);

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method
 *
 * @generated from message cosmos.slashing.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequest = Message<"cosmos.slashing.v1beta1.QueryParamsRequest"> & {
};

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method
 *
 * @generated from message cosmos.slashing.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequestJson = {
};

/**
 * Describes the message cosmos.slashing.v1beta1.QueryParamsRequest.
 * Use `create(QueryParamsRequestSchema)` to create a new message.
 */
export const QueryParamsRequestSchema: GenMessage<QueryParamsRequest, QueryParamsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 0);

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method
 *
 * @generated from message cosmos.slashing.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponse = Message<"cosmos.slashing.v1beta1.QueryParamsResponse"> & {
  /**
   * @generated from field: cosmos.slashing.v1beta1.Params params = 1;
   */
  params?: Params;
};

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method
 *
 * @generated from message cosmos.slashing.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponseJson = {
  /**
   * @generated from field: cosmos.slashing.v1beta1.Params params = 1;
   */
  params?: ParamsJson;
};

/**
 * Describes the message cosmos.slashing.v1beta1.QueryParamsResponse.
 * Use `create(QueryParamsResponseSchema)` to create a new message.
 */
export const QueryParamsResponseSchema: GenMessage<QueryParamsResponse, QueryParamsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 1);

/**
 * QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfoRequest
 */
export type QuerySigningInfoRequest = Message<"cosmos.slashing.v1beta1.QuerySigningInfoRequest"> & {
  /**
   * cons_address is the address to query signing info of
   *
   * @generated from field: string cons_address = 1;
   */
  consAddress: string;
};

/**
 * QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfoRequest
 */
export type QuerySigningInfoRequestJson = {
  /**
   * cons_address is the address to query signing info of
   *
   * @generated from field: string cons_address = 1;
   */
  consAddress?: string;
};

/**
 * Describes the message cosmos.slashing.v1beta1.QuerySigningInfoRequest.
 * Use `create(QuerySigningInfoRequestSchema)` to create a new message.
 */
export const QuerySigningInfoRequestSchema: GenMessage<QuerySigningInfoRequest, QuerySigningInfoRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 2);

/**
 * QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfoResponse
 */
export type QuerySigningInfoResponse = Message<"cosmos.slashing.v1beta1.QuerySigningInfoResponse"> & {
  /**
   * val_signing_info is the signing info of requested val cons address
   *
   * @generated from field: cosmos.slashing.v1beta1.ValidatorSigningInfo val_signing_info = 1;
   */
  valSigningInfo?: ValidatorSigningInfo;
};

/**
 * QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfoResponse
 */
export type QuerySigningInfoResponseJson = {
  /**
   * val_signing_info is the signing info of requested val cons address
   *
   * @generated from field: cosmos.slashing.v1beta1.ValidatorSigningInfo val_signing_info = 1;
   */
  valSigningInfo?: ValidatorSigningInfoJson;
};

/**
 * Describes the message cosmos.slashing.v1beta1.QuerySigningInfoResponse.
 * Use `create(QuerySigningInfoResponseSchema)` to create a new message.
 */
export const QuerySigningInfoResponseSchema: GenMessage<QuerySigningInfoResponse, QuerySigningInfoResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 3);

/**
 * QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfosRequest
 */
export type QuerySigningInfosRequest = Message<"cosmos.slashing.v1beta1.QuerySigningInfosRequest"> & {
  /**
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequest;
};

/**
 * QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfosRequest
 */
export type QuerySigningInfosRequestJson = {
  /**
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 1;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message cosmos.slashing.v1beta1.QuerySigningInfosRequest.
 * Use `create(QuerySigningInfosRequestSchema)` to create a new message.
 */
export const QuerySigningInfosRequestSchema: GenMessage<QuerySigningInfosRequest, QuerySigningInfosRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 4);

/**
 * QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfosResponse
 */
export type QuerySigningInfosResponse = Message<"cosmos.slashing.v1beta1.QuerySigningInfosResponse"> & {
  /**
   * info is the signing info of all validators
   *
   * @generated from field: repeated cosmos.slashing.v1beta1.ValidatorSigningInfo info = 1;
   */
  info: ValidatorSigningInfo[];

  /**
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
 * method
 *
 * @generated from message cosmos.slashing.v1beta1.QuerySigningInfosResponse
 */
export type QuerySigningInfosResponseJson = {
  /**
   * info is the signing info of all validators
   *
   * @generated from field: repeated cosmos.slashing.v1beta1.ValidatorSigningInfo info = 1;
   */
  info?: ValidatorSigningInfoJson[];

  /**
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message cosmos.slashing.v1beta1.QuerySigningInfosResponse.
 * Use `create(QuerySigningInfosResponseSchema)` to create a new message.
 */
export const QuerySigningInfosResponseSchema: GenMessage<QuerySigningInfosResponse, QuerySigningInfosResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_query, 5);

/**
 * Query provides defines the gRPC querier service
 *
 * @generated from service cosmos.slashing.v1beta1.Query
 */
export const Query: GenService<{
  /**
   * Params queries the parameters of slashing module
   *
   * @generated from rpc cosmos.slashing.v1beta1.Query.Params
   */
  params: {
    methodKind: "unary";
    input: typeof QueryParamsRequestSchema;
    output: typeof QueryParamsResponseSchema;
  },
  /**
   * SigningInfo queries the signing info of given cons address
   *
   * @generated from rpc cosmos.slashing.v1beta1.Query.SigningInfo
   */
  signingInfo: {
    methodKind: "unary";
    input: typeof QuerySigningInfoRequestSchema;
    output: typeof QuerySigningInfoResponseSchema;
  },
  /**
   * SigningInfos queries signing info of all validators
   *
   * @generated from rpc cosmos.slashing.v1beta1.Query.SigningInfos
   */
  signingInfos: {
    methodKind: "unary";
    input: typeof QuerySigningInfosRequestSchema;
    output: typeof QuerySigningInfosResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_cosmos_slashing_v1beta1_query, 0);

