import { spawn } from "child_process";
import { existsSync } from "fs";
import * as path from "path";

export interface MockServer {
  gatewayUrl: string;
  grpcAddr: string;
  stop: () => Promise<void>;
}

export async function startMockServer(dataDir: string): Promise<MockServer> {
  const projectRoot = path.resolve(__dirname, "../../..");
  const absoluteDataDir = path.isAbsolute(dataDir) 
    ? dataDir 
    : path.resolve(projectRoot, dataDir);

  const mockServerBin = process.env.MOCK_SERVER_BIN;
  
  let command: string;
  let args: string[];
  let cwd: string;

  if (mockServerBin && existsSync(mockServerBin)) {
    command = mockServerBin;
    args = ["--data-dir", absoluteDataDir];
    cwd = projectRoot;
  } else {
    const goDir = path.join(projectRoot, "go");
    const vendorDir = path.join(goDir, "vendor");
    const modFlag = existsSync(vendorDir) ? "-mod=vendor" : "-mod=readonly";
    
    command = "go";
    args = ["run", modFlag, "testutil/mock/cmd/server/main.go", "--data-dir", absoluteDataDir];
    cwd = goDir;
  }

  const proc = spawn(command, args, {
    stdio: ["ignore", "pipe", "pipe"],
    cwd,
    env: { ...process.env, GOWORK: "off" },
    detached: false,
    killSignal: "SIGTERM",
  });

  let gatewayUrl = "";
  let grpcAddr = "";
  let outputBuffer = "";
  let spawnError: Error | null = null;

  proc.stdout?.on("data", (data: Buffer) => {
    const output = data.toString();
    outputBuffer += output;
    const lines = outputBuffer.split("\n");
    outputBuffer = lines.pop() || "";
    
    for (const line of lines) {
      const trimmed = line.trim();
      const gatewayMatch = trimmed.match(/gateway:\s*(http:\/\/[^\s]+)/i);
      if (gatewayMatch) {
        gatewayUrl = gatewayMatch[1];
      }
      
      const grpcMatch = trimmed.match(/grpc:\s*([^\s]+)/i);
      if (grpcMatch) {
        grpcAddr = grpcMatch[1];
      }
    }
  });

  let stderrBuffer = "";
  proc.stderr?.on("data", (data: Buffer) => {
    const output = data.toString();
    stderrBuffer += output;
    // Only log actual errors, not debug output
    const errorKeywords = ["error", "fail", "panic", "fatal", "exception", "cannot", "unable"];
    if (errorKeywords.some(keyword => output.toLowerCase().includes(keyword))) {
      console.error(`[mock-server] ${output}`);
    }
  });

  proc.on("error", (err) => {
    spawnError = new Error(`Failed to start mock server: ${err.message}`);
  });

  proc.on("exit", (code, signal) => {
    if (code !== null && code !== 0 && code !== 143) { // 143 is SIGTERM
      console.error(`[mock-server] Process exited with code ${code}, signal ${signal}`);
      if (stderrBuffer) {
        console.error(`[mock-server] stderr output:\n${stderrBuffer}`);
      }
    }
  });

  let cleanupOnError = true;
  
  try {
    for (let i = 0; i < 1200; i++) {
      await new Promise(resolve => setTimeout(resolve, 100));
      
      if (spawnError) {
        throw spawnError;
      }
      
      if (proc.exitCode !== null && proc.exitCode !== 0) {
        const errorMsg = stderrBuffer || outputBuffer || "Unknown error";
        throw new Error(`Mock server failed to start (exit code ${proc.exitCode}): ${errorMsg}`);
      }
      
      if (gatewayUrl && grpcAddr) {
        cleanupOnError = false;
        return {
          gatewayUrl,
          grpcAddr,
          stop: async () => {
            if (proc.killed || proc.exitCode !== null) {
              if (proc.stdout && !proc.stdout.destroyed) {
                proc.stdout.destroy();
              }
              if (proc.stderr && !proc.stderr.destroyed) {
                proc.stderr.destroy();
              }
              proc.removeAllListeners();
              return;
            }
            
            proc.removeAllListeners();
            
            if (proc.stdout && !proc.stdout.destroyed) {
              proc.stdout.destroy();
            }
            if (proc.stderr && !proc.stderr.destroyed) {
              proc.stderr.destroy();
            }
            
            return new Promise<void>((resolve) => {
              if (proc.exitCode !== null) {
                resolve();
                return;
              }
              
              let resolved = false;
              const doResolve = () => {
                if (resolved) return;
                resolved = true;
                if (proc.stdout && !proc.stdout.destroyed) {
                  proc.stdout.destroy();
                }
                if (proc.stderr && !proc.stderr.destroyed) {
                  proc.stderr.destroy();
                }
                resolve();
              };
              
              const timeout = setTimeout(() => {
                if (!proc.killed && proc.exitCode === null) {
                  try {
                    proc.kill("SIGKILL");
                  } catch (e) {
                    // Process might already be dead, ignore
                  }
                }
                doResolve();
              }, 500);
              
              const onExit = () => {
                clearTimeout(timeout);
                doResolve();
              };
              
              proc.once("exit", onExit);
              
              try {
                proc.kill("SIGTERM");
              } catch (e) {
                clearTimeout(timeout);
                proc.removeListener("exit", onExit);
                doResolve();
              }
            });
          },
        };
      }
    }
    
    if (spawnError) {
      throw spawnError;
    }
    
    const errorMsg = stderrBuffer || outputBuffer || "No error output captured";
    throw new Error(`Mock server failed to start: timeout waiting for addresses. Last output: ${errorMsg}`);
  } finally {
    if (cleanupOnError) {
      if (!proc.killed && proc.exitCode === null) {
        try {
          proc.kill("SIGTERM");
        } catch {
          // Process already exited, ignore
        }
      }
      
      proc.removeAllListeners();
      
      if (proc.stdout && !proc.stdout.destroyed) {
        proc.stdout.destroy();
      }
      if (proc.stderr && !proc.stderr.destroyed) {
        proc.stderr.destroy();
      }
    }
  }
}

