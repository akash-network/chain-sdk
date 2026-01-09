package v2beta1

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	cerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	atypes "pkg.akt.dev/go/node/audit/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1"
	attr "pkg.akt.dev/go/node/types/attributes/v1"
)

// MakeOrderID returns OrderID instance with provided groupID details and oseq
func MakeOrderID(id dtypes.GroupID, oseq uint32) OrderID {
	return OrderID{
		Owner: id.Owner,
		DSeq:  id.DSeq,
		GSeq:  id.GSeq,
		OSeq:  oseq,
	}
}

// GroupID method returns groupID details for specific order
func (id OrderID) GroupID() dtypes.GroupID {
	return dtypes.GroupID{
		Owner: id.Owner,
		DSeq:  id.DSeq,
		GSeq:  id.GSeq,
	}
}

// Equals method compares specific order with provided order
func (id OrderID) Equals(other OrderID) bool {
	return id.GroupID().Equals(other.GroupID()) && id.OSeq == other.OSeq
}

// Validate method for OrderID and returns nil
func (id OrderID) Validate() error {
	if err := id.GroupID().Validate(); err != nil {
		return cerrors.Wrap(err, "OrderID: Invalid GroupID")
	}
	if id.OSeq == 0 {
		return sdkerrors.ErrInvalidSequence.Wrap("OrderID: Invalid Order Sequence")
	}
	return nil
}

// String provides stringer interface to save reflected formatting.
func (id OrderID) String() string {
	return fmt.Sprintf("%s/%v", id.GroupID(), id.OSeq)
}

// String implements the Stringer interface for a Order object.
func (o *Order) String() string {
	out, _ := yaml.Marshal(o)
	return string(out)
}

// Orders is a collection of Order
type Orders []Order

// String implements the Stringer interface for a Orders object.
func (o Orders) String() string {
	var out string
	for _, order := range o {
		out += order.String() + "\n"
	}

	return strings.TrimSpace(out)
}

// ValidateCanBid method validates whether order is open or not and
// returns error if not
func (o *Order) ValidateCanBid() error {
	switch o.State {
	case OrderOpen:
		return nil
	case OrderActive:
		return ErrOrderActive
	default:
		return ErrOrderClosed
	}
}

// ValidateInactive method validates whether order is open or not and
// returns error if not
func (o *Order) ValidateInactive() error {
	switch o.State {
	case OrderClosed:
		return nil
	case OrderActive:
		return ErrOrderActive
	default:
		return ErrOrderClosed
	}
}

// Price method returns price of specific order
func (o *Order) Prices() sdk.DecCoins {
	return o.Spec.Prices()
}

// MatchAttributes method compares provided attributes with specific order attributes
func (o *Order) MatchAttributes(attrs attr.Attributes) bool {
	return o.Spec.MatchAttributes(attrs)
}

// MatchRequirements method compares provided attributes with specific order attributes
func (o *Order) MatchRequirements(prov []atypes.AuditedProvider) bool {
	return o.Spec.MatchRequirements(prov)
}

// MatchResourcesRequirements method compares provider capabilities with specific order resources attributes
func (o *Order) MatchResourcesRequirements(attr attr.Attributes) bool {
	return o.Spec.MatchResourcesRequirements(attr)
}

// Filters returns whether order filters valid or not
func (o *Order) Filters(filters OrderFilters, stateVal Order_State) bool {
	// Checking owner filter
	if filters.Owner != "" && filters.Owner != o.ID.Owner {
		return false
	}

	// Checking dseq filter
	if filters.DSeq != 0 && filters.DSeq != o.ID.DSeq {
		return false
	}

	// Checking gseq filter
	if filters.GSeq != 0 && filters.GSeq != o.ID.GSeq {
		return false
	}

	// Checking oseq filter
	if filters.OSeq != 0 && filters.OSeq != o.ID.OSeq {
		return false
	}

	// Checking state filter
	if stateVal != 0 && stateVal != o.State {
		return false
	}

	return true
}
