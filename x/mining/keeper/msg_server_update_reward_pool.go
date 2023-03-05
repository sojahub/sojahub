package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
)

func (k msgServer) UpdateRewardPool(goCtx context.Context, msg *types.MsgUpdateRewardPool) (*types.MsgUpdateRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	// find reward pool
	var willUseRewardPool *types.RewardPool
	for _, rewardPool := range stakePool.RewardPools {
		if rewardPool.Index == msg.RewardPoolIndex {
			willUseRewardPool = rewardPool
		}
	}
	if willUseRewardPool == nil {
		return nil, types.ErrRewardPoolNotExist
	}
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) && willUseRewardPool.Creator != msg.Creator {
		return nil, types.ErrUpdateRewardPoolPermissionDeny
	}

	// check minRewardPerSecond
	rewardToken, found := k.Keeper.GetRewardToken(ctx, willUseRewardPool.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}
	if msg.RewardPerSecond.LT(rewardToken.MinRewardPerSecond) {
		return nil, types.ErrRewardPerSecondLessThanLimit
	}

	// check reward second and max lock second
	maxLockSecond := uint64(0)
	for _, stakeItemInfo := range k.Keeper.GetStakeItemList(ctx, msg.StakePoolIndex) {
		if stakeItemInfo.LockSecond > maxLockSecond {
			maxLockSecond = stakeItemInfo.LockSecond
		}
	}
	if willUseRewardPool.TotalRewardAmount.Quo(msg.RewardPerSecond).LT(sdk.NewIntFromUint64(maxLockSecond)) {
		return nil, types.ErrRewardSecondsLessThanMaxLockSeconds
	}

	willUseRewardPool.RewardPerSecond = msg.RewardPerSecond

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgUpdateRewardPoolResponse{}, nil
}
