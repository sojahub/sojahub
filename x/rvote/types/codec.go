package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Content)(nil), nil)
	cdc.RegisterConcrete(&MsgSetProposalLife{}, "rvote/SetProposalLife", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetProposalLife{},
		&MsgSubmitProposal{},
	)

	registry.RegisterInterface(
		"sojahub.rvote.Content",
		(*Content)(nil),
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterProposalTypeCodec(o interface{}, name string) {
	Amino.RegisterConcrete(o, name, nil)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	cryptocodec.RegisterCrypto(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
