// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/cert/v1/msg.proto (package akash.cert.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import { file_cosmos_msg_v1_msg } from "../../../cosmos/msg/v1/msg_pb.ts";
import type { ID, IDJson } from "./cert_pb.ts";
import { file_akash_cert_v1_cert } from "./cert_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/cert/v1/msg.proto.
 */
export const file_akash_cert_v1_msg: GenFile = /*@__PURE__*/
  fileDesc("Chdha2FzaC9jZXJ0L3YxL21zZy5wcm90bxINYWthc2guY2VydC52MSK8AQoUTXNnQ3JlYXRlQ2VydGlmaWNhdGUSQAoFb3duZXIYASABKAlCMereHwVvd25lcvLeHwx5YW1sOiJvd25lciLStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSJQoEY2VydBgCIAEoDEIX6t4fBGNlcnTy3h8LeWFtbDoiY2VydCISKwoGcHVia2V5GAMgASgMQhvq3h8GcHVia2V58t4fDXlhbWw6InB1YmtleSI6DuigHwCC57AqBW93bmVyIh4KHE1zZ0NyZWF0ZUNlcnRpZmljYXRlUmVzcG9uc2UiZwoUTXNnUmV2b2tlQ2VydGlmaWNhdGUSPAoCaWQYASABKAsyES5ha2FzaC5jZXJ0LnYxLklEQh3I3h8A4t4fAklE6t4fAmlk8t4fCXlhbWw6ImlkIjoR6KAfAILnsCoIaWQub3duZXIiHgocTXNnUmV2b2tlQ2VydGlmaWNhdGVSZXNwb25zZUIdWhtwa2cuYWt0LmRldi9nby9ub2RlL2NlcnQvdjFiBnByb3RvMw", [file_gogoproto_gogo, file_cosmos_proto_cosmos, file_cosmos_msg_v1_msg, file_akash_cert_v1_cert]);

/**
 * MsgCreateCertificate defines an SDK message for creating certificate.
 *
 * @generated from message akash.cert.v1.MsgCreateCertificate
 */
export type MsgCreateCertificate = Message<"akash.cert.v1.MsgCreateCertificate"> & {
  /**
   * Owner is the account address of the user who owns the certificate.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 1;
   */
  owner: string;

  /**
   * Cert holds the bytes representing the certificate.
   *
   * @generated from field: bytes cert = 2;
   */
  cert: Uint8Array;

  /**
   * PubKey holds the public key.
   *
   * @generated from field: bytes pubkey = 3;
   */
  pubkey: Uint8Array;
};

/**
 * MsgCreateCertificate defines an SDK message for creating certificate.
 *
 * @generated from message akash.cert.v1.MsgCreateCertificate
 */
export type MsgCreateCertificateJson = {
  /**
   * Owner is the account address of the user who owns the certificate.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 1;
   */
  owner?: string;

  /**
   * Cert holds the bytes representing the certificate.
   *
   * @generated from field: bytes cert = 2;
   */
  cert?: string;

  /**
   * PubKey holds the public key.
   *
   * @generated from field: bytes pubkey = 3;
   */
  pubkey?: string;
};

/**
 * Describes the message akash.cert.v1.MsgCreateCertificate.
 * Use `create(MsgCreateCertificateSchema)` to create a new message.
 */
export const MsgCreateCertificateSchema: GenMessage<MsgCreateCertificate, MsgCreateCertificateJson> = /*@__PURE__*/
  messageDesc(file_akash_cert_v1_msg, 0);

/**
 * MsgCreateCertificateResponse defines the Msg/CreateCertificate response type.
 *
 * @generated from message akash.cert.v1.MsgCreateCertificateResponse
 */
export type MsgCreateCertificateResponse = Message<"akash.cert.v1.MsgCreateCertificateResponse"> & {
};

/**
 * MsgCreateCertificateResponse defines the Msg/CreateCertificate response type.
 *
 * @generated from message akash.cert.v1.MsgCreateCertificateResponse
 */
export type MsgCreateCertificateResponseJson = {
};

/**
 * Describes the message akash.cert.v1.MsgCreateCertificateResponse.
 * Use `create(MsgCreateCertificateResponseSchema)` to create a new message.
 */
export const MsgCreateCertificateResponseSchema: GenMessage<MsgCreateCertificateResponse, MsgCreateCertificateResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_cert_v1_msg, 1);

/**
 * MsgRevokeCertificate defines an SDK message for revoking certificate.
 *
 * @generated from message akash.cert.v1.MsgRevokeCertificate
 */
export type MsgRevokeCertificate = Message<"akash.cert.v1.MsgRevokeCertificate"> & {
  /**
   * Id corresponds to the certificate ID which includes owner and sequence number.
   *
   * @generated from field: akash.cert.v1.ID id = 1;
   */
  id?: ID;
};

/**
 * MsgRevokeCertificate defines an SDK message for revoking certificate.
 *
 * @generated from message akash.cert.v1.MsgRevokeCertificate
 */
export type MsgRevokeCertificateJson = {
  /**
   * Id corresponds to the certificate ID which includes owner and sequence number.
   *
   * @generated from field: akash.cert.v1.ID id = 1;
   */
  id?: IDJson;
};

/**
 * Describes the message akash.cert.v1.MsgRevokeCertificate.
 * Use `create(MsgRevokeCertificateSchema)` to create a new message.
 */
export const MsgRevokeCertificateSchema: GenMessage<MsgRevokeCertificate, MsgRevokeCertificateJson> = /*@__PURE__*/
  messageDesc(file_akash_cert_v1_msg, 2);

/**
 * MsgRevokeCertificateResponse defines the Msg/RevokeCertificate response type.
 *
 * @generated from message akash.cert.v1.MsgRevokeCertificateResponse
 */
export type MsgRevokeCertificateResponse = Message<"akash.cert.v1.MsgRevokeCertificateResponse"> & {
};

/**
 * MsgRevokeCertificateResponse defines the Msg/RevokeCertificate response type.
 *
 * @generated from message akash.cert.v1.MsgRevokeCertificateResponse
 */
export type MsgRevokeCertificateResponseJson = {
};

/**
 * Describes the message akash.cert.v1.MsgRevokeCertificateResponse.
 * Use `create(MsgRevokeCertificateResponseSchema)` to create a new message.
 */
export const MsgRevokeCertificateResponseSchema: GenMessage<MsgRevokeCertificateResponse, MsgRevokeCertificateResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_cert_v1_msg, 3);

