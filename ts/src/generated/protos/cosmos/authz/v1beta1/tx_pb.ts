// Since: cosmos-sdk 0.43

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/authz/v1beta1/tx.proto (package cosmos.authz.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Any, AnyJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_any } from "@bufbuild/protobuf/wkt";
import type { Grant, GrantJson } from "./authz_pb.ts";
import { file_cosmos_authz_v1beta1_authz } from "./authz_pb.ts";
import { file_cosmos_msg_v1_msg } from "../../msg/v1/msg_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/authz/v1beta1/tx.proto.
 */
export const file_cosmos_authz_v1beta1_tx: GenFile = /*@__PURE__*/
  fileDesc("Ch1jb3Ntb3MvYXV0aHovdjFiZXRhMS90eC5wcm90bxIUY29zbW9zLmF1dGh6LnYxYmV0YTEivQEKCE1zZ0dyYW50EikKB2dyYW50ZXIYASABKAlCGNK0LRRjb3Ntb3MuQWRkcmVzc1N0cmluZxIpCgdncmFudGVlGAIgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSNQoFZ3JhbnQYAyABKAsyGy5jb3Ntb3MuYXV0aHoudjFiZXRhMS5HcmFudEIJyN4fAKjnsCoBOiSC57AqB2dyYW50ZXKK57AqE2Nvc21vcy1zZGsvTXNnR3JhbnQiIgoPTXNnRXhlY1Jlc3BvbnNlEg8KB3Jlc3VsdHMYASADKAwimgEKB01zZ0V4ZWMSKQoHZ3JhbnRlZRgBIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEj8KBG1zZ3MYAiADKAsyFC5nb29nbGUucHJvdG9idWYuQW55QhvKtC0XY29zbW9zLmJhc2UudjFiZXRhMS5Nc2c6I4LnsCoHZ3JhbnRlZYrnsCoSY29zbW9zLXNkay9Nc2dFeGVjIhIKEE1zZ0dyYW50UmVzcG9uc2UingEKCU1zZ1Jldm9rZRIpCgdncmFudGVyGAEgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSKQoHZ3JhbnRlZRgCIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEhQKDG1zZ190eXBlX3VybBgDIAEoCTolguewKgdncmFudGVyiuewKhRjb3Ntb3Mtc2RrL01zZ1Jldm9rZSITChFNc2dSZXZva2VSZXNwb25zZTL/AQoDTXNnEk8KBUdyYW50Eh4uY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnR3JhbnQaJi5jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dHcmFudFJlc3BvbnNlEkwKBEV4ZWMSHS5jb3Ntb3MuYXV0aHoudjFiZXRhMS5Nc2dFeGVjGiUuY29zbW9zLmF1dGh6LnYxYmV0YTEuTXNnRXhlY1Jlc3BvbnNlElIKBlJldm9rZRIfLmNvc21vcy5hdXRoei52MWJldGExLk1zZ1Jldm9rZRonLmNvc21vcy5hdXRoei52MWJldGExLk1zZ1Jldm9rZVJlc3BvbnNlGgWA57AqAUIqWiRnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvYXV0aHrI4R4AYgZwcm90bzM", [file_cosmos_proto_cosmos, file_gogoproto_gogo, file_google_protobuf_any, file_cosmos_authz_v1beta1_authz, file_cosmos_msg_v1_msg, file_amino_amino]);

/**
 * MsgGrant is a request type for Grant method. It declares authorization to the grantee
 * on behalf of the granter with the provided expiration time.
 *
 * @generated from message cosmos.authz.v1beta1.MsgGrant
 */
export type MsgGrant = Message<"cosmos.authz.v1beta1.MsgGrant"> & {
  /**
   * @generated from field: string granter = 1;
   */
  granter: string;

  /**
   * @generated from field: string grantee = 2;
   */
  grantee: string;

  /**
   * @generated from field: cosmos.authz.v1beta1.Grant grant = 3;
   */
  grant?: Grant;
};

/**
 * MsgGrant is a request type for Grant method. It declares authorization to the grantee
 * on behalf of the granter with the provided expiration time.
 *
 * @generated from message cosmos.authz.v1beta1.MsgGrant
 */
export type MsgGrantJson = {
  /**
   * @generated from field: string granter = 1;
   */
  granter?: string;

  /**
   * @generated from field: string grantee = 2;
   */
  grantee?: string;

  /**
   * @generated from field: cosmos.authz.v1beta1.Grant grant = 3;
   */
  grant?: GrantJson;
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgGrant.
 * Use `create(MsgGrantSchema)` to create a new message.
 */
export const MsgGrantSchema: GenMessage<MsgGrant, MsgGrantJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 0);

/**
 * MsgExecResponse defines the Msg/MsgExecResponse response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgExecResponse
 */
export type MsgExecResponse = Message<"cosmos.authz.v1beta1.MsgExecResponse"> & {
  /**
   * @generated from field: repeated bytes results = 1;
   */
  results: Uint8Array[];
};

/**
 * MsgExecResponse defines the Msg/MsgExecResponse response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgExecResponse
 */
export type MsgExecResponseJson = {
  /**
   * @generated from field: repeated bytes results = 1;
   */
  results?: string[];
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgExecResponse.
 * Use `create(MsgExecResponseSchema)` to create a new message.
 */
export const MsgExecResponseSchema: GenMessage<MsgExecResponse, MsgExecResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 1);

