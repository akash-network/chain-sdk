import type { Attribute } from "../../generated/protos/index.akash.v1.ts";
import {
  Endpoint,
  Endpoint_Kind,
} from "../../generated/protos/index.akash.v1beta4.ts";
import {
  ServiceExposeHTTPOptions,
} from "../../generated/protos/index.provider.akash.v2beta3.ts";
import { convertCpuResourceString, convertResourceString } from "../sizes.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
import type { StorageAttributesValidation } from "../validateSDL/validateSDLInput.ts";

type SDLService = SDLInput["services"][string];
type SDLExpose = NonNullable<SDLService["expose"]>[number];
type SDLExposeTo = NonNullable<SDLExpose["to"]>[number];
type SDLHttpOptions = SDLExpose["http_options"];
type SDLCompute = SDLInput["profiles"]["compute"][string];
type SDLStorage = SDLCompute["resources"]["storage"];
type SDLStorageVolume = SDLStorage extends (infer T)[] ? T : SDLStorage;
type SDLGpuAttributes = NonNullable<NonNullable<SDLCompute["resources"]["gpu"]>["attributes"]>;

export type { SDLCompute, SDLExpose, SDLExposeTo, SDLGpuAttributes, SDLHttpOptions, SDLService, SDLStorage, SDLStorageVolume };

const encoder = new TextEncoder();

export function encodeResourceValue(value: number): Uint8Array {
  return encoder.encode(value.toString());
}

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

export function isIngress(proto: string, global: boolean, externalPort: number, port: number): boolean {
  const effectivePort = externalPort === 0 ? port : externalPort;
  return global && proto === "TCP" && effectivePort === 80;
}

export function transformGpuAttributes(attributes: SDLGpuAttributes): Attribute[] {
  const vendor = attributes.vendor;
  if (!vendor) return [];

  return Object.keys(vendor)
    .sort((a, b) => a.localeCompare(b))
    .flatMap((vendorName) => {
      const models = vendor[vendorName as keyof typeof vendor];
      if (!models) {
        return [{ key: `vendor/${vendorName}/model/*`, value: "true" }];
      }

      return models.map((model) => {
        let key = `vendor/${vendorName}/model/${model.model}`;
        if (model.ram) key += `/ram/${model.ram}`;
        if (model.interface) key += `/interface/${model.interface}`;
        return { key, value: "true" };
      });
    });
}

export function buildHttpOptions(httpOptions?: SDLHttpOptions): ServiceExposeHTTPOptions {
  return ServiceExposeHTTPOptions.fromPartial({
    maxBodySize: httpOptions?.max_body_size ?? 1048576,
    readTimeout: httpOptions?.read_timeout ?? 60000,
    sendTimeout: httpOptions?.send_timeout ?? 60000,
    nextTries: httpOptions?.next_tries ?? 3,
    nextTimeout: httpOptions?.next_timeout ?? 0,
    nextCases: httpOptions?.next_cases ?? ["error", "timeout"],
  });
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

  pairs.sort((a, b) => a.key.localeCompare(b.key));
  return pairs;
}

export function parseServiceProto(proto?: string): string {
  return proto?.toUpperCase() || "TCP";
}

export function buildServiceEndpoints(
  service: SDLService,
  endpointSequenceNumbers: Record<string, number>,
): Endpoint[] {
  return (service.expose ?? []).flatMap((expose) =>
    (expose.to ?? [])
      .filter((to) => to.global)
      .flatMap((to) => {
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

        return [defaultEp, leasedEp];
      }),
  );
}

export function parseCpuUnits(cpu: SDLCompute["resources"]["cpu"]): number {
  return typeof cpu.units === "string"
    ? convertCpuResourceString(cpu.units)
    : cpu.units * 1000;
}

export function parseMemoryBytes(memory: SDLCompute["resources"]["memory"]): number {
  return convertResourceString(memory.size);
}

export function parseStorageBytes(size: string): number {
  return convertResourceString(size);
}

export function parseGpuUnits(gpu?: SDLCompute["resources"]["gpu"]): number {
  const value = gpu?.units;
  if (value === undefined || value === null) return 0;
  return typeof value === "string" ? parseInt(value, 10) : value;
}

export function buildResourceAttributes(attributes?: Record<string, unknown>): Attribute[] | undefined {
  if (!attributes) return undefined;
  return Object.keys(attributes)
    .sort((a, b) => a.localeCompare(b))
    .map((key) => ({ key, value: String(attributes[key]) }));
}
