// Package sdl provides SDL parsing and validation.
//
// Schema files:
//
//	sdl-input.schema.yaml (go/sdl/)
//	  - Validates user YAML input
//	  - Embedded in Go binary for runtime validation (logs warnings, doesn't block)
//	  - Enforces stricter rules than Go parser (email length, denom pattern, GPU vendor, version enum)
//
// Validation capabilities (sdl-input.schema.yaml):
//   - Types & constraints: required fields, enums, string patterns, min/max, minLength
//   - Patterns: endpoint names (^[a-z]+[-_\da-z]+$), denom (^(uakt|ibc/.*)$)
//   - Conditionals: RAM storage -> persistent=false, IP endpoint -> global=true
//   - Strict rules: email >=5 chars, password >=6 chars, version in {2.0, 2.1}, GPU vendor (nvidia only)
//
// Validation limitations:
// Schema validates structure only. Go/TS parsers handle:
//   - Cross-references (deployment -> profiles, params.storage -> compute.storage)
//   - Semantic constraints (unused endpoints, port collisions, mount uniqueness)
//   - Parser-level checks (count >= 1, unknown fields - TS validates, Go rejects during unmarshal)
//
// Test fixtures:
//   - testdata/sdl/input/invalid/ - Both schema and Go parser reject
//   - testdata/sdl/input/schema-only-invalid/ - Schema rejects, Go parser accepts (stricter rules)
//   - testdata/sdl/input/v2.0/, v2.1/ - Valid fixtures for parity tests (pure fixtures comparison, no output schema validation)
package sdl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

const fixturesInputRoot = "../../testdata/sdl/input"
const fixturesOutputRoot = "../../testdata/sdl/output-fixtures"
const manifestOutputSchemaPath = "manifest-output.schema.yaml"
const groupsOutputSchemaPath = "groups-output.schema.yaml"

func TestParityV2_0(t *testing.T) {
	testParity(t, "v2.0")
}

func TestParityV2_1(t *testing.T) {
	testParity(t, "v2.1")
}

func testParity(t *testing.T, version string) {
	inputDir := filepath.Join(fixturesInputRoot, version)

	entries, err := os.ReadDir(inputDir)
	require.NoError(t, err)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		fixtureName := entry.Name()
		inputPath := filepath.Join(inputDir, fixtureName, "input.yaml")
		manifestPath := filepath.Join(fixturesOutputRoot, version, fixtureName, "manifest.json")
		groupSpecsPath := filepath.Join(fixturesOutputRoot, version, fixtureName, "group-specs.json")

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			t.Fatalf("manifest.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		if _, err := os.Stat(groupSpecsPath); os.IsNotExist(err) {
			t.Fatalf("group-specs.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		t.Run(fixtureName, func(t *testing.T) {
			inputBytes, err := os.ReadFile(inputPath)
			require.NoError(t, err)

			validateInputSchema(t, inputBytes)

			sdl, err := ReadFile(inputPath)
			require.NoError(t, err)

			manifest, err := sdl.Manifest()
			require.NoError(t, err)

			groupSpecs, err := sdl.DeploymentGroups()
			require.NoError(t, err)

			manifestBytes, err := json.Marshal(manifest)
			require.NoError(t, err)

			groupSpecsBytes, err := json.Marshal(groupSpecs)
			require.NoError(t, err)

			validateOutputAgainstSchema(t, manifestBytes, manifestOutputSchemaPath, "Manifest")
			validateOutputAgainstSchema(t, groupSpecsBytes, groupsOutputSchemaPath, "GroupSpecs")

			validateFixtureBytes(t, manifestPath, manifestBytes, "Manifest")
			validateFixtureBytes(t, groupSpecsPath, groupSpecsBytes, "GroupSpecs")
		})
	}
}

func validateInputSchema(t *testing.T, inputBytes []byte) {
	err := validateInputAgainstSchema(inputBytes)
	require.NoError(t, err, "Input schema validation failed")
}

