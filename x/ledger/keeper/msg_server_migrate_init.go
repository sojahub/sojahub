package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	xBridgeTypes "github.com/sojahub/sojahub/x/bridge/types"
	"github.com/sojahub/sojahub/x/ledger/types"
	sudotypes "github.com/sojahub/sojahub/x/sudo/types"
)

// Notice:
// - ensure pool is bonded before migrateInit
// - it will replace pre value with latest vaue if you call migrateInit multi times
func (k msgServer) MigrateInit(goCtx context.Context, msg *types.MsgMigrateInit) (*types.MsgMigrateInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if k.Keeper.MigrateInitIsSealed(ctx) {
		return nil, types.ErrMigrateInitIsSealed
	}

	err := k.CheckAddress(ctx, msg.Denom, msg.GetPool())
	if err != nil {
		return nil, err
	}

	if !k.Keeper.IsBondedPoolExist(ctx, msg.Denom, msg.Pool) {
		return nil, types.ErrPoolNotBonded
	}

	//init exchange rate
	k.Keeper.MigrateExchangeRate(ctx, msg.Denom, msg.ExchangeRate)

	//init pipeline
	pipeline, found := k.GetBondPipeline(ctx, msg.Denom, msg.Pool)
	if !found {
		pipeline = types.BondPipeline{
			Denom: msg.Denom,
			Pool:  msg.Pool,
		}
	}
	pipeline.Chunk.Active = msg.Active
	pipeline.Chunk.Bond = msg.Bond
	pipeline.Chunk.Unbond = msg.Unbond
	k.Keeper.SetBondPipeline(ctx, pipeline)

	//init supply
	shouldMintCoins := sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.TotalSupply))
	moduleAddress := authTypes.NewModuleAddress(xBridgeTypes.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, msg.Denom)
	if balance.Amount.IsPositive() {
		if err := k.bankKeeper.BurnCoins(ctx, xBridgeTypes.ModuleName, sdk.NewCoins(balance)); err != nil {
			return nil, err
		}
	}
	if err := k.bankKeeper.MintCoins(ctx, xBridgeTypes.ModuleName, shouldMintCoins); err != nil {
		return nil, err
	}

	// init total protocol fee
	k.Keeper.SetTotalProtocolFee(ctx, msg.Denom, msg.TotalProtocolFee)

	// seal
	k.Keeper.SealMigrateInit(ctx)

	return &types.MsgMigrateInitResponse{}, nil
}
