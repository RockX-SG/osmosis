package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) InterfluidDelegate(goCtx context.Context, delegate *types.MsgInterfluidDelegate) (*types.MsgInterfluidDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := server.keeper.SuperfluidDelegate(ctx, msg.Sender, msg.LockId, msg.ValAddr)
	if err == nil {
		events.EmitSuperfluidDelegateEvent(ctx, msg.LockId, msg.ValAddr)
	}
	return &types.MsgSuperfluidDelegateResponse{}, err
}

func (server msgServer) InterfluidUndelegate(ctx context.Context, undelegate *types.MsgInterfluidUndelegate) (*types.MsgInterfluidUndelegateResponse, error) {
	//TODO implement me
	panic("implement me")
}
