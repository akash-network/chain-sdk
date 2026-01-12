package v2beta1

import (
	"fmt"
)

// getOrdersPath returns orders path for queries
// nolint: unused
func getOrdersPath(ofilters OrderFilters) string {
	return fmt.Sprintf("%s/%s/%v", ordersPath, ofilters.Owner, ofilters.State)
}

// OrderPath return order path of given order id for queries
func OrderPath(id OrderID) string {
	return fmt.Sprintf("%s/%s", orderPath, orderParts(id))
}

// getBidsPath returns bids path for queries
// nolint: unused
func getBidsPath(bfilters BidFilters) string {
	return fmt.Sprintf("%s/%s/%v", bidsPath, bfilters.Owner, bfilters.State)
}

// getBidPath return bid path of given bid id for queries
// nolint: unused
func getBidPath(id BidID) string {
	return fmt.Sprintf("%s/%s/%s", bidPath, orderParts(id.OrderID()), id.Provider)
}

// getLeasesPath returns leases path for queries
// nolint: unused
func getLeasesPath(lfilters LeaseFilters) string {
	return fmt.Sprintf("%s/%s/%v", leasesPath, lfilters.Owner, lfilters.State)
}
