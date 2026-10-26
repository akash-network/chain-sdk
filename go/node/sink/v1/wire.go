package v1

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

const (
	ModuleName            = "sink"
	StoreKey              = ModuleName
	Denom                 = "uakt"
	DestTypeSolana uint32 = 1
	RecordSize            = 77
)

// Transfer is the cross-chain commitment. Never rewrite or delete its value.
type Transfer struct {
	ID        uint64
	Sender    [20]byte
	Recipient [32]byte
	Amount    uint64
	Height    uint64
}

func TransferKey(id uint64) []byte { return []byte("transfer/" + strconv.FormatUint(id, 10)) }
func VoidKey(id uint64) []byte     { return []byte("void/" + strconv.FormatUint(id, 10)) }

func (t Transfer) Marshal() ([]byte, error) {
	if t.ID == 0 || t.Amount == 0 || t.Height == 0 || t.Recipient == [32]byte{} {
		return nil, fmt.Errorf("invalid sink transfer")
	}
	b := make([]byte, RecordSize)
	b[0] = 1
	binary.LittleEndian.PutUint64(b[1:9], t.ID)
	copy(b[9:29], t.Sender[:])
	copy(b[29:61], t.Recipient[:])
	binary.LittleEndian.PutUint64(b[61:69], t.Amount)
	binary.LittleEndian.PutUint64(b[69:77], t.Height)
	return b, nil
}

func DecodeTransfer(b []byte) (Transfer, error) {
	var t Transfer
	if len(b) != RecordSize || b[0] != 1 {
		return t, fmt.Errorf("expected sink transfer version 1, %d bytes", RecordSize)
	}
	t.ID = binary.LittleEndian.Uint64(b[1:9])
	copy(t.Sender[:], b[9:29])
	copy(t.Recipient[:], b[29:61])
	t.Amount = binary.LittleEndian.Uint64(b[61:69])
	t.Height = binary.LittleEndian.Uint64(b[69:77])
	_, err := t.Marshal()
	return t, err
}

func VoidValue(id uint64) []byte {
	b := make([]byte, 9)
	b[0] = 1
	binary.LittleEndian.PutUint64(b[1:], id)
	return b
}
