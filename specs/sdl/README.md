# SDL Schema Documentation

## Schema Overview

The SDL pipeline uses three schemas that validate different stages:

**`sdl-input.schema.yaml`** — Validates user input (YAML files)
- Used by: VS Code IDE validation and parity tests
- Validates: The SDL YAML files developers write

**`manifest-output.schema.yaml`** — Validates output manifest JSON
- Used by: Parity tests and validation tools
- Validates: The generated manifest JSON output from the SDL parser

**`groups-output.schema.yaml`** — Validates output groups JSON
- Used by: Parity tests and validation tools
- Validates: The generated deployment groups JSON output from the SDL parser

## Validation Capabilities

The `sdl-input.schema.yaml` validates:
- **Structure**: Types, required fields, enums, minLength, min/max values
- **Property name patterns** (using `patternProperties`):
  - Endpoints: Endpoint keys must match `/^[a-z]+[-_\da-z]+$/` (must start with lowercase letters, can contain lowercase letters, hyphens, underscores, and digits)
- **Simple conditional validations** (using `if/then`):
  - Storage: If `class === "ram"`, then `persistent` must be `false`
  - HTTP options: If `next_cases` contains `"off"`, it must be the only value
  - Expose: If `to.ip` is present (non-empty), then `to.global` must be `true`

## Validation Limitations

The `sdl-input.schema.yaml` validates **structure only** (types, required fields, enums). It **cannot** validate semantic relationships that require cross-referencing or document traversal.

**Note:** These limitations apply only to `sdl-input.schema.yaml`. The `manifest-output.schema.yaml` and `groups-output.schema.yaml` validate output JSON that is already processed by the Go parser, so they don't have semantic validation issues.

**Not validated by `sdl-input.schema.yaml`:**
- Unused endpoints (declared but never referenced)
- Missing profile/service references (deployment → compute/placement)
- Storage volume references (params.storage → compute.storage)
- Port collisions and mount path uniqueness
- Cross-field conditional requirements:
  - Persistent storage requires mount in service params (cross-reference validation)
  - Storage class="ram" requires readOnly=false in service params (cross-reference validation)
  - Storage persistent=false cannot have class attribute (hard to express: JSON Schema cannot directly enforce "property must not exist")

**Validated at runtime:** These checks are performed by the Go parser (`go/sdl/v2.go`, `go/sdl/v2_1.go`). Always validate SDL files using the Go parser before deployment - IDE validation will not catch these errors.

