package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	dv1 "pkg.akt.dev/go/node/deployment/v1"
	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
	"pkg.akt.dev/go/node/escrow/module"
	etypes "pkg.akt.dev/go/node/escrow/v1"
	mv1 "pkg.akt.dev/go/node/market/v1"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
	"pkg.akt.dev/go/node/utils"
)

var errNoLeaseMatches = errors.New("leases for deployment do not exist")

func GetQueryEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        module.ModuleName,
		Short:                      "Escrow query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetQueryEscrowAccountsCmd(),
		GetQueryEscrowPaymentsCmd(),
		GetQueryEscrowBlocksRemainingCmd(),
	)

	return cmd
}

func GetQueryEscrowAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts <state> <xid>",
		Short: "Query for escrow account(s)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query escrow accounts.
Arguments are optional. XID cannot be provided without state
<state> - allowed values are open|closed|overdrawn
<xid> - format must follow template - [scope]</xid...>
        allowed scope values are deployment|bid
        xid examples:
        - deployment
        - deployment/akash1...
        - deployment/akash1.../dseq
Examples (pagination limits apply to all examples below):
1. Return all accounts
$ %[1]s query %[2]s accounts
2. Return accounts in open state
$ %[1]s query %[2]s accounts open
3. Return accounts in open state for deployment scope
$ %[1]s query %[2]s accounts open deployment
3. Return accounts in open state for deployment scope
$ %[1]s query %[2]s accounts open deployment/akash1...
`,
				version.AppName, authz.ModuleName)),
		Args:              cobra.RangeArgs(0, 2),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			pageReq, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var state string
			var xid string

			if len(args) > 0 {
				state = args[0]
				switch state {
				case "open", "closed", "overdrawn":
				default:
					return fmt.Errorf("invalid account state. allowed values - open|closed|overdrawn")
				}
			}

			if len(args) > 1 {
				xid = args[1]
				xid = strings.TrimPrefix(xid, "/")
				xid = strings.TrimSuffix(xid, "/")

				parts := strings.Split(xid, "/")

				if len(parts) < 1 {
					return fmt.Errorf("invalid xid format")
				}

				switch parts[0] {
				case "deployment", "bid":
				default:
					return fmt.Errorf("invalid xid scope prefix. allowed values - deployment|bid")
				}
			}
			req := &etypes.QueryAccountsRequest{
				State:      state,
				XID:        xid,
				Pagination: pageReq,
			}

			res, err := cl.Query().Escrow().Accounts(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "escrow")

	return cmd
}

func GetQueryEscrowPaymentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payments <state> <xid>",
		Short: "Query for escrow account(s)",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query escrow accounts.
Arguments are optional. XID cannot be provided without state
<state> - allowed values are open|closed|overdrawn
<xid> - format must follow template - [scope]</xid...>
        allowed scope values are deployment|bid
        xid examples:
        - deployment
        - deployment/akash1...
        - deployment/akash1.../dseq
Examples (pagination limits apply to all examples below):
1. Return all accounts
$ %[1]s query %[2]s accounts
2. Return accounts in open state
$ %[1]s query %[2]s accounts open
3. Return accounts in open state for deployment scope
$ %[1]s query %[2]s accounts open deployment
3. Return accounts in open state for deployment scope
$ %[1]s query %[2]s accounts open deployment/akash1...
`,
				version.AppName, authz.ModuleName)),
		Args:              cobra.RangeArgs(0, 2),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			pageReq, err := sdkclient.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var state string
			var xid string

			if len(args) > 0 {
				state = args[0]
				switch state {
				case "open", "closed", "overdrawn":
				default:
					return fmt.Errorf("invalid account state. allowed values - open|closed|overdrawn")
				}
			}

			if len(args) > 1 {
				xid = args[1]
				xid = strings.TrimPrefix(xid, "/")
				xid = strings.TrimSuffix(xid, "/")

				parts := strings.Split(xid, "/")

				if len(parts) < 1 {
					return fmt.Errorf("invalid xid format")
				}

				switch parts[0] {
				case "deployment", "bid":
				default:
					return fmt.Errorf("invalid xid scope prefix. allowed values - deployment|bid")
				}
			}
			req := &etypes.QueryPaymentsRequest{
				State:      state,
				XID:        xid,
				Pagination: pageReq,
			}

			res, err := cl.Query().Escrow().Payments(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "escrow")

	return cmd
}

func GetQueryEscrowBlocksRemainingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "blocks-remaining",
		Short:             "Compute the number of blocks remaining for an escrow account",
		Args:              cobra.ExactArgs(0),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

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

			balanceRemain := utils.LeaseCalcBalanceRemain(res.EscrowAccount.State.Funds[0].Amount,
				int64(blockchainHeight),
				res.EscrowAccount.State.SettledAt,
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
