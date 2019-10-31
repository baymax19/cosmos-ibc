package types

import "github.com/cosmos/cosmos-sdk/types"

type BankKeeper interface {
	AddCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) (types.Coins, types.Error)
	SubtractCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) (types.Coins, types.Error)
	HasCoins(ctx types.Context, addr types.AccAddress, amt types.Coins) bool
}
