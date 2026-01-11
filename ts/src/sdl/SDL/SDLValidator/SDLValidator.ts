import { AKT_DENOM, type NetworkId, USDC_IBC_DENOMS } from "../../../network/index.ts";
import { dirname, type ErrorMessages, getErrorLocation, humanizeErrors, type ValidationFunction } from "../../../utils/jsonSchemaValidation.ts";
import { castArray, stringToBoolean } from "../utils.ts";
import { schema as sdlSchema, type SDLInput, validate as validateSDLInput } from "./validateSDLInput.ts";

export type { SDLInput };

const ERROR_MESSAGES: ErrorMessages = {
  "#/definitions/storageAttributesValidation"(error) {
    return `"ram" storage${getErrorLocation(dirname(error.instancePath))} cannot be persistent`;
  },
};

/**
 * @private
 */
export class SDLValidator {
  readonly #endpointsUsed = new Set<string>();
  readonly #portsUsed = new Map<string, string>();
  readonly #sdl: SDLInput;
  readonly #errors: string[] = [];

  static validate(sdl: SDLInput, networkId: NetworkId): undefined | string[] {
    validateSDLInput(sdl);
    const schemaErrors = humanizeErrors((validateSDLInput as ValidationFunction).errors, sdlSchema, ERROR_MESSAGES);

    const validator = new SDLValidator(sdl);
    validator.#validate(networkId);

    const allErrors = schemaErrors.concat(validator.#errors);
    return allErrors.length ? allErrors : undefined;
  }

  private constructor(sdl: SDLInput) {
    this.#sdl = sdl;
  }

  #validate(networkId: NetworkId) {
    if (this.#sdl.services) {
      Object.keys(this.#sdl.services).forEach((serviceName) => {
        this.#validateDeploymentWithRelations(serviceName);
        this.#validateLeaseIP(serviceName);
      });
    }

    this.#validateDenom(networkId);
    this.#validateEndpoints();
  }

  #validateDenom(networkId: NetworkId) {
    if (!this.#sdl.profiles?.placement) return;

    const usdcDenom = USDC_IBC_DENOMS[networkId];
    const denoms = Object.entries(this.#sdl.profiles.placement).map(([placementName, placement]) => {
      if (!placement.pricing) return [];
      return Object.entries(placement.pricing).map(([profile, pricing]) => ({
        path: `/profiles/placement/${placementName}/pricing/${profile}/denom`,
        denom: pricing.denom,
      }));
    }).flat();
    const invalidDenom = denoms.find(({ denom }) => denom !== AKT_DENOM && denom !== usdcDenom);
    if (invalidDenom) {
      this.#errors.push(`Invalid denom: "${invalidDenom?.denom}" at path "${invalidDenom?.path}". Only "uakt" and "${usdcDenom}" are supported.`);
    }
  }

  #validateDeploymentWithRelations(serviceName: string) {
    const deployment = this.#sdl.deployment[serviceName];
    if (!deployment) {
      this.#errors.push(`Service "${serviceName}" is not defined in the "deployment" section.`);
      return;
    }

    Object.keys(this.#sdl.deployment[serviceName]).forEach((deploymentName) => {
      this.#validateDeploymentRelations(serviceName, deploymentName);
      this.#validateServiceStorages(serviceName, deploymentName);
      this.#validateStorages(serviceName, deploymentName);
      this.#validateGPU(serviceName, deploymentName);
    });
  }

  #validateDeploymentRelations(serviceName: string, deploymentName: string) {
    const serviceDeployment = this.#sdl.deployment?.[serviceName]?.[deploymentName];
    const compute = this.#sdl.profiles?.compute?.[serviceDeployment?.profile];
    const infra = this.#sdl.profiles?.placement?.[deploymentName];

    if (!infra) {
      this.#errors.push(`The placement "${deploymentName}" is not defined in the "placement" section.`);
    }

    if (infra && !infra.pricing?.[serviceDeployment?.profile]) {
      this.#errors.push(`The pricing for the "${serviceDeployment?.profile}" profile is not defined in the "${deploymentName}" "placement" definition.`);
    }

    if (!compute) {
      this.#errors.push(`The compute requirements for the "${serviceDeployment?.profile}" profile are not defined in the "compute" section.`);
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
        this.#errors.push(`Storage "${storageName}" is not configured.`);
        return;
      }
      const storageNameExists = storages.some(({ name }) => name === storageName);
      if (!storageNameExists) {
        this.#errors.push(`Service "${serviceName}" references to non-existing compute volume names "${storageName}".`);
        return;
      }

      const mount = String(storage.mount);
      const volumeName = mounts[mount];

      if (volumeName && !storage.mount) {
        this.#errors.push("Multiple root ephemeral storages are not allowed");
      }
      if (volumeName && storage.mount) {
        this.#errors.push(`Mount ${mount} already in use by volume "${volumeName}".`);
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
        this.#errors.push(`/compute/storage/${storage.name} has persistent=true which requires /services/${serviceName}/params/storage/${storage.name} to have "mount" field.`);
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

    if (!hasUnits && hasAttributes) {
      this.#errors.push("GPU must not have attributes if units is 0");
    }
    if (hasUnits && !hasAttributes) {
      this.#errors.push("GPU must have attributes if units is not 0");
    }
    if (hasUnits && !hasVendor) {
      this.#errors.push("GPU must specify a vendor if units is not 0");
    }
  }

  #validateLeaseIP(serviceName: string) {
    this.#sdl.services?.[serviceName]?.expose?.forEach((expose) => {
      const proto = expose.proto?.toUpperCase() || "TCP";

      expose.to?.forEach((to) => {
        if (to.ip?.length) {
          if (!to.global) {
            this.#errors.push(`Error on "${serviceName}", if an IP is declared, the directive must be declared as global.`);
          }
          if (!this.#sdl.endpoints?.[to.ip]) {
            this.#errors.push(`Unknown endpoint "${to.ip}" in service "${serviceName}". Add to the list of endpoints in the "endpoints" section.`);
          }

          this.#endpointsUsed.add(to.ip);

          const portKey = `${to.ip}-${expose.as}-${proto}`;
          const otherServiceName = this.#portsUsed.get(portKey);

          if (this.#portsUsed.has(portKey)) {
            this.#errors.push(`IP endpoint ${to.ip} port: ${expose.port} protocol: ${proto} specified by service ${serviceName} already in use by ${otherServiceName}`);
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
        this.#errors.push(`Endpoint ${endpoint} declared but never used.`);
      }
    });
  }
}
