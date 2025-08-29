package v1

import (
	"errors"
	"fmt"
	"strings"

	ev1 "pkg.akt.dev/go/node/escrow/id/v1"
)

func (id BidID) ToEscrowAccountID() ev1.Account {
	return ev1.Account{
		Scope: ev1.ScopeBid,
		XID:   id.String(),
	}
}

func (id LeaseID) ToEscrowPaymentID() ev1.Payment {
	return ev1.Payment{
		AID: id.DeploymentID().ToEscrowAccountID(),
		XID: fmt.Sprintf("%d/%d/%s", id.GSeq, id.OSeq, id.Provider),
	}
}

func LeaseIDFromPaymentID(id ev1.Payment) (LeaseID, error) {
	if id.AID.Scope != ev1.ScopeDeployment {
		return LeaseID{}, errors.New("")
	}

	parts := strings.Split(strings.Join([]string{id.AID.XID, id.XID}, "/"), "/")
	if len(parts) != 3 {
		return LeaseID{}, errors.New("")
	}

	return ParseLeasePath(parts)
}
