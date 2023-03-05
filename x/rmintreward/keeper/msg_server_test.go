package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sojahub/sojahub/testutil/keeper"
	"github.com/sojahub/sojahub/x/rmintreward/keeper"
	"github.com/sojahub/sojahub/x/rmintreward/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RmintrewardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
