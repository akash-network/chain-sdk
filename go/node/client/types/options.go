package types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

const (
	// SignModeDirect is the value of the --sign-mode flag for SIGN_MODE_DIRECT
	SignModeDirect = "direct"
	// SignModeLegacyAminoJSON is the value of the --sign-mode flag for SIGN_MODE_LEGACY_AMINO_JSON
	SignModeLegacyAminoJSON = "amino-json"
	// SignModeDirectAux is the value of the --sign-mode flag for SIGN_MODE_DIRECT_AUX
	SignModeDirectAux = "direct-aux"
	// SignModeEIP191 is the value of the --sign-mode flag for SIGN_MODE_EIP_191
	SignModeEIP191 = "eip-191"
)

// GasSetting encapsulates the possible values passed through the --gas flag.
type GasSetting struct {
	Simulate bool
	Gas      uint64
}

func (v *GasSetting) String() string {
	if v.Simulate {
		return "auto"
	}

	return strconv.FormatUint(v.Gas, 10)
}

type ClientOptions struct {
	AccountNumber    uint64
	AccountSequence  uint64
	GasAdjustment    float64
	Gas              GasSetting
	GasPrices        string
	Fees             string
	Note             string
	TimeoutHeight    uint64
	BroadcastTimeout time.Duration
	SkipConfirm      bool
	SignMode         string
}

type BroadcastOptions struct {
	generateOnly     *bool
	timeoutHeight    *uint64
	gasAdjustment    *float64
	gas              *GasSetting
	gasPrices        *string
	fees             *string
	note             *string
	broadcastTimeout time.Duration
	resultAsError    bool
	skipConfirm      *bool
	confirmFn        ConfirmFn
	broadcastMode    *string
}

type ConfirmFn func(string) (bool, error)

// BroadcastOption is a function that takes as first argument a pointer to BroadcastOptions and returns an error
// if the option cannot be configured. A number of BroadcastOption functions are available in this package.
type BroadcastOption func(*BroadcastOptions) error

type ClientOption func(options *ClientOptions) error

// NewTxFactory creates a new Factory.
func NewTxFactory(cctx client.Context, opts ...ClientOption) (tx.Factory, error) {
	clOpts := &ClientOptions{}

	for _, opt := range opts {
		if err := opt(clOpts); err != nil {
			return tx.Factory{}, err
		}
	}

	txf := tx.Factory{}.
		WithTxConfig(cctx.TxConfig).
		WithAccountRetriever(cctx.AccountRetriever).
		WithAccountNumber(clOpts.AccountNumber).
		WithSequence(clOpts.AccountSequence).
		WithKeybase(cctx.Keyring).
		WithChainID(cctx.ChainID).
		WithGas(clOpts.Gas.Gas).
		WithGasAdjustment(clOpts.GasAdjustment).
		WithGasPrices(clOpts.GasPrices).
		WithSimulateAndExecute(clOpts.Gas.Simulate).
		WithTimeoutHeight(clOpts.TimeoutHeight).
		WithMemo(clOpts.Note).
		WithFees(clOpts.Fees).
		WithFromName(cctx.FromName)

	if !cctx.GenerateOnly {
		var signMode signing.SignMode

		switch cctx.SignModeStr {
		case SignModeDirect:
			signMode = signing.SignMode_SIGN_MODE_DIRECT
		case SignModeDirectAux:
			signMode = signing.SignMode_SIGN_MODE_DIRECT_AUX
		case SignModeLegacyAminoJSON:
			signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
		case SignModeEIP191:
			signMode = signing.SignMode_SIGN_MODE_EIP_191
		default:
			return tx.Factory{}, fmt.Errorf("invalid sign mode \"%s\". expected %s|%s|%s|%s",
				cctx.SignModeStr,
				SignModeDirect,
				SignModeDirectAux,
				SignModeLegacyAminoJSON,
				SignModeEIP191)
		}

		txf = txf.WithSignMode(signMode)
	}

	if !cctx.Offline && cctx.From != "" {
		address := cctx.GetFromAddress()

		if err := txf.AccountRetriever().EnsureExists(cctx, address); err != nil {
			return txf, err
		}

		if txf.AccountNumber() == 0 || txf.Sequence() == 0 {
			num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(cctx, address)
			if err != nil {
				return txf, err
			}

			txf = txf.WithAccountNumber(num).WithSequence(seq)
		}
	}

	return txf, nil
}

func WithAccountNumber(val uint64) ClientOption {
	return func(options *ClientOptions) error {
		options.AccountNumber = val
		return nil
	}
}

