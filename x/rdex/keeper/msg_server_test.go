package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sojahub/sojahub/testutil/keeper"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/sojahub/sojahub/x/rdex/keeper"
	"github.com/sojahub/sojahub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, sdk.Context) {

	k, ctx := keepertest.RdexKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx), ctx
}

func TestMsgServerCreatePoolSuccess(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(30))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10))

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// create pool
	lpDenom := types.GetLpTokenDenom(0)

	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}

	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)
}

func TestMsgServerCreatePoolFailed(t *testing.T) {
	srv, _, ctx, sdkCtx := setupMsgServer(t)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(30))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10))

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	testcases := []struct {
		name    string
		token0  sdk.Coin
		token1  sdk.Coin
		creator sdk.AccAddress
	}{
		{
			name:    "token0 zero amount",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator: poolCreator,
		},
		{
			name:    "token1 zero amount",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(30)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(0)),
			creator: poolCreator,
		},
		{
			name:    "token0 neg amount",
			token0:  sdk.Coin{sample.TestDenom, sdk.NewInt(-10)},
			token1:  sdk.Coin{sample.TestDenom1, sdk.NewInt(10)},
			creator: poolCreator,
		},
		{
			name:    "token1 neg amount",
			token0:  sdk.Coin{sample.TestDenom, sdk.NewInt(30)},
			token1:  sdk.Coin{sample.TestDenom1, sdk.NewInt(-10)},
			creator: poolCreator,
		},
		{
			name:    "not pool creator",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(30)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator: sample.TestAdminAcc,
		},
		{
			name:    "not enough token0",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(30+1)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator: poolCreator,
		},
		{
			name:    "not enough token1",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(30)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10+1)),
			creator: poolCreator,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			msgCreatePool := types.MsgCreatePool{
				Creator: tc.creator.String(),
				Token0:  tc.token0,
				Token1:  tc.token1,
			}
			err := msgCreatePool.ValidateBasic()
			if err != nil {
				t.Log(err)
				return
			}
			_, err = srv.CreatePool(ctx, &msgCreatePool)
			t.Log(err)
			require.Error(t, err)
		})
	}

}

func TestMsgServerAddRmProviderShouldWork(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	creator := sample.TestAdmin
	creator2 := sample.OriginAccAddress()

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     creator,
		UserAddress: creator2.String(),
	}
	_, err := srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, creator2))
	// rm provider
	msgRmProvider := types.MsgRmProvider{
		Creator:     creator,
		UserAddress: creator2.String(),
	}
	_, err = srv.RmProvider(ctx, &msgRmProvider)
	require.NoError(t, err)
	require.False(t, rdexKeeper.HasProvider(sdkCtx, creator2))
}

func TestMsgServerAddRmPoolCreatorShouldWork(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	creator := sample.TestAdmin
	creator2 := sample.OriginAccAddress()

	// add provider
	msgAddProvider := types.MsgAddPoolCreator{
		Creator:     creator,
		UserAddress: creator2.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasPoolCreator(sdkCtx, creator2))
	// rm provider
	msgRmProvider := types.MsgRmPoolCreator{
		Creator:     creator,
		UserAddress: creator2.String(),
	}
	_, err = srv.RmPoolCreator(ctx, &msgRmProvider)
	require.NoError(t, err)
	require.False(t, rdexKeeper.HasPoolCreator(sdkCtx, creator2))
}

func TestMsgServerToggleSwitchShouldWork(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	creator := sample.TestAdmin
	creator2 := sample.OriginAccAddress()

	require.True(t, rdexKeeper.GetProviderSwitch(sdkCtx))

	//toggle with address which is not admin
	msgToggleSwitch := types.MsgToggleProviderSwitch{
		Creator: creator2.String(),
	}
	_, err := srv.ToggleProviderSwitch(ctx, &msgToggleSwitch)
	require.Error(t, err)

	// toggle
	msgToggleSwitch2 := types.MsgToggleProviderSwitch{
		Creator: creator,
	}
	_, err = srv.ToggleProviderSwitch(ctx, &msgToggleSwitch2)
	require.NoError(t, err)
	require.False(t, rdexKeeper.GetProviderSwitch(sdkCtx))
}

