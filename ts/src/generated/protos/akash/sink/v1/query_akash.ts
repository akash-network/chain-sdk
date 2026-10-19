import { QueryAccountRequest, QueryAccountResponse, QueryOpenTransfersRequest, QueryOpenTransfersResponse, QuerySink, QuerySinkResponse, QueryStateRequest, QueryStateResponse, QueryTransferRequest, QueryTransferResponse } from "./query.ts";

export const Query = {
  typeName: "akash.sink.v1.Query",
  methods: {
    sink: {
      name: "Sink",
      httpPath: "/akash/sink/v1/records/{id}",
      input: QuerySink,
      output: QuerySinkResponse,
      get parent() { return Query; },
    },
    state: {
      name: "State",
      httpPath: "/akash/sink/v1/state",
      input: QueryStateRequest,
      output: QueryStateResponse,
      get parent() { return Query; },
    },
    transfer: {
      name: "Transfer",
      httpPath: "/akash/sink/v1/transfers/{id}",
      input: QueryTransferRequest,
      output: QueryTransferResponse,
      get parent() { return Query; },
    },
    account: {
      name: "Account",
      httpPath: "/akash/sink/v1/accounts/{sender}",
      input: QueryAccountRequest,
      output: QueryAccountResponse,
      get parent() { return Query; },
    },
    openTransfers: {
      name: "OpenTransfers",
      httpPath: "/akash/sink/v1/open_transfers",
      input: QueryOpenTransfersRequest,
      output: QueryOpenTransfersResponse,
      get parent() { return Query; },
    },
  },
} as const;
