package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/sojahub/sojahub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgStake_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgStake
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgStake{
				Creator:     "invalid_address",
				StakeAmount: sdk.NewInt(1),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgStake{
				Creator:     sample.AccAddress(),
				StakeAmount: sdk.NewInt(1),
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
