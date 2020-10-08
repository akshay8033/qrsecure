package qrsecure

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/akshay8033/qrsecure/x/qrsecure/types"
	"github.com/akshay8033/qrsecure/x/qrsecure/keeper"
)

func handleMsgCreateProduct(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateProduct) (*sdk.Result, error) {
	var Product = types.Product{
		Creator_id: msg.Creator_id,
		Application_ID:      msg.Application_ID,
    	Data: msg.Data,
	}
	k.CreateProduct(ctx, Product)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
