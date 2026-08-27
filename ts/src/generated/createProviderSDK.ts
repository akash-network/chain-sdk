import { createServiceLoader } from "../sdk/client/createServiceLoader.ts";

import type * as google_protobuf_empty from "./protos/google/protobuf/empty.ts";
import type * as akash_inventory_v1_snapshot from "./protos/akash/inventory/v1/snapshot.ts";
import type * as akash_provider_lease_v1_service from "./protos/akash/provider/lease/v1/service.ts";
import type * as akash_provider_v1_validation from "./protos/akash/provider/v1/validation.ts";
import { createClientFactory } from "../sdk/client/createClientFactory.ts";
import type { Transport, CallOptions } from "../sdk/transport/types.ts";
import { withMetadata } from "../sdk/client/sdkMetadata.ts";
import type { DeepPartial } from "../encoding/typeEncodingHelpers.ts";


export const serviceLoader= createServiceLoader([
  () => import("./protos/akash/inventory/v1/service_akash.ts").then(m => m.NodeRPC),
  () => import("./protos/akash/inventory/v1/service_akash.ts").then(m => m.ClusterRPC),
  () => import("./protos/akash/inventory/v1/snapshot_akash.ts").then(m => m.InventoryService),
  () => import("./protos/akash/provider/lease/v1/service_akash.ts").then(m => m.LeaseRPC),
  () => import("./protos/akash/provider/v1/service_akash.ts").then(m => m.ProviderRPC)
] as const);
export function createSDK(transport: Transport) {
  const getClient = createClientFactory<CallOptions>(transport);
  return {
    akash: {
      inventory: {
        v1: {
          /**
           * queryNode defines a method to query hardware state of the node
           */
          queryNode: withMetadata(async function queryNode(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).queryNode(input, options);
          }, { path: [0, "queryNode"], serviceLoader }),
          /**
           * streamNode defines a method to stream hardware state of the node
           */
          streamNode: withMetadata(async function streamNode(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(0);
            return getClient(service).streamNode(input, options);
          }, { path: [0, "streamNode"], serviceLoader }),
          /**
           * queryCluster defines a method to query hardware state of the cluster
           */
          queryCluster: withMetadata(async function queryCluster(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getClient(service).queryCluster(input, options);
          }, { path: [1, "queryCluster"], serviceLoader }),
          /**
           * streamCluster defines a method to stream hardware state of the cluster
           */
          streamCluster: withMetadata(async function streamCluster(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(1);
            return getClient(service).streamCluster(input, options);
          }, { path: [1, "streamCluster"], serviceLoader }),
          /**
           * getInventorySnapshot returns a fresh provider-signed live challenge snapshot.
           */
          getInventorySnapshot: withMetadata(async function getInventorySnapshot(input: DeepPartial<akash_inventory_v1_snapshot.GetInventorySnapshotRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).getInventorySnapshot(input, options);
          }, { path: [2, "getInventorySnapshot"], serviceLoader }),
          /**
           * getCommittedInventorySnapshot returns an exact provider-signed committed
           * snapshot payload by hash, or the latest committed snapshot when no hash is
           * provided.
           */
          getCommittedInventorySnapshot: withMetadata(async function getCommittedInventorySnapshot(input: DeepPartial<akash_inventory_v1_snapshot.GetCommittedInventorySnapshotRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(2);
            return getClient(service).getCommittedInventorySnapshot(input, options);
          }, { path: [2, "getCommittedInventorySnapshot"], serviceLoader })
        }
      },
      provider: {
        lease: {
          v1: {
            /**
             * sendManifest sends manifest to the provider
             */
            sendManifest: withMetadata(async function sendManifest(input: DeepPartial<akash_provider_lease_v1_service.SendManifestRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).sendManifest(input, options);
            }, { path: [3, "sendManifest"], serviceLoader }),
            /**
             * serviceStatus
             */
            serviceStatus: withMetadata(async function serviceStatus(input: DeepPartial<akash_provider_lease_v1_service.ServiceStatusRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).serviceStatus(input, options);
            }, { path: [3, "serviceStatus"], serviceLoader }),
            /**
             * streamServiceStatus
             */
            streamServiceStatus: withMetadata(async function streamServiceStatus(input: DeepPartial<akash_provider_lease_v1_service.ServiceStatusRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).streamServiceStatus(input, options);
            }, { path: [3, "streamServiceStatus"], serviceLoader }),
            /**
             * serviceLogs
             */
            serviceLogs: withMetadata(async function serviceLogs(input: DeepPartial<akash_provider_lease_v1_service.ServiceLogsRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).serviceLogs(input, options);
            }, { path: [3, "serviceLogs"], serviceLoader }),
            /**
             * streamServiceLogs
             */
            streamServiceLogs: withMetadata(async function streamServiceLogs(input: DeepPartial<akash_provider_lease_v1_service.ServiceLogsRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).streamServiceLogs(input, options);
            }, { path: [3, "streamServiceLogs"], serviceLoader }),
            /**
             * attestationQuote requests hardware-signed attestation evidence from the
             * confidential compute sidecar. The provider forwards the tenant's nonce
             * to the sidecar and returns the hardware-signed quote verbatim.
             */
            attestationQuote: withMetadata(async function attestationQuote(input: DeepPartial<akash_provider_lease_v1_service.AttestationQuoteRequest>, options?: CallOptions) {
              const service = await serviceLoader.loadAt(3);
              return getClient(service).attestationQuote(input, options);
            }, { path: [3, "attestationQuote"], serviceLoader })
          }
        },
        v1: {
          /**
           * getStatus defines a method to query provider state
           */
          getStatus: withMetadata(async function getStatus(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).getStatus(input, options);
          }, { path: [4, "getStatus"], serviceLoader }),
          /**
           * Status defines a method to stream provider state
           */
          streamStatus: withMetadata(async function streamStatus(input: DeepPartial<google_protobuf_empty.Empty> = {}, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).streamStatus(input, options);
          }, { path: [4, "streamStatus"], serviceLoader }),
          /**
           * bidScreening screens a deployment group spec for bid eligibility and returns pricing
           */
          bidScreening: withMetadata(async function bidScreening(input: DeepPartial<akash_provider_v1_validation.BidScreeningRequest>, options?: CallOptions) {
            const service = await serviceLoader.loadAt(4);
            return getClient(service).bidScreening(input, options);
          }, { path: [4, "bidScreening"], serviceLoader })
        }
      }
    }
  };
}
