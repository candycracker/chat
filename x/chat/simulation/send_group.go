package simulation

import (
	"math/rand"

	"chat/x/chat/keeper"
	"chat/x/chat/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSendGroup(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSendGroup{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SendGroup simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendGroup simulation not implemented"), nil, nil
	}
}
