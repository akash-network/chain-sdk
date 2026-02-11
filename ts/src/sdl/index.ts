/**
 * SDL (Stack Definition Language) module exports
 * Provides functionality for parsing and validating Akash deployment manifests
 *
 * @example
 * ```ts
 * import { SDL } from './sdl';
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
 * const sdl = SDL.fromString(yaml);
 * const manifest = sdl.manifest();
 * ```
 */
export { SDL } from "./SDL/SDL.ts";

export { validateSDL, validationSDLSchema } from "./SDL/validateSDL/validateSDL.ts";
export type { SDLInput } from "./SDL/validateSDL/validateSDLInput.ts";
export type { ValidationError } from "../utils/jsonSchemaValidation.ts";

export * from "./types.ts";
export { SdlValidationError } from "./SDL/SdlValidationError.ts";

/**
 * JSON Schema for validating SDL documents
 * Can be used with any JSON Schema validator (e.g., ajv, jsonschema)
 *
 * @example
 * ```ts
 * import Ajv from 'ajv';
 * import { sdlSchema } from './sdl';
 *
 * const ajv = new Ajv();
 * const validate = ajv.compile(sdlSchema);
 * const isValid = validate(sdlData);
 * ```
 */
export { default as sdlSchema } from "./sdl.schema.json";

export {
  createSdlValidator,
  type JsonSchemaCompiler,
  type JsonSchemaValidateFunction,
  type SdlSchemaValidationError,
  type SdlSchemaValidationResult,
} from "./validateSdl.ts";

/**
 * SDL Builder APIs for creating SDL YAML from JavaScript objects
 *
 * @example Simple API
 * ```ts
 * import { createSdl } from '@akashnetwork/chain-sdk';
 *
 * const yaml = createSdl({
 *   image: "nginx",
 *   cpu: "500m",
 *   memory: "512Mi",
 *   port: 80,
 * });
 * ```
 *
 * @example Full Builder API
 * ```ts
 * import { SdlBuilder } from '@akashnetwork/chain-sdk';
 *
 * const yaml = new SdlBuilder()
 *   .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
 *   .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
 *   .placement("akash", {})
 *   .deploy("web", "akash", { profile: "web", count: 1 })
 *   .pricing("akash", { web: 1000 })
 *   .build();
 * ```
 */
export {
  createSdl,
  createSdlWithObject,
  HTTP_OPTIONS_DEFAULTS,
  SDL_DEFAULTS,
  SdlBuilder,
  SdlBuilderError,
  type BuilderComputeResources,
  type BuilderDeployConfig,
  type BuilderExposeConfig,
  type BuilderPlacementConfig,
  type BuilderServiceConfig,
  type CreateSdlOptions,
  type SdlBuildResult,
  type SdlObject,
  type SimpleGpuConfig,
  type SimplePortConfig,
  type SimpleStorageConfig,
} from "./builder/index.ts";
