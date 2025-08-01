// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/nft/module/v1/module.proto (package cosmos.nft.module.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_cosmos_app_v1alpha1_module } from "../../../app/v1alpha1/module_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/nft/module/v1/module.proto.
 */
export const file_cosmos_nft_module_v1_module: GenFile = /*@__PURE__*/
  fileDesc("CiFjb3Ntb3MvbmZ0L21vZHVsZS92MS9tb2R1bGUucHJvdG8SFGNvc21vcy5uZnQubW9kdWxlLnYxIjQKBk1vZHVsZToqusCW2gEkCiJnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvbmZ0YgZwcm90bzM", [file_cosmos_app_v1alpha1_module]);

/**
 * Module is the config object of the nft module.
 *
 * @generated from message cosmos.nft.module.v1.Module
 */
export type Module = Message<"cosmos.nft.module.v1.Module"> & {
};

/**
 * Module is the config object of the nft module.
 *
 * @generated from message cosmos.nft.module.v1.Module
 */
export type ModuleJson = {
};

/**
 * Describes the message cosmos.nft.module.v1.Module.
 * Use `create(ModuleSchema)` to create a new message.
 */
export const ModuleSchema: GenMessage<Module, ModuleJson> = /*@__PURE__*/
  messageDesc(file_cosmos_nft_module_v1_module, 0);

