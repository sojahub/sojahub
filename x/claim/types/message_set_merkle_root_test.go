package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgSetMerkleRoot_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetMerkleRoot
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetMerkleRoot{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetMerkleRoot{
				Creator: sample.AccAddress(),
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
