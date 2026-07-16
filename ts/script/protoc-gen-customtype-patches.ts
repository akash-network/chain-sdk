#!/usr/bin/env -S node

import { type DescField, type DescMessage, ScalarType } from "@bufbuild/protobuf";
import {
  createEcmaScriptPlugin,
  type GeneratedFile,
  type Printable,
  runNodeJs,
  type Schema,
} from "@bufbuild/protoplugin";
import { basename, normalize as normalizePath } from "path";

import type { CustomType } from "../src/encoding/customTypes/CustomType.ts";
import { findPathsToCustomField, getCustomType } from "../src/encoding/customTypes/utils.ts";

export interface PluginOptions {
  /**
   * if true, we will patch the whole tree of the message type, starting from the custom field type and up to the root
   * in case of patching ts-proto generated types which has self-references, we need to patch only leaf level
   * @default false
   */
  patchWholeTree: boolean;
}

runNodeJs(createEcmaScriptPlugin<PluginOptions>({ name: "protoc-gen-customtype-patches", version: "v1", parseOptions, generateTs }));

const PROTO_PATH = "../protos";

function parseOptions(rawOptions: Array<{
  key: string;
  value: string;
}>): PluginOptions {
  const options: PluginOptions = {
    patchWholeTree: false,
  };

  for (const { key, value } of rawOptions) {
    if (key === "patch_whole_tree") {
      options.patchWholeTree = value === "true";
    }
  }

  return options;
}

