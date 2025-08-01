// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true,import_extension=ts"
// @generated from file cosmos/evidence/v1beta1/tx.proto (package cosmos.evidence.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb.ts";
import type { Any, AnyJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_any } from "@bufbuild/protobuf/wkt";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb.ts";
import { file_cosmos_msg_v1_msg } from "../../msg/v1/msg_pb.ts";
import { file_amino_amino } from "../../../amino/amino_pb.ts";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/evidence/v1beta1/tx.proto.
 */
export const file_cosmos_evidence_v1beta1_tx: GenFile = /*@__PURE__*/
  fileDesc("CiBjb3Ntb3MvZXZpZGVuY2UvdjFiZXRhMS90eC5wcm90bxIXY29zbW9zLmV2aWRlbmNlLnYxYmV0YTEixwEKEU1zZ1N1Ym1pdEV2aWRlbmNlEisKCXN1Ym1pdHRlchgBIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEkwKCGV2aWRlbmNlGAIgASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueUIkyrQtIGNvc21vcy5ldmlkZW5jZS52MWJldGExLkV2aWRlbmNlOjeIoB8A6KAfAILnsCoJc3VibWl0dGVyiuewKhxjb3Ntb3Mtc2RrL01zZ1N1Ym1pdEV2aWRlbmNlIikKGU1zZ1N1Ym1pdEV2aWRlbmNlUmVzcG9uc2USDAoEaGFzaBgEIAEoDDJ+CgNNc2cScAoOU3VibWl0RXZpZGVuY2USKi5jb3Ntb3MuZXZpZGVuY2UudjFiZXRhMS5Nc2dTdWJtaXRFdmlkZW5jZRoyLmNvc21vcy5ldmlkZW5jZS52MWJldGExLk1zZ1N1Ym1pdEV2aWRlbmNlUmVzcG9uc2UaBYDnsCoBQjNaLWdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsveC9ldmlkZW5jZS90eXBlc6jiHgFiBnByb3RvMw", [file_gogoproto_gogo, file_google_protobuf_any, file_cosmos_proto_cosmos, file_cosmos_msg_v1_msg, file_amino_amino]);

/**
 * MsgSubmitEvidence represents a message that supports submitting arbitrary
 * Evidence of misbehavior such as equivocation or counterfactual signing.
 *
 * @generated from message cosmos.evidence.v1beta1.MsgSubmitEvidence
 */
export type MsgSubmitEvidence = Message<"cosmos.evidence.v1beta1.MsgSubmitEvidence"> & {
  /**
   * submitter is the signer account address of evidence.
   *
   * @generated from field: string submitter = 1;
   */
  submitter: string;

  /**
   * evidence defines the evidence of misbehavior.
   *
   * @generated from field: google.protobuf.Any evidence = 2;
   */
  evidence?: Any;
};

/**
 * MsgSubmitEvidence represents a message that supports submitting arbitrary
 * Evidence of misbehavior such as equivocation or counterfactual signing.
 *
 * @generated from message cosmos.evidence.v1beta1.MsgSubmitEvidence
 */
export type MsgSubmitEvidenceJson = {
  /**
   * submitter is the signer account address of evidence.
   *
   * @generated from field: string submitter = 1;
   */
  submitter?: string;

  /**
   * evidence defines the evidence of misbehavior.
   *
   * @generated from field: google.protobuf.Any evidence = 2;
   */
  evidence?: AnyJson;
};

/**
 * Describes the message cosmos.evidence.v1beta1.MsgSubmitEvidence.
 * Use `create(MsgSubmitEvidenceSchema)` to create a new message.
 */
export const MsgSubmitEvidenceSchema: GenMessage<MsgSubmitEvidence, MsgSubmitEvidenceJson> = /*@__PURE__*/
  messageDesc(file_cosmos_evidence_v1beta1_tx, 0);

/**
 * MsgSubmitEvidenceResponse defines the Msg/SubmitEvidence response type.
 *
 * @generated from message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse
 */
export type MsgSubmitEvidenceResponse = Message<"cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"> & {
  /**
   * hash defines the hash of the evidence.
   *
   * @generated from field: bytes hash = 4;
   */
  hash: Uint8Array;
};

/**
 * MsgSubmitEvidenceResponse defines the Msg/SubmitEvidence response type.
 *
 * @generated from message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse
 */
export type MsgSubmitEvidenceResponseJson = {
  /**
   * hash defines the hash of the evidence.
   *
   * @generated from field: bytes hash = 4;
   */
  hash?: string;
};

/**
 * Describes the message cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse.
 * Use `create(MsgSubmitEvidenceResponseSchema)` to create a new message.
 */
export const MsgSubmitEvidenceResponseSchema: GenMessage<MsgSubmitEvidenceResponse, MsgSubmitEvidenceResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_evidence_v1beta1_tx, 1);

/**
 * Msg defines the evidence Msg service.
 *
 * @generated from service cosmos.evidence.v1beta1.Msg
 */
export const Msg: GenService<{
  /**
   * SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or
   * counterfactual signing.
   *
   * @generated from rpc cosmos.evidence.v1beta1.Msg.SubmitEvidence
   */
  submitEvidence: {
    methodKind: "unary";
    input: typeof MsgSubmitEvidenceSchema;
    output: typeof MsgSubmitEvidenceResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_cosmos_evidence_v1beta1_tx, 0);

