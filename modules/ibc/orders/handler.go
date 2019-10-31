package orders

import (
	"fmt"
	internal "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
)

func NewHandler(k Keeper) types.Handler {
	return func(ctx types.Context, msg types.Msg) types.Result {
		ctx = ctx.WithEventManager(types.NewEventManager())

		switch msg := msg.(type) {
		case internal.MsgCreateMakeOrder:
			return handleCreateMakeOrder(ctx, k, msg)
		case internal.MsgTakerFillOrder:
			return handleTakerFillOrder(ctx, k, msg)
		case internal.MsgConfirmMakeOrder:
			return handleConfirmMakeOrder(ctx, k, msg)

		default:
			return types.ErrUnknownRequest("21345").Result()
		}

	}
}

func handleCreateMakeOrder(ctx types.Context, k Keeper, msg internal.MsgCreateMakeOrder) types.Result {
	baseMakeOrder := internal.BaseMakeOrder{
		MakerAddress: msg.MakerAddress,
		TakerAddress: nil,
		MakerAsset:   msg.MakerAsset,
		TakerAsset:   msg.TakerAsset,
		OrderHash:    msg.OrderHash,
		OrderStatus:  msg.OrderStatus,
	}

	err := k.UpdateOrder(ctx, baseMakeOrder)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(types.Events{
		types.NewEvent(internal.EventTypeMakeOrder,
			types.NewAttribute(internal.AttributeOrderHash, string(msg.OrderHash))),
	})

	return types.Result{
		Events: ctx.EventManager().Events(),
	}
}

func handleTakerFillOrder(ctx types.Context, k Keeper, msg internal.MsgTakerFillOrder) types.Result {
	baseTakeOrder := internal.BaseTakeOrder{
		TakerFillAmount: msg.TakerFillAmount,
		TakerAddress:    msg.TakerAddress,
		OrderHash:       msg.OrderHash,
	}
	err := k.UpdateTakeOrder(ctx, baseTakeOrder, msg.ChannelID)
	if err != nil {
		return err.Result()
	}

	return types.Result{}
}

func handleConfirmMakeOrder(ctx types.Context, k Keeper, msg internal.MsgConfirmMakeOrder) types.Result {
	makeOrder := k.GetOrder(ctx, msg.OrderHash)
	takeOrders := k.GetTakeOrdersForMakeOrder(ctx, msg.OrderHash)

	if len(takeOrders) < 1 {
		return types.NewError("orders", 19, fmt.Sprintf("takeorders doesn't exist with given %s orderhash", msg.OrderHash)).Result()
	}

	if reflect.DeepEqual(makeOrder, internal.MsgConfirmMakeOrder{}) {
		return types.NewError("orders", 19, "order doesn't exist").Result()
	}

	if !makeOrder.MakerAddress.Equals(msg.FromAddress) {
		return types.NewError("orders", 19, "order not associate with signer").Result()
	}

	err := k.ExchangeTokens(ctx, msg.OrderHash, msg.ChannelID)
	if err != nil {
		return err.Result()
	}

	return types.Result{}
}