func TestMsgServerAddLiquiditySuccess(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	provider := sample.OriginAccAddress()

	addToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(345e8))
	addToken1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(234e8))

	mintToCreator2Tokens := sdk.NewCoins(addToken0, addToken1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToCreator2Tokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, mintToCreator2Tokens)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8))
	lpDenom := types.GetLpTokenDenom(0)
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// crate pool
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}
	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.EqualValues(t, lpBalance, swapPool.LpToken)

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     admin,
		UserAddress: provider.String(),
	}
	_, err = srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, provider))

	// add liquidity
	newTotalLpToken := sdk.NewCoin(lpDenom, sdk.NewInt(76359469067))
	create2LpToken := sdk.NewCoin(lpDenom, sdk.NewInt(26359469067))

	msgAddLiquidity := types.MsgAddLiquidity{
		Creator: provider.String(),
		Token0:  addToken0,
		Token1:  addToken1,
	}

	_, err = srv.AddLiquidity(ctx, &msgAddLiquidity)
	require.NoError(t, err)

	swapPool, found = rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, newTotalLpToken)
	require.Equal(t, swapPool.BaseToken, token0.Add(addToken0))
	require.Equal(t, swapPool.Token, token1.Add(addToken1))

	create2LpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, provider, lpDenom)
	require.Equal(t, create2LpToken, create2LpBalance)

}

func TestMsgServerAddLiquidityFail(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	provider := sample.OriginAccAddress()
	creator3 := sample.OriginAccAddress()

	addToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(345e8))
	addToken1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(234e8))

	mintToProviderTokens := sdk.NewCoins(addToken0, addToken1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToProviderTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, mintToProviderTokens)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8))
	lpDenom := types.GetLpTokenDenom(0)
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// crate pool
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}
	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     admin,
		UserAddress: provider.String(),
	}
	_, err = srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, provider))

	testcases := []struct {
		name    string
		token0  sdk.Coin
		token1  sdk.Coin
		creator string
	}{
		{
			name:    "not provider",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(30)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator: creator3.String(),
		},
		{
			name:    "token0 token1 zero amount",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(0)),
			creator: provider.String(),
		},
		{
			name:    "token0 neg amount",
			token0:  sdk.Coin{sample.TestDenom, sdk.NewInt(-10)},
			token1:  sdk.Coin{sample.TestDenom1, sdk.NewInt(10)},
			creator: provider.String(),
		},
		{
			name:    "token1 neg amount",
			token0:  sdk.Coin{sample.TestDenom, sdk.NewInt(30)},
			token1:  sdk.Coin{sample.TestDenom1, sdk.NewInt(-10)},
			creator: provider.String(),
		},
		{
			name:    "not enough token0",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8+1)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator: provider.String(),
		},
		{
			name:    "not enough token1",
			token0:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(10)),
			token1:  sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8+1)),
			creator: provider.String(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			msgAddLiquidity := types.MsgAddLiquidity{
				Creator: tc.creator,
				Token0:  tc.token0,
				Token1:  tc.token1,
			}
			err := msgAddLiquidity.ValidateBasic()
			if err != nil {
				t.Log(err)
				return
			}

			_, err = srv.AddLiquidity(ctx, &msgAddLiquidity)
			t.Log(err)
			require.Error(t, err)
		})
	}
}

