// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/base/resources/v1beta4/cpu.proto (package akash.base.resources.v1beta4, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../../gogoproto/gogo_pb.ts";
import type { Attribute, AttributeJson } from "../../attributes/v1/attribute_pb.ts";
import { file_akash_base_attributes_v1_attribute } from "../../attributes/v1/attribute_pb.ts";
import type { ResourceValue, ResourceValueJson } from "./resourcevalue_pb.ts";
import { file_akash_base_resources_v1beta4_resourcevalue } from "./resourcevalue_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/base/resources/v1beta4/cpu.proto.
 */
export const file_akash_base_resources_v1beta4_cpu: GenFile = /*@__PURE__*/
  fileDesc("CiZha2FzaC9iYXNlL3Jlc291cmNlcy92MWJldGE0L2NwdS5wcm90bxIcYWthc2guYmFzZS5yZXNvdXJjZXMudjFiZXRhNCL6AQoDQ1BVEkAKBXVuaXRzGAEgASgLMisuYWthc2guYmFzZS5yZXNvdXJjZXMudjFiZXRhNC5SZXNvdXJjZVZhbHVlQgTI3h8AEqoBCgphdHRyaWJ1dGVzGAIgAygLMiMuYWthc2guYmFzZS5hdHRyaWJ1dGVzLnYxLkF0dHJpYnV0ZUJxyN4fAOreHxRhdHRyaWJ1dGVzLG9taXRlbXB0efLeHxt5YW1sOiJhdHRyaWJ1dGVzLG9taXRlbXB0eSKq3x8ycGtnLmFrdC5kZXYvZ28vbm9kZS90eXBlcy9hdHRyaWJ1dGVzL3YxLkF0dHJpYnV0ZXM6BOigHwFCLVorcGtnLmFrdC5kZXYvZ28vbm9kZS90eXBlcy9yZXNvdXJjZXMvdjFiZXRhNGIGcHJvdG8z", [file_gogoproto_gogo, file_akash_base_attributes_v1_attribute, file_akash_base_resources_v1beta4_resourcevalue]);

/**
 * CPU stores resource units and cpu config attributes.
 *
 * @generated from message akash.base.resources.v1beta4.CPU
 */
export type CPU = Message<"akash.base.resources.v1beta4.CPU"> & {
  /**
   * Units of the CPU, which represents the number of CPUs available.
   * This field is required and must be a non-negative integer.
   *
   * @generated from field: akash.base.resources.v1beta4.ResourceValue units = 1;
   */
  units?: ResourceValue;

  /**
   * Attributes holds a list of key-value attributes that describe the GPU, such as its model, memory and interface.
   * This field is required and must be a list of `Attribute` messages.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 2;
   */
  attributes: Attribute[];
};

/**
 * CPU stores resource units and cpu config attributes.
 *
 * @generated from message akash.base.resources.v1beta4.CPU
 */
export type CPUJson = {
  /**
   * Units of the CPU, which represents the number of CPUs available.
   * This field is required and must be a non-negative integer.
   *
   * @generated from field: akash.base.resources.v1beta4.ResourceValue units = 1;
   */
  units?: ResourceValueJson;

  /**
   * Attributes holds a list of key-value attributes that describe the GPU, such as its model, memory and interface.
   * This field is required and must be a list of `Attribute` messages.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 2;
   */
  attributes?: AttributeJson[];
};

/**
 * Describes the message akash.base.resources.v1beta4.CPU.
 * Use `create(CPUSchema)` to create a new message.
 */
export const CPUSchema: GenMessage<CPU, CPUJson> = /*@__PURE__*/
  messageDesc(file_akash_base_resources_v1beta4_cpu, 0);

