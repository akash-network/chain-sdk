import { MsgCancelAuditEscrow, MsgCancelAuditEscrowResponse, MsgOpenAuditEscrow, MsgOpenAuditEscrowResponse, MsgPostAuditorBond, MsgPostAuditorBondResponse, MsgPostProviderBond, MsgPostProviderBondResponse, MsgPostSnapshotHash, MsgPostSnapshotHashResponse, MsgRegisterAuditor, MsgRegisterAuditorResponse, MsgRemoveAttestation, MsgRemoveAttestationResponse, MsgRemoveAuditor, MsgRemoveAuditorResponse, MsgRenewAuditor, MsgRenewAuditorResponse, MsgResignAuditor, MsgResignAuditorResponse, MsgResolveDiscrepancy, MsgResolveDiscrepancyResponse, MsgRevokeAllProviderAttestations, MsgRevokeAllProviderAttestationsResponse, MsgRevokeAttestation, MsgRevokeAttestationResponse, MsgRevokeAuditorAttestations, MsgRevokeAuditorAttestationsResponse, MsgRevokeProviderAttestation, MsgRevokeProviderAttestationResponse, MsgSettleAuditEscrow, MsgSettleAuditEscrowResponse, MsgSlashProviderBond, MsgSlashProviderBondResponse, MsgSubmitAttestation, MsgSubmitAttestationResponse, MsgUpdateParams, MsgUpdateParamsResponse, MsgWithdrawProviderBond, MsgWithdrawProviderBondResponse } from "./msg.ts";

export const Msg = {
  typeName: "akash.verification.v1.Msg",
  methods: {
    postAuditorBond: {
      name: "PostAuditorBond",
      input: MsgPostAuditorBond,
      output: MsgPostAuditorBondResponse,
      get parent() { return Msg; },
    },
    openAuditEscrow: {
      name: "OpenAuditEscrow",
      input: MsgOpenAuditEscrow,
      output: MsgOpenAuditEscrowResponse,
      get parent() { return Msg; },
    },
    cancelAuditEscrow: {
      name: "CancelAuditEscrow",
      input: MsgCancelAuditEscrow,
      output: MsgCancelAuditEscrowResponse,
      get parent() { return Msg; },
    },
    settleAuditEscrow: {
      name: "SettleAuditEscrow",
      input: MsgSettleAuditEscrow,
      output: MsgSettleAuditEscrowResponse,
      get parent() { return Msg; },
    },
    submitAttestation: {
      name: "SubmitAttestation",
      input: MsgSubmitAttestation,
      output: MsgSubmitAttestationResponse,
      get parent() { return Msg; },
    },
    revokeAttestation: {
      name: "RevokeAttestation",
      input: MsgRevokeAttestation,
      output: MsgRevokeAttestationResponse,
      get parent() { return Msg; },
    },
    removeAttestation: {
      name: "RemoveAttestation",
      input: MsgRemoveAttestation,
      output: MsgRemoveAttestationResponse,
      get parent() { return Msg; },
    },
    resignAuditor: {
      name: "ResignAuditor",
      input: MsgResignAuditor,
      output: MsgResignAuditorResponse,
      get parent() { return Msg; },
    },
    postProviderBond: {
      name: "PostProviderBond",
      input: MsgPostProviderBond,
      output: MsgPostProviderBondResponse,
      get parent() { return Msg; },
    },
    withdrawProviderBond: {
      name: "WithdrawProviderBond",
      input: MsgWithdrawProviderBond,
      output: MsgWithdrawProviderBondResponse,
      get parent() { return Msg; },
    },
    postSnapshotHash: {
      name: "PostSnapshotHash",
      input: MsgPostSnapshotHash,
      output: MsgPostSnapshotHashResponse,
      get parent() { return Msg; },
    },
    registerAuditor: {
      name: "RegisterAuditor",
      input: MsgRegisterAuditor,
      output: MsgRegisterAuditorResponse,
      get parent() { return Msg; },
    },
    renewAuditor: {
      name: "RenewAuditor",
      input: MsgRenewAuditor,
      output: MsgRenewAuditorResponse,
      get parent() { return Msg; },
    },
    removeAuditor: {
      name: "RemoveAuditor",
      input: MsgRemoveAuditor,
      output: MsgRemoveAuditorResponse,
      get parent() { return Msg; },
    },
    revokeProviderAttestation: {
      name: "RevokeProviderAttestation",
      input: MsgRevokeProviderAttestation,
      output: MsgRevokeProviderAttestationResponse,
      get parent() { return Msg; },
    },
    revokeAllProviderAttestations: {
      name: "RevokeAllProviderAttestations",
      input: MsgRevokeAllProviderAttestations,
      output: MsgRevokeAllProviderAttestationsResponse,
      get parent() { return Msg; },
    },
    revokeAuditorAttestations: {
      name: "RevokeAuditorAttestations",
      input: MsgRevokeAuditorAttestations,
      output: MsgRevokeAuditorAttestationsResponse,
      get parent() { return Msg; },
    },
    resolveDiscrepancy: {
      name: "ResolveDiscrepancy",
      input: MsgResolveDiscrepancy,
      output: MsgResolveDiscrepancyResponse,
      get parent() { return Msg; },
    },
    slashProviderBond: {
      name: "SlashProviderBond",
      input: MsgSlashProviderBond,
      output: MsgSlashProviderBondResponse,
      get parent() { return Msg; },
    },
    updateParams: {
      name: "UpdateParams",
      input: MsgUpdateParams,
      output: MsgUpdateParamsResponse,
      get parent() { return Msg; },
    },
  },
} as const;