func TestMsgServerSwapSuccess(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	provider := sample.OriginAccAddress()

	addToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(90))
	addToken1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(900))

	mintToProviderTokens := sdk.NewCoins(addToken0, addToken1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToProviderTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, mintToProviderTokens)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(10))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(100))
	lpDenom := types.GetLpTokenDenom(0)
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	swapUser := sample.OriginAccAddress()
	swapUserToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(10))
	mintToSwapUserTokens := sdk.NewCoins(swapUserToken0)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToSwapUserTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, swapUser, mintToSwapUserTokens)

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// crate pool
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}
	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     admin,
		UserAddress: provider.String(),
	}
	_, err = srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, provider))

	// add liquidity
	newTotalLpToken := sdk.NewCoin(lpDenom, sdk.NewInt(100))
	create2LpToken := sdk.NewCoin(lpDenom, sdk.NewInt(90))

	msgAddLiquidity := types.MsgAddLiquidity{
		Creator: provider.String(),
		Token0:  addToken0,
		Token1:  addToken1,
	}

	_, err = srv.AddLiquidity(ctx, &msgAddLiquidity)
	require.NoError(t, err)

	swapPool, found = rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, newTotalLpToken)
	require.Equal(t, swapPool.BaseToken, token0.Add(addToken0))
	require.Equal(t, swapPool.Token, token1.Add(addToken1))

	create2LpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, provider, lpDenom)
	require.Equal(t, create2LpToken, create2LpBalance)

	// swap
	msgSwap := types.MsgSwap{
		Creator:     swapUser.String(),
		InputToken:  swapUserToken0,
		MinOutToken: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(82)),
	}
	_, err = srv.Swap(ctx, &msgSwap)
	require.NoError(t, err)

	outBalance := keepertest.BankKeeper.GetBalance(sdkCtx, swapUser, sample.TestDenom1)
	require.Equal(t, outBalance, sdk.NewCoin(sample.TestDenom1, sdk.NewInt(82)))

	swapPool, found = rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)
	require.Equal(t, swapPool.LpToken, newTotalLpToken)
	require.Equal(t, swapPool.BaseToken, token0.Add(addToken0).Add(swapUserToken0))
	require.Equal(t, swapPool.Token, token1.Add(addToken1).Sub(sdk.NewCoin(sample.TestDenom1, sdk.NewInt(82))))
}

func TestMsgServerSwapFail(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	provider := sample.OriginAccAddress()
	addToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(90))
	addToken1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(900))
	mintToProviderTokens := sdk.NewCoins(addToken0, addToken1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToProviderTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, mintToProviderTokens)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(10))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(100))
	lpDenom := types.GetLpTokenDenom(0)
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	swapUser := sample.OriginAccAddress()
	swapUserToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(10))
	mintToSwapUserTokens := sdk.NewCoins(swapUserToken0)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToSwapUserTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, swapUser, mintToSwapUserTokens)

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// crate pool
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}
	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     admin,
		UserAddress: provider.String(),
	}
	_, err = srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, provider))

	// add liquidity
	newTotalLpToken := sdk.NewCoin(lpDenom, sdk.NewInt(100))
	create2LpToken := sdk.NewCoin(lpDenom, sdk.NewInt(90))

	msgAddLiquidity := types.MsgAddLiquidity{
		Creator: provider.String(),
		Token0:  addToken0,
		Token1:  addToken1,
	}

	_, err = srv.AddLiquidity(ctx, &msgAddLiquidity)
	require.NoError(t, err)

	swapPool, found = rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, newTotalLpToken)
	require.Equal(t, swapPool.BaseToken, token0.Add(addToken0))
	require.Equal(t, swapPool.Token, token1.Add(addToken1))

	create2LpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, provider, lpDenom)
	require.Equal(t, create2LpToken, create2LpBalance)

	// swap
	testcases := []struct {
		name        string
		inputToken  sdk.Coin
		minOutToken sdk.Coin
		creator     string
	}{
		{
			name:        "swap pool not exist",
			inputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(10)),
			minOutToken: sdk.NewCoin(sample.TestDenom2, sdk.NewInt(80)),
			creator:     swapUser.String(),
		},
		{
			name:        "input token amount is zero",
			inputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(0)),
			minOutToken: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(80)),
			creator:     swapUser.String(),
		},
		{
			name:        "token0 neg amount",
			inputToken:  sdk.Coin{sample.TestDenom, sdk.NewInt(-10)},
			minOutToken: sdk.Coin{sample.TestDenom1, sdk.NewInt(10)},
			creator:     swapUser.String(),
		},
		{
			name:        "token1 neg amount",
			inputToken:  sdk.Coin{sample.TestDenom, sdk.NewInt(30)},
			minOutToken: sdk.Coin{sample.TestDenom1, sdk.NewInt(-10)},
			creator:     swapUser.String(),
		},
		{
			name:        "not enough input token balance",
			inputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(11)),
			minOutToken: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(10)),
			creator:     swapUser.String(),
		},
		{
			name:        "min out token to big",
			inputToken:  sdk.NewCoin(sample.TestDenom, sdk.NewInt(10)),
			minOutToken: sdk.NewCoin(sample.TestDenom1, sdk.NewInt(83)),
			creator:     swapUser.String(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			msgSwap := types.MsgSwap{
				Creator:     tc.creator,
				InputToken:  tc.inputToken,
				MinOutToken: tc.minOutToken,
			}
			err := msgSwap.ValidateBasic()
			if err != nil {
				t.Log(err)
				return
			}
			_, err = srv.Swap(ctx, &msgSwap)
			t.Log(err)
			require.Error(t, err)
		})
	}
}

