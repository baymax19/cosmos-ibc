package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var cdc = codec.New()

type MsgUser struct {
	Name      string
	ChannelID string
	Signer    sdk.AccAddress
}

var _ sdk.Msg = MsgUser{}

func NewMsgUser(name, chanID string, from sdk.AccAddress) MsgUser {
	return MsgUser{
		Name:      name,
		ChannelID: chanID,
		Signer:    from,
	}
}

func (msg MsgUser) ValidateBasic() sdk.Error { return nil }

func (msg MsgUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(cdc.MustMarshalJSON(msg))
}

func (MsgUser) Route() string {
	return "send"
}

func (MsgUser) Type() string {
	return "user"
}

func (msg MsgUser) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
