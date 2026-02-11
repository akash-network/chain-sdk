package v1

// AcceptPending returns whether the pending record matches the filter criteria
func (filters *LedgerRecordFilters) AcceptPending(id LedgerRecordID, _ LedgerPendingRecord) bool {
	if filters.Status != "" {
		val, ok := LedgerRecordStatus_value[filters.Status]
		if !ok {
			return false
		}
		if LedgerRecordStatus(val) != LedgerRecordSatusPending {
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

// AcceptExecuted returns whether the executed record matches the filter criteria
func (filters *LedgerRecordFilters) AcceptExecuted(id LedgerRecordID, _ LedgerRecord) bool {
	if filters.Status != "" {
		val, ok := LedgerRecordStatus_value[filters.Status]
		if !ok {
			return false
		}
		if LedgerRecordStatus(val) != LedgerRecordSatusExecuted {
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
