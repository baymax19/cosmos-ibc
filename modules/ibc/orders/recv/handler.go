package recv

import (
	internal "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

func NewHandler(k Keeper) types.Handler {
	return func(ctx types.Context, msg types.Msg) types.Result {
		switch msg := msg.(type) {
		case ibc.MsgPacket:
			switch packet := msg.Packet.(type) {
			case internal.PacketTakeOrder:
				return handlePacketTakeOrder(ctx, k, packet, msg.ChannelID)
			case internal.PacketExchangeOrder:
				return handlePacketExchangeOrder(ctx, k, packet, msg.ChannelID)
			default:
				return types.ErrUnknownRequest("23331345").Result()
			}
		default:
			return types.ErrUnknownRequest("21345").Result()
		}
	}
}

func handlePacketTakeOrder(ctx types.Context, k Keeper, packet internal.PacketTakeOrder, channelID string) types.Result {
	takeOrder := internal.BaseTakeOrder{
		TakerFillAmount: packet.TakerFillAmount,
		TakerAddress:    packet.TakerAddress,
		OrderHash:       packet.OrderHash,
	}

	err := k.UpdateTakeOrder(ctx, takeOrder, channelID)
	if err != nil {
		return err.Result()
	}

	return types.Result{}
}

func handlePacketExchangeOrder(ctx types.Context, k Keeper, packet internal.PacketExchangeOrder, channelID string) types.Result {
	err := k.ExchangeTokens(ctx, packet)
	if err != nil {
		return err.Result()
	}

	return types.Result{}
}
