package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	// ModuleCdc references the global x/verification module codec.
	//
	// Deprecated: ModuleCdc use is deprecated.
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPostAuditorBond{},
		&MsgSubmitAttestation{},
		&MsgOpenAuditEscrow{},
		&MsgCancelAuditEscrow{},
		&MsgSettleAuditEscrow{},
		&MsgRevokeAttestation{},
		&MsgRemoveAttestation{},
		&MsgResignAuditor{},
		&MsgPostProviderBond{},
		&MsgWithdrawProviderBond{},
		&MsgPostSnapshotHash{},
		&MsgRegisterAuditor{},
		&MsgRenewAuditor{},
		&MsgRemoveAuditor{},
		&MsgRevokeProviderAttestation{},
		&MsgRevokeAllProviderAttestations{},
		&MsgRevokeAuditorAttestations{},
		&MsgResolveDiscrepancy{},
		&MsgSlashProviderBond{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &Msg_serviceDesc)
}
