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

func CmdWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [stake-pool-index] [stake-record-index] [withdraw-amount]",
		Short: "Withdraw",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argStakePoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			argStakeRecordIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}

			argWithdrawAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("arg withdrawAmount err")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(
				clientCtx.GetFromAddress().String(),
				uint32(argStakePoolIndex.Uint64()),
				uint32(argStakeRecordIndex.Uint64()),
				argWithdrawAmount,
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
