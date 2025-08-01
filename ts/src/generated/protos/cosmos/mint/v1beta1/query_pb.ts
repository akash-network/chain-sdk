// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/mint/v1beta1/query.proto (package cosmos.mint.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_google_api_annotations } from "../../../google/api/annotations_pb.ts";
import type { Params, ParamsJson } from "./mint_pb.ts";
import { file_cosmos_mint_v1beta1_mint } from "./mint_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/mint/v1beta1/query.proto.
 */
export const file_cosmos_mint_v1beta1_query: GenFile = /*@__PURE__*/
  fileDesc("Ch9jb3Ntb3MvbWludC92MWJldGExL3F1ZXJ5LnByb3RvEhNjb3Ntb3MubWludC52MWJldGExIhQKElF1ZXJ5UGFyYW1zUmVxdWVzdCJNChNRdWVyeVBhcmFtc1Jlc3BvbnNlEjYKBnBhcmFtcxgBIAEoCzIbLmNvc21vcy5taW50LnYxYmV0YTEuUGFyYW1zQgnI3h8AqOewKgEiFwoVUXVlcnlJbmZsYXRpb25SZXF1ZXN0ImAKFlF1ZXJ5SW5mbGF0aW9uUmVzcG9uc2USRgoJaW5mbGF0aW9uGAEgASgMQjPI3h8A2t4fJmdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjqOewKgEiHgocUXVlcnlBbm51YWxQcm92aXNpb25zUmVxdWVzdCJvCh1RdWVyeUFubnVhbFByb3Zpc2lvbnNSZXNwb25zZRJOChFhbm51YWxfcHJvdmlzaW9ucxgBIAEoDEIzyN4fANreHyZnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkRlY6jnsCoBMsUDCgVRdWVyeRKAAQoGUGFyYW1zEicuY29zbW9zLm1pbnQudjFiZXRhMS5RdWVyeVBhcmFtc1JlcXVlc3QaKC5jb3Ntb3MubWludC52MWJldGExLlF1ZXJ5UGFyYW1zUmVzcG9uc2UiI4LT5JMCHRIbL2Nvc21vcy9taW50L3YxYmV0YTEvcGFyYW1zEowBCglJbmZsYXRpb24SKi5jb3Ntb3MubWludC52MWJldGExLlF1ZXJ5SW5mbGF0aW9uUmVxdWVzdBorLmNvc21vcy5taW50LnYxYmV0YTEuUXVlcnlJbmZsYXRpb25SZXNwb25zZSImgtPkkwIgEh4vY29zbW9zL21pbnQvdjFiZXRhMS9pbmZsYXRpb24SqQEKEEFubnVhbFByb3Zpc2lvbnMSMS5jb3Ntb3MubWludC52MWJldGExLlF1ZXJ5QW5udWFsUHJvdmlzaW9uc1JlcXVlc3QaMi5jb3Ntb3MubWludC52MWJldGExLlF1ZXJ5QW5udWFsUHJvdmlzaW9uc1Jlc3BvbnNlIi6C0+STAigSJi9jb3Ntb3MvbWludC92MWJldGExL2FubnVhbF9wcm92aXNpb25zQitaKWdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsveC9taW50L3R5cGVzYgZwcm90bzM", [file_gogoproto_gogo, file_google_api_annotations, file_cosmos_mint_v1beta1_mint, file_amino_amino]);

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequest = Message<"cosmos.mint.v1beta1.QueryParamsRequest"> & {
};

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequestJson = {
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryParamsRequest.
 * Use `create(QueryParamsRequestSchema)` to create a new message.
 */
export const QueryParamsRequestSchema: GenMessage<QueryParamsRequest, QueryParamsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 0);

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponse = Message<"cosmos.mint.v1beta1.QueryParamsResponse"> & {
  /**
   * params defines the parameters of the module.
   *
   * @generated from field: cosmos.mint.v1beta1.Params params = 1;
   */
  params?: Params;
};

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponseJson = {
  /**
   * params defines the parameters of the module.
   *
   * @generated from field: cosmos.mint.v1beta1.Params params = 1;
   */
  params?: ParamsJson;
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryParamsResponse.
 * Use `create(QueryParamsResponseSchema)` to create a new message.
 */
export const QueryParamsResponseSchema: GenMessage<QueryParamsResponse, QueryParamsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 1);

/**
 * QueryInflationRequest is the request type for the Query/Inflation RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryInflationRequest
 */
export type QueryInflationRequest = Message<"cosmos.mint.v1beta1.QueryInflationRequest"> & {
};

