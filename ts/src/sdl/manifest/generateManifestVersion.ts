import { default as stableStringify } from "json-stable-stringify";

import type { GenerateManifestOkResult, Manifest } from "./generateManifest.ts";

const decoder = new TextDecoder();
const encoder = new TextEncoder();
const NULLABLE_MANIFEST_KEYS = new Set(["command", "args", "env", "hosts"]);
const OMITTED_MANIFEST_KEYS = new Set(["kind", "attributes"]);

export async function generateManifestVersion(manifest: Manifest): Promise<Uint8Array> {
  const jsonStr = manifestToSortedJSON(manifest);
  const sortedBytes = encoder.encode(jsonStr);
  const sum = await crypto.subtle.digest("SHA-256", sortedBytes);
  return new Uint8Array(sum);
}

export function manifestToSortedJSON(manifest: Manifest | GenerateManifestOkResult["groupSpecs"]): string {
  const json = stableStringify(manifest, { replacer: manifestReplacer }) || "";
  return escapeHtml(renameFields(json));
}

function manifestReplacer(this: unknown, key: string | number, value: unknown): unknown {
  if (value && value instanceof Uint8Array) {
    return decoder.decode(value);
  }

  if (typeof key !== "string") {
    return value;
  }

  // only top-level "credentials" field can be null, credentials in params should be omitted
  if (typeof this === "object" && this && Object.hasOwn(this, "command") && key === "credentials" && value == null) {
    return null;
  }

  if (NULLABLE_MANIFEST_KEYS.has(key) && ((Array.isArray(value) && value.length === 0) || value == null)) {
    return null;
  }

  if (OMITTED_MANIFEST_KEYS.has(key) && ((Array.isArray(value) && value.length === 0) || value === 0)) {
    return undefined;
  }

  return value;
}

const MANIFEST_VERSION_FIELD_MAPPING: Record<string, string> = { quantity: "size", sequenceNumber: "sequence_number" };
const MANIFEST_VERSION_FIELD_REGEX = new RegExp(`"(${Object.keys(MANIFEST_VERSION_FIELD_MAPPING).join("|")})":`, "g");
function renameFields(jsonStr: string): string {
  MANIFEST_VERSION_FIELD_REGEX.lastIndex = 0; // reset regex state
  return jsonStr.replace(MANIFEST_VERSION_FIELD_REGEX, (_, field) => `"${MANIFEST_VERSION_FIELD_MAPPING[field]}":`);
}

const htmlEscapes: Record<string, string> = {
  "<": "\\u003c",
  ">": "\\u003e",
  "&": "\\u0026",
};

const HTML_SPECIAL_CHARS_REGEX = new RegExp(`[${Object.keys(htmlEscapes).join("")}]`, "g");
function escapeHtml(raw: string): string {
  HTML_SPECIAL_CHARS_REGEX.lastIndex = 0; // reset regex state
  return raw.replace(HTML_SPECIAL_CHARS_REGEX, (ch) => htmlEscapes[ch]);
}
