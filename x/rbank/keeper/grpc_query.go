package keeper

import (
	"github.com/sojahub/sojahub/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
