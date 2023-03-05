package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdRmBannedDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm-banned-denom [chain-id] [denom]",
		Short: "Remove Banned denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChainId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argDenom := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRmBannedDenom(
				clientCtx.GetFromAddress().String(),
				uint32(argChainId.Uint64()),
				argDenom,
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
