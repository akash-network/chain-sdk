import { describe, expect, it } from "@jest/globals";
import type { ErrorObject, ValidateFunction } from "ajv";
import AjvModule from "ajv";
import fs from "fs";
import { load } from "js-yaml";
import path from "path";

import { Group } from "../../generated/protos/akash/manifest/v2beta3/group.ts";
import { generateManifest } from "../manifest/generateManifest.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
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

function validateSchemas(inputBytes: string) {
  const inputYAML = load(inputBytes) as SDLInput;
  validateAgainstSchema("input", inputYAML, inputSchemaPath);

  const result = generateManifest(inputYAML);
  if (!result.ok) throw new Error(`generateManifest failed: ${JSON.stringify(result.value)}`);

  const manifest = manifestToFixtureFormat(result.value.groups);
  return { manifest };
}

const SNAKE_TO_CAMEL: Record<string, string> = {
  endpoint_sequence_number: "endpointSequenceNumber",
  external_port: "externalPort",
  http_options: "httpOptions",
  max_body_size: "maxBodySize",
  next_cases: "nextCases",
  next_timeout: "nextTimeout",
  next_tries: "nextTries",
  read_timeout: "readTimeout",
  send_timeout: "sendTimeout",
};

const SERVICE_DEFAULTS: Record<string, unknown> = {
  args: null,
  command: null,
  credentials: null,
  env: null,
};

const EXPOSE_DEFAULTS: Record<string, unknown> = {
  endpointSequenceNumber: 0,
  externalPort: 0,
  service: "",
  ip: "",
  hosts: null,
  global: false,
};

const HTTP_OPTIONS_DEFAULTS: Record<string, unknown> = {
  nextTimeout: 0,
};

const STORAGE_PARAM_DEFAULTS: Record<string, unknown> = {
  readOnly: false,
};

function manifestToFixtureFormat(groups: unknown[]): unknown[] {
  return groups.map((g) => Group.toJSON(g as never)).map((groupJson) => {
    const obj = JSON.parse(JSON.stringify(groupJson)) as Record<string, unknown>;
    return convertToFixtureFormat(obj);
  });
}

function convertToFixtureFormat(obj: unknown, parentKey?: string): unknown {
  if (obj === null || obj === undefined) return obj;
  if (Array.isArray(obj)) {
    const converted = obj.map((item) => convertToFixtureFormat(item, parentKey));
    if (parentKey === "endpoints") {
      return converted.map((ep) => {
        const e = ep as Record<string, unknown>;
        if (Object.keys(e).length === 0) return { sequence_number: 0 };
        if (e.sequence_number === undefined) return { ...e, sequence_number: 0 };
        return e;
      });
    }
    return converted;
  }
  if (typeof obj === "object") {
    const rec = obj as Record<string, unknown>;
    const out: Record<string, unknown> = {};
    for (const key of Object.keys(rec)) {
      const camelKey = SNAKE_TO_CAMEL[key] ?? key;
      if (key === "quantity" && rec.quantity && typeof rec.quantity === "object") {
        const q = rec.quantity as Record<string, unknown>;
        out.size = q.val !== undefined ? { val: decodeVal(q.val) } : convertToFixtureFormat(rec.quantity);
      } else if (key === "val") {
        out[camelKey] = decodeVal(rec[key]);
      } else if (key === "kind" && typeof rec.kind === "string") {
        out[camelKey] = endpointKindToNumber(rec.kind);
      } else if (key !== "quantity") {
        out[camelKey] = convertToFixtureFormat(rec[key], key);
      }
    }
    if (parentKey === "services") {
      for (const [k, v] of Object.entries(SERVICE_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
    }
    if (parentKey === "resources" && !("endpoints" in out)) {
      out.endpoints = [];
    }
    if (parentKey === "expose") {
      for (const [k, v] of Object.entries(EXPOSE_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
      if (Array.isArray(out.hosts) && out.hosts.length === 0) out.hosts = null;
      if (out.httpOptions && typeof out.httpOptions === "object") {
        const ho = out.httpOptions as Record<string, unknown>;
        for (const [k, v] of Object.entries(HTTP_OPTIONS_DEFAULTS)) {
          if (!(k in ho)) ho[k] = v;
        }
      }
    }
    if (parentKey === "storage" && "name" in out && "mount" in out) {
      for (const [k, v] of Object.entries(STORAGE_PARAM_DEFAULTS)) {
        if (!(k in out)) out[k] = v;
      }
    }
    return out;
  }
  return obj;
}

function decodeVal(v: unknown): string {
  if (v === null || v === undefined) return "";
  if (typeof v === "string") {
    try {
      const bytes = Uint8Array.from(atob(v), (c) => c.charCodeAt(0));
      return new TextDecoder().decode(bytes);
    } catch {
      return v;
    }
  }
  return String(v);
}

function endpointKindToNumber(kind: string): number {
  const map: Record<string, number> = {
    SHARED_HTTP: 0,
    RANDOM_PORT: 1,
    LEASED_IP: 2,
  };
  return map[kind] ?? 0;
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
  const { manifest: actualManifest } = validateSchemas(inputBytes);

  const expectedManifest = JSON.parse(fs.readFileSync(fixture.manifestPath, "utf8"));

  expect(normalizeResourceVal(actualManifest)).toEqual(expectedManifest);
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
