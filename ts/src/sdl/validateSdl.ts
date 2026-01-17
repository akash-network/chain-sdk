/**
 * SDL Schema Validation Utilities
 *
 * Provides types and utilities for validating SDL documents against the JSON Schema.
 * The schema can be used with any JSON Schema Draft-07 compatible validator.
 *
 * @example
 * Using with Ajv (recommended):
 * ```ts
 * import Ajv from 'ajv';
 * import { sdlSchema, createSdlValidator } from '@akashnetwork/chain-sdk';
 *
 * const ajv = new Ajv();
 * const validate = createSdlValidator(ajv, sdlSchema);
 *
 * const result = validate(myYamlParsedObject);
 * if (!result.valid) {
 *   console.error('Validation errors:', result.errors);
 * }
 * ```
 *
 * @example
 * Manual usage:
 * ```ts
 * import Ajv from 'ajv';
 * import { sdlSchema } from '@akashnetwork/chain-sdk';
 *
 * const ajv = new Ajv();
 * const validate = ajv.compile(sdlSchema);
 * const isValid = validate(myData);
 * if (!isValid) {
 *   console.error(validate.errors);
 * }
 * ```
 */

/**
 * Validation error from a JSON Schema validator
 */
export interface SdlSchemaValidationError {
  /** JSON pointer to the location of the error */
  instancePath?: string;
  /** Keyword that failed validation */
  keyword?: string;
  /** Human-readable error message */
  message?: string;
  /** Additional parameters for the error */
  params?: Record<string, unknown>;
  /** The path in schema that failed */
  schemaPath?: string;
}

/**
 * Result of SDL schema validation
 */
export interface SdlSchemaValidationResult {
  /** Whether the SDL document is valid */
  valid: boolean;
  /** Validation errors if not valid */
  errors: SdlSchemaValidationError[];
}

/**
 * Interface for a compiled JSON Schema validator function
 * Compatible with Ajv and similar libraries
 */
export interface JsonSchemaValidateFunction {
  (data: unknown): boolean;
  errors?: SdlSchemaValidationError[] | null;
}

/**
 * Interface for a JSON Schema compiler (like Ajv instance)
 */
export interface JsonSchemaCompiler {
  compile(schema: object): JsonSchemaValidateFunction;
}

/**
 * Creates an SDL validator function using the provided JSON Schema compiler and schema
 *
 * @param compiler - A JSON Schema compiler instance (e.g., Ajv instance)
 * @param schema - The SDL JSON Schema object
 * @returns A validation function that returns a structured result
 *
 * @example
 * ```ts
 * import Ajv from 'ajv';
 * import { createSdlValidator, sdlSchema } from '@akashnetwork/chain-sdk';
 *
 * const ajv = new Ajv({ allErrors: true });
 * const validate = createSdlValidator(ajv, sdlSchema);
 *
 * const sdlData = YAML.parse(yamlString);
 * const result = validate(sdlData);
 *
 * if (result.valid) {
 *   console.log('SDL is valid!');
 * } else {
 *   for (const error of result.errors) {
 *     console.error(`${error.instancePath}: ${error.message}`);
 *   }
 * }
 * ```
 */
export function createSdlValidator(
  compiler: JsonSchemaCompiler,
  schema: object,
): (data: unknown) => SdlSchemaValidationResult {
  const validate = compiler.compile(schema);

  return (data: unknown): SdlSchemaValidationResult => {
    const valid = validate(data);

    return {
      valid,
      errors: valid
        ? []
        : (validate.errors ?? []).map((err) => ({
            instancePath: err.instancePath,
            keyword: err.keyword,
            message: err.message,
            params: err.params,
            schemaPath: err.schemaPath,
          })),
    };
  };
}
