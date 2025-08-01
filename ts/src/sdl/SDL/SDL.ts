/* eslint-disable @typescript-eslint/no-explicit-any */
import YAML from "js-yaml";
import { default as stableStringify } from "json-stable-stringify";

import { AKT_DENOM, MAINNET_ID, USDC_IBC_DENOMS } from "../../network/config.ts";
import type { NetworkId } from "../../network/types.ts";
import { convertCpuResourceString, convertResourceString } from "../sizes.ts";
import type {
  v2ComputeResources,
  v2Expose,
  v2ExposeTo,
  v2HTTPOptions,
  v2Manifest,
  v2ManifestService,
  v2ManifestServiceParams,
  v2ProfileCompute,
  v2ResourceCPU,
  v2ResourceMemory,
  v2ResourceStorage,
  v2ResourceStorageArray,
  v2Sdl,
  v2Service,
  v2ServiceExpose,
  v2ServiceExposeHttpOptions,
  v2ServiceImageCredentials,
  v2ServiceParams,
  v2StorageAttributes,
  v3ComputeResources,
  v3DeploymentGroup,
  v3GPUAttributes,
  v3Manifest,
  v3ManifestService,
  v3ManifestServiceParams,
  v3ProfileCompute,
  v3ResourceGPU,
  v3Sdl,
  v3ServiceExpose,
  v3ServiceExposeHttpOptions } from "../types.ts";
import { SdlValidationError } from "./SdlValidationError.ts";
import { castArray } from "./utils.ts";

const Endpoint_SHARED_HTTP = 0;
const Endpoint_RANDOM_PORT = 1;
const Endpoint_LEASED_IP = 2;
export const GPU_SUPPORTED_VENDORS = ["nvidia", "amd"];
export const GPU_SUPPORTED_INTERFACES = ["pcie", "sxm"];

function isArray<T>(obj: any): obj is Array<T> {
  return Array.isArray(obj);
}

function isString(str: any): str is string {
  return typeof str === "string";
}

type NetworkVersion = "beta2" | "beta3";

/**
 * SDL (Stack Definition Language) parser and validator
 * Handles parsing and validation of Akash deployment manifests
 *
 * @example
 * ```ts
 * import { SDL } from './SDL';
 *
 * const yaml = `
 * version: "2.0"
 * services:
 *   web:
 *     image: nginx
 *     expose:
 *       - port: 80
 *         as: 80
 *         to:
 *           - global: true
 * `;
 *
 * // Parse SDL from YAML string
 * const sdl = SDL.fromString(yaml);
 *
 * // Get deployment manifest
 * const manifest = sdl.manifest();
 *
 * // Get deployment groups
 * const groups = sdl.groups();
 * ```
 */
export class SDL {
  /**
   * Creates an SDL instance from a YAML string.
   *
   * @param {string} yaml - The YAML string containing the SDL definition.
   * @param {NetworkVersion} [version="beta2"] - The SDL version (beta2 or beta3).
   * @param {NetworkId} [networkId=MAINNET_ID] - The network ID to validate against.
   * @returns {SDL} An instance of the SDL class.
   *
   * @example
   * ```ts
   * const yaml = `
   * version: "2.0"
   * services:
   *   web:
   *     image: nginx
   *     expose:
   *       - port: 80
   *         as: 80
   *         to:
   *           - global: true
   * `;
   * const sdl = SDL.fromString(yaml);
   * ```
   */
  static fromString(yaml: string, version: NetworkVersion = "beta2", networkId: NetworkId = MAINNET_ID) {
    const data = YAML.load(yaml) as v3Sdl;
    return new SDL(data, version, networkId);
  }

  /**
   * Validates SDL YAML string (deprecated)
   * @deprecated Use SDL.constructor directly
   */
  static validate(yaml: string) {
    console.warn("SDL.validate is deprecated. Use SDL.constructor directly.");
    const data = YAML.load(yaml) as v3Sdl;

    for (const [name, profile] of Object.entries(data.profiles.compute || {})) {
      this.validateGPU(name, profile.resources.gpu);
      this.validateStorage(name, profile.resources.storage);
    }

    return data;
  }

  /**
   * Validates the GPU configuration for a given service profile.
   *
   * @param {string} name - The name of the service profile.
   * @param {v3ResourceGPU | undefined} gpu - The GPU resource configuration.
   * @throws Will throw an error if the GPU configuration is invalid.
   *
   * @example
   * ```ts
   * const gpuConfig = { units: "1", attributes: { vendor: { nvidia: [{ model: "RTX 3080", ram: "10GB" }] } } };
   * SDL.validateGPU("web", gpuConfig);
   * ```
   */
  static validateGPU(name: string, gpu: v3ResourceGPU | undefined) {
    if (gpu) {
      if (typeof gpu.units === "undefined") {
        throw new Error("GPU units must be specified for profile " + name);
      }

      const units = parseInt(gpu.units.toString());

      if (units === 0 && gpu.attributes !== undefined) {
        throw new Error("GPU must not have attributes if units is 0");
      }

      if (units > 0 && gpu.attributes === undefined) {
        throw new Error("GPU must have attributes if units is not 0");
      }

      if (units > 0 && gpu.attributes?.vendor === undefined) {
        throw new Error("GPU must specify a vendor if units is not 0");
      }

      if (units > 0 && !GPU_SUPPORTED_VENDORS.some((vendor) => vendor in (gpu.attributes?.vendor || {}))) {
        throw new Error(`GPU must be one of the supported vendors (${GPU_SUPPORTED_VENDORS.join(",")}).`);
      }

      const vendor: string = Object.keys(gpu.attributes?.vendor || {})[0];

      if (units > 0 && !!gpu.attributes?.vendor[vendor] && !Array.isArray(gpu.attributes.vendor[vendor])) {
        throw new Error(`GPU configuration must be an array of GPU models with optional ram.`);
      }

      if (
        units > 0
        && Object.values(gpu.attributes?.vendor || {}).some((models) =>
          models?.some((model) => model.interface && !GPU_SUPPORTED_INTERFACES.includes(model.interface)),
        )
      ) {
        throw new Error(`GPU interface must be one of the supported interfaces (${GPU_SUPPORTED_INTERFACES.join(",")}).`);
      }
    }
  }

