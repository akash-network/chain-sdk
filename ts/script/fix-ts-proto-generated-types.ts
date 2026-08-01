#!/usr/bin/env -S node

import { promises as fs } from "node:fs";
import { dirname, relative as relativePath, resolve as resolvePath } from "node:path";

const helperNames = ["isSet", "bytesFromBase64", "base64FromBytes", "toTimestamp", "fromTimestamp", "fromJsonTimestamp", "numberToLong", "isObject"];
const helperRegex = new RegExp(
  `^(function|const)\\s+(${helperNames.join("|")})\\b`,
  "gm",
);
const typeHelpers = ["MessageFns", "DeepPartial"];
const helperTypeRegex = new RegExp(
  `^(interface|type)\\s+(${typeHelpers.join("|")})\\b`,
  "gm",
);

const ROOT_DIR = resolvePath(import.meta.dirname, "..", "src");
const STORAGE_SERVICE_PATH = resolvePath(ROOT_DIR, "generated/protos/akash/manifest/v2beta3/service.ts");

const typesToPatch = new Set<string>();
// Per-message field-type overlays for representation-changing custom types,
// e.g. { "akash.…ResourceValue": { val: "bigint" } }. Emitted by
// protoc-gen-customtype-patches.ts; drives the public `export type` overlay.
const typeOverrides: Record<string, Record<string, string>> = {};
for await (const patchFile of fs.glob(`${ROOT_DIR}/generated/patches/*CustomTypePatches.ts`)) {
  const { patches, typeOverrides: fileOverrides } = await import(patchFile);
  Object.keys(patches).forEach((key) => typesToPatch.add(key));
  Object.assign(typeOverrides, fileOverrides);
}

for await (const path of fs.glob(`${ROOT_DIR}/generated/protos/**/*.ts`)) {
  const source = await fs.readFile(path, "utf8");
  let newSource = source;

  // Remove the `create` method from message objects
  newSource = newSource.replace(/^\s*create\(base\?:\s*DeepPartial<\w+>\):\s*\w+\s*\{\s*return\s*\w+\.fromPartial\(base \?\? \{\}\);\s*\},?\n?/gm, "");
  newSource = injectOwnHelpers(newSource, path);
  newSource = preserveOptionalStorageKeyRef(newSource, path);

  newSource = applyPatching(newSource, path, typesToPatch, typeOverrides);

  if (newSource !== source) {
    await fs.writeFile(path, newSource);
  }
}

function preserveOptionalStorageKeyRef(source: string, path: string): string {
  if (resolvePath(path) !== STORAGE_SERVICE_PATH) return source;

  // A proto3 scalar has no wire-level presence, but keyRef is optional SDL
  // input. Leaving ts-proto's empty-string default in place would add
  // `keyRef: ""` to every existing manifest object. Keep the public TS shape
  // aligned with Go's `omitempty` JSON/YAML contract.
  source = replaceGeneratedText(source, "  keyRef: string;", "  keyRef?: string;");
  source = replaceGeneratedText(source, `, keyRef: ""`, "");
  source = replaceGeneratedText(
    source,
    `if (message.keyRef !== "") {`,
    `if (message.keyRef !== undefined && message.keyRef !== "") {`,
    2,
  );
  source = replaceGeneratedText(
    source,
    `      keyRef: isSet(object.key_ref) ? globalThis.String(object.key_ref) : "",`,
    `      ...(isSet(object.key_ref) ? { keyRef: globalThis.String(object.key_ref) } : {}),`,
  );
  source = replaceGeneratedText(
    source,
    `    message.keyRef = object.keyRef ?? "";`,
    `    if (object.keyRef !== undefined) {\n      message.keyRef = object.keyRef;\n    }`,
  );

  return source;
}

function replaceGeneratedText(source: string, before: string, after: string, expectedCount = 1): string {
  const chunks = source.split(before);
  const count = chunks.length - 1;
  if (count !== expectedCount) {
    throw new Error(`Expected ${expectedCount} generated occurrence(s) of ${JSON.stringify(before)}, found ${count}`);
  }

  return chunks.join(after);
}

