package interfluidreceiver_test

import (
	"testing"

	keepertest "github.com/osmosis-labs/osmosis/v12/testutil/keeper"
	"github.com/osmosis-labs/osmosis/v12/testutil/nullify"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidreceiver"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidreceiver/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InterfluidreceiverKeeper(t)
	interfluidreceiver.InitGenesis(ctx, *k, genesisState)
	got := interfluidreceiver.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
