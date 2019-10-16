package ibcsend

import (
	"fmt"
	"github.com/baymax19/cosmos-ibc/modules/bank/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"strings"
)

type Keeper struct {
	cdc  *codec.Codec
	key  sdk.StoreKey
	port ibc.Port
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, port ibc.Port) Keeper {
	return Keeper{
		cdc:  cdc,
		key:  key,
		port: port,
	}
}

func (k Keeper) GetUser(ctx sdk.Context, chainID string) (string, sdk.Error) {
	store := ctx.KVStore(k.key)

	var res string
	bz := store.Get(types.UserKey(chainID))
	if bz == nil {
		return "", sdk.NewError("ibcsend", 1919, fmt.Sprintf("data is not existed with %s this id", chainID))
	}

	k.cdc.MustUnmarshalBinaryBare(bz, &res)
	return res, nil
}

func (k Keeper) SetUser(ctx sdk.Context, chanID, name string) {
	store := ctx.KVStore(k.key)
	store.Set(types.UserKey(chanID), k.cdc.MustMarshalBinaryBare(name))
}

func (k Keeper) UpdateUser(ctx sdk.Context, chanID, name string) sdk.Error {
	id, err := k.GetUser(ctx, chanID)
	if strings.EqualFold(id, name) {
		return sdk.NewError("ibcsend", 1919, fmt.Sprintf("data already exist %s", err))
	}

	k.SetUser(ctx, chanID, name)
	packet := types.PacketMsgUser{Name: name}

	fmt.Println("Packet++++++++++++++++++++++++++", chanID, packet)
	return k.port.Send(ctx, chanID, packet)
}
