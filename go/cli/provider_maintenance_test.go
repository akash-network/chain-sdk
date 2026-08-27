package cli

import (
	"testing"

	"github.com/stretchr/testify/require"

	types "pkg.akt.dev/go/node/provider/v1beta4"
)

func TestProviderMaintenanceCommandSurface(t *testing.T) {
	tx := requireSubcommand(t, TxCmd(), "provider")
	requireSubcommand(t, tx, "open-maintenance")
	requireSubcommand(t, tx, "close-maintenance")

	query := requireSubcommand(t, QueryCmd(), "provider")
	requireSubcommand(t, query, "maintenances")
	requireSubcommand(t, query, "maintenance")
}

func TestParseProviderMaintenanceType(t *testing.T) {
	testCases := []struct {
		input string
		want  types.ProviderMaintenanceType
	}{
		{"planned", types.ProviderMaintenanceType_provider_maintenance_type_planned},
		{"emergency", types.ProviderMaintenanceType_provider_maintenance_type_emergency},
		{"provider_maintenance_type_security", types.ProviderMaintenanceType_provider_maintenance_type_security},
		{"network", types.ProviderMaintenanceType_provider_maintenance_type_network},
		{"capacity", types.ProviderMaintenanceType_provider_maintenance_type_capacity},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseProviderMaintenanceType(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseProviderMaintenanceType("unknown")
	require.Error(t, err)
}

func TestParseProviderMaintenanceStatus(t *testing.T) {
	testCases := []struct {
		input string
		want  types.ProviderMaintenanceStatus
	}{
		{"scheduled", types.ProviderMaintenanceStatus_provider_maintenance_status_scheduled},
		{"active", types.ProviderMaintenanceStatus_provider_maintenance_status_active},
		{"provider_maintenance_status_elapsed", types.ProviderMaintenanceStatus_provider_maintenance_status_elapsed},
		{"closed", types.ProviderMaintenanceStatus_provider_maintenance_status_closed},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseProviderMaintenanceStatus(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}

	_, err := parseProviderMaintenanceStatus("unknown")
	require.Error(t, err)
}
