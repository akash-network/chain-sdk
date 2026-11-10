package v1

func NewSettlementID(status Status, id ID) SettlementID {
	return SettlementID{
		Status:   status,
		Owner:    id.Owner,
		Height:   id.Height,
		Sequence: id.Sequence,
	}
}

func (key *SettlementID) GetID() *ID {
	return &ID{
		Owner:    key.Owner,
		Height:   key.Height,
		Sequence: key.Sequence,
	}
}

func (key *SettlementID) ValidateBasic() error {
	if !key.Status.IsValid() {
		return ErrInvalidSettlement.Wrap("unknown financial status")
	}

	return key.GetID().ValidateBasic()
}
