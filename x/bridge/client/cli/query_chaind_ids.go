package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/bridge/types"
)

var _ = strconv.Itoa(0)

func CmdChaindIds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-ids",
		Short: "Query supported chain ids",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryChaindIdsRequest{}

			res, err := queryClient.ChaindIds(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
