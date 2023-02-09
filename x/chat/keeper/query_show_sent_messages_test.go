package keeper_test

import (
	"chat/testutil"
	keepertest "chat/testutil/keeper"
	"chat/x/chat/keeper"
	"chat/x/chat/types"
	"reflect"
	"testing"
)

func TestKeeper_ShowSentMessages(t *testing.T) {
	k, ctx := keepertest.ChatKeeper(t)
	s := keeper.NewMsgServerImpl(*k)

	tests := []*types.MsgSend{
		{
			Creator:   testutil.Alice,
			Recipient: testutil.Bob,
			Content:   "Hello",
		},
		{
			Creator:   testutil.Bob,
			Recipient: testutil.Alice,
			Content:   "Hi",
		},
		{
			Creator:   testutil.Alice,
			Recipient: testutil.Bob,
			Content:   "How are you",
		},
	}

	for _, test := range tests {
		s.Send(ctx, test)
	}

	rsp, _ := k.ShowSentMessages(ctx,
		&types.QueryShowSentMessagesRequest{
			Creator: testutil.Alice},
	)

	if !reflect.DeepEqual(rsp.Pagination.Total, uint64(2)) {
		t.Errorf("ShowSentMessages() = %v, want %v", rsp.Pagination.Total, 2)
	}

	rsp2, _ := k.ShowSentMessages(ctx,
		&types.QueryShowSentMessagesRequest{
			Creator: testutil.Bob},
	)

	if !reflect.DeepEqual(rsp2.Pagination.Total, uint64(1)) {
		t.Errorf("ShowSentMessages() = %v, want %v", rsp2.Pagination.Total, 1)
	}
}
