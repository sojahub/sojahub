package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/sojahub/sojahub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdProviderSwitch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider-switch",
		Short: "Query mining provider switch state",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProviderSwitchRequest{}

			res, err := queryClient.ProviderSwitch(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
