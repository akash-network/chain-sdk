package v2

import (
	cerrors "cosmossdk.io/errors"
)

const (
	errPriceEntryExists uint32 = iota + 1
	errInvalidTimestamp
	errUnauthorizedWriterAddress
	errPriceStalled
	errInvalidFeedContractParams
	errInvalidFeedContractConfig
	errTWAPZeroWeight
)

var (
	// ErrPriceEntryExists is the error when price entry already exists
	ErrPriceEntryExists = cerrors.Register(ModuleName+".v2", errPriceEntryExists, "price entry exists")
	// ErrInvalidTimestamp is the error indicating invalid timestamp
	ErrInvalidTimestamp = cerrors.Register(ModuleName+".v2", errInvalidTimestamp, "invalid timestamp")
	// ErrUnauthorizedWriterAddress is the error indicating signer is not allowed to add price records
	ErrUnauthorizedWriterAddress = cerrors.Register(ModuleName+".v2", errUnauthorizedWriterAddress, "unauthorized writer address")
	// ErrPriceStalled is the error when price data is stale
	ErrPriceStalled = cerrors.Register(ModuleName+".v2", errPriceStalled, "price stalled")
	// ErrInvalidFeedContractParams is the error when feed contract params are invalid
	ErrInvalidFeedContractParams = cerrors.Register(ModuleName+".v2", errInvalidFeedContractParams, "invalid feed contract params")
	// ErrInvalidFeedContractConfig is the error when feed contract config is invalid
	ErrInvalidFeedContractConfig = cerrors.Register(ModuleName+".v2", errInvalidFeedContractConfig, "invalid feed contract config")
	ErrTWAPZeroWeight            = cerrors.Register(ModuleName+".v2", errTWAPZeroWeight, "invalid TWAP calculation: zero weight")
)
