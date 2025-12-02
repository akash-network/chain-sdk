# SDL Schema Validation Limitations

The `sdl-input.schema.yaml` validates **structure only** (types, required fields, enums). It **cannot** validate semantic relationships that require cross-referencing or document traversal.

**Not validated by schema:**
- Unused endpoints (declared but never referenced)
- Missing profile/service references (deployment → compute/placement)
- Storage volume references (params.storage → compute.storage)
- Port collisions and mount path uniqueness
- Conditional requirements (e.g., persistent storage requires mount)

**Validated at runtime:** These checks are performed by the Go parser (`go/sdl/v2.go`, `go/sdl/v2_1.go`). Always validate SDL files using the Go parser before deployment - IDE validation will not catch these errors.
