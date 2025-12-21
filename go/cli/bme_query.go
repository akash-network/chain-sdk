package cli

import (
	"github.com/spf13/cobra"

	sdkclient "github.com/cosmos/cosmos-sdk/client"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/bme/v1"
)

func GetQueryBMECmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "BME query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		GetBMEParamsCmd(),
		GetBMEVaultStateCmd(),
		GetBMECollateralRatioCmd(),
		GetBMECircuitBreakerStatusCmd(),
	)

	return cmd
}

// GetBMEParamsCmd returns the command to query BME module parameters
func GetBMEParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "params",
		Short:             "Query the current BME module parameters",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryParamsRequest{}

			res, err := cl.Query().BME().Params(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetBMEVaultStateCmd returns the command to query the BME vault state
func GetBMEVaultStateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "vault-state",
		Short:             "Query the current BME vault state",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryVaultStateRequest{}

			res, err := cl.Query().BME().VaultState(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetBMECollateralRatioCmd returns the command to query the BME collateral ratio
func GetBMECollateralRatioCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "collateral-ratio",
		Short:             "Query the current BME collateral ratio",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryCollateralRatioRequest{}

			res, err := cl.Query().BME().CollateralRatio(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetBMECircuitBreakerStatusCmd returns the command to query the BME circuit breaker status
func GetBMECircuitBreakerStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "circuit-breaker-status",
		Short:             "Query the current BME circuit breaker status",
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			req := &types.QueryCircuitBreakerStatusRequest{}

			res, err := cl.Query().BME().CircuitBreakerStatus(ctx, req)
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}
