import { default as stableStringify } from "json-stable-stringify";

import type { GenerateManifestOkResult, Manifest } from "./generateManifest.ts";

const decoder = new TextDecoder();
const encoder = new TextEncoder();
const NULLABLE_MANIFEST_KEYS = new Set(["command", "args", "env", "hosts", "allOf", "anyOf"]);
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

  // Format price amount as LegacyDec (18 decimal places) to match Go output
  if (key === "amount" && typeof this === "object" && this && Object.hasOwn(this, "denom") && (typeof value === "string" || typeof value === "number")) {
    return formatLegacyDec(String(value));
  }

  if (OMITTED_MANIFEST_KEYS.has(key) && ((Array.isArray(value) && value.length === 0) || value === 0)) {
    // In requirements context (group-specs), empty attributes should be null, not omitted
    if (key === "attributes" && typeof this === "object" && this && Object.hasOwn(this, "signedBy")) {
      return null;
    }
    return undefined;
  }

  return value;
}

function formatLegacyDec(s: string): string {
  if (!s) return "0.000000000000000000";

  // Normalize scientific notation (e.g. "1e-7") to plain decimal
  if (s.includes("e") || s.includes("E")) {
    s = Number(s).toFixed(18);
  }

  if (s.includes(".")) {
    const [int, frac = ""] = s.split(".");
    const truncated = frac.slice(0, 18);
    const pad = 18 - truncated.length;
    return pad > 0 ? `${int}.${truncated}${"0".repeat(pad)}` : `${int}.${truncated}`;
  }
  return `${s}.${"0".repeat(18)}`;
}

const MANIFEST_VERSION_FIELD_MAPPING: Record<string, string> = {
  quantity: "size",
  sequenceNumber: "sequence_number",
  signedBy: "signed_by",
  allOf: "all_of",
  anyOf: "any_of",
};
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
