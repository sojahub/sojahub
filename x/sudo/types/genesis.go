package types

import (
// this line is used by starport scaffolding # genesis/types/import
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Admin: "fis1xrhhus43kjhqdccee7aqnuukh8ugv09affsg84",
		Denoms: []string{},
	    // this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
    // this line is used by starport scaffolding # genesis/types/validate
	_, err := sdk.AccAddressFromBech32(gs.Admin)
	if err != nil {
		return fmt.Errorf("invalid admin address %s", gs.Admin)
	}

	for _, dnm := range gs.Denoms {
		if sdk.ValidateDenom(dnm) != nil {
			return fmt.Errorf("invalid denom %s", dnm)
		}
	}

	return nil
}