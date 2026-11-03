package cli

import (
	"strconv"

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
		GetQuerySinkTransferCmd(),
		GetQuerySinkTransfersCmd(),
		GetQuerySinkDisputesCmd(),
		GetQuerySinkAccountCmd(),
		GetQuerySinkStateCmd(),
		GetQuerySinkParamsCmd(),
		GetQuerySinkReportersCmd(),
	)

	return cmd
}

// GetQuerySinkTransferCmd returns a complete transfer by its composite ID
func GetQuerySinkTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "transfer [owner/height/sequence]",
		Short:             "Query a complete transfer by its composite ID",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.ParseID(args[0])
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Transfer(ctx, &types.QueryTransferRequest{ID: id})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkTransfersCmd pages through transfers by owner, financial status or history
func GetQuerySinkTransfersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "transfers",
		Short:             "List transfers by owner, financial status, or chronological history",
		Long:              "List transfers using key pagination. Keep --height fixed on every page. Without filters, history is ordered by height, sequence, then owner.",
		Args:              cobra.NoArgs,
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			pagination, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			owner, err := cmd.Flags().GetString("owner")
			if err != nil {
				return err
			}

			status, err := cmd.Flags().GetString("status")
			if err != nil {
				return err
			}

			value, ok := types.Status_value[status]
			if !ok {
				return types.ErrInvalidSettlement.Wrap("unknown status filter")
			}

			request := &types.QueryTransfersRequest{
				Owner:      owner,
				Status:     types.Status(value),
				Pagination: pagination,
			}
			if err = request.ValidateBasic(); err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Transfers(ctx, request)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cmd.Flags().String("owner", "", "Filter by a canonical Akash owner")
	cmd.Flags().String("status", "STATUS_INVALID", "STATUS_TRANSFER_PENDING, STATUS_MINTED, STATUS_REFUND_PENDING, STATUS_REFUNDED, or STATUS_INVALID for any")
	cmd.Flags().String(cflags.FlagPageKey, "", "pagination page-key of transfers to query for")
	cmd.Flags().Uint64(cflags.FlagLimit, 100, "pagination limit of transfers to query for")
	cmd.Flags().Bool(cflags.FlagReverse, false, "results are sorted in descending order")
	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkDisputesCmd pages through held transfers
func GetQuerySinkDisputesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "disputes",
		Short:             "List held transfers, including minted and refunded conflicts",
		Args:              cobra.NoArgs,
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			pagination, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			request := &types.QueryDisputesRequest{Pagination: pagination}
			if err = request.ValidateBasic(); err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Disputes(ctx, request)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cmd.Flags().String(cflags.FlagPageKey, "", "pagination page-key of disputes to query for")
	cmd.Flags().Uint64(cflags.FlagLimit, 100, "pagination limit of disputes to query for")
	cmd.Flags().Bool(cflags.FlagReverse, false, "results are sorted in descending order")
	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkAccountCmd returns an owner's lifetime and current epoch usage
func GetQuerySinkAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "account [owner]",
		Short:             "Query owner lifetime and current epoch sending usage",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			request := &types.QueryAccountRequest{Owner: args[0]}
			if err := request.ValidateBasic(); err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Account(ctx, request)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkStateCmd returns params, totals, the current policy, epoch usage and the refund head
func GetQuerySinkStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "state",
		Short:             "Query params, totals, policy, independent epoch usage and refund head",
		Args:              cobra.NoArgs,
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().State(ctx, &types.QueryStateRequest{})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkParamsCmd returns the command to query Sink module parameters
func GetQuerySinkParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "params",
		Short:             "Query the current Sink module parameters",
		Args:              cobra.NoArgs,
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Params(ctx, &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetQuerySinkReportersCmd returns the command to query reporter policies by version
func GetQuerySinkReportersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "reporters [version]",
		Short:             "Query Sink reporter policies by version",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)
			res, err := cl.Query().Sink().Reporters(ctx, &types.QueryReportersRequest{Version: version})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}
