package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed evidence.schema.json
var embeddedEvidenceSchema []byte

type evidenceSchemaValidator struct {
	compileOnce  sync.Once
	schema       *gojsonschema.Schema
	compileError error
}

var defaultEvidenceSchemaValidator evidenceSchemaValidator

func validateEvidenceBytes(raw []byte) error {
	return defaultEvidenceSchemaValidator.validate(raw)
}

func (v *evidenceSchemaValidator) validate(raw []byte) error {
	schema, err := v.compiledSchema()
	if err != nil {
		return err
	}

	result, err := schema.Validate(gojsonschema.NewBytesLoader(raw))
	if err != nil {
		return fmt.Errorf("failed to validate evidence schema: %w", err)
	}
	if result.Valid() {
		return nil
	}

	errs := make([]string, 0, len(result.Errors()))
	for _, desc := range result.Errors() {
		errs = append(errs, desc.String())
	}

	return fmt.Errorf("evidence schema validation failed: %s", strings.Join(errs, "; "))
}

func (v *evidenceSchemaValidator) compiledSchema() (*gojsonschema.Schema, error) {
	v.compileOnce.Do(func() {
		if len(embeddedEvidenceSchema) == 0 {
			v.compileError = fmt.Errorf("embedded evidence schema is empty")
			return
		}

		var schemaJSON map[string]any
		if err := json.Unmarshal(embeddedEvidenceSchema, &schemaJSON); err != nil {
			v.compileError = fmt.Errorf("failed to parse evidence schema: %w", err)
			return
		}

		adaptEvidenceSchemaForDraft7(schemaJSON)

		jsonBytes, err := json.Marshal(schemaJSON)
		if err != nil {
			v.compileError = fmt.Errorf("failed to prepare evidence schema: %w", err)
			return
		}

		loader := gojsonschema.NewSchemaLoader()
		loader.AutoDetect = false
		loader.Draft = gojsonschema.Draft7

		v.schema, err = loader.Compile(gojsonschema.NewBytesLoader(jsonBytes))
		if err != nil {
			v.compileError = fmt.Errorf("failed to compile evidence schema: %w", err)
			return
		}
	})

	return v.schema, v.compileError
}

func adaptEvidenceSchemaForDraft7(node any) {
	switch typed := node.(type) {
	case map[string]any:
		if schemaURL, ok := typed["$schema"].(string); ok && schemaURL == "https://json-schema.org/draft/2020-12/schema" {
			typed["$schema"] = "http://json-schema.org/draft-07/schema#"
		}
		if defs, ok := typed["$defs"]; ok {
			typed["definitions"] = defs
			delete(typed, "$defs")
		}
		if ref, ok := typed["$ref"].(string); ok {
			typed["$ref"] = strings.Replace(ref, "#/$defs/", "#/definitions/", 1)
		}
		for _, value := range typed {
			adaptEvidenceSchemaForDraft7(value)
		}
	case []any:
		for _, item := range typed {
			adaptEvidenceSchemaForDraft7(item)
		}
	}
}
