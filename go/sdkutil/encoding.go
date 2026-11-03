package sdkutil

import (
	"cosmossdk.io/x/tx/signing"
	"cosmossdk.io/x/tx/signing/aminojson"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"

	offchain "pkg.akt.dev/go/node/types/offchain/sign"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	SigningOptions    signing.Options
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for a proto based test configuration.
func MakeEncodingConfig(modules ...module.AppModuleBasic) EncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	co := NewCodecOptions()

	interfaceRegistry := co.NewInterfaceRegistry()

	std.RegisterLegacyAminoCodec(aminoCodec)
	std.RegisterInterfaces(interfaceRegistry)

	aminoCodec.RegisterConcrete(&offchain.MsgSignData{}, "sign/"+(&offchain.MsgSignData{}).Type(), nil)

	if len(modules) > 0 {
		mb := module.NewBasicManager(modules...)
		mb.RegisterLegacyAminoCodec(aminoCodec)
		mb.RegisterInterfaces(interfaceRegistry)
	}

	cdc := codec.NewProtoCodec(interfaceRegistry)

	signingCtx, err := signing.NewContext(co.Options)
	if err != nil {
		panic(err)
	}

	// The global registry leaves nested gogoproto messages unresolved, and Amino JSON
	// would sign them as empty objects. Signer lookups must stay on the global
	// descriptors that every tx config shares, so only this handler resolves them.
	files, err := proto.MergedRegistry()
	if err != nil {
		panic(err)
	}

	txConfig, err := tx.NewTxConfigWithOptions(cdc, tx.ConfigOptions{
		EnabledSignModes: []signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
			signingtypes.SignMode_SIGN_MODE_DIRECT_AUX,
		},
		CustomSignModes: []signing.SignModeHandler{aminojson.NewSignModeHandler(aminojson.SignModeHandlerOptions{
			FileResolver: files,
			TypeResolver: co.Options.TypeResolver,
		})},
		SigningOptions: &co.Options,
		SigningContext: signingCtx,
	})
	if err != nil {
		panic(err)
	}

	encCfg := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          txConfig,
		Amino:             aminoCodec,
		SigningOptions:    co.Options,
	}

	return encCfg
}
