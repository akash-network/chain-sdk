import { PlacementRequirements, SignedBy } from "../../generated/protos/akash/base/attributes/v1/attribute.ts";
import { CPU } from "../../generated/protos/akash/base/resources/v1beta4/cpu.ts";
import { GPU } from "../../generated/protos/akash/base/resources/v1beta4/gpu.ts";
import { Memory } from "../../generated/protos/akash/base/resources/v1beta4/memory.ts";
import { Resources } from "../../generated/protos/akash/base/resources/v1beta4/resources.ts";
import { Storage } from "../../generated/protos/akash/base/resources/v1beta4/storage.ts";
import { GroupSpec } from "../../generated/protos/akash/deployment/v1beta4/groupspec.ts";
import { ResourceUnit } from "../../generated/protos/akash/deployment/v1beta4/resourceunit.ts";
import { Group } from "../../generated/protos/akash/manifest/v2beta3/group.ts";
import { ImageCredentials, Service, ServiceParams, StorageParams } from "../../generated/protos/akash/manifest/v2beta3/service.ts";
import { ServiceExpose } from "../../generated/protos/akash/manifest/v2beta3/serviceexpose.ts";
import { MAINNET_ID } from "../../network/config.ts";
import type { NetworkId } from "../../network/types.ts";
import type { ValidationError } from "../../utils/jsonSchemaValidation.ts";
import { castArray } from "../utils.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
import { validateSDL } from "../validateSDL/validateSDL.ts";
import {
  buildHttpOptions,
  buildResourceAttributes,
  buildServiceEndpoints,
  buildStorageAttributes,
  computeEndpointSequenceNumbers,
  encodeResourceValue,
  parseCpuUnits,
  parseGpuUnits,
  parseMemoryBytes,
  parseServiceProto,
  parseStorageBytes,
  type SDLCompute,
  type SDLService,
  transformGpuAttributes,
} from "./manifestUtils.ts";

export interface GenerateManifestOkResult {
  groups: Group[];
  groupSpecs: GroupSpec[];
}

export type Manifest = GenerateManifestOkResult["groups"];
export type GenerateManifestResult =
  | { ok?: false; value: ValidationError[] }
  | { ok: true; value: GenerateManifestOkResult };

export function generateManifest(sdl: SDLInput, networkId: NetworkId = MAINNET_ID): GenerateManifestResult {
  const errors = validateSDL(sdl, networkId);
  if (errors) return { ok: false, value: errors };

  const endpointSequenceNumbers = computeEndpointSequenceNumbers(sdl.services);
  const groupsMap = new Map<string, {
    dgroup: GroupSpec;
    boundComputes: Record<string, number>;
  }>();
  const resourceIds = new Map<string, number>();

  const deploymentsByPlacement = new Map<string, [string, { profile: string; count: number }][]>();
  for (const [svcName, placements] of Object.entries(sdl.deployment)) {
    for (const [placementName, deployment] of Object.entries(placements)) {
      let list = deploymentsByPlacement.get(placementName);
      if (!list) {
        list = [];
        deploymentsByPlacement.set(placementName, list);
      }
      list.push([svcName, deployment]);
    }
  }
  for (const list of deploymentsByPlacement.values()) {
    list.sort(([a], [b]) => a.localeCompare(b));
  }

  const services = Object.entries(sdl.services).sort(([a], [b]) => a.localeCompare(b));

  for (const [svcName, service] of services) {
    for (const [placementName, svcdepl] of Object.entries(sdl.deployment[svcName])) {
      const compute = sdl.profiles.compute[svcdepl.profile];
      const infra = sdl.profiles.placement[placementName];
      const pricing = infra.pricing[svcdepl.profile];
      const price = {
        denom: pricing.denom,
        amount: pricing.amount?.toString(),
      };

      let group = groupsMap.get(placementName);

      if (!group) {
        group = {
          dgroup: GroupSpec.fromPartial({
            name: placementName,
            resources: [],
            requirements: PlacementRequirements.fromPartial({
              attributes: buildResourceAttributes(infra.attributes),
              signedBy: SignedBy.fromPartial({
                allOf: infra.signedBy?.allOf,
                anyOf: infra.signedBy?.anyOf,
              }),
            }),
          }),
          boundComputes: {},
        };

        groupsMap.set(placementName, group);
      }

      const profileKey = `${placementName}:${svcdepl.profile}`;
      const location = group.boundComputes[svcdepl.profile];

      if (location === undefined) {
        const resId = group.dgroup.resources.length > 0
          ? group.dgroup.resources.length + 1
          : 1;

        resourceIds.set(profileKey, resId);

        const resources = buildResources(resId, compute, service, endpointSequenceNumbers);

        group.dgroup.resources.push(
          ResourceUnit.fromPartial({
            resource: resources,
            count: svcdepl.count,
            price,
          }),
        );

        group.boundComputes[svcdepl.profile] = group.dgroup.resources.length - 1;
      } else {
        if (!resourceIds.has(profileKey)) {
          resourceIds.set(profileKey, group.dgroup.resources[location].resource!.id);
        }

        group.dgroup.resources[location].count += svcdepl.count;
        group.dgroup.resources[location].resource!.endpoints.push(
          ...buildServiceEndpoints(service, endpointSequenceNumbers),
        );
      }
    }
  }

  for (const group of groupsMap.values()) {
    for (const resourceUnit of group.dgroup.resources) {
      resourceUnit.resource!.endpoints.sort(
        (a, b) => a.kind - b.kind || a.sequenceNumber - b.sequenceNumber,
      );
    }
  }

  const sortedGroupNames = [...groupsMap.keys()].sort();
  let groups: Group[] | undefined;
  let groupSpecs: GroupSpec[] | undefined;

  const manifest = {
    get groups() {
      groups ??= sortedGroupNames.map((placementName) => {
        const deployments = deploymentsByPlacement.get(placementName)!;

        return Group.fromPartial({
          name: placementName,
          services: deployments.map(([svcName]) => {
            const service = sdl.services[svcName];
            const deployment = sdl.deployment[svcName][placementName];
            const compute = sdl.profiles.compute[deployment.profile];
            const resourceId = resourceIds.get(`${placementName}:${deployment.profile}`) || 1;

            return buildManifestService(
              resourceId,
              svcName,
              service,
              compute,
              deployment.count,
              endpointSequenceNumbers,
            );
          }),
        });
      });
      return groups;
    },
    get groupSpecs() {
      groupSpecs ??= sortedGroupNames.map((name) => groupsMap.get(name)!.dgroup);
      return groupSpecs;
    },
  };

  return { ok: true, value: manifest };
}

