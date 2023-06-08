package keeper_test

import (
	"context"
	"testing"

	keepertest "cosmic-credit/testutil/keeper"
	"cosmic-credit/x/networking/keeper"
	"cosmic-credit/x/networking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.NetworkingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
