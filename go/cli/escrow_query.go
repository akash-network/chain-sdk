package cli

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	dv1 "pkg.akt.dev/go/node/deployment/v1"
	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
	etypes "pkg.akt.dev/go/node/escrow/v1"
	mv1 "pkg.akt.dev/go/node/market/v1"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
	"pkg.akt.dev/go/node/utils"
)

var errNoLeaseMatches = errors.New("leases for deployment do not exist")

func GetQueryEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        etypes.ModuleName,
		Short:                      "Escrow query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetQueryEscrowBlocksRemainingCmd(),
	)

	return cmd
}

func GetQueryEscrowBlocksRemainingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "blocks-remaining",
		Short:             "Compute the number of blocks remaining for an ecrow account",
		Args:              cobra.ExactArgs(0),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustQueryClientFromContext(ctx)

			id, err := cflags.DeploymentIDFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			// Fetch leases matching owner & dseq
			leaseRequest := mv1beta5.QueryLeasesRequest{
				Filters: mv1.LeaseFilters{
					Owner:    id.Owner,
					DSeq:     id.DSeq,
					GSeq:     0,
					OSeq:     0,
					Provider: "",
					State:    mv1.LeaseActive.String(),
				},
				Pagination: nil,
			}

			leasesResponse, err := cl.Query().Market().Leases(ctx, &leaseRequest)
			if err != nil {
				return err
			}

			if len(leasesResponse.Leases) == 0 {
				return errNoLeaseMatches
			}

			// Fetch the balance of the escrow account
			totalLeaseAmount := leasesResponse.TotalPriceAmount()
			blockchainHeight, err := cl.Node().CurrentBlockHeight(ctx)
			if err != nil {
				return err
			}

			res, err := cl.Query().Deployment().Deployment(cmd.Context(), &dv1beta4.QueryDeploymentRequest{
				ID: dv1.DeploymentID{Owner: id.Owner, DSeq: id.DSeq},
			})
			if err != nil {
				return err
			}

			balanceRemain := utils.LeaseCalcBalanceRemain(res.EscrowAccount.TotalBalance().Amount,
				int64(blockchainHeight),
				res.EscrowAccount.SettledAt,
				totalLeaseAmount)

			blocksRemain := utils.LeaseCalcBlocksRemain(balanceRemain, totalLeaseAmount)

			output := struct {
				BalanceRemain       float64       `json:"balance_remaining" yaml:"balance_remaining"`
				BlocksRemain        int64         `json:"blocks_remaining" yaml:"blocks_remaining"`
				EstimatedTimeRemain time.Duration `json:"estimated_time_remaining" yaml:"estimated_time_remaining"`
			}{
				BalanceRemain: balanceRemain,
				BlocksRemain:  blocksRemain,
				// EstimatedTimeRemain: netutil.AverageBlockTime * time.Duration(blocksRemain),
			}

			outputType, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}

			var data []byte
			if outputType == "json" {
				data, err = json.MarshalIndent(output, " ", "\t")
			} else {
				data, err = yaml.Marshal(output)
			}

			if err != nil {
				return err
			}

			return cl.ClientContext().PrintBytes(data)

		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddDeploymentIDFlags(cmd.Flags())
	cflags.MarkReqDeploymentIDFlags(cmd)

	return cmd
}
