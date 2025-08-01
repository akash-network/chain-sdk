// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/base/resources/v1beta4/storage.proto (package akash.base.resources.v1beta4, syntax proto3)
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
 * Describes the file akash/base/resources/v1beta4/storage.proto.
 */
export const file_akash_base_resources_v1beta4_storage: GenFile = /*@__PURE__*/
  fileDesc("Cipha2FzaC9iYXNlL3Jlc291cmNlcy92MWJldGE0L3N0b3JhZ2UucHJvdG8SHGFrYXNoLmJhc2UucmVzb3VyY2VzLnYxYmV0YTQivwIKB1N0b3JhZ2USJQoEbmFtZRgBIAEoCUIX6t4fBG5hbWXy3h8LeWFtbDoibmFtZSISWgoIcXVhbnRpdHkYAiABKAsyKy5ha2FzaC5iYXNlLnJlc291cmNlcy52MWJldGE0LlJlc291cmNlVmFsdWVCG8jeHwDq3h8Ec2l6ZfLeHwt5YW1sOiJzaXplIhKqAQoKYXR0cmlidXRlcxgDIAMoCzIjLmFrYXNoLmJhc2UuYXR0cmlidXRlcy52MS5BdHRyaWJ1dGVCccjeHwDq3h8UYXR0cmlidXRlcyxvbWl0ZW1wdHny3h8beWFtbDoiYXR0cmlidXRlcyxvbWl0ZW1wdHkiqt8fMnBrZy5ha3QuZGV2L2dvL25vZGUvdHlwZXMvYXR0cmlidXRlcy92MS5BdHRyaWJ1dGVzOgTooB8BQi1aK3BrZy5ha3QuZGV2L2dvL25vZGUvdHlwZXMvcmVzb3VyY2VzL3YxYmV0YTRiBnByb3RvMw", [file_gogoproto_gogo, file_akash_base_attributes_v1_attribute, file_akash_base_resources_v1beta4_resourcevalue]);

/**
 * Storage stores resource quantity and storage attributes.
 *
 * @generated from message akash.base.resources.v1beta4.Storage
 */
export type Storage = Message<"akash.base.resources.v1beta4.Storage"> & {
  /**
   * Name holds an arbitrary name for the storage resource.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * Quantity of storage available, which represents the amount of memory in bytes.
   * This field is required and must be a non-negative integer.
   *
   * @generated from field: akash.base.resources.v1beta4.ResourceValue quantity = 2;
   */
  quantity?: ResourceValue;

  /**
   * Attributes that describe the storage.
   * This field is required and must be a list of Attribute key-values.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 3;
   */
  attributes: Attribute[];
};

/**
 * Storage stores resource quantity and storage attributes.
 *
 * @generated from message akash.base.resources.v1beta4.Storage
 */
export type StorageJson = {
  /**
   * Name holds an arbitrary name for the storage resource.
   *
   * @generated from field: string name = 1;
   */
  name?: string;

  /**
   * Quantity of storage available, which represents the amount of memory in bytes.
   * This field is required and must be a non-negative integer.
   *
   * @generated from field: akash.base.resources.v1beta4.ResourceValue quantity = 2;
   */
  quantity?: ResourceValueJson;

  /**
   * Attributes that describe the storage.
   * This field is required and must be a list of Attribute key-values.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 3;
   */
  attributes?: AttributeJson[];
};

/**
 * Describes the message akash.base.resources.v1beta4.Storage.
 * Use `create(StorageSchema)` to create a new message.
 */
export const StorageSchema: GenMessage<Storage, StorageJson> = /*@__PURE__*/
  messageDesc(file_akash_base_resources_v1beta4_storage, 0);

