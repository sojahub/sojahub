package keeper

import (
	"github.com/sojahub/sojahub/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
