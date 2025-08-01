// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/inventory/v1/node.proto (package akash.inventory.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { NodeResources, NodeResourcesJson } from "./resources_pb.ts";
import { file_akash_inventory_v1_resources } from "./resources_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/inventory/v1/node.proto.
 */
export const file_akash_inventory_v1_node: GenFile = /*@__PURE__*/
  fileDesc("Ch1ha2FzaC9pbnZlbnRvcnkvdjEvbm9kZS5wcm90bxISYWthc2guaW52ZW50b3J5LnYxImwKEE5vZGVDYXBhYmlsaXRpZXMSWAoPc3RvcmFnZV9jbGFzc2VzGAEgAygJQj/i3h8OU3RvcmFnZUNsYXNzZXPq3h8Pc3RvcmFnZV9jbGFzc2Vz8t4fFnlhbWw6InN0b3JhZ2VfY2xhc3NlcyIimAIKBE5vZGUSLQoEbmFtZRgBIAEoCUIf4t4fBE5hbWXq3h8EbmFtZfLeHwt5YW1sOiJuYW1lIhJoCglyZXNvdXJjZXMYAiABKAsyIS5ha2FzaC5pbnZlbnRvcnkudjEuTm9kZVJlc291cmNlc0IyyN4fAOLeHwlSZXNvdXJjZXPq3h8JcmVzb3VyY2Vz8t4fEHlhbWw6InJlc291cmNlcyISdwoMY2FwYWJpbGl0aWVzGAMgASgLMiQuYWthc2guaW52ZW50b3J5LnYxLk5vZGVDYXBhYmlsaXRpZXNCO8jeHwDi3h8MQ2FwYWJpbGl0aWVz6t4fDGNhcGFiaWxpdGllc/LeHxN5YW1sOiJjYXBhYmlsaXRpZXMiQh1aG3BrZy5ha3QuZGV2L2dvL2ludmVudG9yeS92MWIGcHJvdG8z", [file_gogoproto_gogo, file_akash_inventory_v1_resources]);

/**
 * NodeCapabilities extended list of node capabilities
 *
 * @generated from message akash.inventory.v1.NodeCapabilities
 */
export type NodeCapabilities = Message<"akash.inventory.v1.NodeCapabilities"> & {
  /**
   * @generated from field: repeated string storage_classes = 1;
   */
  storageClasses: string[];
};

/**
 * NodeCapabilities extended list of node capabilities
 *
 * @generated from message akash.inventory.v1.NodeCapabilities
 */
export type NodeCapabilitiesJson = {
  /**
   * @generated from field: repeated string storage_classes = 1;
   */
  storageClasses?: string[];
};

/**
 * Describes the message akash.inventory.v1.NodeCapabilities.
 * Use `create(NodeCapabilitiesSchema)` to create a new message.
 */
export const NodeCapabilitiesSchema: GenMessage<NodeCapabilities, NodeCapabilitiesJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_node, 0);

/**
 * Node reports node inventory details
 *
 * @generated from message akash.inventory.v1.Node
 */
export type Node = Message<"akash.inventory.v1.Node"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: akash.inventory.v1.NodeResources resources = 2;
   */
  resources?: NodeResources;

  /**
   * @generated from field: akash.inventory.v1.NodeCapabilities capabilities = 3;
   */
  capabilities?: NodeCapabilities;
};

/**
 * Node reports node inventory details
 *
 * @generated from message akash.inventory.v1.Node
 */
export type NodeJson = {
  /**
   * @generated from field: string name = 1;
   */
  name?: string;

  /**
   * @generated from field: akash.inventory.v1.NodeResources resources = 2;
   */
  resources?: NodeResourcesJson;

  /**
   * @generated from field: akash.inventory.v1.NodeCapabilities capabilities = 3;
   */
  capabilities?: NodeCapabilitiesJson;
};

/**
 * Describes the message akash.inventory.v1.Node.
 * Use `create(NodeSchema)` to create a new message.
 */
export const NodeSchema: GenMessage<Node, NodeJson> = /*@__PURE__*/
  messageDesc(file_akash_inventory_v1_node, 1);

