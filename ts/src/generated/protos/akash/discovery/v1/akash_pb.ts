// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/discovery/v1/akash.proto (package akash.discovery.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { ClientInfo, ClientInfoJson } from "./client_info_pb.ts";
import { file_akash_discovery_v1_client_info } from "./client_info_pb.ts";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/discovery/v1/akash.proto.
 */
export const file_akash_discovery_v1_akash: GenFile = /*@__PURE__*/
  fileDesc("Ch5ha2FzaC9kaXNjb3ZlcnkvdjEvYWthc2gucHJvdG8SEmFrYXNoLmRpc2NvdmVyeS52MSJxCgVBa2FzaBJoCgtjbGllbnRfaW5mbxgBIAEoCzIeLmFrYXNoLmRpc2NvdmVyeS52MS5DbGllbnRJbmZvQjPi3h8KQ2xpZW50SW5mb+reHwtjbGllbnRfaW5mb/LeHxJ5YW1sOiJjbGllbnRfaW5mbyJCHFoacGtnLmFrdC5kZXYvZ28vbm9kZS9jbGllbnRiBnByb3RvMw", [file_akash_discovery_v1_client_info, file_gogoproto_gogo]);

/**
 * Akash akash specific RPC parameters.
 *
 * @generated from message akash.discovery.v1.Akash
 */
export type Akash = Message<"akash.discovery.v1.Akash"> & {
  /**
   * ClientInfo holds information about the client.
   *
   * @generated from field: akash.discovery.v1.ClientInfo client_info = 1;
   */
  clientInfo?: ClientInfo;
};

/**
 * Akash akash specific RPC parameters.
 *
 * @generated from message akash.discovery.v1.Akash
 */
export type AkashJson = {
  /**
   * ClientInfo holds information about the client.
   *
   * @generated from field: akash.discovery.v1.ClientInfo client_info = 1;
   */
  clientInfo?: ClientInfoJson;
};

/**
 * Describes the message akash.discovery.v1.Akash.
 * Use `create(AkashSchema)` to create a new message.
 */
export const AkashSchema: GenMessage<Akash, AkashJson> = /*@__PURE__*/
  messageDesc(file_akash_discovery_v1_akash, 0);

