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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

type parityTestSuite struct {
	fixturesInputRoot  string
	fixturesOutputRoot string
	schemasRoot        string

	manifestSchema *gojsonschema.Schema
	groupsSchema   *gojsonschema.Schema
}

// newParityTestSuite initializes a test suite with compiled output schemas.
func newParityTestSuite(t *testing.T) *parityTestSuite {
	t.Helper()

	s := &parityTestSuite{
		fixturesInputRoot:  "../../testdata/sdl/input",
		fixturesOutputRoot: "../../testdata/sdl/output-fixtures",
		schemasRoot:        "../../specs/sdl",
	}

	var err error
	s.manifestSchema, err = compileSchemaFromPath(filepath.Join(s.schemasRoot, "manifest-output.schema.yaml"))
	require.NoError(t, err, "Failed to compile manifest output schema")

	s.groupsSchema, err = compileSchemaFromPath(filepath.Join(s.schemasRoot, "groups-output.schema.yaml"))
	require.NoError(t, err, "Failed to compile groups output schema")

	return s
}

// validateInputSchema validates raw YAML bytes against the embedded SDL input schema.
func (s *parityTestSuite) validateInputSchema(t *testing.T, inputBytes []byte) {
	t.Helper()
	err := validateInputAgainstSchema(inputBytes)
	require.NoError(t, err, "Input schema validation failed")
}

// validateOutputAgainstSchema validates manifest and group-specs JSON against output schemas.
func (s *parityTestSuite) validateOutputAgainstSchema(t *testing.T, manifestBytes []byte, groupSpecsBytes []byte) {
	t.Helper()

	err := validateDataAgainstCompiledSchema(manifestBytes, s.manifestSchema)
	require.NoError(t, err, "Manifest schema validation failed")

	err = validateDataAgainstCompiledSchema(groupSpecsBytes, s.groupsSchema)
	require.NoError(t, err, "Groups schema validation failed")
}

// validateFixtureBytes compares generated JSON output against a committed fixture file.
func (s *parityTestSuite) validateFixtureBytes(t *testing.T, expectedPath string, actualBytes []byte, name string) {
	t.Helper()
	expectedBytes, err := os.ReadFile(expectedPath)
	require.NoError(t, err, "Failed to read expected %s", name)

	require.JSONEq(t, string(expectedBytes), string(actualBytes), "%s does not match expected output", name)
}

// TestParityV2_0 runs parity tests for SDL v2.0 fixtures.
func TestParityV2_0(t *testing.T) {
	s := newParityTestSuite(t)
	s.testParity(t, "v2.0")
}

// TestParityV2_1 runs parity tests for SDL v2.1 fixtures.
func TestParityV2_1(t *testing.T) {
	s := newParityTestSuite(t)
	s.testParity(t, "v2.1")
}

// testParity validates all fixtures for a given SDL version against Go parser output.
func (s *parityTestSuite) testParity(t *testing.T, version string) {
	inputDir := filepath.Join(s.fixturesInputRoot, version)

	entries, err := os.ReadDir(inputDir)
	require.NoError(t, err)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		fixtureName := entry.Name()
		inputPath := filepath.Join(inputDir, fixtureName, "input.yaml")
		manifestPath := filepath.Join(s.fixturesOutputRoot, version, fixtureName, "manifest.json")
		groupSpecsPath := filepath.Join(s.fixturesOutputRoot, version, fixtureName, "group-specs.json")

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			t.Fatalf("manifest.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		if _, err := os.Stat(groupSpecsPath); os.IsNotExist(err) {
			t.Fatalf("group-specs.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		t.Run(fixtureName, func(t *testing.T) {
			inputBytes, err := os.ReadFile(inputPath)
			require.NoError(t, err)

			s.validateInputSchema(t, inputBytes)

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

			s.validateOutputAgainstSchema(t, manifestBytes, groupSpecsBytes)

			s.validateFixtureBytes(t, manifestPath, manifestBytes, "Manifest")
			s.validateFixtureBytes(t, groupSpecsPath, groupSpecsBytes, "GroupSpecs")
		})
	}
}

// TestInvalidSDLsRejected verifies that all invalid fixtures are rejected by the Go parser.
func TestInvalidSDLsRejected(t *testing.T) {
	s := newParityTestSuite(t)
	invalidDir := filepath.Join(s.fixturesInputRoot, "invalid")

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
	s := newParityTestSuite(t)
	semanticDir := filepath.Join(s.fixturesInputRoot, "semantic-only-invalid")

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
	s := newParityTestSuite(t)
	schemaOnlyDir := filepath.Join(s.fixturesInputRoot, "schema-only-invalid")

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

// compileSchemaFromPath loads a YAML JSON Schema file and compiles it for validation.
func compileSchemaFromPath(schemaPath string) (*gojsonschema.Schema, error) {
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	var schemaData any
	if err := yaml.Unmarshal(schemaBytes, &schemaData); err != nil {
		return nil, fmt.Errorf("failed to parse YAML schema: %w", err)
	}

	jsonBytes, err := json.Marshal(schemaData)
	if err != nil {
		return nil, fmt.Errorf("failed to convert schema to JSON: %w", err)
	}

	schemaLoader := gojsonschema.NewSchemaLoader()
	schema, err := schemaLoader.Compile(gojsonschema.NewBytesLoader(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to compile schema: %w", err)
	}

	return schema, nil
}

// validateDataAgainstCompiledSchema validates JSON bytes against a pre-compiled schema.
func validateDataAgainstCompiledSchema(data []byte, schema *gojsonschema.Schema) error {
	result, err := schema.Validate(gojsonschema.NewBytesLoader(data))
	if err != nil {
		return fmt.Errorf("failed to validate against schema: %w", err)
	}

	if !result.Valid() {
		var errors []string
		for _, desc := range result.Errors() {
			errors = append(errors, desc.String())
		}
		return fmt.Errorf("schema validation failed: %v", errors)
	}

	return nil
}
