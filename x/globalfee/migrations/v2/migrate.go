package v2

import (
	"github.com/bandprotocol/chain/v2/x/globalfee/exported"
	"github.com/bandprotocol/chain/v2/x/globalfee/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "globalfee"
)

// Migrate migrates the x/globalfee module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/globalfee
// module state.
func Migrate(
	ctx sdk.Context,
	store sdk.KVStore,
	legacySubspace exported.Subspace,
	cdc codec.BinaryCodec,
) error {
	var currParams types.Params
	legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&currParams)
	store.Set(types.ParamsKeyPrefix, bz)

	return nil
}
