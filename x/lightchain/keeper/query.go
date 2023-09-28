package keeper

import (
	"lightchain/x/lightchain/types"
)

var _ types.QueryServer = Keeper{}
