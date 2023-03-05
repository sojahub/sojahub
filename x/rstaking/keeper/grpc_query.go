package keeper

import (
	"github.com/sojahub/sojahub/x/rstaking/types"
)

var _ types.QueryServer = Keeper{}
