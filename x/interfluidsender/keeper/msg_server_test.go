package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
    "github.com/osmosis-labs/osmosis/v12/x/interfluidsender/keeper"
    keepertest "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterfluidsenderKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
