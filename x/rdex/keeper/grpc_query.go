package keeper

import (
	"github.com/sojahub/sojahub/x/rdex/types"
)

var _ types.QueryServer = Keeper{}
