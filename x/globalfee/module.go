package globalfee

import (
	"context"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/x/globalfee/client/cli"
	"github.com/bandprotocol/chain/v2/x/globalfee/exported"
	"github.com/bandprotocol/chain/v2/x/globalfee/keeper"
	"github.com/bandprotocol/chain/v2/x/globalfee/types"
)

// ConsensusVersion defines the current x/globalfee module consensus version.
const ConsensusVersion = 2

var (
	_ module.AppModuleBasic   = AppModuleBasic{}
	_ module.AppModuleGenesis = AppModule{}
	_ module.AppModule        = AppModule{}
)

// AppModuleBasic defines the basic application module used by the wasm module.
type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&types.GenesisState{
		Params: types.DefaultParams(),
	})
}

func (a AppModuleBasic) ValidateGenesis(
	marshaler codec.JSONCodec,
	config client.TxEncodingConfig,
	message json.RawMessage,
) error {
	var data types.GenesisState
	err := marshaler.UnmarshalJSON(message, &data)
	if err != nil {
		return err
	}
	if err := data.Params.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "params")
	}
	return nil
}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(amino)
}

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
	// legacySubspace is used solely for migration of x/params managed parameters
	legacySubspace exported.Subspace
}

// NewAppModule constructor
func NewAppModule(keeper keeper.Keeper, ss exported.Subspace) *AppModule {
	return &AppModule{keeper: keeper, legacySubspace: ss}
}

func (a AppModule) InitGenesis(
	ctx sdk.Context,
	marshaler codec.JSONCodec,
	message json.RawMessage,
) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	marshaler.MustUnmarshalJSON(message, &genesisState)
	a.keeper.SetParams(ctx, genesisState.Params)
	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, marshaler codec.JSONCodec) json.RawMessage {
	var genState types.GenesisState
	genState.Params = a.keeper.GetParams(ctx)
	return marshaler.MustMarshalJSON(&genState)
}

func (a AppModule) RegisterInvariants(registry sdk.InvariantRegistry) {
}

// 0.47 TODO: delete these functions
// func (a AppModule) Route() sdk.Route {
// 	return sdk.Route{}
// }

// func (a AppModule) QuerierRoute() string {
// 	return types.QuerierRoute
// }

// func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
// 	return nil
// }

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.Querier{Keeper: am.keeper})

	m := keeper.NewMigrator(am.keeper, am.legacySubspace)

	if err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2); err != nil {
		panic(fmt.Sprintf("failed to migrate x/%s from version 1 to 2: %v", types.ModuleName, err))
	}
}

func (a AppModule) BeginBlock(context sdk.Context, block abci.RequestBeginBlock) {
}

func (a AppModule) EndBlock(context sdk.Context, block abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

// ConsensusVersion is a sequence number for state-breaking change of the
// module. It should be incremented on each consensus-breaking change
// introduced by the module. To avoid wrong/empty versions, the initial version
// should be set to 1.
func (a AppModule) ConsensusVersion() uint64 {
	return ConsensusVersion
}
