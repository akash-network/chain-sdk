package sdl

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"

	"cosmossdk.io/log"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

// Embed the SDL input schema at compile time to make the binary self-contained.
// The schema is embedded directly from this directory, making binaries portable
// and ensuring library users get schema validation with zero configuration.
//
//go:embed sdl-input.schema.yaml
var embeddedSchemaYAML []byte

type SchemaValidator struct {
	compiledSchema     *gojsonschema.Schema
	schemaCompileOnce  sync.Once
	schemaCompileError error
	loggerMu           sync.RWMutex
	logger             log.Logger
}

var defaultValidator = &SchemaValidator{
	logger: &noOpLogger{},
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
	defaultValidator.loggerMu.Lock()
	defer defaultValidator.loggerMu.Unlock()

	// Treat nil as a request to fall back to the no-op logger.
	if logger == nil {
		defaultValidator.logger = &noOpLogger{}
		return
	}

	defaultValidator.logger = logger
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

type noOpLogger struct{}

func (l *noOpLogger) Debug(_ string, _ ...any) {}
func (l *noOpLogger) Info(_ string, _ ...any)  {}
func (l *noOpLogger) Warn(_ string, _ ...any)  {}
func (l *noOpLogger) Error(_ string, _ ...any) {}
func (l *noOpLogger) With(_ ...any) log.Logger { return l }
func (l *noOpLogger) Impl() any                { return l }

// Compile-time assertion that noOpLogger implements log.Logger
var _ log.Logger = (*noOpLogger)(nil)

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
	schemaFailed := schemaErr != nil
	goFailed := goValidationErr != nil

	// Only log if there's a mismatch
	if schemaFailed != goFailed {
		sv.loggerMu.RLock()
		logger := sv.logger
		sv.loggerMu.RUnlock()

		if schemaFailed && !goFailed {
			logger.Warn(
				"SDL schema validation mismatch",
				"schema_error", schemaErr.Error(),
				"go_validation", "passed",
			)
		} else if !schemaFailed && goFailed {
			logger.Warn(
				"SDL schema validation mismatch",
				"schema_validation", "passed",
				"go_error", goValidationErr.Error(),
			)
		}
	}
}
