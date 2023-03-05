package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSetEraSeconds_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetEraSeconds
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetEraSeconds{
				Creator:    "invalid_address",
				EraSeconds: 2000,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetEraSeconds{
				Creator:    sample.AccAddress(),
				EraSeconds: 3000,
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
