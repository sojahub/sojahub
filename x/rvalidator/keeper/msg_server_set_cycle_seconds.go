package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/rvalidator/types"
	sudoTypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) SetCycleSeconds(goCtx context.Context, msg *types.MsgSetCycleSeconds) (*types.MsgSetCycleSecondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	cycleSeconds := k.Keeper.GetCycleSeconds(ctx, msg.Denom)

	k.Keeper.SetCycleSeconds(ctx, &types.CycleSeconds{
		Denom:   msg.Denom,
		Version: cycleSeconds.Version + 1,
		Seconds: msg.Seconds,
	})

	return &types.MsgSetCycleSecondsResponse{}, nil
}
