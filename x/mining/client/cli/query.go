package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sojahub/sojahub/x/mining/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group mining queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdStakePoolInfo())

	cmd.AddCommand(CmdStakeItemList())

	cmd.AddCommand(CmdStakeReward())

	cmd.AddCommand(CmdStakeRecordCount())

	cmd.AddCommand(CmdStakeRecord())

	cmd.AddCommand(CmdStakeRecordList())

	cmd.AddCommand(CmdStakePoolList())

	cmd.AddCommand(CmdMiningProviderList())

	cmd.AddCommand(CmdRewardTokenList())

	cmd.AddCommand(CmdMaxRewardPoolNumber())

	cmd.AddCommand(CmdMaxStakeItemNumber())

	cmd.AddCommand(CmdProviderSwitch())
	// this line is used by starport scaffolding # 1

	return cmd
}
