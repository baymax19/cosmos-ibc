package types

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(PacketMsgUser{}, "ibcrecv/packet-user", nil)
	cdc.RegisterConcrete(PacketTokenTransfer{}, "ibcrecv/packet-token-transfer", nil)
	cdc.RegisterConcrete(MsgUser{}, "ibcsend/msg-user", nil)
	cdc.RegisterConcrete(MsgTokenTransfer{}, "ibcsend/msg-token-transfer", nil)
}

func RegisterSend(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgUser{}, "ibcsend/msg-user", nil)
	cdc.RegisterConcrete(MsgTokenTransfer{}, "ibcsend/msg-token-transfer", nil)
}

func RegisterRecv(cdc *codec.Codec) {
	cdc.RegisterConcrete(PacketMsgUser{}, "ibcrecv/packet-user", nil)
	cdc.RegisterConcrete(PacketTokenTransfer{}, "ibcrecv/packet-token-transfer", nil)
}
