package v1

import (
	sdkerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

const (
	errCodeInvalidParam uint32 = iota + 1
	errCodeInvalidDestination
	errCodeInvalidRecord
	errCodeInvalidObservation
	errCodeInvalidReason
	errCodeInvalidReporters
	errCodeInvalidSettlement
	errCodeInvalidGenesis
	errCodeInvalidGuardianPause
	errCodeInvalidResolution
	errCodeOverflow
	errCodeTransferNotFound
	errCodeTransferExists
	errCodeTransferClosed
	errCodeDepositsDisabled
	errCodeRefundsDisabled
	errCodeBelowMinTransfer
	errCodeAboveMaxTransfer
	errCodeAccountLimitReached
	errCodeAccountDailyAmountExceeded
	errCodeRefundCapReached
	errCodeNotReporter
	errCodeNotGuardian
	errCodeStaleReport
	errCodeNoPendingVoid
	errCodeChallengeWindowEnded
	errCodeChallengeWindowNotElapsed
	errCodeRefundOutOfOrder
	errCodeDepositNotSole
	errCodeSenderNotSigningAccount
	errCodeEscrowTransferNotAuthorized
	errCodeBalanceMismatch
	errCodeInternal
	errCodeInvalidID
	errCodeInvalidClock
	errCodeSendCapReached
	errCodeInvalidPagination
	errCodeSettlementHeld
)

var (
	// ErrInvalidParam is the error when a module parameter is invalid
	ErrInvalidParam = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidParam, codes.InvalidArgument, "invalid parameter")
	// ErrInvalidDestination is the error when a Solana address is invalid
	ErrInvalidDestination = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidDestination, codes.InvalidArgument, "invalid destination")
	// ErrInvalidRecord is the error when a transfer record cannot be encoded or decoded
	ErrInvalidRecord = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidRecord, codes.InvalidArgument, "invalid transfer record")
	// ErrInvalidObservation is the error when a Solana observation is invalid
	ErrInvalidObservation = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidObservation, codes.InvalidArgument, "invalid observation")
	// ErrInvalidReason is the error when a veto or governance reason is empty or too long
	ErrInvalidReason = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidReason, codes.InvalidArgument, "invalid reason")
	// ErrInvalidReporters is the error when the reporter set or threshold is invalid
	ErrInvalidReporters = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidReporters, codes.InvalidArgument, "invalid reporter set")
	// ErrInvalidSettlement is the error when a settlement is internally inconsistent
	ErrInvalidSettlement = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidSettlement, codes.InvalidArgument, "invalid settlement")
	// ErrInvalidGenesis is the error when the genesis state is inconsistent
	ErrInvalidGenesis = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidGenesis, codes.InvalidArgument, "invalid genesis")
	// ErrInvalidGuardianPause is the error when a guardian pause turns nothing off
	ErrInvalidGuardianPause = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidGuardianPause, codes.InvalidArgument, "guardian pause must disable deposits, refunds or both")
	// ErrInvalidResolution is the error when a governance resolution is invalid for the transfer
	ErrInvalidResolution = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidResolution, codes.InvalidArgument, "invalid resolution")
	// ErrOverflow is the error when an amount or counter would overflow
	ErrOverflow = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeOverflow, codes.OutOfRange, "overflow")
	// ErrTransferNotFound is the error when a transfer does not exist
	ErrTransferNotFound = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeTransferNotFound, codes.NotFound, "transfer not found")
	// ErrTransferExists is the error when a transfer record already exists
	ErrTransferExists = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeTransferExists, codes.AlreadyExists, "transfer already exists")
	// ErrTransferClosed is the error when a transfer is already minted or refunded
	ErrTransferClosed = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeTransferClosed, codes.FailedPrecondition, "transfer closed")
	// ErrDepositsDisabled is the error when deposits are disabled
	ErrDepositsDisabled = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeDepositsDisabled, codes.FailedPrecondition, "deposits disabled")
	// ErrRefundsDisabled is the error when refunds are disabled
	ErrRefundsDisabled = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeRefundsDisabled, codes.FailedPrecondition, "refunds disabled")
	// ErrBelowMinTransfer is the error when a deposit is below the minimum transfer
	ErrBelowMinTransfer = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeBelowMinTransfer, codes.FailedPrecondition, "amount below minimum transfer")
	// ErrAboveMaxTransfer is the error when a deposit is above the maximum transfer
	ErrAboveMaxTransfer = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeAboveMaxTransfer, codes.FailedPrecondition, "amount above maximum transfer")
	// ErrAccountLimitReached retains the code for removed account count limits.
	ErrAccountLimitReached = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeAccountLimitReached, codes.ResourceExhausted, "account deposit limit reached")
	// ErrAccountDailyAmountExceeded is the error when a deposit exceeds the account's remaining epoch amount
	ErrAccountDailyAmountExceeded = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeAccountDailyAmountExceeded, codes.ResourceExhausted, "account epoch amount exceeded")
	// ErrRefundCapReached is the error when a refund exceeds the remaining epoch refund cap
	ErrRefundCapReached = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeRefundCapReached, codes.ResourceExhausted, "epoch refund cap reached")
	// ErrNotReporter is the error when the signer is not a current reporter
	ErrNotReporter = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeNotReporter, codes.PermissionDenied, "not a current reporter")
	// ErrNotGuardian is the error when the signer is not the pause guardian
	ErrNotGuardian = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeNotGuardian, codes.PermissionDenied, "not the pause guardian")
	// ErrStaleReport is the error when a report or authorization is from an earlier reporter policy version or round
	ErrStaleReport = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeStaleReport, codes.FailedPrecondition, "stale reporter policy or round")
	// ErrNoPendingVoid is the error when a transfer has no confirmed void to veto
	ErrNoPendingVoid = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeNoPendingVoid, codes.FailedPrecondition, "no pending void")
	// ErrChallengeWindowEnded is the error when the void challenge window has ended
	ErrChallengeWindowEnded = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeChallengeWindowEnded, codes.FailedPrecondition, "challenge window ended")
	// ErrChallengeWindowNotElapsed is the error when the void challenge window has not elapsed
	ErrChallengeWindowNotElapsed = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeChallengeWindowNotElapsed, codes.FailedPrecondition, "challenge window has not elapsed")
	// ErrRefundOutOfOrder is the error when an older refund must be paid first
	ErrRefundOutOfOrder = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeRefundOutOfOrder, codes.FailedPrecondition, "older refund must be paid first")
	// ErrDepositNotSole is the error when a deposit is not the only top-level transaction message
	ErrDepositNotSole = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeDepositNotSole, codes.InvalidArgument, "deposit must be the only transaction message")
	// ErrSenderNotSigningAccount is the error when the sender is a module account or has no public key
	ErrSenderNotSigningAccount = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeSenderNotSigningAccount, codes.PermissionDenied, "sender must be an ordinary signing account")
	// ErrEscrowTransferNotAuthorized is the error when a bank send into or out of the sink escrow is not authorized
	ErrEscrowTransferNotAuthorized = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeEscrowTransferNotAuthorized, codes.PermissionDenied, "escrow transfer not authorized")
	// ErrBalanceMismatch is the error when the sink bank balance does not match the escrow totals
	ErrBalanceMismatch = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeBalanceMismatch, codes.Internal, "bank balance does not match escrow")
	// ErrInternal is the error for internal error
	ErrInternal = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInternal, codes.Internal, "internal error")
	// ErrInvalidID is the error when a composite identity is malformed.
	ErrInvalidID = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidID, codes.InvalidArgument, "invalid transfer ID")
	// ErrInvalidClock is the error for a negative time or backward epoch.
	ErrInvalidClock = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidClock, codes.InvalidArgument, "invalid epoch clock")
	// ErrSendCapReached is the error when aggregate sending exhausts its epoch cap.
	ErrSendCapReached = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeSendCapReached, codes.ResourceExhausted, "send epoch cap reached")
	// ErrInvalidPagination is the error for unsupported limits or cursor scans.
	ErrInvalidPagination = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeInvalidPagination, codes.InvalidArgument, "invalid pagination")
	// ErrSettlementHeld is the error when a settlement hold requires governance resolution.
	ErrSettlementHeld = sdkerrors.RegisterWithGRPCCode(ModuleName, errCodeSettlementHeld, codes.FailedPrecondition, "settlement requires governance resolution")
)
