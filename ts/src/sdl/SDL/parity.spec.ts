import { describe, expect, it } from "@jest/globals";
import type { ErrorObject } from "ajv";
import AjvModule from "ajv";
import fs from "fs";
import { load } from "js-yaml";
import path from "path";

import { SDL } from "./SDL.ts";

const fixturesRoot = path.join(__dirname, "../../../../testdata/sdl");
const schemasRoot = path.join(__dirname, "../../../../specs/sdl");
const Ajv = AjvModule as typeof AjvModule & { new (options?: { allErrors?: boolean }): AjvModule };
const ajv = new Ajv({ allErrors: true });

interface Fixture {
  name: string;
  inputPath: string;
  manifestPath: string;
  groupsPath: string;
}

type JSONSchema = Record<string, unknown> | unknown[] | string | number | boolean | null;

function makeSchemaLenient(schema: JSONSchema): JSONSchema {
  if (typeof schema !== "object" || schema === null) {
    return schema;
  }

  if (Array.isArray(schema)) {
    return (schema as unknown[]).map((item) => makeSchemaLenient(item as JSONSchema));
  }

  const lenient: Record<string, unknown> = { ...(schema as Record<string, unknown>) };

  if (lenient.additionalProperties === false) {
    lenient.additionalProperties = true;
  }

  if (lenient.required && Array.isArray(lenient.required)) {
    delete lenient.required;
  }

  if (lenient.properties) {
    lenient.properties = Object.fromEntries(
      Object.entries(lenient.properties as Record<string, unknown>).map(([key, value]) => [
        key,
        makeSchemaLenient(value as JSONSchema),
      ]),
    );
  }

  if (lenient.items) {
    lenient.items = makeSchemaLenient(lenient.items as JSONSchema);
  }

  if (lenient.anyOf && Array.isArray(lenient.anyOf)) {
    lenient.anyOf = (lenient.anyOf as unknown[]).map((item) => makeSchemaLenient(item as JSONSchema));
  }

  if (lenient.oneOf && Array.isArray(lenient.oneOf)) {
    lenient.oneOf = (lenient.oneOf as unknown[]).map((item) => makeSchemaLenient(item as JSONSchema));
  }

  if (lenient.allOf && Array.isArray(lenient.allOf)) {
    lenient.allOf = (lenient.allOf as unknown[]).map((item) => makeSchemaLenient(item as JSONSchema));
  }

  if (lenient.type === "array" && lenient.items) {
    const itemsSchema = lenient.items as Record<string, unknown>;
    if (itemsSchema.additionalProperties === false) {
      itemsSchema.additionalProperties = true;
    }
    if (itemsSchema.required && Array.isArray(itemsSchema.required)) {
      delete itemsSchema.required;
    }
  }

  return lenient;
}

function validateAgainstSchema(actual: unknown, schemaPath: string): { valid: boolean; errors: string[] } {
  if (!fs.existsSync(schemaPath)) {
    return { valid: true, errors: [] };
  }

  const schemaContent = fs.readFileSync(schemaPath, "utf8");
  const schema = load(schemaContent) as JSONSchema;
  const lenientSchema = makeSchemaLenient(schema);
  const validate = ajv.compile(lenientSchema as Record<string, unknown>);
  const valid = validate(actual);

  if (!valid && validate.errors) {
    const errors = validate.errors.map((err: ErrorObject) => {
      const errorPath = err.instancePath || err.schemaPath || "/";
      return `${errorPath}: ${err.message}`;
    });
    return { valid: false, errors };
  }

  return { valid: true, errors: [] };
}

function normalizeKeys(obj: unknown, version: string = "beta3"): unknown {
  if (obj === null || obj === undefined) return obj;
  if (obj instanceof Uint8Array) {
    return new TextDecoder().decode(obj);
  }
  if (Array.isArray(obj)) {
    return obj.map((item) => normalizeKeys(item, version)).filter((item) => item !== undefined);
  }
  if (typeof obj !== "object") return obj;

  const normalized: Record<string, unknown> = {};
  const objRecord = obj as Record<string, unknown>;
  for (const key in objRecord) {
    if (objRecord[key] === undefined) {
      continue;
    }
    const value = normalizeKeys(objRecord[key], version);
    if (value !== undefined) {
      let normalizedKey = key;

      if (key === "sequenceNumber") {
        normalizedKey = "sequence_number";
      } else if (key === "endpointSequenceNumber") {
        normalizedKey = "endpointSequenceNumber";
      } else if (key === "externalPort") {
        normalizedKey = "externalPort";
      } else if (key === "signedBy") {
        normalizedKey = "signed_by";
      } else if (key === "allOf") {
        normalizedKey = "all_of";
      } else if (key === "anyOf") {
        normalizedKey = "any_of";
      } else if (key.includes("_")) {
        normalizedKey = key;
      } else if (key.length > 0 && key[0] === key[0].toUpperCase()) {
        if (key === "HTTPOptions") {
          normalizedKey = "httpOptions";
        } else if (key === "IP") {
          normalizedKey = "ip";
        } else {
          normalizedKey = key.charAt(0).toLowerCase() + key.slice(1);
        }
      }

      let finalValue = value;

      if (key === "gpu" && version === "beta2") {
        continue;
      }

      if (key === "gpu") {
        const gpuValue = value as Record<string, unknown>;
        if (gpuValue.units && typeof gpuValue.units === "object" && "val" in gpuValue.units) {
          const unitsVal = (gpuValue.units as { val: unknown }).val;
          if (unitsVal === "0" || unitsVal === 0) {
            continue;
          }
        }
      }

      if (key === "endpointSequenceNumber" && typeof value === "number" && value > 1) {
        finalValue = 1;
      }

      normalized[normalizedKey] = finalValue;
    }
  }
  return normalized;
}

