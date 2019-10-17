package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

var _ ibc.Packet = PacketMsgUser{}

var _ ibc.Packet = PacketTokenTransfer{}

type PacketMsgUser struct {
	Name string
}

func (p PacketMsgUser) Marshal() []byte {

	return cdc.MustMarshalBinaryBare(p)
}

func (p PacketMsgUser) SenderPort() string {
	return "send"
}

func (p PacketMsgUser) ReceiverPort() string {
	return "receive"
}

func (p PacketMsgUser) Type() string {
	return "user"
}

func (p PacketMsgUser) ValidateBasic() sdk.Error {
	return nil
}

func (p PacketMsgUser) Timeout() uint64 {
	return 0
}

func (p PacketMsgUser) MarshalJSON() ([]byte, error) {
	bz, err := json.Marshal(p)
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
	cdc := codec.New()
	RegisterCodec(cdc)
	return cdc.MustMarshalBinaryBare(p)
}

func (p PacketTokenTransfer) MarshalJSON() ([]byte, error) {
	bz, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return bz, nil
}