function generateTs(schema: Schema<PluginOptions>): void {
  const allPaths: DescField[][] = [];

  schema.files.forEach((file) => {
    file.messages.forEach((message) => {
      const paths = findPathsToCustomField(message, () => true);
      if (schema.options.patchWholeTree) {
        allPaths.push(...paths);
      } else {
        const leaves = paths.map((path) => path.slice(-1));
        allPaths.push(...leaves);
      }
    });
  });
  if (!allPaths.length) {
    return;
  }

  const messageToCustomFields: Map<DescMessage, Set<DescField>> = new Map();
  allPaths.forEach((path) => {
    path.forEach((field) => {
      if (!messageToCustomFields.has(field.parent)) {
        messageToCustomFields.set(field.parent, new Set());
      }
      messageToCustomFields.get(field.parent)!.add(field);
    });
  });

  const patches: string[] = [];
  const imports: Record<string, Set<string>> = {};
  // Per-message field-type overrides for representation-changing custom types
  // (e.g. `{ "…ResourceValue": { val: "bigint" } }`), consumed by
  // fix-ts-proto-generated-types.ts to overlay the public field type.
  const typeOverrides: Record<string, Record<string, string>> = {};
  const fileName = getOutputFileName(schema);
  const patchesFile = schema.generateFile(fileName);

  Array.from(messageToCustomFields.entries()).forEach(([descMessage, fields]) => {
    const encoded: string[] = [];
    // JSON / fromPartial field overrides, emitted only for representation-changing
    // fields. Dec-only messages leave these empty and keep the plain function form.
    const jsonReadLines: string[] = [];
    const jsonWriteLines: string[] = [];
    const partialLines: string[] = [];

    fields.forEach((field) => {
      const customType = getCustomType(field);

      if (customType) {
        const pathToCustomType = `../../encoding/customTypes/${customType.shortName}`;
        imports[pathToCustomType] ??= new Set();
        imports[pathToCustomType].add(customType.shortName);

        if (field.scalar === ScalarType.BYTES) {
          imports[`../../encoding/binaryEncoding`] ??= new Set(["encodeBinary", "decodeBinary"]);
        }

        encoded.push(generateFieldTransformation(field, customType, {
          value: "value",
          newValue: "newValue",
        }));

        if (customType.jsType) {
          // Representation-changing field: its JS type differs from the wire
          // bytes, so JSON and fromPartial need dedicated normalizers on top of
          // the generated (inner) impl, and the public field type is overridden.
          imports[pathToCustomType].add("bigIntFromJSON");
          imports[pathToCustomType].add("bigIntFromPartial");
          const local = field.localName;
          jsonReadLines.push(`newValue.${local} = bigIntFromJSON(object.${local});`);
          jsonWriteLines.push(`if (value.${local} != null) newValue.${local} = value.${local}.toString();`);
          partialLines.push(`if (newValue.${local} != null) newValue.${local} = bigIntFromPartial(newValue.${local});`);
          (typeOverrides[descMessage.typeName] ??= {})[local] = customType.jsType;
        }
      } else {
        encoded.push(generateNestedFieldUpdate(field, {
          fn: `p["${field.message!.typeName}"]`,
          value: "value",
          newValue: "newValue",
        }));
      }
    });

    const parent = fields.values().next().value!.parent;
    const path = normalizePath(`${PROTO_PATH}/${parent.file.name}`);
    imports[path] ??= new Set(["type *"]);
    const typeRef = `${dirnameToVar(path)}.${parent.name}`;

    // The `transform` callable (encode/decode) is shared by both forms. For
    // representation-changing messages its param straddles the public JS type
    // (on encode) and the wire type (on decode), so it is typed `any`.
    const transformFn = `function(value: ${jsonReadLines.length ? "any" : `${typeRef} | undefined | null`}, transformType: 'encode' | 'decode') {\n${
      indent(`if (value == null) return;`) + "\n"
      + indent(`const newValue = { ...value };`) + "\n"
      + indent(encoded.join("\n")) + "\n"
      + indent("return newValue;")
    }\n}`;

    if (jsonReadLines.length) {
      // Object form: a callable augmented with JSON / fromPartial normalizers
      // that patched() layers over the generated impl.
      patches.push(
        `"${descMessage.typeName}": Object.assign(${transformFn}, {\n${
          indent(`fromJSON(object: any) {\n${indent(`const newValue: any = {};`) + "\n" + indent(jsonReadLines.join("\n")) + "\n" + indent("return newValue;")}\n},`) + "\n"
          + indent(`toJSON(value: any) {\n${indent(`const newValue: any = {};`) + "\n" + indent(jsonWriteLines.join("\n")) + "\n" + indent("return newValue;")}\n},`) + "\n"
          + indent(`fromPartial(newValue: any) {\n${indent(partialLines.join("\n")) + "\n" + indent("return newValue;")}\n},`)
        }\n})`,
      );
    } else {
      // Plain function form (unchanged) — representation-preserving messages
      // (Dec) and nested-tree wrappers.
      patches.push(`"${descMessage.typeName}"${transformFn.replace(/^function/, "")}`);
    }
  });

  const importExtension = schema.options.importExtension ? `.${schema.options.importExtension}` : "";
  Object.entries(imports).forEach(([path, symbols]) => {
    patchesFile.print(`import ${generateImportSymbols(path, symbols)} from "${path}${importExtension}";`);
  });
  patchesFile.print("");
  patchesFile.print(`const p = {\n${indent(patches.join(",\n"))}\n};\n`);
  patchesFile.print(`export const patches = p;`);
  patchesFile.print("");
  patchesFile.print(`export const typeOverrides = ${JSON.stringify(typeOverrides, null, 2)} as const;`);

  const patchesTypeFileName = fileName.replace("CustomTypePatches", "PatchMessage");
  const patchTypeFile = schema.generateFile(patchesTypeFileName);
  patchTypeFile.print(`import { patches } from "./${fileName}";`);
  patchTypeFile.print(`import type { MessageDesc } from "../../sdk/client/types.ts";`);
  patchTypeFile.print(`export const patched = <T extends MessageDesc>(messageDesc: T): T => {`);
  patchTypeFile.print(`  const patchMessage = patches[messageDesc.$type as keyof typeof patches] as any;`);
  patchTypeFile.print(`  if (!patchMessage) return messageDesc;`);
  patchTypeFile.print(`  const wrapped: T = {`);
  patchTypeFile.print(`    ...messageDesc,`);
  patchTypeFile.print(`    encode(message, writer) {`);
  patchTypeFile.print(`      return messageDesc.encode(patchMessage(message, 'encode'), writer);`);
  patchTypeFile.print(`    },`);
  patchTypeFile.print(`    decode(input, length) {`);
  patchTypeFile.print(`      return patchMessage(messageDesc.decode(input, length), 'decode');`);
  patchTypeFile.print(`    },`);
  patchTypeFile.print(`  };`);
  patchTypeFile.print(`  // fromJSON / toJSON / fromPartial are overridden only for representation-`);
  patchTypeFile.print(`  // changing messages, whose patch entry carries these normalizers. Plain`);
  patchTypeFile.print(`  // function patch entries (Dec/DecCoin) leave the generated behavior intact.`);
  patchTypeFile.print(`  // (assigned through an \`any\` view since \`wrapped\` has the generic type \`T\`.)`);
  patchTypeFile.print(`  const overrides = wrapped as any;`);
  patchTypeFile.print(`  if (patchMessage.fromJSON) {`);
  patchTypeFile.print(`    overrides.fromJSON = (object: any) => {`);
  patchTypeFile.print(`      // patchMessage.fromJSON normalizes any accepted input (base64, decimal`);
  patchTypeFile.print(`      // string, number, bigint) to the JS value. The generated fromJSON only`);
  patchTypeFile.print(`      // groks the wire-JSON (base64) form, so feed it a sanitized clone (custom`);
  patchTypeFile.print(`      // fields re-serialized) to fill non-custom fields, then overlay the JS values.`);
  patchTypeFile.print(`      const patchedFields = patchMessage.fromJSON(object);`);
  patchTypeFile.print(`      return { ...(messageDesc.fromJSON({ ...object, ...patchMessage.toJSON(patchedFields) }) as object), ...patchedFields };`);
  patchTypeFile.print(`    };`);
  patchTypeFile.print(`  }`);
  patchTypeFile.print(`  if (patchMessage.toJSON) {`);
  patchTypeFile.print(`    overrides.toJSON = (message: any) => ({ ...(messageDesc.toJSON(patchMessage(message, 'encode')) as object), ...patchMessage.toJSON(message) });`);
  patchTypeFile.print(`  }`);
  patchTypeFile.print(`  if (patchMessage.fromPartial) {`);
  patchTypeFile.print(`    overrides.fromPartial = (object: any) => patchMessage.fromPartial(messageDesc.fromPartial(object));`);
  patchTypeFile.print(`  }`);
  patchTypeFile.print(`  return wrapped;`);
  patchTypeFile.print(`};`);

  const testsFile = schema.generateFile(fileName.replace(/\.ts$/, ".spec.ts"));
  generateTests(basename(fileName), testsFile, messageToCustomFields);
}