  /**
   * Validates the storage configuration for a given service.
   *
   * @param {string} name - The name of the service.
   * @param {v2ResourceStorage | v2ResourceStorageArray} [storage] - The storage resource configuration.
   * @throws Will throw an error if the storage configuration is invalid.
   *
   * @example
   * ```ts
   * const storageConfig = { size: "10Gi", attributes: { class: "ssd" } };
   * SDL.validateStorage("web", storageConfig);
   * ```
   */
  static validateStorage(name: string, storage?: v2ResourceStorage | v2ResourceStorageArray) {
    if (!storage) {
      throw new Error("Storage is required for service " + name);
    }

    const storages = isArray(storage) ? storage : [storage];

    for (const storage of storages) {
      if (typeof storage.size === "undefined") {
        throw new Error("Storage size is required for service " + name);
      }

      if (storage.attributes) {
        for (const [key, value] of Object.entries(storage.attributes)) {
          if (key === "class" && value === "ram" && storage.attributes.persistent === true) {
            throw new Error("Storage attribute 'ram' must have 'persistent' set to 'false' or not defined for service " + name);
          }
        }
      }
    }
  }

  private readonly ENDPOINT_NAME_VALIDATION_REGEX = /^[a-z]+[-_\da-z]+$/;
  private readonly ABSOLUTE_PATH_REGEX = /^(\/|([a-zA-Z]:)?([\\/]))/;

  private readonly ENDPOINT_KIND_IP = "ip";

  private readonly endpointsUsed = new Set<string>();

  private readonly portsUsed = new Map<string, string>();

  constructor(
    public readonly data: v2Sdl,
    public readonly version: NetworkVersion = "beta2",
    private readonly networkId: NetworkId = MAINNET_ID,
  ) {
    this.validate();
  }

  private validate() {
    // TODO: this should really be cast to unknown, then assigned
    // to v2 or v3 SDL only after being validated
    this.validateEndpoints();

    Object.keys(this.data.services).forEach((serviceName) => {
      this.validateDeploymentWithRelations(serviceName);
      this.validateLeaseIP(serviceName);
      this.validateCredentials(serviceName);
    });

    this.validateDenom();
    this.validateEndpointsUtility();
  }

  private validateDenom() {
    const usdcDenom = USDC_IBC_DENOMS[this.networkId];
    const denoms = this.groups()
      .flatMap((g) => g.resources)
      .map((resource) => resource.price.denom);
    const invalidDenom = denoms.find((denom) => denom !== AKT_DENOM && denom !== usdcDenom);

    SdlValidationError.assert(!invalidDenom, `Invalid denom: "${invalidDenom}". Only uakt and ${usdcDenom} are supported.`);
  }

  private validateEndpoints() {
    if (!this.data.endpoints) {
      return;
    }
    Object.keys(this.data.endpoints).forEach((endpointName) => {
      const endpoint = this.data.endpoints[endpointName] || {};
      SdlValidationError.assert(this.ENDPOINT_NAME_VALIDATION_REGEX.test(endpointName), `Endpoint named "${endpointName}" is not a valid name.`);
      SdlValidationError.assert(!!endpoint.kind, `Endpoint named "${endpointName}" has no kind.`);
      SdlValidationError.assert(endpoint.kind === this.ENDPOINT_KIND_IP, `Endpoint named "${endpointName}" has an unknown kind "${endpoint.kind}".`);
    });
  }

  private validateCredentials(serviceName: string) {
    const { credentials } = this.data.services[serviceName];

    if (credentials) {
      const credentialsKeys: (keyof v2ServiceImageCredentials)[] = ["host", "username", "password"];
      credentialsKeys.forEach((key) => {
        SdlValidationError.assert(credentials[key]?.trim().length, `service "${serviceName}" credentials missing "${key}"`);
      });
    }
  }

  private validateDeploymentWithRelations(serviceName: string) {
    const deployment = this.data.deployment[serviceName];
    SdlValidationError.assert(deployment, `Service "${serviceName}" is not defined in the "deployment" section.`);

    Object.keys(this.data.deployment[serviceName]).forEach((deploymentName) => {
      this.validateDeploymentRelations(serviceName, deploymentName);
      this.validateServiceStorages(serviceName, deploymentName);
      this.validateStorages(serviceName, deploymentName);
      this.validateGPU(serviceName, deploymentName);
    });
  }

