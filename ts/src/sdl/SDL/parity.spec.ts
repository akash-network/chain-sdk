import { describe, expect, it } from "@jest/globals";
import type { ErrorObject, ValidateFunction } from "ajv";
import AjvModule from "ajv";
import fs from "fs";
import { load } from "js-yaml";
import path from "path";

import { generateManifest } from "../manifest/generateManifest.ts";
import { manifestToSortedJSON } from "../manifest/generateManifestVersion.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";

const fixturesInputRoot = path.join(__dirname, "../../../../testdata/sdl/input");
const fixturesOutputRoot = path.join(__dirname, "../../../../testdata/sdl/output-fixtures");
const inputSchemaPath = path.join(__dirname, "../../../../go/sdl/sdl-input.schema.yaml");
const manifestOutputSchemaPath = path.join(__dirname, "../../../../go/sdl/manifest-output.schema.yaml");
const groupsOutputSchemaPath = path.join(__dirname, "../../../../go/sdl/groups-output.schema.yaml");
// @ts-expect-error - AjvModule has non-standard export, cast needed for instantiation
const ajv: { compile: (schema: Record<string, unknown>) => ValidateFunction } = new (AjvModule as unknown as new (options?: { allErrors?: boolean }) => typeof AjvModule)({ allErrors: true });

interface Fixture {
  name: string;
  inputPath: string;
  manifestPath: string;
  groupSpecsPath: string;
}

const schemaCache = new Map<string, ValidateFunction>();

