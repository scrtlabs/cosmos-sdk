package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/enigmampc/cosmos-sdk/x/distribution/types"
	"github.com/enigmampc/cosmos-sdk/x/simulation"
)

const (
	keyCommunityTax        = "communitytax"
	keyBaseProposerReward  = "baseproposerreward"
	keyBonusProposerReward = "bonusproposerreward"
	keySecretFoundationTax = "secret_foundation_tax"
)

// ParamChanges defines the parameters that can be modified by param change
// proposals in simulations.
func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyCommunityTax,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenCommunityTax(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keySecretFoundationTax,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenSecretFoundationTax(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyBaseProposerReward,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenBaseProposerReward(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyBonusProposerReward,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenBonusProposerReward(r))
			},
		),
	}
}