  private validateDeploymentRelations(serviceName: string, deploymentName: string) {
    const serviceDeployment = this.data.deployment[serviceName][deploymentName];
    const compute = this.data.profiles.compute?.[serviceDeployment.profile];
    const infra = this.data.profiles.placement?.[deploymentName];

    SdlValidationError.assert(infra, `The placement "${deploymentName}" is not defined in the "placement" section.`);
    SdlValidationError.assert(
      infra.pricing?.[serviceDeployment.profile],
      `The pricing for the "${serviceDeployment.profile}" profile is not defined in the "${deploymentName}" "placement" definition.`,
    );
    SdlValidationError.assert(compute, `The compute requirements for the "${serviceDeployment.profile}" profile are not defined in the "compute" section.`);
  }

  private validateServiceStorages(serviceName: string, deploymentName: string) {
    const service = this.data.services[serviceName];
    const mounts: Record<string, string> = {};
    const serviceDeployment = this.data.deployment[serviceName][deploymentName];
    const compute = this.data.profiles.compute[serviceDeployment.profile];
    const storages = castArray(compute.resources.storage);

    if (!service.params?.storage) {
      return;
    }

    Object.entries(service.params.storage || {}).forEach(([storageName, storage]) => {
      const storageNameExists = storages.some(({ name }) => name === storageName);
      SdlValidationError.assert(storage, `Storage "${storageName}" is not configured.`);
      SdlValidationError.assert(storageNameExists, `Service "${serviceName}" references to non-existing compute volume names "${storageName}".`);

      SdlValidationError.assert(
        !("mount" in storage) || this.ABSOLUTE_PATH_REGEX.test(storage.mount),
        `Invalid value for "service.${serviceName}.params.${storageName}.mount" parameter. expected absolute path.`,
      );

      const mount = storage?.mount;
      const volumeName = mounts[mount];

      SdlValidationError.assert(!volumeName || mount, "Multiple root ephemeral storages are not allowed");
      SdlValidationError.assert(!volumeName || !mount, `Mount ${mount} already in use by volume "${volumeName}".`);

      mounts[mount] = storageName;
    });
  }

  private validateStorages(serviceName: string, deploymentName: string) {
    const service = this.data.services[serviceName];
    const serviceDeployment = this.data.deployment[serviceName][deploymentName];
    const compute = this.data.profiles.compute[serviceDeployment.profile];
    const storages = castArray(compute.resources.storage);

    storages.forEach((storage) => {
      const isRam = storage.attributes?.class === "ram";
      const persistent = this.stringToBoolean(storage.attributes?.persistent || false);

      SdlValidationError.assert(storage.size, `Storage size is required for service "${serviceName}".`);
      SdlValidationError.assert(
        !isRam || !persistent,
        `Storage attribute "ram" must have "persistent" set to "false" or not defined for service "${serviceName}".`,
      );

      const mount = service.params?.storage?.[storage.name]?.mount;
      SdlValidationError.assert(
        !persistent || mount,
        `compute.storage.${storage.name} has persistent=true which requires service.${serviceName}.params.storage.${storage.name} to have mount.`,
      );
    });
  }

  private stringToBoolean(str: string | boolean) {
    if (typeof str === "boolean") {
      return str;
    }

    switch (str.toLowerCase()) {
      case "false":
      case "no":
      case "0":
      case "":
        return false;
      default:
        return true;
    }
  }

  private validateGPU(serviceName: string, deploymentName: string) {
    const deployment = this.data.deployment[serviceName];
    const compute = this.data.profiles.compute[deployment[deploymentName].profile];
    const gpu = (compute.resources as v3ComputeResources).gpu;

    if (!gpu) {
      return;
    }
    const hasUnits = gpu.units !== 0;
    const hasAttributes = typeof gpu.attributes !== "undefined";
    const hasVendor = hasAttributes && typeof gpu.attributes?.vendor !== "undefined";

    SdlValidationError.assert(typeof gpu.units === "number", `GPU units must be specified for profile "${serviceName}".`);
    SdlValidationError.assert(hasUnits || !hasAttributes, "GPU must not have attributes if units is 0");
    SdlValidationError.assert(!hasUnits || hasAttributes, "GPU must have attributes if units is not 0");
    SdlValidationError.assert(!hasUnits || hasVendor, "GPU must specify a vendor if units is not 0");
    const hasUnsupportedVendor = hasVendor && GPU_SUPPORTED_VENDORS.some((vendor) => vendor in (gpu.attributes?.vendor || {}));
    SdlValidationError.assert(!hasUnits || hasUnsupportedVendor, `GPU must be one of the supported vendors (${GPU_SUPPORTED_VENDORS.join(",")}).`);

    const vendor: string = Object.keys(gpu.attributes?.vendor || {})[0];

    SdlValidationError.assert(
      !hasUnits || !gpu.attributes?.vendor[vendor] || Array.isArray(gpu.attributes.vendor[vendor]),
      `GPU configuration must be an array of GPU models with optional ram.`,
    );
    SdlValidationError.assert(
      !hasUnits
      || !Object.values(gpu.attributes?.vendor || {}).some((models) =>
        models?.some((model) => model.interface && !GPU_SUPPORTED_INTERFACES.includes(model.interface)),
      ),
      `GPU interface must be one of the supported interfaces (${GPU_SUPPORTED_INTERFACES.join(",")}).`,
    );
  }

