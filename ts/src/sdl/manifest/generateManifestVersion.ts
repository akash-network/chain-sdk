import { default as stableStringify } from "json-stable-stringify";

import type { Resources, ResourceValue } from "../../generated/protos/index.akash.v1beta4.ts";
import type { Group, Service, ServiceExpose } from "../../generated/protos/index.provider.akash.v2beta3.ts";

const decoder = new TextDecoder();
const encoder = new TextEncoder();

export async function generateManifestVersion(manifest: Group[]): Promise<Uint8Array> {
  const jsonStr = manifestToSortedJSON(manifest);
  const sortedBytes = encoder.encode(jsonStr);
  const sum = await crypto.subtle.digest("SHA-256", sortedBytes);
  return new Uint8Array(sum);
}

export function manifestToSortedJSON(manifest: Group[]): string {
  const jsonReady = manifest.map((group) => ({
    name: group.name,
    services: group.services.map(serviceToJSON),
  }));

  const sorted = stableStringify(jsonReady) || "";
  return escapeHtml(sorted);
}

function serviceToJSON(svc: Service): unknown {
  return {
    name: svc.name,
    image: svc.image,
    command: svc.command.length > 0 ? svc.command : null,
    args: svc.args.length > 0 ? svc.args : null,
    env: svc.env.length > 0 ? svc.env : null,
    resources: resourcesToJSON(svc.resources!),
    count: svc.count,
    expose: svc.expose.map(exposeToJSON),
    ...(svc.params
      ? {
          params: {
            storage: svc.params.storage.map((s) => ({
              name: s.name,
              mount: s.mount,
              readOnly: s.readOnly,
            })),
            ...(Object.hasOwn(svc.params, "permissions")
              ? { permissions: (svc.params as unknown as Record<string, unknown>).permissions }
              : {}),
          },
        }
      : {}),
    credentials: svc.credentials
      ? {
          host: svc.credentials.host,
          email: svc.credentials.email,
          username: svc.credentials.username,
          password: svc.credentials.password,
        }
      : null,
  };
}

function exposeToJSON(e: ServiceExpose): unknown {
  return {
    port: e.port,
    externalPort: e.externalPort,
    proto: e.proto,
    service: e.service,
    global: e.global,
    hosts: e.hosts.length > 0 ? e.hosts : null,
    httpOptions: e.httpOptions
      ? {
          maxBodySize: e.httpOptions.maxBodySize,
          readTimeout: e.httpOptions.readTimeout,
          sendTimeout: e.httpOptions.sendTimeout,
          nextTries: e.httpOptions.nextTries,
          nextTimeout: e.httpOptions.nextTimeout,
          nextCases: e.httpOptions.nextCases,
        }
      : undefined,
    ip: e.ip,
    endpointSequenceNumber: e.endpointSequenceNumber,
  };
}

function resourceValueToString(rv: ResourceValue | undefined): string {
  if (!rv || rv.val.length === 0) return "0";
  return decoder.decode(rv.val);
}

function resourcesToJSON(resources: Resources): unknown {
  return {
    id: resources.id,
    cpu: {
      units: { val: resourceValueToString(resources.cpu?.units) },
      ...(resources.cpu?.attributes && resources.cpu.attributes.length > 0
        ? { attributes: resources.cpu.attributes }
        : {}),
    },
    memory: {
      size: { val: resourceValueToString(resources.memory?.quantity) },
      ...(resources.memory?.attributes && resources.memory.attributes.length > 0
        ? { attributes: resources.memory.attributes }
        : {}),
    },
    storage: resources.storage.map((s) => ({
      name: s.name,
      size: { val: resourceValueToString(s.quantity) },
      ...(s.attributes && s.attributes.length > 0 ? { attributes: s.attributes } : {}),
    })),
    gpu: {
      units: { val: resourceValueToString(resources.gpu?.units) },
      ...(resources.gpu?.attributes && resources.gpu.attributes.length > 0
        ? { attributes: resources.gpu.attributes }
        : {}),
    },
    endpoints: resources.endpoints.map((e) => ({
      ...(e.kind !== 0 ? { kind: e.kind } : {}),
      sequence_number: e.sequenceNumber,
    })),
  };
}

function escapeHtml(raw: string): string {
  return raw
    .replace(/</g, "\\u003c")
    .replace(/>/g, "\\u003e")
    .replace(/&/g, "\\u0026");
}
