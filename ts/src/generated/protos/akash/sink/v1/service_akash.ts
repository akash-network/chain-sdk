import { MsgConfirmMint, MsgConfirmMintResponse, MsgConfirmVoid, MsgConfirmVoidResponse, MsgFinalizeRefund, MsgFinalizeRefundResponse, MsgGuardianPause, MsgGuardianPauseResponse, MsgReportOutcome, MsgReportOutcomeResponse, MsgRequestVoid, MsgRequestVoidResponse, MsgRestartTransfer, MsgRestartTransferResponse, MsgSink, MsgSinkResponse, MsgUpdateParams, MsgUpdateParamsResponse, MsgVetoVoid, MsgVetoVoidResponse } from "./msgs.ts";

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
    restartTransfer: {
      name: "RestartTransfer",
      input: MsgRestartTransfer,
      output: MsgRestartTransferResponse,
      get parent() { return Msg; },
    },
    confirmMint: {
      name: "ConfirmMint",
      input: MsgConfirmMint,
      output: MsgConfirmMintResponse,
      get parent() { return Msg; },
    },
    confirmVoid: {
      name: "ConfirmVoid",
      input: MsgConfirmVoid,
      output: MsgConfirmVoidResponse,
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
