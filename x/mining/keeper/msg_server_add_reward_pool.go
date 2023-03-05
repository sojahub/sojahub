package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
)

func (k msgServer) AddRewardPool(goCtx context.Context, msg *types.MsgAddRewardPool) (*types.MsgAddRewardPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// check stake pool
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	if stakePool.EmergencySwitch {
		return nil, types.ErrEmergencySwitchOpen
	}

	// check reward token denom
	denomMap := make(map[string]bool)
	for _, rewardPool := range stakePool.RewardPools {
		if denomMap[rewardPool.RewardTokenDenom] {
			return nil, types.ErrRewardTokenDenomDuplicate
		}
		denomMap[rewardPool.RewardTokenDenom] = true
	}
	if denomMap[msg.RewardTokenDenom] {
		return nil, types.ErrRewardTokenDenomDuplicate
	}

	// check reward pool number
	maxRewardPoolNumber := k.Keeper.GetMaxRewardPoolNumber(ctx)
	if len(stakePool.RewardPools) >= int(maxRewardPoolNumber) {
		return nil, types.ErrRewardPoolNumberReachLimit
	}

	// check reward token
	rewardToken, found := k.Keeper.GetRewardToken(ctx, msg.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}
	if msg.TotalRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
		return nil, types.ErrTotalRewardAmountLessThanLimit
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
	if msg.TotalRewardAmount.Quo(msg.RewardPerSecond).LT(sdk.NewIntFromUint64(maxLockSecond)) {
		return nil, types.ErrRewardSecondsLessThanMaxLockSeconds
	}

	curBlockTime := uint64(ctx.BlockTime().Unix())

	// store reward pool
	willUseIndex := uint32(len(stakePool.RewardPools))
	willUseLastRewardTimestamp := msg.StartTimestamp
	if msg.StartTimestamp < curBlockTime {
		willUseLastRewardTimestamp = curBlockTime
	}

	rewardTokens := sdk.NewCoins(sdk.NewCoin(msg.RewardTokenDenom, msg.TotalRewardAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	stakePool.RewardPools = append(stakePool.RewardPools, &types.RewardPool{
		Index:               willUseIndex,
		RewardTokenDenom:    msg.RewardTokenDenom,
		TotalRewardAmount:   msg.TotalRewardAmount,
		LeftRewardAmount:    msg.TotalRewardAmount,
		RewardPerSecond:     msg.RewardPerSecond,
		StartTimestamp:      willUseLastRewardTimestamp,
		RewardPerPower:      sdk.ZeroInt(),
		LastRewardTimestamp: willUseLastRewardTimestamp,
		Creator:             msg.Creator,
		TotalClaimedAmount:  sdk.ZeroInt(),
	})

	k.Keeper.SetStakePool(ctx, stakePool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddRewardPool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyRewardTokenDenom, msg.RewardTokenDenom),
			sdk.NewAttribute(types.AttributeKeyTotalRewardAmount, msg.TotalRewardAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRewardPerSecond, msg.RewardPerSecond.String()),
			sdk.NewAttribute(types.AttributeKeyStartTimestamp, fmt.Sprintf("%d", msg.StartTimestamp)),
			sdk.NewAttribute(types.AttributeKeyLastRewardTimestamp, fmt.Sprintf("%d", willUseLastRewardTimestamp)),
		),
	)

	return &types.MsgAddRewardPoolResponse{}, nil
}
