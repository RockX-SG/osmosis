package events

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/x/gamm/utils"
	"github.com/osmosis-labs/osmosis/v12/x/interfluidsender/types"
)

func EmitSetInterfluidAssetEvent(ctx sdk.Context, denom string, assetType types.InterfluidAssetType) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newSetInterfluidAssetEvent(denom, assetType),
	})
}

func newSetInterfluidAssetEvent(denom string, assetType types.InterfluidAssetType) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtSetInterfluidAsset,
		sdk.NewAttribute(types.AttributeDenom, denom),
		sdk.NewAttribute(types.AttributeInterfluidAssetType, assetType.String()),
	)
}

func EmitRemoveInterfluidAsset(ctx sdk.Context, denom string) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newRemoveInterfluidAssetEvent(denom),
	})
}

func newRemoveInterfluidAssetEvent(denom string) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtRemoveInterfluidAsset,
		sdk.NewAttribute(types.AttributeDenom, denom),
	)
}

func EmitInterfluidDelegateEvent(ctx sdk.Context, lockId uint64, valAddress string) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newInterfluidDelegateEvent(lockId, valAddress),
	})
}

func newInterfluidDelegateEvent(lockId uint64, valAddress string) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtInterfluidDelegate,
		sdk.NewAttribute(types.AttributeLockId, utils.Uint64ToString(lockId)),
		sdk.NewAttribute(types.AttributeValidator, valAddress),
	)
}

func EmitInterfluidIncreaseDelegationEvent(ctx sdk.Context, lockId uint64, amount sdk.Coins) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newInterfluidIncreaseDelegationEvent(lockId, amount),
	})
}

func newInterfluidIncreaseDelegationEvent(lockId uint64, amount sdk.Coins) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtInterfluidIncreaseDelegation,
		sdk.NewAttribute(types.AttributeLockId, fmt.Sprintf("%d", lockId)),
		sdk.NewAttribute(types.AttributeAmount, amount.String()),
	)
}

func EmitInterfluidUndelegateEvent(ctx sdk.Context, lockId uint64) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newInterfluidUndelegateEvent(lockId),
	})
}

func newInterfluidUndelegateEvent(lockId uint64) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtInterfluidUndelegate,
		sdk.NewAttribute(types.AttributeLockId, fmt.Sprintf("%d", lockId)),
	)
}

func EmitInterfluidUnbondLockEvent(ctx sdk.Context, lockId uint64) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newInterfluidUnbondLockEvent(lockId),
	})
}

func newInterfluidUnbondLockEvent(lockId uint64) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtInterfluidUnbondLock,
		sdk.NewAttribute(types.AttributeLockId, fmt.Sprintf("%d", lockId)),
	)
}

func EmitUnpoolIdEvent(ctx sdk.Context, sender string, lpShareDenom string, allExitedLockIDsSerialized []byte) {
	if ctx.EventManager() == nil {
		return
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		newUnpoolIdEvent(sender, lpShareDenom, allExitedLockIDsSerialized),
	})
}

func newUnpoolIdEvent(sender string, lpShareDenom string, allExitedLockIDsSerialized []byte) sdk.Event {
	return sdk.NewEvent(
		types.TypeEvtUnpoolId,
		sdk.NewAttribute(sdk.AttributeKeySender, sender),
		sdk.NewAttribute(types.AttributeDenom, lpShareDenom),
		sdk.NewAttribute(types.AttributeNewLockIds, string(allExitedLockIDsSerialized)),
	)
}
