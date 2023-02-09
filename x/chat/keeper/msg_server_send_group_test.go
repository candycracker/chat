package keeper_test

import (
	"chat/testutil"
	keepertest "chat/testutil/keeper"
	"chat/x/chat/keeper"
	"chat/x/chat/types"
	"reflect"
	"testing"
)

func TestSendGroupMessage(t *testing.T) {
	k, ctx := keepertest.ChatKeeper(t)
	s := keeper.NewMsgServerImpl(*k)
	rst, _ := s.SendGroup(ctx,
		&types.MsgSendGroup{
			Creator: testutil.Alice,
			Gid:     5,
			Content: "Hello"},
	)

	if !reflect.DeepEqual(rst.Id, uint64(0)) {
		t.Errorf("Send() = %v, want %v", rst.Id, 0)
	}

}
