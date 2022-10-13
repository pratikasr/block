package keeper

import (
	"block/x/block/types"
)

var _ types.QueryServer = Keeper{}
