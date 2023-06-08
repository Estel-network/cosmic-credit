package keeper_test

import (
	"testing"

	testkeeper "cosmic-credit/testutil/keeper"
	"cosmic-credit/x/credit/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CreditKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