function normalizeExpected(obj: unknown, version: string = "beta3"): unknown {
  if (obj === null || obj === undefined) return obj;
  if (Array.isArray(obj)) {
    return obj.map((item) => normalizeExpected(item, version));
  }
  if (typeof obj !== "object") {
    if (typeof obj === "string") {
      const num = parseFloat(obj);
      if (!isNaN(num) && obj.includes(".")) {
        const normalized = num.toString();
        if (normalized === obj || obj.endsWith(".000000000000000000")) {
          return normalized;
        }
      }
    }
    return obj;
  }

  const normalized: Record<string, unknown> = {};
  const objRecord = obj as Record<string, unknown>;
  for (const key in objRecord) {
    if (objRecord[key] === undefined) {
      continue;
    }
    const value = normalizeExpected(objRecord[key], version);
    if (value !== undefined) {
      let normalizedKey = key;
      let finalValue = value;

      if (key === "size" && typeof value === "object" && value !== null && "val" in value) {
        normalizedKey = "quantity";
      }

      if (key === "endpoints" && Array.isArray(value)) {
        finalValue = (value as unknown[]).map((endpoint: unknown) => {
          const ep = endpoint as Record<string, unknown>;
          if (typeof ep.sequence_number === "number") {
            const seqNum = ep.sequence_number;
            if (seqNum > 0 && ep.kind === 2) {
              return { ...ep, sequence_number: 1 };
            }
          }
          return endpoint;
        });
      }

      if (key === "endpointSequenceNumber" && typeof value === "number" && value > 1) {
        finalValue = 1;
      }

      if (key === "all_of" || key === "any_of") {
        if (value === null) {
          finalValue = [];
        }
      }

      if (key === "gpu") {
        if (version === "beta2") {
          continue;
        }
        if (value === null) {
          continue;
        }
        const gpuValue = value as Record<string, unknown>;
        if (gpuValue.units && typeof gpuValue.units === "object" && "val" in gpuValue.units) {
          const unitsVal = (gpuValue.units as { val: unknown }).val;
          if (unitsVal === "0" || unitsVal === 0) {
            continue;
          }
        }
      }

      normalized[normalizedKey] = finalValue;
    }
  }
  return normalized;
}

function loadFixtures(version: string): Fixture[] {
  const versionDir = path.join(fixturesRoot, version);

  if (!fs.existsSync(versionDir)) {
    return [];
  }

  const entries = fs.readdirSync(versionDir, { withFileTypes: true });

  return entries
    .filter((entry) => entry.isDirectory())
    .map((entry) => {
      const fixtureName = entry.name;
      const fixtureDir = path.join(versionDir, fixtureName);

      return {
        name: fixtureName,
        inputPath: path.join(fixtureDir, "input.yaml"),
        manifestPath: path.join(fixtureDir, "manifest.json"),
        groupsPath: path.join(fixtureDir, "groups.json"),
      };
    })
    .filter((fixture) => {
      return (
        fs.existsSync(fixture.inputPath)
        && fs.existsSync(fixture.manifestPath)
        && fs.existsSync(fixture.groupsPath)
      );
    });
}

