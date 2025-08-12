package cli

import (
	"context"

	sdksrv "github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"

	cflags "pkg.akt.dev/go/cli/flags"
	aclient "pkg.akt.dev/go/node/client/discovery"
)

func QueryPersistentPreRunE(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()

	cctx, err := GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	cl, err := aclient.DiscoverLightClient(ctx, cctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, ContextTypeQueryClient, cl)

	cmd.SetContext(ctx)

	return nil
}

func QueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	cmd.AddCommand(
		GetQueryAuthCmd(),
		GetQueryAuthzCmd(),
		GetQueryBankCmd(),
		GetQueryDistributionCmd(),
		GetQueryEvidenceCmd(),
		GetQueryFeegrantCmd(),
		GetQueryMintCmd(),
		GetQueryParamsCmd(),
		cflags.LineBreak,
		sdksrv.QueryBlockCmd(),
		sdksrv.QueryBlocksCmd(),
		GetQueryAuthTxsByEventsCmd(),
		GetQueryAuthTxCmd(),
		GetQueryGovCmd(),
		GetQuerySlashingCmd(),
		GetQueryStakingCmd(),
		cflags.LineBreak,
		GetQueryAuditCmd(),
		GetQueryCertCmd(),
		GetQueryDeploymentCmds(),
		GetQueryMarketCmds(),
		GetQueryEscrowCmd(),
		GetQueryProviderCmds(),
	)

	cmd.PersistentFlags().String(cflags.FlagChainID, "", "The network chain ID")

	return cmd
}
