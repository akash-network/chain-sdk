package main

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

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
		return validateEvidenceSemantics(raw)
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

func validateEvidenceSemantics(raw []byte) error {
	var evidence EvidenceDocument
	if err := json.Unmarshal(raw, &evidence); err != nil {
		return fmt.Errorf("decode evidence: %w", err)
	}

	errs := make([]string, 0)
	if tierRank(evidence.AttestedTier) > tierRank(evidence.TargetTier) {
		errs = append(errs, fmt.Sprintf("attested_tier %q exceeds target_tier %q", evidence.AttestedTier, evidence.TargetTier))
	}
	if err := validateCapabilitySet("attested_capabilities", evidence.AttestedCapabilities); err != nil {
		errs = append(errs, err.Error())
	}
	if err := validateUint64String("audit_escrow_id", evidence.AuditEscrowID); err != nil {
		errs = append(errs, err.Error())
	}
	if err := validateUint64String("block_height", evidence.BlockHeight); err != nil {
		errs = append(errs, err.Error())
	}
	if err := validateBase64Field("inventory_nonce", evidence.InventoryNonce); err != nil {
		errs = append(errs, err.Error())
	}
	if evidence.Software.Signature != "" {
		if err := validateBase64Field("software.signature", evidence.Software.Signature); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if err := validateTimestampField("collected_at", evidence.CollectedAt); err != nil {
		errs = append(errs, err.Error())
	}
	if err := validateTimestampField("sustained_validation.last_checked_at", evidence.SustainedValidation.LastCheckedAt); err != nil {
		errs = append(errs, err.Error())
	}
	for idx, check := range evidence.Checks {
		if check.ObservedAt == "" {
			continue
		}
		if err := validateTimestampField(fmt.Sprintf("checks[%d].observed_at", idx), check.ObservedAt); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("evidence semantic validation failed: %s", strings.Join(errs, "; "))
	}

	return nil
}

func validateUint64String(field, value string) error {
	if _, err := strconv.ParseUint(value, 10, 64); err != nil {
		return fmt.Errorf("%s must be a uint64 string: %w", field, err)
	}

	return nil
}

func validateBase64Field(field, value string) error {
	if _, err := base64.StdEncoding.DecodeString(value); err != nil {
		return fmt.Errorf("%s must be base64: %w", field, err)
	}

	return nil
}

func validateTimestampField(field, value string) error {
	if _, err := time.Parse(time.RFC3339, value); err != nil {
		return fmt.Errorf("%s must be RFC3339 date-time: %w", field, err)
	}

	return nil
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
