package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// Simulation parameter constants
const (
	DepositParamsMinDeposit    = "deposit_params_min_deposit"
	DepositParamsMinExpeditedDeposit  = "deposit_params_min_expedited_deposit"
	DepositParamsDepositPeriod = "deposit_params_deposit_period"
	DepositMinInitialRatio     = "deposit_params_min_initial_ratio"
	VotingParamsVotingPeriod   = "voting_params_voting_period"
	ExpeditedVotingParamsVotingPeriod = "expedited_voting_params_voting_period"
	TallyParamsQuorum          = "tally_params_quorum"
	TallyParamsThreshold       = "tally_params_threshold"
	TallyParamsExpeditedThreshold     = "tally_params_expedited_threshold"
	TallyParamsVeto            = "tally_params_veto"

	// ExpeditedThreshold must be at least as large as the regular Threshold
	// Therefore, we use this break out point in randomization.
	tallyNonExpeditedMax = 500

	// Similarly, expedited voting period must be strictly less than the regular
	// voting period to be valid. Therefore, we use this break out point in randomization.
	expeditedMaxVotingPeriod = 60 * 60 * 24 * 2
)

// GenDepositParamsDepositPeriod returns randomized DepositParamsDepositPeriod
func GenDepositParamsDepositPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, 2*60*60*24*2)) * time.Second
}

// GenDepositParamsMinDeposit returns randomized DepositParamsMinDeposit
func GenDepositParamsMinDeposit(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1, 1e3/2))))
}

// GenDepositParamsMinExpeditedDeposit randomized DepositParamsMinExpeditedDeposit
func GenDepositParamsMinExpeditedDeposit(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simulation.RandIntBetween(r, 1e3/2, 1e3))))
}

// GenDepositMinInitialRatio returns randomized DepositMinInitialRatio
func GenDepositMinInitialDepositRatio(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(int64(simulation.RandIntBetween(r, 0, 99))).Quo(sdk.NewDec(100))
}

// GenVotingParamsVotingPeriod returns randomized VotingParamsVotingPeriod
func GenVotingParamsVotingPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, expeditedMaxVotingPeriod, 2*expeditedMaxVotingPeriod)) * time.Second
}

// GenVotingParamsExpeditedVotingPeriod randomized VotingParamsExpeditedVotingPeriod
func GenVotingParamsExpeditedVotingPeriod(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 1, expeditedMaxVotingPeriod)) * time.Second
}

// GenTallyParamsQuorum returns randomized TallyParamsQuorum
func GenTallyParamsQuorum(r *rand.Rand) math.LegacyDec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 334, 500)), 3)
}

// GenTallyParamsThreshold returns randomized TallyParamsThreshold
func GenTallyParamsThreshold(r *rand.Rand) math.LegacyDec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 450, 550)), 3)
// GenTallyParamsThreshold randomized TallyParamsThreshold

// GenTallyParamsExpeditedThreshold randomized TallyParamsExpeditedThreshold
func GenTallyParamsExpeditedThreshold(r *rand.Rand) math.LegacyDec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, tallyNonExpeditedMax, 550)), 3)
}

// GenTallyParamsVeto returns randomized TallyParamsVeto
func GenTallyParamsVeto(r *rand.Rand) math.LegacyDec {
	return sdk.NewDecWithPrec(int64(simulation.RandIntBetween(r, 250, 334)), 3)
}

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) {
	startingProposalID := uint64(simState.Rand.Intn(100))

	var minDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinDeposit, &minDeposit, simState.Rand,
		func(r *rand.Rand) { minDeposit = GenDepositParamsMinDeposit(r) },
	)

	var minExpeditedDeposit sdk.Coins
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsMinExpeditedDeposit, &minExpeditedDeposit, simState.Rand,
		func(r *rand.Rand) { minExpeditedDeposit = GenDepositParamsMinExpeditedDeposit(r) },
	)

	var depositPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositParamsDepositPeriod, &depositPeriod, simState.Rand,
		func(r *rand.Rand) { depositPeriod = GenDepositParamsDepositPeriod(r) },
	)

	var minInitialDepositRatio sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DepositMinInitialRatio, &minInitialDepositRatio, simState.Rand,
		func(r *rand.Rand) { minInitialDepositRatio = GenDepositMinInitialDepositRatio(r) },
	)

	var votingPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, VotingParamsVotingPeriod, &votingPeriod, simState.Rand,
		func(r *rand.Rand) { votingPeriod = GenVotingParamsVotingPeriod(r) },
	)

	var expeditedVotingPeriod time.Duration
	simState.AppParams.GetOrGenerate(
		simState.Cdc, ExpeditedVotingParamsVotingPeriod, &expeditedVotingPeriod, simState.Rand,
		func(r *rand.Rand) { expeditedVotingPeriod = GenVotingParamsExpeditedVotingPeriod(r) },
	)

	var quorum sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsQuorum, &quorum, simState.Rand,
		func(r *rand.Rand) { quorum = GenTallyParamsQuorum(r) },
	)

	var threshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsThreshold, &threshold, simState.Rand,
		func(r *rand.Rand) { threshold = GenTallyParamsThreshold(r) },
	)

	var expeditedThreshold sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsExpeditedThreshold, &expeditedThreshold, simState.Rand,
		func(r *rand.Rand) { expeditedThreshold = GenTallyParamsExpeditedThreshold(r) },
	)

	var veto sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, TallyParamsVeto, &veto, simState.Rand,
		func(r *rand.Rand) { veto = GenTallyParamsVeto(r) },
	)

	govGenesis := v1.NewGenesisState(
		startingProposalID,
		v1.NewParams(minDeposit, depositPeriod, votingPeriod, quorum.String(), threshold.String(), veto.String(), minInitialDepositRatio.String(), simState.Rand.Intn(2) == 0, simState.Rand.Intn(2) == 0, simState.Rand.Intn(2) == 0),
	)

	bz, err := json.MarshalIndent(&govGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated governance parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(govGenesis)
}
