package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendInterfluidDelegatePacketData = "send_interfluid_delegate_packet_data"

var _ sdk.Msg = &MsgSendInterfluidDelegatePacketData{}

func NewMsgSendInterfluidDelegatePacketData(
    creator string,
    port string,
    channelID string,
    timeoutTimestamp uint64,
    delAddr string,
    valAddr string,
    coin string,
) *MsgSendInterfluidDelegatePacketData {
    return &MsgSendInterfluidDelegatePacketData{
		Creator: creator,
		Port: port,
		ChannelID: channelID,
		TimeoutTimestamp: timeoutTimestamp,
        DelAddr: delAddr,
        ValAddr: valAddr,
        Coin: coin,
	}
}

func (msg *MsgSendInterfluidDelegatePacketData) Route() string {
    return RouterKey
}

func (msg *MsgSendInterfluidDelegatePacketData) Type() string {
    return TypeMsgSendInterfluidDelegatePacketData
}

func (msg *MsgSendInterfluidDelegatePacketData) GetSigners() []sdk.AccAddress {
    creator, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        panic(err)
    }
    return []sdk.AccAddress{creator}
}

func (msg *MsgSendInterfluidDelegatePacketData) GetSignBytes() []byte {
    bz := ModuleCdc.MustMarshalJSON(msg)
    return sdk.MustSortJSON(bz)
}

func (msg *MsgSendInterfluidDelegatePacketData) ValidateBasic() error {
    _, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
    }
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
    return nil
}
