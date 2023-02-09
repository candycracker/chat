package keeper

import (
	"context"
	"time"

	"chat/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	id := k.GetPostCount(ctx)
	var message = types.Message{
		Sender:    msg.Creator,
		Recipient: msg.Recipient,
		Content:   msg.Content,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		Id:        id,
	}
	k.SetMsgStore(ctx, message, types.MsgSenderKey+"/"+message.Sender)
	k.SetMsgStore(ctx, message, types.MsgRecipientKey+"/"+message.Recipient)

	return &types.MsgSendResponse{Id: id}, nil
}
