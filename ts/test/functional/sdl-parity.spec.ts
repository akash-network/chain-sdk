import fs from "node:fs";
import path from "node:path";

import { describe, expect, it } from "@jest/globals";
import type { ErrorObject, ValidateFunction } from "ajv";
import { Ajv } from "ajv";

import { LegacyDec } from "../../src/encoding/customTypes/LegacyDec.ts";
import { generateManifest } from "../../src/sdl/manifest/generateManifest.ts";
import { manifestToSortedJSON } from "../../src/sdl/manifest/generateManifestVersion.ts";
import type { SDLInput } from "../../src/sdl/validateSDL/validateSDL.ts";
import { yaml } from "../../src/utils/yaml.ts";

const PROJECT_ROOT = path.join(__dirname, "../../..");
const FIXTURES_INPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/input");
const FIXTURES_OUTPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/output-fixtures");
const INPUT_SCHEMA_PATH = path.join(PROJECT_ROOT, "go/sdl/sdl-input.schema.yaml");
const MANIFEST_OUTPUT_SCHEMA_PATH = path.join(PROJECT_ROOT, "specs/sdl/manifest-output.schema.yaml");
const GROUPS_OUTPUT_SCHEMA_PATH = path.join(PROJECT_ROOT, "specs/sdl/groups-output.schema.yaml");

describe("SDL Parity Tests", () => {
  describe("v2.0", () => {
    loadFixtures("v2.0").forEach((fixture) => {
      it(fixture.name, () => {
        const { manifest, expectedManifest, groupSpecs, expectedGroupSpecs } = setup(fixture);
        expect(manifest).toEqual(expectedManifest);
        expect(groupSpecs).toEqual(expectedGroupSpecs);
      });
    });
  });

  describe("v2.1", () => {
    loadFixtures("v2.1").forEach((fixture) => {
      it(fixture.name, () => {
        const { manifest, expectedManifest, groupSpecs, expectedGroupSpecs } = setup(fixture);
        expect(manifest).toEqual(expectedManifest);
        expect(groupSpecs).toEqual(expectedGroupSpecs);
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

  describe("semantic-only-invalid SDLs", () => {
    const semanticInvalidDir = path.join(FIXTURES_INPUT_ROOT, "semantic-only-invalid");

    if (!fs.existsSync(semanticInvalidDir)) {
      it("semantic-only-invalid fixtures directory must exist", () => {
        throw new Error(`Semantic-only-invalid fixtures directory not found: ${semanticInvalidDir}`);
      });
      return;
    }

    fs.globSync("*.yaml", { cwd: semanticInvalidDir }).forEach((filename) => {
      it(`schema accepts: ${filename}`, () => {
        const fixturePath = path.join(semanticInvalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const inputYAML: SDLInput = yaml.raw(input);

        const schemaValidator = compileSchema(INPUT_SCHEMA_PATH);
        const schemaValid = schemaValidator(inputYAML);
        expect(schemaValid).toBe(true);
      });

      it(`parser rejects: ${filename}`, () => {
        const fixturePath = path.join(semanticInvalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const inputYAML: SDLInput = yaml.raw(input);

        const result = generateManifest(inputYAML);
        expect(result.ok).toBe(false);
      });
    });
  });

  describe("schema-only-invalid SDLs", () => {
    const schemaOnlyInvalidDir = path.join(FIXTURES_INPUT_ROOT, "schema-only-invalid");

    if (!fs.existsSync(schemaOnlyInvalidDir)) {
      it("schema-only-invalid fixtures directory must exist", () => {
        throw new Error(`Schema-only-invalid fixtures directory not found: ${schemaOnlyInvalidDir}`);
      });
      return;
    }

    fs.globSync("*.yaml", { cwd: schemaOnlyInvalidDir }).forEach((filename) => {
      it(`schema rejects: ${filename}`, () => {
        const fixturePath = path.join(schemaOnlyInvalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const inputYAML: SDLInput = yaml.raw(input);

        const schemaValidator = compileSchema(INPUT_SCHEMA_PATH);
        const schemaValid = schemaValidator(inputYAML);
        expect(schemaValid).toBe(false);
      });

      it(`generateManifest also rejects: ${filename}`, () => {
        const fixturePath = path.join(schemaOnlyInvalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const inputYAML: SDLInput = yaml.raw(input);

        const result = generateManifest(inputYAML);
        expect(result.ok).toBe(false);
      });
    });
  });

  describe("canonical byte-level equality", () => {
    const canonicalFixtures = [
      { version: "v2.0", name: "simple" },
      { version: "v2.0", name: "ip-endpoint" },
      { version: "v2.1", name: "credentials" },
    ];

    canonicalFixtures.forEach(({ version, name }) => {
      it(`${version}/${name} manifest canonical JSON matches`, () => {
        const inputPath = path.join(FIXTURES_INPUT_ROOT, version, name, "input.yaml");
        const manifestPath = path.join(FIXTURES_OUTPUT_ROOT, version, name, "manifest.json");

        const rawSDL = fs.readFileSync(inputPath, "utf8");
        const untrustedSDL: SDLInput = yaml.raw(rawSDL);
        const result = generateManifest(untrustedSDL);
        if (!result.ok) throw new Error(`generateManifest failed`);

        const canonicalTS = manifestToSortedJSON(result.value.groups);
        const goFixture = JSON.parse(fs.readFileSync(manifestPath, "utf8"));
        const canonicalGo = manifestToSortedJSON(goFixture);

        expect(canonicalTS).toBe(canonicalGo);
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

    const groupSpecs = JSON.parse(manifestToSortedJSON(result.value.groupSpecs));
    const expectedGroupSpecs = JSON.parse(fs.readFileSync(fixture.groupSpecsPath, "utf8"));

    validateAgainstSchema("manifest output", manifest, MANIFEST_OUTPUT_SCHEMA_PATH);
    validateAgainstSchema("groups output", groupSpecs, GROUPS_OUTPUT_SCHEMA_PATH);

    return {
      manifest,
      expectedManifest,
      groupSpecs,
      expectedGroupSpecs,
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
      const groupSpecsPath = path.join(FIXTURES_OUTPUT_ROOT, version, fixtureName, "group-specs.json");

      if (!fs.existsSync(manifestPath)) {
        throw new Error(`manifest.json not generated for ${fixtureName} (run: make generate-sdl-fixtures)`);
      }

      if (!fs.existsSync(groupSpecsPath)) {
        throw new Error(`group-specs.json not generated for ${fixtureName} (run: make generate-sdl-fixtures)`);
      }

      return {
        name: fixtureName,
        inputPath,
        manifestPath,
        groupSpecsPath,
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
  groupSpecsPath: string;
}
