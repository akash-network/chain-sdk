#!/usr/bin/env -S node --experimental-strip-types --no-warnings

import { type DescEnum, type DescMessage } from "@bufbuild/protobuf";
import {
  createEcmaScriptPlugin,
  type GeneratedFile,
  runNodeJs,
  type Schema,
} from "@bufbuild/protoplugin";

runNodeJs(
  createEcmaScriptPlugin({
    name: "protoc-gen-type-index-files",
    version: "v1",
    generateTs,
  }),
);

function generateTs(schema: Schema): void {
  const protoSource = process.env.PROTO_SOURCE;
  if (!protoSource) {
    throw new Error("PROTO_SOURCE is required and should be set to 'node', 'provider', or 'cosmos'");
  }

  const indexFiles: Record<string, {
    file: GeneratedFile;
    symbols: Set<string>;
  }> = {};
  const namespacePrefix = protoSource === "provider" ? "provider." : "";
  schema.files.forEach((file) => {
    const packageParts = file.proto.package.split(".");
    const namespace = namespacePrefix + packageParts[0];
    const version = packageParts.at(-1);
    const path = `index.${namespace}.${version}.ts`;
    indexFiles[path] ??= {
      file: schema.generateFile(path),
      symbols: new Set(),
    };
    const { file: indexFile, symbols: fileSymbols } = indexFiles[path];

    const typesToExport: Array<{ exportedName: string; name: string }> = [];
    for (const type of schema.typesInFile(file)) {
      if (type.kind === "service" || type.kind === "extension") continue;

      const name = genName(type);
      const exportedName = fileSymbols.has(name) ? genUniqueName(type, fileSymbols) : name;
      fileSymbols.add(exportedName);
      typesToExport.push({ exportedName, name });
    }

    if (typesToExport.length > 0) {
      const symbolsToExport = typesToExport.map((type) => type.exportedName === type.name ? type.exportedName : `${type.name} as ${type.exportedName}`).join(", ");
      indexFile.print(`export { ${symbolsToExport} } from "./${file.name}.ts";`);
    }
  });
}

function genName(type: DescMessage | DescEnum): string {
  return type.typeName.slice(type.file.proto.package.length + 1).replace(/\./g, "_");
}

let uniqueNameCounter = 0;
function genUniqueName(type: DescMessage | DescEnum, allSymbols: Set<string>, attempt = 0): string {
  const name = genName(type);
  if (allSymbols.has(name)) {
    const packageParts = type.file.proto.package.split(".");
    const prefix = packageParts.slice(-2 - attempt, -1).map(capitalize).join("_");
    let newName = `${prefix}_${name}`;
    if (newName === name) {
      newName = `${prefix}_${name}_${uniqueNameCounter++}`;
    }
    return allSymbols.has(newName) ? genUniqueName(type, allSymbols, attempt + 1) : newName;
  }
  return name;
}

function capitalize(str: string): string {
  return str[0].toUpperCase() + str.slice(1);
}
