import { load } from "js-yaml";

/**
 * @example
 * ```ts
 * const version = "2.1";
 * const expose = [{ port: 80, as: 80, to: [{ global: true }] }];
 * const pricing = { web: { denom: "uakt", amount: 1000 } };
 *
 * const sdl: SDLInput = yaml`
 * version: ${version}
 * services:
 *   web:
 *     image: nginx
 *     expose: ${expose}
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
 *       pricing: ${pricing}
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
    const value = args[i];
    const intermediateResult = acc + chunk;
    if (value === undefined) return intermediateResult;

    return intermediateResult + JSON.stringify(value);
  }, "");

  return load(str) as T;
}

/**
 * Use this function to parse YAML template defined in external resource (e.g., file, http response, etc.).
 * Prefer `yaml` function for inline YAML templates defined in code, as it provides better ergonomics.
 */
yaml.template = function yamlTemplate<T>(template: string, vars?: Record<string, unknown>): T {
  const finalYaml = template.replace(/\$\{(\w+)\}/g, (_, varName) => {
    const value = vars?.[varName];
    if (value === undefined) {
      throw new ReferenceError(`Variable "${varName}" is not provided to yaml template`);
    }
    return JSON.stringify(value);
  });
  return load(finalYaml) as T;
};
