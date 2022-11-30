package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/keeper"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterfluidsenderKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}