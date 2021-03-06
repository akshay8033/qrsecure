package qrsecure

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/akshay8033/qrsecure/x/qrsecure/keeper"
	"github.com/akshay8033/qrsecure/x/qrsecure/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding # 1
		case types.MsgCreateProduct:
			return handleMsgCreateProduct(ctx, k, msg)
		case types.MsgSetProduct:
			return handleMsgSetProduct(ctx, k, msg)
		case types.MsgDeleteProduct:
			return handleMsgDeleteProduct(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
