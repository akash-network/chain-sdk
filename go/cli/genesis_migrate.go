package cli

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	v043 "github.com/cosmos/cosmos-sdk/x/genutil/migrations/v043"
	v046 "github.com/cosmos/cosmos-sdk/x/genutil/migrations/v046"
	v047 "github.com/cosmos/cosmos-sdk/x/genutil/migrations/v047"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"

	cflags "pkg.akt.dev/go/cli/flags"
)

// Allow applications to extend and modify the migration process.
//
// Ref: https://github.com/cosmos/cosmos-sdk/issues/5041
var migrationMap = types.MigrationMap{
	"v0.43": v043.Migrate, // NOTE: v0.43, v0.44 and v0.45 are genesis compatible.
	"v0.46": v046.Migrate,
	"v0.47": v047.Migrate,
}

// GetMigrationCallback returns a MigrationCallback for a given version.
func GetMigrationCallback(version string) types.MigrationCallback {
	return migrationMap[version]
}

// MigrateGenesisCmd returns a command to execute genesis state migration.
// Applications should pass their own migration map to this function.
// When the application migration includes a SDK migration, the Cosmos SDK migration function should as well be called.
func MigrateGenesisCmd(migrations types.MigrationMap) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate [target-version] [genesis-file]",
		Short:   "Migrate genesis to a specified target version",
		Long:    "Migrate the source genesis into the target version and print to STDOUT",
		Example: fmt.Sprintf("%s genesis migrate v0.47 /path/to/genesis.json --chain-id=cosmoshub-3 --genesis-time=2019-04-22T17:00:00Z", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return MigrateHandler(cmd, args, migrations)
		},
	}

	cmd.Flags().String(cflags.FlagGenesisTime, "", "Override genesis_time with this flag")
	cmd.Flags().String(flags.FlagChainID, "", "Override chain_id with this flag")
	cmd.Flags().String(flags.FlagOutputDocument, "", "Exported state is written to the given file instead of STDOUT")

	return cmd
}

// MigrateHandler handles the migration command with a migration map as input,
// returning an error upon failure.
func MigrateHandler(cmd *cobra.Command, args []string, migrations types.MigrationMap) error {
	clientCtx := client.GetClientContextFromCmd(cmd)

	target := args[0]
	migrationFunc, ok := migrations[target]
	if !ok || migrationFunc == nil {
		versions := slices.Sorted(maps.Keys(migrations))
		return fmt.Errorf("unknown migration function for version: %s (supported versions %s)", target, strings.Join(versions, ", "))
	}

	importGenesis := args[1]
	appGenesis, err := types.AppGenesisFromFile(importGenesis)
	if err != nil {
		return err
	}

	if err := appGenesis.ValidateAndComplete(); err != nil {
		return fmt.Errorf("make sure that you have correctly migrated all CometBFT consensus params. Refer the UPGRADING.md (%s): %w", chainUpgradeGuide, err)
	}

	// Since some default values are valid values, we just print to
	// make sure the user didn't forget to update these values.
	if appGenesis.Consensus.Params.Evidence.MaxBytes == 0 {
		fmt.Printf("Warning: consensus.params.evidence.max_bytes is set to 0. If this is"+
			" deliberate, feel free to ignore this warning. If not, please have a look at the chain"+
			" upgrade guide at %s.\n", chainUpgradeGuide)
	}

	var initialState types.AppMap
	if err := json.Unmarshal(appGenesis.AppState, &initialState); err != nil {
		return fmt.Errorf("failed to JSON unmarshal initial genesis state: %w", err)
	}

	newGenState, err := migrationFunc(initialState, clientCtx)
	if err != nil {
		return fmt.Errorf("failed to migrate genesis state: %w", err)
	}

	appGenesis.AppState, err = json.Marshal(newGenState)
	if err != nil {
		return fmt.Errorf("failed to JSON marshal migrated genesis state: %w", err)
	}

	genesisTime, _ := cmd.Flags().GetString(cflags.FlagGenesisTime)
	if genesisTime != "" {
		var t time.Time

		err := t.UnmarshalText([]byte(genesisTime))
		if err != nil {
			return fmt.Errorf("failed to unmarshal genesis time: %w", err)
		}

		appGenesis.GenesisTime = t
	}

	chainID, _ := cmd.Flags().GetString(cflags.FlagChainID)
	if chainID != "" {
		appGenesis.ChainID = chainID
	}

	bz, err := json.Marshal(appGenesis)
	if err != nil {
		return fmt.Errorf("failed to marshal app genesis: %w", err)
	}

	outputDocument, _ := cmd.Flags().GetString(cflags.FlagOutputDocument)
	if outputDocument == "" {
		cmd.Println(string(bz))
		return nil
	}

	if err = appGenesis.SaveAs(outputDocument); err != nil {
		return err
	}

	return nil
}