  private validateLeaseIP(serviceName: string) {
    this.data.services[serviceName].expose?.forEach((expose) => {
      const proto = this.parseServiceProto(expose.proto);

      expose.to?.forEach((to) => {
        if (to.ip?.length > 0) {
          SdlValidationError.assert(to.global, `Error on "${serviceName}", if an IP is declared, the directive must be declared as global.`);
          SdlValidationError.assert(
            this.data.endpoints?.[to.ip],
            `Unknown endpoint "${to.ip}" in service "${serviceName}". Add to the list of endpoints in the "endpoints" section.`,
          );

          this.endpointsUsed.add(to.ip);

          const portKey = `${to.ip}-${expose.as}-${proto}`;
          const otherServiceName = this.portsUsed.get(portKey);
          SdlValidationError.assert(
            !this.portsUsed.has(portKey),
            `IP endpoint ${to.ip} port: ${expose.port} protocol: ${proto} specified by service ${serviceName} already in use by ${otherServiceName}`,
          );
          this.portsUsed.set(portKey, serviceName);
        }
      });
    });
  }

  private validateEndpointsUtility() {
    if (this.data.endpoints) {
      Object.keys(this.data.endpoints).forEach((endpoint) => {
        SdlValidationError.assert(this.endpointsUsed.has(endpoint), `Endpoint ${endpoint} declared but never used.`);
      });
    }
  }

  services() {
    if (this.data) {
      return this.data.services;
    }

    return {};
  }

  deployments() {
    if (this.data) {
      return this.data.deployment;
    }

    return {};
  }

  profiles() {
    if (this.data) {
      return this.data.profiles;
    }

    return {};
  }

  placements() {
    const { placement } = this.data.profiles;

    return placement || {};
  }

  serviceNames() {
    const names = this.data ? Object.keys(this.data.services) : [];

    // TODO: sort these
    return names;
  }

  deploymentsByPlacement(placement: string) {
    const deployments = this.data ? this.data.deployment : [];

    return Object.entries(deployments as object).filter(({ 1: deployment }) => Object.prototype.hasOwnProperty.call(deployment, placement));
  }

  resourceUnit(val: string, asString: boolean) {
    return asString ? { val: `${convertResourceString(val)}` } : { val: convertResourceString(val) };
  }

  resourceValue(value: { toString: () => string } | null, asString: boolean) {
    if (value === null) {
      return value;
    }

    const strVal = value.toString();
    const encoder = new TextEncoder();

    return asString ? strVal : encoder.encode(strVal);
  }

  serviceResourceCpu(resource: v2ResourceCPU) {
    const units = isString(resource.units) ? convertCpuResourceString(resource.units) : resource.units * 1000;

    return resource.attributes
      ? {
          units: { val: `${units}` },
          attributes: this.serviceResourceAttributes(resource.attributes),
        }
      : {
          units: { val: `${units}` },
        };
  }

  serviceResourceMemory(resource: v2ResourceMemory, asString: boolean) {
    const key = asString ? "quantity" : "size";

    return resource.attributes
      ? {
          [key]: this.resourceUnit(resource.size, asString),
          attributes: this.serviceResourceAttributes(resource.attributes),
        }
      : {
          [key]: this.resourceUnit(resource.size, asString),
        };
  }

  serviceResourceStorage(resource: v2ResourceStorageArray | v2ResourceStorage, asString: boolean) {
    const key = asString ? "quantity" : "size";
    const storage = isArray(resource) ? resource : [resource];

    return storage.map((storage) =>
      storage.attributes
        ? {
            name: storage.name || "default",
            [key]: this.resourceUnit(storage.size, asString),
            attributes: this.serviceResourceStorageAttributes(storage.attributes),
          }
        : {
            name: storage.name || "default",
            [key]: this.resourceUnit(storage.size, asString),
          },
    );
  }

  serviceResourceAttributes(attributes?: Record<string, any>) {
    return (
      attributes
      && Object.keys(attributes)
        .sort()
        .map((key) => ({ key, value: attributes[key].toString() }))
    );
  }

  serviceResourceStorageAttributes(attributes?: v2StorageAttributes) {
    if (!attributes) return undefined;

    const pairs = Object.keys(attributes).map((key) => ({ key, value: attributes[key].toString() }));

    if (attributes.class === "ram" && !("persistent" in attributes)) {
      pairs.push({ key: "persistent", value: "false" });
    }

    pairs.sort((a, b) => a.key.localeCompare(b.key));

    return pairs;
  }

  serviceResourceGpu(resource: v3ResourceGPU | undefined, asString: boolean) {
    const value = resource?.units || 0;
    const numVal = isString(value) ? Buffer.from(value, "ascii") : value;
    const strVal = !isString(value) ? value.toString() : value;

    return resource?.attributes
      ? {
          units: asString ? { val: strVal } : { val: numVal },
          attributes: this.transformGpuAttributes(resource?.attributes),
        }
      : {
          units: asString ? { val: strVal } : { val: numVal },
        };
  }

  v2ServiceResourceEndpoints(service: v2Service) {
    const endpointSequenceNumbers = this.computeEndpointSequenceNumbers(this.data);
    const endpoints = service.expose.flatMap((expose) =>
      expose.to
        ? expose.to
            .filter((to) => to.global && to.ip?.length > 0)
            .map((to) => ({
              kind: Endpoint_LEASED_IP,
              sequence_number: endpointSequenceNumbers[to.ip] || 0,
            }))
        : [],
    );

    return endpoints.length > 0 ? endpoints : null;
  }

