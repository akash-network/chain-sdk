import { describe, expect, it } from "@jest/globals";
import YAML from "js-yaml";

import { createSdl, createSdlWithObject } from "./createSdl.ts";
import { SDL_DEFAULTS } from "./defaults.ts";
import type { SdlObject } from "./types.ts";

describe("createSdl", () => {
  describe("minimal configuration", () => {
    it("should create SDL with just an image", () => {
      const yaml = createSdl({ image: "nginx" });
      const sdl = YAML.load(yaml) as SdlObject;

      expect(sdl.version).toBe(SDL_DEFAULTS.version);
      expect(sdl.services.app).toBeDefined();
      expect((sdl.services.app as { image: string }).image).toBe("nginx");
    });

    it("should use default values", () => {
      const yaml = createSdl({ image: "nginx" });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      expect((resources.cpu as { units: string }).units).toBe(SDL_DEFAULTS.cpu);
      expect((resources.memory as { size: string }).size).toBe(
        SDL_DEFAULTS.memory,
      );

      const deployment = sdl.deployment.app as Record<
        string,
        { count: number }
      >;
      expect(deployment[SDL_DEFAULTS.placement].count).toBe(
        SDL_DEFAULTS.replicas,
      );
    });

    it("should expose port 80 by default", () => {
      const yaml = createSdl({ image: "nginx" });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { expose: Array<{ port: number }> };
      expect(service.expose[0].port).toBe(80);
    });
  });

  describe("custom resources", () => {
    it("should accept string CPU units", () => {
      const yaml = createSdl({ image: "nginx", cpu: "1000m" });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      expect((resources.cpu as { units: string }).units).toBe("1000m");
    });

    it("should convert numeric CPU to millicores", () => {
      const yaml = createSdl({ image: "nginx", cpu: 0.5 });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      expect((resources.cpu as { units: string }).units).toBe("500m");
    });

    it("should accept memory configuration", () => {
      const yaml = createSdl({ image: "nginx", memory: "2Gi" });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      expect((resources.memory as { size: string }).size).toBe("2Gi");
    });

    it("should accept simple storage string", () => {
      const yaml = createSdl({ image: "nginx", storage: "10Gi" });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      const storage = resources.storage as Array<{ size: string }>;
      expect(storage[0].size).toBe("10Gi");
    });

    it("should accept storage array with persistent storage", () => {
      const yaml = createSdl({
        image: "nginx",
        storage: [
          { size: "1Gi" },
          {
            name: "data",
            size: "50Gi",
            persistent: true,
            class: "beta3",
            mount: "/data",
          },
        ],
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      const storage = resources.storage as Array<{
        name: string;
        size: string;
        attributes?: { persistent?: boolean; class?: string };
      }>;

      expect(storage).toHaveLength(2);
      expect(storage[1].name).toBe("data");
      expect(storage[1].size).toBe("50Gi");
      expect(storage[1].attributes?.persistent).toBe(true);
      expect(storage[1].attributes?.class).toBe("beta3");

      // Check mount params
      const service = sdl.services.app as {
        params?: { storage?: Record<string, { mount: string }> };
      };
      expect(service.params?.storage?.data?.mount).toBe("/data");
    });
  });

  describe("port configuration", () => {
    it("should accept single port number", () => {
      const yaml = createSdl({ image: "nginx", port: 8080 });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { expose: Array<{ port: number }> };
      expect(service.expose[0].port).toBe(8080);
    });

    it("should accept multiple ports", () => {
      const yaml = createSdl({ image: "nginx", ports: [80, 443] });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { expose: Array<{ port: number }> };
      expect(service.expose).toHaveLength(2);
      expect(service.expose[0].port).toBe(80);
      expect(service.expose[1].port).toBe(443);
    });

    it("should accept port configuration objects", () => {
      const yaml = createSdl({
        image: "nginx",
        ports: [
          { port: 80, as: 8080, global: true },
          { port: 443, proto: "tcp", accept: ["example.com"] },
        ],
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as {
        expose: Array<{
          port: number;
          as?: number;
          proto?: string;
          accept?: string[];
        }>;
      };
      expect(service.expose[0].port).toBe(80);
      expect(service.expose[0].as).toBe(8080);
      expect(service.expose[1].accept).toContain("example.com");
    });
  });

  describe("environment variables", () => {
    it("should accept env as object", () => {
      const yaml = createSdl({
        image: "nginx",
        env: { NODE_ENV: "production", PORT: "8080" },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { env: string[] };
      expect(service.env).toContain("NODE_ENV=production");
      expect(service.env).toContain("PORT=8080");
    });

    it("should accept env as array", () => {
      const yaml = createSdl({
        image: "nginx",
        env: ["NODE_ENV=production", "PORT=8080"],
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { env: string[] };
      expect(service.env).toContain("NODE_ENV=production");
      expect(service.env).toContain("PORT=8080");
    });
  });

  describe("command and args", () => {
    it("should accept command", () => {
      const yaml = createSdl({ image: "node", command: ["node", "server.js"] });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { command: string[] };
      expect(service.command).toEqual(["node", "server.js"]);
    });

    it("should accept args", () => {
      const yaml = createSdl({ image: "nginx", args: ["--config", "/etc/nginx.conf"] });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as { args: string[] };
      expect(service.args).toEqual(["--config", "/etc/nginx.conf"]);
    });
  });

  describe("GPU configuration", () => {
    it("should accept GPU configuration", () => {
      const yaml = createSdl({
        image: "pytorch/pytorch",
        gpu: { units: 1, vendor: "nvidia", model: "a100" },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      const gpu = resources.gpu as {
        units: number;
        attributes: { vendor: Record<string, unknown> };
      };
      expect(gpu.units).toBe(1);
      expect(gpu.attributes.vendor.nvidia).toBeDefined();
    });

    it("should accept GPU with RAM and interface", () => {
      const yaml = createSdl({
        image: "pytorch/pytorch",
        gpu: {
          units: 1,
          vendor: "nvidia",
          model: "a100",
          ram: "40Gi",
          interface: "pcie",
        },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const resources = (
        sdl.profiles.compute.app as { resources: Record<string, unknown> }
      ).resources;
      const gpu = resources.gpu as {
        units: number;
        attributes: {
          vendor: {
            nvidia: Array<{ model: string; ram?: string; interface?: string }>;
          };
        };
      };
      expect(gpu.attributes.vendor.nvidia[0].ram).toBe("40Gi");
      expect(gpu.attributes.vendor.nvidia[0].interface).toBe("pcie");
    });
  });

  describe("deployment configuration", () => {
    it("should accept custom replicas", () => {
      const yaml = createSdl({ image: "nginx", replicas: 3 });
      const sdl = YAML.load(yaml) as SdlObject;

      const deployment = sdl.deployment.app as Record<
        string,
        { count: number }
      >;
      expect(deployment.akash.count).toBe(3);
    });

    it("should accept custom pricing", () => {
      const yaml = createSdl({ image: "nginx", pricing: 5000 });
      const sdl = YAML.load(yaml) as SdlObject;

      const placement = sdl.profiles.placement.akash as {
        pricing: { app: { amount: number } };
      };
      expect(placement.pricing.app.amount).toBe(5000);
    });

    it("should accept custom denom", () => {
      const yaml = createSdl({ image: "nginx", denom: "ibc/..." });
      const sdl = YAML.load(yaml) as SdlObject;

      const placement = sdl.profiles.placement.akash as {
        pricing: { app: { denom: string } };
      };
      expect(placement.pricing.app.denom).toBe("ibc/...");
    });

    it("should accept custom placement name", () => {
      const yaml = createSdl({ image: "nginx", placement: "us-west" });
      const sdl = YAML.load(yaml) as SdlObject;

      expect(sdl.profiles.placement["us-west"]).toBeDefined();
      expect(sdl.deployment.app).toHaveProperty("us-west");
    });

    it("should accept placement attributes", () => {
      const yaml = createSdl({
        image: "nginx",
        attributes: { region: "us-west" },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const placement = sdl.profiles.placement.akash as {
        attributes: Record<string, string>;
      };
      expect(placement.attributes.region).toBe("us-west");
    });

    it("should accept signedBy requirements", () => {
      const yaml = createSdl({
        image: "nginx",
        signedBy: { anyOf: ["akash1..."] },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const placement = sdl.profiles.placement.akash as {
        signedBy: { anyOf: string[] };
      };
      expect(placement.signedBy.anyOf).toContain("akash1...");
    });
  });

  describe("service configuration", () => {
    it("should accept custom service name", () => {
      const yaml = createSdl({ image: "nginx", name: "web" });
      const sdl = YAML.load(yaml) as SdlObject;

      expect(sdl.services.web).toBeDefined();
      expect(sdl.profiles.compute.web).toBeDefined();
      expect(sdl.deployment.web).toBeDefined();
    });

    it("should accept credentials", () => {
      const yaml = createSdl({
        image: "private/image",
        credentials: {
          host: "https://registry.example.com",
          username: "user",
          password: "pass",
        },
      });
      const sdl = YAML.load(yaml) as SdlObject;

      const service = sdl.services.app as {
        credentials: { host: string; username: string; password: string };
      };
      expect(service.credentials.host).toBe("https://registry.example.com");
      expect(service.credentials.username).toBe("user");
    });
  });

  describe("createSdlWithObject", () => {
    it("should return both yaml and object", () => {
      const result = createSdlWithObject({ image: "nginx" });

      expect(typeof result.yaml).toBe("string");
      expect(result.object.version).toBe(SDL_DEFAULTS.version);
      expect(result.object.services.app).toBeDefined();
    });
  });
});
