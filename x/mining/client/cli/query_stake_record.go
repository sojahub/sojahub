package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdStakeRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-record [user-address] [stake-pool-index] [stake-record-index]",
		Short: "Query stake record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUserAddress := args[0]
			stakePoolIndex, err := sdk.ParseUint(args[1])
			if err != nil {
				return err
			}
			reqStakeRecordIndex, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStakeRecordRequest{
				UserAddress:      reqUserAddress,
				StakePoolIndex:   uint32(stakePoolIndex.Uint64()),
				StakeRecordIndex: uint32(reqStakeRecordIndex.Uint64()),
			}

			res, err := queryClient.StakeRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