func validateOutputAgainstSchema(t *testing.T, outputBytes []byte, schemaPath, name string) {
	schemaBytes, err := os.ReadFile(schemaPath)
	require.NoError(t, err, "Failed to read %s schema", name)

	validator := &SchemaValidator{}

	var schemaJSON map[string]any
	require.NoError(t, yaml.Unmarshal(schemaBytes, &schemaJSON), "Failed to parse %s schema YAML", name)
	require.NoError(t, sanitizeSchemaRefs(schemaJSON), "Invalid %s schema refs", name)

	jsonSchemaBytes, err := json.Marshal(schemaJSON)
	require.NoError(t, err, "Failed to marshal %s schema to JSON", name)

	schemaLoader := gojsonschema.NewSchemaLoader()
	compiled, err := schemaLoader.Compile(gojsonschema.NewBytesLoader(jsonSchemaBytes))
	require.NoError(t, err, "Failed to compile %s schema", name)

	_ = validator
	result, err := compiled.Validate(gojsonschema.NewBytesLoader(outputBytes))
	require.NoError(t, err, "Failed to validate %s against schema", name)

	if !result.Valid() {
		var errors []string
		for _, desc := range result.Errors() {
			errors = append(errors, desc.String())
		}
		t.Errorf("%s output schema validation failed: %v", name, errors)
	}
}

func validateFixtureBytes(t *testing.T, expectedPath string, actualBytes []byte, name string) {
	expectedBytes, err := os.ReadFile(expectedPath)
	require.NoError(t, err, "Failed to read expected %s", name)

	require.JSONEq(t, string(expectedBytes), string(actualBytes), "%s does not match expected output", name)
}

func TestInvalidSDLsRejected(t *testing.T) {
	invalidDir := filepath.Join(fixturesInputRoot, "invalid")

	entries, err := os.ReadDir(invalidDir)
	if os.IsNotExist(err) {
		t.Skip("Invalid fixtures directory does not exist yet")
		return
	}
	require.NoError(t, err)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fixturePath := filepath.Join(invalidDir, entry.Name())
		t.Run(entry.Name(), func(t *testing.T) {
			_, err := ReadFile(fixturePath)
			require.Error(t, err)
		})
	}
}

// TestSemanticOnlyInvalid tests SDL files that pass schema validation but are
// rejected by both Go and TS parsers due to semantic constraints not expressible
// in JSON Schema (e.g., unused endpoints, cross-reference errors, duplicate mounts).
func TestSemanticOnlyInvalid(t *testing.T) {
	semanticDir := filepath.Join(fixturesInputRoot, "semantic-only-invalid")

	entries, err := os.ReadDir(semanticDir)
	if os.IsNotExist(err) {
		t.Skip("Semantic-only invalid fixtures directory does not exist yet")
		return
	}
	require.NoError(t, err)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fixturePath := filepath.Join(semanticDir, entry.Name())
		t.Run(entry.Name(), func(t *testing.T) {
			inputBytes, err := os.ReadFile(fixturePath)
			require.NoError(t, err)

			// Schema should accept (structurally valid)
			schemaErr := validateInputAgainstSchema(inputBytes)
			require.NoError(t, schemaErr, "Schema should accept this input (semantic-only invalid)")

			// Go parser should reject (semantically invalid)
			_, goErr := ReadFile(fixturePath)
			require.Error(t, goErr, "Go parser should reject this input (semantic validation)")
		})
	}
}

// TestSchemaOnlyValidations tests SDL files that are rejected by the JSON schema
// but accepted by the Go parser. These represent validations that exist only in
// the schema layer (e.g., string length limits, enum value constraints) but are
// not enforced in the Go validation logic.
func TestSchemaOnlyValidations(t *testing.T) {
	schemaOnlyDir := filepath.Join(fixturesInputRoot, "schema-only-invalid")

	entries, err := os.ReadDir(schemaOnlyDir)
	if os.IsNotExist(err) {
		t.Skip("Schema-only invalid fixtures directory does not exist yet")
		return
	}
	require.NoError(t, err)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fixturePath := filepath.Join(schemaOnlyDir, entry.Name())
		t.Run(entry.Name(), func(t *testing.T) {
			inputBytes, err := os.ReadFile(fixturePath)
			require.NoError(t, err)

			schemaErr := validateInputAgainstSchema(inputBytes)
			require.Error(t, schemaErr, "Schema should reject this input")

			_, goErr := ReadFile(fixturePath)
			require.NoError(t, goErr, "Go should accept this input (schema-only validation)")
		})
	}
}
