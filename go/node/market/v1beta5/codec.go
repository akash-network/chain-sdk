package v1beta5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	v1 "pkg.akt.dev/go/node/market/v1"
)

var (
	// amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/market module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/market and
	// defined at the application level.
	//
	// Deprecated: ModuleCdc use is deprecated
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// RegisterLegacyAminoCodec registers the necessary x/market interfaces and concrete types
// on the provided Amino codec. These types are used for Amino JSON serialization.
//
// Deprecated: RegisterLegacyAminoCodec is deprecated
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateBid{}, "akash-sdk/x/"+v1.ModuleName+"/"+(&MsgCreateBid{}).Type(), nil)
	cdc.RegisterConcrete(&MsgCloseBid{}, "akash-sdk/x/"+v1.ModuleName+"/"+(&MsgCloseBid{}).Type(), nil)
	cdc.RegisterConcrete(&MsgCreateLease{}, "akash-sdk/x/"+v1.ModuleName+"/"+(&MsgCreateLease{}).Type(), nil)
	cdc.RegisterConcrete(&MsgCloseLease{}, "akash-sdk/x/"+v1.ModuleName+"/"+(&MsgCloseLease{}).Type(), nil)
	cdc.RegisterConcrete(&MsgWithdrawLease{}, "akash-sdk/x/"+v1.ModuleName+"/"+(&MsgWithdrawLease{}).Type(), nil)
}

// RegisterInterfaces registers the x/market interfaces types with the interface registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBid{},
		&MsgCloseBid{},
		&MsgCreateLease{},
		&MsgCloseLease{},
		&MsgWithdrawLease{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
