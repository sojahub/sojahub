package sudo_test

import (
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/sudo"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SudoKeeper(t)
	sudo.InitGenesis(ctx, *k, genesisState)
	got := sudo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}