package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/rdex/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	if !k.Keeper.HasPoolCreator(ctx, userAddress) {
		return nil, types.ErrPermissionDeny
	}

	orderTokens := sdk.Coins{msg.Token0, msg.Token1}.Sort()

	willUseSwapPoolIndex := k.Keeper.GetSwapPoolNextIndex(ctx)
	lpDenom := types.GetLpTokenDenom(willUseSwapPoolIndex)

	poolTotalUnit, addLpUnit := CalPoolUnit(sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), orderTokens[0].Amount, orderTokens[1].Amount)
	if !addLpUnit.IsPositive() {
		return nil, types.ErrAddLpUnitZero
	}

	// send coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, orderTokens); err != nil {
		return nil, err
	}
	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(lpDenom, addLpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool := types.SwapPool{
		Index:     willUseSwapPoolIndex,
		LpToken:   sdk.NewCoin(lpDenom, poolTotalUnit),
		BaseToken: orderTokens[0],
		Token:     orderTokens[1],
	}

	k.Keeper.SetSwapPool(ctx, lpDenom, &swapPool)
	k.Keeper.SetSwapPoolIndex(ctx, willUseSwapPoolIndex)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, poolTotalUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, addLpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyPoolBaseTokenBalance, swapPool.BaseToken.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokenBalance, swapPool.Token.String()),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
