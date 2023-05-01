package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// Simulation parameter constants
const (
	CommunityTax            = "community_tax"
	WithdrawEnabled         = "withdraw_enabled"
	MinimumRestakeThreshold = "minimum_restake_threshold"
	RestakePeriod           = "restake_period"
)

// GenMinimumRestakeThreshold returns a randomized restake threshold.
func GenMinimumRestakeThreshold(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(int64(r.Intn(100_000_000)))
}

// GenMinimumRestakeThreshold returns a randomized restake threshold.
func GenRestakePeriod(r *rand.Rand) sdk.Int {
	return sdk.NewInt(int64(r.Intn(100_000_000)))
}

// GenCommunityTax randomized CommunityTax
func GenCommunityTax(r *rand.Rand) math.LegacyDec {
	return sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2))
}

// GenWithdrawEnabled returns a randomized WithdrawEnabled parameter.
func GenWithdrawEnabled(r *rand.Rand) bool {
	return r.Int63n(101) <= 95 // 95% chance of withdraws being enabled
}

// RandomizedGenState generates a random GenesisState for distribution
func RandomizedGenState(simState *module.SimulationState) {
	var communityTax sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, CommunityTax, &communityTax, simState.Rand,
		func(r *rand.Rand) { communityTax = GenCommunityTax(r) },
	)

	var withdrawEnabled bool
	simState.AppParams.GetOrGenerate(
		simState.Cdc, WithdrawEnabled, &withdrawEnabled, simState.Rand,
		func(r *rand.Rand) { withdrawEnabled = GenWithdrawEnabled(r) },
	)

	var restakeThreshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, MinimumRestakeThreshold, &restakeThreshold, simState.Rand,
		func(r *rand.Rand) { restakeThreshold = GenMinimumRestakeThreshold(r) },
	)

	var restakePeriod sdk.Int
	simState.AppParams.GetOrGenerate(
		simState.Cdc, RestakePeriod, &restakePeriod, simState.Rand,
		func(r *rand.Rand) { restakePeriod = GenRestakePeriod(r) },
	)

	foundationTaxAcc, _ := simulation.RandomAcc(simState.Rand, simState.Accounts)

	distrGenesis := types.GenesisState{
		FeePool: types.InitialFeePool(),
		Params: types.Params{
			CommunityTax:        communityTax,
			WithdrawAddrEnabled: withdrawEnabled,
			MinimumRestakeThreshold: restakeThreshold,
			RestakePeriod:           restakePeriod,
		},
	}

	bz, err := json.MarshalIndent(&distrGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated distribution parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&distrGenesis)
}
