package keeper

import (
	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
)

var _ types.QueryServer = Keeper{}
