// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/cert/v1/filters.proto (package akash.cert.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/cert/v1/filters.proto.
 */
export const file_akash_cert_v1_filters: GenFile = /*@__PURE__*/
  fileDesc("Chtha2FzaC9jZXJ0L3YxL2ZpbHRlcnMucHJvdG8SDWFrYXNoLmNlcnQudjEisgEKEUNlcnRpZmljYXRlRmlsdGVyEkAKBW93bmVyGAEgASgJQjHq3h8Fb3duZXLy3h8MeWFtbDoib3duZXIi0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEisKBnNlcmlhbBgCIAEoCUIb6t4fBnNlcmlhbPLeHw15YW1sOiJzZXJpYWwiEigKBXN0YXRlGAMgASgJQhnq3h8Fc3RhdGXy3h8MeWFtbDoic3RhdGUiOgTooB8AQh1aG3BrZy5ha3QuZGV2L2dvL25vZGUvY2VydC92MWIGcHJvdG8z", [file_gogoproto_gogo, file_cosmos_proto_cosmos]);

/**
 * CertificateFilter defines filters used to filter certificates.
 *
 * @generated from message akash.cert.v1.CertificateFilter
 */
export type CertificateFilter = Message<"akash.cert.v1.CertificateFilter"> & {
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
   * Serial is a sequence number for the certificate.
   *
   * @generated from field: string serial = 2;
   */
  serial: string;

  /**
   * State is the state of the certificate.
   * CertificateValid denotes state for deployment active.
   * CertificateRevoked denotes state for deployment closed.
   *
   * @generated from field: string state = 3;
   */
  state: string;
};

/**
 * CertificateFilter defines filters used to filter certificates.
 *
 * @generated from message akash.cert.v1.CertificateFilter
 */
export type CertificateFilterJson = {
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
   * Serial is a sequence number for the certificate.
   *
   * @generated from field: string serial = 2;
   */
  serial?: string;

  /**
   * State is the state of the certificate.
   * CertificateValid denotes state for deployment active.
   * CertificateRevoked denotes state for deployment closed.
   *
   * @generated from field: string state = 3;
   */
  state?: string;
};

/**
 * Describes the message akash.cert.v1.CertificateFilter.
 * Use `create(CertificateFilterSchema)` to create a new message.
 */
export const CertificateFilterSchema: GenMessage<CertificateFilter, CertificateFilterJson> = /*@__PURE__*/
  messageDesc(file_akash_cert_v1_filters, 0);

