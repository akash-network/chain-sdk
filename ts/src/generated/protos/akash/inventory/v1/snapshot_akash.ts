import { GetInventorySnapshotRequest, GetInventorySnapshotResponse } from "./snapshot.ts";

export const InventoryService = {
  typeName: "akash.inventory.v1.InventoryService",
  methods: {
    getInventorySnapshot: {
      name: "GetInventorySnapshot",
      httpMethod: "post",
      httpPath: "/v1/inventory/snapshot",
      input: GetInventorySnapshotRequest,
      output: GetInventorySnapshotResponse,
      get parent() { return InventoryService; },
    },
  },
} as const;
