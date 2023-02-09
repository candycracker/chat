package keeper

import (
	"context"
	"strconv"
	"time"

	"chat/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendGroup(goCtx context.Context, msg *types.MsgSendGroup) (*types.MsgSendGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	groupId := strconv.FormatUint(msg.Gid, 10)
	id := k.GetPostCount(ctx)
	var message = types.Message{
		Sender:    msg.Creator,
		Recipient: groupId,
		Content:   msg.Content,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
		Id:        id,
	}
	k.SetMsgStore(ctx, message, types.MsgGroupKey+"/"+groupId)

	return &types.MsgSendGroupResponse{Id: id}, nil
}
