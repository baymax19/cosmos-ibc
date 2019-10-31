package ibcrecv

import (
	types2 "github.com/baymax19/cosmos-ibc/modules/ibc/bank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case ibc.MsgPacket:
			switch packet := msg.Packet.(type) {
			case types2.PacketMsgUser:
				return handleMyPacket(ctx, k, packet, msg.ChannelID)
			case types2.PacketTokenTransfer:
				return handleTokenTransfer(ctx, k, packet, msg.ChannelID)

			default:
				return sdk.ErrUnknownRequest("19191").Result()
			}
		default:
			return sdk.ErrUnknownRequest("1919").Result()
		}
	}
}

func handleMyPacket(ctx sdk.Context, k Keeper, packet types2.PacketMsgUser, chainID string) sdk.Result {
	err := k.UpdateUser(ctx, chainID, packet.Name)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleTokenTransfer(ctx sdk.Context, k Keeper, packet types2.PacketTokenTransfer, chanID string) sdk.Result {
	err := k.ReceiveTokens(ctx, packet.Receiver, packet.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}