func WithAccountSequence(val uint64) ClientOption {
	return func(options *ClientOptions) error {
		options.AccountSequence = val
		return nil
	}
}

func WithGasAdjustment(val float64) ClientOption {
	return func(options *ClientOptions) error {
		options.GasAdjustment = val
		return nil
	}
}

func WithNote(val string) ClientOption {
	return func(options *ClientOptions) error {
		options.Note = val
		return nil
	}
}

func WithGas(val GasSetting) ClientOption {
	return func(options *ClientOptions) error {
		options.Gas = val
		return nil
	}
}

func WithGasPrices(val string) ClientOption {
	return func(options *ClientOptions) error {
		options.GasPrices = val
		return nil
	}
}

func WithFees(val string) ClientOption {
	return func(options *ClientOptions) error {
		options.Fees = val
		return nil
	}
}

func WithTimeoutHeight(val uint64) ClientOption {
	return func(options *ClientOptions) error {
		options.TimeoutHeight = val
		return nil
	}
}

func WithSkipConfirm(val bool) ClientOption {
	return func(options *ClientOptions) error {
		options.SkipConfirm = val
		return nil
	}
}

func WithSignMode(val string) ClientOption {
	return func(options *ClientOptions) error {
		options.SignMode = val
		return nil
	}
}

//// WithGasAdjustment returns a BroadcastOption that sets the gas adjustment configuration for the transaction.
//func WithGasAdjustment(val float64) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.gasAdjustment = new(float64)
//		*options.gasAdjustment = val
//		return nil
//	}
//}
//
//// WithNote returns a BroadcastOption that sets the note configuration for the transaction.
//func WithNote(val string) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.note = new(string)
//		*options.note = val
//		return nil
//	}
//}
//
//// WithGas returns a BroadcastOption that sets the gas setting configuration for the transaction.
//func WithGas(val cltypes.GasSetting) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.gas = new(cltypes.GasSetting)
//		*options.gas = val
//		return nil
//	}
//}
//
//// WithGasPrices returns a BroadcastOption that sets the gas price configuration for the transaction.
//// Gas price is a string of the amount. E.g. "0.25uakt".
//func WithGasPrices(val string) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.gasPrices = new(string)
//		*options.gasPrices = val
//		return nil
//	}
//}
//
//// WithFees returns a BroadcastOption that sets the fees configuration for the transaction.
//func WithFees(val string) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.fees = new(string)
//		*options.fees = val
//		return nil
//	}
//}
//
//// WithTimeoutHeight returns a BroadcastOption that sets the timeout height configuration for the transaction.
//func WithTimeoutHeight(val uint64) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.timeoutHeight = new(uint64)
//		*options.timeoutHeight = val
//		return nil
//	}
//}
//
//// WithBroadcastTimeout returns a BroadcastOption that sets the timeout configuration for the transaction.
//func WithBroadcastTimeout(val time.Duration) BroadcastOption {
//	return func(options *BroadcastOptions) error {
//		options.broadcastTimeout = val
//		return nil
//	}
//}
//
//// WithResultCodeAsError returns a BroadcastOption that enables the result code as error configuration for the transaction.
//func WithResultCodeAsError() BroadcastOption {
//	return func(opts *BroadcastOptions) error {
//		opts.resultAsError = true
//		return nil
//	}
//}
//
//// WithSkipConfirm returns a BroadcastOption that sets whether to skip or not the confirmation for the transaction.
//func WithSkipConfirm(val bool) BroadcastOption {
//	return func(opts *BroadcastOptions) error {
//		opts.skipConfirm = new(bool)
//		*opts.skipConfirm = val
//		return nil
//	}
//}
//
//// WithConfirmFn returns a BroadcastOption that sets the ConfirmFn function configuration for the transaction.
//func WithConfirmFn(val ConfirmFn) BroadcastOption {
//	return func(opts *BroadcastOptions) error {
//		opts.confirmFn = val
//		return nil
//	}
//}
//
//// WithBroadcastMode returns a BroadcastOption that sets the broadcast for particular tx
//func WithBroadcastMode(val string) BroadcastOption {
//	return func(opts *BroadcastOptions) error {
//		opts.broadcastMode = new(string)
//		*opts.broadcastMode = val
//		return nil
//	}
//}
//
//// WithGenerateOnly returns a BroadcastOption that sets transaction generator to generate only mode
//func WithGenerateOnly(val bool) BroadcastOption {
//	return func(opts *BroadcastOptions) error {
//		opts.generateOnly = new(bool)
//		*opts.generateOnly = val
//
//		return nil
//	}
//}
