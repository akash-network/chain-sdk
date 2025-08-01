// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file akash/staking/v1beta3/service.proto (package akash.staking.v1beta3, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { MsgUpdateParamsResponseSchema, MsgUpdateParamsSchema } from "./paramsmsg_pb.ts";
import { file_akash_staking_v1beta3_paramsmsg } from "./paramsmsg_pb.ts";
import { file_cosmos_msg_v1_msg } from "../../../cosmos/msg/v1/msg_pb.ts";

/**
 * Describes the file akash/staking/v1beta3/service.proto.
 */
export const file_akash_staking_v1beta3_service: GenFile = /*@__PURE__*/
  fileDesc("CiNha2FzaC9zdGFraW5nL3YxYmV0YTMvc2VydmljZS5wcm90bxIVYWthc2guc3Rha2luZy52MWJldGEzMnQKA01zZxJmCgxVcGRhdGVQYXJhbXMSJi5ha2FzaC5zdGFraW5nLnYxYmV0YTMuTXNnVXBkYXRlUGFyYW1zGi4uYWthc2guc3Rha2luZy52MWJldGEzLk1zZ1VwZGF0ZVBhcmFtc1Jlc3BvbnNlGgWA57AqAUIlWiNwa2cuYWt0LmRldi9nby9ub2RlL3N0YWtpbmcvdjFiZXRhM2IGcHJvdG8z", [file_akash_staking_v1beta3_paramsmsg, file_cosmos_msg_v1_msg]);

/**
 * Msg defines the market Msg service.
 *
 * @generated from service akash.staking.v1beta3.Msg
 */
export const Msg: GenService<{
  /**
   * UpdateParams defines a governance operation for updating the x/market module
   * parameters. The authority is hard-coded to the x/gov module account.
   *
   * Since: akash v1.0.0
   *
   * @generated from rpc akash.staking.v1beta3.Msg.UpdateParams
   */
  updateParams: {
    methodKind: "unary";
    input: typeof MsgUpdateParamsSchema;
    output: typeof MsgUpdateParamsResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_akash_staking_v1beta3_service, 0);

