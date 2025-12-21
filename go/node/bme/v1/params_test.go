package v1_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	bmev1 "pkg.akt.dev/go/node/bme/v1"
)

func TestParamsValidate(t *testing.T) {
	tests := []struct {
		name      string
		params    bmev1.Params
		expectErr bool
		errMsg    string
	}{
		{
			name:      "default params are valid",
			params:    bmev1.DefaultParams(),
			expectErr: false,
		},
		{
			name: "negative oracle_twap_window",
			params: bmev1.Params{
				OracleTwapWindow:            -1 * time.Second,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "oracle_twap_window must be positive",
		},
		{
			name: "oracle_outlier_threshold_bps exceeds 10000",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   10001,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "oracle_outlier_threshold_bps cannot exceed 10000",
		},
		{
			name: "settlement_epoch_blocks is zero",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       0,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "settlement_epoch_blocks must be positive",
		},
		{
			name: "settlement_epoch_blocks is negative",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       -1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "settlement_epoch_blocks must be positive",
		},
		{
			name: "min_runway_blocks is zero",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             0,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "min_runway_blocks must be positive",
		},
		{
			name: "min_runway_blocks is negative",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             -100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "min_runway_blocks must be positive",
		},
		{
			name: "circuit_breaker_warn_threshold exceeds 10000",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 10001,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "circuit_breaker_warn_threshold cannot exceed 10000",
		},
		{
			name: "circuit_breaker_halt_threshold exceeds 10000",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 10001,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "circuit_breaker_halt_threshold cannot exceed 10000",
		},
		{
			name: "warn threshold not greater than halt threshold (equal)",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9000,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "warn threshold must be greater than halt threshold",
		},
		{
			name: "warn threshold not greater than halt threshold (less)",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 8000,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "warn threshold must be greater than halt threshold",
		},
		{
			name: "mint_spread_bps exceeds 1000",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               1001,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "mint_spread_bps cannot exceed 1000 (10%)",
		},
		{
			name: "settle_spread_bps exceeds 1000",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 9500,
				CircuitBreakerHaltThreshold: 9000,
				MintSpreadBps:               25,
				SettleSpreadBps:             1001,
				Enabled:                     true,
			},
			expectErr: true,
			errMsg:    "settle_spread_bps cannot exceed 1000 (10%)",
		},
		{
			name: "circuit breaker thresholds at max valid values",
			params: bmev1.Params{
				OracleTwapWindow:            time.Hour,
				OracleOutlierThresholdBps:   150,
				SettlementEpochBlocks:       1,
				MinRunwayBlocks:             100,
				CircuitBreakerWarnThreshold: 10000,
				CircuitBreakerHaltThreshold: 9999,
				MintSpreadBps:               25,
				SettleSpreadBps:             0,
				Enabled:                     true,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
