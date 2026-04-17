package client

import (
	"testing"
)

func TestNegotiateVersion_MultiVersion(t *testing.T) {
	tests := []struct {
		name     string
		result   *Akash
		expected string
	}{
		{
			name: "new node with both versions, client picks v1beta4",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: VersionV1beta3},
				SupportedVersions: []VersionInfo{
					{ApiVersion: VersionV1beta4},
					{ApiVersion: VersionV1beta3},
				},
			},
			expected: VersionV1beta4,
		},
		{
			name: "new node with only v1beta3, client falls back",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: VersionV1beta3},
				SupportedVersions: []VersionInfo{
					{ApiVersion: VersionV1beta3},
				},
			},
			expected: VersionV1beta3,
		},
		{
			name: "new node with unknown versions only",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: VersionV1beta3},
				SupportedVersions: []VersionInfo{
					{ApiVersion: "v99"},
				},
			},
			expected: "",
		},
		{
			name: "new node with future + known versions",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: VersionV1beta3},
				SupportedVersions: []VersionInfo{
					{ApiVersion: "v1"},
					{ApiVersion: VersionV1beta4},
					{ApiVersion: VersionV1beta3},
				},
			},
			expected: VersionV1beta4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := negotiateVersion(tt.result)
			if got != tt.expected {
				t.Errorf("negotiateVersion() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestNegotiateVersion_Legacy(t *testing.T) {
	tests := []struct {
		name     string
		result   *Akash
		expected string
	}{
		{
			name: "old node returns v1beta3 only",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: VersionV1beta3},
			},
			expected: VersionV1beta3,
		},
		{
			name: "old node returns unknown version",
			result: &Akash{
				ClientInfo: ClientInfo{ApiVersion: "v0"},
			},
			expected: "v0",
		},
		{
			name: "empty SupportedVersions falls back to ClientInfo",
			result: &Akash{
				ClientInfo:        ClientInfo{ApiVersion: VersionV1beta4},
				SupportedVersions: []VersionInfo{},
			},
			expected: VersionV1beta4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := negotiateVersion(tt.result)
			if got != tt.expected {
				t.Errorf("negotiateVersion() = %q, want %q", got, tt.expected)
			}
		})
	}
}
