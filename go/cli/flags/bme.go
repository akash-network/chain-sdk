package flags

import (
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "pkg.akt.dev/go/node/bme/v1"
)

// AddBMELedgerFilterFlags add flags to filter for ledger record list
func AddBMELedgerFilterFlags(flags *pflag.FlagSet) {
	flags.String(FlagOwner, "", "source address to filter")
	flags.String(FlagDenom, "", "burn denomination to filter")
	flags.String(FlagToDenom, "", "mint denomination to filter")
	flags.String(FlagStatus, "", "record status to filter (ledger_record_status_pending,ledger_record_status_executed)")
}

// BMELedgerFiltersFromFlags returns LedgerRecordFilters with given flags and error if occurred
func BMELedgerFiltersFromFlags(flags *pflag.FlagSet) (types.LedgerRecordFilters, error) {
	var filters types.LedgerRecordFilters

	owner, err := flags.GetString(FlagOwner)
	if err != nil {
		return filters, err
	}

	if owner != "" {
		_, err = sdk.AccAddressFromBech32(owner)
		if err != nil {
			return filters, err
		}
	}

	filters.Source = owner

	if filters.Denom, err = flags.GetString(FlagDenom); err != nil {
		return filters, err
	}

	if filters.ToDenom, err = flags.GetString(FlagToDenom); err != nil {
		return filters, err
	}

	if filters.Status, err = flags.GetString(FlagStatus); err != nil {
		return filters, err
	}

	return filters, nil
}