  v3ServiceResourceEndpoints(service: v2Service) {
    const endpointSequenceNumbers = this.computeEndpointSequenceNumbers(this.data);
    const endpoints = service.expose.flatMap((expose) =>
      expose.to
        ? expose.to
            .filter((to) => to.global)
            .flatMap((to) => {
              const exposeSpec = {
                port: expose.port,
                externalPort: expose.as || 0,
                proto: this.parseServiceProto(expose.proto),
                global: !!to.global,
              };

              const kind = this.exposeShouldBeIngress(exposeSpec) ? Endpoint_SHARED_HTTP : Endpoint_RANDOM_PORT;

              const defaultEp = kind !== 0 ? { kind: kind, sequence_number: 0 } : { sequence_number: 0 };

              const leasedEp
                = to.ip?.length > 0
                  ? {
                      kind: Endpoint_LEASED_IP,
                      sequence_number: endpointSequenceNumbers[to.ip] || 0,
                    }
                  : undefined;

              return leasedEp ? [defaultEp, leasedEp] : [defaultEp];
            })
        : [],
    );

    return endpoints;
  }

  serviceResourcesBeta2(profile: v2ProfileCompute, service: v2Service, asString: boolean = false) {
    return {
      cpu: this.serviceResourceCpu(profile.resources.cpu),
      memory: this.serviceResourceMemory(profile.resources.memory, asString),
      storage: this.serviceResourceStorage(profile.resources.storage, asString),
      endpoints: this.v2ServiceResourceEndpoints(service),
    };
  }

  serviceResourcesBeta3(id: number, profile: v3ProfileCompute, service: v2Service, asString: boolean = false) {
    return {
      id: id,
      cpu: this.serviceResourceCpu(profile.resources.cpu),
      memory: this.serviceResourceMemory(profile.resources.memory, asString),
      storage: this.serviceResourceStorage(profile.resources.storage, asString),
      endpoints: this.v3ServiceResourceEndpoints(service),
      gpu: this.serviceResourceGpu(profile.resources.gpu, asString),
    };
  }

  /**
   * Parses the service protocol.
   *
   * @param proto - The protocol string (e.g., "TCP", "UDP").
   * @returns The parsed protocol.
   * @throws Will throw an error if the protocol is unsupported.
   *
   * @example
   * ```ts
   * const protocol = SDL.parseServiceProto("TCP");
   * // protocol is "TCP"
   * ```
   */
  parseServiceProto(proto?: string) {
    const raw = proto?.toUpperCase();
    let result = "TCP";

    switch (raw) {
      case "TCP":
      case "":
      case undefined:
        result = "TCP";
        break;
      case "UDP":
        result = "UDP";
        break;
      default:
        throw new Error("ErrUnsupportedServiceProtocol");
    }

    return result;
  }

  manifestExposeService(to: v2ExposeTo) {
    return to.service || "";
  }

  manifestExposeGlobal(to: v2ExposeTo) {
    return to.global || false;
  }

  manifestExposeHosts(expose: v2Expose) {
    return expose.accept || null;
  }

  v2HttpOptions(http_options: v2HTTPOptions | undefined) {
    const defaults = {
      MaxBodySize: 1048576,
      ReadTimeout: 60000,
      SendTimeout: 60000,
      NextTries: 3,
      NextTimeout: 0,
      NextCases: ["error", "timeout"],
    };

    if (!http_options) {
      return { ...defaults };
    }

    return {
      MaxBodySize: http_options.max_body_size || defaults.MaxBodySize,
      ReadTimeout: http_options.read_timeout || defaults.ReadTimeout,
      SendTimeout: http_options.send_timeout || defaults.SendTimeout,
      NextTries: http_options.next_tries || defaults.NextTries,
      NextTimeout: http_options.next_timeout || defaults.NextTimeout,
      NextCases: http_options.next_cases || defaults.NextCases,
    };
  }

  v3HttpOptions(http_options: v2HTTPOptions | undefined) {
    const defaults = {
      maxBodySize: 1048576,
      readTimeout: 60000,
      sendTimeout: 60000,
      nextTries: 3,
      nextTimeout: 0,
      nextCases: ["error", "timeout"],
    };

    if (!http_options) {
      return { ...defaults };
    }

    return {
      maxBodySize: http_options.max_body_size || defaults.maxBodySize,
      readTimeout: http_options.read_timeout || defaults.readTimeout,
      sendTimeout: http_options.send_timeout || defaults.sendTimeout,
      nextTries: http_options.next_tries || defaults.nextTries,
      nextTimeout: http_options.next_timeout || defaults.nextTimeout,
      nextCases: http_options.next_cases || defaults.nextCases,
    };
  }

  v2ManifestExposeHttpOptions(expose: v2Expose): v2ServiceExposeHttpOptions {
    return this.v2HttpOptions(expose.http_options);
  }

  v3ManifestExposeHttpOptions(expose: v2Expose): v3ServiceExposeHttpOptions {
    return this.v3HttpOptions(expose.http_options);
  }

