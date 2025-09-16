import { createServiceLoader } from "../client/createServiceLoader.ts";

import type * as google_protobuf_empty from "./protos/google/protobuf/empty.ts";
import type * as akash_provider_lease_v1_service from "./protos/akash/provider/lease/v1/service.ts";
import { createClientFactory } from "../client/createClientFactory.ts";
import type { Transport, CallOptions } from "../transport/types.ts";
import type { SDKOptions } from "../sdk/types.ts";
import { withMetadata } from "../utils/sdkMetadata.ts";


export const serviceLoader= createServiceLoader([
  () => import("./protos/akash/inventory/v1/service_akash.ts").then(m => m.NodeRPC),
  () => import("./protos/akash/inventory/v1/service_akash.ts").then(m => m.ClusterRPC),
  () => import("./protos/akash/provider/lease/v1/service_akash.ts").then(m => m.LeaseRPC),
  () => import("./protos/akash/provider/v1/service_akash.ts").then(m => m.ProviderRPC)
] as const);
export function createSDK(transport: Transport, options?: SDKOptions) {
  const getClient = createClientFactory<CallOptions>(transport, options?.clientOptions);
  return {
    akash: {
      inventory: {
        v1: {
          /**
           * queryNode defines a method to query hardware state of the node
           */
          queryNode: withMetadata(async function queryNode(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).queryNode(input, options);
          }, { path: [0, 0] }),
          /**
           * streamNode defines a method to stream hardware state of the node
           */
          streamNode: withMetadata(async function streamNode(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).streamNode(input, options);
          }, { path: [0, 1] }),
          /**
           * queryCluster defines a method to query hardware state of the cluster
           */
          queryCluster: withMetadata(async function queryCluster(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getClient(service).queryCluster(input, options);
          }, { path: [1, 0] }),
          /**
           * streamCluster defines a method to stream hardware state of the cluster
           */
          streamCluster: withMetadata(async function streamCluster(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getClient(service).streamCluster(input, options);
          }, { path: [1, 1] })
        }
      },
      provider: {
        lease: {
          v1: {
            /**
             * sendManifest sends manifest to the provider
             */
            sendManifest: withMetadata(async function sendManifest(input: akash_provider_lease_v1_service.SendManifestRequest, options?: CallOptions) {
              const service = await serviceLoader.loadAt(2);
              return getClient(service).sendManifest(input, options);
            }, { path: [2, 0] }),
            /**
             * serviceStatus
             */
            serviceStatus: withMetadata(async function serviceStatus(input: akash_provider_lease_v1_service.ServiceStatusRequest, options?: CallOptions) {
              const service = await serviceLoader.loadAt(2);
              return getClient(service).serviceStatus(input, options);
            }, { path: [2, 1] }),
            /**
             * streamServiceStatus
             */
            streamServiceStatus: withMetadata(async function streamServiceStatus(input: akash_provider_lease_v1_service.ServiceStatusRequest, options?: CallOptions) {
              const service = await serviceLoader.loadAt(2);
              return getClient(service).streamServiceStatus(input, options);
            }, { path: [2, 2] }),
            /**
             * serviceLogs
             */
            serviceLogs: withMetadata(async function serviceLogs(input: akash_provider_lease_v1_service.ServiceLogsRequest, options?: CallOptions) {
              const service = await serviceLoader.loadAt(2);
              return getClient(service).serviceLogs(input, options);
            }, { path: [2, 3] }),
            /**
             * streamServiceLogs
             */
            streamServiceLogs: withMetadata(async function streamServiceLogs(input: akash_provider_lease_v1_service.ServiceLogsRequest, options?: CallOptions) {
              const service = await serviceLoader.loadAt(2);
              return getClient(service).streamServiceLogs(input, options);
            }, { path: [2, 4] })
          }
        },
        v1: {
          /**
           * getStatus defines a method to query provider state
           */
          getStatus: withMetadata(async function getStatus(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getClient(service).getStatus(input, options);
          }, { path: [3, 0] }),
          /**
           * Status defines a method to stream provider state
           */
          streamStatus: withMetadata(async function streamStatus(input: google_protobuf_empty.Empty = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(3);
            return getClient(service).streamStatus(input, options);
          }, { path: [3, 1] })
        }
      }
    }
  };
}
