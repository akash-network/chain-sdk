/**
 * SDL Builder module exports
 *
 * Provides simple and advanced APIs for creating SDL YAML configurations.
 *
 * @example Simple API
 * ```ts
 * import { createSdl } from "@akashnetwork/chain-sdk";
 *
 * // Minimal - just an image
 * const yaml = createSdl({ image: "nginx" });
 *
 * // With resources
 * const yaml = createSdl({
 *   image: "nginx:latest",
 *   cpu: "500m",
 *   memory: "512Mi",
 *   port: 80,
 * });
 * ```
 *
 * @example Full Builder API
 * ```ts
 * import { SdlBuilder } from "@akashnetwork/chain-sdk";
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

export { createSdl, createSdlWithObject, SdlBuilderError } from "./createSdl.ts";
export { SdlBuilder } from "./SdlBuilder.ts";
export { SDL_DEFAULTS, HTTP_OPTIONS_DEFAULTS } from "./defaults.ts";
export type {
  BuilderComputeResources,
  BuilderDeployConfig,
  BuilderExposeConfig,
  BuilderPlacementConfig,
  BuilderServiceConfig,
  CreateSdlOptions,
  SdlBuildResult,
  SdlObject,
  SimpleGpuConfig,
  SimplePortConfig,
  SimpleStorageConfig,
} from "./types.ts";
