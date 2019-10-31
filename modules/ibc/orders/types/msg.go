package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

const Route = "orders"

type MsgCreateMakeOrder struct {
	MakerAddress types.AccAddress
	TakerAddress types.AccAddress

	MakerAsset types.Coin
	TakerAsset types.Coin

	OrderHash   string
	OrderStatus OrderStatus
}

func NewMsgCreateMakeOrder(makerAddress types.AccAddress, makerAsset, takerAsset types.Coin,
	orderHash string, status OrderStatus) MsgCreateMakeOrder {
	return MsgCreateMakeOrder{
		MakerAddress: makerAddress,
		TakerAddress: nil,
		MakerAsset:   makerAsset,
		TakerAsset:   takerAsset,
		OrderHash:    orderHash,
		OrderStatus:  status,
	}
}

func (m MsgCreateMakeOrder) Route() string {
	return Route
}

func (m MsgCreateMakeOrder) Type() string {
	return "create-make-order"
}

func (m MsgCreateMakeOrder) ValidateBasic() types.Error {
	return nil
}

func (m MsgCreateMakeOrder) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgCreateMakeOrder) GetSigners() []types.AccAddress {
	return []types.AccAddress{m.MakerAddress}
}

var _ types.Msg = MsgCreateMakeOrder{}

type MsgTakerFillOrder struct {
	TakerFillAmount types.Coin
	TakerAddress    types.AccAddress
	OrderHash       string
	ChannelID       string
}

func NewMsgTakerFillAmount(fillAmount types.Coin, takerAddress types.AccAddress, hash, channel string) MsgTakerFillOrder {
	return MsgTakerFillOrder{
		TakerFillAmount: fillAmount,
		TakerAddress:    takerAddress,
		OrderHash:       hash,
		ChannelID:       channel,
	}
}
func (m MsgTakerFillOrder) Route() string {
	return Route
}

func (m MsgTakerFillOrder) Type() string {
	return "taker-fill-order"
}

func (m MsgTakerFillOrder) ValidateBasic() types.Error {
	return nil
}

func (m MsgTakerFillOrder) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgTakerFillOrder) GetSigners() []types.AccAddress {
	return []types.AccAddress{m.TakerAddress}
}

var _ types.Msg = MsgTakerFillOrder{}

type MsgConfirmMakeOrder struct {
	FromAddress types.AccAddress
	OrderHash   string
	ChannelID   string
}

func NewMsgConfirmMakeOrder(from types.AccAddress, hash, chanID string) MsgConfirmMakeOrder {
	return MsgConfirmMakeOrder{
		FromAddress: from,
		OrderHash:   hash,
		ChannelID:   chanID,
	}
}

func (m MsgConfirmMakeOrder) Route() string {
	return Route
}

func (m MsgConfirmMakeOrder) Type() string {
	return "confirm-make-order"
}

func (m MsgConfirmMakeOrder) ValidateBasic() types.Error {
	return nil
}

func (m MsgConfirmMakeOrder) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgConfirmMakeOrder) GetSigners() []types.AccAddress {
	return []types.AccAddress{m.FromAddress}
}

var _ types.Msg = MsgConfirmMakeOrder{}
