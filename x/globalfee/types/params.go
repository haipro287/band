package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{MinimumGasPrices: sdk.DecCoins{}}
}

// ValidateBasic performs basic validation.
func (p Params) ValidateBasic() error {
	return validateMinimumGasPrices(p.MinimumGasPrices)
}

// this requires the fee non-negative
func validateMinimumGasPrices(i interface{}) error {
	v, ok := i.(sdk.DecCoins)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "type: %T, expected sdk.DecCoins", i)
	}

	return v.Validate()
}

func (p Params) Validate() error {
	return validateMinimumGasPrices(p.MinimumGasPrices)
}
