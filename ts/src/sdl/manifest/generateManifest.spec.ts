import fs from "node:fs";
import path from "node:path";

import { describe, expect, it } from "@jest/globals";

import { GroupSpec } from "../../generated/protos/index.akash.v1beta4.ts";
import { Group } from "../../generated/protos/index.provider.akash.v2beta3.ts";
import { yaml } from "../../utils/yaml.ts";
import type { SDLInput } from "../validateSDL/validateSDL.ts";
import type { GenerateManifestResult } from "./generateManifest.ts";
import { generateManifest } from "./generateManifest.ts";

describe(generateManifest.name, () => {
  describe("basic manifest generation", () => {
    it("generates manifest with correct structure", () => {
      const { result } = setup();

      expect(result.groups).toHaveLength(1);
      expect(result.groups[0].name).toBe("dcloud");
      expect(result.groups[0].services).toHaveLength(1);
      expect(result.groups[0].services[0].name).toBe("web");
      expect(result.groups[0].services[0].image).toBe("nginx");
    });

    it("generates groupSpecs with correct structure", () => {
      const { result } = setup();

      expect(result.groupSpecs).toHaveLength(1);
      expect(result.groupSpecs[0].name).toBe("dcloud");
      expect(result.groupSpecs[0].resources).toHaveLength(1);
      expect(result.groupSpecs[0].resources[0].count).toBe(1);
      expect(result.groupSpecs[0].resources[0].price).toEqual({
        denom: "uakt",
        amount: "1000",
      });
    });

    it("parses CPU units correctly", () => {
      const sdl = createBasicSdl({ cpu: 0.5 });
      const { result } = setup({ sdl });

      const cpuVal = result.groups[0].services[0].resources?.cpu?.units?.val;
      expect(new TextDecoder().decode(cpuVal)).toBe("500");
    });

    it("parses CPU units from string format", () => {
      const sdl = createBasicSdl({ cpu: "100m" });
      const { result } = setup({ sdl });

      const cpuVal = result.groups[0].services[0].resources?.cpu?.units?.val;
      expect(new TextDecoder().decode(cpuVal)).toBe("100");
    });

    it("parses memory size correctly", () => {
      const sdl = createBasicSdl({ memory: "512Mi" });
      const { result } = setup({ sdl });

      const memVal = result.groups[0].services[0].resources?.memory?.quantity?.val;
      expect(new TextDecoder().decode(memVal)).toBe("536870912");
    });

    it("parses storage correctly", () => {
      const sdl = createBasicSdl({ storage: [{ size: "1Gi" }] });
      const { result } = setup({ sdl });

      const storage = result.groups[0].services[0].resources?.storage;
      expect(storage).toHaveLength(1);
      expect(storage?.[0].name).toBe("default");
      const storageVal = storage?.[0].quantity?.val;
      expect(new TextDecoder().decode(storageVal)).toBe("1073741824");
    });

    it("parses storage params in ASC order", () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        services:
          web:
            image: nginx
            params:
              storage:
                data:
                  mount: /mnt/data
                  readOnly: true
                cache:
                  mount: /mnt/cache
        profiles:
          compute:
            web:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  - name: data
                    size: 1Gi
                  - name: cache
                    size: 512Mi
          placement:
            datacenter:
              pricing:
                web:
                  denom: uakt
                  amount: 200
        deployment:
          web:
            datacenter:
              profile: web
              count: 1
      `;
      const { result } = setup({ sdl });
      expect(result.groups[0].services[0].params?.storage).toEqual([
        expect.objectContaining({ name: "cache" }),
        expect.objectContaining({ name: "data" }),
      ]);
    });
  });

  describe("service features", () => {
    it("includes credentials when provided", () => {
      const credentials = {
        host: "registry.example.com",
        username: "user",
        password: "password123",
      };
      const sdl = createBasicSdl({ credentials });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].credentials).toMatchObject({
        host: "registry.example.com",
        username: "user",
        password: "password123",
        email: "",
      });
    });

    it("excludes credentials when not provided", () => {
      const sdl = createBasicSdl();
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].credentials).toBeUndefined();
    });

    it("includes environment variables", () => {
      const env = ["ENV1=value1", "ENV2=value2"];
      const sdl = createBasicSdl({ env });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].env).toEqual(env);
    });

    it("includes command and args", () => {
      const command = ["/bin/sh"];
      const args = ["-c", "echo hello"];
      const sdl = createBasicSdl({ command, args });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].command).toEqual(command);
      expect(result.groups[0].services[0].args).toEqual(args);
    });
  });

  describe("GPU configuration", () => {
    it("handles GPU with units=0", () => {
      const sdl = createBasicSdl({ gpu: { units: 0 } });
      const { result } = setup({ sdl });

      const gpuVal = result.groups[0].services[0].resources?.gpu?.units?.val;
      expect(new TextDecoder().decode(gpuVal)).toBe("0");
    });

    it("handles GPU with units and vendor attributes", () => {
      const sdl = createBasicSdl({
        gpu: {
          units: 1,
          attributes: {
            vendor: {
              nvidia: [{ model: "rtxa6000" }],
            },
          },
        },
      });
      const { result } = setup({ sdl });

      const gpuVal = result.groups[0].services[0].resources?.gpu?.units?.val;
      expect(new TextDecoder().decode(gpuVal)).toBe("1");
      expect(result.groups[0].services[0].resources?.gpu?.attributes).toContainEqual({
        key: "vendor/nvidia/model/rtxa6000",
        value: "true",
      });
    });

    it("handles GPU with ram and interface", () => {
      const sdl = createBasicSdl({
        gpu: {
          units: 1,
          attributes: {
            vendor: {
              nvidia: [{ model: "rtxa6000", ram: "48Gi", interface: "pcie" }],
            },
          },
        },
      });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].resources?.gpu?.attributes).toContainEqual({
        key: "vendor/nvidia/model/rtxa6000/ram/48Gi/interface/pcie",
        value: "true",
      });
    });
  });

  describe("storage configuration", () => {
    it("handles array of storage volumes", () => {
      const sdl = createBasicSdl({
        storage: [
          { name: "data", size: "1Gi" },
          { name: "cache", size: "512Mi" },
        ],
      });
      const { result } = setup({ sdl });

      const storage = result.groups[0].services[0].resources?.storage;
      expect(storage).toHaveLength(2);
      expect(storage?.[0].name).toBe("data");
      expect(storage?.[1].name).toBe("cache");
    });

    it("handles persistent storage attributes", () => {
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
            params:
              storage:
                data:
                  mount: /mnt/data
        profiles:
          compute:
            web:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  - name: data
                    size: 1Gi
                    attributes:
                      persistent: true
                      class: beta2
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
      const { result } = setup({ sdl });

      const storage = result.groups[0].services[0].resources?.storage;
      expect(storage?.[0].attributes).toContainEqual({ key: "persistent", value: "true" });
      expect(storage?.[0].attributes).toContainEqual({ key: "class", value: "beta2" });
    });
  });

  describe("placement and deployment", () => {
    it("includes placement attributes", () => {
      const sdl = createBasicSdl({
        placementAttributes: { region: "us-west", datacenter: "equinix" },
      });
      const { result } = setup({ sdl });

      expect(result.groupSpecs[0].requirements?.attributes).toContainEqual({
        key: "datacenter",
        value: "equinix",
      });
      expect(result.groupSpecs[0].requirements?.attributes).toContainEqual({
        key: "region",
        value: "us-west",
      });
    });

    it("includes signedBy requirements", () => {
      const sdl = createBasicSdl({
        signedBy: {
          anyOf: ["akash1abc", "akash1def"],
          allOf: ["akash1xyz"],
        },
      });
      const { result } = setup({ sdl });

      expect(result.groupSpecs[0].requirements?.signedBy?.anyOf).toEqual(["akash1abc", "akash1def"]);
      expect(result.groupSpecs[0].requirements?.signedBy?.allOf).toEqual(["akash1xyz"]);
    });

    it("sorts placements alphabetically", () => {
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
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  size: 512Mi
          placement:
            zebra:
              pricing:
                web:
                  denom: uakt
                  amount: 1000
            alpha:
              pricing:
                web:
                  denom: uakt
                  amount: 1000
        deployment:
          web:
            zebra:
              profile: web
              count: 1
            alpha:
              profile: web
              count: 1
      `;
      const { result } = setup({ sdl });

      expect(result.groups.map((g) => g.name)).toEqual(["alpha", "zebra"]);
      expect(result.groupSpecs.map((g) => g.name)).toEqual(["alpha", "zebra"]);
    });

    it("sorts services alphabetically within placement", () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        services:
          zebra:
            image: nginx
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
          alpha:
            image: nginx
            expose:
              - port: 81
                as: 81
                to:
                  - global: true
        profiles:
          compute:
            zebra:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  size: 512Mi
            alpha:
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
                zebra:
                  denom: uakt
                  amount: 1000
                alpha:
                  denom: uakt
                  amount: 1000
        deployment:
          zebra:
            dcloud:
              profile: zebra
              count: 1
          alpha:
            dcloud:
              profile: alpha
              count: 1
      `;
      const { result } = setup({ sdl });

      expect(result.groups[0].services.map((s) => s.name)).toEqual(["alpha", "zebra"]);
    });
  });

  describe("endpoints and expose", () => {
    it("handles global endpoint", () => {
      const { result } = setup();

      const expose = result.groups[0].services[0].expose;
      expect(expose).toHaveLength(1);
      expect(expose[0].port).toBe(80);
      expect(expose[0].externalPort).toBe(80);
      expect(expose[0].global).toBe(true);
      expect(expose[0].proto).toBe("TCP");
    });

    it("handles UDP protocol", () => {
      const sdl = createBasicSdl({
        expose: [{ port: 12345, as: 12345, proto: "udp", to: [{ global: true }] }],
      });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].expose[0].proto).toBe("UDP");
    });

    it("handles HTTP options", () => {
      const sdl = createBasicSdl({
        expose: [
          {
            port: 80,
            as: 80,
            to: [{ global: true }],
            http_options: {
              max_body_size: 104857600,
              read_timeout: 50,
              send_timeout: 100,
              next_tries: 24,
              next_timeout: 48,
              next_cases: ["500"],
            },
          },
        ],
      });
      const { result } = setup({ sdl });

      const httpOptions = result.groups[0].services[0].expose[0].httpOptions;
      expect(httpOptions?.maxBodySize).toBe(104857600);
      expect(httpOptions?.readTimeout).toBe(50);
      expect(httpOptions?.sendTimeout).toBe(100);
      expect(httpOptions?.nextTries).toBe(24);
      expect(httpOptions?.nextTimeout).toBe(48);
      expect(httpOptions?.nextCases).toEqual(["500"]);
    });

    it("handles accept hosts", () => {
      const sdl = createBasicSdl({
        expose: [{ port: 80, as: 80, to: [{ global: true }], accept: ["example.com", "api.example.com"] }],
      });
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].expose[0].hosts).toEqual(["example.com", "api.example.com"]);
    });

    it("handles IP lease endpoints", () => {
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
      const { result } = setup({ sdl });

      expect(result.groups[0].services[0].expose[0].ip).toBe("myip");
      expect(result.groups[0].services[0].expose[0].endpointSequenceNumber).toBe(1);

      // Check endpoint in resources
      const endpoints = result.groups[0].services[0].resources?.endpoints;
      expect(endpoints).toContainEqual(expect.objectContaining({ kind: 2, sequenceNumber: 1 }));
    });

    it("assigns sequence numbers to global endpoints in sorted order", async () => {
      const sdlContent = readFileSync("../fixtures/shared-ip.yml");
      const sdl = yaml.template<SDLInput>(sdlContent);
      const { result } = setup({ sdl });

      const sequenceNumbers = result.groups
        .flatMap((g) => g.services)
        .flatMap((s) => s.expose ?? [])
        .map((e) => e.endpointSequenceNumber);

      expect(sequenceNumbers).toEqual([2, 2]);
    });
  });

  describe("complex scenarios", () => {
    it("handles multi-service SDL", () => {
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
      const { result } = setup({ sdl });

      expect(result.groups).toHaveLength(1);
      expect(result.groups[0].services).toHaveLength(2);
      expect(result.groups[0].services.map((s) => s.name)).toEqual(["db", "web"]);

      // GroupSpecs have 2 resource units (different profiles)
      expect(result.groupSpecs[0].resources).toHaveLength(2);
    });

    it("handles service-to-service expose", () => {
      const sdl: SDLInput = yaml`
        version: "2.0"
        services:
          api:
            image: api:latest
            expose:
              - port: 3000
                to:
                  - service: web
          web:
            image: web:latest
            expose:
              - port: 80
                as: 80
                to:
                  - global: true
        profiles:
          compute:
            api:
              resources:
                cpu:
                  units: 0.5
                memory:
                  size: 512Mi
                storage:
                  size: 512Mi
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
                api:
                  denom: uakt
                  amount: 1000
                web:
                  denom: uakt
                  amount: 1000
        deployment:
          api:
            dcloud:
              profile: api
              count: 1
          web:
            dcloud:
              profile: web
              count: 1
      `;
      const { result } = setup({ sdl });

      const apiService = result.groups[0].services.find((s) => s.name === "api");
      expect(apiService?.expose[0].service).toBe("web");
      expect(apiService?.expose[0].global).toBe(false);
    });
  });

  describe("fixture file tests", () => {
    describe("GPU basic", () => {
      const sdlContent = readFileSync("../fixtures/gpu_basic.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("GPU basic manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("GPU basic groupSpecs");
      });
    });

    describe("GPU basic with RAM", () => {
      const sdlContent = readFileSync("../fixtures/gpu_basic_ram.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("GPU basic RAM manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("GPU basic RAM groupSpecs");
      });
    });

    describe("GPU basic with RAM and interface", () => {
      const sdlContent = readFileSync("../fixtures/gpu_basic_ram_interface.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("GPU basic RAM interface manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("GPU basic RAM interface groupSpecs");
      });
    });

    describe("WordPress multi-service", () => {
      const sdlContent = readFileSync("../fixtures/wordpress.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("WordPress manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("WordPress groupSpecs");
      });
    });

    describe("Persistent storage", () => {
      const sdlContent = readFileSync("../fixtures/persistent_storage_valid.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("Persistent storage manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("Persistent storage groupSpecs");
      });
    });

    describe("IP lease", () => {
      const sdlContent = readFileSync("../fixtures/ip_lease_valid.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("IP lease manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("IP lease groupSpecs");
      });
    });

    describe("GPU without model", () => {
      const sdlContent = readFileSync("../fixtures/gpu_basic_no_model.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("GPU no model manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("GPU no model groupSpecs");
      });
    });

    describe("GPU without GPU (units=0)", () => {
      const sdlContent = readFileSync("../fixtures/gpu_no_gpu_valid.sdl.yml");

      it("generates manifest (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroups(result.groups)).toMatchSnapshot("GPU zero units manifest");
      });

      it("generates groupSpecs (snapshot)", () => {
        const sdl = yaml.template<SDLInput>(sdlContent);
        const { result } = setup({ sdl });

        expect(stringifyGroupSpecs(result.groupSpecs)).toMatchSnapshot("GPU zero units groupSpecs");
      });
    });
  });

  function setup(input?: {
    sdl: SDLInput;
  }) {
    const result = generateManifest(input?.sdl ?? createBasicSdl());
    assertBuildResult(result);
    return { result: result.value };
  }

  function assertBuildResult(result: ReturnType<typeof generateManifest>): asserts result is Extract<GenerateManifestResult, { ok: true }> {
    if (!result.ok) {
      throw new Error(`Expected BuildResult but got validation errors: ${JSON.stringify(result)}`);
    }
  }

  function createBasicSdl(input: {
    image?: string;
    port?: number;
    cpu?: number | string;
    memory?: string;
    storage?: SDLInput["profiles"]["compute"][string]["resources"]["storage"];
    gpu?: SDLInput["profiles"]["compute"][string]["resources"]["gpu"];
    denom?: string;
    amount?: number;
    placementAttributes?: SDLInput["profiles"]["placement"][string]["attributes"];
    signedBy?: SDLInput["profiles"]["placement"][string]["signedBy"];
    env?: SDLInput["services"][string]["env"];
    command?: SDLInput["services"][string]["command"];
    args?: SDLInput["services"][string]["args"];
    credentials?: SDLInput["services"][string]["credentials"];
    expose?: SDLInput["services"][string]["expose"];
    endpoints?: SDLInput["endpoints"];
  } = {}): SDLInput {
    const {
      port = 80,
      cpu = 0.5,
      memory = "512Mi",
      storage = "512Mi",
      gpu,
      denom = "uakt",
      amount = 1000,
    } = input;

    return yaml`
      version: "2.1"
      services:
        web:
          image: ${input.image ?? "nginx"}
          env: ${input.env}
          command: ${input.command}
          args: ${input.args}
          credentials: ${input.credentials}
          expose: ${input.expose ?? [{ port, as: port, to: [{ global: true }] }]}
      profiles:
        compute:
          web:
            resources:
              cpu:
                units: ${cpu}
              memory:
                size: ${memory}
              storage: ${typeof storage === "string" ? { size: storage } : storage}
              gpu: ${gpu}
        placement:
          dcloud:
            pricing:
              web: ${{ denom, amount }}
            attributes: ${input.placementAttributes}
            signedBy: ${input.signedBy}
      deployment:
        web:
          dcloud:
            profile: web
            count: 1
      endpoints: ${input.endpoints}
    `;
  }

  type RelativeFilePath<T extends string> = string & { baseDir?: T };
  function readFileSync(filePath: RelativeFilePath<"$dir">): string {
    const fullPath = path.join(__dirname, filePath);
    return fs.readFileSync(fullPath, "utf8");
  }

  function stringifyGroups(manifest: Group[]): string {
    return JSON.stringify(manifest.map((g) => Group.toJSON(g)), null, 2);
  }

  function stringifyGroupSpecs(manifest: GroupSpec[]): string {
    return JSON.stringify(manifest.map((g) => GroupSpec.toJSON(g)), null, 2);
  }
});
