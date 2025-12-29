package v1

import (
	cerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

const (
	errNameDoesNotExist uint32 = iota + 1
	errInvalidRequest
	errDeploymentExists
	errDeploymentNotFound
	errDeploymentClosed
	errOwnerAcctMissing
	errInvalidGroups
	errInvalidDeploymentID
	errEmptyHash
	errInvalidHash
	errInternal
	errInvalidDeployment
	errInvalidGroupID
	errGroupNotFound
	errGroupClosed
	errGroupOpen
	errGroupPaused
	errGroupNotOpen
	errGroupSpec
	errInvalidDeposit
	errInvalidIDPath
	errInvalidParam
	errInvalidEscrowID
)

var (
	// ErrNameDoesNotExist is the error when name does not exist
	ErrNameDoesNotExist = cerrors.RegisterWithGRPCCode(ModuleName, errNameDoesNotExist, codes.NotFound, "Name does not exist")
	// ErrInvalidRequest is the error for invalid request
	ErrInvalidRequest = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidRequest, codes.InvalidArgument, "Invalid request")
	// ErrDeploymentExists is the error when already deployment exists
	ErrDeploymentExists = cerrors.RegisterWithGRPCCode(ModuleName, errDeploymentExists, codes.AlreadyExists, "Deployment exists")
	// ErrDeploymentNotFound is the error when deployment not found
	ErrDeploymentNotFound = cerrors.RegisterWithGRPCCode(ModuleName, errDeploymentNotFound, codes.NotFound, "Deployment not found")
	// ErrDeploymentClosed is the error when deployment is closed
	ErrDeploymentClosed = cerrors.RegisterWithGRPCCode(ModuleName, errDeploymentClosed, codes.FailedPrecondition, "Deployment closed")
	// ErrOwnerAcctMissing is the error for owner account missing
	ErrOwnerAcctMissing = cerrors.RegisterWithGRPCCode(ModuleName, errOwnerAcctMissing, codes.InvalidArgument, "Owner account missing")
	// ErrInvalidGroups is the error when groups are empty
	ErrInvalidGroups = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidGroups, codes.InvalidArgument, "Invalid groups")
	// ErrInvalidDeploymentID is the error for invalid deployment id
	ErrInvalidDeploymentID = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidDeploymentID, codes.InvalidArgument, "Invalid: deployment id")
	// ErrEmptyHash is the error when version is empty
	ErrEmptyHash = cerrors.RegisterWithGRPCCode(ModuleName, errEmptyHash, codes.InvalidArgument, "Invalid: empty hash")
	// ErrInvalidHash is the error when version is invalid
	ErrInvalidHash = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidHash, codes.InvalidArgument, "Invalid: deployment hash")
	// ErrInternal is the error for internal error
	ErrInternal = cerrors.RegisterWithGRPCCode(ModuleName, errInternal, codes.Internal, "internal error")
	// ErrInvalidDeployment = is the error when deployment does not pass validation
	ErrInvalidDeployment = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidDeployment, codes.InvalidArgument, "Invalid deployment")
	// ErrInvalidGroupID is the error when the deployment's group ID is invalid
	ErrInvalidGroupID = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidGroupID, codes.InvalidArgument, "Invalid deployment's group ID")
	// ErrGroupNotFound is the keeper's error for not finding a group
	ErrGroupNotFound = cerrors.RegisterWithGRPCCode(ModuleName, errGroupNotFound, codes.NotFound, "Group not found")
	// ErrGroupClosed is the error when deployment is closed
	ErrGroupClosed = cerrors.RegisterWithGRPCCode(ModuleName, errGroupClosed, codes.FailedPrecondition, "Group already closed")
	// ErrGroupOpen is the error when deployment is closed
	ErrGroupOpen = cerrors.RegisterWithGRPCCode(ModuleName, errGroupOpen, codes.FailedPrecondition, "Group open")
	// ErrGroupPaused is the error when deployment is closed
	ErrGroupPaused = cerrors.RegisterWithGRPCCode(ModuleName, errGroupPaused, codes.FailedPrecondition, "Group paused")
	// ErrGroupNotOpen indicates the Group state has progressed beyond initial Open.
	ErrGroupNotOpen = cerrors.RegisterWithGRPCCode(ModuleName, errGroupNotOpen, codes.FailedPrecondition, "Group not open")
	// ErrGroupSpecInvalid indicates a GroupSpec has invalid configuration
	ErrGroupSpecInvalid = cerrors.RegisterWithGRPCCode(ModuleName, errGroupSpec, codes.InvalidArgument, "GroupSpec invalid")
	// ErrInvalidDeposit indicates an invalid deposit
	ErrInvalidDeposit = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidDeposit, codes.InvalidArgument, "Deposit invalid")
	// ErrInvalidIDPath indicates an invalid ID path
	ErrInvalidIDPath = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidIDPath, codes.InvalidArgument, "ID path invalid")
	// ErrInvalidParam indicates an invalid chain parameter
	ErrInvalidParam = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidParam, codes.InvalidArgument, "parameter invalid")
	// ErrInvalidEscrowID indicates an invalid escrow ID
	ErrInvalidEscrowID = cerrors.RegisterWithGRPCCode(ModuleName, errInvalidEscrowID, codes.InvalidArgument, "invalid escrow id")
)
