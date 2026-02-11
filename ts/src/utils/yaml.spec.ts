import { describe, expect, it } from "@jest/globals";

import { yaml } from "./yaml.ts";

describe(yaml.name, () => {
  it("parses yaml", () => {
    const result = yaml`
      version: "2.0"
      services:
        web:
          image: nginx
          expose:
            - port: 80
    `;

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: "nginx",
          expose: [
            { port: 80 },
          ],
        },
      },
    });
  });

  it("parses yaml with primitive interpolated values", () => {
    const port = 80;
    const serviceImage = "nginx:latest";
    const isPrivate = true;
    const result = yaml`
      version: "2.0"
      services:
        web:
          image: ${serviceImage}
          private: ${isPrivate}
          expose:
            - port: ${port}
    `;

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: serviceImage,
          private: isPrivate,
          expose: [
            { port },
          ],
        },
      },
    });
  });

  it("parses yaml with object interpolated values", () => {
    const serviceConfig = {
      image: "nginx:latest",
      port: 90,
    };

    const result = yaml`
      version: "2.0"
      services:
        web: ${serviceConfig}
    `;

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: serviceConfig,
      },
    });
  });

  it("parses yaml with array interpolated values", () => {
    const servicePorts = [{ port: 80 }, { port: 443 }];
    const result = yaml`
      version: "2.0"
      services:
        web:
          image: nginx:latest
          expose: ${servicePorts}
    `;

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: "nginx:latest",
          expose: servicePorts,
        },
      },
    });
  });
});

describe(yaml.template.name, () => {
  it("parses yaml with primitive variables", () => {
    const result = yaml.template(`
      version: \${version}
      services:
        web:
          image: \${image}
          private: \${isPrivate}
          expose:
            - port: \${port}
      `,
    { version: "2.0", image: "nginx:latest", isPrivate: true, port: 80 },
    );

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: "nginx:latest",
          private: true,
          expose: [{ port: 80 }],
        },
      },
    });
  });

  it("parses yaml with object variables", () => {
    const serviceConfig = { image: "nginx:latest", port: 90 };
    const result = yaml.template(`
      version: "2.0"
      services:
        web: \${serviceConfig}
      `,
    { serviceConfig },
    );

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: serviceConfig,
      },
    });
  });

  it("parses yaml with array variables", () => {
    const ports = [{ port: 80 }, { port: 443 }];
    const result = yaml.template(`
      version: "2.0"
      services:
        web:
          image: nginx:latest
          expose: \${ports}
      `,
    { ports },
    );

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: "nginx:latest",
          expose: ports,
        },
      },
    });
  });

  it("throws ReferenceError for missing variables", () => {
    expect(() =>
      yaml.template(`version: \${version}`, {}),
    ).toThrow(new ReferenceError("Variable \"version\" is not provided to yaml template"));
  });

  it("leaves text without placeholders unchanged", () => {
    const result = yaml.template(`
      version: "2.0"
      services:
        web:
          image: nginx
      `,
    );

    expect(result).toEqual({
      version: "2.0",
      services: {
        web: {
          image: "nginx",
        },
      },
    });
  });
});
