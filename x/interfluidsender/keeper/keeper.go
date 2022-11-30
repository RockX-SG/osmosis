package keeper

import (
	"fmt"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		*ibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramSpace paramtypes.Subspace

		ibcKeeper *ibckeeper.Keeper

		ak authkeeper.AccountKeeper
		bk types.BankKeeper
		sk types.StakingKeeper
		ck types.CommunityPoolKeeper
		ek types.EpochKeeper
		lk types.LockupKeeper
		gk types.GammKeeper
		ik types.IncentivesKeeper

		lms types.LockupMsgServer
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ibcKeeper *ibckeeper.Keeper,
	ak authkeeper.AccountKeeper,
	bk types.BankKeeper,
	sk types.StakingKeeper,
	ck types.CommunityPoolKeeper,
	ek types.EpochKeeper,
	lk types.LockupKeeper,
	gk types.GammKeeper,
	ik types.IncentivesKeeper,

	lms types.LockupMsgServer,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		Keeper:     ibcKeeper,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: ps,
		ak:         authkeeper.AccountKeeper{},
		bk:         bk,
		sk:         sk,
		ck:         ck,
		ek:         ek,
		lk:         lk,
		gk:         gk,
		ik:         ik,
		lms:        lms,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
