package v1

import (
	cerrors "cosmossdk.io/errors"

	attr "pkg.akt.dev/go/node/types/attributes/v1"
)

const (
	errInvalidDepositor = iota + attr.ErrLast
	errInvalidDepositSource
)

var (
	// ErrInvalidDepositor indicates an invalid chain parameter
	ErrInvalidDepositor = cerrors.Register(attr.ModuleName, errInvalidDepositor, "invalid depositor")
	// ErrInvalidDepositSource indicates invalid deposit source for the deployment
	ErrInvalidDepositSource = cerrors.Register(attr.ModuleName, errInvalidDepositSource, "invalid deposit source")
)
