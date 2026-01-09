package v1

import (
	"fmt"
	"time"
)

const (
	DefaultOracleOutlierThresholdBps   = uint32(150)
	DefaultSettlementEpochBlocks       = int64(1)
	DefaultCircuitBreakerWarnThreshold = uint32(9500)
	DefaultCircuitBreakerHaltThreshold = uint32(9000)
	DefaultMintSpreadBps               = uint32(25)
	DefaultSettleSpreadBps             = uint32(0)
	DefaultMinRunwayBlocks             = int64(7 * ((time.Hour * 24) / (6 * time.Second)))
	DefaultEnabled                     = true
)

var DefaultOracleTWAPWindow = time.Hour

func DefaultParams() Params {
	return Params{
		OracleTwapWindow:            DefaultOracleTWAPWindow,
		OracleOutlierThresholdBps:   DefaultOracleOutlierThresholdBps,
		SettlementEpochBlocks:       DefaultSettlementEpochBlocks,
		CircuitBreakerWarnThreshold: DefaultCircuitBreakerWarnThreshold,
		CircuitBreakerHaltThreshold: DefaultCircuitBreakerHaltThreshold,
		MintSpreadBps:               DefaultMintSpreadBps,
		SettleSpreadBps:             DefaultSettleSpreadBps,
		MinRunwayBlocks:             DefaultMinRunwayBlocks,
		Enabled:                     DefaultEnabled,
	}
}

func (p Params) Validate() error {
	if p.OracleTwapWindow <= 0 {
		return fmt.Errorf("oracle_twap_window must be positive")
	}
	if p.OracleOutlierThresholdBps > 10000 {
		return fmt.Errorf("oracle_outlier_threshold_bps cannot exceed 10000")
	}
	if p.SettlementEpochBlocks <= 0 {
		return fmt.Errorf("settlement_epoch_blocks must be positive")
	}
	if p.MinRunwayBlocks <= 0 {
		return fmt.Errorf("min_runway_blocks must be positive")
	}
	if p.CircuitBreakerWarnThreshold > 10000 {
		return fmt.Errorf("circuit_breaker_warn_threshold cannot exceed 10000")
	}
	if p.CircuitBreakerHaltThreshold > 10000 {
		return fmt.Errorf("circuit_breaker_halt_threshold cannot exceed 10000")
	}
	if p.CircuitBreakerWarnThreshold <= p.CircuitBreakerHaltThreshold {
		return fmt.Errorf("warn threshold must be greater than halt threshold")
	}
	if p.MintSpreadBps > 1000 {
		return fmt.Errorf("mint_spread_bps cannot exceed 1000 (10%%)")
	}
	if p.SettleSpreadBps > 1000 {
		return fmt.Errorf("settle_spread_bps cannot exceed 1000 (10%%)")
	}
	return nil
}
