package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TotalProtocolFee(goCtx context.Context, req *types.QueryTotalProtocolFeeRequest) (*types.QueryTotalProtocolFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryTotalProtocolFeeResponse{
		TotalProtocolFeeList: k.GetAllTotalProtocolFee(ctx),
	}, nil
}
