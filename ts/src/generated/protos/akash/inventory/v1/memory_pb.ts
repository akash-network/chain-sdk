// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/inventory/v1/memory.proto (package akash.inventory.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { ResourcePair, ResourcePairJson } from "./resourcepair_pb.ts";
import { file_akash_inventory_v1_resourcepair } from "./resourcepair_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/inventory/v1/memory.proto.
 */
export const file_akash_inventory_v1_memory: GenFile = /*@__PURE__*/
  fileDesc("Ch9ha2FzaC9pbnZlbnRvcnkvdjEvbWVtb3J5LnByb3RvEhJha2FzaC5pbnZlbnRvcnkudjEi6wEKCk1lbW9yeUluZm8SNQoGdmVuZG9yGAEgASgJQiXi3h8GVmVuZG9y6t4fBnZlbmRvcvLeHw15YW1sOiJ2ZW5kb3IiEi0KBHR5cGUYAiABKAlCH+LeHwRUeXBl6t4fBHR5cGXy3h8LeWFtbDoidHlwZSISRAoKdG90YWxfc2l6ZRgDIAEoCUIw4t4fCVRvdGFsU2l6ZereHwp0b3RhbF9zaXpl8t4fEXlhbWw6InRvdGFsX3NpemUiEjEKBXNwZWVkGAQgASgJQiLi3h8FU3BlZWTq3h8Fc3BlZWTy3h8MeWFtbDoic3BlZWQiIs8BCgZNZW1vcnkSYwoIcXVhbnRpdHkYASABKAsyIC5ha2FzaC5pbnZlbnRvcnkudjEuUmVzb3VyY2VQYWlyQi/I3h8A4t4fCFF1YW50aXR56t4fCHF1YW50aXR58t4fD3lhbWw6InF1YW50aXR5IhJgCgRpbmZvGAIgAygLMh4uYWthc2guaW52ZW50b3J5LnYxLk1lbW9yeUluZm9CMsjeHwDi3h8ESW5mb+reHwRpbmZv8t4fC3lhbWw6ImluZm8iqt8fC01lbW9yeUluZm9TQh1aG3BrZy5ha3QuZGV2L2dvL2ludmVudG9yeS92MWIGcHJvdG8z", [file_gogoproto_gogo, file_akash_inventory_v1_resourcepair]);

/**
 * MemoryInfo reports Memory details
 *
 * @generated from message akash.inventory.v1.MemoryInfo
 */
export type MemoryInfo = Message<"akash.inventory.v1.MemoryInfo"> & {
  /**
   * @generated from field: string vendor = 1;
   */
  vendor: string;

  /**
   * @generated from field: string type = 2;
   */
  type: string;

  /**
   * @generated from field: string total_size = 3;
   */
  totalSize: string;

  /**
   * @generated from field: string speed = 4;
   */
  speed: string;
};

/**
 * MemoryInfo reports Memory details
 *
 * @generated from message akash.inventory.v1.MemoryInfo
 */
export type MemoryInfoJson = {
  /**
   * @generated from field: string vendor = 1;
   */
  vendor?: string;

  /**
   * @generated from field: string type = 2;
   */
  type?: string;

  /**
   * @generated from field: string total_size = 3;
   */
  totalSize?: string;

  /**
   * @generated from field: string speed = 4;
   */
  speed?: string;
};

/**
 * Describes the message akash.inventory.v1.MemoryInfo.
 * Use `create(MemoryInfoSchema)` to create a new message.
 */
export const MemoryInfoSchema: GenMessage<MemoryInfo, MemoryInfoJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_memory, 0);

/**
 * Memory reports Memory inventory details
 *
 * @generated from message akash.inventory.v1.Memory
 */
export type Memory = Message<"akash.inventory.v1.Memory"> & {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePair;

  /**
   * @generated from field: repeated akash.inventory.v1.MemoryInfo info = 2;
   */
  info: MemoryInfo[];
};

/**
 * Memory reports Memory inventory details
 *
 * @generated from message akash.inventory.v1.Memory
 */
export type MemoryJson = {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePairJson;

  /**
   * @generated from field: repeated akash.inventory.v1.MemoryInfo info = 2;
   */
  info?: MemoryInfoJson[];
};

/**
 * Describes the message akash.inventory.v1.Memory.
 * Use `create(MemorySchema)` to create a new message.
 */
export const MemorySchema: GenMessage<Memory, MemoryJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_memory, 1);

