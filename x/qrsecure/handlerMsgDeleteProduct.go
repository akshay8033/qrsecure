package qrsecure

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/akshay8033/qrsecure/x/qrsecure/types"
	"github.com/akshay8033/qrsecure/x/qrsecure/keeper"
)

// Handle a message to delete name
func handleMsgDeleteProduct(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteProduct) (*sdk.Result, error) {
	if !k.ProductExists(ctx, msg.Application_ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Application_ID)
	}
	if !msg.Creator_id.Equals(k.GetProductOwner(ctx, msg.Application_ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteProduct(ctx, msg.Application_ID)
	return &sdk.Result{}, nil
}
