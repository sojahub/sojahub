package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	userStakeRecord, found := k.Keeper.GetUserStakeRecord(ctx, msg.Creator, msg.StakePoolIndex, msg.StakeRecordIndex)
	if !found {
		return nil, types.ErrUserStakeRecordNotExist
	}

	if msg.WithdrawAmount.GT(userStakeRecord.StakedAmount) {
		return nil, types.ErrWithdrawAmountMoreThanStakeRecord
	}

	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakePoolIndex)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	if stakePool.EmergencySwitch {
		return nil, types.ErrEmergencySwitchOpen
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	if userStakeRecord.LockEndTimestamp > curBlockTime {
		return nil, types.ErrStakeTokenStillLocked
	}

	updateStakePool(stakePool, curBlockTime)
	willClaimCoins := calRewardTokens(stakePool, userStakeRecord)

	willRmPower := msg.WithdrawAmount.Mul(userStakeRecord.StakedPower).Quo(userStakeRecord.StakedAmount)

	stakePool.TotalStakedAmount = stakePool.TotalStakedAmount.Sub(msg.WithdrawAmount)
	if stakePool.TotalStakedAmount.IsNegative() {
		stakePool.TotalStakedAmount = sdk.ZeroInt()
	}
	stakePool.TotalStakedPower = stakePool.TotalStakedPower.Sub(willRmPower)
	if stakePool.TotalStakedPower.IsNegative() {
		stakePool.TotalStakedPower = sdk.ZeroInt()
	}

	userStakeRecord.StakedAmount = userStakeRecord.StakedAmount.Sub(msg.WithdrawAmount)
	if userStakeRecord.StakedAmount.IsNegative() {
		userStakeRecord.StakedAmount = sdk.ZeroInt()
	}
	userStakeRecord.StakedPower = userStakeRecord.StakedPower.Sub(willRmPower)
	if userStakeRecord.StakedPower.IsNegative() {
		userStakeRecord.StakedPower = sdk.ZeroInt()
	}

	// recheck stake amount and power
	if userStakeRecord.StakedAmount.IsZero() {
		if userStakeRecord.StakedPower.IsPositive() {
			stakePool.TotalStakedPower = stakePool.TotalStakedPower.Sub(userStakeRecord.StakedPower)
			if stakePool.TotalStakedPower.IsNegative() {
				stakePool.TotalStakedPower = sdk.ZeroInt()
			}
			userStakeRecord.StakedPower = sdk.ZeroInt()
		}
	}

	setNewRewardDebt(stakePool, userStakeRecord)

	willClaimCoins = willClaimCoins.Add(sdk.NewCoin(stakePool.StakeTokenDenom, msg.WithdrawAmount))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, willClaimCoins); err != nil {
		return nil, err
	}

	k.SetStakePool(ctx, stakePool)
	k.SetUserStakeRecord(ctx, userStakeRecord)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", msg.StakePoolIndex)),
			sdk.NewAttribute(types.AttributeKeyStakeRecordIndex, fmt.Sprintf("%d", msg.StakeRecordIndex)),
			sdk.NewAttribute(types.AttributeKeyClaimedTokens, willClaimCoins.String()),
			sdk.NewAttribute(types.AttributeKeyWithdrawAmount, msg.WithdrawAmount.String()),
		),
	)

	return &types.MsgWithdrawResponse{}, nil
}
