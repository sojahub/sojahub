package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/rdex/types"
)

var _ = strconv.Itoa(0)

func CmdRemoveLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-liquidity [swap-pool-index] [rm-unit] [swap-unit] [min-out-token0] [min-out-token1] [input-token-denom]",
		Short: "Remove liquidity",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			swapPoolIndex, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			argRmUnit, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("rm unit params err")
			}
			argSwapUnit, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("swap unit params err")
			}

			minOutToken0, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}
			minOutToken1, err := sdk.ParseCoinNormalized(args[4])
			if err != nil {
				return err
			}

			inputTokenDenom := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveLiquidity(
				clientCtx.GetFromAddress().String(),
				uint32(swapPoolIndex.Uint64()),
				argRmUnit,
				argSwapUnit,
				minOutToken0,
				minOutToken1,
				inputTokenDenom,
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
