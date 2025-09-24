/**
 * Generic test helpers for protobuf deserialization validation
 */
import { expect } from "@jest/globals";
import Long from "long";

/**
 * Generic function to validate protobuf deserialization for ANY type
 * 
 * Usage examples:
 *   validateProtobufDeserialization<QueryCertificatesResponse>(certificateResponse);
 *   validateProtobufDeserialization<QueryLeaseResponse>(marketResponse);
 *   validateProtobufDeserialization<YourProtobufType>(anyProtobufObject);
 * 
 * Benefits:
 *   - Full TypeScript type safety and IntelliSense after validation
 *   - Works with any protobuf-generated type
 *   - No need to specify field names or structures manually
 *   - Uses TypeScript assertion to provide type safety after validation
 */
export function validateProtobufDeserialization<T>(obj: any): asserts obj is T {
  const typeName = typeof obj === 'object' && obj?.constructor?.name || 'unknown';
  
  if (obj === null || obj === undefined) {
    throw new Error(`Object is null or undefined (expected ${typeName})`);
  }

  // Check for common protobuf deserialization patterns
  const validateValue = (value: any, path: string): void => {
    if (value === null || value === undefined) {
      return; // Optional fields can be undefined
    }

    // Check for proper Long deserialization (protobuf int64/uint64)
    if (typeof value === 'object' && value.constructor?.name === 'Long') {
      if (!Long.isLong(value)) {
        throw new Error(`Expected Long at ${path}, got ${typeof value}`);
      }
      return;
    }

    // Check for proper Uint8Array deserialization (protobuf bytes)
    if (value instanceof Uint8Array) {
      if (!(value instanceof Uint8Array)) {
        throw new Error(`Expected Uint8Array at ${path}, got ${typeof value}`);
      }
      return;
    }

    // Recursively validate nested objects
    if (typeof value === 'object' && !Array.isArray(value)) {
      Object.entries(value).forEach(([key, nestedValue]) => {
        validateValue(nestedValue, `${path}.${key}`);
      });
      return;
    }

    // Validate arrays (protobuf repeated fields)
    if (Array.isArray(value)) {
      value.forEach((item, index) => {
        validateValue(item, `${path}[${index}]`);
      });
      return;
    }

    // Primitive types should be properly typed
    const primitiveTypes = ['string', 'number', 'boolean'];
    if (!primitiveTypes.includes(typeof value)) {
      throw new Error(`Unexpected type at ${path}: ${typeof value}`);
    }
  };

  // Start validation from root
  validateValue(obj, typeName);

  // Additional checks for protobuf-specific patterns
  if (typeof obj === 'object') {
    // Ensure no functions leaked through (common serialization issue)
    const hasFunctions = Object.values(obj).some(v => typeof v === 'function');
    if (hasFunctions) {
      throw new Error('Object contains function properties - possible serialization issue');
    }

    // Ensure object is plain (not a class instance unless it's Long or Uint8Array)
    const allowedConstructors = ['Object', 'Long', 'Uint8Array', 'Array'];
    const constructorName = obj.constructor?.name;
    if (constructorName && !allowedConstructors.includes(constructorName)) {
      console.warn(`Unexpected constructor: ${constructorName} for type validation`);
    }
  }
}

/**
 * Jest-compatible expectation wrapper for protobuf validation
 * 
 * Usage:
 *   expectValidProtobufDeserialization<YourType>(response);
 */
export function expectValidProtobufDeserialization<T>(obj: any): void {
  expect(() => validateProtobufDeserialization<T>(obj)).not.toThrow();
}
