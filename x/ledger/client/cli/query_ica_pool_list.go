package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/ledger/types"
)

var _ = strconv.Itoa(0)

func CmdIcaPoolList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ica-pool-list [denom]",
		Short: "Query ica pool list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryIcaPoolListRequest{

				Denom: reqDenom,
			}

			res, err := queryClient.IcaPoolList(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
