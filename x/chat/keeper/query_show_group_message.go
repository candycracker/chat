package keeper

import (
	"context"
	"strconv"

	"chat/x/chat/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowGroupMessage(goCtx context.Context, req *types.QueryShowGroupMessagesRequest) (*types.QueryShowGroupMessagesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	var msgs []types.Message

	postStore := prefix.NewStore(store, types.KeyPrefix(types.MsgGroupKey+"/"+strconv.FormatUint(req.Gid, 10)))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var msg types.Message
		if err := k.cdc.Unmarshal(value, &msg); err != nil {
			return err
		}

		msgs = append(msgs, msg)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowGroupMessagesResponse{Message: msgs, Pagination: pageRes}, nil
}
