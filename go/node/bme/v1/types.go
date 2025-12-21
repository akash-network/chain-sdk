package v1

// ToEvent this function exists bc protobuf does not support typedefs
func (m *BMRecord) ToEvent() *EventBMRecord {
	return &EventBMRecord{
		BurnedFrom: m.BurnedFrom,
		MintedTo:   m.MintedTo,
		Burner:     m.Burner,
		Minter:     m.Minter,
		Burned:     m.Burned,
		Minted:     m.Minted,
	}
}
