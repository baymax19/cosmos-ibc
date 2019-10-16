package types

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

var _ ibc.Packet = PacketMsgUser{}

type PacketMsgUser struct {
	Name string
}

func (p PacketMsgUser) Marshal() []byte {
	cdc := codec.New()
	RegisterCodec(cdc)
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

	fmt.Println("-------------------------------------------", p)
	bz, err := json.Marshal(p.Name)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
