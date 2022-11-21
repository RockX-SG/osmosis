package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/osmosis-labs/osmosis/v12/x/interfluidreceiver/types"
    "github.com/osmosis-labs/osmosis/v12/x/interfluidreceiver/keeper"
    keepertest "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterfluidreceiverKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
