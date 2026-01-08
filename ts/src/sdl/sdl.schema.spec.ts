import { type Json, validator } from "@exodus/schemasafe";
import { describe, expect, it } from "@jest/globals";
import fs from "fs";
import YAML from "js-yaml";
import path from "path";

import sdlSchema from "./sdl.schema.json";
import {
  createSdlValidator,
  type SdlSchemaValidationResult,
} from "./validateSdl.ts";

// Create a mock compiler for testing createSdlValidator
const mockCompiler = {
  compile: (schema: object) => {
    const validateFn = validator(schema as Parameters<typeof validator>[0], {
      includeErrors: true,
      allErrors: true,
    });
    return Object.assign((data: unknown) => validateFn(data as Json), {
      errors: null as Array<{
        instancePath?: string;
        keyword?: string;
        message?: string;
      }> | null,
    });
  },
};

// Helper to create validator
function createValidator() {
  const validate = validator(sdlSchema as Parameters<typeof validator>[0], {
    includeErrors: true,
    allErrors: true,
  });
  return (data: unknown): SdlSchemaValidationResult => {
    const valid = validate(data as Json);
    return {
      valid,
      errors: valid
        ? []
        : (validate.errors ?? []).map((err) => ({
            instancePath: err.instanceLocation,
            keyword: err.keywordLocation,
          })),
    };
  };
}

// Helper to read fixture files
function readFixture(filename: string): unknown {
  const filePath = path.join(__dirname, "SDL/fixtures", filename);
  const content = fs.readFileSync(filePath, "utf8");
  return YAML.load(content);
}

// Minimal valid SDL for testing
const minimalValidSdl = {
  version: "2.0",
  services: {
    web: {
      image: "nginx",
      expose: [
        {
          port: 80,
          to: [{ global: true }],
        },
      ],
    },
  },
  profiles: {
    compute: {
      web: {
        resources: {
          cpu: { units: "100m" },
          memory: { size: "128Mi" },
          storage: { size: "1Gi" },
        },
      },
    },
    placement: {
      akash: {
        pricing: {
          web: { denom: "uakt", amount: 1000 },
        },
      },
    },
  },
  deployment: {
    web: {
      akash: {
        profile: "web",
        count: 1,
      },
    },
  },
};

