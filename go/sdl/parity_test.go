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

const fixturesRoot = "../../testdata/sdl"
const schemasRoot = "../../specs/sdl"

func validateAgainstSchema(t *testing.T, name string, data []byte, schemaPath string) {
	t.Helper()

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("Schema file %s not found: %v", schemaPath, err)
	}

	var schemaJSON map[string]any
	err = yaml.Unmarshal(schemaBytes, &schemaJSON)
	require.NoError(t, err, "Failed to parse YAML schema")

	jsonBytes, err := json.Marshal(schemaJSON)
	require.NoError(t, err, "Failed to convert schema to JSON")

	schemaLoader := gojsonschema.NewSchemaLoader()
	schema, err := schemaLoader.Compile(gojsonschema.NewBytesLoader(jsonBytes))
	require.NoError(t, err, "Failed to compile schema")

	result, err := schema.Validate(gojsonschema.NewBytesLoader(data))
	require.NoError(t, err, "Failed to validate against schema")

	if !result.Valid() {
		var errors []string
		for _, desc := range result.Errors() {
			errors = append(errors, desc.String())
		}
		require.Failf(t, "%s validation failed", name, "Errors: %v", errors)
	}
}

func TestParityV2_0(t *testing.T) {
	testParity(t, "v2.0")
}

func TestParityV2_1(t *testing.T) {
	testParity(t, "v2.1")
}

func testParity(t *testing.T, version string) {
	fixturesDir := filepath.Join(fixturesRoot, version)

	entries, err := os.ReadDir(fixturesDir)
	if os.IsNotExist(err) {
		t.Fatalf("Fixtures directory %s does not exist", fixturesDir)
	}
	require.NoError(t, err)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		fixtureName := entry.Name()
		fixtureDir := filepath.Join(fixturesDir, fixtureName)
		inputPath := filepath.Join(fixtureDir, "input.yaml")
		manifestPath := filepath.Join(fixtureDir, "manifest.json")
		groupsPath := filepath.Join(fixtureDir, "groups.json")

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			t.Fatalf("manifest.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		if _, err := os.Stat(groupsPath); os.IsNotExist(err) {
			t.Fatalf("groups.json not generated for %s (run: make generate-sdl-fixtures)", fixtureName)
		}

		t.Run(fixtureName, func(t *testing.T) {
			sdl, err := ReadFile(inputPath)
			require.NoError(t, err)

			expectedManifestBytes, err := os.ReadFile(manifestPath)
			require.NoError(t, err, "Failed to read expected manifest.json")

			expectedGroupsBytes, err := os.ReadFile(groupsPath)
			require.NoError(t, err, "Failed to read expected groups.json")

			manifest, err := sdl.Manifest()
			require.NoError(t, err)

			actualManifestBytes, err := json.Marshal(manifest)
			require.NoError(t, err)

			var expectedManifest, actualManifest any
			require.NoError(t, json.Unmarshal(expectedManifestBytes, &expectedManifest))
			require.NoError(t, json.Unmarshal(actualManifestBytes, &actualManifest))
			require.Equal(t, expectedManifest, actualManifest, "Manifest does not match expected output")

			validateAgainstSchema(t, "manifest", actualManifestBytes, schemasRoot+"/manifest.schema.yaml")

			groups, err := sdl.DeploymentGroups()
			require.NoError(t, err)

			actualGroupsBytes, err := json.Marshal(groups)
			require.NoError(t, err)

			var expectedGroups, actualGroups any
			require.NoError(t, json.Unmarshal(expectedGroupsBytes, &expectedGroups))
			require.NoError(t, json.Unmarshal(actualGroupsBytes, &actualGroups))
			require.Equal(t, expectedGroups, actualGroups, "Groups does not match expected output")

			validateAgainstSchema(t, "groups", actualGroupsBytes, schemasRoot+"/groups.schema.yaml")
		})
	}
}

func TestInvalidSDLsRejected(t *testing.T) {
	invalidDir := filepath.Join(fixturesRoot, "invalid")

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
