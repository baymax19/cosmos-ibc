package ibcrecv

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

func (k Keeper) SetUser(ctx sdk.Context, chainID, name string) {
	store := ctx.KVStore(k.key)
	store.Set(types.UserKey(chainID), k.cdc.MustMarshalBinaryBare(name))
}

func (k Keeper) UpdateUser(ctx sdk.Context, chaiID, name string) sdk.Error {
	id, err := k.GetUser(ctx, chaiID)

	fmt.Println("Equal fold Data", name, id)
	if strings.EqualFold(id, name) {

		return sdk.NewError("ibcsend", 1919, fmt.Sprintf("data already exist %s", err))
	}

	k.SetUser(ctx, chaiID, name)
	return nil
}
