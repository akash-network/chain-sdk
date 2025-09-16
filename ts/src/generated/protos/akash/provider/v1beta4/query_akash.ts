import { QueryProviderRequest, QueryProviderResponse, QueryProvidersRequest, QueryProvidersResponse } from "./query.ts";

export const Query = {
  typeName: "akash.provider.v1beta4.Query",
  methods: {
    providers: {
      name: "Providers",
      httpPath: "/akash/provider/v1beta4/providers",
      input: QueryProvidersRequest,
      output: QueryProvidersResponse,
      get parent() { return Query; },
    },
    provider: {
      name: "Provider",
      httpPath: "/akash/provider/v1beta4/providers/{owner}",
      input: QueryProviderRequest,
      output: QueryProviderResponse,
      get parent() { return Query; },
    },
  },
} as const;
