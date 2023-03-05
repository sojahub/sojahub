package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

func (k msgServer) ToggleUnbondSwitch(goCtx context.Context, msg *types.MsgToggleUnbondSwitch) (*types.MsgToggleUnbondSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.ToggleUnbondSwitch(ctx, msg.Denom)

	return &types.MsgToggleUnbondSwitchResponse{}, nil
}
