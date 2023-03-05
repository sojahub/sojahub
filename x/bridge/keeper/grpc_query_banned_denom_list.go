package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BannedDenomList(goCtx context.Context, req *types.QueryBannedDenomListRequest) (*types.QueryBannedDenomListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryBannedDenomListResponse{List: k.GetBannedDenomList(ctx)}, nil
}
