package keeper

import (
	"fmt"

	sdk "github.com/enigmampc/cosmos-sdk/types"
	sdkerrors "github.com/enigmampc/cosmos-sdk/types/errors"
	"github.com/enigmampc/cosmos-sdk/x/distribution/types"
)

// HandleCommunityPoolSpendProposal is a handler for executing a passed community spend proposal
func HandleCommunityPoolSpendProposal(ctx sdk.Context, k Keeper, p types.CommunityPoolSpendProposal) error {
	if k.blacklistedAddrs[p.Recipient.String()] {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is blacklisted from receiving external funds", p.Recipient)
	}

	err := k.DistributeFromFeePool(ctx, p.Amount, p.Recipient)
	if err != nil {
		return err
	}

	logger := k.Logger(ctx)
	logger.Info(fmt.Sprintf("transferred %s from the community pool to recipient %s", p.Amount, p.Recipient))
	return nil
}

// HandleSecretFoundationTaxProposal executes a SecretFoundationTaxProposal
// handler by settting the tax and address of the secret foundation tax pool.
func HandleSecretFoundationTaxProposal(ctx sdk.Context, k Keeper, p types.SecretFoundationTaxProposal) error {
	k.SetSecretFoundationTax(ctx, p.Tax)
	k.SetSecretFoundationAddr(ctx, p.Address)

	logger := k.Logger(ctx)
	logger.Info("updated secret foundation tax", "address", p.Address.String(), "tax", p.Tax.String())
	return nil
}
