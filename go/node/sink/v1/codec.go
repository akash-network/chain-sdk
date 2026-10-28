package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers sink messages using their original signing names.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSink{}, "sink/MsgSink", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "sink/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgRequestVoid{}, "sink/MsgRequestVoid", nil)
	cdc.RegisterConcrete(&MsgReportOutcome{}, "sink/MsgReportOutcome", nil)
	cdc.RegisterConcrete(&MsgVetoVoid{}, "sink/MsgVetoVoid", nil)
	cdc.RegisterConcrete(&MsgResolveTransfer{}, "sink/MsgResolveTransfer", nil)
	cdc.RegisterConcrete(&MsgFinalizeRefund{}, "sink/MsgFinalizeRefund", nil)
	cdc.RegisterConcrete(&MsgGuardianPause{}, "sink/MsgGuardianPause", nil)
}

// RegisterInterfaces registers the sink messages and their response descriptors.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSink{}, &MsgUpdateParams{}, &MsgRequestVoid{}, &MsgReportOutcome{},
		&MsgVetoVoid{}, &MsgResolveTransfer{}, &MsgFinalizeRefund{}, &MsgGuardianPause{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
