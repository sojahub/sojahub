package rbank_test

import (
	"testing"

	keepertest "github.com/sojahub/sojahub/testutil/keeper"
	"github.com/sojahub/sojahub/testutil/nullify"
	"github.com/sojahub/sojahub/x/rbank"
	"github.com/sojahub/sojahub/x/rbank/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RbankKeeper(t)
	rbank.InitGenesis(ctx, *k, genesisState)
	got := rbank.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
