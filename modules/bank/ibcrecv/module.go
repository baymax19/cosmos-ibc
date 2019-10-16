package ibcrecv

import (
	"encoding/json"
	"github.com/baymax19/cosmos-ibc/modules/bank/ibcrecv/cli"
	"github.com/baymax19/cosmos-ibc/modules/bank/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

const ModuleName = "receive"

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return ModuleName
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterRecv(cdc)
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
	return nil
}

func (a AppModuleBasic) ValidateGenesis(json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(context.CLIContext, *mux.Router) {

}

func (a AppModuleBasic) GetTxCmd(*codec.Codec) *cobra.Command {
	return nil
}

func (a AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(ModuleName, cdc)
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
