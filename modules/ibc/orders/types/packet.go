package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

var _ ibc.Packet = PacketTakeOrder{}

type PacketTakeOrder struct {
	TakerFillAmount types.Coin
	TakerAddress    types.AccAddress
	OrderHash       string
}

func (p PacketTakeOrder) SenderPort() string {
	return "orders"
}

func (p PacketTakeOrder) ReceiverPort() string {
	return "ordersrecv"
}

func (p PacketTakeOrder) Type() string {
	return "packet-take-order"
}

func (p PacketTakeOrder) ValidateBasic() types.Error {
	return nil
}

func (p PacketTakeOrder) Timeout() uint64 {
	return 0
}

func (p PacketTakeOrder) Marshal() []byte {
	return ModuleCdc.MustMarshalBinaryBare(p)
}

func (p PacketTakeOrder) MarshalJSON() ([]byte, error) {

	var pkt = struct {
		TakerFillAmount types.Coin
		TakerAddress    types.AccAddress
		OrderHash       string
	}{
		TakerFillAmount: p.TakerFillAmount,
		TakerAddress:    p.TakerAddress,
		OrderHash:       p.OrderHash,
	}

	return json.Marshal(pkt)
}

func (p *PacketTakeOrder) UnmarshalJSON(bz []byte) (err error) {
	return json.Unmarshal(bz, &p)
}

type PacketExchangeOrder struct {
	TakerGetAmount    types.Coin
	TakerDeductAmount types.Coin
	TakerAddress      types.AccAddress
	OrderHash         string
}

func (p PacketExchangeOrder) SenderPort() string {
	return "orders"
}

func (p PacketExchangeOrder) ReceiverPort() string {
	return "ordersrecv"
}

func (p PacketExchangeOrder) Type() string {
	return "packet-exchange-order"
}

func (p PacketExchangeOrder) ValidateBasic() types.Error {
	return nil
}

func (p PacketExchangeOrder) Timeout() uint64 {
	return 0
}

func (p PacketExchangeOrder) Marshal() []byte {
	return ModuleCdc.MustMarshalBinaryBare(p)
}

func (p PacketExchangeOrder) MarshalJSON() ([]byte, error) {
	var pkt = struct {
		TakerGetAmount    types.Coin
		TakerDeductAmount types.Coin
		TakerAddress      types.AccAddress
		OrderHash         string
	}{
		TakerGetAmount:    p.TakerGetAmount,
		TakerDeductAmount: p.TakerDeductAmount,
		TakerAddress:      p.TakerAddress,
		OrderHash:         p.OrderHash,
	}

	return json.Marshal(pkt)
}

var _ ibc.Packet = PacketExchangeOrder{}
