package v1

import (
	"encoding/json"

	collcodec "cosmossdk.io/collections/codec"
)

var SettlementIDKey collcodec.KeyCodec[SettlementID] = settlementIDCodec{}

type settlementIDCodec struct{}

const SettlementIDKeySize = 1 + IDKeySize

func NewSettlementID(status Status, id ID) SettlementID {
	return SettlementID{
		Status:   status,
		Owner:    id.Owner,
		Height:   id.Height,
		Sequence: id.Sequence,
	}
}

func (key SettlementID) ID() ID {
	return ID{
		Owner:    key.Owner,
		Height:   key.Height,
		Sequence: key.Sequence,
	}
}

func statusByte(status Status) (byte, error) {
	if !status.IsValid() {
		return 0, ErrInvalidSettlement.Wrap("unknown financial status")
	}

	return byte(status), nil //nolint:gosec // IsValid admits only statuses 1 through 5.
}

func (settlementIDCodec) Encode(b []byte, key SettlementID) (int, error) {
	if len(b) < SettlementIDKeySize {
		return 0, ErrInvalidID.Wrap("short settlement key buffer")
	}

	status, err := statusByte(key.Status)
	if err != nil {
		return 0, err
	}

	n, err := IDKey.Encode(b[1:], key.ID())
	if err != nil {
		return 0, err
	}

	b[0] = status

	return 1 + n, nil
}

func (settlementIDCodec) Decode(b []byte) (int, SettlementID, error) {
	if len(b) < SettlementIDKeySize {
		return 0, SettlementID{}, ErrInvalidID.Wrap("short settlement key")
	}

	status := Status(b[0])
	if _, err := statusByte(status); err != nil {
		return 0, SettlementID{}, err
	}

	n, id, err := IDKey.Decode(b[1:])
	if err != nil {
		return 0, SettlementID{}, err
	}

	return 1 + n, NewSettlementID(status, id), nil
}

func (settlementIDCodec) Size(SettlementID) int { return SettlementIDKeySize }

func (c settlementIDCodec) EncodeNonTerminal(b []byte, key SettlementID) (int, error) {
	return c.Encode(b, key)
}

func (c settlementIDCodec) DecodeNonTerminal(b []byte) (int, SettlementID, error) {
	return c.Decode(b)
}

func (settlementIDCodec) SizeNonTerminal(SettlementID) int { return SettlementIDKeySize }

func (settlementIDCodec) KeyType() string { return "sink.SettlementID" }

func (settlementIDCodec) Stringify(key SettlementID) string {
	return key.Status.String() + "/" + key.ID().String()
}

func (settlementIDCodec) EncodeJSON(key SettlementID) ([]byte, error) {
	if _, err := statusByte(key.Status); err != nil {
		return nil, err
	}

	if err := key.ID().ValidateBasic(); err != nil {
		return nil, err
	}

	return json.Marshal(key)
}

func (settlementIDCodec) DecodeJSON(b []byte) (SettlementID, error) {
	var key SettlementID
	if err := json.Unmarshal(b, &key); err != nil {
		return SettlementID{}, err
	}

	if _, err := statusByte(key.Status); err != nil {
		return SettlementID{}, err
	}

	return key, key.ID().ValidateBasic()
}
