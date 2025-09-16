import { QueryBidRequest, QueryBidResponse, QueryBidsRequest, QueryBidsResponse, QueryLeaseRequest, QueryLeaseResponse, QueryLeasesRequest, QueryLeasesResponse, QueryOrderRequest, QueryOrderResponse, QueryOrdersRequest, QueryOrdersResponse, QueryParamsRequest, QueryParamsResponse } from "./query.ts";

export const Query = {
  typeName: "akash.market.v1beta5.Query",
  methods: {
    orders: {
      name: "Orders",
      httpPath: "/akash/market/v1beta5/orders/list",
      input: QueryOrdersRequest,
      output: QueryOrdersResponse,
      get parent() { return Query; },
    },
    order: {
      name: "Order",
      httpPath: "/akash/market/v1beta5/orders/info",
      input: QueryOrderRequest,
      output: QueryOrderResponse,
      get parent() { return Query; },
    },
    bids: {
      name: "Bids",
      httpPath: "/akash/market/v1beta5/bids/list",
      input: QueryBidsRequest,
      output: QueryBidsResponse,
      get parent() { return Query; },
    },
    bid: {
      name: "Bid",
      httpPath: "/akash/market/v1beta5/bids/info",
      input: QueryBidRequest,
      output: QueryBidResponse,
      get parent() { return Query; },
    },
    leases: {
      name: "Leases",
      httpPath: "/akash/market/v1beta5/leases/list",
      input: QueryLeasesRequest,
      output: QueryLeasesResponse,
      get parent() { return Query; },
    },
    lease: {
      name: "Lease",
      httpPath: "/akash/market/v1beta5/leases/info",
      input: QueryLeaseRequest,
      output: QueryLeaseResponse,
      get parent() { return Query; },
    },
    params: {
      name: "Params",
      httpPath: "/akash/market/v1beta5/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
  },
} as const;
