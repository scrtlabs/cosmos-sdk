package distribution

import (
	"time"
    "fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// determine the total power signing the block
	var previousTotalPower int64
	for _, voteInfo := range ctx.VoteInfos() {
		previousTotalPower += voteInfo.Validator.Power
	}

	// TODO this is Tendermint-dependent
	// ref https://github.com/cosmos/cosmos-sdk/issues/3095
	if ctx.BlockHeight() > 1 {
		if err := k.AllocateTokens(ctx, previousTotalPower, ctx.VoteInfos()); err != nil {
			return err
		}
	}

	restakeFunc := func(delegator sdk.AccAddress, validator sdk.ValAddress) (stop bool) {
		err := k.PerformRestake(ctx, delegator, validator)
		if err != nil {
			k.Logger(ctx).Info(fmt.Sprintf("Err: %s, Failed to perform restake for delegator-validator %s - %s", err, delegator, validator))
		}

		return err != nil
	}

    restakePeriod, err := k.GetRestakePeriod(ctx)
    if err != nil {
        return err
    }
	if ctx.BlockHeight()%restakePeriod.Int64() == 0 {
        staleKeys := k.IterateRestakeEntries(ctx, restakeFunc)

		for _, stale := range staleKeys {

			err := k.DeleteAutoRestakeEntry(ctx, stale.Delegator, stale.Validator)
			if err != nil {
				k.Logger(ctx).Info(fmt.Sprintf("Err: %s, Failed to delete restake key", err))
			}
		}
	}

	// record the proposer for when we payout on the next block
	consAddr := sdk.ConsAddress(ctx.BlockHeader().ProposerAddress)
	return k.SetPreviousProposerConsAddr(ctx, consAddr)
}
