import { describe, expect, it } from "@jest/globals";
import YAML from "js-yaml";

import { SDL_DEFAULTS } from "./defaults.ts";
import { SdlBuilder } from "./SdlBuilder.ts";
import type { SdlObject } from "./types.ts";

describe("SdlBuilder", () => {
  describe("basic usage", () => {
    it("should create a minimal SDL", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 80, global: true }],
        })
        .computeProfile("web", {
          cpu: "500m",
          memory: "512Mi",
          storage: "1Gi",
        })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;

      expect(sdl.version).toBe(SDL_DEFAULTS.version);
      expect(sdl.services.web).toBeDefined();
      expect(sdl.profiles.compute.web).toBeDefined();
      expect(sdl.profiles.placement.akash).toBeDefined();
      expect(sdl.deployment.web).toBeDefined();
    });

    it("should set custom version", () => {
      const yaml = new SdlBuilder()
        .version("2.1")
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      expect(sdl.version).toBe("2.1");
    });
  });

  describe("multi-service configuration", () => {
    it("should create SDL with multiple services", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 80, global: true }],
        })
        .service("api", {
          image: "node:18",
          command: ["node", "server.js"],
          env: ["PORT=3000"],
          expose: [{ port: 3000, global: true }],
        })
        .service("db", {
          image: "postgres:15",
          expose: [{ port: 5432, service: "api" }],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .computeProfile("api", { cpu: "1", memory: "1Gi", storage: "2Gi" })
        .computeProfile("db", { cpu: "2", memory: "4Gi", storage: "50Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 2 })
        .deploy("api", "akash", { profile: "api", count: 1 })
        .deploy("db", "akash", { profile: "db", count: 1 })
        .pricing("akash", { web: 1000, api: 2000, db: 5000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;

      expect(Object.keys(sdl.services)).toHaveLength(3);
      expect(Object.keys(sdl.profiles.compute)).toHaveLength(3);
      expect(Object.keys(sdl.deployment)).toHaveLength(3);
    });
  });

  describe("multi-placement configuration", () => {
    it("should create SDL with multiple placements", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 80, global: true }],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("us-west", { attributes: { region: "us-west" } })
        .placement("us-east", { attributes: { region: "us-east" } })
        .deploy("web", "us-west", { profile: "web", count: 2 })
        .deploy("web", "us-east", { profile: "web", count: 2 })
        .pricing("us-west", { web: 1000 })
        .pricing("us-east", { web: 1200 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;

      expect(Object.keys(sdl.profiles.placement)).toHaveLength(2);
      const webDeployment = sdl.deployment.web as Record<string, unknown>;
      expect(webDeployment["us-west"]).toBeDefined();
      expect(webDeployment["us-east"]).toBeDefined();
    });
  });

  describe("service configuration", () => {
    it("should support command and args", () => {
      const yaml = new SdlBuilder()
        .service("app", {
          image: "node",
          command: ["node"],
          args: ["--max-old-space-size=4096", "server.js"],
          expose: [{ port: 3000, global: true }],
        })
        .computeProfile("app", { cpu: "1", memory: "4Gi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("app", "akash", { profile: "app", count: 1 })
        .pricing("akash", { app: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.app as { command: string[]; args: string[] };

      expect(service.command).toEqual(["node"]);
      expect(service.args).toEqual(["--max-old-space-size=4096", "server.js"]);
    });

    it("should support environment variables", () => {
      const yaml = new SdlBuilder()
        .service("app", {
          image: "node",
          env: ["NODE_ENV=production", "PORT=3000"],
          expose: [{ port: 3000, global: true }],
        })
        .computeProfile("app", { cpu: "1", memory: "1Gi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("app", "akash", { profile: "app", count: 1 })
        .pricing("akash", { app: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.app as { env: string[] };

      expect(service.env).toContain("NODE_ENV=production");
      expect(service.env).toContain("PORT=3000");
    });

    it("should support depends_on", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          depends_on: ["api", "db"],
          expose: [{ port: 80, global: true }],
        })
        .service("api", { image: "node", expose: [{ port: 3000, global: true }] })
        .service("db", { image: "postgres", expose: [{ port: 5432, service: "api" }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .computeProfile("api", { cpu: "1", memory: "1Gi", storage: "1Gi" })
        .computeProfile("db", { cpu: "1", memory: "1Gi", storage: "10Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .deploy("api", "akash", { profile: "api", count: 1 })
        .deploy("db", "akash", { profile: "db", count: 1 })
        .pricing("akash", { web: 1000, api: 1000, db: 2000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.web as { depends_on: string[] };

      expect(service.depends_on).toEqual(["api", "db"]);
    });

    it("should support credentials", () => {
      const yaml = new SdlBuilder()
        .service("app", {
          image: "private-registry/app",
          expose: [{ port: 3000, global: true }],
          credentials: {
            host: "https://private-registry.com",
            username: "user",
            password: "pass",
          },
        })
        .computeProfile("app", { cpu: "1", memory: "1Gi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("app", "akash", { profile: "app", count: 1 })
        .pricing("akash", { app: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.app as {
        credentials: { host: string; username: string };
      };

      expect(service.credentials.host).toBe("https://private-registry.com");
      expect(service.credentials.username).toBe("user");
    });

    it("should support storage params", () => {
      const yaml = new SdlBuilder()
        .service("app", {
          image: "postgres",
          expose: [{ port: 5432, global: true }],
          params: {
            storage: {
              data: { mount: "/var/lib/postgresql/data", readOnly: false },
            },
          },
        })
        .computeProfile("app", {
          cpu: "1",
          memory: "1Gi",
          storage: [
            { size: "1Gi" },
            { name: "data", size: "50Gi", persistent: true, class: "beta3" },
          ],
        })
        .placement("akash", {})
        .deploy("app", "akash", { profile: "app", count: 1 })
        .pricing("akash", { app: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.app as {
        params: { storage: Record<string, { mount: string }> };
      };

      expect(service.params.storage.data.mount).toBe(
        "/var/lib/postgresql/data",
      );
    });
  });

  describe("expose configuration", () => {
    it("should support external port mapping", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 3000, as: 80, global: true }],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.web as {
        expose: Array<{ port: number; as: number }>;
      };

      expect(service.expose[0].port).toBe(3000);
      expect(service.expose[0].as).toBe(80);
    });

    it("should support UDP protocol", () => {
      const yaml = new SdlBuilder()
        .service("dns", {
          image: "coredns",
          expose: [{ port: 53, proto: "udp", global: true }],
        })
        .computeProfile("dns", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("dns", "akash", { profile: "dns", count: 1 })
        .pricing("akash", { dns: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.dns as {
        expose: Array<{ proto: string }>;
      };

      expect(service.expose[0].proto).toBe("udp");
    });

    it("should support service-to-service exposure", () => {
      const yaml = new SdlBuilder()
        .service("db", {
          image: "postgres",
          expose: [{ port: 5432, service: "api" }],
        })
        .service("api", {
          image: "node",
          expose: [{ port: 3000, global: true }],
        })
        .computeProfile("db", { cpu: "1", memory: "1Gi", storage: "10Gi" })
        .computeProfile("api", { cpu: "1", memory: "1Gi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("db", "akash", { profile: "db", count: 1 })
        .deploy("api", "akash", { profile: "api", count: 1 })
        .pricing("akash", { db: 1000, api: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.db as {
        expose: Array<{ to: Array<{ service: string }> }>;
      };

      expect(service.expose[0].to[0].service).toBe("api");
    });

    it("should support accept hostnames", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 80, global: true, accept: ["example.com", "www.example.com"] }],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.web as {
        expose: Array<{ accept: string[] }>;
      };

      expect(service.expose[0].accept).toContain("example.com");
    });

    it("should support HTTP options", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [
            {
              port: 80,
              global: true,
              http_options: {
                max_body_size: 104857600,
                read_timeout: 30000,
              },
            },
          ],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const service = sdl.services.web as {
        expose: Array<{ http_options: { max_body_size: number } }>;
      };

      expect(service.expose[0].http_options.max_body_size).toBe(104857600);
    });
  });

  describe("compute resources", () => {
    it("should convert numeric CPU to millicores", () => {
      const yaml = new SdlBuilder()
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: 0.5, memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const resources = (
        sdl.profiles.compute.web as { resources: Record<string, unknown> }
      ).resources;

      expect((resources.cpu as { units: string }).units).toBe("500m");
    });

    it("should support GPU configuration", () => {
      const yaml = new SdlBuilder()
        .service("ml", { image: "pytorch", expose: [{ port: 8080, global: true }] })
        .computeProfile("ml", {
          cpu: "4",
          memory: "16Gi",
          storage: "100Gi",
          gpu: { units: 1, vendor: "nvidia", model: "a100", ram: "40Gi" },
        })
        .placement("akash", {})
        .deploy("ml", "akash", { profile: "ml", count: 1 })
        .pricing("akash", { ml: 10000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const resources = (
        sdl.profiles.compute.ml as { resources: Record<string, unknown> }
      ).resources;
      const gpu = resources.gpu as {
        units: number;
        attributes: {
          vendor: { nvidia: Array<{ model: string; ram: string }> };
        };
      };

      expect(gpu.units).toBe(1);
      expect(gpu.attributes.vendor.nvidia[0].model).toBe("a100");
      expect(gpu.attributes.vendor.nvidia[0].ram).toBe("40Gi");
    });

    it("should support multiple storage volumes", () => {
      const yaml = new SdlBuilder()
        .service("db", { image: "postgres", expose: [{ port: 5432, global: true }] })
        .computeProfile("db", {
          cpu: "2",
          memory: "4Gi",
          storage: [
            { size: "1Gi" },
            { name: "data", size: "100Gi", persistent: true, class: "beta3" },
          ],
        })
        .placement("akash", {})
        .deploy("db", "akash", { profile: "db", count: 1 })
        .pricing("akash", { db: 5000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const resources = (
        sdl.profiles.compute.db as { resources: Record<string, unknown> }
      ).resources;
      const storage = resources.storage as Array<{
        name: string;
        attributes?: { persistent: boolean; class: string };
      }>;

      expect(storage).toHaveLength(2);
      expect(storage[1].name).toBe("data");
      expect(storage[1].attributes?.persistent).toBe(true);
    });
  });

  describe("placement configuration", () => {
    it("should support attributes", () => {
      const yaml = new SdlBuilder()
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", { attributes: { region: "us-west", tier: "premium" } })
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const placement = sdl.profiles.placement.akash as {
        attributes: Record<string, string>;
      };

      expect(placement.attributes.region).toBe("us-west");
      expect(placement.attributes.tier).toBe("premium");
    });

    it("should support signedBy", () => {
      const yaml = new SdlBuilder()
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {
          signedBy: {
            allOf: ["akash1abc..."],
            anyOf: ["akash1def...", "akash1ghi..."],
          },
        })
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .build();

      const sdl = YAML.load(yaml) as SdlObject;
      const placement = sdl.profiles.placement.akash as {
        signedBy: { allOf: string[]; anyOf: string[] };
      };

      expect(placement.signedBy.allOf).toContain("akash1abc...");
      expect(placement.signedBy.anyOf).toHaveLength(2);
    });
  });

  describe("endpoints", () => {
    it("should support IP endpoints", () => {
      const yaml = new SdlBuilder()
        .service("web", {
          image: "nginx",
          expose: [{ port: 80, global: true, ip: "my-ip" }],
        })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .endpoint("my-ip", "ip")
        .build();

      const sdl = YAML.load(yaml) as SdlObject;

      expect(sdl.endpoints).toBeDefined();
      expect(sdl.endpoints!["my-ip"]).toEqual({ kind: "ip" });
    });
  });

  describe("output methods", () => {
    it("should return object with toObject()", () => {
      const sdl = new SdlBuilder()
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .toObject();

      expect(sdl.version).toBe(SDL_DEFAULTS.version);
      expect(sdl.services.web).toBeDefined();
    });

    it("should return both with buildWithObject()", () => {
      const result = new SdlBuilder()
        .service("web", { image: "nginx", expose: [{ port: 80, global: true }] })
        .computeProfile("web", { cpu: "500m", memory: "512Mi", storage: "1Gi" })
        .placement("akash", {})
        .deploy("web", "akash", { profile: "web", count: 1 })
        .pricing("akash", { web: 1000 })
        .buildWithObject();

      expect(typeof result.yaml).toBe("string");
      expect(result.object.version).toBe(SDL_DEFAULTS.version);
    });
  });
});
