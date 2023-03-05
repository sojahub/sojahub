package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/rmintreward/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClaimInfoDetail(goCtx context.Context, req *types.QueryClaimInfoDetailRequest) (*types.QueryClaimInfoDetailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	claimInfo, found := k.GetUserClaimInfo(ctx, address, req.Denom, req.Cycle, req.MintIndex)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryClaimInfoDetailResponse{ClaimInfo: claimInfo}, nil
}
