import { QueryPricesRequest, QueryPricesResponse } from "./prices.ts";
import { QueryAggregatedPriceRequest, QueryAggregatedPriceResponse, QueryParamsRequest, QueryParamsResponse, QueryPriceFeedConfigRequest, QueryPriceFeedConfigResponse } from "./query.ts";

export const Query = {
  typeName: "akash.oracle.v1.Query",
  methods: {
    prices: {
      name: "Prices",
      httpPath: "/akash/oracle/v1/prices",
      input: QueryPricesRequest,
      output: QueryPricesResponse,
      get parent() { return Query; },
    },
    params: {
      name: "Params",
      httpPath: "/akash/oracle/v1/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
    priceFeedConfig: {
      name: "PriceFeedConfig",
      httpPath: "/akash/oracle/v1/price_feed_config/{denom}",
      input: QueryPriceFeedConfigRequest,
      output: QueryPriceFeedConfigResponse,
      get parent() { return Query; },
    },
    aggregatedPrice: {
      name: "AggregatedPrice",
      httpPath: "/akash/oracle/v1/aggregated_price/{denom}",
      input: QueryAggregatedPriceRequest,
      output: QueryAggregatedPriceResponse,
      get parent() { return Query; },
    },
  },
} as const;
