package bank

//
//import (
//	"encoding/json"
//	"github.com/baymax19/cosmos-ibc/modules/bank/types"
//
//	"github.com/gorilla/mux"
//	"github.com/spf13/cobra"
//
//	abci "github.com/tendermint/tendermint/abci/types"
//
//	"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//)
//
//const (
//	ModuleName = "ibcbank"
//	StoreKey   = ModuleName
//)
//
//var (
//	_ module.AppModule      = AppModule{}
//	_ module.AppModuleBasic = AppModuleBasic{}
//)
//
//type AppModuleBasic struct{}
//
//var _ module.AppModuleBasic = AppModuleBasic{}
//
//func (AppModuleBasic) Name() string {
//	return ModuleName
//}
//
//func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
//	types.RegisterCodec(cdc)
//}
//
//func (AppModuleBasic) DefaultGenesis() json.RawMessage {
//	return nil
//}
//
//func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
//	return nil
//}
//
//func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
//	//noop
//}
//
//func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
//	return nil
//}
//
//func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
//	return nil
//}
//
//type AppModule struct {
//	AppModuleBasic
//}
//
//func NewAppModule() AppModule {
//	return AppModule{
//	}
//}
//
//func (AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
//
//}
//
//func (AppModule) Name() string {
//	return ModuleName
//}
//
//func (AppModule) Route() string {
//	return ModuleName
//}
//
//func (am AppModule) NewHandler() sdk.Handler {
//	return nil
//}
//
//func (am AppModule) QuerierRoute() string {
//	return ModuleName
//}
//
//func (am AppModule) NewQuerierHandler() sdk.Querier {
//	return nil
//}
//
//func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
//	return []abci.ValidatorUpdate{}
//}
//
//func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
//	return nil
//}
//
//func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
//
//}
//
//func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
//	return []abci.ValidatorUpdate{}
//}
