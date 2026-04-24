package client

import (
	"testing"
)

func TestDefaultRegistry(t *testing.T) {
	r := DefaultRegistry()

	if !r.SupportsVersion(VersionV1beta3) {
		t.Error("default registry should support v1beta3")
	}
	if !r.SupportsVersion(VersionV1beta4) {
		t.Error("default registry should support v1beta4")
	}
	if r.SupportsVersion("v999") {
		t.Error("default registry should not support v999")
	}
}

func TestRegistryOldestVersion(t *testing.T) {
	r := DefaultRegistry()

	oldest := r.OldestVersion()
	if oldest != VersionV1beta3 {
		t.Errorf("expected oldest version %q, got %q", VersionV1beta3, oldest)
	}
}

func TestRegistryOldestVersionEmpty(t *testing.T) {
	r := NewRegistry(nil)
	if oldest := r.OldestVersion(); oldest != "" {
		t.Errorf("expected empty string for empty registry, got %q", oldest)
	}
}

func TestRegistryToAkash(t *testing.T) {
	r := DefaultRegistry(
		WithChainID("akashnet-2"),
		WithNodeVersion("v0.36.0"),
		WithMinClientVersion("v0.30.0"),
	)

	akash := r.ToAkash()

	// ClientInfo.ApiVersion should be the oldest supported version for backward compat
	if akash.ClientInfo.ApiVersion != VersionV1beta3 {
		t.Errorf("expected ClientInfo.ApiVersion=%q, got %q", VersionV1beta3, akash.ClientInfo.ApiVersion)
	}

	if len(akash.SupportedVersions) != 2 {
		t.Fatalf("expected 2 supported versions, got %d", len(akash.SupportedVersions))
	}

	// First entry should be the newest version
	if akash.SupportedVersions[0].ApiVersion != VersionV1beta4 {
		t.Errorf("expected first supported version %q, got %q", VersionV1beta4, akash.SupportedVersions[0].ApiVersion)
	}
	if akash.SupportedVersions[1].ApiVersion != VersionV1beta3 {
		t.Errorf("expected second supported version %q, got %q", VersionV1beta3, akash.SupportedVersions[1].ApiVersion)
	}

	if akash.ChainID != "akashnet-2" {
		t.Errorf("expected ChainID=%q, got %q", "akashnet-2", akash.ChainID)
	}
	if akash.NodeVersion != "v0.36.0" {
		t.Errorf("expected NodeVersion=%q, got %q", "v0.36.0", akash.NodeVersion)
	}
	if akash.MinClientVersion != "v0.30.0" {
		t.Errorf("expected MinClientVersion=%q, got %q", "v0.30.0", akash.MinClientVersion)
	}
}

func TestRegistryToAkashModules(t *testing.T) {
	r := DefaultRegistry()
	akash := r.ToAkash()

	for _, vi := range akash.SupportedVersions {
		if len(vi.Modules) == 0 {
			t.Errorf("version %q should have module list", vi.ApiVersion)
		}

		// Check that deployment module is present
		found := false
		for _, m := range vi.Modules {
			if m.Module == "deployment" {
				found = true
				if m.Version != "v1beta4" {
					t.Errorf("version %q: expected deployment module version %q, got %q", vi.ApiVersion, "v1beta4", m.Version)
				}
			}
		}
		if !found {
			t.Errorf("version %q: deployment module not found", vi.ApiVersion)
		}
	}
}

func TestRegistrySetChainID(t *testing.T) {
	r := DefaultRegistry()
	r.SetChainID("testnet-1")

	akash := r.ToAkash()
	if akash.ChainID != "testnet-1" {
		t.Errorf("expected ChainID=%q after SetChainID, got %q", "testnet-1", akash.ChainID)
	}
}

func TestNewRegistryDefensiveCopy(t *testing.T) {
	versions := []VersionInfo{
		{ApiVersion: "v2", Modules: []ModuleVersion{{Module: "m", Version: "v1"}}},
		{ApiVersion: "v1"},
	}
	r := NewRegistry(versions)

	// Mutate the caller's slice and its nested fields.
	versions[0].ApiVersion = "MUTATED"
	versions[0].Modules[0].Module = "MUTATED"

	// Registry must be unaffected.
	if !r.SupportsVersion("v2") {
		t.Error("registry should still support v2 after caller mutates input slice")
	}
	akash := r.ToAkash()
	if akash.SupportedVersions[0].Modules[0].Module != "m" {
		t.Error("registry module name should be unaffected by caller mutation")
	}
}

func TestToAkashDefensiveCopy(t *testing.T) {
	r := DefaultRegistry()
	akash := r.ToAkash()

	// Mutate the returned slice.
	akash.SupportedVersions[0].ApiVersion = "MUTATED"

	// A subsequent call must return the original value.
	akash2 := r.ToAkash()
	if akash2.SupportedVersions[0].ApiVersion == "MUTATED" {
		t.Error("mutating ToAkash result should not affect registry internal state")
	}
}

func TestNewRegistryCustomVersions(t *testing.T) {
	versions := []VersionInfo{
		{ApiVersion: "v2"},
		{ApiVersion: "v1"},
	}
	r := NewRegistry(versions)

	if !r.SupportsVersion("v2") {
		t.Error("should support v2")
	}
	if !r.SupportsVersion("v1") {
		t.Error("should support v1")
	}
	if r.SupportsVersion(VersionV1beta3) {
		t.Error("should not support v1beta3")
	}
	if r.OldestVersion() != "v1" {
		t.Errorf("expected oldest %q, got %q", "v1", r.OldestVersion())
	}
}
