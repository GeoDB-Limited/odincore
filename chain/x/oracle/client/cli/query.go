package cli

import (
	"fmt"
	"github.com/GeoDB-Limited/odincore/chain/x/common/client/cli"
	"github.com/cosmos/cosmos-sdk/version"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	clientcmn "github.com/GeoDB-Limited/odincore/chain/x/oracle/client/common"
	"github.com/GeoDB-Limited/odincore/chain/x/oracle/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	oracleCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	oracleCmd.AddCommand(flags.GetCommands(
		GetQueryCmdParams(storeKey, cdc),
		GetQueryCmdCounts(storeKey, cdc),
		GetQueryCmdDataSource(storeKey, cdc),
		GetQueryCmdOracleScript(storeKey, cdc),
		GetQueryCmdRequest(storeKey, cdc),
		GetQueryCmdRequestSearch(storeKey, cdc),
		GetQueryCmdValidatorStatus(storeKey, cdc),
		GetQueryCmdReporters(storeKey, cdc),
		GetQueryActiveValidators(storeKey, cdc),
		GetQueryPendingRequests(storeKey, cdc),
		GetCmdQueryDataProvidersPool(storeKey, cdc),
	)...)
	return oracleCmd
}

// GetQueryCmdParams implements the query parameters command.
func GetQueryCmdParams(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "params",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", route, types.QueryParams))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.Params{})
		},
	}
}

// GetQueryCmdCounts implements the query counts command.
func GetQueryCmdCounts(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "counts",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", route, types.QueryCounts))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.QueryCountsResult{})
		},
	}
}

// GetQueryCmdDataSource implements the query data source command.
func GetQueryCmdDataSource(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "data-source [id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", route, types.QueryDataSources, args[0]))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.DataSource{})
		},
	}
}

// GetQueryCmdOracleScript implements the query oracle script command.
func GetQueryCmdOracleScript(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "oracle-script [id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", route, types.QueryOracleScripts, args[0]))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.OracleScript{})
		},
	}
}

// GetQueryCmdRequest implements the query request command.
func GetQueryCmdRequest(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "request [id]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", route, types.QueryRequests, args[0]))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.QueryRequestResult{})
		},
	}
}

// GetQueryCmdRequestSearch implements the search request command.
func GetQueryCmdRequestSearch(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "request-search [oracle-script-id] [calldata] [ask-count] [min-count]",
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := clientcmn.QuerySearchLatestRequest(route, cliCtx, args[0], args[1], args[2], args[3])
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.QueryRequestResult{})
		},
	}
}

// GetQueryCmdValidatorStatus implements the query reporter list of validator command.
func GetQueryCmdValidatorStatus(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "validator [validator]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", route, types.QueryValidatorStatus, args[0]))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &types.ValidatorStatus{})
		},
	}
}

// GetQueryCmdReporters implements the query reporter list of validator command.
func GetQueryCmdReporters(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "reporters [validator]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", route, types.QueryReporters, args[0]))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &[]sdk.AccAddress{})
		},
	}
}

// GetQueryActiveValidators implements the query active validators command.
func GetQueryActiveValidators(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "active-validators",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", route, types.QueryActiveValidators))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &[]types.QueryActiveValidatorResult{})
		},
	}
}

// GetQueryPendingRequests implements the query pending requests command.
func GetQueryPendingRequests(route string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "pending-requests [validator]",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			path := fmt.Sprintf("custom/%s/%s", route, types.QueryPendingRequests)
			if len(args) == 1 {
				path += "/" + args[0]
			}

			bz, _, err := cliCtx.Query(path)
			if err != nil {
				return err
			}

			return cli.PrintOutput(cliCtx, cdc, bz, &[]types.RequestID{})
		},
	}
}

// GetCmdQueryDataProvidersPool returns the command for fetching community pool info
func GetCmdQueryDataProvidersPool(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "data-providers-pool",
		Args:  cobra.NoArgs,
		Short: "Query the amount of coins in the data providers pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all coins in the data providers pool.

Example:
$ %s query oracle data-providers-pool
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			bz, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryDataProvidersPool))
			if err != nil {
				return err
			}
			return cli.PrintOutput(cliCtx, cdc, bz, &[]sdk.DecCoin{})
		},
	}
}