/**
 * MsgExec attempts to execute the provided messages using
 * authorizations granted to the grantee. Each message should have only
 * one signer corresponding to the granter of the authorization.
 *
 * @generated from message cosmos.authz.v1beta1.MsgExec
 */
export type MsgExec = Message<"cosmos.authz.v1beta1.MsgExec"> & {
  /**
   * @generated from field: string grantee = 1;
   */
  grantee: string;

  /**
   * Execute Msg.
   * The x/authz will try to find a grant matching (msg.signers[0], grantee, MsgTypeURL(msg))
   * triple and validate it.
   *
   * @generated from field: repeated google.protobuf.Any msgs = 2;
   */
  msgs: Any[];
};

/**
 * MsgExec attempts to execute the provided messages using
 * authorizations granted to the grantee. Each message should have only
 * one signer corresponding to the granter of the authorization.
 *
 * @generated from message cosmos.authz.v1beta1.MsgExec
 */
export type MsgExecJson = {
  /**
   * @generated from field: string grantee = 1;
   */
  grantee?: string;

  /**
   * Execute Msg.
   * The x/authz will try to find a grant matching (msg.signers[0], grantee, MsgTypeURL(msg))
   * triple and validate it.
   *
   * @generated from field: repeated google.protobuf.Any msgs = 2;
   */
  msgs?: AnyJson[];
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgExec.
 * Use `create(MsgExecSchema)` to create a new message.
 */
export const MsgExecSchema: GenMessage<MsgExec, MsgExecJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 2);

/**
 * MsgGrantResponse defines the Msg/MsgGrant response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgGrantResponse
 */
export type MsgGrantResponse = Message<"cosmos.authz.v1beta1.MsgGrantResponse"> & {
};

/**
 * MsgGrantResponse defines the Msg/MsgGrant response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgGrantResponse
 */
export type MsgGrantResponseJson = {
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgGrantResponse.
 * Use `create(MsgGrantResponseSchema)` to create a new message.
 */
export const MsgGrantResponseSchema: GenMessage<MsgGrantResponse, MsgGrantResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 3);

/**
 * MsgRevoke revokes any authorization with the provided sdk.Msg type on the
 * granter's account with that has been granted to the grantee.
 *
 * @generated from message cosmos.authz.v1beta1.MsgRevoke
 */
export type MsgRevoke = Message<"cosmos.authz.v1beta1.MsgRevoke"> & {
  /**
   * @generated from field: string granter = 1;
   */
  granter: string;

  /**
   * @generated from field: string grantee = 2;
   */
  grantee: string;

  /**
   * @generated from field: string msg_type_url = 3;
   */
  msgTypeUrl: string;
};

/**
 * MsgRevoke revokes any authorization with the provided sdk.Msg type on the
 * granter's account with that has been granted to the grantee.
 *
 * @generated from message cosmos.authz.v1beta1.MsgRevoke
 */
export type MsgRevokeJson = {
  /**
   * @generated from field: string granter = 1;
   */
  granter?: string;

  /**
   * @generated from field: string grantee = 2;
   */
  grantee?: string;

  /**
   * @generated from field: string msg_type_url = 3;
   */
  msgTypeUrl?: string;
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgRevoke.
 * Use `create(MsgRevokeSchema)` to create a new message.
 */
export const MsgRevokeSchema: GenMessage<MsgRevoke, MsgRevokeJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 4);

/**
 * MsgRevokeResponse defines the Msg/MsgRevokeResponse response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgRevokeResponse
 */
export type MsgRevokeResponse = Message<"cosmos.authz.v1beta1.MsgRevokeResponse"> & {
};

/**
 * MsgRevokeResponse defines the Msg/MsgRevokeResponse response type.
 *
 * @generated from message cosmos.authz.v1beta1.MsgRevokeResponse
 */
export type MsgRevokeResponseJson = {
};

/**
 * Describes the message cosmos.authz.v1beta1.MsgRevokeResponse.
 * Use `create(MsgRevokeResponseSchema)` to create a new message.
 */
export const MsgRevokeResponseSchema: GenMessage<MsgRevokeResponse, MsgRevokeResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_authz_v1beta1_tx, 5);

/**
 * Msg defines the authz Msg service.
 *
 * @generated from service cosmos.authz.v1beta1.Msg
 */
export const Msg: GenService<{
  /**
   * Grant grants the provided authorization to the grantee on the granter's
   * account with the provided expiration time. If there is already a grant
   * for the given (granter, grantee, Authorization) triple, then the grant
   * will be overwritten.
   *
   * @generated from rpc cosmos.authz.v1beta1.Msg.Grant
   */
  grant: {
    methodKind: "unary";
    input: typeof MsgGrantSchema;
    output: typeof MsgGrantResponseSchema;
  },
  /**
   * Exec attempts to execute the provided messages using
   * authorizations granted to the grantee. Each message should have only
   * one signer corresponding to the granter of the authorization.
   *
   * @generated from rpc cosmos.authz.v1beta1.Msg.Exec
   */
  exec: {
    methodKind: "unary";
    input: typeof MsgExecSchema;
    output: typeof MsgExecResponseSchema;
  },
  /**
   * Revoke revokes any authorization corresponding to the provided method name on the
   * granter's account that has been granted to the grantee.
   *
   * @generated from rpc cosmos.authz.v1beta1.Msg.Revoke
   */
  revoke: {
    methodKind: "unary";
    input: typeof MsgRevokeSchema;
    output: typeof MsgRevokeResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_cosmos_authz_v1beta1_tx, 0);

