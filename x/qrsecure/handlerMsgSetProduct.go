package qrsecure

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/akshay8033/qrsecure/x/qrsecure/types"
	"github.com/akshay8033/qrsecure/x/qrsecure/keeper"
)

func handleMsgSetProduct(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetProduct) (*sdk.Result, error) {
	var Product = types.Product{
		Creator_id: msg.Creator_id,
		Application_ID:      msg.Application_ID,
    	Data: msg.Data,
	}
	if !msg.Creator_id.Equals(k.GetProductOwner(ctx, msg.Application_ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetProduct(ctx, Product)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
