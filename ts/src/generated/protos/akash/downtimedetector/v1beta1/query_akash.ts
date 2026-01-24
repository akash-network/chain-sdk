import { RecoveredSinceDowntimeOfLengthRequest, RecoveredSinceDowntimeOfLengthResponse } from "./query.ts";

export const Query = {
  typeName: "akash.downtimedetector.v1beta1.Query",
  methods: {
    recoveredSinceDowntimeOfLength: {
      name: "RecoveredSinceDowntimeOfLength",
      httpPath: "/akash/downtime-detector/v1beta1/RecoveredSinceDowntimeOfLength",
      input: RecoveredSinceDowntimeOfLengthRequest,
      output: RecoveredSinceDowntimeOfLengthResponse,
      get parent() { return Query; },
    },
  },
} as const;
