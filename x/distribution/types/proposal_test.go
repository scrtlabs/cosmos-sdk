package types_test

import (
	"testing"

	sdk "github.com/enigmampc/cosmos-sdk/types"
	"github.com/enigmampc/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/require"
)

func TestCommunityPoolSpendProposal_String(t *testing.T) {
	csp := types.NewCommunityPoolSpendProposal("test", "send me money!", sdk.AccAddress("test"), sdk.NewCoins(sdk.NewInt64Coin("foo", 50000)))
	require.Equal(t, "title: test\ndescription: send me money!\nrecipient: cosmos1w3jhxaq8w2lx0\namount:\n- denom: foo\n  amount: \"50000\"\n", csp.String())
}

func TestSecretFoundationTaxProposal_ValidateBasic(t *testing.T) {
	testCases := map[string]struct {
		prop      types.SecretFoundationTaxProposal
		expectErr bool
	}{
		"valid proposal": {
			prop:      types.NewSecretFoundationTaxProposal("test", "test", sdk.NewDec(5), sdk.AccAddress("test")),
			expectErr: false,
		},
		"invalid title": {
			prop:      types.NewSecretFoundationTaxProposal("", "test", sdk.NewDec(5), sdk.AccAddress("test")),
			expectErr: true,
		},
		"invalid description": {
			prop:      types.NewSecretFoundationTaxProposal("test", "", sdk.NewDec(5), sdk.AccAddress("test")),
			expectErr: true,
		},
		"invalid tax": {
			prop:      types.NewSecretFoundationTaxProposal("test", "test", sdk.NewDec(-5), sdk.AccAddress("test")),
			expectErr: true,
		},
		"invalid address": {
			prop:      types.NewSecretFoundationTaxProposal("test", "test", sdk.NewDec(5), sdk.AccAddress("")),
			expectErr: true,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tc.expectErr, tc.prop.ValidateBasic() != nil)
		})
	}
}
