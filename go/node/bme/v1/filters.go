package v1

// Accept returns whether the status of the record matches the filter criteria
func (filters *LedgerRecordFilters) Accept(id LedgerRecordID, status LedgerRecordStatus) bool {
	if filters.Status != "" {
		val, ok := LedgerRecordStatus_value[filters.Status]
		if !ok {
			return false
		}
		if LedgerRecordStatus(val) != status {
			return false
		}
	}

	if filters.Denom != "" && filters.Denom != id.Denom {
		return false
	}

	if filters.ToDenom != "" && filters.ToDenom != id.ToDenom {
		return false
	}

	if filters.Source != "" && filters.Source != id.Source {
		return false
	}

	return true
}
