package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) AddStakeToken(goCtx context.Context, msg *types.MsgAddStakeToken) (*types.MsgAddStakeTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.AddStakeToken(ctx, msg.Denom)

	return &types.MsgAddStakeTokenResponse{}, nil
}
