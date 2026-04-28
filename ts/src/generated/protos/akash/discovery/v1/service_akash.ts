import { GetInfoRequest, GetInfoResponse } from "./service.ts";

export const Discovery = {
  typeName: "akash.discovery.v1.Discovery",
  methods: {
    getInfo: {
      name: "GetInfo",
      httpPath: "/akash/discovery/v1/info",
      input: GetInfoRequest,
      output: GetInfoResponse,
      get parent() { return Discovery; },
    },
  },
} as const;
