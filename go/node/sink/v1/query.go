package v1

import "github.com/cosmos/cosmos-sdk/types/query"

const (
	DefaultPageLimit uint64 = 100
	MaxPageLimit     uint64 = 500
)

// ValidatePagination rejects offset scans; the selected index validates cursor bytes.
func ValidatePagination(p *query.PageRequest) error {
	if p != nil && (p.Offset != 0 || p.CountTotal || p.Limit > MaxPageLimit) {
		return ErrInvalidPagination.Wrapf("use key, reverse and limit up to %d; offset and count_total are unsupported", MaxPageLimit)
	}

	return nil
}

func (q QueryTransferRequest) ValidateBasic() error {
	return q.ID.ValidateBasic()
}

func (q QueryTransfersRequest) ValidateBasic() error {
	if q.Owner != "" {
		if _, err := ParseAccount(q.Owner, AccountLength); err != nil {
			return err
		}
	}

	if q.Status != StatusInvalid && !q.Status.IsValid() {
		return ErrInvalidSettlement.Wrap("invalid status filter")
	}

	return ValidatePagination(q.Pagination)
}

func (q QueryDisputesRequest) ValidateBasic() error {
	return ValidatePagination(q.Pagination)
}

func (q QueryAccountRequest) ValidateBasic() error {
	_, err := ParseAccount(q.Owner, AccountLength)
	return err
}
