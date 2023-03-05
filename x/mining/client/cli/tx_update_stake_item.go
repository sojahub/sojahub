package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/utils"
	"github.com/sojahub/sojahub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdUpdateStakeItem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-stake-item [index] [lock-second] [power-reward-rate] [enable]",
		Short: "Update stake item",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			argLockSecond, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}
			argPowerRewardRate, err := utils.NewDecFromStr(args[2])
			if err != nil {
				return err
			}

			enable := true
			if args[3] == "false" {
				enable = false
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateStakeItem(
				clientCtx.GetFromAddress().String(),
				uint32(argIndex.Uint64()),
				argLockSecond.Uint64(),
				argPowerRewardRate,
				enable,
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
