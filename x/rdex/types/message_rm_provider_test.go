package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/sojahub/sojahub/x/rdex/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRmProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRmProvider
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRmProvider{
				Creator:     "invalid_address",
				UserAddress: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRmProvider{
				Creator:     sample.AccAddress(),
				UserAddress: sample.AccAddress(),
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