function getOutputFileName(schema: Schema): string {
  if (process.env.PROTO_SOURCE) {
    return `${process.env.PROTO_SOURCE}CustomTypePatches.ts`;
  }

  if (process.env.BUF_PLUGIN_CUSTOMTYPE_TYPES_PATCHES_OUTPUT_FILE) {
    return process.env.BUF_PLUGIN_CUSTOMTYPE_TYPES_PATCHES_OUTPUT_FILE;
  }

  for (const file of schema.files) {
    if (file.name.includes("akash/provider/lease")) {
      return "providerCustomTypePatches.ts";
    }
    if (file.name.includes("akash/cert/v1/msg")) {
      return "nodeCustomTypePatches.ts";
    }
    if (file.name.includes("cosmos/base/tendermint/v1beta1/query") || file.name.includes("cosmos/base/query/v1/query")) {
      return "cosmosCustomTypePatches.ts";
    }
  }

  throw new Error("Cannot determine file name for custom patches");
}

const indent = (value: string, tab = "  ") => tab + value.replace(/\n/g, "\n" + tab);

function generateNestedFieldUpdate(field: DescField, vars: VarNames) {
  const fieldRef = `${vars.value}.${field.localName}`;
  const newValueRef = `${vars.newValue}.${field.localName}`;
  if (field.fieldKind === "list") {
    return `if (${fieldRef}) ${newValueRef} = ${fieldRef}.map((item) => ${vars.fn}(item, transformType)!);`;
  }

  if (field.fieldKind === "map") {
    return `if (${fieldRef}) {\n`
      + `  ${newValueRef} = {};\n`
      + `  Object.keys(${fieldRef}).forEach(k => ${newValueRef}[k] = ${vars.fn}(${fieldRef}[k], transformType)!);\n`
      + `}`;
  }

  if (field.oneof && field.message) {
    const oneofValueRef = `${vars.value}.${field.oneof.localName}`;
    return `if (${oneofValueRef}?.case === "${field.localName}") {\n`
      + `  ${newValueRef} = {\n`
      + `    ...${oneofValueRef},\n`
      + `    value: ${vars.fn}(${oneofValueRef}.value, transformType)\n`
      + `  };\n`
      + `}`;
  }

  return `if (${fieldRef} != null) ${newValueRef} = ${vars.fn}(${fieldRef}, transformType);`;
}

function generateFieldTransformation(field: DescField, customType: CustomType<unknown, unknown>, vars: Omit<VarNames, "fn">) {
  const valueRef = `${vars.value}.${field.localName}`;
  const newValueRef = `${vars.newValue}.${field.localName}`;
  const fn = `${customType.shortName}[transformType]`;

  if (field.scalar !== ScalarType.BYTES) {
    return `if (${valueRef} != null) ${newValueRef} = ${fn}(${valueRef});`;
  }

  // Representation-changing bytes types (e.g. Int: wire bytes <-> bigint) need
  // asymmetric handling: on encode the JS value (bigint) is stringified then
  // written as bytes; on decode the bytes are read as a string then parsed.
  if (customType.jsType) {
    return `if (${valueRef} != null) ${newValueRef} = transformType === 'encode'`
      + ` ? encodeBinary(${customType.shortName}.encode(${valueRef}))`
      + ` : ${customType.shortName}.decode(decodeBinary(${valueRef}));`;
  }

  // Representation-preserving bytes types (e.g. Dec) round-trip bytes<->bytes.
  return `if (${valueRef} != null) ${newValueRef} = encodeBinary(${fn}(decodeBinary(${valueRef})));`;
}

