package types

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(PacketMsgUser{}, "ibcrecv/packetuser", nil)
	cdc.RegisterConcrete(MsgUser{}, "ibcsend/msguser", nil)
}

func RegisterSend(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgUser{}, "ibcsend/msguser", nil)
}

func RegisterRecv(cdc *codec.Codec) {
	cdc.RegisterConcrete(PacketMsgUser{}, "ibcrecv/packetuser", nil)
}
