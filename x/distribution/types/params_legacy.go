package types

import paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

// Parameter keys
var (
	ParamStoreKeyCommunityTax        = []byte("communitytax")
	ParamStoreKeyWithdrawAddrEnabled = []byte("withdrawaddrenabled")
	ParamSecretFoundationTax         = []byte("secretfoundationtax")
	ParamSecretFoundationAddress     = []byte("secretfoundationaddress")
	ParamMinimumRestakeThreshold     = []byte("minimumrestakethreshold")
	ParamRestakePeriod               = []byte("restakeperiod")
)

// Deprecated: ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Deprecated: ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, validateCommunityTax),
		paramtypes.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, validateWithdrawAddrEnabled),
		paramtypes.NewParamSetPair(ParamSecretFoundationTax, &p.SecretFoundationTax, validateSecretFoundationTax),
		paramtypes.NewParamSetPair(ParamSecretFoundationAddress, &p.SecretFoundationAddress, validateSecretFoundationAddress),
		paramtypes.NewParamSetPair(ParamMinimumRestakeThreshold, &p.MinimumRestakeThreshold, validateMinimumRestakeThreshold),
		paramtypes.NewParamSetPair(ParamRestakePeriod, &p.RestakePeriod, validateRestakePeriod),
	}
}
