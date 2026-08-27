import { QueryParamsRequest, QueryParamsResponse, QueryProviderMaintenanceRequest, QueryProviderMaintenanceResponse, QueryProviderMaintenancesRequest, QueryProviderMaintenancesResponse, QueryProviderRequest, QueryProviderResponse, QueryProvidersRequest, QueryProvidersResponse, QueryRegistrationRequest, QueryRegistrationResponse } from "./query.ts";

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
    providerMaintenance: {
      name: "ProviderMaintenance",
      httpPath: "/akash/provider/v1beta4/providers/{provider}/maintenance/{maintenance_id}",
      input: QueryProviderMaintenanceRequest,
      output: QueryProviderMaintenanceResponse,
      get parent() { return Query; },
    },
    providerMaintenances: {
      name: "ProviderMaintenances",
      httpPath: "/akash/provider/v1beta4/providers/{provider}/maintenance",
      input: QueryProviderMaintenancesRequest,
      output: QueryProviderMaintenancesResponse,
      get parent() { return Query; },
    },
    params: {
      name: "Params",
      httpPath: "/akash/provider/v1beta4/params",
      input: QueryParamsRequest,
      output: QueryParamsResponse,
      get parent() { return Query; },
    },
    registration: {
      name: "Registration",
      httpPath: "/akash/provider/v1beta4/providers/{provider}/registration",
      input: QueryRegistrationRequest,
      output: QueryRegistrationResponse,
      get parent() { return Query; },
    },
  },
} as const;