function buildResources(
  id: number,
  compute: SDLCompute,
  service: SDLService,
  endpointSequenceNumbers: Record<string, number>,
): Resources {
  const res = compute.resources;
  const cpuAttributes = buildResourceAttributes(res.cpu.attributes);
  const gpuAttributes = res.gpu?.attributes ? transformGpuAttributes(res.gpu.attributes) : [];

  return Resources.fromPartial({
    id,
    cpu: CPU.fromPartial({
      units: { val: encodeResourceValue(parseCpuUnits(res.cpu)) },
      attributes: cpuAttributes,
    }),
    memory: Memory.fromPartial({
      quantity: { val: encodeResourceValue(parseMemoryBytes(res.memory)) },
    }),
    storage: castArray(res.storage).map((s) =>
      Storage.fromPartial({
        name: s.name || "default",
        quantity: { val: encodeResourceValue(parseStorageBytes(s.size)) },
        attributes: buildStorageAttributes(s.attributes),
      }),
    ),
    gpu: GPU.fromPartial({
      units: { val: encodeResourceValue(parseGpuUnits(res.gpu)) },
      attributes: gpuAttributes,
    }),
    endpoints: buildServiceEndpoints(service, endpointSequenceNumbers),
  });
}

function buildManifestService(
  resourceId: number,
  name: string,
  service: SDLService,
  compute: SDLCompute,
  count: number,
  endpointSequenceNumbers: Record<string, number>,
): Service {
  const credentials = service.credentials
    ? ImageCredentials.fromPartial({
        host: service.credentials.host,
        email: service.credentials.email || "",
        username: service.credentials.username,
        password: service.credentials.password,
      })
    : undefined;

  const params = buildParams(service);

  return Service.fromPartial({
    name,
    image: service.image,
    command: service.command || [],
    args: service.args || [],
    env: service.env || [],
    resources: buildResources(resourceId, compute, service, endpointSequenceNumbers),
    count,
    expose: buildManifestExpose(service, endpointSequenceNumbers),
    params,
    credentials,
  });
}

function buildParams(service: SDLService): ServiceParams | undefined {
  if (!service.params) return undefined;

  const storage = service.params.storage || {};
  const storageNames = service.params.storage ? Object.keys(storage).sort() : [];
  const result = ServiceParams.fromPartial({
    storage: storageNames.map((name) => {
      const config = storage[name];
      return StorageParams.fromPartial({
        name,
        mount: config.mount || "",
        readOnly: config.readOnly || false,
      });
    }),
  });

  if (service.params.permissions) {
    (result as unknown as Record<string, unknown>).permissions = service.params.permissions;
  }

  return result;
}

