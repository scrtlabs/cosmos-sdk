package v1beta1

import (
	"time"

	"sigs.k8s.io/yaml"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default period for deposits & voting
const (
	DefaultPeriod          time.Duration = time.Hour * 24 * 2 // 2 days
	DefaultExpeditedPeriod time.Duration = time.Hour * 24     // 1 day
)

// Default governance params
var (
	DefaultMinDepositTokens          = sdk.NewInt(10000000)
	DefaultMinExpeditedDepositTokens = sdk.NewInt(10000000 * 5)
	DefaultQuorum                    = sdk.NewDecWithPrec(334, 3)
	DefaultThreshold                 = sdk.NewDecWithPrec(5, 1)
	DefaultExpeditedThreshold        = sdk.NewDecWithPrec(667, 3)
	DefaultVetoThreshold             = sdk.NewDecWithPrec(334, 3)
)

// NewDepositParams creates a new DepositParams object
func NewDepositParams(minDeposit sdk.Coins, maxDepositPeriod time.Duration, minExpeditedDeposit sdk.Coins) DepositParams {
	return DepositParams{
		MinDeposit:          minDeposit,
		MaxDepositPeriod:    maxDepositPeriod,
		MinExpeditedDeposit: minExpeditedDeposit,
	}
}

// DefaultDepositParams returns the default parameters for deposits
func DefaultDepositParams() DepositParams {
	return NewDepositParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultMinDepositTokens)),
		DefaultPeriod,
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultMinExpeditedDepositTokens)),
	)
}

// String implements stringer insterface
func (dp DepositParams) String() string {
	out, _ := yaml.Marshal(dp)
	return string(out)
}

// Equal checks equality of DepositParams
func (dp DepositParams) Equal(dp2 DepositParams) bool {
	return dp.MinDeposit.IsEqual(dp2.MinDeposit) && dp.MaxDepositPeriod == dp2.MaxDepositPeriod
}

// NewTallyParams creates a new TallyParams object
func NewTallyParams(quorum, threshold, expeditedThreshold, vetoThreshold sdk.Dec) TallyParams {
	return TallyParams{
		Quorum:             quorum,
		Threshold:          threshold,
		ExpeditedThreshold: expeditedThreshold,
		VetoThreshold:      vetoThreshold,
	}
}

// DefaultTallyParams returns default parameters for tallying
func DefaultTallyParams() TallyParams {
	return NewTallyParams(DefaultQuorum, DefaultThreshold, DefaultExpeditedThreshold, DefaultVetoThreshold)
}

// GetThreshold returns threshold based on the value isExpedited
func (tp TallyParams) GetThreshold(isExpedited bool) sdk.Dec {
	if isExpedited {
		return tp.ExpeditedThreshold
	}
	return tp.Threshold
}

// Equal checks equality of TallyParams
func (tp TallyParams) Equal(other TallyParams) bool {
	return tp.Quorum.Equal(other.Quorum) && tp.Threshold.Equal(other.Threshold) && tp.ExpeditedThreshold.Equal(other.ExpeditedThreshold) && tp.VetoThreshold.Equal(other.VetoThreshold)
}

// String implements stringer insterface
func (tp TallyParams) String() string {
	out, _ := yaml.Marshal(tp)
	return string(out)
}

// NewVotingParams creates a new VotingParams object
func NewVotingParams(votingPeriod time.Duration, expeditedPeriod time.Duration) VotingParams {
	return VotingParams{
		VotingPeriod:          votingPeriod,
		ExpeditedVotingPeriod: expeditedPeriod,
	}
}

// DefaultVotingParams default parameters for voting
func DefaultVotingParams() VotingParams {
	return NewVotingParams(DefaultPeriod, DefaultExpeditedPeriod)
}

// GetVotingPeriod returns voting period based on whether isExpedited is requested.
func (vp VotingParams) GetVotingPeriod(isExpedited bool) time.Duration {
	if isExpedited {
		return vp.ExpeditedVotingPeriod
	}
	return vp.VotingPeriod
}

// String implements stringer interface
func (vp VotingParams) String() string {
	out, _ := yaml.Marshal(vp)
	return string(out)
}

// Params returns all of the governance params
type Params struct {
	VotingParams  VotingParams  `json:"voting_params" yaml:"voting_params"`
	TallyParams   TallyParams   `json:"tally_params" yaml:"tally_params"`
	DepositParams DepositParams `json:"deposit_params" yaml:"deposit_params"`
}

// String implements stringer interface
func (gp Params) String() string {
	return gp.VotingParams.String() + "\n" +
		gp.TallyParams.String() + "\n" + gp.DepositParams.String()
}

// NewParams creates a new gov Params instance
func NewParams(vp VotingParams, tp TallyParams, dp DepositParams) Params {
	return Params{
		VotingParams:  vp,
		DepositParams: dp,
		TallyParams:   tp,
	}
}

// DefaultParams returns the default governance params
func DefaultParams() Params {
	return NewParams(DefaultVotingParams(), DefaultTallyParams(), DefaultDepositParams())
}