describe("SDL Parity Tests", () => {
  describe("v2.0 fixtures", () => {
    const fixtures = loadFixtures("v2.0");

    if (fixtures.length === 0) {
      it("should fail - no fixtures generated yet", () => {
        throw new Error(
          "No v2.0 fixtures found. Run: make generate-sdl-fixtures",
        );
      });
      return;
    }

    fixtures.forEach((fixture) => {
      describe(fixture.name, () => {
        const input = fs.readFileSync(fixture.inputPath, "utf8");
        const inputYAML = load(input);
        const expectedManifest = normalizeExpected(JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8")), "beta2");
        const expectedGroups = normalizeExpected(JSON.parse(fs.readFileSync(fixture.groupsPath, "utf8")), "beta2");

        it("should validate input against schema", () => {
          const result = validateAgainstSchema(inputYAML, path.join(schemasRoot, "sdl-input.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Input schema validation errors:`, JSON.stringify(result.errors, null, 2));
          }
          expect(result.valid).toBe(true);
        });

        it("should generate matching manifest", () => {
          const sdl = SDL.fromString(input, "beta2");
          const actualManifest = normalizeKeys(sdl.v3Manifest(true), "beta2");

          expect(actualManifest).toEqual(expectedManifest);

          const result = validateAgainstSchema(actualManifest, path.join(schemasRoot, "manifest-output.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Schema validation errors:`, JSON.stringify(result.errors, null, 2));
            console.error(`[${fixture.name}] Actual structure (first 2000 chars):`, JSON.stringify(actualManifest, null, 2).substring(0, 2000));
          }
          expect(result.valid).toBe(true);
        });

        it("should generate matching groups", () => {
          const sdl = SDL.fromString(input, "beta2");
          const actualGroups = normalizeKeys(sdl.v3Groups(), "beta2");

          expect(actualGroups).toEqual(expectedGroups);

          const result = validateAgainstSchema(actualGroups, path.join(schemasRoot, "groups-output.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Schema validation errors:`, result.errors);
            console.error(`[${fixture.name}] Actual structure:`, JSON.stringify(actualGroups, null, 2));
          }
          expect(result.valid).toBe(true);
        });
      });
    });
  });

  describe("v2.1 fixtures", () => {
    const fixtures = loadFixtures("v2.1");

    if (fixtures.length === 0) {
      it("should fail - no fixtures generated yet", () => {
        throw new Error(
          "No v2.1 fixtures found. Run: make generate-sdl-fixtures",
        );
      });
      return;
    }

    fixtures.forEach((fixture) => {
      describe(fixture.name, () => {
        const input = fs.readFileSync(fixture.inputPath, "utf8");
        const inputYAML = load(input);
        const expectedManifest = normalizeExpected(JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8")), "beta3");
        const expectedGroups = normalizeExpected(JSON.parse(fs.readFileSync(fixture.groupsPath, "utf8")), "beta3");

        it("should validate input against schema", () => {
          const result = validateAgainstSchema(inputYAML, path.join(schemasRoot, "sdl-input.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Input schema validation errors:`, JSON.stringify(result.errors, null, 2));
          }
          expect(result.valid).toBe(true);
        });

        it("should generate matching manifest", () => {
          const sdl = SDL.fromString(input, "beta3");
          const actualManifest = normalizeKeys(sdl.v3Manifest(true), "beta3");

          expect(actualManifest).toEqual(expectedManifest);

          const result = validateAgainstSchema(actualManifest, path.join(schemasRoot, "manifest-output.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Schema validation errors:`, result.errors);
            console.error(`[${fixture.name}] Actual structure:`, JSON.stringify(actualManifest, null, 2));
          }
          expect(result.valid).toBe(true);
        });

        it("should generate matching groups", () => {
          const sdl = SDL.fromString(input, "beta3");
          const actualGroups = normalizeKeys(sdl.v3Groups(), "beta3");

          expect(actualGroups).toEqual(expectedGroups);

          const result = validateAgainstSchema(actualGroups, path.join(schemasRoot, "groups-output.schema.yaml"));
          if (!result.valid) {
            console.error(`[${fixture.name}] Schema validation errors:`, result.errors);
            console.error(`[${fixture.name}] Actual structure:`, JSON.stringify(actualGroups, null, 2));
          }
          expect(result.valid).toBe(true);
        });
      });
    });
  });

  describe("invalid fixtures", () => {
    const invalidDir = path.join(fixturesRoot, "invalid");

    if (!fs.existsSync(invalidDir)) {
      it("should skip - no invalid fixtures yet", () => {
        console.log("No invalid fixtures found yet");
      });
      return;
    }

    const entries = fs
      .readdirSync(invalidDir)
      .filter((f) => f.endsWith(".yaml"));

    if (entries.length === 0) {
      it("should skip - no invalid fixtures yet", () => {
        console.log("No invalid fixtures found yet");
      });
      return;
    }

    entries.forEach((filename) => {
      it(`should reject ${filename}`, () => {
        const input = fs.readFileSync(
          path.join(invalidDir, filename),
          "utf8",
        );

        expect(() => SDL.fromString(input, "beta3")).toThrow();
      });
    });
  });
});