function compileSchema(schemaPath: string): ValidateFunction {
  const cached = schemaCache.get(schemaPath);
  if (cached) {
    return cached;
  }

  const schemaContent = fs.readFileSync(schemaPath, "utf8");
  const schema = load(schemaContent);
  const validator = ajv.compile(schema as Record<string, unknown>);
  schemaCache.set(schemaPath, validator);
  return validator;
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

function loadFixtures(version: string): Fixture[] {
  const inputVersionDir = path.join(fixturesInputRoot, version);

  if (!fs.existsSync(inputVersionDir)) {
    throw new Error(`Fixtures directory ${inputVersionDir} does not exist`);
  }

  const entries = fs.readdirSync(inputVersionDir, { withFileTypes: true });

  return entries
    .filter((entry) => entry.isDirectory())
    .map((entry) => {
      const fixtureName = entry.name;
      const inputPath = path.join(inputVersionDir, fixtureName, "input.yaml");
      const manifestPath = path.join(fixturesOutputRoot, version, fixtureName, "manifest.json");
      const groupSpecsPath = path.join(fixturesOutputRoot, version, fixtureName, "group-specs.json");

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

function validateSchemas(inputBytes: string) {
  const inputYAML = load(inputBytes) as SDLInput;
  validateAgainstSchema("input", inputYAML, inputSchemaPath);

  const result = generateManifest(inputYAML);
  if (!result.ok) throw new Error(`generateManifest failed: ${JSON.stringify(result.value)}`);

  const manifest = JSON.parse(manifestToSortedJSON(result.value.groups)) as unknown;
  const groupSpecs = JSON.parse(manifestToSortedJSON(result.value.groupSpecs)) as unknown;

  validateAgainstSchema("manifest output", manifest, manifestOutputSchemaPath);
  validateAgainstSchema("groups output", groupSpecs, groupsOutputSchemaPath);

  return { manifest, groupSpecs };
}

function strVal(v: unknown): string {
  if (v === null || v === undefined) return "";
  return String(v);
}

function formatAmount(amount: unknown): string {
  const s = strVal(amount);
  if (!s) return "0.000000000000000000";
  if (s.includes(".")) {
    const [, frac] = s.split(".");
    const pad = 18 - (frac?.length ?? 0);
    return pad > 0 ? s + "0".repeat(pad) : s.slice(0, s.length + (pad ?? 0));
  }
  return `${s}.${"0".repeat(18)}`;
}

function normalizeResourceVal(obj: unknown): unknown {
  if (obj === null || obj === undefined) return obj;
  if (Array.isArray(obj)) return obj.map(normalizeResourceVal);
  if (typeof obj === "object") {
    const out: Record<string, unknown> = {};
    const rec = obj as Record<string, unknown>;
    for (const key of Object.keys(rec)) {
      if (key === "val") {
        out[key] = strVal(rec[key]);
      } else if (key === "amount" && typeof rec.denom !== "undefined") {
        out[key] = formatAmount(rec[key]);
      } else {
        out[key] = normalizeResourceVal(rec[key]);
      }
    }
    return out;
  }
  return obj;
}

function validateFixtures(fixture: Fixture) {
  const inputBytes = fs.readFileSync(fixture.inputPath, "utf8");
  const { manifest: actualManifest, groupSpecs: actualGroupSpecs } = validateSchemas(inputBytes);

  const expectedManifest = JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8"));
  const expectedGroupSpecs = JSON.parse(fs.readFileSync(fixture.groupSpecsPath, "utf8"));

  expect(normalizeResourceVal(actualManifest)).toEqual(expectedManifest);
  expect(normalizeResourceVal(actualGroupSpecs)).toEqual(expectedGroupSpecs);
}

describe("SDL Parity Tests", () => {
  describe("v2.0", () => {
    loadFixtures("v2.0").forEach((fixture) => {
      it(fixture.name, () => validateFixtures(fixture));
    });
  });

  describe("v2.1", () => {
    loadFixtures("v2.1").forEach((fixture) => {
      it(fixture.name, () => validateFixtures(fixture));
    });
  });

  describe("invalid SDLs rejected", () => {
    const invalidDir = path.join(fixturesInputRoot, "invalid");

    if (!fs.existsSync(invalidDir)) {
      it("invalid fixtures directory must exist", () => {
        throw new Error(`Invalid fixtures directory not found: ${invalidDir}`);
      });
      return;
    }

    fs.readdirSync(invalidDir)
      .filter((f) => f.endsWith(".yaml"))
      .forEach((filename) => {
        it(filename, () => {
          const fixturePath = path.join(invalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          const inputYAML = load(input) as SDLInput;

          // generateManifest runs both schema validation and semantic validation,
          // so it should reject all invalid fixtures regardless of whether the
          // issue is structural (schema) or semantic (parser logic)
          const result = generateManifest(inputYAML);
          expect(result.ok).not.toBe(true);
        });
      });
  });

  describe("semantic-only-invalid SDLs", () => {
    const semanticInvalidDir = path.join(fixturesInputRoot, "semantic-only-invalid");

    if (!fs.existsSync(semanticInvalidDir)) {
      it("semantic-only-invalid fixtures directory must exist", () => {
        throw new Error(`Semantic-only-invalid fixtures directory not found: ${semanticInvalidDir}`);
      });
      return;
    }

    fs.readdirSync(semanticInvalidDir)
      .filter((f) => f.endsWith(".yaml"))
      .forEach((filename) => {
        it(`schema accepts: ${filename}`, () => {
          const fixturePath = path.join(semanticInvalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          const inputYAML = load(input) as SDLInput;

          const schemaValidator = compileSchema(inputSchemaPath);
          const schemaValid = schemaValidator(inputYAML);
          expect(schemaValid).toBe(true);
        });

        it(`parser rejects: ${filename}`, () => {
          const fixturePath = path.join(semanticInvalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          const inputYAML = load(input) as SDLInput;

          const result = generateManifest(inputYAML);
          expect(result.ok).toBe(false);
        });
      });
  });

  describe("schema-only-invalid SDLs", () => {
    const schemaOnlyInvalidDir = path.join(fixturesInputRoot, "schema-only-invalid");

    if (!fs.existsSync(schemaOnlyInvalidDir)) {
      it("schema-only-invalid fixtures directory must exist", () => {
        throw new Error(`Schema-only-invalid fixtures directory not found: ${schemaOnlyInvalidDir}`);
      });
      return;
    }

    fs.readdirSync(schemaOnlyInvalidDir)
      .filter((f) => f.endsWith(".yaml"))
      .forEach((filename) => {
        it(`schema rejects: ${filename}`, () => {
          const fixturePath = path.join(schemaOnlyInvalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          const inputYAML = load(input) as SDLInput;

          const schemaValidator = compileSchema(inputSchemaPath);
          const schemaValid = schemaValidator(inputYAML);
          expect(schemaValid).toBe(false);
        });

        // Note: Unlike the Go parser which accepts these inputs (Go doesn't run
        // schema validation in its parser), the TS generateManifest() also rejects
        // them because it embeds schema validation via validateSDL(). This confirms
        // that TS is at least as strict as the schema — a superset of Go's behavior.
        it(`generateManifest also rejects: ${filename}`, () => {
          const fixturePath = path.join(schemaOnlyInvalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          const inputYAML = load(input) as SDLInput;

          const result = generateManifest(inputYAML);
          expect(result.ok).toBe(false);
        });
      });
  });

  describe("canonical byte-level equality", () => {
    // For representative fixtures, verify that the TS-generated canonical JSON
    // is byte-identical to the committed Go-generated fixture.
    // This catches subtle serialization differences (field ordering, whitespace,
    // number formatting) that structural equality tests might miss.
    const canonicalFixtures = [
      { version: "v2.0", name: "simple" },
      { version: "v2.0", name: "ip-endpoint" },
      { version: "v2.1", name: "credentials" },
    ];

    canonicalFixtures.forEach(({ version, name }) => {
      it(`${version}/${name} manifest canonical JSON matches`, () => {
        const inputPath = path.join(fixturesInputRoot, version, name, "input.yaml");
        const manifestPath = path.join(fixturesOutputRoot, version, name, "manifest.json");

        const inputBytes = fs.readFileSync(inputPath, "utf8");
        const inputYAML = load(inputBytes) as SDLInput;
        const result = generateManifest(inputYAML);
        if (!result.ok) throw new Error(`generateManifest failed`);

        // Use the same canonical serialization as manifest version hashing
        const canonicalTS = manifestToSortedJSON(result.value.groups);

        // Canonicalize the Go fixture the same way (parse + re-serialize with sorted keys)
        const goFixture = JSON.parse(fs.readFileSync(manifestPath, "utf8"));
        const canonicalGo = manifestToSortedJSON(goFixture);

        expect(canonicalTS).toBe(canonicalGo);
      });

      // Group-specs use structural equality (toEqual) in the main parity tests.
      // Canonical byte-level comparison is only applied to manifests because
      // manifest version hashing depends on byte-identical canonical JSON.
    });
  });
});
