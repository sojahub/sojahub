package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IcaPoolList(goCtx context.Context, req *types.QueryIcaPoolListRequest) (*types.QueryIcaPoolListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryIcaPoolListResponse{
		IcaPoolList: k.GetIcaPoolDetailList(ctx, req.Denom),
	}, nil
}
