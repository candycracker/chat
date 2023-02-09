package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendGroup = "send_group"

var _ sdk.Msg = &MsgSendGroup{}

func NewMsgSendGroup(creator string, content string, gid uint64) *MsgSendGroup {
	return &MsgSendGroup{
		Creator: creator,
		Content: content,
		Gid:     gid,
	}
}

func (msg *MsgSendGroup) Route() string {
	return RouterKey
}

func (msg *MsgSendGroup) Type() string {
	return TypeMsgSendGroup
}

func (msg *MsgSendGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
