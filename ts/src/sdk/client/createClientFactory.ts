import type { Transport } from "../transport/types.ts";
import type { Client, ServiceClientOptions } from "./createServiceClient.ts";
import { createServiceClient } from "./createServiceClient.ts";
import type { ServiceDesc } from "./types.ts";

export function createClientFactory<TCallOptions>(transport: Transport, options?: ServiceClientOptions) {
  const services: Record<string, Client<ServiceDesc, TCallOptions>> = Object.create(null);

  return function getClient<T extends ServiceDesc>(service: T): Client<T, TCallOptions> {
    if (!services[service.typeName]) {
      services[service.typeName] = createServiceClient<T, TCallOptions>(service, transport, options);
    }
    return services[service.typeName] as Client<T, TCallOptions>;
  };
}
