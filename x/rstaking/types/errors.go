package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rstaking module sentinel errors
var (
	ErrValAlreadyInWhitelist       = sdkerrors.Register(ModuleName, 1100, "validator already in whitelist error")
	ErrInsufficientFunds           = sdkerrors.Register(ModuleName, 1101, "insufficient funds error")
	ErrDelegatorAlreadyInWhitelist = sdkerrors.Register(ModuleName, 1102, "delegator already in whitelist error")
	ErrDenomUnmatch                = sdkerrors.Register(ModuleName, 1103, "denom unmatch")
	ErrDelegatorNotInWhitelist     = sdkerrors.Register(ModuleName, 1104, "delegator not in whitelist error")
)
