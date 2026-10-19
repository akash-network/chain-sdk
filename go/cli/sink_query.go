package cli

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/sink/v1"
)

// GetQuerySinkCmd returns the query commands for the sink module
func GetQuerySinkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Sink query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetQuerySinkRecordCmd(),
		GetQuerySinkOpenTransfersCmd(),
	)

	return cmd
}

// GetQuerySinkRecordCmd returns a transfer's status by its record ID
func GetQuerySinkRecordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "record [id]",
		Short:             "Query status and settlement by canonical record ID",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			res, err := cl.Query().Sink().Sink(ctx, &types.QuerySink{Id: args[0]})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkOpenTransfersCmd pages through transfers that are neither minted
// nor refunded, in ascending ID order
func GetQuerySinkOpenTransfersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-transfers",
		Short: "List pending and voided transfers in ascending ID order",
		Long: `List transfers with IDs above --start-after.
Pass the returned next_start_after as --start-after to continue; zero means there are no more.`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			startAfter, err := cmd.Flags().GetUint64("start-after")
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetUint32("limit")
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			res, err := cl.Query().Sink().OpenTransfers(ctx, &types.QueryOpenTransfersRequest{
				StartAfter: startAfter,
				Limit:      limit,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cmd.Flags().Uint64("start-after", 0, "Return transfers with IDs above this one")
	cmd.Flags().Uint32("limit", 0, "Maximum transfers to return; 0 means 100, at most 500")
	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}
