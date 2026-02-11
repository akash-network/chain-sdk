package v1

import (
	"cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrInvalidParams          = errors.RegisterWithGRPCCode(ModuleName, 2, codes.InvalidArgument, "invalid parameters")
	ErrUnauthorized           = errors.RegisterWithGRPCCode(ModuleName, 3, codes.Unauthenticated, "unauthorized")
	ErrOracleUnhealthy        = errors.RegisterWithGRPCCode(ModuleName, 5, codes.Aborted, "oracle is unhealthy")
	ErrCircuitBreakerActive   = errors.RegisterWithGRPCCode(ModuleName, 6, codes.Aborted, "circuit breaker is active")
	ErrInsufficientVaultFunds = errors.RegisterWithGRPCCode(ModuleName, 7, codes.Aborted, "insufficient vault funds")
	ErrInvalidAmount          = errors.RegisterWithGRPCCode(ModuleName, 8, codes.InvalidArgument, "invalid amount")
	ErrInvalidDenom           = errors.RegisterWithGRPCCode(ModuleName, 9, codes.InvalidArgument, "invalid denomination")
	ErrInvalidAddress         = errors.RegisterWithGRPCCode(ModuleName, 10, codes.InvalidArgument, "invalid address")
	ErrInvalidSender          = errors.RegisterWithGRPCCode(ModuleName, 11, codes.InvalidArgument, "invalid sender")
	ErrZeroPrice              = errors.RegisterWithGRPCCode(ModuleName, 12, codes.Internal, "oracle price is zero")
	ErrRecordExists           = errors.RegisterWithGRPCCode(ModuleName, 13, codes.AlreadyExists, "ledger record already exists")
	ErrMintFailed             = errors.RegisterWithGRPCCode(ModuleName, 14, codes.Aborted, "failed to mint tokens")
	ErrBurnFailed             = errors.RegisterWithGRPCCode(ModuleName, 15, codes.Aborted, "failed to burn tokens")
)
