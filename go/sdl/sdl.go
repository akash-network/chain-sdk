package sdl

import (
	"errors"
	"fmt"
	"os"

	"github.com/blang/semver/v4"
	"gopkg.in/yaml.v3"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
	dtypes "pkg.akt.dev/go/node/deployment/v1beta4"
)

const (
	sdlVersionField = "version"
)

var (
	errUninitializedConfig = errors.New("sdl: uninitialized")
	errSDLInvalidNoVersion = fmt.Errorf("%w: no version found", errSDLInvalid)
)

// SDL is the interface which wraps Validate, Deployment and Manifest methods
type SDL interface {
	DeploymentGroups() (dtypes.GroupSpecs, error)
	Manifest() (manifest.Manifest, error)
	Version() ([]byte, error)
	validate() error
}

var _ SDL = (*sdl)(nil)

type sdl struct {
	Ver  semver.Version `yaml:"version,-"`
	data SDL            `yaml:"-"`
}

func (s *sdl) UnmarshalYAML(node *yaml.Node) error {
	var result sdl

	foundVersion := false
	for idx := range node.Content {
		if node.Content[idx].Value == sdlVersionField {
			var err error
			if result.Ver, err = semver.ParseTolerant(node.Content[idx+1].Value); err != nil {
				return err
			}
			foundVersion = true
			break
		}
	}

	if !foundVersion {
		return errSDLInvalidNoVersion
	}

	// nolint: gocritic
	if result.Ver.EQ(semver.MustParse("2.0.0")) {
		var decoded v2
		if err := node.Decode(&decoded); err != nil {
			return err
		}

		result.data = &decoded
	} else if result.Ver.GE(semver.MustParse("2.1.0")) {
		var decoded v2_1
		if err := node.Decode(&decoded); err != nil {
			return err
		}

		result.data = &decoded
	} else {
		return fmt.Errorf("%w: config: unsupported version %q", errSDLInvalid, result.Ver)
	}

	*s = result

	return nil
}

// ReadFile read from given path and returns SDL instance
func ReadFile(path string) (SDL, error) {
	buf, err := os.ReadFile(path) //nolint: gosec
	if err != nil {
		return nil, err
	}
	return Read(buf)
}

// Read reads buffer data and returns SDL instance
func Read(buf []byte) (sdlObj SDL, err error) {
	schemaErr := validateInputAgainstSchema(buf)

	// Soft check if schema validation passed but the SDL is rejected by the Go parser
	defer func() {
		checkSchemaValidationResult(schemaErr, err)
	}()

	obj := &sdl{}
	if err = yaml.Unmarshal(buf, obj); err != nil {
		return nil, err
	}

	if err = validateSDL(obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// validateSDL runs semantic validation, deployment group validation,
// and manifest validation on an SDL instance.
func validateSDL(obj *sdl) error {
	if err := obj.validate(); err != nil {
		return fmt.Errorf("sdl validation: %w", err)
	}

	dgroups, err := obj.DeploymentGroups()
	if err != nil {
		return fmt.Errorf("deployment groups: %w", err)
	}

	vgroups := make([]dtypes.GroupSpec, 0, len(dgroups))
	for _, dgroup := range dgroups {
		vgroups = append(vgroups, dgroup)
	}

	if err = dtypes.ValidateDeploymentGroups(vgroups); err != nil {
		return fmt.Errorf("validate deployment groups: %w", err)
	}

	m, err := obj.Manifest()
	if err != nil {
		return fmt.Errorf("manifest: %w", err)
	}

	if err = m.Validate(); err != nil {
		return fmt.Errorf("validate manifest: %w", err)
	}

	return nil
}

// Version creates the deterministic Deployment Version hash from the SDL.
func (s *sdl) Version() ([]byte, error) {
	if s.data == nil {
		return nil, errUninitializedConfig
	}

	return s.data.Version()
}

func (s *sdl) DeploymentGroups() (dtypes.GroupSpecs, error) {
	if s.data == nil {
		return dtypes.GroupSpecs{}, errUninitializedConfig
	}

	return s.data.DeploymentGroups()
}

func (s *sdl) Manifest() (manifest.Manifest, error) {
	if s.data == nil {
		return manifest.Manifest{}, errUninitializedConfig
	}

	return s.data.Manifest()
}

func (s *sdl) validate() error {
	if s.data == nil {
		return errUninitializedConfig
	}

	return s.data.validate()
}

// ReadFileStrict reads from given path and returns SDL instance with strict
// schema enforcement. Unlike ReadFile, it returns an error immediately if the
// input fails schema validation against sdl-input.schema.yaml.
func ReadFileStrict(path string) (SDL, error) {
	buf, err := os.ReadFile(path) //nolint: gosec
	if err != nil {
		return nil, err
	}
	return ReadStrict(buf)
}

// ReadStrict reads buffer data and returns SDL instance with strict schema
// enforcement. It validates against sdl-input.schema.yaml first and returns
// an error immediately on schema failure, then continues with existing Go
// semantic validation.
func ReadStrict(buf []byte) (SDL, error) {
	if err := validateInputAgainstSchema(buf); err != nil {
		return nil, fmt.Errorf("strict schema validation failed: %w", err)
	}

	obj := &sdl{}
	if err := yaml.Unmarshal(buf, obj); err != nil {
		return nil, err
	}

	if err := validateSDL(obj); err != nil {
		return nil, err
	}

	return obj, nil
}
