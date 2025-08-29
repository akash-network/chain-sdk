package module

import (
	cerrors "cosmossdk.io/errors"
)

const (
	errAccountExists uint32 = iota + 1
	errAccountClosed
	errAccountNotFound
	errAccountOverdrawn
	errInvalidDenomination
	errPaymentExists
	errPaymentClosed
	errPaymentNotFound
	errPaymentRateZero
	errInvalidPayment
	errInvalidSettlement
	errInvalidID
	errInvalidAccount
	errInvalidAccountDepositor
	errUnauthorizedDepositScope
	errInvalidDeposit
	errInvalidAuthzScope
)

var (
	ErrAccountExists            = cerrors.Register(ModuleName, errAccountExists, "account exists")
	ErrAccountClosed            = cerrors.Register(ModuleName, errAccountClosed, "account closed")
	ErrAccountNotFound          = cerrors.Register(ModuleName, errAccountNotFound, "account not found")
	ErrAccountOverdrawn         = cerrors.Register(ModuleName, errAccountOverdrawn, "account overdrawn")
	ErrInvalidDenomination      = cerrors.Register(ModuleName, errInvalidDenomination, "invalid denomination")
	ErrPaymentExists            = cerrors.Register(ModuleName, errPaymentExists, "payment exists")
	ErrPaymentClosed            = cerrors.Register(ModuleName, errPaymentClosed, "payment closed")
	ErrPaymentNotFound          = cerrors.Register(ModuleName, errPaymentNotFound, "payment not found")
	ErrPaymentRateZero          = cerrors.Register(ModuleName, errPaymentRateZero, "payment rate zero")
	ErrInvalidPayment           = cerrors.Register(ModuleName, errInvalidPayment, "invalid payment")
	ErrInvalidSettlement        = cerrors.Register(ModuleName, errInvalidSettlement, "invalid settlement")
	ErrInvalidID                = cerrors.Register(ModuleName, errInvalidID, "invalid ID")
	ErrInvalidAccount           = cerrors.Register(ModuleName, errInvalidAccount, "invalid account")
	ErrInvalidAccountDepositor  = cerrors.Register(ModuleName, errInvalidAccountDepositor, "invalid account depositor")
	ErrUnauthorizedDepositScope = cerrors.Register(ModuleName, errUnauthorizedDepositScope, "unauthorized deposit scope")
	ErrInvalidDeposit           = cerrors.Register(ModuleName, errInvalidDeposit, "invalid deposit")
	ErrInvalidAuthzScope        = cerrors.Register(ModuleName, errInvalidAuthzScope, "invalid authz scope")
)
