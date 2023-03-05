package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sojahub/sojahub/x/mining/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAddStakePool())
	cmd.AddCommand(CmdAddStakeItem())
	cmd.AddCommand(CmdAddRewardPool())
	cmd.AddCommand(CmdStake())
	cmd.AddCommand(CmdClaimReward())
	cmd.AddCommand(CmdWithdraw())
	cmd.AddCommand(CmdUpdateStakeItem())
	cmd.AddCommand(CmdAddMiningProvider())
	cmd.AddCommand(CmdRmMiningProvider())
	cmd.AddCommand(CmdAddRewardToken())
	cmd.AddCommand(CmdSetMaxRewardPoolNumber())
	cmd.AddCommand(CmdUpdateRewardPool())
	cmd.AddCommand(CmdToggleProviderSwitch())
	cmd.AddCommand(CmdSetMaxStakeItemNumber())
	cmd.AddCommand(CmdAddReward())
	cmd.AddCommand(CmdToggleEmergencySwitch())
	cmd.AddCommand(CmdEmergencyWithdraw())
	cmd.AddCommand(CmdAddStakeToken())
	cmd.AddCommand(CmdRmRewardToken())
	cmd.AddCommand(CmdRmStakeToken())
	cmd.AddCommand(CmdSetStakeItemLimit())
	cmd.AddCommand(CmdWithdrawRewardToken())
	// this line is used by starport scaffolding # 1

	return cmd
}
