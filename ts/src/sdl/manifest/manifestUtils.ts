import type { Attribute } from "../../generated/protos/index.akash.v1.ts";
import {
  Endpoint,
  Endpoint_Kind,
} from "../../generated/protos/index.akash.v1beta4.ts";
import {
  ProxyOptions,
  ServiceExposeHTTPOptions,
} from "../../generated/protos/index.provider.akash.v2beta3.ts";
import { parseHTTPTimeout } from "../httpTimeout.ts";
import { convertCpuResourceString, convertResourceString } from "../sizes.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
import type { StorageAttributesValidation } from "../validateSDL/validateSDLInput.ts";

type SDLService = SDLInput["services"][string];
type SDLExpose = NonNullable<SDLService["expose"]>[number];
type SDLExposeTo = NonNullable<SDLExpose["to"]>[number];
type SDLHttpOptions = SDLExpose["http_options"];
type SDLHttpProxyOptions = NonNullable<SDLHttpOptions>["proxy"];
type SDLCompute = SDLInput["profiles"]["compute"][string];
type SDLStorage = SDLCompute["resources"]["storage"];
type SDLStorageVolume = SDLStorage extends (infer T)[] ? T : SDLStorage;
type SDLGpuAttributes = NonNullable<NonNullable<SDLCompute["resources"]["gpu"]>["attributes"]>;

export type { SDLCompute, SDLExpose, SDLExposeTo, SDLGpuAttributes, SDLHttpOptions, SDLService, SDLStorage, SDLStorageVolume };

export function computeEndpointSequenceNumbers(services: SDLInput["services"]): Record<string, number> {
  const endpointNames: string[] = [];

  for (const service of Object.values(services)) {
    if (!service.expose) continue;
    for (const expose of service.expose) {
      if (!expose.to) continue;
      for (const to of expose.to) {
        if (to.global && to.ip && to.ip.length > 0) {
          endpointNames.push(to.ip);
        }
      }
    }
  }

  return endpointNames.sort().reduce<Record<string, number>>((result, name, seqNumber) => {
    result[name] = seqNumber + 1;
    return result;
  }, {});
}

// Go orders every string it sorts in the SDL builder with a plain `<`:
// `Attributes.Less` compares keys, `ServiceExposes.Less` compares service names,
// and placement/service names go through `sort.Strings`. `localeCompare` is not
// the same relation — it folds case and demotes punctuation, so e.g. "a_b" and
// "aB" come out in the opposite order. Attribute keys carry "/" and placement,
// service and volume names are unconstrained strings, so the two can genuinely
// disagree, and the resulting group spec is signed. Every sort on this path uses
// this comparator instead.
export function compareStrings(a: string, b: string): number {
  if (a === b) return 0;
  return a < b ? -1 : 1;
}

export function isIngress(proto: string, global: boolean, externalPort: number, port: number): boolean {
  const effectivePort = externalPort === 0 ? port : externalPort;
  return global && proto === "TCP" && effectivePort === 80;
}

// INTERCONNECT_GROUP_AUTO is the reserved name the SDL parser assigns to
// every `interconnect: []` opt-in within one placement. Tenants cannot
// write it explicitly under `interconnect: { group: ... }` — mirrors
// `InterconnectGroupAuto` in go/sdl/gpu.go.
export const INTERCONNECT_GROUP_AUTO = "auto";

// resolveInterconnectGroup returns the group string the parser derives
// from gpu.attributes.interconnect:
//   - empty sequence `[]`         → INTERCONNECT_GROUP_AUTO ("auto")
//   - mapping `{ group: <name> }` → <name>
//   - anything else / absent       → ""  (non-interconnect)
// Both the on-chain attribute emitter and the off-chain manifest builder
// route through this helper so they agree on the resolved value.
export function resolveInterconnectGroup(interconnect: SDLGpuAttributes["interconnect"]): string {
  if (Array.isArray(interconnect)) {
    return interconnect.length === 0 ? INTERCONNECT_GROUP_AUTO : "";
  }
  if (interconnect && typeof interconnect === "object" && "group" in interconnect) {
    const g = (interconnect as { group?: unknown }).group;
    return typeof g === "string" ? g : "";
  }
  return "";
}

