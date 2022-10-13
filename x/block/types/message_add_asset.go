package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAsset = "add_asset"

var _ sdk.Msg = &MsgAddAsset{}

func NewMsgAddAsset(creator string, name string, denom string, decimal string, price string, appId string, ibcStatus string) *MsgAddAsset {
	return &MsgAddAsset{
		Creator:   creator,
		Name:      name,
		Denom:     denom,
		Decimal:   decimal,
		Price:     price,
		AppId:     appId,
		IbcStatus: ibcStatus,
	}
}

func (msg *MsgAddAsset) Route() string {
	return RouterKey
}

func (msg *MsgAddAsset) Type() string {
	return TypeMsgAddAsset
}

func (msg *MsgAddAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
