import { describe, expect, it } from "@jest/globals";
import { exec } from "child_process";
import { access, constants as fsConst, readFile, rmdir } from "fs/promises";
import { tmpdir } from "os";
import { join as joinPath } from "path";
import { promisify } from "util";

const execAsync = promisify(exec);

async function vendorExists(vendorPath: string): Promise<boolean> {
  try {
    await access(vendorPath, fsConst.F_OK);
    return true;
  } catch {
    return false;
  }
}

describe("protoc-sdk-objec plugin", () => {
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

  it("generates SDK object from proto files", async () => {
    const repoRoot = joinPath(__dirname, "..", "..", "..");
    const cosmosSdkVendor = joinPath(repoRoot, "go/vendor/github.com/cosmos/cosmos-sdk/proto");
    
    const outputDir = joinPath(tmpdir(), `ts-bufplugin-${process.pid.toString()}`);
    const protoDir = "./ts/test/functional/proto";
    
    let bufConfig;
    if (await vendorExists(cosmosSdkVendor)) {
      bufConfig = {
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
    } else {
      console.log("Skipping test - vendor directory not found. Run 'make modvendor' to set up vendor.");
      return;
    }
    const command = [
      `npx --package=@bufbuild/buf buf generate`,
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
