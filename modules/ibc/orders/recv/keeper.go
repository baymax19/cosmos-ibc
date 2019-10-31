package recv

import (
	internal "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"

	"sort"
)

type Keeper struct {
	cdc        *codec.Codec
	storeKey   types.StoreKey
	sendKeeper OrderSenderKeeper
	bankKeeper BankKeeper
	port       ibc.Port
}

func NewKeeper(cdc *codec.Codec, key types.StoreKey, port ibc.Port, sendKeeper OrderSenderKeeper, bankKeeper BankKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		sendKeeper: sendKeeper,
		bankKeeper: bankKeeper,
		port:       port,
	}
}

func (k Keeper) UpdateTakeOrder(ctx types.Context, takeOrder internal.BaseTakeOrder, channelID string) types.Error {
	orders := k.sendKeeper.GetTakeOrdersForMakeOrder(ctx, takeOrder.OrderHash)

	orders = append(orders, takeOrder)

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].TakerFillAmount.Amount.GT(orders[j].TakerFillAmount.Amount)
	})

	err := k.sendKeeper.UpdateTakeOrder(ctx, takeOrder, "")
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) ExchangeTokens(ctx types.Context, packet internal.PacketExchangeOrder) types.Error {

	if !k.bankKeeper.HasCoins(ctx, packet.TakerAddress, types.Coins{packet.TakerDeductAmount}) {
		return types.NewError("orders", 19, "insufficient funds  ")
	}

	_, err := k.bankKeeper.SubtractCoins(ctx, packet.TakerAddress, types.Coins{packet.TakerDeductAmount})
	if err != nil {
		return err
	}

	_, err = k.bankKeeper.AddCoins(ctx, packet.TakerAddress, types.Coins{packet.TakerGetAmount})
	if err != nil {
		return err
	}

	return nil
}
