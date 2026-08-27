package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cflags "pkg.akt.dev/go/cli/flags"
	types "pkg.akt.dev/go/node/provider/v1beta4"
)

const (
	flagExpectedEndsAt    = "expected-ends-at"
	flagMaintenanceType   = "type"
	flagStartsAt          = "starts-at"
	flagMaintenanceStatus = "status"
)

func GetTxProviderOpenMaintenanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "open-maintenance",
		Short:             "Open provider maintenance",
		Args:              cobra.NoArgs,
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			maintenanceType, err := readProviderMaintenanceTypeFlag(cmd)
			if err != nil {
				return err
			}
			startsAt, err := readTimeFlag(cmd, flagStartsAt)
			if err != nil {
				return err
			}
			expectedEndsAt, err := readTimeFlag(cmd, flagExpectedEndsAt)
			if err != nil {
				return err
			}
			metadataHash, err := readOptionalHashFlag(cmd, flagMetadataHash)
			if err != nil {
				return err
			}

			msg := &types.MsgOpenProviderMaintenance{
				Provider:        cctx.GetFromAddress().String(),
				MaintenanceType: maintenanceType,
				StartsAt:        startsAt,
				ExpectedEndsAt:  expectedEndsAt,
				MetadataHash:    metadataHash,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagMaintenanceType, "", "Maintenance type")
	cmd.Flags().String(flagStartsAt, "", "Maintenance start time in RFC3339 format")
	cmd.Flags().String(flagExpectedEndsAt, "", "Expected maintenance end time in RFC3339 format")
	cmd.Flags().String(flagMetadataHash, "", "Optional metadata hash in hex or sha256:<hex> form")
	mustMarkRequired(cmd, flagMaintenanceType, flagStartsAt, flagExpectedEndsAt)

	return cmd
}

func GetTxProviderCloseMaintenanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "close-maintenance [maintenance-id]",
		Short:             "Close provider maintenance",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: TxPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustClientFromContext(ctx)
			cctx := cl.ClientContext()

			maintenanceID, err := parseUint64Arg(args[0], "maintenance id")
			if err != nil {
				return err
			}

			msg := &types.MsgCloseProviderMaintenance{
				Provider:      cctx.GetFromAddress().String(),
				MaintenanceID: maintenanceID,
			}

			resp, err := cl.Tx().BroadcastMsgs(ctx, []sdk.Msg{msg})
			if err != nil {
				return err
			}

			return cl.PrintMessage(resp)
		},
	}

	cflags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetQueryProviderMaintenancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "maintenances [provider]",
		Short:             "Query provider maintenances",
		Args:              cobra.ExactArgs(1),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			status, err := readProviderMaintenanceStatusFlag(cmd)
			if err != nil {
				return err
			}
			pageReq, err := ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := cl.Query().Provider().ProviderMaintenances(ctx, &types.QueryProviderMaintenancesRequest{
				Provider:     provider.String(),
				StatusFilter: status,
				Pagination:   pageReq,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)
	cflags.AddPaginationFlagsToCmd(cmd, "provider maintenances")
	cmd.Flags().String(flagMaintenanceStatus, "", "Optional maintenance status filter")

	return cmd
}

func GetQueryProviderMaintenanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "maintenance [provider] [maintenance-id]",
		Short:             "Query provider maintenance",
		Args:              cobra.ExactArgs(2),
		PersistentPreRunE: QueryPersistentPreRunE,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl := MustLightClientFromContext(ctx)

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			maintenanceID, err := parseUint64Arg(args[1], "maintenance id")
			if err != nil {
				return err
			}

			res, err := cl.Query().Provider().ProviderMaintenance(ctx, &types.QueryProviderMaintenanceRequest{
				Provider:      provider.String(),
				MaintenanceId: maintenanceID,
			})
			if err != nil {
				return err
			}

			return cl.ClientContext().PrintProto(res)
		},
	}

	cflags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func readProviderMaintenanceTypeFlag(cmd *cobra.Command) (types.ProviderMaintenanceType, error) {
	val, err := cmd.Flags().GetString(flagMaintenanceType)
	if err != nil {
		return types.ProviderMaintenanceType_provider_maintenance_type_unspecified, err
	}
	return parseProviderMaintenanceType(val)
}

func readProviderMaintenanceStatusFlag(cmd *cobra.Command) (types.ProviderMaintenanceStatus, error) {
	val, err := cmd.Flags().GetString(flagMaintenanceStatus)
	if err != nil {
		return types.ProviderMaintenanceStatus_provider_maintenance_status_unspecified, err
	}
	if val == "" {
		return types.ProviderMaintenanceStatus_provider_maintenance_status_unspecified, nil
	}
	return parseProviderMaintenanceStatus(val)
}

func parseProviderMaintenanceType(val string) (types.ProviderMaintenanceType, error) {
	switch normalizeEnumInput(val) {
	case "planned", "provider_maintenance_type_planned":
		return types.ProviderMaintenanceType_provider_maintenance_type_planned, nil
	case "emergency", "provider_maintenance_type_emergency":
		return types.ProviderMaintenanceType_provider_maintenance_type_emergency, nil
	case "security", "provider_maintenance_type_security":
		return types.ProviderMaintenanceType_provider_maintenance_type_security, nil
	case "network", "provider_maintenance_type_network":
		return types.ProviderMaintenanceType_provider_maintenance_type_network, nil
	case "capacity", "provider_maintenance_type_capacity":
		return types.ProviderMaintenanceType_provider_maintenance_type_capacity, nil
	default:
		return types.ProviderMaintenanceType_provider_maintenance_type_unspecified, fmt.Errorf("invalid provider maintenance type %q", val)
	}
}

func parseProviderMaintenanceStatus(val string) (types.ProviderMaintenanceStatus, error) {
	switch normalizeEnumInput(val) {
	case "scheduled", "provider_maintenance_status_scheduled":
		return types.ProviderMaintenanceStatus_provider_maintenance_status_scheduled, nil
	case "active", "provider_maintenance_status_active":
		return types.ProviderMaintenanceStatus_provider_maintenance_status_active, nil
	case "elapsed", "provider_maintenance_status_elapsed":
		return types.ProviderMaintenanceStatus_provider_maintenance_status_elapsed, nil
	case "closed", "provider_maintenance_status_closed":
		return types.ProviderMaintenanceStatus_provider_maintenance_status_closed, nil
	default:
		return types.ProviderMaintenanceStatus_provider_maintenance_status_unspecified, fmt.Errorf("invalid provider maintenance status %q", val)
	}
}
