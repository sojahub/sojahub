package keeper

import (
	"github.com/sojahub/sojahub/x/mining/types"
)

var _ types.QueryServer = Keeper{}
