package recv

import (
	internal "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/types"
)

type OrderSenderKeeper interface {
	UpdateTakeOrder(ctx types.Context, takeOrder internal.BaseTakeOrder, channelID string) types.Error
	GetTakeOrdersForMakeOrder(ctx types.Context, orderHash string) []internal.BaseTakeOrder
}

type BankKeeper interface {
	AddCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) (types.Coins, types.Error)
	SubtractCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) (types.Coins, types.Error)
	HasCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) bool
}
