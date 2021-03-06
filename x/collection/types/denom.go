package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewDenom return a new denom
func NewDenom(id, name, schema, symbol string, creator sdk.AccAddress, mintRestricted, updateRestricted bool) Denom {
	return Denom{
		ID:               id,
		Name:             name,
		Schema:           schema,
		Creator:          creator.String(),
		Symbol:           symbol,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
	}
}