function injectOwnHelpers(source: string, path: string) {
  const foundHelperNames = new Set<string>();
  source = source.replace(helperRegex, (_, kind, name) => {
    foundHelperNames.add(name);
    return `${kind} _unused_${name}`;
  });

  const foundTypeHelperNames = new Set<string>();
  source = source.replace(helperTypeRegex, (_, kind, name) => {
    foundTypeHelperNames.add(name);
    return `${kind} _unused_${name}`;
  });

  const importHelpers = foundHelperNames.size
    ? `import { ${Array.from(foundHelperNames).join(", ")} } from "${relativePath(dirname(path), `${ROOT_DIR}/encoding/typeEncodingHelpers.ts`)}"\n`
    : "";
  const importTypeHelpers = foundTypeHelperNames.size
    ? `import type { ${Array.from(foundTypeHelperNames).join(", ")} } from "${relativePath(dirname(path), `${ROOT_DIR}/encoding/typeEncodingHelpers.ts`)}"\n`
    : "";

  return importHelpers + importTypeHelpers + source;
}

function applyPatching(
  source: string,
  filePath: string,
  typesToPatch: Set<string>,
  typeOverrides: Record<string, Record<string, string>>,
) {
  const imports = new Set<string>();
  const exports: string[] = [];
  const overlayTypes: string[] = [];

  let newSource = source.replace(
    /^export const (\w+)(:\s*MessageFns<[^>]+,\s*["']([^"']+)["']>\s*=)/gm,
    (match, symbolName, typeAnnotation, fullName) => {
      if (!typesToPatch.has(fullName)) return match;

      const namespace = fullName.split(".")[0];
      const prefix = namespace === "akash" ? "node" : namespace;
      const importPath = relativePath(filePath, `${ROOT_DIR}/generated/protos/patches/${prefix}PatchMessage.ts`);
      imports.add(`import { patched } from "${importPath}";`);

      const overrideFields = typeOverrides[fullName];
      if (overrideFields) {
        // Representation-changing message: its wire interface (bytes) is renamed
        // to `_${symbolName}Wire` below, a public `export type ${symbolName}`
        // overlay is emitted, and the patched value is retyped to it.
        const omitKeys = Object.keys(overrideFields).map((field) => `"${field}"`).join(" | ");
        const overlayFields = Object.entries(overrideFields).map(([field, jsType]) => `${field}: ${jsType}`).join("; ");
        overlayTypes.push(`export type ${symbolName} = Omit<_${symbolName}Wire, ${omitKeys}> & { ${overlayFields} };`);
        exports.push(`export const ${symbolName} = patched(_${symbolName}) as unknown as MessageFns<${symbolName}, "${fullName}">;`);
      } else {
        exports.push(`export const ${symbolName} = patched(_${symbolName});`);
      }

      return `const _${symbolName}${typeAnnotation}`;
    },
  );

  if (!exports.length) return source;

  // For each overlaid message defined in this file, rename its wire interface and
  // every in-file *type* reference to it (`Name` -> `_NameWire`), leaving the
  // inner impl operating on the wire (bytes) shape. The lookbehind/lookahead skip
  // value positions (`_Name`, `createBaseName`) and string literals (the `$type`
  // and fully-qualified name), which is why this runs after the const rename.
  for (const fullName of Object.keys(typeOverrides)) {
    const name = fullName.split(".").pop()!;
    if (!newSource.includes(`interface ${name} `)) continue;
    newSource = newSource
      .replace(new RegExp(`export interface ${name} `), `interface _${name}Wire `)
      .replace(new RegExp(`(?<![\\w."])${name}(?![\\w"])`, "g"), `_${name}Wire`);
  }

  return Array.from(imports).join("\n") + "\n"
    + newSource + "\n"
    + (overlayTypes.length ? overlayTypes.join("\n") + "\n" : "")
    + exports.join("\n") + "\n";
}
