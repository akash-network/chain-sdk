package client

import "sync"

// VersionRegistry holds the server's supported API versions and metadata.
// It is used by both the CometBFT JSON-RPC handler and the gRPC Discovery service
// to produce version responses.
type VersionRegistry struct {
	mu                sync.RWMutex
	supportedVersions []VersionInfo
	chainID           string
	nodeVersion       string
	minClientVersion  string
}

// RegistryOption configures a VersionRegistry.
type RegistryOption func(*VersionRegistry)

// WithChainID sets the chain ID on the registry.
func WithChainID(chainID string) RegistryOption {
	return func(r *VersionRegistry) {
		r.chainID = chainID
	}
}

// WithNodeVersion sets the node software version on the registry.
func WithNodeVersion(version string) RegistryOption {
	return func(r *VersionRegistry) {
		r.nodeVersion = version
	}
}

// WithMinClientVersion sets the minimum required client version.
func WithMinClientVersion(version string) RegistryOption {
	return func(r *VersionRegistry) {
		r.minClientVersion = version
	}
}

// NewRegistry creates a new VersionRegistry with the given supported versions.
func NewRegistry(versions []VersionInfo, opts ...RegistryOption) *VersionRegistry {
	r := &VersionRegistry{
		supportedVersions: copyVersions(versions),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// v1beta3Modules defines the per-module versions for the v1beta3 API version.
var v1beta3Modules = []ModuleVersion{
	{Module: "audit", Version: "v1"},
	{Module: "bme", Version: "v1"},
	{Module: "cert", Version: "v1"},
	{Module: "deployment", Version: "v1beta4"},
	{Module: "escrow", Version: "v1"},
	{Module: "market", Version: "v1beta5"},
	{Module: "oracle", Version: "v2"},
	{Module: "provider", Version: "v1beta4"},
	{Module: "take", Version: "v1"},
	{Module: "wasm", Version: "v1"},
}

// v1beta4Modules defines the per-module versions for the v1beta4 API version.
// Currently identical to v1beta3 — will diverge when module versions are bumped.
var v1beta4Modules = []ModuleVersion{
	{Module: "audit", Version: "v1"},
	{Module: "bme", Version: "v1"},
	{Module: "cert", Version: "v1"},
	{Module: "deployment", Version: "v1beta4"},
	{Module: "escrow", Version: "v1"},
	{Module: "market", Version: "v1beta5"},
	{Module: "oracle", Version: "v2"},
	{Module: "provider", Version: "v1beta4"},
	{Module: "take", Version: "v1"},
	{Module: "wasm", Version: "v1"},
}

// DefaultRegistry creates a registry with v1beta3 and v1beta4 support.
func DefaultRegistry(opts ...RegistryOption) *VersionRegistry {
	return NewRegistry([]VersionInfo{
		{
			ApiVersion: VersionV1beta4,
			Modules:    v1beta4Modules,
		},
		{
			ApiVersion: VersionV1beta3,
			Modules:    v1beta3Modules,
		},
	}, opts...)
}

// OldestVersion returns the lowest supported version string.
// Used as the value for ClientInfo.ApiVersion for backward compatibility with old clients.
func (r *VersionRegistry) OldestVersion() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.supportedVersions) == 0 {
		return ""
	}
	return r.supportedVersions[len(r.supportedVersions)-1].ApiVersion
}

// SupportsVersion checks if a specific version is in the supported set.
func (r *VersionRegistry) SupportsVersion(v string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, sv := range r.supportedVersions {
		if sv.ApiVersion == v {
			return true
		}
	}
	return false
}

// SetChainID updates the chain ID. Safe for concurrent use.
func (r *VersionRegistry) SetChainID(chainID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.chainID = chainID
}

// copyVersions returns a deep copy of a []VersionInfo slice.
// Each element's Modules and Features sub-slices are independently copied
// so the returned slice shares no backing arrays with the original.
func copyVersions(src []VersionInfo) []VersionInfo {
	if src == nil {
		return nil
	}
	dst := make([]VersionInfo, len(src))
	for i, v := range src {
		dst[i] = VersionInfo{
			ApiVersion: v.ApiVersion,
		}
		if v.Modules != nil {
			dst[i].Modules = make([]ModuleVersion, len(v.Modules))
			copy(dst[i].Modules, v.Modules)
		}
		if v.Features != nil {
			dst[i].Features = make([]string, len(v.Features))
			copy(dst[i].Features, v.Features)
		}
	}
	return dst
}

// ToAkash converts the registry state to a proto Akash response.
// ClientInfo.ApiVersion is set to the oldest supported version for backward compatibility.
func (r *VersionRegistry) ToAkash() *Akash {
	r.mu.RLock()
	defer r.mu.RUnlock()

	oldest := ""
	if len(r.supportedVersions) > 0 {
		oldest = r.supportedVersions[len(r.supportedVersions)-1].ApiVersion
	}

	return &Akash{
		ClientInfo:        ClientInfo{ApiVersion: oldest},
		SupportedVersions: copyVersions(r.supportedVersions),
		ChainID:           r.chainID,
		NodeVersion:       r.nodeVersion,
		MinClientVersion:  r.minClientVersion,
	}
}
