import { Empty } from "../../../google/protobuf/empty.ts";
import { Status } from "./status.ts";
import { BidScreeningRequest, BidScreeningResponse } from "./validation.ts";

export const ProviderRPC = {
  typeName: "akash.provider.v1.ProviderRPC",
  methods: {
    getStatus: {
      name: "GetStatus",
      httpPath: "/v1/status",
      input: Empty,
      output: Status,
      get parent() { return ProviderRPC; },
    },
    streamStatus: {
      name: "StreamStatus",
      kind: "server_streaming",
      input: Empty,
      output: Status,
      get parent() { return ProviderRPC; },
    },
    bidScreening: {
      name: "BidScreening",
      input: BidScreeningRequest,
      output: BidScreeningResponse,
      get parent() { return ProviderRPC; },
    },
  },
} as const;
