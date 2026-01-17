/**
 * Simple API for creating SDL YAML from minimal configuration
 */

import YAML from "js-yaml";

import type { SdlSchemaValidationError } from "../validateSdl.ts";
import { SDL_DEFAULTS } from "./defaults.ts";
import type {
  CreateSdlOptions,
  SdlObject,
  SimpleGpuConfig,
  SimplePortConfig,
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
 * Normalize CPU value to string format
 */
function normalizeCpu(cpu: string | number): string {
  if (typeof cpu === "number") {
    // Convert decimal to millicores if less than 10, otherwise treat as millicores
    if (cpu < 10) {
      return `${Math.round(cpu * 1000)}m`;
    }
    return `${cpu}m`;
  }
  return cpu;
}

/**
 * Normalize storage configuration to array format
 */
function normalizeStorage(
  storage: string | SimpleStorageConfig | SimpleStorageConfig[],
): SimpleStorageConfig[] {
  if (typeof storage === "string") {
    return [{ size: storage }];
  }
  if (Array.isArray(storage)) {
    return storage;
  }
  return [storage];
}

/**
 * Normalize ports configuration
 */
function normalizePorts(
  port?: number,
  ports?: (number | SimplePortConfig)[],
): SimplePortConfig[] {
  const result: SimplePortConfig[] = [];

  if (port !== undefined) {
    result.push({ port, global: true });
  }

  if (ports) {
    for (const p of ports) {
      if (typeof p === "number") {
        result.push({ port: p, global: true });
      } else {
        result.push({ ...p, global: p.global ?? true });
      }
    }
  }

  // Default to port 80 if no ports specified
  if (result.length === 0) {
    result.push({ port: 80, global: true });
  }

  return result;
}

/**
 * Normalize environment variables to array format
 */
function normalizeEnv(
  env?: Record<string, string> | string[],
): string[] | undefined {
  if (!env) return undefined;

  if (Array.isArray(env)) {
    return env;
  }

  return Object.entries(env).map(([key, value]) => `${key}=${value}`);
}

/**
 * Build GPU resource configuration
 */
function buildGpuResource(gpu?: SimpleGpuConfig): object | undefined {
  if (!gpu) {
    return { units: 0 };
  }

  const models: Array<{ model: string; ram?: string; interface?: string }> = [];

  if (gpu.model) {
    const modelConfig: { model: string; ram?: string; interface?: string } = {
      model: gpu.model,
    };
    if (gpu.ram) modelConfig.ram = gpu.ram;
    if (gpu.interface) modelConfig.interface = gpu.interface;
    models.push(modelConfig);
  }

  return {
    units: gpu.units,
    attributes: {
      vendor: {
        [gpu.vendor]: models.length > 0 ? models : undefined,
      },
    },
  };
}

/**
 * Build storage resources for compute profile
 */
function buildStorageResources(
  storages: SimpleStorageConfig[],
): Array<object> {
  return storages.map((storage, index) => {
    const result: Record<string, unknown> = {
      name: storage.name ?? (index === 0 ? "default" : `storage-${index}`),
      size: storage.size,
    };

    if (storage.persistent || storage.class) {
      result.attributes = {};
      if (storage.persistent) {
        (result.attributes as Record<string, unknown>).persistent = true;
      }
      if (storage.class) {
        (result.attributes as Record<string, unknown>).class = storage.class;
      }
    }

    return result;
  });
}

/**
 * Build storage params for service
 */
function buildStorageParams(
  storages: SimpleStorageConfig[],
): Record<string, { mount: string; readOnly?: boolean }> | undefined {
  const params: Record<string, { mount: string; readOnly?: boolean }> = {};
  let hasParams = false;

  for (const storage of storages) {
    if (storage.mount) {
      const name
        = storage.name ?? (storages.indexOf(storage) === 0 ? "default" : `storage-${storages.indexOf(storage)}`);
      params[name] = { mount: storage.mount, readOnly: false };
      hasParams = true;
    }
  }

  return hasParams ? params : undefined;
}

/**
 * Build expose configuration for service
 */
function buildExposeConfig(ports: SimplePortConfig[]): Array<object> {
  return ports.map((port) => {
    const expose: Record<string, unknown> = {
      port: port.port,
    };

    if (port.as !== undefined) {
      expose.as = port.as;
    }

    if (port.proto) {
      expose.proto = port.proto;
    }

    if (port.accept && port.accept.length > 0) {
      expose.accept = port.accept;
    }

    expose.to = [{ global: port.global ?? true }];

    return expose;
  });
}

/**
 * Creates SDL YAML from a simplified configuration object.
 *
 * @param options - Simplified SDL configuration
 * @returns The SDL as a YAML string
 * @throws {SdlBuilderError} If the generated SDL fails validation
 *
 * @example
 * ```ts
 * // Minimal - just an image
 * const yaml = createSdl({ image: "nginx" });
 *
 * // With resources
 * const yaml = createSdl({
 *   image: "nginx:latest",
 *   cpu: "500m",
 *   memory: "512Mi",
 *   storage: "1Gi",
 *   port: 80,
 *   env: { NODE_ENV: "production" },
 * });
 *
 * // With GPU
 * const yaml = createSdl({
 *   image: "pytorch/pytorch",
 *   cpu: 4,
 *   memory: "16Gi",
 *   storage: "100Gi",
 *   gpu: { units: 1, vendor: "nvidia", model: "a100" },
 *   port: 8080,
 * });
 * ```
 */
export function createSdl(options: CreateSdlOptions): string {
  const {
    image,
    name = SDL_DEFAULTS.serviceName,
    cpu = SDL_DEFAULTS.cpu,
    memory = SDL_DEFAULTS.memory,
    storage = SDL_DEFAULTS.storage,
    port,
    ports,
    env,
    command,
    args,
    gpu,
    replicas = SDL_DEFAULTS.replicas,
    pricing = SDL_DEFAULTS.pricing,
    denom = SDL_DEFAULTS.denom,
    placement = SDL_DEFAULTS.placement,
    attributes,
    signedBy,
    credentials,
  } = options;

  // Normalize inputs
  const normalizedCpu = normalizeCpu(cpu);
  const normalizedStorage = normalizeStorage(storage);
  const normalizedPorts = normalizePorts(port, ports);
  const normalizedEnv = normalizeEnv(env);
  const storageParams = buildStorageParams(normalizedStorage);

  // Build service configuration
  const service: Record<string, unknown> = {
    image,
    expose: buildExposeConfig(normalizedPorts),
  };

  if (command) service.command = command;
  if (args) service.args = args;
  if (normalizedEnv) service.env = normalizedEnv;
  if (storageParams) service.params = { storage: storageParams };
  if (credentials) service.credentials = credentials;

  // Build compute resources
  const computeResources: Record<string, unknown> = {
    cpu: { units: normalizedCpu },
    memory: { size: memory },
    storage: buildStorageResources(normalizedStorage),
  };

  // Add GPU if specified
  const gpuResource = buildGpuResource(gpu);
  if (gpuResource) {
    computeResources.gpu = gpuResource;
  }

  // Build placement configuration
  const placementConfig: Record<string, unknown> = {
    pricing: {
      [name]: {
        denom,
        amount: pricing,
      },
    },
  };

  if (attributes) {
    placementConfig.attributes = attributes;
  }

  if (signedBy) {
    placementConfig.signedBy = signedBy;
  }

  // Build full SDL object
  const sdlObject: SdlObject = {
    version: SDL_DEFAULTS.version,
    services: {
      [name]: service,
    },
    profiles: {
      compute: {
        [name]: {
          resources: computeResources,
        },
      },
      placement: {
        [placement]: placementConfig,
      },
    },
    deployment: {
      [name]: {
        [placement]: {
          profile: name,
          count: replicas,
        },
      },
    },
  };

  // Convert to YAML
  return YAML.dump(sdlObject, {
    indent: 2,
    lineWidth: -1,
    noRefs: true,
    sortKeys: false,
  });
}

/**
 * Creates SDL and returns both YAML and the raw object
 *
 * @param options - Simplified SDL configuration
 * @returns Object containing both YAML string and raw SDL object
 */
export function createSdlWithObject(options: CreateSdlOptions): {
  yaml: string;
  object: SdlObject;
} {
  const yaml = createSdl(options);
  const object = YAML.load(yaml) as SdlObject;
  return { yaml, object };
}
