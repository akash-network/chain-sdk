package v1

import (
	"encoding/binary"
	"encoding/json"
	"math"
	"strconv"
	"strings"

	collcodec "cosmossdk.io/collections/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const IDKeySize = 37

var IDKey collcodec.KeyCodec[ID] = idKey{}

type idKey struct{}

func (id ID) ValidateBasic() error {
	if _, err := ParseAccount(id.Owner, AccountLength); err != nil {
		return ErrInvalidID.Wrap(err.Error())
	}

	if id.Height <= 0 || id.Sequence == 0 {
		return ErrInvalidID.Wrap("height and sequence must be positive")
	}

	return nil
}

func (id ID) String() string {
	return id.Owner + "/" + strconv.FormatInt(id.Height, 10) + "/" + strconv.FormatUint(id.Sequence, 10)
}

func ParseID(s string) (ID, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 3 {
		return ID{}, ErrInvalidID.Wrap("expected owner/height/sequence")
	}

	height, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return ID{}, ErrInvalidID.Wrap("invalid height")
	}

	sequence, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return ID{}, ErrInvalidID.Wrap("invalid sequence")
	}

	id := ID{Owner: parts[0], Height: height, Sequence: sequence}
	if err = id.ValidateBasic(); err != nil {
		return ID{}, err
	}

	if id.String() != s {
		return ID{}, ErrInvalidID.Wrap("noncanonical ID")
	}

	return id, nil
}

func OwnerRange(owner sdk.AccAddress) (first, last ID) {
	return ID{Owner: owner.String(), Height: 1, Sequence: 1}, ID{Owner: owner.String(), Height: math.MaxInt64, Sequence: math.MaxUint64}
}

func (idKey) Encode(b []byte, id ID) (int, error) {
	if len(b) < IDKeySize {
		return 0, ErrInvalidID.Wrap("short key buffer")
	}

	if err := id.ValidateBasic(); err != nil {
		return 0, err
	}

	owner, err := ParseAccount(id.Owner, AccountLength)
	if err != nil {
		return 0, err
	}

	b[0] = AccountLength
	copy(b[1:21], owner)
	binary.BigEndian.PutUint64(b[21:29], uint64(id.Height)) //nolint:gosec // ValidateBasic requires a positive height.
	binary.BigEndian.PutUint64(b[29:37], id.Sequence)

	return IDKeySize, nil
}

func (idKey) Decode(b []byte) (int, ID, error) {
	if len(b) < IDKeySize || b[0] != AccountLength {
		return 0, ID{}, ErrInvalidID.Wrap("invalid key length or owner length")
	}

	height := binary.BigEndian.Uint64(b[21:29])
	if height == 0 || height > math.MaxInt64 {
		return 0, ID{}, ErrInvalidID.Wrap("height out of range")
	}

	id := ID{Owner: sdk.AccAddress(b[1:21]).String(), Height: int64(height), Sequence: binary.BigEndian.Uint64(b[29:37])}
	if err := id.ValidateBasic(); err != nil {
		return 0, ID{}, err
	}

	return IDKeySize, id, nil
}

func (idKey) Size(ID) int { return IDKeySize }

func (k idKey) EncodeNonTerminal(b []byte, id ID) (int, error) {
	return k.Encode(b, id)
}

func (k idKey) DecodeNonTerminal(b []byte) (int, ID, error) {
	return k.Decode(b)
}

func (idKey) SizeNonTerminal(ID) int { return IDKeySize }

func (idKey) KeyType() string { return "sink.ID" }

func (idKey) Stringify(id ID) string { return id.String() }

func (idKey) EncodeJSON(id ID) ([]byte, error) {
	if err := id.ValidateBasic(); err != nil {
		return nil, err
	}

	return json.Marshal(id)
}

func (idKey) DecodeJSON(b []byte) (ID, error) {
	var id ID
	if err := json.Unmarshal(b, &id); err != nil {
		return ID{}, err
	}

	return id, id.ValidateBasic()
}
