package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakeReward(goCtx context.Context, req *types.QueryStakeRewardRequest) (*types.QueryStakeRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	curBlockTime := uint64(ctx.BlockTime().Unix())

	userStakeRecord, found := k.GetUserStakeRecord(ctx, req.StakeUserAddress, req.StakePoolIndex, req.StakeRecordIndex)
	if !found {
		return nil, types.ErrUserStakeRecordNotExist
	}
	stakePool, found := k.GetStakePool(ctx, req.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	updateStakePool(stakePool, curBlockTime)
	rewardTokens := calRewardTokens(stakePool, userStakeRecord)

	return &types.QueryStakeRewardResponse{
		RewardTokens: rewardTokens,
	}, nil
}
