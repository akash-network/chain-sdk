/**
 * Types for the SDL Builder APIs
 */

/**
 * GPU configuration for simple API
 */
export interface SimpleGpuConfig {
  /** Number of GPU units */
  units: number;
  /** GPU vendor (nvidia or amd) */
  vendor: "nvidia" | "amd";
  /** GPU model (e.g., "a100", "h100", "rtx4090") */
  model?: string;
  /** GPU RAM (e.g., "40Gi", "80Gi") */
  ram?: string;
  /** GPU interface (pcie or sxm) */
  interface?: "pcie" | "sxm";
}

/**
 * Storage configuration for simple API
 */
export interface SimpleStorageConfig {
  /** Storage size (e.g., "10Gi") */
  size: string;
  /** Volume name (defaults to "default") */
  name?: string;
  /** Whether storage is persistent */
  persistent?: boolean;
  /** Storage class */
  class?: "default" | "beta1" | "beta2" | "beta3" | "ram";
  /** Mount path (required for persistent storage) */
  mount?: string;
}

/**
 * Port exposure configuration for simple API
 */
export interface SimplePortConfig {
  /** Internal port number */
  port: number;
  /** External port (defaults to same as port) */
  as?: number;
  /** Protocol (defaults to TCP) */
  proto?: "tcp" | "udp";
  /** Whether to expose globally (defaults to true) */
  global?: boolean;
  /** Accepted hostnames for HTTP routing */
  accept?: string[];
}

/**
 * Simple API input options for creating SDL with minimal configuration
 */
export interface CreateSdlOptions {
  /** Docker image (required) */
  image: string;

  /** Service name (defaults to "app") */
  name?: string;

  /** CPU units (e.g., "500m", "1", 0.5) - defaults to "500m" */
  cpu?: string | number;

  /** Memory size (e.g., "512Mi", "1Gi") - defaults to "512Mi" */
  memory?: string;

  /** Storage size or configuration - defaults to "1Gi" */
  storage?: string | SimpleStorageConfig | SimpleStorageConfig[];

  /** Single port to expose globally */
  port?: number;

  /** Multiple ports configuration */
  ports?: (number | SimplePortConfig)[];

  /** Environment variables (object or array of "KEY=value" strings) */
  env?: Record<string, string> | string[];

  /** Command to run */
  command?: string[];

  /** Arguments to command */
  args?: string[];

  /** GPU configuration */
  gpu?: SimpleGpuConfig;

  /** Number of replicas (defaults to 1) */
  replicas?: number;

  /** Price amount (defaults to 1000) */
  pricing?: number;

  /** Token denomination (defaults to "uakt") */
  denom?: string;

  /** Placement name (defaults to "akash") */
  placement?: string;

  /** Placement attributes */
  attributes?: Record<string, string>;

  /** Signed by requirements */
  signedBy?: {
    allOf?: string[];
    anyOf?: string[];
  };

  /** Private registry credentials */
  credentials?: {
    host: string;
    username: string;
    password: string;
    email?: string;
  };
}

/**
 * Service configuration for the full builder API
 */
export interface BuilderServiceConfig {
  /** Docker image */
  image: string;
  /** Command to run */
  command?: string[];
  /** Arguments */
  args?: string[];
  /** Environment variables */
  env?: string[];
  /** Port exposures */
  expose?: BuilderExposeConfig[];
  /** Service dependencies */
  depends_on?: string[];
  /** Storage params */
  params?: {
    storage?: Record<
      string,
      {
        mount: string;
        readOnly?: boolean;
      } | null
    >;
  };
  /** Private registry credentials */
  credentials?: {
    host: string;
    username: string;
    password: string;
    email?: string;
  };
}

/**
 * Expose configuration for builder API
 */
export interface BuilderExposeConfig {
  /** Port number */
  port: number;
  /** External port */
  as?: number;
  /** Protocol */
  proto?: "tcp" | "udp" | "TCP" | "UDP";
  /** Whether to expose globally */
  global?: boolean;
  /** Target service for internal exposure */
  service?: string;
  /** Accepted hostnames */
  accept?: string[];
  /** HTTP options */
  http_options?: {
    max_body_size?: number;
    read_timeout?: number;
    send_timeout?: number;
    next_tries?: number;
    next_timeout?: number;
    next_cases?: string[];
  };
  /** IP endpoint name */
  ip?: string;
}

/**
 * Compute resources for builder API
 */
export interface BuilderComputeResources {
  /** CPU units */
  cpu: string | number;
  /** Memory size */
  memory: string;
  /** Storage configuration */
  storage: string | SimpleStorageConfig | SimpleStorageConfig[];
  /** GPU configuration */
  gpu?: SimpleGpuConfig | { units: 0 };
}

/**
 * Placement configuration for builder API
 */
export interface BuilderPlacementConfig {
  /** Provider attributes */
  attributes?: Record<string, string>;
  /** Signed by requirements */
  signedBy?: {
    allOf?: string[];
    anyOf?: string[];
  };
}

/**
 * Deployment configuration for builder API
 */
export interface BuilderDeployConfig {
  /** Compute profile name */
  profile: string;
  /** Number of instances */
  count: number;
}

/**
 * Result of SDL building
 */
export interface SdlBuildResult {
  /** The YAML string */
  yaml: string;
  /** The raw SDL object */
  object: SdlObject;
}

/**
 * Raw SDL object structure
 */
export interface SdlObject {
  version: string;
  services: Record<string, unknown>;
  profiles: {
    compute: Record<string, unknown>;
    placement: Record<string, unknown>;
  };
  deployment: Record<string, unknown>;
  endpoints?: Record<string, unknown>;
}
