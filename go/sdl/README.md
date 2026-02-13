# SDL Schema Documentation

## Schema Files

**`sdl-input.schema.yaml`** (`go/sdl/`)
- Validates user YAML input
- Embedded in Go binary for runtime validation (logs warnings, doesn't block)
- Enforces stricter rules than Go parser (email length, denom pattern, GPU vendor, version enum)

**`manifest-output.schema.yaml`** (`specs/sdl/`)
- Validates generated manifest JSON

**`groups-output.schema.yaml`** (`specs/sdl/`)
- Validates generated deployment groups JSON

## Validation Capabilities

`sdl-input.schema.yaml` validates:
- **Types & Constraints**: Required fields, enums, string patterns, min/max, minLength
- **Patterns**: Endpoint names (`^[a-z]+[-_\da-z]+$`), denom (`^(uakt|ibc/.*)$`)
- **Conditionals**: RAM storage → persistent=false, IP endpoint → global=true
- **Strict Rules**: Email ≥5 chars, password ≥6 chars, version ∈ {2.0, 2.1}, GPU vendor (nvidia only)

## Validation Limitations

Schema validates structure only. **Go/TS parsers handle:**
- Cross-references (deployment → profiles, params.storage → compute.storage)
- Semantic constraints (unused endpoints, port collisions, mount uniqueness)
- Parser-level checks (count ≥ 1, unknown fields — TS validates, Go rejects during unmarshal)

## Test Fixtures

`testdata/sdl/input/invalid/` — Both schema and Go parser reject
`testdata/sdl/input/schema-only-invalid/` — Schema rejects, Go parser accepts (stricter rules)
`testdata/sdl/input/v2.0/`, `v2.1/` — Valid fixtures for parity tests (Go ↔ TS output comparison)
