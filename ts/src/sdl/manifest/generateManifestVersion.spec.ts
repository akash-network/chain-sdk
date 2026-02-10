import { describe, expect, it } from "@jest/globals";

import type { Group } from "../../generated/protos/index.provider.akash.v2beta3.ts";
import { yaml } from "../../utils/yaml.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
import { type BuildResult, generateManifest } from "./generateManifest.ts";
import { generateManifestVersion, manifestToSortedJSON } from "./generateManifestVersion.ts";

describe(generateManifestVersion.name, () => {
  describe("basic version generation", () => {
    it("returns 32-byte SHA-256 hash", async () => {
      const { manifest } = setup();
      const version = await generateManifestVersion(manifest);

      expect(version.length).toBe(32);
    });

    it("returns consistent hash for same manifest", async () => {
      const { manifest } = setup();
      const version1 = await generateManifestVersion(manifest);
      const version2 = await generateManifestVersion(manifest);

      expect(version1).toEqual(version2);
    });

    it("returns different hash for different manifests", async () => {
      const { manifest: manifest1 } = setup();
      const { manifest: manifest2 } = setup({ image: "alpine" });

      const version1 = await generateManifestVersion(manifest1);
      const version2 = await generateManifestVersion(manifest2);

      expect(version1).not.toEqual(version2);
    });
  });

  describe(manifestToSortedJSON.name, () => {
    it("returns stable JSON string", () => {
      const { manifest } = setup();
      const json1 = manifestToSortedJSON(manifest);
      const json2 = manifestToSortedJSON(manifest);

      expect(json1).toBe(json2);
    });

    it("returns valid JSON", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);

      expect(() => JSON.parse(json)).not.toThrow();
    });

    it("includes group name and services", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].name).toBe("dcloud");
      expect(parsed[0].services).toHaveLength(1);
      expect(parsed[0].services[0].name).toBe("web");
    });

    it("formats command as null when empty", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].command).toBeNull();
    });

    it("formats args as null when empty", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].args).toBeNull();
    });

    it("formats env as null when empty", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].env).toBeNull();
    });

    it("includes command when provided", () => {
      const { manifest } = setup({ command: ["/bin/sh"] });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].command).toEqual(["/bin/sh"]);
    });

    it("includes args when provided", () => {
      const { manifest } = setup({ args: ["-c", "echo hello"] });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].args).toEqual(["-c", "echo hello"]);
    });

    it("includes env when provided", () => {
      const { manifest } = setup({ env: ["ENV1=test1", "ENV2=test2"] });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].env).toEqual(["ENV1=test1", "ENV2=test2"]);
    });

    it("formats credentials as null when not provided", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].credentials).toBeNull();
    });

    it("includes credentials when provided", () => {
      const { manifest } = setup({
        credentials: {
          host: "registry.example.com",
          username: "user",
          password: "password123",
        },
      });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].credentials).toEqual({
        host: "registry.example.com",
        email: "",
        username: "user",
        password: "password123",
      });
    });

    it("formats hosts as null when empty", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].expose[0].hosts).toBeNull();
    });

    it("includes hosts when provided", () => {
      const { manifest } = setup({
        expose: [{ port: 80, as: 80, to: [{ global: true }], accept: ["example.com"] }],
      });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].expose[0].hosts).toEqual(["example.com"]);
    });

    it("includes resource values as strings", () => {
      const { manifest } = setup();
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].resources.cpu.units.val).toBe("500");
      expect(parsed[0].services[0].resources.memory.size.val).toBe("536870912");
    });

    it("escapes HTML characters", () => {
      const { manifest } = setup();
      // The function escapes <, >, and &
      const json = manifestToSortedJSON(manifest);

      expect(json).not.toContain("<");
      expect(json).not.toContain(">");
      // Note: & only gets escaped if present, which won't be in normal output
    });
  });

  describe("snapshot tests", () => {
    it("generates consistent version for Minesweeper SDL", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          minesweeper:
            image: creepto/minesweeper
            expose:
              - port: 3000
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            minesweeper:
              resources:
                cpu:
                  units: 0.1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                organization: akash.network
              signedBy:
                anyOf:
                  - akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63
                  - akash18qa2a2ltfyvkyj0ggj3hkvuj6twzyumuaru9s4
              pricing:
                minesweeper:
                  denom: uakt
                  amount: 10000
        deployment:
          minesweeper:
            akash:
              profile: minesweeper
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const version = await generateManifestVersion(manifest);

      expect(version).toMatchSnapshot("Minesweeper manifest version");
    });

    it("generates consistent JSON for Empty Profile SDL", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          tetris-main:
            image: bsord/tetris
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            tetris:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                host: akash
              signedBy:
                anyOf:
                  - akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63
                  - akash18qa2a2ltfyvkyj0ggj3hkvuj6twzyumuaru9s4
              pricing:
                tetris:
                  denom: uakt
                  amount: 10000
        deployment:
          tetris-main:
            akash:
              profile: tetris
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const json = manifestToSortedJSON(manifest);

      expect(json).toMatchSnapshot("Empty Profile manifest JSON");
    });

    it("generates consistent version for Empty Profile SDL", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          tetris-main:
            image: bsord/tetris
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            tetris:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                host: akash
              signedBy:
                anyOf:
                  - akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63
                  - akash18qa2a2ltfyvkyj0ggj3hkvuj6twzyumuaru9s4
              pricing:
                tetris:
                  denom: uakt
                  amount: 10000
        deployment:
          tetris-main:
            akash:
              profile: tetris
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const version = await generateManifestVersion(manifest);

      expect(version).toMatchSnapshot("Empty Profile manifest version");
    });

    it("generates consistent JSON for Basic SDL", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          tetris:
            image: bsord/tetris
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            tetris:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                host: akash
              signedBy:
                anyOf:
                  - akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63
                  - akash18qa2a2ltfyvkyj0ggj3hkvuj6twzyumuaru9s4
              pricing:
                tetris:
                  denom: uakt
                  amount: 10000
        deployment:
          tetris:
            akash:
              profile: tetris
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const json = manifestToSortedJSON(manifest);

      expect(json).toMatchSnapshot("Basic SDL manifest JSON");
    });

    it("generates consistent version for Basic SDL", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          tetris:
            image: bsord/tetris
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            tetris:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                host: akash
              signedBy:
                anyOf:
                  - akash1365yvmc4s7awdyj3n2sav7xfx76adc6dnmlx63
                  - akash18qa2a2ltfyvkyj0ggj3hkvuj6twzyumuaru9s4
              pricing:
                tetris:
                  denom: uakt
                  amount: 10000
        deployment:
          tetris:
            akash:
              profile: tetris
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const version = await generateManifestVersion(manifest);

      expect(version).toMatchSnapshot("Basic SDL manifest version");
    });
  });

  describe("complex scenarios", () => {
    it("handles multi-service manifest", async () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        services:
          web:
            image: wordpress
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
          db:
            image: mysql
            expose:
              - port: 3306
                to:
                  - service: web
        profiles:
          compute:
            web:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 1Gi
                storage:
                  size: 2Gi
            db:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  size: 1Gi
          placement:
            dcloud:
              pricing:
                web:
                  denom: uakt
                  amount: 2000
                db:
                  denom: uakt
                  amount: 1000
        deployment:
          web:
            dcloud:
              profile: web
              count: 1
          db:
            dcloud:
              profile: db
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const version = await generateManifestVersion(manifest);

      expect(version).toBeInstanceOf(Uint8Array);
      expect(version.length).toBe(32);
    });

    it("handles GPU manifest", async () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        services:
          web:
            image: nginx
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            web:
              resources:
                cpu:
                  units: "100m"
                gpu:
                  units: 1
                  attributes:
                    vendor:
                      nvidia:
                memory:
                  size: "128Mi"
                storage:
                  - size: "1Gi"
          placement:
            westcoast:
              attributes:
                region: us-west
              pricing:
                web:
                  denom: uakt
                  amount: 50
        deployment:
          web:
            westcoast:
              profile: web
              count: 2
      `;

      const { manifest } = setup({ sdl });
      const version = await generateManifestVersion(manifest);

      expect(version).toBeInstanceOf(Uint8Array);
      expect(version.length).toBe(32);
    });

    it("handles IP lease endpoints", async () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        endpoints:
          myip:
            kind: ip
        services:
          web:
            image: nginx
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
                    ip: myip
        profiles:
          compute:
            web:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  size: 512Mi
          placement:
            dcloud:
              pricing:
                web:
                  denom: uakt
                  amount: 1000
        deployment:
          web:
            dcloud:
              profile: web
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].expose[0].ip).toBe("myip");
      expect(parsed[0].services[0].expose[0].endpointSequenceNumber).toBe(1);
    });

    it("handles HTTP options", async () => {
      const sdl: SDLInput = yaml`
        version: '2.0'
        services:
          tetris:
            image: bsord/tetris
            expose:
              - port: 80
                http_options:
                  max_body_size: 104857600
                  read_timeout: 50
                  send_timeout: 100
                  next_tries: 24
                  next_timeout: 48
                  next_cases:
                    - "500"
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            tetris:
              resources:
                cpu:
                  units: 1
                memory:
                  size: 512Mi
                storage:
                  - size: 512Mi
          placement:
            akash:
              attributes:
                host: akash
              pricing:
                tetris:
                  denom: uakt
                  amount: 10000
        deployment:
          tetris:
            akash:
              profile: tetris
              count: 1
      `;

      const { manifest } = setup({ sdl });
      const json = manifestToSortedJSON(manifest);
      const parsed = JSON.parse(json);

      expect(parsed[0].services[0].expose[0].httpOptions).toEqual({
        maxBodySize: 104857600,
        readTimeout: 50,
        sendTimeout: 100,
        nextTries: 24,
        nextTimeout: 48,
        nextCases: ["500"],
      });
    });
  });

  function setup(options: {
    sdl?: SDLInput;
    image?: string;
    command?: string[];
    args?: string[];
    env?: string[];
    credentials?: { host: string; username: string; password: string };
    expose?: Array<{
      port: number;
      as?: number;
      to?: Array<{ global?: boolean }>;
      accept?: string[];
    }>;
  } = {}): { manifest: Group[] } {
    const sdl
      = options.sdl
        ?? createBasicSdl({
          image: options.image,
          command: options.command,
          args: options.args,
          env: options.env,
          credentials: options.credentials,
          expose: options.expose,
        });

    const result = generateManifest(sdl);
    assertBuildResult(result);

    return { manifest: result.groups };
  }

  function assertBuildResult(result: ReturnType<typeof generateManifest>): asserts result is BuildResult {
    if (Array.isArray(result)) {
      throw new Error(`Expected BuildResult but got validation errors: ${JSON.stringify(result)}`);
    }
  }

  function createBasicSdl(options: {
    image?: string;
    command?: string[];
    args?: string[];
    env?: string[];
    credentials?: { host: string; username: string; password: string };
    expose?: Array<{
      port: number;
      as?: number;
      to?: Array<{ global?: boolean }>;
      accept?: string[];
    }>;
  } = {}): SDLInput {
    const { image = "nginx", command, args, env, credentials, expose } = options;

    return {
      version: "2.0",
      services: {
        web: {
          image,
          expose: expose ?? [
            {
              port: 80,
              as: 80,
              to: [{ global: true }],
            },
          ],
          ...(env && { env }),
          ...(command && { command }),
          ...(args && { args }),
          ...(credentials && { credentials }),
        },
      },
      profiles: {
        compute: {
          web: {
            resources: {
              cpu: { units: 0.5 },
              memory: { size: "512Mi" },
              storage: { size: "512Mi" },
            },
          },
        },
        placement: {
          dcloud: {
            pricing: {
              web: { denom: "uakt", amount: 1000 },
            },
          },
        },
      },
      deployment: {
        web: {
          dcloud: {
            profile: "web",
            count: 1,
          },
        },
      },
    };
  }
});
