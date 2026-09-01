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
import { toGoGroupSpecJSON } from "./goGroupSpecJSON.ts";

const PROJECT_ROOT = path.join(__dirname, "../../..");
const FIXTURES_INPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/input");
const FIXTURES_OUTPUT_ROOT = path.join(PROJECT_ROOT, "testdata/sdl/output-fixtures");
const INPUT_SCHEMA_PATH = path.join(PROJECT_ROOT, "go/sdl/sdl-input.schema.yaml");

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

    // Fixtures that exist to lock a specific *semantic* rejection (not just "the
    // schema rejects it") pin their exact error, so they can't silently pass for
    // the wrong reason if the schema or the fixture later drifts.
    const EXPECTED_ERRORS: Record<string, { instancePath: string; messageIncludes: string }> = {
      "v2.1-reclamation-bad-window.yaml": {
        instancePath: "/reclamation/min_window",
        messageIncludes: "whole number followed by s, m, or h",
      },
      "v2.1-tee-cpu-gpu-no-gpu.yaml": {
        instancePath: "/services/web/params/tee",
        messageIncludes: "tee type requires gpu resources",
      },
      // Both arch fixtures are rejected by Go's parser too (`v2CPUAttributes.
      // UnmarshalYAML`), so pinning the position here is what makes "same SDL,
      // same answer, same place" a checked claim rather than a coincidence.
      "cpu-arch-unknown-value.yaml": {
        instancePath: "/profiles/compute/web/resources/cpu/attributes/arch",
        messageIncludes: "amd64, arm64",
      },
      "cpu-unknown-attribute.yaml": {
        instancePath: "/profiles/compute/web/resources/cpu/attributes",
        messageIncludes: '"vendor" is not allowed',
      },
    };

    fs.globSync("*.yaml", { cwd: invalidDir }).forEach((filename) => {
      it(filename, () => {
        const fixturePath = path.join(invalidDir, filename);
        const input = fs.readFileSync(fixturePath, "utf8");
        const sdl: SDLInput = yaml.raw(input);
        const result = generateManifest(sdl);
        expect(result.ok).toBe(false);

        const expected = EXPECTED_ERRORS[filename];
        if (!result.ok && expected) {
          expect(result.value).toContainEqual(expect.objectContaining({
            instancePath: expected.instancePath,
            message: expect.stringContaining(expected.messageIncludes),
          }));
        }
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

    // Resource attributes (cpu/gpu/memory/storage) never reach a manifest — they
    // only exist in the group spec — so the manifest comparison above cannot see
    // them drift. Both sides go through `normalizeGroupSpecsJSON` so the Go
    // 18-decimal price and this side's plain decimal compare as the same number.
    const groupSpecs = JSON.parse(JSON.stringify(toGoGroupSpecJSON(result.value.groupSpecs)), normalizeGroupSpecsJSON);
    const expectedGroupSpecs = JSON.parse(fs.readFileSync(fixture.groupSpecsPath, "utf8"), normalizeGroupSpecsJSON);

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

function normalizeGroupSpecsJSON(this: unknown, key: string, value: unknown): unknown {
  if (typeof this !== "object" || this === null) return value;

  if (key === "amount" && "denom" in this) {
    return LegacyDec.encode(value as string);
  }

  return value;
}

interface Fixture {
  name: string;
  inputPath: string;
  manifestPath: string;
  groupSpecsPath: string;
}
