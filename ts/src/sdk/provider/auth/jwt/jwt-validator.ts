import { base64Decode } from "./base64.ts";
import type { AnyRecord, JwtTokenPayload } from "./types.ts";
import { schema as tokenPayloadSchema, validate as validatePayload } from "./validate-payload.ts";

export interface JwtValidationResult {
  isValid: boolean;
  errors: string[];
  decodedToken?: {
    header: AnyRecord;
    payload: AnyRecord;
    signature: string;
  };
}

export class JwtValidator {
  private compiledSchema = validatePayload as {
    (data: unknown): boolean;
    errors: Array<{
      keywordLocation: string;
      instanceLocation: string;
    }>;
  };

  /**
   * Validate a JWT token against the Akash JWT schema
   * @param token The JWT token to validate
   * @returns Validation result with errors if any
   */
  validateToken(token: string | JwtTokenPayload): JwtValidationResult {
    const result: JwtValidationResult = {
      isValid: false,
      errors: [],
    };

    try {
      // Check for empty or null input
      if (typeof token === "string" && !token.trim()) {
        result.errors.push("Error validating token: Empty token provided");
        return result;
      } else if (typeof token !== "string" && (!token || Object.keys(token).length === 0)) {
        result.errors.push("Error validating token: Empty payload provided");
        return result;
      }

      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      let payload: Record<string, any>;

      if (typeof token === "string") {
        const parts = token.split(".", 3);
        if (parts.length !== 3) {
          result.errors.push("Error validating token: Invalid token format");
          return result;
        }

        const [headerB64, payloadB64, signature] = parts;

        const header = base64Decode(headerB64);
        payload = base64Decode(payloadB64);

        result.decodedToken = {
          header,
          payload,
          signature,
        };

        // Validate header
        if (!header.alg) {
          result.errors.push("Missing required field in header: alg");
          return result;
        }
      } else {
        payload = token;
      }

      // Validate payload with the schema
      let valid = this.compiledSchema(payload);

      if (!valid) {
        result.errors
          = this.compiledSchema.errors?.map((error) => {
            const field = getFieldName(error.instanceLocation);

            if (error.keywordLocation.endsWith("/required")) {
              return `Missing required field: ${field}`;
            }
            if (error.keywordLocation.endsWith("/pattern")) {
              return `Invalid format: ${field} does not match pattern "${getSchemaFieldByPath(error.keywordLocation)}"`;
            }
            if (error.keywordLocation.endsWith("/additionalProperties")) {
              return "Additional properties are not allowed";
            }
            if (error.keywordLocation.endsWith("/type")) {
              return `${field} should be ${getSchemaFieldByPath(error.keywordLocation)}`;
            }
            if (error.keywordLocation.endsWith("/enum")) {
              return `${field} should be one of: ${getSchemaFieldByPath<string[]>(error.keywordLocation).join(", ")}`;
            }
            return `${field}: does not satisfy ${getSchemaFieldByPath(error.keywordLocation)}`;
          }) || [];
      }

      // Additional validation for granular access
      if (payload.leases?.access === "granular") {
        if (!payload.leases?.permissions) {
          result.errors.push("Missing required field: permissions");
          valid = false;
        } else {
          // Check for duplicate providers
          const providers = new Set<string>();
          for (const perm of payload.leases.permissions) {
            if (providers.has(perm.provider)) {
              result.errors.push("Duplicate provider in permissions");
              valid = false;
              break;
            }
            providers.add(perm.provider);
          }

          for (const perm of payload.leases.permissions) {
            // Validate access type specific rules
            if (perm.access === "scoped") {
              if (!perm.scope) {
                result.errors.push("Missing required field: scope for scoped access");
                valid = false;
              } else if (perm.deployments) {
                result.errors.push("Deployments not allowed for scoped access");
                valid = false;
              }
            } else if (perm.access === "granular") {
              if (!perm.deployments) {
                result.errors.push("Missing required field: deployments for granular access");
                valid = false;
              } else if (perm.scope) {
                result.errors.push("Scope not allowed for granular access");
                valid = false;
              }
            }

            // Check for duplicate services and validate deployment dependencies
            if (perm.deployments) {
              for (const deployment of perm.deployments) {
                if (!this.validateDeployment(deployment, result)) {
                  valid = false;
                }
              }
            }
          }
        }
      } else if (payload.leases?.access === "full" && payload.leases?.permissions) {
        result.errors.push("Permissions not allowed for full access");
        valid = false;
      }

      result.isValid = result.errors.length === 0;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : String(error);
      result.errors.push(`Error during JWT validation: ${errorMessage}`);
    }

    return result;
  }

  /**
   * Validates deployment structure and dependencies
   * @param deployment The deployment to validate
   * @param result The validation result to update
   * @returns Whether the validation passed
   */
  private validateDeployment(deployment: AnyRecord, result: JwtValidationResult): boolean {
    let valid = true;

    // Validate deployment dependencies
    if (deployment.gseq && !deployment.dseq) {
      result.errors.push("gseq requires dseq");
      valid = false;
    }
    if (deployment.oseq && (!deployment.dseq || !deployment.gseq)) {
      result.errors.push("oseq requires dseq and gseq");
      valid = false;
    }
    if (deployment.dseq && !deployment.services) {
      result.errors.push("services required when dseq is present");
      valid = false;
    }
    if (deployment.services && !deployment.dseq) {
      result.errors.push("services requires dseq");
      valid = false;
    }

    return valid;
  }
}

function getFieldName(instanceLocation: string): string {
  const lastPart = basename(instanceLocation);
  return Number.isNaN(Number(lastPart)) ? lastPart : `${basename(dirname(instanceLocation))} ${lastPart}`;
}

function basename(path: string): string {
  const lastPartIndex = path.lastIndexOf("/");
  if (lastPartIndex === -1) return path;
  return path.slice(lastPartIndex + 1);
}

function dirname(path: string): string {
  const lastPartIndex = path.lastIndexOf("/");
  if (lastPartIndex === -1) return path;
  return path.slice(0, lastPartIndex);
}

function getSchemaFieldByPath<T = string>(keywordLocation: string): T {
  return keywordLocation.split("/").slice(1)
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    .reduce<T>((schema, key) => (schema as Record<string, any>)[key], tokenPayloadSchema as T);
}
