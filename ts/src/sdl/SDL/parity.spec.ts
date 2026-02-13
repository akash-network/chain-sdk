import { describe, expect, it } from "@jest/globals";
import type { ErrorObject, ValidateFunction } from "ajv";
import AjvModule from "ajv";
import fs from "fs";
import { load } from "js-yaml";
import path from "path";

import { SDL } from "./SDL.ts";

const fixturesInputRoot = path.join(__dirname, "../../../../testdata/sdl/input");
const fixturesOutputRoot = path.join(__dirname, "../../../../testdata/sdl/output-fixtures");
const inputSchemaPath = path.join(__dirname, "../../../../go/sdl/sdl-input.schema.yaml");
// @ts-expect-error - AjvModule has non-standard export, cast needed for instantiation
const ajv: { compile: (schema: Record<string, unknown>) => ValidateFunction } = new (AjvModule as unknown as new (options?: { allErrors?: boolean }) => typeof AjvModule)({ allErrors: true });

interface Fixture {
  name: string;
  inputPath: string;
  manifestPath: string;
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

function validateSchemas(inputBytes: string, version: "beta2" | "beta3") {
  const inputYAML = load(inputBytes);
  validateAgainstSchema("input", inputYAML, inputSchemaPath);

  const sdl = SDL.fromString(inputBytes, version);
  const manifest = sdl.v3Manifest(false);

  return { manifest };
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

function validateFixtures(fixture: Fixture, version: "beta2" | "beta3") {
  const inputBytes = fs.readFileSync(fixture.inputPath, "utf8");
  const { manifest: actualManifest } = validateSchemas(inputBytes, version);

  const expectedManifest = JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8"));

  expect(normalizeResourceVal(actualManifest)).toEqual(expectedManifest);
}

describe("SDL Parity Tests", () => {
  describe("v2.0", () => {
    loadFixtures("v2.0").forEach((fixture) => {
      it(fixture.name, () => validateFixtures(fixture, "beta2"));
    });
  });

  describe("v2.1", () => {
    loadFixtures("v2.1").forEach((fixture) => {
      it(fixture.name, () => validateFixtures(fixture, "beta3"));
    });
  });

  describe("invalid SDLs rejected", () => {
    const invalidDir = path.join(fixturesInputRoot, "invalid");

    if (!fs.existsSync(invalidDir)) {
      it.skip("invalid fixtures directory not found", () => {});
      return;
    }

    fs.readdirSync(invalidDir)
      .filter((f) => f.endsWith(".yaml"))
      .forEach((filename) => {
        it(filename, () => {
          const fixturePath = path.join(invalidDir, filename);
          const input = fs.readFileSync(fixturePath, "utf8");
          expect(() => SDL.fromString(input, "beta3")).toThrow();
        });
      });
  });
});
