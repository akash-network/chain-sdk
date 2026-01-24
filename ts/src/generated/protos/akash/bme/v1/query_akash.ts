import { QueryParamsRequest, QueryParamsResponse, QueryStatusRequest, QueryStatusResponse, QueryVaultStateRequest, QueryVaultStateResponse } from "./query.ts";

export const Query = {
  typeName: "akash.bme.v1.Query",
  methods: {
    params: {
      name: "Params",
      httpPath: "/akash/bme/v1/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
    vaultState: {
      name: "VaultState",
      httpPath: "/akash/bme/v1/vault",
      input: QueryVaultStateRequest,
      output: QueryVaultStateResponse,
      get parent() { return Query; },
    },
    status: {
      name: "Status",
      httpPath: "/akash/bme/v1/status",
      input: QueryStatusRequest,
      output: QueryStatusResponse,
      get parent() { return Query; },
    },
  },
} as const;
