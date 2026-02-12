import { GroupSpec } from "../generated/protos/akash/deployment/v1beta4/groupspec.ts";
import type { v3DeploymentGroup } from "./types.ts";

function mapEndpoint(ep: { kind?: number; sequence_number?: number }) {
  return {
    kind: ep.kind ?? 0,
    sequenceNumber: ep.sequence_number ?? 0,
  };
}

function formatAmount(amount: unknown): string {
  const s = String(amount ?? "");
  if (!s) return "0.000000000000000000";
  if (s.includes(".")) {
    const [, frac] = s.split(".");
    const pad = 18 - (frac?.length ?? 0);
    return pad > 0 ? s + "0".repeat(pad) : s.slice(0, s.length + (pad ?? 0));
  }
  return `${s}.${"0".repeat(18)}`;
}

type DomainResource = {
  id?: number;
  cpu?: { units?: { val?: unknown }; attributes?: unknown[] };
  memory?: { quantity?: { val?: unknown }; size?: { val?: unknown }; attributes?: unknown[] };
  storage?: Array<{ name?: string; quantity?: { val?: unknown }; size?: { val?: unknown }; attributes?: unknown[] }> | { name?: string; quantity?: { val?: unknown }; size?: { val?: unknown }; attributes?: unknown[] };
  gpu?: { units?: { val?: unknown }; attributes?: unknown[] };
  endpoints?: Array<{ kind?: number; sequence_number?: number }>;
};

function mapResource(domainRes: DomainResource) {
  const mem = domainRes.memory;
  const memVal = mem?.quantity ?? mem?.size;
  const r: Record<string, unknown> = {
    id: domainRes.id ?? 0,
    endpoints: (domainRes.endpoints ?? []).map(mapEndpoint),
  };
  if (domainRes.cpu) r.cpu = domainRes.cpu;
  if (mem) {
    r.memory = memVal ? { quantity: memVal, attributes: mem.attributes ?? [] } : undefined;
  }
  if (domainRes.storage) {
    const arr = Array.isArray(domainRes.storage) ? domainRes.storage : [domainRes.storage];
    r.storage = arr.map((s) => ({
      name: s.name ?? "default",
      quantity: s.quantity ?? s.size,
      attributes: s.attributes ?? [],
    }));
  }
  if (domainRes.gpu) {
    r.gpu = domainRes.gpu;
  } else {
    r.gpu = { units: { val: "0" }, attributes: [] };
  }
  return r;
}

export function domainToGroupSpec(dgroup: v3DeploymentGroup) {
  const partial = {
    name: dgroup.name,
    requirements: dgroup.requirements
      ? {
          signedBy: dgroup.requirements.signedBy
            ? {
                allOf: dgroup.requirements.signedBy.allOf ?? [],
                anyOf: dgroup.requirements.signedBy.anyOf ?? [],
              }
            : { allOf: [], anyOf: [] },
          attributes: dgroup.requirements.attributes ?? [],
        }
      : undefined,
    resources: dgroup.resources.map((ru) => {
      const price = ru.price as { denom?: string; amount?: string | number };
      return {
        resource: mapResource(ru.resource as unknown as DomainResource),
        count: ru.count,
        price: {
          denom: price.denom ?? "uakt",
          amount: formatAmount(price.amount),
        },
      };
    }),
  };
  return GroupSpec.fromPartial(partial);
}

export function groupsToProtobufJson(groups: v3DeploymentGroup[]): unknown[] {
  return groups.map((g) => GroupSpec.toJSON(domainToGroupSpec(g)));
}

function base64ToDecimalStr(b64: string): string {
  const bin = globalThis.atob(b64);
  const arr = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; ++i) arr[i] = bin.charCodeAt(i);
  return new TextDecoder().decode(arr);
}

function isBase64(s: string): boolean {
  return /^[A-Za-z0-9+/]+=*$/.test(s) && s.length % 4 === 0;
}

const ENDPOINT_KIND_STR_TO_NUM: Record<string, number> = {
  SHARED_HTTP: 0,
  RANDOM_PORT: 1,
  LEASED_IP: 2,
};

function normalizeEndpoint(ep: Record<string, unknown>): Record<string, unknown> {
  const out = { ...ep };
  if (out.sequence_number === undefined) out.sequence_number = 0;
  if (typeof out.kind === "string") out.kind = ENDPOINT_KIND_STR_TO_NUM[out.kind] ?? out.kind;
  return out;
}

export function protobufJsonToGoFormat(obj: unknown, parentKey?: string): unknown {
  if (obj === null || obj === undefined) return obj;
  if (Array.isArray(obj)) {
    const arr = obj.map((item, i) => protobufJsonToGoFormat(item, parentKey));
    if (parentKey === "endpoints") {
      return arr.map((item) => normalizeEndpoint((item as Record<string, unknown>) ?? {}));
    }
    return arr;
  }
  if (typeof obj === "object") {
    const out: Record<string, unknown> = {};
    const rec = obj as Record<string, unknown>;
    for (const key of Object.keys(rec)) {
      const val = rec[key];
      if (key === "quantity" && typeof val === "object" && val !== null) {
        out.size = protobufJsonToGoFormat(val, key);
      } else if (key === "val" && typeof val === "string" && isBase64(val)) {
        out[key] = base64ToDecimalStr(val);
      } else if (key === "amount" && typeof rec.denom !== "undefined") {
        out[key] = formatAmount(val);
      } else if (key === "signed_by" && typeof val === "object" && val !== null) {
        const sb = protobufJsonToGoFormat(val, key) as Record<string, unknown>;
        const allOf = sb.all_of as unknown[] | undefined;
        const anyOf = sb.any_of as unknown[] | undefined;
        out[key] = {
          all_of: allOf?.length ? allOf : null,
          any_of: anyOf?.length ? anyOf : null,
        };
      } else if (key === "kind" && typeof val === "string") {
        out[key] = ENDPOINT_KIND_STR_TO_NUM[val] ?? val;
      } else {
        const mapped = protobufJsonToGoFormat(val, key);
        if (mapped !== undefined) out[key] = mapped;
      }
    }
    if (parentKey === "resource" && out.endpoints === undefined) out.endpoints = [];
    if (parentKey === "requirements" && (out.attributes === undefined || (Array.isArray(out.attributes) && out.attributes.length === 0))) out.attributes = null;
    return out;
  }
  return obj;
}
