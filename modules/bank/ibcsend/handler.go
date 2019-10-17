package ibcsend

import (
	"github.com/baymax19/cosmos-ibc/modules/bank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgUser:
			return handleMsgUser(ctx, k, msg)
		case types.MsgTokenTransfer:
			return handleTokenTransfer(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("1919").Result()
		}
	}
}

func handleMsgUser(ctx sdk.Context, k Keeper, msg types.MsgUser) ( sdk.Result) {
	err := k.UpdateUser(ctx, msg.ChannelID, msg.Name)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleTokenTransfer(ctx sdk.Context, k Keeper, msg types.MsgTokenTransfer) sdk.Result {
	err := k.TransferTokens(ctx, msg.Signer, msg.ToAddress, msg.Amount, msg.ChannelID)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}
