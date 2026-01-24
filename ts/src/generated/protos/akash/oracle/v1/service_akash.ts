import { MsgAddPriceEntry, MsgAddPriceEntryResponse, MsgUpdateParams, MsgUpdateParamsResponse } from "./msgs.ts";

export const Msg = {
  typeName: "akash.oracle.v1.Msg",
  methods: {
    addPriceEntry: {
      name: "AddPriceEntry",
      input: MsgAddPriceEntry,
      output: MsgAddPriceEntryResponse,
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
