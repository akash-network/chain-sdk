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

const typesToPatch = new Set<string>();
for await (const patchFile of fs.glob(`${ROOT_DIR}/generated/patches/*CustomTypePatches.ts`)) {
  const { patches } = await import(patchFile);
  Object.keys(patches).forEach((key) => typesToPatch.add(key));
}

for await (const path of fs.glob(`${ROOT_DIR}/generated/protos/**/*.ts`)) {
  const source = await fs.readFile(path, "utf8");
  let newSource = source;

  // Remove the `create` method from message objects
  newSource = newSource.replace(/^\s*create\(base\?:\s*DeepPartial<\w+>\):\s*\w+\s*\{\s*return\s*\w+\.fromPartial\(base \?\? \{\}\);\s*\},?\n?/gm, "");
  newSource = coerceBigIntFromPartial(newSource);
  newSource = injectOwnHelpers(newSource, path);

  newSource = applyPatching(newSource, path, typesToPatch);

  if (newSource !== source) {
    await fs.writeFile(path, newSource);
  }
}

// ts-proto's fromPartial assigns bigint fields verbatim (`x ?? 0n` for scalars,
// `.map((e) => e)` for repeated), so a caller passing a string or number would be
// stored uncoerced. Wrap those assignments with BigInt(...) so fromPartial actually
// accepts `string | number | bigint`, matching the DeepPartial<bigint> contract.
function coerceBigIntFromPartial(source: string) {
  // Scalar fields: `message.x = object.x ?? 0n;`
  // `?? 0n` only appears for bigint scalars in fromPartial, so this is unambiguous.
  let result = source.replace(
    /(message\.[A-Za-z0-9_]+ = )(object\.[A-Za-z0-9_]+) \?\? 0n;/g,
    "$1($2 !== undefined && $2 !== null) ? BigInt($2) : 0n;",
  );

  // Repeated fields: `message.x = object.x?.map((e) => e) || [];`
  // This shape is shared by all scalar arrays, so restrict to fields the file
  // declares as `bigint[]`, skipping any name also used as a non-bigint array.
  const bigintArrays = new Set<string>();
  for (const [, name] of source.matchAll(/^\s*([A-Za-z0-9_]+): bigint\[\];/gm)) {
    bigintArrays.add(name);
  }
  for (const [, name] of source.matchAll(/^\s*([A-Za-z0-9_]+): (?:string|number|boolean|Uint8Array)\[\];/gm)) {
    if (bigintArrays.delete(name)) {
      console.warn(`fix-ts-proto: array field "${name}" is bigint[] and non-bigint[] in the same file; skipping BigInt coercion`);
    }
  }
  for (const name of bigintArrays) {
    result = result.replace(
      new RegExp(`(message\\.${name} = object\\.[A-Za-z0-9_]+\\?\\.map\\(\\(e\\) => )e(\\) \\|\\| \\[\\];)`, "g"),
      "$1BigInt(e)$2",
    );
  }

  return result;
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

function applyPatching(source: string, filePath: string, typesToPatch: Set<string>) {
  const imports = new Set<string>();
  const exports: string[] = [];

  const newSource = source.replace(
    /^export const (\w+)(:\s*MessageFns<[^>]+,\s*["']([^"']+)["']>\s*=)/gm,
    (match, symbolName, typeAnnotation, fullName) => {
      if (!typesToPatch.has(fullName)) return match;

      const namespace = fullName.split(".")[0];
      const prefix = namespace === "akash" ? "node" : namespace;
      const importPath = relativePath(filePath, `${ROOT_DIR}/generated/protos/patches/${prefix}PatchMessage.ts`);
      imports.add(`import { patched } from "${importPath}";`);
      exports.push(`export const ${symbolName} = patched(_${symbolName});`);

      return `const _${symbolName}${typeAnnotation}`;
    },
  );

  if (!exports.length) return source;

  return Array.from(imports).join("\n") + "\n" + newSource + "\n" + exports.join("\n") + "\n";
}
