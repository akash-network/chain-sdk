package v1

import (
	"encoding/binary"
	"encoding/json"

	collcodec "cosmossdk.io/collections/codec"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"pkg.akt.dev/go/sdkutil"
)

const (
	WireVersion    byte = 2
	RecordDataSize int  = 41
)

var (
	TransferKeyPrefix                                   = []byte{0x12, 0x00}
	VoidOrderKeyPrefix                                  = []byte{0x12, 0x01}
	RecordDataValue    collcodec.ValueCodec[RecordData] = recordDataValue{}
	VoidOrderValue     collcodec.ValueCodec[VoidOrder]  = voidOrderValue{}
)

type VoidOrder struct{}

type recordDataValue struct{}

type voidOrderValue struct{}

func TransferKey(id ID) ([]byte, error) {
	return proofKey(TransferKeyPrefix, id)
}

func VoidOrderKey(id ID) ([]byte, error) {
	return proofKey(VoidOrderKeyPrefix, id)
}

func proofKey(prefix []byte, id ID) ([]byte, error) {
	b := make([]byte, len(prefix)+IDKeySize)
	copy(b, prefix)

	if _, err := IDKey.Encode(b[len(prefix):], id); err != nil {
		return nil, err
	}

	return b, nil
}

func (r RecordData) ValidateBasic() error {
	if _, err := ParseSolanaAddress(r.Dest); err != nil {
		return err
	}

	if err := ValidateAmount(r.Amount); err != nil {
		return err
	}

	if !r.Amount.Amount.IsPositive() || !r.Amount.Amount.IsUint64() {
		return sdkerrors.ErrInvalidCoins.Wrap("transfer amount must be a positive uint64")
	}

	return nil
}

func (recordDataValue) Encode(r RecordData) ([]byte, error) {
	if err := r.ValidateBasic(); err != nil {
		return nil, err
	}

	recipient, err := ParseSolanaAddress(r.Dest)
	if err != nil {
		return nil, err
	}

	b := make([]byte, RecordDataSize)
	b[0] = WireVersion
	copy(b[1:33], recipient[:])
	binary.LittleEndian.PutUint64(b[33:], r.Amount.Amount.Uint64())

	return b, nil
}

func (recordDataValue) Decode(b []byte) (RecordData, error) {
	if len(b) != RecordDataSize || b[0] != WireVersion {
		return RecordData{}, ErrInvalidRecord.Wrapf("expected version %d and %d bytes", WireVersion, RecordDataSize)
	}

	var recipient SolanaAddress
	copy(recipient[:], b[1:33])
	r := RecordData{Dest: recipient.String(), Amount: sdk.NewCoin(sdkutil.DenomUakt, sdkmath.NewIntFromUint64(binary.LittleEndian.Uint64(b[33:])))}
	return r, r.ValidateBasic()
}

func (recordDataValue) EncodeJSON(r RecordData) ([]byte, error) {
	if err := r.ValidateBasic(); err != nil {
		return nil, err
	}

	return json.Marshal(r)
}

func (recordDataValue) DecodeJSON(b []byte) (RecordData, error) {
	var r RecordData
	if err := json.Unmarshal(b, &r); err != nil {
		return r, err
	}

	return r, r.ValidateBasic()
}

func (recordDataValue) Stringify(r RecordData) string { return r.String() }

func (recordDataValue) ValueType() string { return "sink.RecordData" }

func (voidOrderValue) Encode(VoidOrder) ([]byte, error) {
	return []byte{WireVersion}, nil
}

func (voidOrderValue) Decode(b []byte) (VoidOrder, error) {
	if len(b) != 1 || b[0] != WireVersion {
		return VoidOrder{}, ErrInvalidRecord.Wrapf("expected version %d void order", WireVersion)
	}

	return VoidOrder{}, nil
}

func (voidOrderValue) EncodeJSON(v VoidOrder) ([]byte, error) {
	return json.Marshal(v)
}

func (voidOrderValue) DecodeJSON(b []byte) (VoidOrder, error) {
	var v VoidOrder
	if string(b) != "{}" {
		return v, ErrInvalidRecord.Wrap("expected empty void order object")
	}

	return v, nil
}

func (voidOrderValue) Stringify(VoidOrder) string { return "{}" }

func (voidOrderValue) ValueType() string { return "sink.VoidOrder" }
