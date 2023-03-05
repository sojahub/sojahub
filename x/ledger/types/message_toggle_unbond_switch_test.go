package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgToggleUnbondSwitch_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgToggleUnbondSwitch
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgToggleUnbondSwitch{
				Creator: "invalid_address",
				Denom:   sample.TestDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgToggleUnbondSwitch{
				Creator: sample.AccAddress(),
				Denom:   sample.TestDenom,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
