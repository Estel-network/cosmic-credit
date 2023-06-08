package keeper

import (
	"cosmic-credit/x/credit/types"
)

var _ types.QueryServer = Keeper{}
