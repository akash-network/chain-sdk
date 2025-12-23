import { describe, expect, it } from "@jest/globals";
import { exec } from "child_process";
import { existsSync } from "fs";
import { access, constants as fsConst, readFile, rmdir } from "fs/promises";
import { tmpdir } from "os";
import { join as joinPath } from "path";
import { promisify } from "util";

const execAsync = promisify(exec);

describe("protoc-sdk-object plugin", () => {
  const config = {
    version: "v2",
    clean: true,
      plugins: [
        {
          local: "ts/script/protoc-gen-sdk-object-wrapper.sh",
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
  const cosmosSdkVendor = joinPath(repoRoot, "go/vendor/github.com/cosmos/cosmos-sdk/proto");
  const bufBin = process.env.AKASH_DEVCACHE_BIN 
    ? joinPath(process.env.AKASH_DEVCACHE_BIN, "buf")
    : null;
  
  const hasVendor = existsSync(cosmosSdkVendor);
  const hasBuf = bufBin ? existsSync(bufBin) : false;
  const canRun = hasVendor && hasBuf;

  (canRun ? it : it.skip)("generates SDK object from proto files", async () => {
    const outputDir = joinPath(tmpdir(), `ts-bufplugin-${process.pid.toString()}`);
    const protoDir = "./ts/test/functional/proto";
    
    const bufConfig = {
      version: "v2",
      modules: [
        { path: "go/vendor/github.com/cosmos/cosmos-sdk/proto" },
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
      `--path ${protoDir}/msg.proto`,
      `--path ${protoDir}/query.proto`,
      protoDir,
    ].join(" ");

    try {
      await execAsync(command, {
        cwd: joinPath(__dirname, "..", "..", ".."),
        env: {
          ...process.env,
          BUF_PLUGIN_SDK_OBJECT_OUTPUT_FILE: "sdk.ts",
          NODE_OPTIONS: "--experimental-strip-types --no-warnings",
        },
      });

      expect(await readFile(joinPath(outputDir, "sdk.ts"), "utf-8")).toMatchSnapshot();
      expect(await readFile(joinPath(outputDir, "protos", "msg_akash.ts"), "utf-8")).toMatchSnapshot();
      expect(await readFile(joinPath(outputDir, "protos", "query_akash.ts"), "utf-8")).toMatchSnapshot();
    } finally {
      if (await access(outputDir, fsConst.W_OK).catch(() => false)) {
        await rmdir(outputDir, { recursive: true });
      }
    }
  });
});
