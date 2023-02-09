package keeper_test

import (
	"chat/testutil"
	keepertest "chat/testutil/keeper"
	"chat/x/chat/keeper"
	"chat/x/chat/types"
	"reflect"
	"testing"
)

func TestKeeper_ShowGrooupMessages(t *testing.T) {
	k, ctx := keepertest.ChatKeeper(t)
	s := keeper.NewMsgServerImpl(*k)

	tests := []*types.MsgSendGroup{
		{
			Creator: testutil.Alice,
			Gid:     5,
			Content: "Any plans this weekend??",
		},
		{
			Creator: testutil.Bob,
			Gid:     5,
			Content: "No",
		},
		{
			Creator: testutil.Carol,
			Gid:     5,
			Content: "Nope",
		},
	}

	for _, test := range tests {
		s.SendGroup(ctx, test)
	}

	rsp, _ := k.ShowGroupMessage(ctx,
		&types.QueryShowGroupMessagesRequest{
			Gid: 5},
	)

	if !reflect.DeepEqual(rsp.Pagination.Total, uint64(3)) {
		t.Errorf("ShowSentMessages() = %v, want %v", rsp.Pagination.Total, 3)
	}

	rsp2, _ := k.ShowGroupMessage(ctx,
		&types.QueryShowGroupMessagesRequest{
			Gid: 1},
	)

	if !reflect.DeepEqual(rsp2.Pagination.Total, uint64(0)) {
		t.Errorf("ShowSentMessages() = %v, want %v", rsp2.Pagination.Total, 0)
	}
}
