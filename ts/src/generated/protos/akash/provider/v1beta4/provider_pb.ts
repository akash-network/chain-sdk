// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/provider/v1beta4/provider.proto (package akash.provider.v1beta4, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import type { Attribute, AttributeJson } from "../../base/attributes/v1/attribute_pb.ts";
import { file_akash_base_attributes_v1_attribute } from "../../base/attributes/v1/attribute_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/provider/v1beta4/provider.proto.
 */
export const file_akash_provider_v1beta4_provider: GenFile = /*@__PURE__*/
  fileDesc("CiVha2FzaC9wcm92aWRlci92MWJldGE0L3Byb3ZpZGVyLnByb3RvEhZha2FzaC5wcm92aWRlci52MWJldGE0ImkKBEluZm8SMQoFZW1haWwYASABKAlCIuLeHwVFTWFpbOreHwVlbWFpbPLeHwx5YW1sOiJlbWFpbCISLgoHd2Vic2l0ZRgCIAEoCUId6t4fB3dlYnNpdGXy3h8OeWFtbDoid2Vic2l0ZSIi9gIKCFByb3ZpZGVyEkAKBW93bmVyGAEgASgJQjHq3h8Fb3duZXLy3h8MeWFtbDoib3duZXIi0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEjwKCGhvc3RfdXJpGAIgASgJQiri3h8HSG9zdFVSSereHwhob3N0X3VyafLeHw95YW1sOiJob3N0X3VyaSISlgEKCmF0dHJpYnV0ZXMYAyADKAsyIy5ha2FzaC5iYXNlLmF0dHJpYnV0ZXMudjEuQXR0cmlidXRlQl3I3h8A6t4fCmF0dHJpYnV0ZXPy3h8ReWFtbDoiYXR0cmlidXRlcyKq3x8ycGtnLmFrdC5kZXYvZ28vbm9kZS90eXBlcy9hdHRyaWJ1dGVzL3YxLkF0dHJpYnV0ZXMSRwoEaW5mbxgEIAEoCzIcLmFrYXNoLnByb3ZpZGVyLnYxYmV0YTQuSW5mb0IbyN4fAOreHwRpbmZv8t4fC3lhbWw6ImluZm8iOgiYoB8A6KAfAEImWiRwa2cuYWt0LmRldi9nby9ub2RlL3Byb3ZpZGVyL3YxYmV0YTRiBnByb3RvMw", [file_gogoproto_gogo, file_cosmos_proto_cosmos, file_akash_base_attributes_v1_attribute]);

/**
 * Info contains information on the provider.
 *
 * @generated from message akash.provider.v1beta4.Info
 */
export type Info = Message<"akash.provider.v1beta4.Info"> & {
  /**
   * Email is the email address to contact the provider.
   *
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * Website is the URL to the landing page or socials of the provider.
   *
   * @generated from field: string website = 2;
   */
  website: string;
};

/**
 * Info contains information on the provider.
 *
 * @generated from message akash.provider.v1beta4.Info
 */
export type InfoJson = {
  /**
   * Email is the email address to contact the provider.
   *
   * @generated from field: string email = 1;
   */
  email?: string;

  /**
   * Website is the URL to the landing page or socials of the provider.
   *
   * @generated from field: string website = 2;
   */
  website?: string;
};

/**
 * Describes the message akash.provider.v1beta4.Info.
 * Use `create(InfoSchema)` to create a new message.
 */
export const InfoSchema: GenMessage<Info, InfoJson> = /*@__PURE__*/
  messageDesc(file_akash_provider_v1beta4_provider, 0);

/**
 * Provider stores owner and host details.
 * Akash providers are entities that contribute computing resources to the network.
 * They can be individuals or organizations with underutilized computing resources, such as data centers or personal servers.
 * Providers participate in the network by running the Akash node software and setting the price for their services.
 * Users can then choose a provider based on factors such as cost, performance, and location.
 *
 * @generated from message akash.provider.v1beta4.Provider
 */
export type Provider = Message<"akash.provider.v1beta4.Provider"> & {
  /**
   * Owner is the bech32 address of the account of the provider.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 1;
   */
  owner: string;

  /**
   * HostURI is the Uniform Resource Identifier for provider connection.
   * This URI is used to directly connect to the provider to perform tasks such as sending the manifest.
   *
   * @generated from field: string host_uri = 2;
   */
  hostUri: string;

  /**
   * Attributes is a list of arbitrary attribute key-value pairs.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 3;
   */
  attributes: Attribute[];

  /**
   * Info contains additional provider information.
   *
   * @generated from field: akash.provider.v1beta4.Info info = 4;
   */
  info?: Info;
};

/**
 * Provider stores owner and host details.
 * Akash providers are entities that contribute computing resources to the network.
 * They can be individuals or organizations with underutilized computing resources, such as data centers or personal servers.
 * Providers participate in the network by running the Akash node software and setting the price for their services.
 * Users can then choose a provider based on factors such as cost, performance, and location.
 *
 * @generated from message akash.provider.v1beta4.Provider
 */
export type ProviderJson = {
  /**
   * Owner is the bech32 address of the account of the provider.
   * It is a string representing a valid account address.
   *
   * Example:
   *   "akash1..."
   *
   * @generated from field: string owner = 1;
   */
  owner?: string;

  /**
   * HostURI is the Uniform Resource Identifier for provider connection.
   * This URI is used to directly connect to the provider to perform tasks such as sending the manifest.
   *
   * @generated from field: string host_uri = 2;
   */
  hostUri?: string;

  /**
   * Attributes is a list of arbitrary attribute key-value pairs.
   *
   * @generated from field: repeated akash.base.attributes.v1.Attribute attributes = 3;
   */
  attributes?: AttributeJson[];

  /**
   * Info contains additional provider information.
   *
   * @generated from field: akash.provider.v1beta4.Info info = 4;
   */
  info?: InfoJson;
};

/**
 * Describes the message akash.provider.v1beta4.Provider.
 * Use `create(ProviderSchema)` to create a new message.
 */
export const ProviderSchema: GenMessage<Provider, ProviderJson> = /*@__PURE__*/
  messageDesc(file_akash_provider_v1beta4_provider, 1);

