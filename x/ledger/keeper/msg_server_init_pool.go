package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

// init multisig pool or ica pool
func (k msgServer) InitPool(goCtx context.Context, msg *types.MsgInitPool) (*types.MsgInitPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	err := k.Keeper.CheckAddress(ctx, msg.Denom, msg.Pool)
	if err != nil {
		return nil, err
	}
	// must set withdrawal addres if it is ica pool
	if icaPool, found := k.Keeper.GetIcaPoolByDelegationAddr(ctx, msg.Pool); found && icaPool.Status != types.IcaPoolStatusSetWithdrawal {
		return nil, types.ErrIcaPoolStatusUnmatch
	}

	if !k.IsBondedPoolExist(ctx, msg.Denom, msg.Pool) {
		_, found := k.GetExchangeRate(ctx, msg.Denom)
		if !found {
			k.SetExchangeRate(ctx, msg.Denom, sdk.NewInt(0), sdk.NewInt(0))
		}
		k.AddBondedPool(ctx, msg.Denom, msg.Pool)
		k.SetBondPipeline(ctx, types.NewBondPipeline(msg.Denom, msg.Pool))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInitPool,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Pool),
		),
	)

	return &types.MsgInitPoolResponse{}, nil
}
