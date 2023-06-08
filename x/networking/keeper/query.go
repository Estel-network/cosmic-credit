package keeper

import (
	"cosmic-credit/x/networking/types"
)

var _ types.QueryServer = Keeper{}
