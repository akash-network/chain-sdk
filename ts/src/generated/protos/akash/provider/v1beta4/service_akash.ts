import { MsgCloseProviderMaintenance, MsgCloseProviderMaintenanceResponse, MsgCreateProvider, MsgCreateProviderResponse, MsgDeleteProvider, MsgDeleteProviderResponse, MsgOpenProviderMaintenance, MsgOpenProviderMaintenanceResponse, MsgUpdateProvider, MsgUpdateProviderResponse } from "./msg.ts";
import { MsgUpdateParams, MsgUpdateParamsResponse } from "./paramsmsg.ts";

export const Msg = {
  typeName: "akash.provider.v1beta4.Msg",
  methods: {
    createProvider: {
      name: "CreateProvider",
      input: MsgCreateProvider,
      output: MsgCreateProviderResponse,
      get parent() { return Msg; },
    },
    updateProvider: {
      name: "UpdateProvider",
      input: MsgUpdateProvider,
      output: MsgUpdateProviderResponse,
      get parent() { return Msg; },
    },
    deleteProvider: {
      name: "DeleteProvider",
      input: MsgDeleteProvider,
      output: MsgDeleteProviderResponse,
      get parent() { return Msg; },
    },
    openProviderMaintenance: {
      name: "OpenProviderMaintenance",
      input: MsgOpenProviderMaintenance,
      output: MsgOpenProviderMaintenanceResponse,
      get parent() { return Msg; },
    },
    closeProviderMaintenance: {
      name: "CloseProviderMaintenance",
      input: MsgCloseProviderMaintenance,
      output: MsgCloseProviderMaintenanceResponse,
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
