package types

import (
	// this line is used by starport scaffolding # genesis/types/import
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sojahub/sojahub/utils"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                   DefaultParams(),
		InflationBase:            sdk.NewIntFromUint64(1000000000000),
		CoinToBeBurned:           sdk.Coin{Denom: utils.FuryDenom, Amount: sdk.NewIntFromUint64(1000000000000)},
		ValidatorWhitelist:       []string{},
		ValidatorWhitelistSwitch: false,
		DelegatorWhitelist:       []string{},
		DelegatorWhitelistSwitch: false,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
