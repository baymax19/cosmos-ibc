package ibcsend

import (
	"encoding/json"
	cli2 "github.com/baymax19/cosmos-ibc/modules/ibc/bank/ibcsend/cli"
	types2 "github.com/baymax19/cosmos-ibc/modules/ibc/bank/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

const ModuleName = "send"

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types2.RegisterSend(cdc)
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(context.CLIContext, *mux.Router) {

}

func (a AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli2.GetTxCmd(cdc)
}

func (a AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli2.GetQueryCmd(ModuleName, cdc)
}

type AppModule struct {
	AppModuleBasic
	k Keeper
}

func NewAppModule(k Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		k:              k,
	}
}

func (a AppModule) Name() string {
	return ModuleName
}

func (a AppModule) InitGenesis(sdk.Context, json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (a AppModule) ExportGenesis(sdk.Context) json.RawMessage {
	return nil
}

func (a AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (a AppModule) Route() string {
	return ModuleName
}

func (a AppModule) NewHandler() sdk.Handler {
	return NewHandler(a.k)
}

func (a AppModule) QuerierRoute() string {
	return ModuleName
}

func (a AppModule) NewQuerierHandler() sdk.Querier {
	return nil
}

func (a AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
}

func (a AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
