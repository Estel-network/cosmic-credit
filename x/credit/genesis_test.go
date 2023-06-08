package credit_test

import (
	"testing"

	keepertest "cosmic-credit/testutil/keeper"
	"cosmic-credit/testutil/nullify"
	"cosmic-credit/x/credit"
	"cosmic-credit/x/credit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ModuleInfo: types.ModuleInfo{
			Enabled:              true,
			TotalPositions:       28,
			TotalCredited:        24,
			CreditFee:            66,
			RewardAmount:         54,
			RewardTime:           "85",
			LiquidationThreshold: 65,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CreditKeeper(t)
	credit.InitGenesis(ctx, *k, genesisState)
	got := credit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.ModuleInfo, got.ModuleInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
