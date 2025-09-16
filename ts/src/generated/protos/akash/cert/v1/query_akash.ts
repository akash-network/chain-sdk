import { QueryCertificatesRequest, QueryCertificatesResponse } from "./query.ts";

export const Query = {
  typeName: "akash.cert.v1.Query",
  methods: {
    certificates: {
      name: "Certificates",
      httpPath: "/akash/cert/v1/certificates/list",
      input: QueryCertificatesRequest,
      output: QueryCertificatesResponse,
      get parent() { return Query; },
    },
  },
} as const;
