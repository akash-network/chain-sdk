import { MsgFinalizeRefund, MsgFinalizeRefundResponse, MsgGuardianPause, MsgGuardianPauseResponse, MsgReportOutcome, MsgReportOutcomeResponse, MsgRequestVoid, MsgRequestVoidResponse, MsgResolveTransfer, MsgResolveTransferResponse, MsgSink, MsgSinkResponse, MsgUpdateParams, MsgUpdateParamsResponse, MsgVetoVoid, MsgVetoVoidResponse } from "./msgs.ts";

export const Msg = {
  typeName: "akash.sink.v1.Msg",
  methods: {
    sink: {
      name: "Sink",
      input: MsgSink,
      output: MsgSinkResponse,
      get parent() { return Msg; },
    },
    updateParams: {
      name: "UpdateParams",
      input: MsgUpdateParams,
      output: MsgUpdateParamsResponse,
      get parent() { return Msg; },
    },
    requestVoid: {
      name: "RequestVoid",
      input: MsgRequestVoid,
      output: MsgRequestVoidResponse,
      get parent() { return Msg; },
    },
    reportOutcome: {
      name: "ReportOutcome",
      input: MsgReportOutcome,
      output: MsgReportOutcomeResponse,
      get parent() { return Msg; },
    },
    vetoVoid: {
      name: "VetoVoid",
      input: MsgVetoVoid,
      output: MsgVetoVoidResponse,
      get parent() { return Msg; },
    },
    resolveTransfer: {
      name: "ResolveTransfer",
      input: MsgResolveTransfer,
      output: MsgResolveTransferResponse,
      get parent() { return Msg; },
    },
    finalizeRefund: {
      name: "FinalizeRefund",
      input: MsgFinalizeRefund,
      output: MsgFinalizeRefundResponse,
      get parent() { return Msg; },
    },
    guardianPause: {
      name: "GuardianPause",
      input: MsgGuardianPause,
      output: MsgGuardianPauseResponse,
      get parent() { return Msg; },
    },
  },
} as const;
