package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Product struct {
	Creator_id sdk.AccAddress `json:"creator" yaml:"creator"`
	Application_ID      string         `json:"id" yaml:"id"`
    Data string `json:"data" yaml:"data"`
}