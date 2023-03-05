package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChaindIds(goCtx context.Context, req *types.QueryChaindIdsRequest) (*types.QueryChaindIdsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryChaindIdsResponse{
		ChainId: k.GetChainIdList(ctx),
	}, nil
}
