package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteProduct{}

type MsgDeleteProduct struct {
  Application_ID      string         `json:"id" yaml:"id"`
  Creator_id sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteProduct(id string, creator sdk.AccAddress) MsgDeleteProduct {
  return MsgDeleteProduct{
    Application_ID: id,
		Creator_id: creator,
	}
}

func (msg MsgDeleteProduct) Route() string {
  return RouterKey
}

func (msg MsgDeleteProduct) Type() string {
  return "DeleteProduct"
}

func (msg MsgDeleteProduct) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator_id)}
}

func (msg MsgDeleteProduct) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteProduct) ValidateBasic() error {
  if msg.Creator_id.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}