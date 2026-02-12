/**
 * @deprecated will be removed in favor of `Manifest` type
 */

/* eslint-disable @typescript-eslint/no-explicit-any */
export type v2Manifest = v2Group[];

/** @deprecated Use `Manifest` type instead. */
export type v3Manifest = v3Group[];

/** @deprecated Use `Manifest` type instead. */
export type v3Group = {
  name: string;
  services: v3ManifestService[];
};

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2Group = {
  Name: string;
  Services: v2ManifestService[];
};

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2ManifestService = {
  Name: string;
  Image: string;
  Command: string[] | null;
  Args: string[] | null;
  Env: string[] | null;
  Resources: ResourceUnits;
  Count: number;
  Expose: v2ServiceExpose[];
  params?: v2ManifestServiceParams;
};

/** @deprecated Use `Manifest` type instead. */
export type v3ManifestService = {
  name: string;
  image: string;
  command: string[] | null;
  args: string[] | null;
  env: string[] | null;
  resources: ResourceUnits;
  count: number;
  expose: v3ServiceExpose[];
  params?: v3ManifestServiceParams | null;
  credentials: v2ServiceImageCredentials | null;
};

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2ServiceExposeHttpOptions = {
  MaxBodySize: number;
  ReadTimeout: number;
  SendTimeout: number;
  NextTries: number;
  NextTimeout: number;
  NextCases: string[];
};

/** @deprecated Use `Manifest[number]['services'][number]['expose'][number]['httpOptions']` type instead. */
export type v3ServiceExposeHttpOptions = {
  maxBodySize: number;
  readTimeout: number;
  sendTimeout: number;
  nextTries: number;
  nextTimeout: number;
  nextCases: string[];
};

/** @deprecated will be removed in favor of `Manifest` type. */
export type ResourceUnits = Record<string, any>;

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2ServiceExpose = {
  Port: number;
  ExternalPort: number;
  Proto: string;
  Service: any;
  Global: boolean;
  Hosts: any;
  HTTPOptions: v2ServiceExposeHttpOptions;
  IP: string;
  EndpointSequenceNumber: number;
};

/** @deprecated Use `Manifest[number]['services'][number]['expose'][number]` type instead. */
export type v3ServiceExpose = {
  port: number;
  externalPort: number;
  proto: string;
  service: any;
  global: boolean;
  hosts: any;
  httpOptions: v3ServiceExposeHttpOptions;
  ip: string;
  endpointSequenceNumber: number;
};

type v2ServicePermissionsScope = "deployment" | "logs";

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2ServicePermissions = {
  read?: v2ServicePermissionsScope[];
};

/** @deprecated will be removed in favor of `Manifest` type. */
export type v2ManifestServiceParams = {
  Storage: v2ServiceStorageParams[];
  Permissions?: v2ServicePermissions;
};

/** @deprecated Use `Manifest[number]['services'][number]['params']` type instead. */
export type v3ManifestServiceParams = {
  storage: v2ServiceStorageParams[] | null;
  permissions?: v2ServicePermissions;
};

/** @deprecated will be removed in favor of `Manifest` type */
export type v2Sdl = {
  services: Record<string, v2Service>;
  profiles: v2Profiles;
  deployment: Record<string, v2Deployment>;
  endpoints: Record<string, v2Endpoint>;
};

/** @deprecated Use `SDLInput` type instead. */
export type v3Sdl = {
  services: Record<string, v2Service>;
  profiles: v3Profiles;
  deployment: Record<string, v2Deployment>;
  endpoints: Record<string, v2Endpoint>;
};

/** @deprecated will be removed. */
export type v2Endpoint = {
  kind: string;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ExposeTo = {
  service?: string;
  global?: boolean;
  http_options: v2HTTPOptions;
  ip: string;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2HTTPOptions = {
  max_body_size: number;
  read_timeout: number;
  send_timeout: number;
  next_tries: number;
  next_timeout: number;
  next_cases: string[];
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Accept = {
  items?: string[];
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Expose = {
  port: number;
  as: number;
  proto?: string;
  to?: v2ExposeTo[];
  accept: v2Accept;
  http_options: v2HTTPOptions;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Dependency = {
  service: string;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ServiceStorageParams = {
  name: string;
  mount: string;
  readOnly: boolean;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ServiceParams = {
  storage?: Record<string, v2ServiceStorageParams>;
  permissions?: v2ServicePermissions;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ServiceImageCredentials = {
  host: string;
  email?: string;
  username: string;
  password: string;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Service = {
  image: string;
  command: string[] | null;
  args: string[] | null;
  env: string[] | null;
  expose: v2Expose[];
  dependencies?: v2Dependency[];
  params?: v2ServiceParams;
  credentials?: v2ServiceImageCredentials;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ServiceDeployment = {
  profile: string;
  count: number;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Deployment = Record<string, v2ServiceDeployment>;

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2CPUAttributes = Record<string, any>;

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ResourceCPU = {
  units: number | string;
  attributes?: v2CPUAttributes;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ResourceMemory = {
  size: string;
  attributes?: Record<string, any>;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3GPUAttributes = {
  vendor: {
    [vendor: string]: Array<{ model: string; ram?: string; interface?: string }>;
  };
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3ResourceGPU = {
  units: number | string;
  attributes?: v3GPUAttributes;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2StorageAttributes = Record<string, any>;

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ResourceStorage = {
  name: string;
  size: string;
  attributes: v2StorageAttributes;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ResourceStorageArray = v2ResourceStorage[];

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ComputeResources = {
  cpu: v2ResourceCPU;
  memory: v2ResourceMemory;
  storage: v2ResourceStorageArray | v2ResourceStorage;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3ComputeResources = {
  cpu: v2ResourceCPU;
  memory: v2ResourceMemory;
  storage: v2ResourceStorageArray | v2ResourceStorage;
  gpu: v3ResourceGPU;
  id: number;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ProfileCompute = {
  resources: v2ComputeResources;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3ProfileCompute = {
  resources: v3ComputeResources;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2PlacementAttributes = Attributes;

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Coin = {
  denom: string;
  value: number;
  amount: number;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2PlacementPricing = Record<string, v2Coin>;

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type SignedBy = {
  allOf: string[];
  anyOf: string[];
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2ProfilePlacement = {
  attributes: v2PlacementAttributes;
  signedBy: SignedBy;
  pricing: v2PlacementPricing;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v2Profiles = {
  compute: Record<string, v2ProfileCompute>;
  placement: Record<string, v2ProfilePlacement>;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3Profiles = {
  compute: Record<string, v3ProfileCompute>;
  placement: Record<string, v2ProfilePlacement>;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type Attribute = {
  key: string;
  value: string;
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type v3DeploymentGroup = {
  name: string;
  
  resources: Array<{
    resource: v3ComputeResources;
    price: number;
    count: number;
    endpoints: Array<{ kind: number; sequence_number: number }>;
  }>;
  requirements: {
    attributes: Array<Attribute>;
    signedBy: {
      allOf: string[];
      anyOf: string[];
    };
  };
};

/** @deprecated will be removed in favor of `SDLInput` and `Manifest` types. */
export type Attributes = Attribute[];
