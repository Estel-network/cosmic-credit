package keeper_test

import (
	"context"
	"testing"

	keepertest "cosmic-credit/testutil/keeper"
	"cosmic-credit/x/credit/keeper"
	"cosmic-credit/x/credit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CreditKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