/**
 * QueryInflationRequest is the request type for the Query/Inflation RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryInflationRequest
 */
export type QueryInflationRequestJson = {
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryInflationRequest.
 * Use `create(QueryInflationRequestSchema)` to create a new message.
 */
export const QueryInflationRequestSchema: GenMessage<QueryInflationRequest, QueryInflationRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 2);

/**
 * QueryInflationResponse is the response type for the Query/Inflation RPC
 * method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryInflationResponse
 */
export type QueryInflationResponse = Message<"cosmos.mint.v1beta1.QueryInflationResponse"> & {
  /**
   * inflation is the current minting inflation value.
   *
   * @generated from field: bytes inflation = 1;
   */
  inflation: Uint8Array;
};

/**
 * QueryInflationResponse is the response type for the Query/Inflation RPC
 * method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryInflationResponse
 */
export type QueryInflationResponseJson = {
  /**
   * inflation is the current minting inflation value.
   *
   * @generated from field: bytes inflation = 1;
   */
  inflation?: string;
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryInflationResponse.
 * Use `create(QueryInflationResponseSchema)` to create a new message.
 */
export const QueryInflationResponseSchema: GenMessage<QueryInflationResponse, QueryInflationResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 3);

/**
 * QueryAnnualProvisionsRequest is the request type for the
 * Query/AnnualProvisions RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryAnnualProvisionsRequest
 */
export type QueryAnnualProvisionsRequest = Message<"cosmos.mint.v1beta1.QueryAnnualProvisionsRequest"> & {
};

/**
 * QueryAnnualProvisionsRequest is the request type for the
 * Query/AnnualProvisions RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryAnnualProvisionsRequest
 */
export type QueryAnnualProvisionsRequestJson = {
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryAnnualProvisionsRequest.
 * Use `create(QueryAnnualProvisionsRequestSchema)` to create a new message.
 */
export const QueryAnnualProvisionsRequestSchema: GenMessage<QueryAnnualProvisionsRequest, QueryAnnualProvisionsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 4);

/**
 * QueryAnnualProvisionsResponse is the response type for the
 * Query/AnnualProvisions RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryAnnualProvisionsResponse
 */
export type QueryAnnualProvisionsResponse = Message<"cosmos.mint.v1beta1.QueryAnnualProvisionsResponse"> & {
  /**
   * annual_provisions is the current minting annual provisions value.
   *
   * @generated from field: bytes annual_provisions = 1;
   */
  annualProvisions: Uint8Array;
};

/**
 * QueryAnnualProvisionsResponse is the response type for the
 * Query/AnnualProvisions RPC method.
 *
 * @generated from message cosmos.mint.v1beta1.QueryAnnualProvisionsResponse
 */
export type QueryAnnualProvisionsResponseJson = {
  /**
   * annual_provisions is the current minting annual provisions value.
   *
   * @generated from field: bytes annual_provisions = 1;
   */
  annualProvisions?: string;
};

/**
 * Describes the message cosmos.mint.v1beta1.QueryAnnualProvisionsResponse.
 * Use `create(QueryAnnualProvisionsResponseSchema)` to create a new message.
 */
export const QueryAnnualProvisionsResponseSchema: GenMessage<QueryAnnualProvisionsResponse, QueryAnnualProvisionsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_mint_v1beta1_query, 5);

/**
 * Query provides defines the gRPC querier service.
 *
 * @generated from service cosmos.mint.v1beta1.Query
 */
export const Query: GenService<{
  /**
   * Params returns the total set of minting parameters.
   *
   * @generated from rpc cosmos.mint.v1beta1.Query.Params
   */
  params: {
    methodKind: "unary";
    input: typeof QueryParamsRequestSchema;
    output: typeof QueryParamsResponseSchema;
  },
  /**
   * Inflation returns the current minting inflation value.
   *
   * @generated from rpc cosmos.mint.v1beta1.Query.Inflation
   */
  inflation: {
    methodKind: "unary";
    input: typeof QueryInflationRequestSchema;
    output: typeof QueryInflationResponseSchema;
  },
  /**
   * AnnualProvisions current minting annual provisions value.
   *
   * @generated from rpc cosmos.mint.v1beta1.Query.AnnualProvisions
   */
  annualProvisions: {
    methodKind: "unary";
    input: typeof QueryAnnualProvisionsRequestSchema;
    output: typeof QueryAnnualProvisionsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_cosmos_mint_v1beta1_query, 0);

