import type { ErrorMessages, ValidationError, ValidationFunction } from "../../utils/jsonSchemaValidation.ts";
import { dirname, getErrorLocation, humanizeErrors } from "../../utils/jsonSchemaValidation.ts";
import { INTERCONNECT_GROUP_AUTO, resolveInterconnectGroup } from "../manifest/manifestUtils.ts";
import { castArray, stringToBoolean } from "../utils.ts";
import { schema as validationSDLSchema, type SDLInput, validate as validateSDLInput } from "./validateSDLInput.ts";

export { validationSDLSchema };
export type { SDLInput };

const ERROR_MESSAGES: ErrorMessages = {
  "#/definitions/storageRamClassMustNotBePersistent"(error) {
    return `"ram" storage${getErrorLocation(dirname(error.instancePath))} cannot be persistent`;
  },
  "#/definitions/exposeToWithIpEnforcesGlobal"() {
    return `If an IP is declared, the directive must be declared as global.`;
  },
  // Mirrors the shared schema's `min_window` pattern (`^[1-9][0-9]*(s|m|h)$`).
  // Go stays the lenient layer (`go/sdl/reclamation.go` accepts any `> 0`
  // `time.ParseDuration`), so this is a sanctioned schema-only-stricter rule.
  "#/properties/reclamation/properties/min_window/pattern"() {
    return `Reclamation min_window must be a whole number followed by s, m, or h (e.g. "24h", "30m").`;
  },
};

export function validateSDL(sdl: SDLInput): undefined | ValidationError[] {
  validateSDLInput(sdl);
  const schemaErrors = humanizeErrors((validateSDLInput as ValidationFunction).errors, validationSDLSchema, ERROR_MESSAGES);
  if (schemaErrors.length) return schemaErrors;

  const validator = new SDLValidator(sdl);
  const errors = validator.validate();

  const allErrors = schemaErrors.concat(errors);
  return allErrors.length ? allErrors : undefined;
}

class SDLValidator {
  readonly #endpointsUsed = new Set<string>();
  readonly #portsUsed = new Map<string, string>();
  readonly #teeTypeByPlacement = new Map<string, string>();
  readonly #sdl: SDLInput;
  readonly #errors: ValidationError[] = [];

  constructor(sdl: SDLInput) {
    this.#sdl = sdl;
  }

  validate() {
    if (this.#sdl.services) {
      Object.keys(this.#sdl.services).forEach((serviceName) => {
        this.#validateDeploymentWithRelations(serviceName);
        this.#validateLeaseIP(serviceName);
      });
    }

