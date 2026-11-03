package v1

import (
	"strconv"
	"strings"
)

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

	id := ID{
		Owner:    parts[0],
		Height:   height,
		Sequence: sequence,
	}
	if err = id.ValidateBasic(); err != nil {
		return ID{}, err
	}

	if id.String() != s {
		return ID{}, ErrInvalidID.Wrap("noncanonical ID")
	}

	return id, nil
}
