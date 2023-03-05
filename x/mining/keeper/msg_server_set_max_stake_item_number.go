package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/mining/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) SetMaxStakeItemNumber(goCtx context.Context, msg *types.MsgSetMaxStakeItemNumber) (*types.MsgSetMaxStakeItemNumberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.SetMaxStakeItemNumber(ctx, msg.Number)

	return &types.MsgSetMaxStakeItemNumberResponse{}, nil
}
