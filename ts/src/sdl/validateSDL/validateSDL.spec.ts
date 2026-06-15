import { merge } from "lodash";
import { describe, expect, it } from "vitest";

import type { DeepPartial } from "../../encoding/typeEncodingHelpers.ts";
import type { NetworkId } from "../../network/index.ts";
import { AKT_DENOM } from "../../network/index.ts";
import { type SDLInput, validateSDL } from "./validateSDL.ts";

describe(validateSDL.name, () => {
  describe("valid SDL", () => {
    it("returns undefined for a valid SDL", () => {
      const { validate } = setup();
      expect(validate()).toBeUndefined();
    });
  });

  describe("denom validation", () => {
    it("accets uakt denom", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: {
                  amount: "1000",
                  denom: "uakt",
                },
              },
            },
          },
        },
      });

      const errors = validate();
      expect(errors).toBeUndefined();
    });

    it("accets uact denom", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: {
                  amount: "1000",
                  denom: "uact",
                },
              },
            },
          },
        },
      });

      const errors = validate();
      expect(errors).toBeUndefined();
    });
  });

  describe("deployment validation", () => {
    it("returns an error when service is not defined in deployment", () => {
      const sdl: SDLInput = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          // web is missing here
          other: {
            dcloud: { count: 1, profile: "web" },
          },
        },
      };

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Service \"web\" is not defined at \"/deployment\" section.",
        instancePath: "/deployment",
        schemaPath: "#/properties/deployment",
        keyword: "required",
        params: { missingProperty: "web" },
      }));
    });

    it("returns an error when placement is not defined", () => {
      const { validate } = setup({
        deployment: {
          web: {
            "nonexistent-placement": {
              count: 1,
              profile: "web",
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "The placement \"nonexistent-placement\" is not defined in the \"placement\" section.",
        instancePath: "/profiles/placement",
        schemaPath: "#/properties/profiles/properties/placement",
        keyword: "required",
        params: { missingProperty: "nonexistent-placement" },
      }));
    });

    it("returns an error when compute profile is not defined", () => {
      const { validate } = setup({
        deployment: {
          web: {
            dcloud: {
              count: 1,
              profile: "nonexistent-profile",
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "The compute requirements for the \"nonexistent-profile\" profile are not defined in the \"compute\" section.",
        instancePath: "/profiles/compute",
        schemaPath: "#/properties/profiles/properties/compute",
        keyword: "required",
        params: { missingProperty: "nonexistent-profile" },
      }));
    });

    it("returns an error when pricing for profile is not defined", () => {
      const sdl: SDLInput = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                other: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: {
            dcloud: { count: 1, profile: "web" },
          },
        },
      };

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: "The pricing for the \"web\" profile is not defined in the \"dcloud\" placement.",
        instancePath: "/profiles/placement/dcloud/pricing",
        schemaPath: "#/properties/profiles/properties/placement/additionalProperties/properties/pricing",
        keyword: "required",
        params: { missingProperty: "web" },
      }));
    });
  });

  describe("storage validation", () => {
    it("returns an error when service references non-existing storage volume", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { mount: "/data" },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Service \"web\" references non-existing compute volume \"data\".",
        instancePath: "/profiles/compute/web/resources/storage",
        schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/storage",
        keyword: "required",
        params: { missingProperty: "data" },
      }));
    });

    it("returns an error for multiple root ephemeral storages", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: {},
                logs: {},
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: [
                  { name: "data", size: "1Gi" },
                  { name: "logs", size: "1Gi" },
                ],
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Multiple root ephemeral storages are not allowed.",
        schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage",
        keyword: "uniqueItems",
      }));
    });

    it("returns an error when mount is used by multiple volumes", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { mount: "/mnt" },
                logs: { mount: "/mnt" },
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: [
                  { name: "data", size: "1Gi" },
                  { name: "logs", size: "1Gi" },
                ],
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Mount \"/mnt\" already in use by volume \"data\".",
        instancePath: "/services/web/params/storage/logs/mount",
        schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage/additionalProperties/properties/mount",
        keyword: "uniqueItems",
        params: { duplicate: "/mnt" },
      }));
    });

    it("returns an error when persistent storage has no mount", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { readOnly: false },
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {
                  name: "data",
                  size: "1Gi",
                  attributes: { class: "default", persistent: true },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Persistent storage \"data\" requires a mount path in /services/web/params/storage/data/mount.",
        instancePath: "/services/web/params/storage/data",
        schemaPath: "#/properties/services/additionalProperties/properties/params/properties/storage/additionalProperties/properties/mount",
        keyword: "required",
        params: { missingProperty: "mount" },
      }));
    });

    it("accepts persistent storage with mount defined", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { mount: "/data", readOnly: false },
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {
                  name: "data",
                  size: "1Gi",
                  attributes: { class: "default", persistent: true },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("permissions validation", () => {
    it("accepts valid permissions params", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              permissions: {
                read: ["deployment", "logs"],
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("returns an error for unknown properties in permissions", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              permissions: {
                read: ["deployment"],
                write: ["logs"],
              },
            },
          },
        },
      } as DeepPartial<SDLInput>);

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        keyword: "additionalProperties",
        params: { additionalProperty: "write" },
      }));
    });
  });

  describe("GPU validation", () => {
    it("returns an error when GPU units > 0 but no attributes", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 1 },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringMatching(/GPU must have attributes|Missing required field: "attributes"/),
        instancePath: "/profiles/compute/web/resources/gpu",
      }));
    });

    it("returns an error when GPU units > 0 but no vendor", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 1, attributes: {} },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "GPU must specify a vendor if units is not 0.",
        instancePath: "/profiles/compute/web/resources/gpu/attributes",
        schemaPath: "#/properties/profiles/properties/compute/additionalProperties/properties/resources/properties/gpu/properties/attributes/properties/vendor",
        keyword: "required",
        params: { missingProperty: "vendor" },
      }));
    });

    it("returns an error when GPU units = 0 but has attributes", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 0, attributes: { vendor: { nvidia: [] } } },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringMatching(/GPU must not have attributes|must be greater than 0/),
        instancePath: expect.stringMatching(/\/profiles\/compute\/web\/resources\/gpu/),
      }));
    });

    it("accepts GPU with units > 0 and valid vendor", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 1, attributes: { vendor: { nvidia: [] } } },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("accepts GPU with units = 0 and no attributes", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 0 },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("IP lease validation", () => {
    it("returns an error when IP is declared but not global", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                to: [{ ip: "myendpoint", global: false }],
              },
            ],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        instancePath: "/services/web/expose/0/to/0",
        schemaPath: "#/definitions/exposeToWithIpEnforcesGlobal/if",
        keyword: "if",
        params: { failingKeyword: "then" },
        message: "If an IP is declared, the directive must be declared as global.",
      }));
    });

    it("returns an error when IP references unknown endpoint", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                to: [{ ip: "unknown-endpoint", global: true }],
              },
            ],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Unknown endpoint \"unknown-endpoint\" for service \"web\". Add it to the \"endpoints\" section.",
        instancePath: "/endpoints/unknown-endpoint",
        schemaPath: "#/properties/endpoints",
        keyword: "required",
        params: { missingProperty: "unknown-endpoint" },
      }));
    });

    it("returns an error when same IP endpoint port is used by multiple services", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
          api: {
            image: "node:latest",
            expose: [
              {
                port: 3000,
                as: 80,
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
            api: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
                api: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: {
            dcloud: { count: 1, profile: "web" },
          },
          api: {
            dcloud: { count: 1, profile: "api" },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("IP endpoint \"myendpoint\" port"),
        keyword: "uniqueItems",
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("already in use by"),
      }));
    });

    it("accepts valid IP lease configuration", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("endpoint validation", () => {
    it("returns an error when endpoint is declared but never used", () => {
      const { validate } = setup({
        endpoints: {
          "unused-endpoint": { kind: "ip" },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Endpoint \"unused-endpoint\" declared but never used.",
        instancePath: "/endpoints/unused-endpoint",
        schemaPath: "#/properties/endpoints",
        keyword: "additionalProperties",
        params: { additionalProperty: "unused-endpoint" },
      }));
    });

    it("does not return an error when all endpoints are used", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("multiple services validation", () => {
    it("validates all services in the SDL", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
          api: {
            image: "node:latest",
            expose: [{ port: 3000, as: 3000, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
            api: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
                api: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: {
            dcloud: { count: 1, profile: "web" },
          },
          api: {
            dcloud: { count: 1, profile: "api" },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("returns errors for multiple invalid services", () => {
      const sdl: SDLInput = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
          api: {
            image: "node:latest",
            expose: [{ port: 3000, as: 3000, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
            api: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
                api: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          // both web and api are missing
          other: {
            dcloud: { count: 1, profile: "web" },
          },
        },
      };

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: "Service \"web\" is not defined at \"/deployment\" section.",
        keyword: "required",
        params: { missingProperty: "web" },
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: "Service \"api\" is not defined at \"/deployment\" section.",
        keyword: "required",
        params: { missingProperty: "api" },
      }));
    });
  });

  describe("protocol handling", () => {
    it("handles uppercase TCP protocol", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                proto: "TCP",
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("handles lowercase tcp protocol", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                proto: "tcp",
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("handles UDP protocol", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 53,
                as: 53,
                proto: "UDP",
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("allows same port with different protocols on same IP", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [
              {
                port: 80,
                as: 80,
                proto: "TCP",
                to: [{ ip: "myendpoint", global: true }],
              },
              {
                port: 80,
                as: 80,
                proto: "UDP",
                to: [{ ip: "myendpoint", global: true }],
              },
            ],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("schema validation: services", () => {
    it("returns an error when service image is empty", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"image\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("at least 1 character"),
      }));
    });

    it("returns an error for invalid port number (0)", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 0 as number, as: 80, to: [{ global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"port\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("at least 1"),
      }));
    });

    it("returns an error for invalid port number (65536)", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 65536 as number, as: 80, to: [{ global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"port\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("at most 65535"),
      }));
    });

    it("returns an error for invalid protocol", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, proto: "HTTP" as "TCP", to: [{ global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"proto\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("TCP, UDP, tcp, udp"),
      }));
    });
  });

  describe("schema validation: credentials", () => {
    it("returns an error when credentials host is missing", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            credentials: {
              username: "user",
              password: "password123",
            } as SDLInput["services"][string]["credentials"],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"host\""),
      }));
    });

    it("returns an error when credentials username is missing", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            credentials: {
              host: "registry.example.com",
              password: "password123",
            } as SDLInput["services"][string]["credentials"],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"username\""),
      }));
    });

    it("returns an error when credentials password is missing", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            credentials: {
              host: "registry.example.com",
              username: "user",
            } as SDLInput["services"][string]["credentials"],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"password\""),
      }));
    });

    it("returns an error when credentials password is too short", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            credentials: {
              host: "registry.example.com",
              username: "user",
              password: "short",
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"password\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("at least 6 characters"),
      }));
    });

    it("accepts valid credentials", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            credentials: {
              host: "registry.example.com",
              username: "user",
              password: "password123",
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("schema validation: http_options", () => {
    it("returns an error when max_body_size exceeds 100MB", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{
              port: 80,
              as: 80,
              to: [{ global: true }],
              http_options: {
                max_body_size: 104857601,
              },
            }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"max_body_size\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("104857600"),
      }));
    });

    it("returns an error when read_timeout exceeds 60000ms", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{
              port: 80,
              as: 80,
              to: [{ global: true }],
              http_options: {
                read_timeout: 60001,
              },
            }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"read_timeout\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("60000"),
      }));
    });

    it("returns an error when send_timeout exceeds 60000ms", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{
              port: 80,
              as: 80,
              to: [{ global: true }],
              http_options: {
                send_timeout: 60001,
              },
            }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"send_timeout\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("60000"),
      }));
    });

    it("accepts valid http_options", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{
              port: 80,
              as: 80,
              to: [{ global: true }],
              http_options: {
                max_body_size: 104857600,
                read_timeout: 60000,
                send_timeout: 60000,
                next_tries: 3,
                next_timeout: 0,
                next_cases: ["error", "timeout"],
              },
            }],
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("schema validation: storage mount path", () => {
    it("returns an error when mount path is not absolute", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { mount: "relative/path" as `/relative/path` },
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { name: "data", size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"mount\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("pattern \"^/\""),
      }));
    });

    it("accepts absolute mount path", () => {
      const { validate } = setup({
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: { mount: "/data" },
              },
            },
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { name: "data", size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("schema validation: endpoints", () => {
    it("returns an error for invalid endpoint name pattern", () => {
      const { validate } = setup({
        endpoints: {
          "123invalid": { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ ip: "123invalid", global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"123invalid\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("^[a-z]+[-_0-9a-z]+$"),
      }));
    });

    it("returns an error when endpoint kind is missing", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: {} as { kind: "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ ip: "myendpoint", global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"kind\""),
      }));
    });

    it("returns an error for invalid endpoint kind", () => {
      const { validate } = setup({
        endpoints: {
          myendpoint: { kind: "dns" as "ip" },
        },
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ ip: "myendpoint", global: true }] }],
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"kind\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("ip"),
      }));
    });
  });

  describe("schema validation: profiles.compute.resources", () => {
    it("returns an error when cpu is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"cpu\""),
      }));
    });

    it("returns an error when memory is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"memory\""),
      }));
    });

    it("returns an error when storage is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"storage\""),
      }));
    });

    it("returns an error when cpu.units is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: {},
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"units\""),
      }));
    });

    it("returns an error when memory.size is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: {},
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"size\""),
      }));
    });

    it("returns an error when storage.size is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {},
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"size\""),
      }));
    });
  });

  describe("schema validation: GPU attributes", () => {
    it("returns an error for invalid GPU interface", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      nvidia: [{ interface: "invalid" as "pcie" }],
                    },
                  },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"interface\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("pcie, sxm"),
      }));
    });

    it("accepts valid GPU with pcie interface", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      nvidia: [{ model: "a100", interface: "pcie" }],
                    },
                  },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("accepts valid GPU with sxm interface", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      nvidia: [{ model: "a100", interface: "sxm" }],
                    },
                  },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });

    it("returns an error for invalid GPU vendor", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      amd: [],
                    },
                  },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"amd\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("not allowed"),
      }));
    });
  });

  describe("schema validation: storage attributes", () => {
    it("returns an error when RAM storage is persistent", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {
                  name: "data",
                  size: "1Gi",
                  attributes: { class: "ram", persistent: true },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"ram\" storage"),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("cannot be persistent"),
      }));
    });

    it("returns an error when RAM storage persistent is string \"true\"", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {
                  name: "data",
                  size: "1Gi",
                  attributes: { class: "ram", persistent: "true" },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"ram\" storage"),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("cannot be persistent"),
      }));
    });

    it("accepts RAM storage when not persistent", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: {
                  size: "1Gi",
                  attributes: { class: "ram", persistent: false },
                },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
      });

      expect(validate()).toBeUndefined();
    });
  });

  describe("schema validation: pricing", () => {
    it("returns an error for invalid denom", () => {
      const { validate } = setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: "usdt" as "uact" },
              },
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"denom\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("should be one of: uakt, uact."),
      }));
    });

    it("returns an error when denom is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000" },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"denom\""),
      }));
    });

    it("returns an error when amount is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: { dcloud: { count: 1, profile: "web" } },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"amount\""),
      }));
    });
  });

  describe("schema validation: deployment", () => {
    it("returns an error when deployment count is 0", () => {
      const { validate } = setup({
        deployment: {
          web: {
            dcloud: {
              count: 0 as number,
              profile: "web",
            },
          },
        },
      });

      const errors = validate();

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"count\""),
      }));
      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("at least 1"),
      }));
    });

    it("returns an error when deployment profile is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: {
            dcloud: {
              count: 1,
            },
          },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"profile\""),
      }));
    });

    it("returns an error when deployment count is missing", () => {
      const sdl = {
        version: "2.0",
        services: {
          web: {
            image: "nginx:latest",
            expose: [{ port: 80, as: 80, to: [{ global: true }] }],
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            dcloud: {
              pricing: {
                web: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          web: {
            dcloud: {
              profile: "web",
            },
          },
        },
      } as unknown as SDLInput;

      const errors = validateSDL(sdl);

      expect(errors).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("\"count\""),
      }));
    });
  });

  describe("reclamation validation", () => {
    it.each(["24h", "720h", "8760h", "30m", "1s"])("accepts a valid min_window %j", (minWindow) => {
      const { validate } = setup({ reclamation: { min_window: minWindow } });
      expect(validate()).toBeUndefined();
    });

    it("accepts reclamation on a v2.0 SDL (version-agnostic)", () => {
      const { validate } = setup({ version: "2.0", reclamation: { min_window: "720h" } });
      expect(validate()).toBeUndefined();
    });

    it("accepts reclamation on a v2.1 SDL", () => {
      const { validate } = setup({ version: "2.1", reclamation: { min_window: "30m" } });
      expect(validate()).toBeUndefined();
    });

    it("returns no reclamation error when the block is absent", () => {
      const { validate } = setup();
      expect(validate()).toBeUndefined();
    });

    // The schema pattern is intentionally stricter than Go's `time.ParseDuration`:
    // it rejects compound ("1h30m"), fractional ("1.5h"), sub-second units
    // ("500ms"), signs, zero, and unitless values, leaving only whole s/m/h
    // windows. Go stays the lenient layer (see go/sdl/reclamation.go).
    it.each(["abc", "0s", "-1h", "100", "1h30m", "1.5h", "500ms"])("rejects an invalid min_window %j", (minWindow) => {
      const { validate } = setup({ reclamation: { min_window: minWindow } });
      expect(validate()).toContainEqual(expect.objectContaining({
        instancePath: "/reclamation/min_window",
        schemaPath: "#/properties/reclamation/properties/min_window/pattern",
        keyword: "pattern",
        message: expect.stringContaining("whole number followed by s, m, or h"),
      }));
    });
  });

  // Mirrors the Go SDL parser's interconnect cross-field rules (go/sdl/v2.go
  // validateInterconnect + go/sdl/gpu.go parse-time guards). Without these TS
  // checks, tenants using @akashnetwork/chain-sdk could broadcast SDLs
  // that the Go CLI would have rejected outright.
  describe("GPU interconnect validation", () => {
    function interconnectSetup(opts: {
      interconnect?: boolean;
      interconnectGroup?: string;
      units?: number;
      placementRequiresInterconnect?: boolean;
    } = {}) {
      const { interconnect, interconnectGroup, units = 1, placementRequiresInterconnect = true } = opts;
      const placementAttrs: Record<string, string> = {};
      if (placementRequiresInterconnect) placementAttrs["capabilities/gpu-interconnect"] = "true";

      const gpuAttrs: Record<string, unknown> = { vendor: { nvidia: [{ model: "a100" }] } };
      if (interconnect !== undefined) gpuAttrs.interconnect = interconnect;
      if (interconnectGroup !== undefined) gpuAttrs.interconnect_group = interconnectGroup;

      return setup({
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: { units, attributes: gpuAttrs },
              },
            },
          },
          placement: {
            dcloud: {
              attributes: placementAttrs,
            },
          },
        },
      });
    }

    it("accepts a valid interconnect profile under an interconnect-capable placement", () => {
      const { validate } = interconnectSetup({  interconnect: true });
      expect(validate()).toBeUndefined();
    });

    // units==0 + interconnect / interconnect_group is rejected by the schema-level
    // gpuAttributesRequireUnitsGt0 rule (any attribute present requires
    // units > 0), so the semantic validator never runs for that case.
    // Pinning the schema's behavior here so a future schema relaxation
    // doesn't silently open a hole.
    it("rejects gpu.attributes.interconnect=true when gpu.units is 0 (schema-level)", () => {
      const { validate } = interconnectSetup({  interconnect: true, units: 0 });
      expect(validate()).toContainEqual(expect.objectContaining({
        schemaPath: expect.stringContaining("gpuAttributesRequireUnitsGt0"),
      }));
    });

    it("rejects gpu.attributes.interconnect_group when gpu.units is 0 (schema-level)", () => {
      const { validate } = interconnectSetup({ interconnectGroup: "pair0", units: 0 });
      expect(validate()).toContainEqual(expect.objectContaining({
        schemaPath: expect.stringContaining("gpuAttributesRequireUnitsGt0"),
      }));
    });

    it("rejects interconnect_group set without interconnect=true", () => {
      const { validate } = interconnectSetup({ interconnectGroup: "pair0" });
      expect(validate()).toContainEqual(expect.objectContaining({
        message: expect.stringContaining(`sets gpu.attributes.interconnect_group="pair0" but does not set gpu.attributes.interconnect: true`),
      }));
    });

    it("rejects interconnect=true under a placement that does not require capabilities/gpu-interconnect=true", () => {
      const { validate } = interconnectSetup({ interconnect: true, placementRequiresInterconnect: false });
      expect(validate()).toContainEqual(expect.objectContaining({
        message: expect.stringContaining(`but placement does not require capabilities/gpu-interconnect=true`),
      }));
    });

    it("rejects implicit + explicit interconnect_group mixing within one placement", () => {
      // Two profiles, both interconnect, but only one sets interconnect_group.
      const { validate } = setup({
        services: {
          worker: {
            image: "nginx:latest",
          },
        },
        profiles: {
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: { nvidia: [{ model: "a100" }] },
                     interconnect: true,
                    interconnect_group: "pair0",
                  },
                },
              },
            },
            worker: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: { nvidia: [{ model: "a100" }] },
                     interconnect: true,
                    // interconnect_group intentionally omitted -> rule 3 violation
                  },
                },
              },
            },
          },
          placement: {
            dcloud: {
              attributes: { "capabilities/gpu-interconnect": "true" },
              pricing: {
                worker: { amount: "1000", denom: AKT_DENOM },
              },
            },
          },
        },
        deployment: {
          worker: {
            dcloud: { count: 1, profile: "worker" },
          },
        },
      });

      expect(validate()).toContainEqual(expect.objectContaining({
        message: expect.stringContaining("mixes explicit and implicit interconnect_group"),
      }));
    });
  });

  function setup(overrides: DeepPartial<SDLInput> = {}, networkId: NetworkId = "sandbox") {
    const defaultSDL: SDLInput = {
      version: "2.0",
      services: {
        web: {
          image: "nginx:latest",
          expose: [
            {
              port: 80,
              as: 80,
              to: [{ global: true }],
            },
          ],
        },
      },
      profiles: {
        compute: {
          web: {
            resources: {
              cpu: { units: 1 },
              memory: { size: "512Mi" },
              storage: { size: "1Gi" },
            },
          },
        },
        placement: {
          dcloud: {
            pricing: {
              web: {
                amount: "1000",
                denom: AKT_DENOM,
              },
            },
          },
        },
      },
      deployment: {
        web: {
          dcloud: {
            count: 1,
            profile: "web",
          },
        },
      },
    };

    const sdl = merge(defaultSDL, overrides);

    return {
      sdl,
      networkId,
      validate: () => validateSDL(sdl),
    };
  }
});
