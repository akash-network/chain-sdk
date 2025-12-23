import { describe, expect, it } from "@jest/globals";
import { exec } from "child_process";
import { existsSync } from "fs";
import { access, constants as fsConst, readFile, rmdir } from "fs/promises";
import { tmpdir } from "os";
import { join as joinPath } from "path";
import { promisify } from "util";

const execAsync = promisify(exec);

describe("protoc-gen-customtype-patches plugin", () => {
  const config = {
    version: "v2",
    clean: true,
      plugins: [
        {
          local: "ts/script/protoc-gen-customtype-patches-wrapper.sh",
          strategy: "all",
          out: ".",
          opt: [
            "target=ts",
            "import_extension=ts"
          ],
        },
      ],
  };

  const repoRoot = joinPath(__dirname, "..", "..", "..");
  const gogoprotoVendor = joinPath(repoRoot, "go/vendor/github.com/cosmos/gogoproto");
  const bufBin = process.env.AKASH_DEVCACHE_BIN 
    ? joinPath(process.env.AKASH_DEVCACHE_BIN, "buf")
    : null;

  it("generates `Set` instance with all the types that have reference to fields with custom type option", async () => {
    if (!existsSync(gogoprotoVendor)) {
      throw new Error(`Go vendor missing at ${gogoprotoVendor}. Run 'make modvendor' from repo root.`);
    }
    if (!bufBin || !existsSync(bufBin)) {
      throw new Error(`buf binary missing at ${bufBin}. AKASH_DEVCACHE_BIN=${process.env.AKASH_DEVCACHE_BIN}`);
    }
    const outputDir = joinPath(tmpdir(), `ts-bufplugin-${process.pid.toString()}`);
    const protoDir = "./ts/test/functional/proto";
    
    const bufConfig = {
      version: "v2",
      modules: [
        { path: "go/vendor/github.com/cosmos/gogoproto" },
        { path: protoDir },
      ],
      deps: [
        "buf.build/googleapis/googleapis",
        "buf.build/protocolbuffers/wellknowntypes",
      ],
    };
    
    const command = [
      `${bufBin} generate`,
      `--config '${JSON.stringify(bufConfig)}'`,
      `--template '${JSON.stringify(config)}'`,
      `-o '${outputDir}'`,
      `--path ${protoDir}/customtype.proto`,
      protoDir,
    ].join(" ");

    try {
      await execAsync(command, {
        cwd: joinPath(__dirname, "..", "..", ".."),
        env: {
          ...process.env,
          BUF_PLUGIN_CUSTOMTYPE_TYPES_PATCHES_OUTPUT_FILE: "customPatches.ts",
        },
      });

      expect(await readFile(joinPath(outputDir, "customPatches.ts"), "utf-8")).toMatchSnapshot();
    } finally {
      if (await access(outputDir, fsConst.W_OK).catch(() => false)) {
        await rmdir(outputDir, { recursive: true });
      }
    }
  });
});
