package orders

import (
	internal "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"reflect"
	"sort"
)

type Keeper struct {
	cdc        *codec.Codec
	storeKey   types.StoreKey
	bankKeeper internal.BankKeeper
	port       ibc.Port
}

func NewKeeper(cdc *codec.Codec, storeKey types.StoreKey, port ibc.Port, bankKeeper internal.BankKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		port:       port,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) SetOrder(ctx types.Context, order internal.BaseMakeOrder) {
	store := ctx.KVStore(k.storeKey)
	store.Set(internal.GetMakeOrderKey(order.OrderHash), k.cdc.MustMarshalBinaryBare(order))
}

func (k Keeper) UpdateOrder(ctx types.Context, order internal.BaseMakeOrder) types.Error {
	baseMakeOrder := k.GetOrder(ctx, order.OrderHash)
	if !reflect.DeepEqual(baseMakeOrder, internal.BaseMakeOrder{}) {
		return types.NewError("orders", 19, "order already exist")
	}

	k.SetOrder(ctx, order)
	return nil
}

func (k Keeper) GetOrder(ctx types.Context, orderHash string) internal.BaseMakeOrder {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(internal.GetMakeOrderKey(orderHash))
	if bz == nil {
		return internal.BaseMakeOrder{}
	}

	var baseMakeOrder internal.BaseMakeOrder
	k.cdc.MustUnmarshalBinaryBare(bz, &baseMakeOrder)

	return baseMakeOrder
}

func (k Keeper) GetOrders(ctx types.Context) []internal.BaseMakeOrder {
	var orders []internal.BaseMakeOrder

	store := ctx.KVStore(k.storeKey)

	iterator := types.KVStorePrefixIterator(store, internal.MakeOrderKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var order internal.BaseMakeOrder
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &order)

		orders = append(orders, order)
	}
	return orders
}

func (k Keeper) GetOrderBook(ctx types.Context) []internal.BaseMakeOrder {
	var activeOrders []internal.BaseMakeOrder
	orders := k.GetOrders(ctx)

	for _, order := range orders {
		if order.OrderStatus == internal.StatusUnFilled {
			activeOrders = append(activeOrders, order)
		}
	}

	return activeOrders
}

func (k Keeper) SetTakeOrder(ctx types.Context, takeOrder []internal.BaseTakeOrder) {
	store := ctx.KVStore(k.storeKey)
	store.Set(internal.GetTakeOrderKey(takeOrder[0].OrderHash), k.cdc.MustMarshalBinaryBare(takeOrder))
}

func (k Keeper) GetTakeOrdersForMakeOrder(ctx types.Context, orderHash string) []internal.BaseTakeOrder {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(internal.GetTakeOrderKey(orderHash))
	if bz == nil {
		return nil
	}

	var orders []internal.BaseTakeOrder
	k.cdc.MustUnmarshalBinaryBare(bz, &orders)

	return orders
}

func (k Keeper) UpdateTakeOrder(ctx types.Context, takeOrder internal.BaseTakeOrder, channelID string) types.Error {
	// considering intra chain transaction
	if channelID == "" {
		orders := k.GetTakeOrdersForMakeOrder(ctx, takeOrder.OrderHash)
		if orders == nil {
			k.SetTakeOrder(ctx, []internal.BaseTakeOrder{takeOrder})
		}

		orders = append(orders, takeOrder)

		sort.Slice(orders, func(i, j int) bool {
			return orders[i].TakerFillAmount.Amount.GT(orders[j].TakerFillAmount.Amount)
		})

		k.SetTakeOrder(ctx, orders)
	} else {
		// inter chain transaction

		packet := internal.PacketTakeOrder{
			TakerFillAmount: takeOrder.TakerFillAmount,
			TakerAddress:    takeOrder.TakerAddress,
			OrderHash:       takeOrder.OrderHash,
		}

		return k.port.Send(ctx, channelID, packet)
	}
	return nil
}

func (k Keeper) ExchangeTokens(ctx types.Context, orderHash, channelID string) types.Error {
	makeOrder := k.GetOrder(ctx, orderHash)
	takeOrders := k.GetTakeOrdersForMakeOrder(ctx, orderHash)

	// TODO handle floating value; decimal implementation; selection of take order from list ;
	takeOrder := takeOrders[0] // Based on highest TakerFillAmount
	takerGetAmount := makeOrder.MakerAsset.Amount.Mul(takeOrder.TakerFillAmount.Amount).Quo(makeOrder.TakerAsset.Amount)

	if !k.bankKeeper.HasCoins(ctx, makeOrder.MakerAddress, types.Coins{types.Coin{makeOrder.MakerAsset.Denom, takerGetAmount}}) {
		return types.NewError("orders", 19, "insufficient funds in maker ")
	}

	_, err := k.bankKeeper.SubtractCoins(ctx, makeOrder.MakerAddress, types.Coins{types.Coin{makeOrder.MakerAsset.Denom, takerGetAmount}})
	if err != nil {
		return err
	}

	_, err = k.bankKeeper.AddCoins(ctx, makeOrder.MakerAddress, types.Coins{takeOrder.TakerFillAmount})
	if err != nil {
		return err
	}

	makeOrder.TakerAddress = takeOrder.TakerAddress
	makeOrder.OrderStatus = internal.StatusFilled
	k.SetOrder(ctx, makeOrder)

	if channelID != "" {
		packet := internal.PacketExchangeOrder{
			TakerGetAmount:    types.Coin{makeOrder.MakerAsset.Denom, takerGetAmount},
			TakerDeductAmount: takeOrder.TakerFillAmount,
			TakerAddress:      takeOrder.TakerAddress,
			OrderHash:         makeOrder.OrderHash,
		}
		return k.port.Send(ctx, channelID, packet)
	}

	// intra TakerAddress
	if !k.bankKeeper.HasCoins(ctx, takeOrder.TakerAddress, types.Coins{takeOrder.TakerFillAmount}) {
		return types.NewError("orders", 19, "insufficient funds in taker ")
	}

	_, err = k.bankKeeper.SubtractCoins(ctx, takeOrder.TakerAddress, types.Coins{takeOrder.TakerFillAmount})
	if err != nil {
		return err
	}

	_, err = k.bankKeeper.AddCoins(ctx, takeOrder.TakerAddress, types.Coins{types.Coin{makeOrder.MakerAsset.Denom, takerGetAmount}})
	if err != nil {
		return err
	}

	return nil
}
