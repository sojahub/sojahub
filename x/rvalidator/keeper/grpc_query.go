package keeper

import (
	"github.com/sojahub/sojahub/x/rvalidator/types"
)

var _ types.QueryServer = Keeper{}
