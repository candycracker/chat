package keeper

import (
	"chat/x/chat/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetMsgStore(ctx sdk.Context, msg types.Message, key string) uint64 {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(key))
	b := k.cdc.MustMarshal(&msg)
	store.Set(GetPostIDBytes(msg.Id), b)

	return msg.Id
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MsgCountKey)
	bz := store.Get(byteKey)
	var idx uint64 = 0
	if bz == nil {
		idx = 0
	} else {
		idx = binary.BigEndian.Uint64(bz)
	}
	k.SetPostCount(ctx, idx+1)
	return idx
}

func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MsgCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
