package bank

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/sojahub/sojahub/utils"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the staking module.
type AppModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultGenesisState := types.DefaultGenesisState()
	// {
	//     "description": "The native staking token of the SoJa Hub",
	//     "denom_units": [
	//         {
	//             "denom": "ufury",
	//             "exponent": 0,
	//             "aliases": [
	//                 "microfury"
	//             ]
	//         },
	//         {
	//             "denom": "mfury",
	//             "exponent": 3,
	//             "aliases": [
	//               "millifury"
	//             ]
	//         },
	//         {
	//             "denom": "fury",
	//             "exponent": 6
	//         }
	//     ],
	//     "base": "ufury",
	//     "display": "fury",
	//     "name": "FURY",
	//     "symbol": "FURY"
	// }
	defaultGenesisState.DenomMetadata = append(defaultGenesisState.DenomMetadata,
		types.Metadata{
			Description: "The native staking token of the SoJa Hub",
			DenomUnits: []*types.DenomUnit{
				{
					Denom:    utils.FuryDenom,
					Exponent: 0,
					Aliases:  []string{"microfury"},
				},
				{
					Denom:    "mfury",
					Exponent: 3,
					Aliases:  []string{"millifury"},
				},
				{
					Denom:    "fury",
					Exponent: 6,
				},
			},
			Base:    utils.FuryDenom,
			Display: "fury",
			Name:    "FURY",
			Symbol:  "FURY",
		},
	)
	return cdc.MustMarshalJSON(defaultGenesisState)
}
