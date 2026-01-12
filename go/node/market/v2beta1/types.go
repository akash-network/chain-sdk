package v2beta1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GatewayVersion = "v2beta2"
)

type LeaseClosedReasonRange int

const (
	LeaseClosedReasonRangeOwner LeaseClosedReasonRange = iota
	LeaseClosedReasonRangeProvider
	LeaseClosedReasonRangeNetwork
)

func (m LeaseClosedReason) IsRange(r LeaseClosedReasonRange) bool {
	switch r {
	case LeaseClosedReasonRangeOwner:
		return m >= 0 && m <= 9999
	case LeaseClosedReasonRangeProvider:
		return m >= 10000 && m <= 19999
	case LeaseClosedReasonRangeNetwork:
		return m >= 20000 && m <= 29999
	}

	return false
}

func (m QueryLeasesResponse) TotalPricesAmount() sdk.DecCoins {
	total := sdk.DecCoins{}

	for i, lease := range m.Leases {
		if i == 0 {
			total = lease.Lease.Prices
		} else {
			total = total.Add(lease.Lease.Prices...)
		}
	}

	return total
}
