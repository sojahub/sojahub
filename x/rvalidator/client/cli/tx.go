package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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

	cmd.AddCommand(CmdInitRValidator())
	cmd.AddCommand(CmdUpdateRValidator())
	cmd.AddCommand(CmdSetCycleSeconds())
	cmd.AddCommand(CmdSetShuffleSeconds())
	cmd.AddCommand(CmdUpdateRValidatorReport())
	cmd.AddCommand(CmdAddRValidator())
	cmd.AddCommand(CmdRmRValidator())
	// this line is used by starport scaffolding # 1

	return cmd
}
