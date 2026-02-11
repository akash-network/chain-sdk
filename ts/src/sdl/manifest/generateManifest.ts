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

export interface BuildResult {
  groups: Group[];
  groupSpecs: GroupSpec[];
}

export type Manifest = BuildResult["groups"];
export type GenerateManifestResult =
  | { ok?: false; value: ValidationError[] }
  | { ok: true; value: BuildResult };
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
              attributes: buildResourceAttributes(infra.attributes) ?? [],
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

  const storageEntries = Object.entries(service.params.storage ?? {});
  const result = ServiceParams.fromPartial({
    storage: storageEntries.map(([name, config]) =>
      StorageParams.fromPartial({
        name,
        mount: config.mount || "",
        readOnly: config.readOnly || false,
      }),
    ),
  });

  // Permissions are not in the protobuf type but need to be preserved
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
