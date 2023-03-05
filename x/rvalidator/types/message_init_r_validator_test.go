package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sojahub/sojahub/testutil/sample"
	"github.com/sojahub/sojahub/x/rvalidator/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddRValidator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgInitRValidator
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgInitRValidator{
				Creator:        "invalid_address",
				Denom:          sample.AccAddress(),
				ValAddressList: []string{"cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgInitRValidator{
				Creator:        sample.AccAddress(),
				Denom:          sample.AccAddress(),
				ValAddressList: []string{"cosmosvaloper1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u2lcnj0"},
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
