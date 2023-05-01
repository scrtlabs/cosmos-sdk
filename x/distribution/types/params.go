package types

import (
	"fmt"

	"cosmossdk.io/math"
	"sigs.k8s.io/yaml"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:            sdk.NewDecWithPrec(2, 2), // 2%
		BaseProposerReward:      sdk.ZeroDec(),            // deprecated
		BonusProposerReward:     sdk.ZeroDec(),            // deprecated
		WithdrawAddrEnabled:     true,
		MinimumRestakeThreshold: sdk.NewDec(10_000_000),
		RestakePeriod:           sdk.NewInt(1000),
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, validateCommunityTax),
		paramtypes.NewParamSetPair(ParamStoreKeyBaseProposerReward, &p.BaseProposerReward, validateBaseProposerReward),
		paramtypes.NewParamSetPair(ParamStoreKeyBonusProposerReward, &p.BonusProposerReward, validateBonusProposerReward),
		paramtypes.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, validateWithdrawAddrEnabled),
		paramtypes.NewParamSetPair(ParamMinimumRestakeThreshold, &p.MinimumRestakeThreshold, validateMinimumRestakeThreshold),
		paramtypes.NewParamSetPair(ParamRestakePeriod, &p.RestakePeriod, validateRestakePeriod),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() || p.CommunityTax.GT(math.LegacyOneDec()) {
		return fmt.Errorf(
			"community tax should be non-negative and less than one: %s", p.CommunityTax,
		)
	}

	return nil
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("community tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("community tax must be positive: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateWithdrawAddrEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMinimumRestakeThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid minimum restake threshold parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("minimum restake threshold must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("minimum restake threshold must be positive: %s", v)
	}

	return nil
}

func validateRestakePeriod(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid minimum period parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("minimum period must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("minimum period must be positive: %s", v)
	}
	if v.LT(sdk.NewInt(1000)) {
		return fmt.Errorf("minimum period must be greater than 1000 blocks")
	}

	return nil
}
