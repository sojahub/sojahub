package cli

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAddNewPool())
	cmd.AddCommand(CmdRemovePool())
	cmd.AddCommand(CmdSetEraUnbondLimit())
	cmd.AddCommand(CmdSetInitBond())
	cmd.AddCommand(CmdSetChainBondingDuration())
	cmd.AddCommand(CmdSetPoolDetail())
	cmd.AddCommand(CmdSetLeastBond())
	cmd.AddCommand(CmdClearCurrentEraSnapShots())
	cmd.AddCommand(CmdSetCommission())
	cmd.AddCommand(CmdSetReceiver())
	cmd.AddCommand(CmdSetUnbondFee())
	cmd.AddCommand(CmdSetUnbondCommission())
	cmd.AddCommand(CmdSetChainEra())
	cmd.AddCommand(CmdActiveReport())
	cmd.AddCommand(CmdBondReport())
	cmd.AddCommand(CmdBondAndReportActive())
	cmd.AddCommand(CmdWithdrawReport())
	cmd.AddCommand(CmdTransferReport())
	cmd.AddCommand(CmdLiquidityUnbond())
	cmd.AddCommand(CmdExecuteBondProposal())
	cmd.AddCommand(CmdSubmitSignature())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdLiquidityUnbond() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidity-unbond [pool] [value] [recipient]",
		Short: "Broadcast message liquidity_unbond",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPool := args[0]
			argValue, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			argRecipient := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLiquidityUnbond(
				clientCtx.GetFromAddress(),
				argPool,
				argValue,
				argRecipient,
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

func CmdSubmitSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-signature [denom] [era] [pool] [tx-type] [prop-id] [signature]",
		Short: "Broadcast message submit_signature",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			argPool := args[2]
			argTxType, ok := types.OriginalTxType_value[args[3]]
			if !ok {
				return fmt.Errorf("invalid txtype")
			}

			argPropId, err := hex.DecodeString(args[4])
			if err != nil {
				return err
			}
			argSignature := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitSignature(
				clientCtx.GetFromAddress().String(),
				argDenom,
				uint32(argEra),
				argPool,
				types.OriginalTxType(argTxType),
				argPropId,
				argSignature,
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
