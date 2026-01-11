/**
 * Default values for SDL Builder
 */

/**
 * Default SDL configuration values
 */
export const SDL_DEFAULTS = {
  /** Default SDL version */
  version: "2.0",

  /** Default service name for simple API */
  serviceName: "app",

  /** Default CPU units */
  cpu: "500m",

  /** Default memory size */
  memory: "512Mi",

  /** Default storage size */
  storage: "1Gi",

  /** Default replica count */
  replicas: 1,

  /** Default placement name */
  placement: "akash",

  /** Default pricing in uakt */
  pricing: 1000,

  /** Default token denomination */
  denom: "uakt",
} as const;

/**
 * Default HTTP options for expose configurations
 */
export const HTTP_OPTIONS_DEFAULTS = {
  max_body_size: 1048576,
  read_timeout: 60000,
  send_timeout: 60000,
  next_tries: 3,
  next_timeout: 0,
  next_cases: ["error", "timeout"],
} as const;
