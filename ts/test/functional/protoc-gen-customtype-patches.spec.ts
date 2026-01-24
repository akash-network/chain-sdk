import { afterEach, describe, expect, it } from "@jest/globals";
import { exec } from "child_process";
import { access, constants as fsConst, readFile, rmdir } from "fs/promises";
import { tmpdir } from "os";
import { join as joinPath } from "path";
import type { PluginOptions } from "../../script/protoc-gen-customtype-patches.ts";
import { promisify } from "util";

const execAsync = promisify(exec);

describe("protoc-gen-customtype-patches plugin", () => {
  const outputDir = joinPath(tmpdir(), `ts-bufplugin-${process.pid.toString()}`);
  const protoDir = "./ts/test/functional/proto";

  afterEach(async () => {
    if (await access(outputDir, fsConst.W_OK).then(() => true, () => false)) {
      await rmdir(outputDir, { recursive: true });
    }
  });

  describe('when patch_whole_tree is true', () => {
    it("generates `Set` instance with all the types that have reference to fields with custom type option", async () => {
      const command = [
        `buf generate`,
        `--config '${JSON.stringify({
          version: "v2",
          modules: [
            { path: "go/vendor/github.com/cosmos/gogoproto" },
            { path: "./ts/test/functional/proto" },
          ],
        })}'`,
        `--template '${JSON.stringify(createBufGenerateConfig({ patchWholeTree: true }))}'`,
        `-o '${outputDir}'`,
        `--path ${protoDir}/customtype.proto`,
        protoDir,
      ].join(" ");

      await execAsync(command, {
        cwd: joinPath(__dirname, "..", "..", ".."),
        env: {
          ...process.env,
          BUF_PLUGIN_CUSTOMTYPE_TYPES_PATCHES_OUTPUT_FILE: "customPatches.ts",
        },
      });

      expect(await readFile(joinPath(outputDir, "customPatches.ts"), "utf-8")).toMatchSnapshot();
    });
  });

  describe('when patch_whole_tree is false', () => {
    it("generates `Set` instance with all the leaf types that have reference to fields with custom type option", async () => {
      const command = [
        `buf generate`,
        `--config '${JSON.stringify({
          version: "v2",
          modules: [
            { path: "go/vendor/github.com/cosmos/gogoproto" },
            { path: "./ts/test/functional/proto" },
          ],
        })}'`,
        `--template '${JSON.stringify(createBufGenerateConfig({ patchWholeTree: false }))}'`,
        `-o '${outputDir}'`,
        `--path ${protoDir}/customtype.proto`,
        protoDir,
      ].join(" ");

      await execAsync(command, {
        cwd: joinPath(__dirname, "..", "..", ".."),
        env: {
          ...process.env,
          BUF_PLUGIN_CUSTOMTYPE_TYPES_PATCHES_OUTPUT_FILE: "customPatches.ts",
        },
      });

      expect(await readFile(joinPath(outputDir, "customPatches.ts"), "utf-8")).toMatchSnapshot();
    });
  });

  function createBufGenerateConfig(options: PluginOptions) {
    return {
      version: "v2",
      clean: true,
      plugins: [
        {
          local: "ts/script/protoc-gen-customtype-patches.ts",
          strategy: "all",
          out: ".",
          opt: [
            "target=ts",
            "import_extension=ts",
            `patch_whole_tree=${options.patchWholeTree}`,
          ],
        },
      ],
    };
  }
});
