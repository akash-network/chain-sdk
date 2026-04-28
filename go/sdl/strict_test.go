package sdl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestReadStrictValidInputs verifies that ReadFileStrict accepts all valid fixtures.
func TestReadStrictValidInputs(t *testing.T) {
	s := newParityTestSuite(t)

	for _, version := range []string{"v2.0", "v2.1"} {
		inputDir := filepath.Join(s.fixturesInputRoot, version)

		entries, err := os.ReadDir(inputDir)
		require.NoError(t, err)

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			fixtureName := entry.Name()
			inputPath := filepath.Join(inputDir, fixtureName, "input.yaml")

			t.Run(version+"/"+fixtureName, func(t *testing.T) {
				sdl, err := ReadFileStrict(inputPath)
				require.NoError(t, err, "ReadFileStrict should accept valid fixture %s", fixtureName)
				require.NotNil(t, sdl)
			})
		}
	}
}

// TestReadStrictRejectsInvalid verifies that ReadStrict rejects all invalid fixtures.
func TestReadStrictRejectsInvalid(t *testing.T) {
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
			_, err := ReadFileStrict(fixturePath)
			require.Error(t, err, "ReadFileStrict should reject invalid fixture")
		})
	}
}

// TestReadStrictRejectsSchemaOnlyInvalid verifies that ReadStrict rejects
// inputs that the schema rejects but the lenient Go parser accepts.
// This is the key difference from ReadFile.
func TestReadStrictRejectsSchemaOnlyInvalid(t *testing.T) {
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
			// Lenient ReadFile should accept these
			_, lenientErr := ReadFile(fixturePath)
			require.NoError(t, lenientErr, "ReadFile (lenient) should accept this input")

			// Strict ReadFileStrict should reject these
			_, strictErr := ReadFileStrict(fixturePath)
			require.Error(t, strictErr, "ReadFileStrict should reject schema-only-invalid input")
			require.Contains(t, strictErr.Error(), "strict schema validation failed")
		})
	}
}