  v2ManifestExpose(service: v2Service): v2ServiceExpose[] {
    const endpointSequenceNumbers = this.computeEndpointSequenceNumbers(this.data);
    return service.expose.flatMap((expose) =>
      expose.to
        ? expose.to.map((to) => ({
            Port: expose.port,
            ExternalPort: expose.as || 0,
            Proto: this.parseServiceProto(expose.proto),
            Service: this.manifestExposeService(to),
            Global: this.manifestExposeGlobal(to),
            Hosts: this.manifestExposeHosts(expose),
            HTTPOptions: this.v2ManifestExposeHttpOptions(expose),
            IP: to.ip || "",
            EndpointSequenceNumber: endpointSequenceNumbers[to.ip] || 0,
          }))
        : [],
    );
  }

  v3ManifestExpose(service: v2Service): v3ServiceExpose[] {
    const endpointSequenceNumbers = this.computeEndpointSequenceNumbers(this.data);
    return service.expose
      .flatMap((expose) =>
        expose.to
          ? expose.to.map((to) => ({
              port: expose.port,
              externalPort: expose.as || 0,
              proto: this.parseServiceProto(expose.proto),
              service: this.manifestExposeService(to),
              global: this.manifestExposeGlobal(to),
              hosts: this.manifestExposeHosts(expose),
              httpOptions: this.v3ManifestExposeHttpOptions(expose),
              ip: to.ip || "",
              endpointSequenceNumber: endpointSequenceNumbers[to.ip] || 0,
            }))
          : [],
      )
      .sort((a, b) => {
        if (a.service != b.service) return a.service.localeCompare(b.service);
        if (a.port != b.port) return a.port - b.port;
        if (a.proto != b.proto) return a.proto.localeCompare(b.proto);
        if (a.global != b.global) return a.global ? -1 : 1;

        return 0;
      });
  }

  v2ManifestServiceParams(params: v2ServiceParams): v2ManifestServiceParams | undefined {
    return {
      Storage: Object.keys(params?.storage ?? {}).map((name) => {
        if (!params?.storage) throw new Error("Storage is undefined");
        return {
          name: name,
          mount: params.storage[name].mount,
          readOnly: params.storage[name].readOnly || false,
        };
      }),
    };
  }

  v3ManifestServiceParams(params: v2ServiceParams | undefined): v3ManifestServiceParams | null {
    if (params === undefined) {
      return null;
    }

    return {
      storage: Object.keys(params?.storage ?? {}).map((name) => {
        if (!params?.storage) throw new Error("Storage is undefined");
        return {
          name: name,
          mount: params.storage[name]?.mount,
          readOnly: params.storage[name]?.readOnly || false,
        };
      }),
    };
  }

  v2ManifestService(placement: string, name: string, asString: boolean): v2ManifestService {
    const service = this.data.services[name];
    const deployment = this.data.deployment[name];
    const profile = this.data.profiles.compute[deployment[placement].profile];

    const manifestService: v2ManifestService = {
      Name: name,
      Image: service.image,
      Command: service.command || null,
      Args: service.args || null,
      Env: service.env || null,
      Resources: this.serviceResourcesBeta2(profile, service, asString),
      Count: deployment[placement].count,
      Expose: this.v2ManifestExpose(service),
    };

    if (service.params) {
      manifestService.params = this.v2ManifestServiceParams(service.params);
    }

    return manifestService;
  }

  v3ManifestService(id: number, placement: string, name: string, asString: boolean): v3ManifestService {
    const service = this.data.services[name];
    const deployment = this.data.deployment[name];
    const profile = this.data.profiles.compute[deployment[placement].profile];
    const credentials = service.credentials || null;

    if (credentials && !credentials.email) {
      credentials.email = "";
    }

    return {
      name: name,
      image: service.image,
      command: service.command || null,
      args: service.args || null,
      env: service.env || null,
      resources: this.serviceResourcesBeta3(id, profile as v3ProfileCompute, service, asString),
      count: deployment[placement].count,
      expose: this.v3ManifestExpose(service),
      params: this.v3ManifestServiceParams(service.params),
      credentials,
    };
  }

  v2Manifest(asString: boolean = false): v2Manifest {
    return Object.keys(this.placements()).map((name) => ({
      Name: name,
      Services: this.deploymentsByPlacement(name).map(([service]) => this.v2ManifestService(name, service, asString)),
    }));
  }

  v3Manifest(asString: boolean = false): v3Manifest {
    const groups = this.v3Groups();
    const serviceId = (pIdx: number, sIdx: number) => groups[pIdx].resources[sIdx].resource.id;

    return Object.keys(this.placements()).map((name, pIdx) => ({
      name: name,
      services: this.deploymentsByPlacement(name)
        .sort(([a], [b]) => a.localeCompare(b))
        .map(([service], idx) => this.v3ManifestService(serviceId(pIdx, idx), name, service, asString)),
    }));
  }

  manifest(asString: boolean = false): v2Manifest | v3Manifest {
    return this.version === "beta2" ? this.v2Manifest(asString) : this.v3Manifest(asString);
  }

