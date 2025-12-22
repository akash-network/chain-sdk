import { describe, expect, it } from "@jest/globals";
import { exec } from "child_process";
import { existsSync } from "fs";
import { access, constants as fsConst, readFile, rmdir } from "fs/promises";
import { tmpdir } from "os";
import { join as joinPath } from "path";
import { promisify } from "util";

const execAsync = promisify(exec);

const MIN_NODE_VERSION = "22.6.0";

function checkNodeVersion(): void {
  const currentVersion = process.version.slice(1);
  const [currentMajor, currentMinor] = currentVersion.split(".").map(Number);
  const [minMajor, minMinor] = MIN_NODE_VERSION.split(".").map(Number);
  
  if (currentMajor < minMajor || (currentMajor === minMajor && currentMinor < minMinor)) {
    throw new Error(
      `Node.js ${MIN_NODE_VERSION} or higher is required for --experimental-strip-types support. ` +
      `Current version: ${process.version}`
    );
  }
}

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
  const hasVendor = existsSync(cosmosSdkVendor);

  (hasVendor ? it : it.skip)("generates SDK object from proto files", async () => {
    checkNodeVersion();
    
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
      "npx --package=@bufbuild/buf buf generate",
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
