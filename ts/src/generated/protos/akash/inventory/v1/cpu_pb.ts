// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/inventory/v1/cpu.proto (package akash.inventory.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { ResourcePair, ResourcePairJson } from "./resourcepair_pb.ts";
import { file_akash_inventory_v1_resourcepair } from "./resourcepair_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/inventory/v1/cpu.proto.
 */
export const file_akash_inventory_v1_cpu: GenFile = /*@__PURE__*/
  fileDesc("Chxha2FzaC9pbnZlbnRvcnkvdjEvY3B1LnByb3RvEhJha2FzaC5pbnZlbnRvcnkudjEi0QEKB0NQVUluZm8SJQoCaWQYASABKAlCGeLeHwJJROreHwJpZPLeHwl5YW1sOiJpZCISNQoGdmVuZG9yGAIgASgJQiXi3h8GVmVuZG9y6t4fBnZlbmRvcvLeHw15YW1sOiJ2ZW5kb3IiEjEKBW1vZGVsGAMgASgJQiLi3h8FTW9kZWzq3h8FbW9kZWzy3h8MeWFtbDoibW9kZWwiEjUKBnZjb3JlcxgEIAEoDUIl4t4fBlZjb3Jlc+reHwZ2Y29yZXPy3h8NeWFtbDoidmNvcmVzIiLGAQoDQ1BVEmMKCHF1YW50aXR5GAEgASgLMiAuYWthc2guaW52ZW50b3J5LnYxLlJlc291cmNlUGFpckIvyN4fAOLeHwhRdWFudGl0eereHwhxdWFudGl0efLeHw95YW1sOiJxdWFudGl0eSISWgoEaW5mbxgCIAMoCzIbLmFrYXNoLmludmVudG9yeS52MS5DUFVJbmZvQi/I3h8A4t4fBEluZm/q3h8EaW5mb/LeHwt5YW1sOiJpbmZvIqrfHwhDUFVJbmZvU0IdWhtwa2cuYWt0LmRldi9nby9pbnZlbnRvcnkvdjFiBnByb3RvMw", [file_gogoproto_gogo, file_akash_inventory_v1_resourcepair]);

/**
 * CPUInfo reports CPU details
 *
 * @generated from message akash.inventory.v1.CPUInfo
 */
export type CPUInfo = Message<"akash.inventory.v1.CPUInfo"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string vendor = 2;
   */
  vendor: string;

  /**
   * @generated from field: string model = 3;
   */
  model: string;

  /**
   * @generated from field: uint32 vcores = 4;
   */
  vcores: number;
};

/**
 * CPUInfo reports CPU details
 *
 * @generated from message akash.inventory.v1.CPUInfo
 */
export type CPUInfoJson = {
  /**
   * @generated from field: string id = 1;
   */
  id?: string;

  /**
   * @generated from field: string vendor = 2;
   */
  vendor?: string;

  /**
   * @generated from field: string model = 3;
   */
  model?: string;

  /**
   * @generated from field: uint32 vcores = 4;
   */
  vcores?: number;
};

/**
 * Describes the message akash.inventory.v1.CPUInfo.
 * Use `create(CPUInfoSchema)` to create a new message.
 */
export const CPUInfoSchema: GenMessage<CPUInfo, CPUInfoJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_cpu, 0);

/**
 * CPU reports CPU inventory details
 *
 * @generated from message akash.inventory.v1.CPU
 */
export type CPU = Message<"akash.inventory.v1.CPU"> & {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePair;

  /**
   * @generated from field: repeated akash.inventory.v1.CPUInfo info = 2;
   */
  info: CPUInfo[];
};

/**
 * CPU reports CPU inventory details
 *
 * @generated from message akash.inventory.v1.CPU
 */
export type CPUJson = {
  /**
   * @generated from field: akash.inventory.v1.ResourcePair quantity = 1;
   */
  quantity?: ResourcePairJson;

  /**
   * @generated from field: repeated akash.inventory.v1.CPUInfo info = 2;
   */
  info?: CPUInfoJson[];
};

/**
 * Describes the message akash.inventory.v1.CPU.
 * Use `create(CPUSchema)` to create a new message.
 */
export const CPUSchema: GenMessage<CPU, CPUJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_cpu, 1);

