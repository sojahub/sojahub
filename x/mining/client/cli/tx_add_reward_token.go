package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdAddRewardToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-reward-token [denom] [min-total-reward-amount] [min-reward-per-second]",
		Short: "Add reward token",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argMinTotalRewardAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("argMinTotalRewardAmount err")
			}
			argMinRewardPerSecond, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("argMinRewardPerSecond err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddRewardToken(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argMinTotalRewardAmount,
				argMinRewardPerSecond,
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
