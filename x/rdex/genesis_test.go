package rdex_test

import (
	"testing"

	keepertest "github.com/sojahub/sojahub/testutil/keeper"
	"github.com/sojahub/sojahub/testutil/nullify"
	"github.com/sojahub/sojahub/x/rdex"
	"github.com/sojahub/sojahub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RdexKeeper(t)
	rdex.InitGenesis(ctx, *k, genesisState)
	got := rdex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
