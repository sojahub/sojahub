package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ledgermodule "github.com/sojahub/sojahub/x/ledger"
	ledgertypes "github.com/sojahub/sojahub/x/ledger/types"
	"github.com/sojahub/sojahub/x/rvote/keeper"
	"github.com/sojahub/sojahub/x/rvote/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	rvoteStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	rvoteMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	rvoteOnce        sync.Once
)

func RvoteKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	rvoteOnce.Do(func() {
		stateStore.MountStoreWithDB(rvoteStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(rvoteMemStoreKey, sdk.StoreTypeMemory, nil)
	})

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper, _ := RelayersKeeper(t)
	ledgerKeeper, _ := LedgerKeeper(t)
	require.NoError(t, stateStore.LoadLatestVersion())

	rvoteRouter := types.NewRouter()
	rvoteRouter.AddRoute(ledgertypes.RouterKey, ledgermodule.NewProposalHandler(ledgerKeeper))
	rvoteKeeper := keeper.NewKeeper(
		cdc,
		rvoteStoreKey,
		rvoteMemStoreKey,

		sudoKeeper,
		relayersKeeper,
		rvoteRouter,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return rvoteKeeper, ctx
}
