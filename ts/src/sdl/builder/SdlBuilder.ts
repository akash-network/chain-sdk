/**
 * Fluent builder API for creating complex SDL configurations
 */

import YAML from "js-yaml";

import type { SdlSchemaValidationError } from "../validateSdl.ts";
import { SDL_DEFAULTS } from "./defaults.ts";
import type {
  BuilderComputeResources,
  BuilderDeployConfig,
  BuilderExposeConfig,
  BuilderPlacementConfig,
  BuilderServiceConfig,
  SdlObject,
  SimpleGpuConfig,
  SimpleStorageConfig,
} from "./types.ts";

/**
 * Error thrown when SDL validation fails
 */
export class SdlBuilderError extends Error {
  constructor(
    message: string,
    public errors: SdlSchemaValidationError[],
  ) {
    super(message);
    this.name = "SdlBuilderError";
  }
}

/**
 * Fluent builder for creating SDL YAML configurations.
 *
 * @example
 * ```ts
 * const yaml = new SdlBuilder()
 *   .version("2.0")
 *   .service("web", {
 *     image: "nginx",
 *     expose: [{ port: 80, global: true }],
 *   })
 *   .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
 *   .placement("akash", { attributes: { region: "us-west" } })
 *   .deploy("web", "akash", { profile: "web", count: 2 })
 *   .pricing("akash", { web: 1000 })
 *   .build();
 * ```
 */
export class SdlBuilder {
  private _version: string = SDL_DEFAULTS.version;

  private _services: Map<string, BuilderServiceConfig> = new Map();

  private _computeProfiles: Map<string, BuilderComputeResources> = new Map();

  private _placements: Map<string, BuilderPlacementConfig> = new Map();

  private _deployments: Map<string, Map<string, BuilderDeployConfig>>
    = new Map();

  private _pricing: Map<string, { denom: string; prices: Map<string, number> }>
    = new Map();

  private _endpoints: Map<string, { kind: string }> = new Map();

  /**
   * Set the SDL version
   *
   * @param version - SDL version (e.g., "2.0", "2.1")
   * @returns The builder instance for chaining
   */
  version(version: string): this {
    this._version = version;
    return this;
  }

  /**
   * Add a service definition
   *
   * @param name - Service name
   * @param config - Service configuration
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.service("web", {
   *   image: "nginx",
   *   expose: [{ port: 80, global: true }],
   *   env: ["PORT=80"],
   * });
   * ```
   */
  service(name: string, config: BuilderServiceConfig): this {
    this._services.set(name, config);
    return this;
  }

  /**
   * Add a compute profile with resource requirements
   *
   * @param name - Profile name
   * @param resources - Resource requirements
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.computeProfile("web", {
   *   cpu: "500m",
   *   memory: "512Mi",
   *   storage: "1Gi",
   * });
   *
   * // With GPU
   * builder.computeProfile("ml", {
   *   cpu: "4",
   *   memory: "16Gi",
   *   storage: "100Gi",
   *   gpu: { units: 1, vendor: "nvidia", model: "a100" },
   * });
   * ```
   */
  computeProfile(name: string, resources: BuilderComputeResources): this {
    this._computeProfiles.set(name, resources);
    return this;
  }

  /**
   * Add a placement configuration
   *
   * @param name - Placement name
   * @param config - Placement configuration
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.placement("akash", {
   *   attributes: { region: "us-west" },
   *   signedBy: { anyOf: ["akash1..."] },
   * });
   * ```
   */
  placement(name: string, config: BuilderPlacementConfig = {}): this {
    this._placements.set(name, config);
    return this;
  }

  /**
   * Add a deployment mapping service to placement
   *
   * @param serviceName - Service name
   * @param placementName - Placement name
   * @param config - Deployment configuration
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.deploy("web", "akash", { profile: "web", count: 2 });
   * ```
   */
  deploy(
    serviceName: string,
    placementName: string,
    config: BuilderDeployConfig,
  ): this {
    if (!this._deployments.has(serviceName)) {
      this._deployments.set(serviceName, new Map());
    }
    this._deployments.get(serviceName)!.set(placementName, config);
    return this;
  }

  /**
   * Set pricing for compute profiles in a placement
   *
   * @param placementName - Placement name
   * @param prices - Map of profile names to prices in uakt
   * @param denom - Token denomination (defaults to "uakt")
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.pricing("akash", { web: 1000, api: 2000 });
   * ```
   */
  pricing(
    placementName: string,
    prices: Record<string, number>,
    denom: string = SDL_DEFAULTS.denom,
  ): this {
    if (!this._pricing.has(placementName)) {
      this._pricing.set(placementName, { denom, prices: new Map() });
    }
    const placementPricing = this._pricing.get(placementName)!;
    placementPricing.denom = denom;
    for (const [profile, amount] of Object.entries(prices)) {
      placementPricing.prices.set(profile, amount);
    }
    return this;
  }

  /**
   * Add an endpoint definition
   *
   * @param name - Endpoint name
   * @param kind - Endpoint kind (currently only "ip" is supported)
   * @returns The builder instance for chaining
   *
   * @example
   * ```ts
   * builder.endpoint("my-ip", "ip");
   * ```
   */
  endpoint(name: string, kind: "ip" = "ip"): this {
    this._endpoints.set(name, { kind });
    return this;
  }

