package keeper

import (
	"context"

	"cosmossdk.io/math"
)

// GetCommunityTax returns the current distribution community tax.
func (k Keeper) GetCommunityTax(ctx context.Context) (math.LegacyDec, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return params.CommunityTax, nil
}

// GetWithdrawAddrEnabled returns the current distribution withdraw address
// enabled parameter.
func (k Keeper) GetWithdrawAddrEnabled(ctx context.Context) (enabled bool, err error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return false, err
	}

	return params.WithdrawAddrEnabled, nil
}

// GetSecretFoundationTax returns the current secret foundation tax.
func (k Keeper) GetSecretFoundationTax(ctx context.Context) (tax math.LegacyDec, err error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return math.LegacyDec{}, err
    }

	return params.SecretFoundationTax, nil
}

// GetSecretFoundationAddr returns the current secret foundation address.
func (k Keeper) GetSecretFoundationAddr(ctx context.Context) (addr string, err error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return "", err
    }

	return params.SecretFoundationAddress, nil
}

// GetSecretFoundationTax returns the current secret foundation tax.
func (k Keeper) GetMinimumRestakeThreshold(ctx context.Context) (amount math.LegacyDec, err error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return math.LegacyDec{}, err
    }

	return params.MinimumRestakeThreshold, nil
}

func (k Keeper) GetRestakePeriod(ctx context.Context) (amount math.Int, err error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return math.Int{}, err
    }

	return params.RestakePeriod, nil
}
