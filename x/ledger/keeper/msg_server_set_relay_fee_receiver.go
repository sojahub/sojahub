package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) SetRelayFeeReceiver(goCtx context.Context, msg *types.MsgSetRelayFeeReceiver) (*types.MsgSetRelayFeeReceiverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetRelayFeeReceiver(ctx, msg.Denom, receiver)

	return &types.MsgSetRelayFeeReceiverResponse{}, nil
}