export function transformGpuAttributes(attributes: SDLGpuAttributes): Attribute[] {
  const result: Attribute[] = [];

  const vendor = attributes.vendor;
  if (vendor) {
    Object.keys(vendor)
      .sort(compareStrings)
      .forEach((vendorName) => {
        const models = vendor[vendorName as keyof typeof vendor];
        if (!models) {
          result.push({ key: `vendor/${vendorName}/model/*`, value: "true" });
          return;
        }
        for (const model of models) {
          let key = `vendor/${vendorName}/model/${model.model}`;
          if (model.ram) key += `/ram/${model.ram}`;
          if (model.interface) key += `/interface/${model.interface}`;
          result.push({ key, value: "true" });
        }
      });
  }

  // interconnect emits a single on-chain attribute `interconnect/group` —
  // the group is the entire opt-in signal. Keep parity with the Go parser
  // in go/sdl/gpu.go: empty sequence `[]` resolves to the reserved literal
  // `auto`, the explicit mapping form `{ group: <name> }` carries the
  // tenant-chosen name. See docs/sdl-interconnect-spec.md.
  const group = resolveInterconnectGroup(attributes.interconnect);
  if (group !== "") {
    result.push({ key: "interconnect/group", value: group });
  }

  // Go SDL parser canonicalizes the slice via sort.Sort(res) before
  // returning. Mirror that here so the on-chain attribute order matches
  // byte-for-byte across both implementations.
  result.sort((a, b) => compareStrings(a.key, b.key));

  return result;
}

export function buildHttpOptions(httpOptions?: SDLHttpOptions): ServiceExposeHTTPOptions {
  return ServiceExposeHTTPOptions.fromPartial({
    maxBodySize: httpOptions?.max_body_size ?? 1048576,
    readTimeout: normalizedHTTPTimeout(httpOptions?.read_timeout),
    sendTimeout: normalizedHTTPTimeout(httpOptions?.send_timeout),
    nextTries: httpOptions?.next_tries ?? 3,
    nextTimeout: httpOptions?.next_timeout ?? 0,
    nextCases: httpOptions?.next_cases ?? ["error", "timeout"],
    proxy: buildProxyOptions(httpOptions?.proxy),
  });
}

function buildProxyOptions(proxy?: SDLHttpProxyOptions): ProxyOptions | undefined {
  if (!proxy) return undefined;

  const bufferingDisable = proxy.buffering_disable ?? false;
  const bufferSize = proxy.buffer_size ?? 0;
  const buffersNumber = proxy.buffers_number ?? 0;
  const buffersSize = proxy.buffers_size ?? 0;
  const busyBuffersSize = proxy.busy_buffers_size ?? 0;
  const connectTimeout = proxy.connect_timeout ?? 0;

  if ((buffersNumber === 0) !== (buffersSize === 0)) {
    throw new Error("proxy.buffers_number and proxy.buffers_size must be set together");
  }

  if (!bufferingDisable && bufferSize === 0 && buffersNumber === 0 && buffersSize === 0 && busyBuffersSize === 0 && connectTimeout === 0) {
    return undefined;
  }

  return ProxyOptions.fromPartial({ bufferingDisable, bufferSize, buffersNumber, buffersSize, busyBuffersSize, connectTimeout });
}

function normalizedHTTPTimeout(value: number | string | undefined): number {
  const result = parseHTTPTimeout(value ?? 60_000);
  if (!result.ok) {
    throw new Error(`Invalid HTTP timeout: ${result.message}`);
  }

  return result.milliseconds;
}

