package sdl

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"cosmossdk.io/log"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"

	"pkg.akt.dev/go/util/noop"
)

// Embed the SDL input schema at compile time.
//
//go:embed sdl-input.schema.yaml
var embeddedSchemaYAML []byte

type loggerHolder struct{ log.Logger }

type SchemaValidator struct {
	compiledSchema     *gojsonschema.Schema
	schemaCompileOnce  sync.Once
	schemaCompileError error
	logger             atomic.Value
}

var defaultValidator = &SchemaValidator{}

func init() {
	defaultValidator.logger.Store(loggerHolder{noop.NewLogger()})
}

// SetSchemaLogger configures the logger for schema validation warnings.
// This allows the application to inject an observability-aware logger that
// will properly flow to Grafana, Loki, Prometheus, etc.
//
// Example usage in application startup:
//
//	sdl.SetSchemaLogger(appLogger.With("component", "sdl-schema-validator"))
//
// By default, a no-op logger is used, so validation warnings are silent.
// This function is thread-safe and can be called concurrently with validation.
func SetSchemaLogger(logger log.Logger) {
	if logger == nil {
		logger = noop.NewLogger()
	}
	defaultValidator.logger.Store(loggerHolder{logger})
}

// loadSchemaBytes loads the embedded schema compiled into the binary.
// The schema is automatically embedded from sdl-input.schema.yaml in this directory.
func (sv *SchemaValidator) loadSchemaBytes() ([]byte, error) {
	if len(embeddedSchemaYAML) == 0 {
		return nil, fmt.Errorf("embedded schema is empty; this should not happen as the schema file is tracked in git")
	}
	return embeddedSchemaYAML, nil
}

func (sv *SchemaValidator) getCompiledSchema() (*gojsonschema.Schema, error) {
	sv.schemaCompileOnce.Do(func() {
		schemaBytes, err := sv.loadSchemaBytes()
		if err != nil {
			sv.schemaCompileError = fmt.Errorf("failed to load schema: %w", err)
			return
		}

		var schemaJSON map[string]any
		if err := yaml.Unmarshal(schemaBytes, &schemaJSON); err != nil {
			sv.schemaCompileError = fmt.Errorf("failed to parse YAML schema: %w", err)
			return
		}

		if err := sanitizeSchemaRefs(schemaJSON); err != nil {
			sv.schemaCompileError = fmt.Errorf("invalid schema: %w", err)
			return
		}

		jsonBytes, err := json.Marshal(schemaJSON)
		if err != nil {
			sv.schemaCompileError = fmt.Errorf("failed to convert schema to JSON: %w", err)
			return
		}

		schemaLoader := gojsonschema.NewSchemaLoader()
		sv.compiledSchema, err = schemaLoader.Compile(gojsonschema.NewBytesLoader(jsonBytes))
		if err != nil {
			sv.schemaCompileError = fmt.Errorf("failed to compile schema: %w", err)
			return
		}
	})

	return sv.compiledSchema, sv.schemaCompileError
}

func validateInputAgainstSchema(buf []byte) error {
	return defaultValidator.validateInputAgainstSchema(buf)
}

func (sv *SchemaValidator) validateInputAgainstSchema(buf []byte) error {
	schema, err := sv.getCompiledSchema()
	if err != nil {
		return err
	}

	var inputYAML any
	if err := yaml.Unmarshal(buf, &inputYAML); err != nil {
		return fmt.Errorf("failed to parse input YAML: %w", err)
	}

	inputJSONBytes, err := json.Marshal(inputYAML)
	if err != nil {
		return fmt.Errorf("failed to convert input to JSON: %w", err)
	}

	result, err := schema.Validate(gojsonschema.NewBytesLoader(inputJSONBytes))
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

func checkSchemaValidationResult(schemaErr error, goValidationErr error) {
	defaultValidator.checkSchemaValidationResult(schemaErr, goValidationErr)
}

func (sv *SchemaValidator) checkSchemaValidationResult(schemaErr error, goValidationErr error) {
	if schemaErr == nil && goValidationErr != nil {
		logger := sv.logger.Load().(loggerHolder).Logger
		logger.Warn(
			"SDL schema validation mismatch",
			"schema_validation", "passed",
			"go_error", goValidationErr.Error(),
		)
	}
}

func sanitizeSchemaRefs(node any) error {
	switch typed := node.(type) {
	case map[string]any:
		for key, value := range typed {
			if key == "$ref" {
				ref, ok := value.(string)
				if !ok {
					return fmt.Errorf("schema $ref must be string")
				}

				if !strings.HasPrefix(ref, "#") {
					return fmt.Errorf("external schema reference %q is not allowed", ref)
				}
			}

			if err := sanitizeSchemaRefs(value); err != nil {
				return err
			}
		}
	case []any:
		for _, item := range typed {
			if err := sanitizeSchemaRefs(item); err != nil {
				return err
			}
		}
	}

	return nil
}