  /**
   * Build the SDL object without converting to YAML
   *
   * @returns The raw SDL object
   */
  toObject(): SdlObject {
    // Build services
    const services: Record<string, unknown> = {};
    for (const [name, config] of this._services) {
      services[name] = this.buildService(config);
    }

    // Build compute profiles
    const compute: Record<string, unknown> = {};
    for (const [name, resources] of this._computeProfiles) {
      compute[name] = {
        resources: this.buildComputeResources(resources),
      };
    }

    // Build placements with pricing
    const placement: Record<string, unknown> = {};
    for (const [name, config] of this._placements) {
      const placementConfig: Record<string, unknown> = {};

      if (config.attributes) {
        placementConfig.attributes = config.attributes;
      }

      if (config.signedBy) {
        placementConfig.signedBy = config.signedBy;
      }

      // Add pricing for this placement
      const pricingData = this._pricing.get(name);
      if (pricingData) {
        const pricing: Record<string, unknown> = {};
        for (const [profile, amount] of pricingData.prices) {
          pricing[profile] = { denom: pricingData.denom, amount };
        }
        placementConfig.pricing = pricing;
      }

      placement[name] = placementConfig;
    }

    // Build deployments
    const deployment: Record<string, unknown> = {};
    for (const [serviceName, placements] of this._deployments) {
      const serviceDeployment: Record<string, unknown> = {};
      for (const [placementName, config] of placements) {
        serviceDeployment[placementName] = config;
      }
      deployment[serviceName] = serviceDeployment;
    }

    const sdl: SdlObject = {
      version: this._version,
      services,
      profiles: { compute, placement },
      deployment,
    };

    // Add endpoints if any
    if (this._endpoints.size > 0) {
      const endpoints: Record<string, unknown> = {};
      for (const [name, config] of this._endpoints) {
        endpoints[name] = config;
      }
      sdl.endpoints = endpoints;
    }

    return sdl;
  }

  /**
   * Build and validate the SDL, returning YAML
   *
   * @returns The SDL as a YAML string
   * @throws {SdlBuilderError} If validation fails
   */
  build(): string {
    const sdlObject = this.toObject();

    return YAML.dump(sdlObject, {
      indent: 2,
      lineWidth: -1,
      noRefs: true,
      sortKeys: false,
    });
  }

  /**
   * Build and return both YAML and object
   *
   * @returns Object containing YAML and raw SDL object
   */
  buildWithObject(): { yaml: string; object: SdlObject } {
    const object = this.toObject();
    const yaml = YAML.dump(object, {
      indent: 2,
      lineWidth: -1,
      noRefs: true,
      sortKeys: false,
    });
    return { yaml, object };
  }

  private buildService(config: BuilderServiceConfig): Record<string, unknown> {
    const service: Record<string, unknown> = {
      image: config.image,
    };

    if (config.command) service.command = config.command;
    if (config.args) service.args = config.args;
    if (config.env) service.env = config.env;
    if (config.depends_on) service.depends_on = config.depends_on;
    if (config.credentials) service.credentials = config.credentials;
    if (config.params) service.params = config.params;

    if (config.expose && config.expose.length > 0) {
      service.expose = config.expose.map((e) => this.buildExpose(e));
    }

    return service;
  }

  private buildExpose(config: BuilderExposeConfig): Record<string, unknown> {
    const expose: Record<string, unknown> = {
      port: config.port,
    };

    if (config.as !== undefined) expose.as = config.as;
    if (config.proto) expose.proto = config.proto;
    if (config.accept) expose.accept = config.accept;
    if (config.http_options) expose.http_options = config.http_options;

    // Build "to" array
    const to: Array<Record<string, unknown>> = [];

    if (config.global !== undefined || config.service || config.ip) {
      const toEntry: Record<string, unknown> = {};
      if (config.global !== undefined) toEntry.global = config.global;
      if (config.service) toEntry.service = config.service;
      if (config.ip) toEntry.ip = config.ip;
      to.push(toEntry);
    }

    if (to.length > 0) {
      expose.to = to;
    }

    return expose;
  }

  private buildComputeResources(
    resources: BuilderComputeResources,
  ): Record<string, unknown> {
    const cpu
      = typeof resources.cpu === "number"
        ? resources.cpu < 10
          ? `${Math.round(resources.cpu * 1000)}m`
          : `${resources.cpu}m`
        : resources.cpu;

    const result: Record<string, unknown> = {
      cpu: { units: cpu },
      memory: { size: resources.memory },
      storage: this.buildStorageResources(resources.storage),
    };

    if (resources.gpu) {
      result.gpu = this.buildGpuResource(resources.gpu);
    }

    return result;
  }

  private buildStorageResources(
    storage: string | SimpleStorageConfig | SimpleStorageConfig[],
  ): Array<Record<string, unknown>> {
    const storages = Array.isArray(storage)
      ? storage
      : typeof storage === "string"
        ? [{ size: storage }]
        : [storage];

    return storages.map((s, index) => {
      const result: Record<string, unknown> = {
        name: s.name ?? (index === 0 ? "default" : `storage-${index}`),
        size: s.size,
      };

      if (s.persistent || s.class) {
        const attrs: Record<string, unknown> = {};
        if (s.persistent) attrs.persistent = true;
        if (s.class) attrs.class = s.class;
        result.attributes = attrs;
      }

      return result;
    });
  }

  private buildGpuResource(
    gpu: SimpleGpuConfig | { units: 0 },
  ): Record<string, unknown> {
    if ("vendor" in gpu) {
      const gpuConfig = gpu as SimpleGpuConfig;
      const models: Array<{ model: string; ram?: string; interface?: string }>
        = [];

      if (gpuConfig.model) {
        const modelConfig: { model: string; ram?: string; interface?: string }
          = { model: gpuConfig.model };
        if (gpuConfig.ram) modelConfig.ram = gpuConfig.ram;
        if (gpuConfig.interface) modelConfig.interface = gpuConfig.interface;
        models.push(modelConfig);
      }

      return {
        units: gpuConfig.units,
        attributes: {
          vendor: {
            [gpuConfig.vendor]: models.length > 0 ? models : undefined,
          },
        },
      };
    }

    return { units: 0 };
  }
}