  /**
   * Computes the endpoint sequence numbers for the given SDL.
   *
   * @param sdl - The SDL data.
   * @returns An object mapping IPs to their sequence numbers.
   *
   * @example
   * ```ts
   * const sequenceNumbers = sdl.computeEndpointSequenceNumbers(sdlData);
   * // sequenceNumbers might be { "192.168.1.1": 1, "192.168.1.2": 2 }
   * ```
   */
  computeEndpointSequenceNumbers(sdl: v2Sdl) {
    return Object.fromEntries(
      Object.values(sdl.services).flatMap((service) =>
        service.expose.flatMap((expose) =>
          expose.to
            ? expose.to
                .filter((to) => to.global && to.ip?.length > 0)
                .map((to) => to.ip)
                .sort()
                .map((ip, index) => [ip, index + 1])
            : [],
        ),
      ),
    );
  }

  resourceUnitCpu(computeResources: v2ComputeResources, asString: boolean) {
    const attributes = computeResources.cpu.attributes;
    const cpu = isString(computeResources.cpu.units) ? convertCpuResourceString(computeResources.cpu.units) : computeResources.cpu.units * 1000;

    return {
      units: { val: this.resourceValue(cpu, asString) },
      attributes:
        attributes
        && Object.entries(attributes)
          .sort(([k0], [k1]) => k0.localeCompare(k1))
          .map(([key, value]) => ({
            key: key,
            value: value.toString(),
          })),
    };
  }

  resourceUnitMemory(computeResources: v2ComputeResources, asString: boolean) {
    const attributes = computeResources.memory.attributes;

    return {
      quantity: {
        val: this.resourceValue(convertResourceString(computeResources.memory.size), asString),
      },
      attributes:
        attributes
        && Object.entries(attributes)
          .sort(([k0], [k1]) => k0.localeCompare(k1))
          .map(([key, value]) => ({
            key: key,
            value: value.toString(),
          })),
    };
  }

  resourceUnitStorage(computeResources: v2ComputeResources, asString: boolean) {
    const storages = isArray(computeResources.storage) ? computeResources.storage : [computeResources.storage];

    return storages.map((storage) => ({
      name: storage.name || "default",
      quantity: {
        val: this.resourceValue(convertResourceString(storage.size), asString),
      },
      attributes: this.serviceResourceStorageAttributes(storage.attributes),
    }));
  }

  transformGpuAttributes(attributes: v3GPUAttributes): Array<{ key: string; value: string }> {
    return Object.entries(attributes.vendor).flatMap(([vendor, models]) =>
      models
        ? models.map((model) => {
            let key = `vendor/${vendor}/model/${model.model}`;

            if (model.ram) {
              key += `/ram/${model.ram}`;
            }

            if (model.interface) {
              key += `/interface/${model.interface}`;
            }

            return {
              key: key,
              value: "true",
            };
          })
        : [
            {
              key: `vendor/${vendor}/model/*`,
              value: "true",
            },
          ],
    );
  }

  resourceUnitGpu(computeResources: v3ComputeResources, asString: boolean) {
    const attributes = computeResources.gpu?.attributes;
    const units = computeResources.gpu?.units || "0";
    const gpu = isString(units) ? parseInt(units) : units;

    return {
      units: { val: this.resourceValue(gpu, asString) },
      attributes: attributes && this.transformGpuAttributes(attributes),
    };
  }

  groupResourceUnits(resource: v2ComputeResources | undefined, asString: boolean) {
    if (!resource) return {};

    const units = {
      endpoints: null,
    } as any;

    if (resource.cpu) {
      units.cpu = this.resourceUnitCpu(resource, asString);
    }

    if (resource.memory) {
      units.memory = this.resourceUnitMemory(resource, asString);
    }

    if (resource.storage) {
      units.storage = this.resourceUnitStorage(resource, asString);
    }

    if (this.version === "beta3") {
      units.gpu = this.resourceUnitGpu(resource as v3ComputeResources, asString);
    }

    return units;
  }

  exposeShouldBeIngress(expose: { proto: string; global: boolean; externalPort: number; port: number }) {
    const externalPort = expose.externalPort === 0 ? expose.port : expose.externalPort;

    return expose.global && expose.proto === "TCP" && externalPort === 80;
  }

  groups() {
    return this.version === "beta2" ? this.v2Groups() : this.v3Groups();
  }

