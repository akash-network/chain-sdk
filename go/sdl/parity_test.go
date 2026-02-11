package sdl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

const fixturesInputRoot = "../../testdata/sdl/input"
const fixturesOutputRoot = "../../testdata/sdl/output-fixtures"
const schemasRoot = "../../specs/sdl" // Output schemas for tests

var (
	manifestSchema     *gojsonschema.Schema
	manifestSchemaOnce sync.Once
	manifestSchemaErr  error

	groupsSchema     *gojsonschema.Schema
	groupsSchemaOnce sync.Once
	groupsSchemaErr  error
)

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
		groupsPath := filepath.Join(fixturesOutputRoot, version, fixtureName, "groups.json")

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			t.Fatalf("manifest.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		if _, err := os.Stat(groupsPath); os.IsNotExist(err) {
			t.Fatalf("groups.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		t.Run(fixtureName, func(t *testing.T) {
			inputBytes, err := os.ReadFile(inputPath)
			require.NoError(t, err)

			validateInputSchema(t, inputBytes)

			sdl, err := ReadFile(inputPath)
			require.NoError(t, err)

			manifest, err := sdl.Manifest()
			require.NoError(t, err)

			groups, err := sdl.DeploymentGroups()
			require.NoError(t, err)

			manifestBytes, err := json.Marshal(manifest)
			require.NoError(t, err)

			groupsBytes, err := json.Marshal(groups)
			require.NoError(t, err)

			validateManifestSchemaBytes(t, manifestBytes)
			validateGroupsSchemaBytes(t, groupsBytes)

			validateFixtureBytes(t, manifestPath, manifestBytes, "Manifest")
			validateFixtureBytes(t, groupsPath, groupsBytes, "Groups")
		})
	}
}

func validateInputSchema(t *testing.T, inputBytes []byte) {
	err := validateInputAgainstSchema(inputBytes)
	require.NoError(t, err, "Input schema validation failed")
}

func validateManifestSchemaBytes(t *testing.T, manifestBytes []byte) {
	manifestSchemaOnce.Do(func() {
		manifestSchema, manifestSchemaErr = compileSchemaFromPath(filepath.Join(schemasRoot, "manifest-output.schema.yaml"))
	})
	require.NoError(t, manifestSchemaErr, "Failed to compile manifest schema")

	err := validateDataAgainstCompiledSchema(manifestBytes, manifestSchema)
	require.NoError(t, err, "Manifest schema validation failed")
}

func validateGroupsSchemaBytes(t *testing.T, groupsBytes []byte) {
	groupsSchemaOnce.Do(func() {
		groupsSchema, groupsSchemaErr = compileSchemaFromPath(filepath.Join(schemasRoot, "groups-output.schema.yaml"))
	})
	require.NoError(t, groupsSchemaErr, "Failed to compile groups schema")

	err := validateDataAgainstCompiledSchema(groupsBytes, groupsSchema)
	require.NoError(t, err, "Groups schema validation failed")
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

func compileSchemaFromPath(schemaPath string) (*gojsonschema.Schema, error) {
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	var schemaData any
	if err := yaml.Unmarshal(schemaBytes, &schemaData); err != nil {
		return nil, fmt.Errorf("failed to parse YAML schema: %w", err)
	}

	if err := sanitizeSchemaRefs(schemaData); err != nil {
		return nil, fmt.Errorf("invalid schema: %w", err)
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
