import { MsgCreateCertificate, MsgCreateCertificateResponse, MsgRevokeCertificate, MsgRevokeCertificateResponse } from "./msg.ts";

export const Msg = {
  typeName: "akash.cert.v1.Msg",
  methods: {
    createCertificate: {
      name: "CreateCertificate",
      input: MsgCreateCertificate,
      output: MsgCreateCertificateResponse,
      get parent() { return Msg; },
    },
    revokeCertificate: {
      name: "RevokeCertificate",
      input: MsgRevokeCertificate,
      output: MsgRevokeCertificateResponse,
      get parent() { return Msg; },
    },
  },
} as const;
