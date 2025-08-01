// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/deployment/v1beta4/groupspec.proto (package akash.deployment.v1beta4, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { PlacementRequirements, PlacementRequirementsJson } from "../../base/attributes/v1/attribute_pb.ts";
import { file_akash_base_attributes_v1_attribute } from "../../base/attributes/v1/attribute_pb.ts";
import type { ResourceUnit, ResourceUnitJson } from "./resourceunit_pb.ts";
import { file_akash_deployment_v1beta4_resourceunit } from "./resourceunit_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/deployment/v1beta4/groupspec.proto.
 */
export const file_akash_deployment_v1beta4_groupspec: GenFile = /*@__PURE__*/
  fileDesc("Cihha2FzaC9kZXBsb3ltZW50L3YxYmV0YTQvZ3JvdXBzcGVjLnByb3RvEhhha2FzaC5kZXBsb3ltZW50LnYxYmV0YTQiowIKCUdyb3VwU3BlYxIlCgRuYW1lGAEgASgJQhfq3h8EbmFtZfLeHwt5YW1sOiJuYW1lIhJyCgxyZXF1aXJlbWVudHMYAiABKAsyLy5ha2FzaC5iYXNlLmF0dHJpYnV0ZXMudjEuUGxhY2VtZW50UmVxdWlyZW1lbnRzQivI3h8A6t4fDHJlcXVpcmVtZW50c/LeHxN5YW1sOiJyZXF1aXJlbWVudHMiEnEKCXJlc291cmNlcxgDIAMoCzImLmFrYXNoLmRlcGxveW1lbnQudjFiZXRhNC5SZXNvdXJjZVVuaXRCNsjeHwDq3h8JcmVzb3VyY2Vz8t4fEHlhbWw6InJlc291cmNlcyKq3x8NUmVzb3VyY2VVbml0czoIiKAfAOigHwBCKFomcGtnLmFrdC5kZXYvZ28vbm9kZS9kZXBsb3ltZW50L3YxYmV0YTRiBnByb3RvMw", [file_gogoproto_gogo, file_akash_base_attributes_v1_attribute, file_akash_deployment_v1beta4_resourceunit]);

/**
 * GroupSpec defines a specification for a group in a deployment on the network.
 * This includes attributes like the group name, placement requirements, and resource units.
 *
 * @generated from message akash.deployment.v1beta4.GroupSpec
 */
export type GroupSpec = Message<"akash.deployment.v1beta4.GroupSpec"> & {
  /**
   * Name is the name of the group.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * Requirements specifies the placement requirements for the group.
   * This determines where the resources in the group can be deployed.
   *
   * @generated from field: akash.base.attributes.v1.PlacementRequirements requirements = 2;
   */
  requirements?: PlacementRequirements;

  /**
   * Resources is a list containing the resource units allocated to the group.
   * Each ResourceUnit defines the specific resources (e.g., CPU, memory) assigned.
   *
   * @generated from field: repeated akash.deployment.v1beta4.ResourceUnit resources = 3;
   */
  resources: ResourceUnit[];
};

/**
 * GroupSpec defines a specification for a group in a deployment on the network.
 * This includes attributes like the group name, placement requirements, and resource units.
 *
 * @generated from message akash.deployment.v1beta4.GroupSpec
 */
export type GroupSpecJson = {
  /**
   * Name is the name of the group.
   *
   * @generated from field: string name = 1;
   */
  name?: string;

  /**
   * Requirements specifies the placement requirements for the group.
   * This determines where the resources in the group can be deployed.
   *
   * @generated from field: akash.base.attributes.v1.PlacementRequirements requirements = 2;
   */
  requirements?: PlacementRequirementsJson;

  /**
   * Resources is a list containing the resource units allocated to the group.
   * Each ResourceUnit defines the specific resources (e.g., CPU, memory) assigned.
   *
   * @generated from field: repeated akash.deployment.v1beta4.ResourceUnit resources = 3;
   */
  resources?: ResourceUnitJson[];
};

/**
 * Describes the message akash.deployment.v1beta4.GroupSpec.
 * Use `create(GroupSpecSchema)` to create a new message.
 */
export const GroupSpecSchema: GenMessage<GroupSpec, GroupSpecJson> = /*@__PURE__*/
  messageDesc(file_akash_deployment_v1beta4_groupspec, 0);

