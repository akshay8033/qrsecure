package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateProduct{}

type MsgCreateProduct struct {
  Application_ID      string
  Creator_id sdk.AccAddress `json:"creator" yaml:"creator"`
  Data string `json:"data" yaml:"data"`
}

func NewMsgCreateProduct(creator sdk.AccAddress, data string) MsgCreateProduct {
  return MsgCreateProduct{
    Application_ID: uuid.New().String(),
		Creator_id: creator,
    Data: data,
	}
}

func (msg MsgCreateProduct) Route() string {
  return RouterKey
}

func (msg MsgCreateProduct) Type() string {
  return "CreateProduct"
}

func (msg MsgCreateProduct) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator_id)}
}

func (msg MsgCreateProduct) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateProduct) ValidateBasic() error {
  if msg.Creator_id.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}