package keeper_test

import (
	"testing"

	testkeeper "github.com/sojahub/sojahub/testutil/keeper"
	"github.com/sojahub/sojahub/x/claim/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ClaimKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
