import { load } from "js-yaml";

/**
 * @example
 * ```ts
 * const sdl = yaml<SDLInput>`
 * # yaml-language-server: $schema=https://console.akash.network/sdl-schema.yaml
 *
 * version: "2.0"
 * services:
 *   web:
 *     image: nginx
 *     expose:
 *       - port: 80
 *         as: 80
 *         to:
 *           - global: true
 * profiles:
 *   compute:
 *     web:
 *       resources:
 *         cpu:
 *           units: 0.5
 *         memory:
 *           size: 512Mi
 *         storage:
 *           size: 1Gi
 *   placement:
 *     dcloud:
 *       pricing:
 *         web:
 *           denom: uakt
 *           amount: 1000
 * deployment:
 *   web:
 *     dcloud:
 *       profile: web
 *       count: 1
 * `;
 * ```
 */
export function yaml<T>(chunks: TemplateStringsArray, ...args: unknown[]): T {
  const str = chunks.reduce((acc, chunk, i) => {
    return acc + chunk + (args[i] !== undefined ? String(args[i]) : "");
  }, "");
  return load(str) as T;
}
