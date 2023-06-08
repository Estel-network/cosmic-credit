package types_test

import (
	"testing"

	"cosmic-credit/x/credit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ModuleInfo: types.ModuleInfo{
					Enabled:              true,
					TotalPositions:       37,
					TotalCredited:        94,
					CreditFee:            45,
					RewardAmount:         55,
					RewardTime:           "71",
					LiquidationThreshold: 20,
				},
				CreditList: []types.Credit{
					{
						Owner: "0",
					},
					{
						Owner: "1",
					},
				},
				CollateralList: []types.Collateral{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated credit",
			genState: &types.GenesisState{
				CreditList: []types.Credit{
					{
						Owner: "0",
					},
					{
						Owner: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated collateral",
			genState: &types.GenesisState{
				CollateralList: []types.Collateral{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
