package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/rstaking/types"
)

var _ = strconv.Itoa(0)

func CmdAddDelegatorToWhitelist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-delegator-to-whitelist [del-address]",
		Short: "Add delegator address to whitelist",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDelAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDelegatorToWhitelist(
				clientCtx.GetFromAddress().String(),
				argDelAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
