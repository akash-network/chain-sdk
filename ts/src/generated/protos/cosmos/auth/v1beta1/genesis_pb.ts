// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/auth/v1beta1/genesis.proto (package cosmos.auth.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Any, AnyJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_any } from "@bufbuild/protobuf/wkt";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Params, ParamsJson } from "./auth_pb.ts";
import { file_cosmos_auth_v1beta1_auth } from "./auth_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/auth/v1beta1/genesis.proto.
 */
export const file_cosmos_auth_v1beta1_genesis: GenFile = /*@__PURE__*/
  fileDesc("CiFjb3Ntb3MvYXV0aC92MWJldGExL2dlbmVzaXMucHJvdG8SE2Nvc21vcy5hdXRoLnYxYmV0YTEibgoMR2VuZXNpc1N0YXRlEjYKBnBhcmFtcxgBIAEoCzIbLmNvc21vcy5hdXRoLnYxYmV0YTEuUGFyYW1zQgnI3h8AqOewKgESJgoIYWNjb3VudHMYAiADKAsyFC5nb29nbGUucHJvdG9idWYuQW55QitaKWdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsveC9hdXRoL3R5cGVzYgZwcm90bzM", [file_google_protobuf_any, file_gogoproto_gogo, file_cosmos_auth_v1beta1_auth, file_amino_amino]);

/**
 * GenesisState defines the auth module's genesis state.
 *
 * @generated from message cosmos.auth.v1beta1.GenesisState
 */
export type GenesisState = Message<"cosmos.auth.v1beta1.GenesisState"> & {
  /**
   * params defines all the parameters of the module.
   *
   * @generated from field: cosmos.auth.v1beta1.Params params = 1;
   */
  params?: Params;

  /**
   * accounts are the accounts present at genesis.
   *
   * @generated from field: repeated google.protobuf.Any accounts = 2;
   */
  accounts: Any[];
};

/**
 * GenesisState defines the auth module's genesis state.
 *
 * @generated from message cosmos.auth.v1beta1.GenesisState
 */
export type GenesisStateJson = {
  /**
   * params defines all the parameters of the module.
   *
   * @generated from field: cosmos.auth.v1beta1.Params params = 1;
   */
  params?: ParamsJson;

  /**
   * accounts are the accounts present at genesis.
   *
   * @generated from field: repeated google.protobuf.Any accounts = 2;
   */
  accounts?: AnyJson[];
};

/**
 * Describes the message cosmos.auth.v1beta1.GenesisState.
 * Use `create(GenesisStateSchema)` to create a new message.
 */
export const GenesisStateSchema: GenMessage<GenesisState, GenesisStateJson> = /*@__PURE__*/
  messageDesc(file_cosmos_auth_v1beta1_genesis, 0);