  v3Groups() {
    const groups = new Map<
      string,
      {
        dgroup: v3DeploymentGroup;
        boundComputes: Record<string, Record<string, number>>;
      }
    >();
    const services = Object.entries(this.data.services).sort(([a], [b]) => a.localeCompare(b));

    for (const [svcName, service] of services) {
      for (const [placementName, svcdepl] of Object.entries(this.data.deployment[svcName])) {
        // objects below have been ensured to exist
        const compute = this.data.profiles.compute[svcdepl.profile];
        const infra = this.data.profiles.placement[placementName];
        const pricing = infra.pricing[svcdepl.profile];
        const price = {
          ...pricing,
          amount: pricing.amount?.toString(),
        };

        let group = groups.get(placementName);

        if (!group) {
          const attributes = (infra.attributes
            ? Object.entries(infra.attributes).map(([key, value]) => ({
                key,
                value,
              }))
            : []) as unknown as Array<{ key: string; value: string }>;

          attributes.sort((a, b) => a.key.localeCompare(b.key));

          group = {
            dgroup: {
              name: placementName,
              resources: [],
              requirements: {
                attributes: attributes,
                signedBy: {
                  allOf: infra.signedBy?.allOf || [],
                  anyOf: infra.signedBy?.anyOf || [],
                },
              },
            },
            boundComputes: {},
          };

          groups.set(placementName, group);
        }

        if (!group.boundComputes[placementName]) {
          group.boundComputes[placementName] = {};
        }

        // const resources = this.serviceResourcesBeta3(0, compute as v3ProfileCompute, service, false);
        const location = group.boundComputes[placementName][svcdepl.profile];

        if (!location) {
          const res = this.groupResourceUnits(compute.resources, false);
          res.endpoints = this.v3ServiceResourceEndpoints(service);

          const resID = group.dgroup.resources.length > 0 ? group.dgroup.resources.length + 1 : 1;
          res.id = resID;
          // resources.id = res.id;

          group.dgroup.resources.push({
            resource: res,
            price: price,
            count: svcdepl.count,
          } as any);

          group.boundComputes[placementName][svcdepl.profile] = group.dgroup.resources.length - 1;
        } else {
          const endpoints = this.v3ServiceResourceEndpoints(service);
          // resources.id = group.dgroup.resources[location].id;

          group.dgroup.resources[location].count += svcdepl.count;
          group.dgroup.resources[location].endpoints += endpoints as any;
          group.dgroup.resources[location].endpoints.sort();
        }
      }
    }

    // keep ordering stable
    const names: string[] = [...groups.keys()].sort();
    return names.map((name) => groups.get(name)).map((group) => (group ? (group.dgroup as typeof group.dgroup) : {})) as Array<v3DeploymentGroup>;
  }

  v2Groups() {
    const yamlJson = this.data;
    const ipEndpointNames = this.computeEndpointSequenceNumbers(yamlJson);

    const groups = {} as any;

    Object.keys(yamlJson.services).forEach((svcName) => {
      const svc = yamlJson.services[svcName];
      const depl = yamlJson.deployment[svcName];

      Object.keys(depl).forEach((placementName) => {
        const svcdepl = depl[placementName];
        const compute = yamlJson.profiles.compute[svcdepl.profile];
        const infra = yamlJson.profiles.placement[placementName];

        const pricing = infra.pricing[svcdepl.profile];
        const price = {
          ...pricing,
          amount: pricing.amount.toString(),
        };

        let group = groups[placementName];

        if (!group) {
          group = {
            name: placementName,
            requirements: {
              attributes: infra.attributes
                ? Object.entries(infra.attributes).map(([key, value]) => ({
                    key,
                    value,
                  }))
                : [],
              signedBy: {
                allOf: infra.signedBy?.allOf || [],
                anyOf: infra.signedBy?.anyOf || [],
              },
            },
            resources: [],
          };

          if (group.requirements.attributes) {
            group.requirements.attributes = group.requirements.attributes.sort((a: any, b: any) => a.key < b.key);
          }

          groups[group.name] = group;
        }

        const resources = {
          resources: this.groupResourceUnits(compute.resources, false), // Changed resources => unit
          price: price,
          count: svcdepl.count,
        };

        const endpoints = [] as any[];
        svc?.expose?.forEach((expose) => {
          expose?.to
            ?.filter((to) => to.global)
            .forEach((to) => {
              const exposeSpec = {
                port: expose.port,
                externalPort: expose.as || 0,
                proto: this.parseServiceProto(expose.proto),
                global: !!to.global,
              };

              if (to.ip?.length > 0) {
                const seqNo = ipEndpointNames[to.ip];
                endpoints.push({
                  kind: Endpoint_LEASED_IP,
                  sequence_number: seqNo,
                });
              }

              const kind = this.exposeShouldBeIngress(exposeSpec) ? Endpoint_SHARED_HTTP : Endpoint_RANDOM_PORT;

              endpoints.push({ kind: kind, sequence_number: 0 });
            });
        });

        resources.resources.endpoints = endpoints;
        group.resources.push(resources);
      });
    });

    return Object.keys(groups)
      .sort((a, b) => (a < b ? 1 : 0))
      .map((name) => groups[name]);
  }

  /**
   * Escapes HTML characters in a string.
   *
   * @param raw - The raw string to escape.
   * @returns The escaped string.
   *
   * @example
   * ```ts
   * const escaped = sdl.escapeHtml("<div>Hello</div>");
   * // escaped is "\\u003cdiv\\u003eHello\\u003c/div\\u003e"
   * ```
   */
  escapeHtml(raw: string) {
    return raw.replace(/</g, "\\u003c").replace(/>/g, "\\u003e").replace(/&/g, "\\u0026");
  }

  SortJSON(jsonStr: string) {
    return this.escapeHtml(stableStringify(JSON.parse(jsonStr)) || "");
  }

  manifestSortedJSON() {
    const manifest = this.manifest(true);
    let jsonStr = JSON.stringify(manifest);

    if (jsonStr) {
      jsonStr = jsonStr.replaceAll("\"quantity\":{\"val", "\"size\":{\"val");
    }

    return this.SortJSON(jsonStr);
  }

  async manifestVersion() {
    const jsonStr = this.manifestSortedJSON();
    const enc = new TextEncoder();
    const sortedBytes = enc.encode(jsonStr);
    const sum = await crypto.subtle.digest("SHA-256", sortedBytes);

    return new Uint8Array(sum);
  }

  manifestSorted() {
    const sorted = this.manifestSortedJSON();
    return JSON.parse(sorted);
  }
}