export function buildStorageAttributes(attributes?: StorageAttributesValidation): Attribute[] {
  if (!attributes) return [];

  const pairs: Attribute[] = Object.entries(attributes).map(([key, value]) => ({
    key,
    value: String(value),
  }));

  if (attributes.class === "ram" && !("persistent" in attributes)) {
    pairs.push({ key: "persistent", value: "false" });
  }

  pairs.sort((a, b) => compareStrings(a.key, b.key));
  return pairs;
}

export function parseServiceProto(proto?: string): string {
  return proto?.toUpperCase() || "TCP";
}

export interface ExposeSortKey {
  service: string;
  port: number;
  proto: string;
  global: boolean;
}

// Mirrors `ServiceExposes.Less` in go/manifest/v2beta3/serviceexposes.go.
export function compareExposes(a: ExposeSortKey, b: ExposeSortKey): number {
  if (a.service !== b.service) return compareStrings(a.service, b.service);
  if (a.port !== b.port) return a.port - b.port;
  if (a.proto !== b.proto) return compareStrings(a.proto, b.proto);
  if (a.global !== b.global) return a.global ? -1 : 1;
  return 0;
}

// Go expands a service's `expose` entries into one manifest expose per `to`
// target and sorts the whole list before anything reads it — `toManifestExpose`
// in go/sdl/expose.go. Both the manifest expose list and the resource endpoint
// list are then built by walking it in that order, and the endpoint order is
// part of the signed group spec, so ordering has to be decided once, here.
export function sortedExposeTargets(service: SDLService): { expose: SDLExpose; to: SDLExposeTo }[] {
  return (service.expose ?? [])
    .flatMap((expose) => (expose.to ?? []).map((to) => ({ expose, to })))
    .sort((a, b) => compareExposes(exposeSortKey(a), exposeSortKey(b)));
}

function exposeSortKey({ expose, to }: { expose: SDLExpose; to: SDLExposeTo }): ExposeSortKey {
  return {
    service: to.service || "",
    port: expose.port,
    proto: parseServiceProto(expose.proto),
    global: to.global || false,
  };
}

export function buildServiceEndpoints(
  service: SDLService,
  endpointSequenceNumbers: Record<string, number>,
): Endpoint[] {
  return sortedExposeTargets(service)
    .filter(({ to }) => to.global)
    .flatMap(({ expose, to }) => {
      const externalPort = expose.as || 0;
      const proto = parseServiceProto(expose.proto);
      const kind = isIngress(proto, !!to.global, externalPort, expose.port)
        ? Endpoint_Kind.SHARED_HTTP
        : Endpoint_Kind.RANDOM_PORT;

      const defaultEp = Endpoint.fromPartial({
        kind,
        sequenceNumber: 0,
      });

      if (!to.ip?.length) {
        return [defaultEp];
      }

      const leasedEp = Endpoint.fromPartial({
        kind: Endpoint_Kind.LEASED_IP,
        sequenceNumber: endpointSequenceNumbers[to.ip] ?? 0,
      });

      // Go emits [LEASED_IP, <kind>] then sorts this pair on its own
      // (`ServiceExpose.GetEndpoints`), which puts the lower kind first.
      return [defaultEp, leasedEp];
    });
}

export function parseCpuUnits(cpu: SDLCompute["resources"]["cpu"]): bigint {
  return typeof cpu.units === "string"
    ? convertCpuResourceString(cpu.units)
    : convertCpuResourceString(String(cpu.units));
}

export function parseMemoryBytes(memory: SDLCompute["resources"]["memory"]): bigint {
  return convertResourceString(memory.size);
}

export function parseStorageBytes(size: string): bigint {
  return convertResourceString(size);
}

export function parseGpuUnits(gpu?: SDLCompute["resources"]["gpu"]): bigint {
  const value = gpu?.units;
  if (value === undefined || value === null) return 0n;
  return BigInt(value);
}

export function buildResourceAttributes(attributes?: Record<string, unknown>): Attribute[] | undefined {
  if (!attributes) return undefined;
  return Object.keys(attributes)
    .sort(compareStrings)
    .map((key) => ({ key, value: String(attributes[key]) }));
}
