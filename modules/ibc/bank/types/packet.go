package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

var _ ibc.Packet = PacketMsgUser{}

var _ ibc.Packet = PacketTokenTransfer{}

type PacketMsgUser struct {
	Name string
}

func (p PacketMsgUser) Marshal() []byte {
	return ModuleCdc.MustMarshalBinaryBare(p)
}

func (p PacketMsgUser) SenderPort() string {
	return "send"
}

func (p PacketMsgUser) ReceiverPort() string {
	return "receive"
}

func (p PacketMsgUser) Type() string {
	return "packet-user"
}

func (p PacketMsgUser) ValidateBasic() sdk.Error {
	return nil
}

func (p PacketMsgUser) Timeout() uint64 {
	return 0
}

func (p PacketMsgUser) MarshalJSON() ([]byte, error) {
	bz, err := json.Marshal(p.Name)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

type PacketTokenTransfer struct {
	Sender   sdk.AccAddress
	Receiver sdk.AccAddress
	Amount   sdk.Coins
}

func (p PacketTokenTransfer) SenderPort() string {
	return "send"
}

func (p PacketTokenTransfer) ReceiverPort() string {
	return "receive"
}

func (p PacketTokenTransfer) Type() string {
	return "token-transfer"
}

func (p PacketTokenTransfer) ValidateBasic() sdk.Error {
	return nil
}

func (p PacketTokenTransfer) Timeout() uint64 {
	return 0
}

func (p PacketTokenTransfer) Marshal() []byte {
	return ModuleCdc.MustMarshalBinaryBare(p)
}

func (p PacketTokenTransfer) MarshalJSON() ([]byte, error) {

	var a = struct {
		Sender   sdk.AccAddress
		Receiver sdk.AccAddress
		Amount   sdk.Coins
	}{
		Sender:   p.Sender,
		Receiver: p.Receiver,
		Amount:   p.Amount,
	}

	return json.Marshal(a)
}