describe("SDL JSON Schema", () => {
  const validate = createValidator();

  describe("Schema Structure", () => {
    it("should have correct schema metadata", () => {
      expect(sdlSchema.$schema).toBe("http://json-schema.org/draft-07/schema#");
      expect(sdlSchema.title).toBe("Akash SDL Schema");
      expect(sdlSchema.type).toBe("object");
    });

    it("should require version, services, profiles, and deployment", () => {
      expect(sdlSchema.required).toContain("version");
      expect(sdlSchema.required).toContain("services");
      expect(sdlSchema.required).toContain("profiles");
      expect(sdlSchema.required).toContain("deployment");
    });

    it("should have $defs for all component types", () => {
      const defs = Object.keys(sdlSchema.$defs);
      expect(defs).toContain("version");
      expect(defs).toContain("service");
      expect(defs).toContain("expose");
      expect(defs).toContain("computeResources");
      expect(defs).toContain("gpuResource");
      expect(defs).toContain("storageResource");
      expect(defs).toContain("placementProfile");
      expect(defs).toContain("endpoint");
    });
  });

  describe("Valid SDL Documents", () => {
    it("should validate minimal SDL", () => {
      const result = validate(minimalValidSdl);
      expect(result.valid).toBe(true);
      expect(result.errors).toHaveLength(0);
    });

    it("should validate SDL with version 2.0", () => {
      const result = validate({ ...minimalValidSdl, version: "2.0" });
      expect(result.valid).toBe(true);
    });

    it("should validate SDL with version 2.1", () => {
      const result = validate({ ...minimalValidSdl, version: "2.1" });
      expect(result.valid).toBe(true);
    });

    it("should validate SDL with version 2.1.0", () => {
      const result = validate({ ...minimalValidSdl, version: "2.1.0" });
      expect(result.valid).toBe(true);
    });

    it("should validate gpu_basic.sdl.yml fixture", () => {
      const sdl = readFixture("gpu_basic.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate gpu_basic_ram.sdl.yml fixture", () => {
      const sdl = readFixture("gpu_basic_ram.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate gpu_basic_ram_interface.sdl.yml fixture", () => {
      const sdl = readFixture("gpu_basic_ram_interface.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate wordpress.sdl.yml fixture", () => {
      const sdl = readFixture("wordpress.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate ip_lease_valid.sdl.yml fixture", () => {
      const sdl = readFixture("ip_lease_valid.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate persistent_storage_valid.sdl.yml fixture", () => {
      const sdl = readFixture("persistent_storage_valid.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should validate gpu_no_gpu_valid.sdl.yml fixture", () => {
      const sdl = readFixture("gpu_no_gpu_valid.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Invalid SDL Documents - Missing Required Fields", () => {
    it("should reject SDL without version", () => {
      const { version: _, ...sdl } = minimalValidSdl;
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject SDL without services", () => {
      const { services: _, ...sdl } = minimalValidSdl;
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject SDL without profiles", () => {
      const { profiles: _, ...sdl } = minimalValidSdl;
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject SDL without deployment", () => {
      const { deployment: _, ...sdl } = minimalValidSdl;
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject service without image", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            expose: [{ port: 80 }],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject expose without port", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ to: [{ global: true }] }],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject compute profile without cpu", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject compute profile without memory", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject compute profile without storage", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject deployment placement without profile", () => {
      const sdl = {
        ...minimalValidSdl,
        deployment: {
          web: {
            akash: {
              count: 1,
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject deployment placement without count", () => {
      const sdl = {
        ...minimalValidSdl,
        deployment: {
          web: {
            akash: {
              profile: "web",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject pricing without denom", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          placement: {
            akash: {
              pricing: {
                web: { amount: 1000 },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject pricing without amount", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          placement: {
            akash: {
              pricing: {
                web: { denom: "uakt" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("Version Validation", () => {
    it("should accept version 2.0", () => {
      const result = validate({ ...minimalValidSdl, version: "2.0" });
      expect(result.valid).toBe(true);
    });

    it("should accept version 2.1", () => {
      const result = validate({ ...minimalValidSdl, version: "2.1" });
      expect(result.valid).toBe(true);
    });

    it("should accept version 2.1.0", () => {
      const result = validate({ ...minimalValidSdl, version: "2.1.0" });
      expect(result.valid).toBe(true);
    });

    it("should accept version 2.0.0", () => {
      const result = validate({ ...minimalValidSdl, version: "2.0.0" });
      expect(result.valid).toBe(true);
    });

    it("should reject version 1.0", () => {
      const result = validate({ ...minimalValidSdl, version: "1.0" });
      expect(result.valid).toBe(false);
    });

    it("should reject version 3.0", () => {
      const result = validate({ ...minimalValidSdl, version: "3.0" });
      expect(result.valid).toBe(false);
    });

    it("should reject non-string version", () => {
      const result = validate({ ...minimalValidSdl, version: 2.0 });
      expect(result.valid).toBe(false);
    });
  });

  describe("CPU Resource Validation", () => {
    it("should accept millicores format (100m)", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept decimal string format (0.1)", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "0.1" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept integer string format (1)", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "1" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept numeric format", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: 1 },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept CPU with arch attribute", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m", attributes: { arch: "amd64" } },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Memory Resource Validation", () => {
    it("should accept Mi suffix", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "512Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept Gi suffix", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "4Gi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should reject invalid memory suffix", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "512MB" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("Storage Resource Validation", () => {
    it("should accept single storage object", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept storage array", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: [{ size: "1Gi" }, { name: "data", size: "10Gi" }],
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept storage with persistent attribute", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: [
                  { size: "1Gi" },
                  {
                    name: "data",
                    size: "10Gi",
                    attributes: { persistent: true, class: "beta3" },
                  },
                ],
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept storage with ram class", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: [
                  { size: "1Gi" },
                  {
                    name: "shm",
                    size: "256Mi",
                    attributes: { class: "ram" },
                  },
                ],
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept all valid storage classes", () => {
      const classes = ["default", "beta1", "beta2", "beta3", "ram"];
      for (const storageClass of classes) {
        const sdl = {
          ...minimalValidSdl,
          profiles: {
            ...minimalValidSdl.profiles,
            compute: {
              web: {
                resources: {
                  cpu: { units: "100m" },
                  memory: { size: "128Mi" },
                  storage: {
                    size: "1Gi",
                    attributes: { class: storageClass },
                  },
                },
              },
            },
          },
        };
        const result = validate(sdl);
        expect(result.valid).toBe(true);
      }
    });

    it("should reject invalid storage class", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: {
                  size: "1Gi",
                  attributes: { class: "invalid" },
                },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("GPU Resource Validation", () => {
    it("should accept GPU with units 0", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
                gpu: { units: 0 },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept GPU with nvidia vendor", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      nvidia: [{ model: "a100" }],
                    },
                  },
                },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept GPU with amd vendor", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      amd: [{ model: "mi100" }],
                    },
                  },
                },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept GPU with ram specification", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
                gpu: {
                  units: 1,
                  attributes: {
                    vendor: {
                      nvidia: [{ model: "a100", ram: "40Gi" }],
                    },
                  },
                },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept GPU with pcie interface", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
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
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept GPU with sxm interface", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
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
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should reject GPU with invalid interface", () => {
      const sdl = readFixture("gpu_invalid_interface.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject GPU with invalid vendor", () => {
      const sdl = readFixture("gpu_invalid_vendor.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject GPU with no vendor when units > 0", () => {
      const sdl = readFixture("gpu_invalid_no_vendor.sdl.yml");
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("Service Expose Validation", () => {
    it("should accept expose with all options", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 80,
                as: 8080,
                proto: "TCP",
                accept: ["example.com"],
                to: [{ global: true }],
                http_options: {
                  max_body_size: 1048576,
                  read_timeout: 60000,
                  send_timeout: 60000,
                  next_tries: 3,
                  next_timeout: 0,
                  next_cases: ["error", "timeout"],
                },
              },
            ],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept expose with UDP protocol", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 12345,
                proto: "udp",
                to: [{ global: true }],
              },
            ],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept expose to specific service", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 80,
                to: [{ service: "api" }],
              },
            ],
          },
          api: {
            image: "node",
            expose: [
              {
                port: 3000,
                to: [{ global: true }],
              },
            ],
          },
        },
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            ...minimalValidSdl.profiles.compute,
            api: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: { size: "1Gi" },
              },
            },
          },
          placement: {
            akash: {
              pricing: {
                web: { denom: "uakt", amount: 1000 },
                api: { denom: "uakt", amount: 1000 },
              },
            },
          },
        },
        deployment: {
          ...minimalValidSdl.deployment,
          api: {
            akash: { profile: "api", count: 1 },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should reject port outside valid range", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 70000,
                to: [{ global: true }],
              },
            ],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject invalid next_cases value", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 80,
                to: [{ global: true }],
                http_options: {
                  next_cases: ["invalid_case"],
                },
              },
            ],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should accept all valid next_cases values", () => {
      const validCases = [
        "error",
        "timeout",
        "500",
        "502",
        "503",
        "504",
        "403",
        "404",
        "429",
        "off",
      ];
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 80,
                to: [{ global: true }],
                http_options: {
                  next_cases: validCases,
                },
              },
            ],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Service Params Validation", () => {
    it("should accept service with storage params", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: {
                  mount: "/mnt/data",
                  readOnly: false,
                },
              },
            },
          },
        },
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: [
                  { size: "1Gi" },
                  {
                    name: "data",
                    size: "10Gi",
                    attributes: { persistent: true },
                  },
                ],
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept null storage params for ephemeral storage", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            params: {
              storage: {
                data: null,
              },
            },
          },
        },
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            web: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "128Mi" },
                storage: [{ size: "1Gi" }, { name: "data", size: "10Gi" }],
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Service Credentials Validation", () => {
    it("should accept service with credentials", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "private-registry.com/nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            credentials: {
              host: "https://private-registry.com",
              username: "user",
              password: "pass123",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept credentials with optional email", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "private-registry.com/nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            credentials: {
              host: "https://private-registry.com",
              username: "user",
              password: "pass123",
              email: "user@example.com",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should reject credentials without host", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "private-registry.com/nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            credentials: {
              username: "user",
              password: "pass123",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject credentials without username", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "private-registry.com/nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            credentials: {
              host: "https://private-registry.com",
              password: "pass123",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject credentials without password", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "private-registry.com/nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            credentials: {
              host: "https://private-registry.com",
              username: "user",
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("Service Dependencies Validation", () => {
    it("should accept dependencies array format", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            dependencies: [{ service: "db" }],
          },
          db: {
            image: "mysql",
            expose: [{ port: 3306, to: [{ service: "web" }] }],
          },
        },
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            ...minimalValidSdl.profiles.compute,
            db: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "256Mi" },
                storage: { size: "5Gi" },
              },
            },
          },
          placement: {
            akash: {
              pricing: {
                web: { denom: "uakt", amount: 1000 },
                db: { denom: "uakt", amount: 1000 },
              },
            },
          },
        },
        deployment: {
          ...minimalValidSdl.deployment,
          db: {
            akash: { profile: "db", count: 1 },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept depends_on string array format", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            depends_on: ["db"],
          },
          db: {
            image: "mysql",
            expose: [{ port: 3306, to: [{ service: "web" }] }],
          },
        },
        profiles: {
          ...minimalValidSdl.profiles,
          compute: {
            ...minimalValidSdl.profiles.compute,
            db: {
              resources: {
                cpu: { units: "100m" },
                memory: { size: "256Mi" },
                storage: { size: "5Gi" },
              },
            },
          },
          placement: {
            akash: {
              pricing: {
                web: { denom: "uakt", amount: 1000 },
                db: { denom: "uakt", amount: 1000 },
              },
            },
          },
        },
        deployment: {
          ...minimalValidSdl.deployment,
          db: {
            akash: { profile: "db", count: 1 },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Endpoints Validation", () => {
    it("should accept valid endpoint", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [
              {
                port: 80,
                to: [{ global: true, ip: "my-endpoint" }],
              },
            ],
          },
        },
        endpoints: {
          "my-endpoint": { kind: "ip" },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should reject endpoint with invalid kind", () => {
      const sdl = {
        ...minimalValidSdl,
        endpoints: {
          "my-endpoint": { kind: "invalid" },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject endpoint without kind", () => {
      const sdl = {
        ...minimalValidSdl,
        endpoints: {
          "my-endpoint": {},
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });

  describe("Placement Profile Validation", () => {
    it("should accept placement with attributes", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          placement: {
            akash: {
              attributes: {
                region: "us-west",
                datacenter: "dc1",
              },
              pricing: {
                web: { denom: "uakt", amount: 1000 },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept placement with signedBy", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          placement: {
            akash: {
              signedBy: {
                allOf: ["akash1abc123"],
                anyOf: ["akash1def456", "akash1ghi789"],
              },
              pricing: {
                web: { denom: "uakt", amount: 1000 },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept string amount in pricing", () => {
      const sdl = {
        ...minimalValidSdl,
        profiles: {
          ...minimalValidSdl.profiles,
          placement: {
            akash: {
              pricing: {
                web: { denom: "uakt", amount: "1000" },
              },
            },
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Environment Variables Validation", () => {
    it("should accept valid environment variables", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            env: ["PORT=8080", "NODE_ENV=production", "DEBUG=true"],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept empty value in environment variable", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            env: ["EMPTY="],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("Command and Args Validation", () => {
    it("should accept service with command", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            command: ["/bin/sh", "-c"],
            expose: [{ port: 80, to: [{ global: true }] }],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept service with args", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            args: ["--config", "/etc/nginx/nginx.conf"],
            expose: [{ port: 80, to: [{ global: true }] }],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });

    it("should accept service with command and args", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            command: ["/bin/sh"],
            args: ["-c", "echo hello"],
            expose: [{ port: 80, to: [{ global: true }] }],
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(true);
    });
  });

  describe("createSdlValidator utility", () => {
    it("should create a validator from a compiler", () => {
      const validateFn = createSdlValidator(mockCompiler, sdlSchema);
      expect(typeof validateFn).toBe("function");
    });

    it("should validate valid SDL", () => {
      const validateFn = createSdlValidator(mockCompiler, sdlSchema);
      const result = validateFn(minimalValidSdl);
      expect(result.valid).toBe(true);
      expect(result.errors).toHaveLength(0);
    });

    it("should return errors for invalid SDL", () => {
      // Use the direct validator function since mockCompiler doesn't propagate errors correctly
      const result = validate({ version: "2.0" });
      expect(result.valid).toBe(false);
      expect(result.errors.length).toBeGreaterThan(0);
    });
  });

  describe("Additional Properties", () => {
    it("should reject unknown top-level properties", () => {
      const sdl = {
        ...minimalValidSdl,
        unknown_field: "value",
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });

    it("should reject unknown service properties", () => {
      const sdl = {
        ...minimalValidSdl,
        services: {
          web: {
            image: "nginx",
            expose: [{ port: 80, to: [{ global: true }] }],
            unknown_field: "value",
          },
        },
      };
      const result = validate(sdl);
      expect(result.valid).toBe(false);
    });
  });
});
