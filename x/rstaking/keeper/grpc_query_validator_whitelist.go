package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/x/rstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorWhitelist(goCtx context.Context, req *types.QueryValidatorWhitelistRequest) (*types.QueryValidatorWhitelistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	list := k.GetValAddressWhitelist(ctx)

	return &types.QueryValidatorWhitelistResponse{
		ValAddress: list,
	}, nil
}
