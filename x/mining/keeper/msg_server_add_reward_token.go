package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) AddRewardToken(goCtx context.Context, msg *types.MsgAddRewardToken) (*types.MsgAddRewardTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rewardToken := types.RewardToken{
		RewardTokenDenom:     msg.Denom,
		MinTotalRewardAmount: msg.MinTotalRewardAmount,
		MinRewardPerSecond:   msg.MinRewardPerSecond,
	}

	k.Keeper.AddRewardToken(ctx, &rewardToken)

	return &types.MsgAddRewardTokenResponse{}, nil
}
