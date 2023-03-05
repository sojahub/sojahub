package keeper

import (
	"github.com/sojahub/sojahub/x/claim/types"
)

var _ types.QueryServer = Keeper{}
