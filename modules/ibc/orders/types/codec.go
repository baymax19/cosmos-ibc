package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateMakeOrder{}, "ibc/orders/create-mak-order", nil)
	cdc.RegisterConcrete(MsgTakerFillOrder{}, "ibc/orders/taker-fill-order", nil)
	cdc.RegisterConcrete(MsgConfirmMakeOrder{}, "ibc/orders/confirm-make-orders", nil)
	cdc.RegisterConcrete(PacketTakeOrder{}, "orders/packet-take-order", nil)
	cdc.RegisterConcrete(PacketExchangeOrder{}, "orders/packet-exchange-order", nil)
}

func RegisterSend(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateMakeOrder{}, "ibc/orders/create-make-order", nil)
	cdc.RegisterConcrete(MsgTakerFillOrder{}, "ibc/orders/taker-fill-order", nil)
	cdc.RegisterConcrete(MsgConfirmMakeOrder{}, "ibc/orders/confirm-make-orders", nil)

}

func RegisterRecv(cdc *codec.Codec) {
	cdc.RegisterConcrete(PacketTakeOrder{}, "orders/packet-take-order", nil)
	cdc.RegisterConcrete(PacketExchangeOrder{}, "orders/packet-exchange-order", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
}
