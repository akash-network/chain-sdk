package sdl

import (
	"errors"
	"net/url"
	"strings"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
)

const (
	kbsModeProvider = "provider"
	kbsModeTenant   = "tenant"
)

type v2ServiceKBSParams struct {
	Mode                   string `yaml:"mode"`
	URL                    string `yaml:"url,omitempty"`
	Certificate            string `yaml:"certificate,omitempty"`
	ImageSecurityPolicyURI string `yaml:"imageSecurityPolicyURI,omitempty"`
	AgentPolicy            string `yaml:"agentPolicy,omitempty"`
}

func (service v2Service) validateKBS() error {
	requiresKBS := service.Credentials != nil && strings.TrimSpace(service.Credentials.URI) != ""
	for _, entry := range service.Env {
		_, value, ok := strings.Cut(entry, "=")
		if ok && strings.HasPrefix(value, "sealed.") {
			requiresKBS = true
		}
	}
	if service.Params != nil {
		for _, storage := range service.Params.Storage {
			if strings.TrimSpace(storage.KeyRef) != "" {
				requiresKBS = true
			}
		}
	}

	if service.Params == nil || service.Params.KBS == nil {
		if requiresKBS {
			return errors.New("KBS-backed workload data requires an explicit params.kbs selection")
		}
		return nil
	}
	return service.Params.KBS.validate(service.Params.TEE != "")
}

func (params v2ServiceKBSParams) validate(hasTEE bool) error {
	if !hasTEE {
		return errors.New("KBS configuration requires tee")
	}

	hasTenantFields := params.URL != "" || params.Certificate != "" ||
		params.ImageSecurityPolicyURI != "" || params.AgentPolicy != ""
	switch params.Mode {
	case kbsModeProvider:
		if hasTenantFields {
			return errors.New("provider mode cannot include tenant KBS fields")
		}
		return nil
	case kbsModeTenant:
		if params.URL == "" || params.Certificate == "" ||
			params.ImageSecurityPolicyURI == "" || params.AgentPolicy == "" {
			return errors.New("tenant mode requires a complete public configuration")
		}
		if err := validateTenantKBSURL(params.URL); err != nil {
			return err
		}
		if len(params.Certificate) > 64*1024 || strings.ContainsAny(params.Certificate, "\x00\r") {
			return errors.New("tenant KBS certificate must be a bounded PEM value")
		}
		if err := validateImageSecurityPolicyURI(params.ImageSecurityPolicyURI); err != nil {
			return err
		}
		if len(params.AgentPolicy) > 1024*1024 ||
			!strings.Contains(params.AgentPolicy, "package agent_policy") ||
			strings.ContainsAny(params.AgentPolicy, "\x00\r") {
			return errors.New("tenant agent policy must be a bounded agent_policy document")
		}
		return nil
	default:
		return errors.New("KBS mode must be provider or tenant")
	}
}

func validateTenantKBSURL(value string) error {
	parsed, err := url.Parse(value)
	if err != nil || len(value) > 2048 || value != strings.TrimSpace(value) ||
		parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil ||
		parsed.Path != "" || parsed.RawPath != "" || parsed.Opaque != "" ||
		parsed.RawQuery != "" || parsed.ForceQuery || parsed.Fragment != "" {
		return errors.New("tenant KBS URL must be a canonical HTTPS origin")
	}
	return nil
}

func validateImageSecurityPolicyURI(value string) error {
	parsed, err := url.Parse(value)
	if err != nil || len(value) > 1024 || value != strings.TrimSpace(value) ||
		parsed.Scheme != "kbs" || parsed.Host != "" || parsed.User != nil ||
		parsed.RawPath != "" || parsed.Opaque != "" || parsed.RawQuery != "" ||
		parsed.ForceQuery || parsed.Fragment != "" {
		return errors.New("image security policy URI must be a canonical kbs:/// resource URI")
	}
	parts := strings.Split(strings.TrimPrefix(parsed.Path, "/"), "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] != "security-policy" ||
		len(parts[2]) != len("sha256-")+64 || !strings.HasPrefix(parts[2], "sha256-") {
		return errors.New("image security policy URI must be content addressed")
	}
	for _, value := range strings.TrimPrefix(parts[2], "sha256-") {
		if !strings.ContainsRune("0123456789abcdef", value) {
			return errors.New("image security policy URI must end in a lowercase SHA-256 digest")
		}
	}
	return nil
}

func (params *v2ServiceKBSParams) toManifest() *manifest.KBSParams {
	if params == nil {
		return nil
	}

	result := &manifest.KBSParams{}
	if params.Mode == kbsModeProvider {
		result.Source = &manifest.KBSParams_Provider{Provider: &manifest.ProviderKBSParams{}}
	} else {
		result.Source = &manifest.KBSParams_Tenant{Tenant: &manifest.TenantKBSParams{
			URL:                    params.URL,
			Certificate:            params.Certificate,
			ImageSecurityPolicyURI: params.ImageSecurityPolicyURI,
			AgentPolicy:            params.AgentPolicy,
		}}
	}
	return result
}
