package keeper

import (
	"github.com/osmosis-labs/osmosis/v12/x/interfluidreceiver/types"
)

var _ types.QueryServer = Keeper{}