    this.#validateInterconnect();
    this.#validateEndpoints();
    return this.#errors;
  }

  #validateDeploymentWithRelations(serviceName: string) {
    const deployment = this.#sdl.deployment[serviceName];
    if (!deployment) {
      this.#errors.push({
        message: `Service "${serviceName}" is not defined at "/deployment" section.`,
        instancePath: `/deployment`,
        schemaPath: "#/properties/deployment",
        keyword: "required",
        params: {
          missingProperty: serviceName,
        },
      });
      return;
    }

    Object.keys(this.#sdl.deployment[serviceName]).forEach((deploymentName) => {
      this.#validateDeploymentRelations(serviceName, deploymentName);
      this.#validateServiceStorages(serviceName, deploymentName);
      this.#validateStorages(serviceName, deploymentName);
      this.#validateGPU(serviceName, deploymentName);
      this.#validateTEE(serviceName, deploymentName);
    });
  }

  #validateDeploymentRelations(serviceName: string, deploymentName: string) {
    const serviceDeployment = this.#sdl.deployment?.[serviceName]?.[deploymentName];
    const compute = this.#sdl.profiles?.compute?.[serviceDeployment?.profile];
    const infra = this.#sdl.profiles?.placement?.[deploymentName];

    if (!infra) {
      this.#errors.push({
        message: `The placement "${deploymentName}" is not defined in the "placement" section.`,
        instancePath: `/profiles/placement`,
        schemaPath: "#/properties/profiles/properties/placement",
        keyword: "required",
        params: {
          missingProperty: deploymentName,
        },
      });
    }

    if (infra && !infra.pricing?.[serviceDeployment?.profile]) {
      this.#errors.push({
        message: `The pricing for the "${serviceDeployment?.profile}" profile is not defined in the "${deploymentName}" placement.`,
        instancePath: `/profiles/placement/${deploymentName}/pricing`,
        schemaPath: "#/properties/profiles/properties/placement/additionalProperties/properties/pricing",
        keyword: "required",
        params: {
          missingProperty: serviceDeployment?.profile,
        },
      });
    }

    if (!compute) {
      this.#errors.push({
        message: `The compute requirements for the "${serviceDeployment?.profile}" profile are not defined in the "compute" section.`,
        instancePath: `/profiles/compute`,
        schemaPath: "#/properties/profiles/properties/compute",
        keyword: "required",
        params: {
          missingProperty: serviceDeployment?.profile,
        },
      });
    }
  }

  #validateServiceStorages(serviceName: string, deploymentName: string) {
    const service = this.#sdl.services?.[serviceName];
    const mounts: Record<string, string> = {};
    const serviceDeployment = this.#sdl.deployment[serviceName][deploymentName];
    const compute = this.#sdl.profiles?.compute?.[serviceDeployment.profile];
    const storages = castArray(compute?.resources.storage);

    if (!service?.params?.storage) {
      return;
    }

    Object.entries(service.params.storage).forEach(([storageName, storage]) => {
      if (!storage) {
        this.#errors.push({
          message: `Storage "${storageName}" is not configured.`,
          instancePath: `/services/${serviceName}/params/storage/${storageName}`,
          schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage/additionalProperties",
          keyword: "required",
          params: {
            missingProperty: storageName,
          },
        });
        return;
      }
      const storageNameExists = storages.some(({ name }) => name === storageName);
      if (!storageNameExists) {
        this.#errors.push({
          message: `Service "${serviceName}" references non-existing compute volume "${storageName}".`,
          instancePath: `/profiles/compute/${serviceDeployment.profile}/resources/storage`,
          schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/storage",
          keyword: "required",
          params: {
            missingProperty: storageName,
          },
        });
        return;
      }

      const mount = String(storage.mount);
      const volumeName = mounts[mount];

      if (volumeName && !storage.mount) {
        this.#errors.push({
          message: "Multiple root ephemeral storages are not allowed.",
          instancePath: `/services/${serviceName}/params/storage/${storageName}`,
          schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage",
          keyword: "uniqueItems",
          params: {
            duplicate: volumeName,
          },
        });
      }
      if (volumeName && storage.mount) {
        this.#errors.push({
          message: `Mount "${mount}" already in use by volume "${volumeName}".`,
          instancePath: `/services/${serviceName}/params/storage/${storageName}/mount`,
          schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage/additionalProperties/properties/mount",
          keyword: "uniqueItems",
          params: {
            duplicate: mount,
          },
        });
      }

      mounts[mount] = storageName;
    });
  }

  #validateStorages(serviceName: string, deploymentName: string) {
    const service = this.#sdl.services?.[serviceName];
    const serviceDeployment = this.#sdl.deployment[serviceName][deploymentName];
    const compute = this.#sdl.profiles?.compute?.[serviceDeployment.profile];
    const storages = castArray(compute?.resources.storage);

    storages.forEach((storage) => {
      const persistent = stringToBoolean(storage.attributes?.persistent as string | boolean || false);

      if (persistent && !service?.params?.storage?.[storage.name || ""]?.mount) {
        this.#errors.push({
          message: `Persistent storage "${storage.name || "default"}" requires a mount path in /services/${serviceName}/params/storage/${storage.name || "default"}/mount.`,
          instancePath: `/services/${serviceName}/params/storage/${storage.name || "default"}`,
          schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage/additionalProperties/properties/mount",
          keyword: "required",
          params: {
            missingProperty: "mount",
          },
        });
      }
    });
  }

  #validateGPU(serviceName: string, deploymentName: string) {
    const deployment = this.#sdl.deployment[serviceName];
    const compute = this.#sdl.profiles?.compute?.[deployment[deploymentName]?.profile];
    const gpu = compute?.resources.gpu;
    if (!gpu) return;

    const hasUnits = gpu.units !== undefined && gpu.units !== 0;
    const hasAttributes = typeof gpu.attributes !== "undefined";
    const hasVendor = hasAttributes && typeof gpu.attributes?.vendor !== "undefined";

    const profile = deployment[deploymentName]?.profile;
    const gpuPath = `/profiles/compute/${profile}/resources/gpu`;

    if (!hasUnits && hasAttributes) {
      this.#errors.push({
        message: "GPU must not have attributes if units is 0.",
        instancePath: `${gpuPath}/attributes`,
        schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu/properties/attributes",
        keyword: "additionalProperties",
        params: {
          additionalProperty: "attributes",
        },
      });
    }
    if (hasUnits && !hasAttributes) {
      this.#errors.push({
        message: "GPU must have attributes if units is not 0.",
        instancePath: gpuPath,
        schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu",
        keyword: "required",
        params: {
          missingProperty: "attributes",
        },
      });
    }
    if (hasUnits && !hasVendor) {
      this.#errors.push({
        message: "GPU must specify a vendor if units is not 0.",
        instancePath: `${gpuPath}/attributes`,
        schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu/properties/attributes/properties/vendor",
        keyword: "required",
        params: {
          missingProperty: "vendor",
        },
      });
    }

    // Note: the units==0 + interconnect/interconnect_group case is already rejected by
    // the schema-level rule `gpuAttributesRequireUnitsGt0` — it requires
    // units > 0 whenever any attribute (including interconnect / interconnect_group) is
    // present. So we don't need a duplicate semantic check here. The
    // Go parser (go/sdl/gpu.go) carries its own parse-time check because
    // YAML decoding has no schema layer above it.
  }

  // Mirror the Go-side validateInterconnect in go/sdl/v2.go. The TS
  // SDK enforces the same cross-field rules so tenants using
  // @akashnetwork/chain-sdk get the same fail-fast feedback as Go CLI
  // users (the chain doesn't validate SDL semantics). See
  // docs/sdl-interconnect-spec.md for the spec.
  //
  //   1. Any opt-in (implicit `[]` or explicit `{ group: <name> }`)
  //      requires the placement to require
  //      capabilities/gpu-interconnect="true".
  //   2. The reserved name `auto` may not be written explicitly under
  //      `interconnect: { group: ... }`. (Rule 2 is enforced at the
  //      schema level via the JSON-schema oneOf, but checked here too
  //      for a friendlier error message.)
  //   3. Within one placement, no mixing of implicit and explicit
  //      opt-in forms — either every service uses `[]` or every service
  //      uses a `{ group: ... }`.
  //   4. (Schema-level) `gpu.units == 0` with any interconnect opt-in
  //      is rejected.
  #validateInterconnect() {
    const profiles = this.#sdl.profiles?.compute;
    const placements = this.#sdl.profiles?.placement;
    if (!profiles) return;

    // Rule 2: tenants cannot write `group: auto` under the explicit
    // mapping form. The JSON schema's `oneOf` already rules this out at
    // the structural level when the value is bare, but we surface a
    // friendly error here.
    for (const [profileName, compute] of Object.entries(profiles)) {
      const gpu = compute?.resources?.gpu;
      if (!gpu || gpu.units === 0 || gpu.units === undefined) continue;
      const ic = gpu.attributes?.interconnect;
      if (ic && typeof ic === "object" && !Array.isArray(ic) && (ic as { group?: unknown }).group === INTERCONNECT_GROUP_AUTO) {
        this.#errors.push({
          message: `Compute profile "${profileName}" uses gpu.attributes.interconnect.group="${INTERCONNECT_GROUP_AUTO}", which is reserved (the parser auto-assigns it to "interconnect: []" resources). Pick a different group name.`,
          instancePath: `/profiles/compute/${profileName}/resources/gpu/attributes/interconnect/group`,
          schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu/properties/attributes/properties/interconnect",
          keyword: "const",
          params: { allowedValue: `non-${INTERCONNECT_GROUP_AUTO}` },
        });
      }
    }

    // Rules 1 + 3 are per-placement. Walk every (service, placement) and
    // capture which profile each service uses and whether it opted into
    // interconnect (implicit, explicit, or not at all).
    interface profileUsage {
      profileName: string;
      serviceName: string;
      group: string; // "" if non-interconnect, INTERCONNECT_GROUP_AUTO for implicit, else explicit name
    }
    const usagesByPlacement = new Map<string, profileUsage[]>();

    for (const [serviceName, perPlacement] of Object.entries(this.#sdl.deployment ?? {})) {
      for (const [placementName, svcdepl] of Object.entries(perPlacement ?? {})) {
        const compute = profiles[svcdepl.profile];
        if (!compute) continue;
        const gpu = compute.resources?.gpu;
        const usage: profileUsage = {
          profileName: svcdepl.profile,
          serviceName,
          group: "",
        };
        if (gpu && gpu.units !== 0 && gpu.units !== undefined) {
          usage.group = resolveInterconnectGroup(gpu.attributes?.interconnect);
        }

        const bucket = usagesByPlacement.get(placementName) ?? [];
        bucket.push(usage);
        usagesByPlacement.set(placementName, bucket);

        // Rule 1: any interconnect opt-in must be deployed under a
        // placement whose attributes require capabilities/gpu-interconnect=true.
        if (usage.group !== "") {
          const placement = placements?.[placementName];
          const attrs = (placement?.attributes ?? {}) as Record<string, unknown>;
          const capability = attrs["capabilities/gpu-interconnect"];
          const requiresInterconnect = capability === "true" || capability === true;
          if (!requiresInterconnect) {
            this.#errors.push({
              message: `Service "${serviceName}" uses interconnect profile "${svcdepl.profile}" under placement "${placementName}" but placement does not require capabilities/gpu-interconnect=true.`,
              instancePath: `/profiles/placement/${placementName}/attributes`,
              schemaPath: "#/properties/profiles/properties/placement/additionalProperties/properties/attributes",
              keyword: "required",
              params: { missingProperty: "capabilities/gpu-interconnect" },
            });
          }
        }
      }
    }

    // Rule 3: within one placement, no mixing of implicit (`auto`) and
    // explicit (named) opt-ins. Either form alone is fine; combining
    // them is the failure mode this rule catches.
    for (const [placementName, usages] of usagesByPlacement.entries()) {
      let firstImplicit: string | undefined;
      let firstExplicit: string | undefined;
      for (const u of usages) {
        if (u.group === "") continue;
        if (u.group === INTERCONNECT_GROUP_AUTO) {
          firstImplicit ??= u.profileName;
        } else {
          firstExplicit ??= u.profileName;
        }
      }
      if (firstImplicit !== undefined && firstExplicit !== undefined) {
        this.#errors.push({
          message: `Placement "${placementName}" mixes implicit "interconnect: []" (profile "${firstImplicit}") and explicit "interconnect: { group: ... }" (profile "${firstExplicit}"); use one form across the placement.`,
          instancePath: `/profiles/compute/${firstImplicit}/resources/gpu/attributes/interconnect`,
          schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu/properties/attributes/properties/interconnect",
          keyword: "oneOf",
          params: { passingSchemas: [] },
        });
      }
    }
  }

  // Mirrors Go's TEE rules from `buildGroups` (go/sdl/groupBuilder_v2.go:58-83,
  // groupBuilder_v2_1.go:59-84). Two checks, both on `deploymentName` (the
  // placement-group name):
  //   1. `validateTEEWithGPU` — the `cpu-gpu` type requires GPU resources on the
  //      resolved compute profile (`cpu`/absent need none).
  //   2. `errTEETypeMismatch` — a placementg diff group may carry only one tee type,
  //      since `generateManifest` projects it as a single `tee/type` requirement
  //      attribute. Detecting the conflict here (the validation layer) means
  //      standalone `validateSDL` callers catch it, and `generateManifest` can
  //      project unconditionally.
  // Unsupported tee values are rejected structurally by the input schema (`tee`
  // enum in validateSDLInput.ts) before this runs.
  #validateTEE(serviceName: string, deploymentName: string) {
    const tee = this.#sdl.services?.[serviceName]?.params?.tee;
    if (!tee) return;

    if (tee === "cpu-gpu") {
      const profile = this.#sdl.deployment[serviceName]?.[deploymentName]?.profile;
      const gpu = this.#sdl.profiles?.compute?.[profile]?.resources.gpu;
      const hasGpu = gpu?.units !== undefined && gpu.units !== 0;

      if (!hasGpu) {
        this.#errors.push({
          message: `Service "${serviceName}" tee type requires gpu resources.`,
          instancePath: `/services/${serviceName}/params/tee`,
          schemaPath: "#/properties/services/additionalProperties/properties/params/properties/tee",
          keyword: "required",
          params: {
            missingProperty: "gpu",
          },
        });
      }
    }

    const existing = this.#teeTypeByPlacement.get(deploymentName);
    if (existing === undefined) {
      this.#teeTypeByPlacement.set(deploymentName, tee);
    } else if (existing !== tee) {
      this.#errors.push({
        message: `conflicting tee types in placement group "${deploymentName}": "${existing}" and "${tee}"`,
        instancePath: `/services/${serviceName}/params/tee`,
        schemaPath: "#/properties/services/additionalProperties/properties/params/properties/tee",
        keyword: "tee",
        params: {},
      });
    }
  }

  #validateLeaseIP(serviceName: string) {
    this.#sdl.services?.[serviceName]?.expose?.forEach((expose, exposeIndex) => {
      const proto = expose.proto?.toUpperCase() || "TCP";

      expose.to?.forEach((to, toIndex) => {
        if (to.ip?.length) {
          const toPath = `/services/${serviceName}/expose/${exposeIndex}/to/${toIndex}`;

          if (!to.global) {
            this.#errors.push({
              message: `If an IP is declared, the directive must be declared as global.`,
              instancePath: `${toPath}/global`,
              schemaPath: "#/definitions/exposeToWithIpEnforcesGlobal/then/properties/global/const",
              keyword: "const",
              params: {
                allowedValue: true,
              },
            });
          }
          if (!this.#sdl.endpoints?.[to.ip]) {
            this.#errors.push({
              message: `Unknown endpoint "${to.ip}" for service "${serviceName}". Add it to the "endpoints" section.`,
              instancePath: `/endpoints/${to.ip}`,
              schemaPath: "#/properties/endpoints",
              keyword: "required",
              params: {
                missingProperty: to.ip,
              },
            });
          }

          this.#endpointsUsed.add(to.ip);

          const externalPort = expose.as ?? expose.port;
          const portKey = `${to.ip}-${externalPort}-${proto}`;
          const otherServiceName = this.#portsUsed.get(portKey);

          if (this.#portsUsed.has(portKey)) {
            this.#errors.push({
              message: `IP endpoint "${to.ip}" port ${externalPort} protocol ${proto} already in use by service "${otherServiceName}".`,
              instancePath: `${toPath}/ip`,
              schemaPath: "#/properties/services/additionalProperties/properties/expose/items/properties/to/items",
              keyword: "uniqueItems",
              params: {
                duplicate: portKey,
              },
            });
          }
          this.#portsUsed.set(portKey, serviceName);
        }
      });
    });
  }

  #validateEndpoints() {
    if (!this.#sdl.endpoints) return;

    Object.keys(this.#sdl.endpoints).forEach((endpoint) => {
      if (!this.#endpointsUsed.has(endpoint)) {
        this.#errors.push({
          message: `Endpoint "${endpoint}" declared but never used.`,
          instancePath: `/endpoints/${endpoint}`,
          schemaPath: "#/properties/endpoints",
          keyword: "additionalProperties",
          params: {
            additionalProperty: endpoint,
          },
        });
      }
    });
  }
}
