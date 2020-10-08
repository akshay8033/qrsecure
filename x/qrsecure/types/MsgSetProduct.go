package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetProduct{}

type MsgSetProduct struct {
  Application_ID      string      `json:"id" yaml:"id"`
  Creator_id sdk.AccAddress `json:"creator" yaml:"creator"`
  Data string `json:"data" yaml:"data"`
}

func NewMsgSetProduct(creator sdk.AccAddress, id string, data string) MsgSetProduct {
  return MsgSetProduct{
    Application_ID: id,
		Creator_id: creator,
    Data: data,
	}
}

func (msg MsgSetProduct) Route() string {
  return RouterKey
}

func (msg MsgSetProduct) Type() string {
  return "SetProduct"
}

func (msg MsgSetProduct) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator_id)}
}

func (msg MsgSetProduct) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetProduct) ValidateBasic() error {
  if msg.Creator_id.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}