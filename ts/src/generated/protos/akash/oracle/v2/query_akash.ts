import { QueryAggregatedPriceRequest, QueryAggregatedPriceResponse, QueryParamsRequest, QueryParamsResponse, QueryPricesRequest, QueryPricesResponse } from "./query.ts";

export const Query = {
  typeName: "akash.oracle.v2.Query",
  methods: {
    prices: {
      name: "Prices",
      httpPath: "/akash/oracle/v2/prices",
      input: QueryPricesRequest,
      output: QueryPricesResponse,
      get parent() { return Query; },
    },
    params: {
      name: "Params",
      httpPath: "/akash/oracle/v2/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
    aggregatedPrice: {
      name: "AggregatedPrice",
      httpPath: "/akash/oracle/v2/aggregated_price/{denom=**}",
      input: QueryAggregatedPriceRequest,
      output: QueryAggregatedPriceResponse,
      get parent() { return Query; },
    },
  },
} as const;
