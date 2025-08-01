// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/inventory/v1/storage.proto (package akash.inventory.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { ResourcePair, ResourcePairJson } from "./resourcepair_pb.ts";
import { file_akash_inventory_v1_resourcepair } from "./resourcepair_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/inventory/v1/storage.proto.
 */
export const file_akash_inventory_v1_storage: GenFile = /*@__PURE__*/
  fileDesc("CiBha2FzaC9pbnZlbnRvcnkvdjEvc3RvcmFnZS5wcm90bxISYWthc2guaW52ZW50b3J5LnYxIm8KC1N0b3JhZ2VJbmZvEjEKBWNsYXNzGAEgASgJQiLi3h8FQ2xhc3Pq3h8FY2xhc3Py3h8MeWFtbDoiY2xhc3MiEi0KBGlvcHMYAiABKAlCH+LeHwRJT1BT6t4fBGlvcHPy3h8LeWFtbDoiaW9wcyIiwgEKB1N0b3JhZ2USYwoIcXVhbnRpdHkYASABKAsyIC5ha2FzaC5pbnZlbnRvcnkudjEuUmVzb3VyY2VQYWlyQi/I3h8A4t4fCFF1YW50aXR56t4fCHF1YW50aXR58t4fD3lhbWw6InF1YW50aXR5IhJSCgRpbmZvGAIgASgLMh8uYWthc2guaW52ZW50b3J5LnYxLlN0b3JhZ2VJbmZvQiPI3h8A4t4fBEluZm/q3h8EaW5mb/LeHwt5YW1sOiJpbmZvIkIdWhtwa2cuYWt0LmRldi9nby9pbnZlbnRvcnkvdjFiBnByb3RvMw", [file_gogoproto_gogo, file_akash_inventory_v1_resourcepair]);

/**
 * StorageInfo reports Storage details
 *
 * @generated from message akash.inventory.v1.StorageInfo
 */
export type StorageInfo = Message<"akash.inventory.v1.StorageInfo"> & {
  /**
   * @generated from field: string class = 1;
   */
  class: string;

  /**
   * @generated from field: string iops = 2;
   */
  iops: string;
};

/**
 * StorageInfo reports Storage details
 *
 * @generated from message akash.inventory.v1.StorageInfo
 */
export type StorageInfoJson = {
  /**
   * @generated from field: string class = 1;
   */
  class?: string;

  /**
   * @generated from field: string iops = 2;
   */
  iops?: string;
};

/**
 * Describes the message akash.inventory.v1.StorageInfo.
 * Use `create(StorageInfoSchema)` to create a new message.
 */
export const StorageInfoSchema: GenMessage<StorageInfo, StorageInfoJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_storage, 0);

/**
 * Storage reports Storage inventory details
 *
 * @generated from message akash.inventory.v1.Storage
 */
export type Storage = Message<"akash.inventory.v1.Storage"> & {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePair;

  /**
   * @generated from field: akash.inventory.v1.StorageInfo info = 2;
   */
  info?: StorageInfo;
};

/**
 * Storage reports Storage inventory details
 *
 * @generated from message akash.inventory.v1.Storage
 */
export type StorageJson = {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePairJson;

  /**
   * @generated from field: akash.inventory.v1.StorageInfo info = 2;
   */
  info?: StorageInfoJson;
};

/**
 * Describes the message akash.inventory.v1.Storage.
 * Use `create(StorageSchema)` to create a new message.
 */
export const StorageSchema: GenMessage<Storage, StorageJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_storage, 1);

