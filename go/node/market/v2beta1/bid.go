package v2beta1

import (
	"fmt"
	"sort"
	"strings"

	cerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gopkg.in/yaml.v3"

	dtypesv1 "pkg.akt.dev/go/node/deployment/v1"
	dtypes "pkg.akt.dev/go/node/deployment/v1beta5"
)

// MakeBidID returns BidID instance with provided order details and provider
func MakeBidID(id OrderID, provider sdk.AccAddress) BidID {
	return BidID{
		Owner:    id.Owner,
		DSeq:     id.DSeq,
		GSeq:     id.GSeq,
		OSeq:     id.OSeq,
		Provider: provider.String(),
	}
}

// Equals method compares specific bid with provided bid
func (id BidID) Equals(other BidID) bool {
	return id.OrderID().Equals(other.OrderID()) &&
		id.Provider == other.Provider
}

// LeaseID method returns lease details of bid
func (id BidID) LeaseID() LeaseID {
	return LeaseID(id)
}

// OrderID method returns OrderID details with specific bid details
func (id BidID) OrderID() OrderID {
	return OrderID{
		Owner: id.Owner,
		DSeq:  id.DSeq,
		GSeq:  id.GSeq,
		OSeq:  id.OSeq,
	}
}

// String method for consistent output.
func (id BidID) String() string {
	return fmt.Sprintf("%s/%s", id.OrderID().String(), id.Provider)
}

// GroupID method returns GroupID details with specific bid details
func (id BidID) GroupID() dtypesv1.GroupID {
	return id.OrderID().GroupID()
}

// DeploymentID method returns deployment details with specific bid details
func (id BidID) DeploymentID() dtypesv1.DeploymentID {
	return id.GroupID().DeploymentID()
}

// Validate validates bid instance and returns nil
func (id BidID) Validate() error {
	if err := id.OrderID().Validate(); err != nil {
		return cerrors.Wrap(err, "BidID: Invalid OrderID")
	}
	if _, err := sdk.AccAddressFromBech32(id.Provider); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("BidID: Invalid Provider Address")
	}
	if id.Owner == id.Provider {
		return sdkerrors.ErrConflict.Wrap("BidID: self-bid")
	}

	// BSeq is hardcoded to 0 for now. This will be lifted
	// once multiple bids support enable
	if id.BSeq != 0 {
		return sdkerrors.ErrConflict.Wrap("BidID: invalid bseq")
	}

	return nil
}

type ResourcesOffer []ResourceOffer

// Bids is a collection of Bid
type Bids []Bid

var _ sort.Interface = (*ResourcesOffer)(nil)

// String implements the Stringer interface for a Bid object.
func (o *Bid) String() string {
	out, _ := yaml.Marshal(o)
	return string(out)
}

// String implements the Stringer interface for a Bids object.
func (b Bids) String() string {
	var out string
	for _, bid := range b {
		out += bid.String() + "\n"
	}

	return strings.TrimSpace(out)
}

// Filters returns whether bid filters valid or not
func (o *Bid) Filters(filters BidFilters, stateVal Bid_State) bool {
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

	// Checking provider filter
	if filters.Provider != "" && filters.Provider != o.ID.Provider {
		return false
	}

	if filters.BSeq != 0 && filters.BSeq != o.ID.BSeq {
		return false
	}

	// Checking state filter
	if stateVal != 0 && stateVal != o.State {
		return false
	}

	return true
}

func (s ResourcesOffer) MatchGSpec(gspec dtypes.GroupSpec) bool {
	if len(s) == 0 {
		return true
	}

	ru := make(map[uint32]*dtypes.ResourceUnit)

	for idx := range gspec.Resources {
		ru[gspec.Resources[idx].ID] = &gspec.Resources[idx]
	}

	for _, ro := range s {
		res, exists := ru[ro.Resources.ID]
		if !exists {
			return false
		}

		ru[ro.Resources.ID] = nil

		if res.Count != ro.Count {
			return false
		}

		// TODO @troian check resources boundaries
	}

	return true
}

func (r *ResourceOffer) Dup() ResourceOffer {
	return ResourceOffer{
		Resources: r.Resources.Dup(),
		Count:     r.Count,
	}
}

func (s ResourcesOffer) Len() int {
	return len(s)
}

func (s ResourcesOffer) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ResourcesOffer) Less(i, j int) bool {
	return s[i].Resources.ID < s[j].Resources.ID
}

func (s ResourcesOffer) Dup() ResourcesOffer {
	res := make(ResourcesOffer, 0, len(s))

	for _, ru := range s {
		res = append(res, ru.Dup())
	}

	return res
}

func ResourceOfferFromRU(ru dtypes.ResourceUnits) ResourcesOffer {
	res := make(ResourcesOffer, 0, len(ru))

	for _, r := range ru {
		res = append(res, ResourceOffer{
			Resources: r.Resources,
			Count:     r.Count,
		})
	}

	return res
}
