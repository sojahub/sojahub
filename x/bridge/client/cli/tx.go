package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sojahub/sojahub/x/bridge/types"
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

	cmd.AddCommand(CmdSetResourceidToDenom())
	cmd.AddCommand(CmdDeposit())
	cmd.AddCommand(CmdAddChainId())
	cmd.AddCommand(CmdVoteProposal())
	cmd.AddCommand(CmdRmChainId())
	cmd.AddCommand(CmdSetRelayFeeReceiver())
	cmd.AddCommand(CmdSetRelayFee())
	cmd.AddCommand(CmdAddBannedDenom())
	cmd.AddCommand(CmdRmBannedDenom())
	// this line is used by starport scaffolding # 1

	return cmd
}
