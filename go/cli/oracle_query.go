package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/oracle/v2"
)

func GetQueryOracleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Oracle query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetOraclePricesCmd(),
		GetOracleAggregatedPriceCmd(),
		GetQueryOracleParamsCmd(),
	)

	return cmd
}

func GetOraclePricesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "prices",
		Aliases:           []string{"p"},
		Short:             "Query price history for denoms",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			// Get filter flags
			assetDenom, _ := cmd.Flags().GetString(cflags.FlagAssetDenom)
			baseDenom, _ := cmd.Flags().GetString(cflags.FlagBaseDenom)
			startTimeStr, _ := cmd.Flags().GetString("start-time")
			endTimeStr, _ := cmd.Flags().GetString("end-time")

			filters := types.PricesFilter{
				AssetDenom: assetDenom,
				BaseDenom:  baseDenom,
			}

			if startTimeStr != "" {
				ts, err := time.Parse(time.RFC3339, startTimeStr)
				if err != nil {
					return err
				}
				filters.StartTime = ts
			}

			if endTimeStr != "" {
				ts, err := time.Parse(time.RFC3339, endTimeStr)
				if err != nil {
					return err
				}
				filters.EndTime = ts
			}

			if !filters.StartTime.IsZero() && !filters.EndTime.IsZero() && filters.StartTime.After(filters.EndTime) {
				return fmt.Errorf("start-time %q must be before end-time %q", startTimeStr, endTimeStr)
			}

			req := &types.QueryPricesRequest{
				Filters:    filters,
				Pagination: pageReq,
			}

			res, err := cl.Query().Oracle().Prices(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "prices")
	cmd.Flags().String(cflags.FlagAssetDenom, "", "Filter by asset denomination (e.g., uakt)")
	cmd.Flags().String(cflags.FlagBaseDenom, "", "Filter by base denomination (e.g., usd)")
	cmd.Flags().String("start-time", "", "Filter by start time (RFC3339 format, e.g., 2024-01-01T00:00:00Z)")
	cmd.Flags().String("end-time", "", "Filter by end time (RFC3339 format, e.g., 2024-01-01T00:00:00Z)")

	return cmd
}

func GetOracleAggregatedPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "aggregated-price [denom]",
		Aliases:           []string{"ap"},
		Short:             "Query aggregated price for a denom",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryAggregatedPriceRequest{
				Denom: args[0],
			}

			res, err := cl.Query().Oracle().AggregatedPrice(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetQueryOracleParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "params",
		Short:             "Query the current oracle parameters",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryParamsRequest{}

			res, err := cl.Query().Oracle().Params(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}
