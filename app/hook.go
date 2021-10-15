package band

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Hooks []Hook

// Hook is an interface of hook that can process along with abci application
type Hook interface {
	AfterInitChain(ctx sdk.Context, req abci.RequestInitChain, res abci.ResponseInitChain)
	AfterBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock)
	AfterDeliverTx(ctx sdk.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx)
	AfterEndBlock(ctx sdk.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock)
	ApplyQuery(req abci.RequestQuery) (res abci.ResponseQuery, stop bool)
	BeforeCommit()
}

func (h Hooks) AfterInitChain(ctx sdk.Context, req abci.RequestInitChain, res abci.ResponseInitChain) {
	for _, hook := range h {
		hook.AfterInitChain(ctx, req, res)
	}
}

func (h Hooks) AfterBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) {
	for _, hook := range h {
		hook.AfterBeginBlock(ctx, req, res)
	}
}

func (h Hooks) AfterDeliverTx(ctx sdk.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) {
	for _, hook := range h {
		hook.AfterDeliverTx(ctx, req, res)
	}
}

func (h Hooks) AfterEndBlock(ctx sdk.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) {
	for _, hook := range h {
		hook.AfterEndBlock(ctx, req, res)
	}
}

func (h Hooks) BeforeCommit() {
	for _, hook := range h {
		hook.BeforeCommit()
	}
}

func (h Hooks) ApplyQuery(req abci.RequestQuery) (res abci.ResponseQuery, stop bool) {
	for _, hook := range h {
		res, stop = hook.ApplyQuery(req)
		if stop {
			return
		}
	}

	return
}