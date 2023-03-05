package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolUnbondNextSequence(goCtx context.Context, req *types.QueryPoolUnbondNextSequenceRequest) (*types.QueryPoolUnbondNextSequenceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryPoolUnbondNextSequenceResponse{
		NextSequence: k.GetPoolUnbondNextSequence(ctx, req.Denom, req.Pool, req.UnlockEra),
	}, nil
}