interface VarNames {
  fn: string;
  value: string;
  newValue: string;
}

const dirnameToVar = (path: string) => path.replace(/\.+\//g, "_").replace(/\//g, "_").replace(/_pb$/, "");
function generateImportSymbols(path: string, symbols: Set<string>): string {
  if (symbols.has("type *")) return `type * as ${dirnameToVar(path)}`;
  if (symbols.has("*")) return `* as ${dirnameToVar(path)}`;
  return `{ ${Array.from(symbols).join(", ")} }`;
}

function generateTests(fileName: string, testsFile: GeneratedFile, messageToCustomFields: Map<DescMessage, Set<DescField>>) {
  testsFile.print(`import { expect, describe, it } from "vitest";`);
  testsFile.print(`import { patches } from "./${basename(fileName)}";`);
  testsFile.print(`import { generateMessage, type MessageSchema } from "@test/helpers/generateMessage.ts";`);
  testsFile.print(`import type { TypePatches } from "../../sdk/client/types.ts";`);
  testsFile.print("");
  testsFile.print(`const messageTypes: Record<string, MessageSchema> = {`);
  for (const [message, fields] of messageToCustomFields.entries()) {
    testsFile.print(`  "${message.typeName}": {`);
    testsFile.print(`    type: `, testsFile.import(message.name, `${PROTO_PATH}/${message.file.name}.ts`), `,`);
    testsFile.print(`    fields: [`, ...Array.from(fields, (f) => serializeField(f, testsFile)), `],`);
    testsFile.print(`  },`);
  }
  testsFile.print(`};`);
  testsFile.print(`describe("${fileName}", () => {`);
  testsFile.print(`  describe.each(Object.entries(patches))('patch %s', (typeName, patch: TypePatches[keyof TypePatches]) => {`);
  testsFile.print(`    it('returns undefined if receives null or undefined', () => {`);
  testsFile.print(`      expect(patch(null, 'encode')).toBe(undefined);`);
  testsFile.print(`      expect(patch(null, 'decode')).toBe(undefined);`);
  testsFile.print(`      expect(patch(undefined, 'encode')).toBe(undefined);`);
  testsFile.print(`      expect(patch(undefined, 'decode')).toBe(undefined);`);
  testsFile.print(`    });`);
  testsFile.print("");
  testsFile.print(`    it.each(generateTestCases(typeName, messageTypes))('patches and returns cloned value: %s', (name, value) => {`);
  testsFile.print(`      const transformedValue = patch(patch(value, 'encode'), 'decode');`);
  testsFile.print(`      expect(value).toEqual(transformedValue);`);
  testsFile.print(`      expect(value).not.toBe(transformedValue);`);
  testsFile.print(`    });`);
  testsFile.print(`  });`);
  testsFile.print("");
  testsFile.print(`  function generateTestCases(typeName: string, messageTypes: Record<string, MessageSchema>) {`);
  testsFile.print(`    const type = messageTypes[typeName];`);
  testsFile.print(`    const cases = type.fields.map((field) => ["single " + field.name + " field", generateMessage(typeName, {`);
  testsFile.print(`      ...messageTypes,`);
  testsFile.print(`      [typeName]: {`);
  testsFile.print(`        ...type,`);
  testsFile.print(`        fields: [field],`);
  testsFile.print(`      }`);
  testsFile.print(`    })]);`);
  testsFile.print(`    cases.push(["all fields", generateMessage(typeName, messageTypes)]);`);
  testsFile.print(`    return cases;`);
  testsFile.print(`  }`);
  testsFile.print("});");
}

function serializeField(f: DescField, file: GeneratedFile): Printable {
  const field: Printable[] = [
    `{`,
    `name: "${f.localName}",`,
    `kind: "${f.fieldKind}",`,
  ];
  if (f.fieldKind === "scalar") {
    field.push(`scalarType: ${f.scalar},`);
  }
  if (f.fieldKind === "enum") {
    field.push(`enum: `, JSON.stringify(f.enum.values.map((v) => v.localName)), `,`);
  }
  if (getCustomType(f)) {
    field.push(`customType: "${getCustomType(f)!.shortName}",`);
  }
  if (f.fieldKind === "map") {
    field.push(`mapKeyType: ${f.mapKey},`);
  }
  if (f.message) {
    field.push(`message: {fields: [`,
      ...f.message.fields.map((nf) => serializeField(nf, file)),
      `],`,
      `type: `, file.import(f.message.name, `${PROTO_PATH}/${f.message.file.name}.ts`),
      `},`,
    );
  }
  field.push(`},`);
  return field;
}