function buildManifestExpose(
  service: SDLService,
  endpointSequenceNumbers: Record<string, number>,
): ServiceExpose[] {
  return (service.expose ?? [])
    .flatMap((expose) =>
      (expose.to ?? []).map((to) =>
        ServiceExpose.fromPartial({
          port: expose.port,
          externalPort: expose.as,
          proto: parseServiceProto(expose.proto),
          service: to.service || "",
          global: to.global || false,
          hosts: expose.accept || [],
          httpOptions: buildHttpOptions(expose.http_options),
          ip: to.ip || "",
          endpointSequenceNumber: endpointSequenceNumbers[to.ip!],
        }),
      ),
    )
    .sort((a, b) => {
      if (a.service !== b.service) return a.service.localeCompare(b.service);
      if (a.port !== b.port) return a.port - b.port;
      if (a.proto !== b.proto) return a.proto.localeCompare(b.proto);
      if (a.global !== b.global) return a.global ? -1 : 1;
      return 0;
    });
}

const SNAKE_TO_CAMEL: Record<string, string> = {
  endpoint_sequence_number: "endpointSequenceNumber",
  external_port: "externalPort",
  http_options: "httpOptions",
  max_body_size: "maxBodySize",
  next_cases: "nextCases",
  next_timeout: "nextTimeout",
  next_tries: "nextTries",
  read_timeout: "readTimeout",
  send_timeout: "sendTimeout",
};

export function manifestToFixtureFormat(groups: Group[]): unknown[] {
  return groups.map((g) => Group.toJSON(g)).map((groupJson) => {
    const obj = JSON.parse(JSON.stringify(groupJson)) as Record<string, unknown>;
    return convertToFixtureFormat(obj);
  });
}

const SERVICE_DEFAULTS: Record<string, unknown> = {
  args: null,
  command: null,
  credentials: null,
  env: null,
};

const EXPOSE_DEFAULTS: Record<string, unknown> = {
  endpointSequenceNumber: 0,
  externalPort: 0,
  service: "",
  ip: "",
  hosts: null,
  global: false,
};

const HTTP_OPTIONS_DEFAULTS: Record<string, unknown> = {
  nextTimeout: 0,
};

const STORAGE_PARAM_DEFAULTS: Record<string, unknown> = {
  readOnly: false,
};

function convertToFixtureFormat(obj: unknown, parentKey?: string): unknown {
  if (obj === null || obj === undefined) return obj;
  if (Array.isArray(obj)) {
    const converted = obj.map((item) => convertToFixtureFormat(item, parentKey));
    if (parentKey === "endpoints") {
      return converted.map((ep) => {
        const e = ep as Record<string, unknown>;
        if (Object.keys(e).length === 0) return { sequence_number: 0 };
        if (e.sequence_number === undefined) return { ...e, sequence_number: 0 };
        return e;
      });
    }
    return converted;
  }
  if (typeof obj === "object") {
    const rec = obj as Record<string, unknown>;
    const out: Record<string, unknown> = {};
    for (const key of Object.keys(rec)) {
      const camelKey = SNAKE_TO_CAMEL[key] ?? key;
      if (key === "quantity" && rec.quantity && typeof rec.quantity === "object") {
        const q = rec.quantity as Record<string, unknown>;
        out.size = q.val !== undefined ? { val: decodeVal(q.val) } : convertToFixtureFormat(rec.quantity);
      } else if (key === "val") {
        out[camelKey] = decodeVal(rec[key]);
      } else if (key === "kind" && typeof rec.kind === "string") {
        out[camelKey] = endpointKindToNumber(rec.kind);
      } else if (key !== "quantity") {
        out[camelKey] = convertToFixtureFormat(rec[key], key);
      }
    }
    if (parentKey === "services") {
      for (const [k, v] of Object.entries(SERVICE_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
    }
    if (parentKey === "resources" && !("endpoints" in out)) {
      out.endpoints = [];
    }
    if (parentKey === "expose") {
      for (const [k, v] of Object.entries(EXPOSE_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
      if (Array.isArray(out.hosts) && out.hosts.length === 0) out.hosts = null;
      if (out.httpOptions && typeof out.httpOptions === "object") {
        const ho = out.httpOptions as Record<string, unknown>;
        for (const [k, v] of Object.entries(HTTP_OPTIONS_DEFAULTS)) {
          if (!(k in ho)) ho[k] = v;
        }
      }
    }
    if (parentKey === "storage" && "name" in out && "mount" in out) {
      for (const [k, v] of Object.entries(STORAGE_PARAM_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
    }
    return out;
  }
  return obj;
}

function decodeVal(v: unknown): string {
  if (v === null || v === undefined) return "";
  if (typeof v === "string") {
    try {
      const bytes = Uint8Array.from(atob(v), (c) => c.charCodeAt(0));
      return new TextDecoder().decode(bytes);
    } catch {
      return v;
    }
  }
  return String(v);
}

function endpointKindToNumber(kind: string): number {
  const map: Record<string, number> = {
    SHARED_HTTP: 0,
    RANDOM_PORT: 1,
    LEASED_IP: 2,
  };
  return map[kind] ?? 0;
}
