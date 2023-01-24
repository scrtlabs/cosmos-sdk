package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax        = []byte("communitytax")
	ParamStoreKeyBaseProposerReward  = []byte("baseproposerreward")
	ParamStoreKeyBonusProposerReward = []byte("bonusproposerreward")
	ParamStoreKeyWithdrawAddrEnabled = []byte("withdrawaddrenabled")
	ParamSecretFoundationTax         = []byte("secretfoundationtax")
	ParamSecretFoundationAddress     = []byte("secretfoundationaddress")
	ParamMinimumRestakeThreshold     = []byte("minimumrestakethreshold")
	ParamRestakePeriod               = []byte("restakeperiod")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:            sdk.NewDecWithPrec(2, 2), // 2%
		SecretFoundationTax:     sdk.ZeroDec(),            // 0%
		SecretFoundationAddress: sdk.AccAddress{}.String(),
		BaseProposerReward:      sdk.NewDecWithPrec(1, 2), // 1%
		BonusProposerReward:     sdk.NewDecWithPrec(4, 2), // 4%
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
		paramtypes.NewParamSetPair(ParamSecretFoundationTax, &p.SecretFoundationTax, validateSecretFoundationTax),
		paramtypes.NewParamSetPair(ParamSecretFoundationAddress, &p.SecretFoundationAddress, validateSecretFoundationAddress),
		paramtypes.NewParamSetPair(ParamMinimumRestakeThreshold, &p.MinimumRestakeThreshold, validateMinimumRestakeThreshold),
		paramtypes.NewParamSetPair(ParamRestakePeriod, &p.RestakePeriod, validateRestakePeriod),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() || p.CommunityTax.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"community tax should be non-negative and less than one: %s", p.CommunityTax,
		)
	}
	if p.BaseProposerReward.IsNegative() {
		return fmt.Errorf(
			"base proposer reward should be positive: %s", p.BaseProposerReward,
		)
	}
	if p.BonusProposerReward.IsNegative() {
		return fmt.Errorf(
			"bonus proposer reward should be positive: %s", p.BonusProposerReward,
		)
	}
	if v := p.BaseProposerReward.Add(p.BonusProposerReward).Add(p.CommunityTax); v.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"sum of base, bonus proposer rewards, and community tax cannot be greater than one: %s", v,
		)
	}

	if p.SecretFoundationTax.IsNegative() || p.SecretFoundationTax.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"secret foundation tax should non-negative and less than one: %s", p.SecretFoundationTax,
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
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateBaseProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("base proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("base proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("base proposer reward too large: %s", v)
	}

	return nil
}

func validateBonusProposerReward(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("bonus proposer reward must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("bonus proposer reward must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("bonus proposer reward too large: %s", v)
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

func validateSecretFoundationTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid foundation tax parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("secret foundation tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("secret foundation tax must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("secret foundation tax too large: %s", v)
	}

	return nil
}

func validateSecretFoundationAddress(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) > 0 {
		addr, err := sdk.AccAddressFromBech32(v)
		if err != nil {
			return fmt.Errorf("invalid parameter for foundation address: %s", err.Error())
		}
		return sdk.VerifyAddressFormat(addr)
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
