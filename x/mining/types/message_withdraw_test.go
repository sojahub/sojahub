package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/sojahub/sojahub/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgWithdraw_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgWithdraw
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgWithdraw{
				Creator:        "invalid_address",
				WithdrawAmount: sdk.NewInt(10),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgWithdraw{
				Creator:        sample.AccAddress(),
				WithdrawAmount: sdk.NewInt(10),
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
