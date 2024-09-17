package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:            math.LegacyNewDecWithPrec(2, 2), // 2%
		BaseProposerReward:      math.LegacyZeroDec(),            // deprecated
		BonusProposerReward:     math.LegacyZeroDec(),            // deprecated
		WithdrawAddrEnabled:     true,
		SecretFoundationTax:     math.LegacyZeroDec(), // 0%
		SecretFoundationAddress: sdk.AccAddress{}.String(),
		MinimumRestakeThreshold: math.LegacyNewDec(10_000_000),
		RestakePeriod:           math.NewInt(1000),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.SecretFoundationTax.IsNegative() || p.SecretFoundationTax.GT(math.LegacyOneDec()) {
		return fmt.Errorf(
			"secret foundation tax should non-negative and less than one: %s", p.SecretFoundationTax,
		)
	}

	return validateCommunityTax(p.CommunityTax)
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(math.LegacyDec)
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

func validateSecretFoundationTax(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid foundation tax parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("secret foundation tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("secret foundation tax must be positive: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
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
	v, ok := i.(math.LegacyDec)
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
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid minimum period parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("minimum period must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("minimum period must be positive: %s", v)
	}
	if v.LT(math.NewInt(1000)) {
		return fmt.Errorf("minimum period must be greater than 1000 blocks")
	}

	return nil
}
