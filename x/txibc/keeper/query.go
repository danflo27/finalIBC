package keeper

import (
	"final/x/txibc/types"
)

var _ types.QueryServer = Keeper{}