func TestMsgServerRmLiquiditySuccess(t *testing.T) {
	srv, rdexKeeper, ctx, sdkCtx := setupMsgServer(t)

	provider := sample.OriginAccAddress()
	addToken0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8))
	addToken1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8))

	mintToProviderTokens := sdk.NewCoins(addToken0, addToken1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, mintToProviderTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, provider, mintToProviderTokens)

	token0 := sdk.NewCoin(sample.TestDenom, sdk.NewInt(500e8))
	token1 := sdk.NewCoin(sample.TestDenom1, sdk.NewInt(500e8))
	lpDenom := types.GetLpTokenDenom(0)
	willMintLpToken := sdk.NewCoin(lpDenom, token0.Amount)

	// add pool creator
	admin := sample.TestAdmin
	poolCreator := sample.OriginAccAddress()

	willMintTokens := sdk.NewCoins(token0, token1)
	keepertest.BankKeeper.MintCoins(sdkCtx, types.ModuleName, willMintTokens)
	keepertest.BankKeeper.SendCoinsFromModuleToAccount(sdkCtx, types.ModuleName, poolCreator, willMintTokens)

	msgPoolCreator := types.MsgAddPoolCreator{
		Creator:     admin,
		UserAddress: poolCreator.String(),
	}
	_, err := srv.AddPoolCreator(ctx, &msgPoolCreator)
	require.NoError(t, err)

	// crate pool
	msgCreatePool := types.MsgCreatePool{
		Creator: poolCreator.String(),
		Token0:  token0,
		Token1:  token1,
	}
	_, err = srv.CreatePool(ctx, &msgCreatePool)
	require.NoError(t, err)

	swapPool, found := rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, willMintLpToken)
	require.Equal(t, swapPool.BaseToken, token0)
	require.Equal(t, swapPool.Token, token1)

	lpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, poolCreator, lpDenom)
	require.Equal(t, lpBalance, swapPool.LpToken)

	// add provider
	msgAddProvider := types.MsgAddProvider{
		Creator:     admin,
		UserAddress: provider.String(),
	}
	_, err = srv.AddProvider(ctx, &msgAddProvider)
	require.NoError(t, err)
	require.True(t, rdexKeeper.HasProvider(sdkCtx, provider))

	// add liquidity
	newTotalLpToken := sdk.NewCoin(lpDenom, sdk.NewInt(1000e8))
	create2LpToken := sdk.NewCoin(lpDenom, sdk.NewInt(500e8))

	msgAddLiquidity := types.MsgAddLiquidity{
		Creator: provider.String(),
		Token0:  addToken0,
		Token1:  addToken1,
	}

	_, err = srv.AddLiquidity(ctx, &msgAddLiquidity)
	require.NoError(t, err)

	swapPool, found = rdexKeeper.GetSwapPool(sdkCtx, lpDenom)
	require.True(t, found)

	require.Equal(t, swapPool.LpToken, newTotalLpToken)
	require.Equal(t, swapPool.BaseToken, token0.Add(addToken0))
	require.Equal(t, swapPool.Token, token1.Add(addToken1))

	create2LpBalance := keepertest.BankKeeper.GetBalance(sdkCtx, provider, lpDenom)
	require.Equal(t, create2LpToken, create2LpBalance)

	// rm liqidity
	msgRmLp := types.MsgRemoveLiquidity{
		Creator:         provider.String(),
		RmUnit:          create2LpToken.Amount,
		SwapUnit:        sdk.ZeroInt(),
		MinOutToken0:    addToken1,
		MinOutToken1:    addToken0,
		InputTokenDenom: "",
	}

	_, err = srv.RemoveLiquidity(ctx, &msgRmLp)
	require.NoError(t, err)

}
