import { QueryAccountRequest, QueryAccountResponse, QueryDisputesRequest, QueryDisputesResponse, QueryParamsRequest, QueryParamsResponse, QueryReportersRequest, QueryReportersResponse, QueryStateRequest, QueryStateResponse, QueryTransferRequest, QueryTransferResponse, QueryTransfersRequest, QueryTransfersResponse } from "./query.ts";

export const Query = {
  typeName: "akash.sink.v1.Query",
  methods: {
    transfer: {
      name: "Transfer",
      httpPath: "/akash/sink/v1/transfers/{id.owner}/{id.height}/{id.sequence}",
      input: QueryTransferRequest,
      output: QueryTransferResponse,
      get parent() { return Query; },
    },
    transfers: {
      name: "Transfers",
      httpPath: "/akash/sink/v1/transfers",
      input: QueryTransfersRequest,
      output: QueryTransfersResponse,
      get parent() { return Query; },
    },
    disputes: {
      name: "Disputes",
      httpPath: "/akash/sink/v1/disputes",
      input: QueryDisputesRequest,
      output: QueryDisputesResponse,
      get parent() { return Query; },
    },
    account: {
      name: "Account",
      httpPath: "/akash/sink/v1/accounts/{owner}",
      input: QueryAccountRequest,
      output: QueryAccountResponse,
      get parent() { return Query; },
    },
    state: {
      name: "State",
      httpPath: "/akash/sink/v1/state",
      input: QueryStateRequest,
      output: QueryStateResponse,
      get parent() { return Query; },
    },
    params: {
      name: "Params",
      httpPath: "/akash/sink/v1/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
    reporters: {
      name: "Reporters",
      httpPath: "/akash/sink/v1/reporters",
      input: QueryReportersRequest,
      output: QueryReportersResponse,
      get parent() { return Query; },
    },
  },
} as const;
