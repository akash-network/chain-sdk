import fs from "node:fs";
import path from "node:path";

import type { ErrorObject, ValidateFunction } from "ajv";
import { Ajv } from "ajv";
import { describe, expect, it } from "vitest";

import { LegacyDec } from "../../src/encoding/customTypes/LegacyDec.ts";
import { generateManifest } from "../../src/sdl/manifest/generateManifest.ts";
import { manifestToSortedJSON } from "../../src/sdl/manifest/generateManifestVersion.ts";
import type { SDLInput } from "../../src/sdl/validateSDL/validateSDL.ts";
import { yaml } from "../../src/utils/yaml.ts";

const PROJECT_ROOT = path.join(__dirname, "../../..");
const FIXTURES_INPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/input");
const FIXTURES_OUTPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/output-fixtures");
const INPUT_SCHEMA_PATH = path.join(PROJECT_ROOT, "go/sdl/sdl-input.schema.yaml");

describe("SDL Parity Tests", () => {
  describe("v2.0", () => {
    loadFixtures("v2.0").forEach((fixture) => {
      it(fixture.name, () => {
        const { manifest, expectedManifest } = setup(fixture);
        expect(manifest).toEqual(expectedManifest);
      });
    });
  });

  describe("v2.1", () => {
    loadFixtures("v2.1").forEach((fixture) => {
      it(fixture.name, () => {
        const { manifest, expectedManifest } = setup(fixture);
        expect(manifest).toEqual(expectedManifest);
      });
    });
  });

  describe("invalid SDLs rejected", () => {
    const invalidDir = path.join(FIXTURES_INPUT_ROOT, "invalid");

    if (!fs.existsSync(invalidDir)) {
      it("invalid fixtures directory must exist", () => {
        throw new Error(`Invalid fixtures directory not found: ${invalidDir}`);
      });
      return;
    }

    fs.globSync("*.yaml", { cwd: invalidDir }).forEach((filename) => {
      it(filename, () => {
        const fixturePath = path.join(invalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const sdl: SDLInput = yaml.raw(input);
        const result = generateManifest(sdl);
        expect(result.ok).toBe(false);
      });
    });
  });

  function setup(fixture: Fixture) {
    const rawSDL = fs.readFileSync(fixture.inputPath, "utf8");
    const untrustedSDL: SDLInput = yaml.raw(rawSDL);

    validateAgainstSchema("input", untrustedSDL, INPUT_SCHEMA_PATH);

    const result = generateManifest(untrustedSDL);
    if (!result.ok) throw new Error(`generateManifest failed: ${JSON.stringify(result.value)}`);

    const manifest = JSON.parse(manifestToSortedJSON(result.value.groups), normalizeManifestJSON);
    const expectedManifest = JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8"));

    return {
      manifest,
      expectedManifest,
    };
  }
});

function loadFixtures(version: string): Fixture[] {
  const inputVersionDir = path.join(FIXTURES_INPUT_ROOT, version);

  if (!fs.existsSync(inputVersionDir)) {
    throw new Error(`Fixtures directory ${inputVersionDir} does not exist`);
  }

  const entries = fs.readdirSync(inputVersionDir, { withFileTypes: true });

  return entries
    .filter((entry) => entry.isDirectory())
    .map((entry) => {
      const fixtureName = entry.name;
      const inputPath = path.join(inputVersionDir, fixtureName, "input.yaml");
      const manifestPath = path.join(FIXTURES_OUTPUT_ROOT, version, fixtureName, "manifest.json");

      if (!fs.existsSync(manifestPath)) {
        throw new Error(`manifest.json not generated for ${fixtureName} (run: make generate-sdl-fixtures)`);
      }

      return {
        name: fixtureName,
        inputPath,
        manifestPath,
      };
    });
}

function validateAgainstSchema(name: string, data: unknown, schemaPath: string): void {
  const validate = compileSchema(schemaPath);
  const valid = validate(data);

  if (!valid && validate.errors) {
    const errors = validate.errors.map((err: ErrorObject) => {
      const errorPath = err.instancePath || "(root)";
      return `${errorPath}: ${err.message} [${err.keyword}]`;
    });
    throw new Error(`${name} validation failed. Errors: ${JSON.stringify(errors, null, 2)}`);
  }
}

const schemaCache = new Map<string, ValidateFunction>();
const ajv = new Ajv({ allErrors: true, strict: false });

function compileSchema(schemaPath: string): ValidateFunction {
  const cached = schemaCache.get(schemaPath);
  if (cached) {
    return cached;
  }

  const schemaContent = fs.readFileSync(schemaPath, "utf8");
  const schema = yaml.raw(schemaContent);
  const validator = ajv.compile(schema as Record<string, unknown>);
  schemaCache.set(schemaPath, validator);
  return validator;
}

function normalizeManifestJSON(this: unknown, key: string, value: unknown): unknown {
  if (typeof this !== "object" || this === null) return value;

  if (key === "amount" && "denom" in this && this.denom !== undefined) {
    return LegacyDec.encode(value as string);
  }

  if (key === "val") {
    return value ?? "";
  }

  return value;
}

interface Fixture {
  name: string;
  inputPath: string;
  manifestPath: string;
}
