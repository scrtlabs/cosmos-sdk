package types

import (
	sdk "github.com/enigmampc/cosmos-sdk/types"
	govtypes "github.com/enigmampc/cosmos-sdk/x/gov/types"
	"gopkg.in/yaml.v2"
)

const (
	ProposalTypeCommunityPoolSpend   = "CommunityPoolSpend"
	ProposalTypeSecretFoundationTax  = "SecretFoundationTax"
	ProposalRouteSecretFoundationTax = "SecretFoundationTax"
)

var (
	_ govtypes.Content = CommunityPoolSpendProposal{}
	_ govtypes.Content = SecretFoundationTaxProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolSpend)
	govtypes.RegisterProposalTypeCodec(CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal")
	govtypes.RegisterProposalType(ProposalTypeSecretFoundationTax)
	govtypes.RegisterProposalTypeCodec(SecretFoundationTaxProposal{}, "cosmos-sdk/SecretFoundationTaxProposal")
}

// CommunityPoolSpendProposal spends from the community pool
type CommunityPoolSpendProposal struct {
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Recipient   sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewCommunityPoolSpendProposal creates a new community pool spned proposal.
func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.Coins) CommunityPoolSpendProposal {
	return CommunityPoolSpendProposal{title, description, recipient, amount}
}

// GetTitle returns the title of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) GetTitle() string { return csp.Title }

// GetDescription returns the description of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) GetDescription() string { return csp.Description }

// GetDescription returns the routing key of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a community pool spend proposal.
func (csp CommunityPoolSpendProposal) ProposalType() string { return ProposalTypeCommunityPoolSpend }

// ValidateBasic runs basic stateless validity checks
func (csp CommunityPoolSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount
	}
	if csp.Recipient.Empty() {
		return ErrEmptyProposalRecipient
	}

	return nil
}

// String implements the Stringer interface.
func (csp CommunityPoolSpendProposal) String() string {
	out, _ := yaml.Marshal(csp)
	return string(out)
}

// SecretFoundationTaxProposal defines a governance proposal type that allows
// for the modification of the secret foundation tax and/or address.
type SecretFoundationTaxProposal struct {
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Tax         sdk.Dec        `json:"tax" yaml:"tax"`
	Address     sdk.AccAddress `json:"address" yaml:"address"`
}

func NewSecretFoundationTaxProposal(title, descr string, tax sdk.Dec, addr sdk.AccAddress) SecretFoundationTaxProposal {
	return SecretFoundationTaxProposal{
		Title:       title,
		Description: descr,
		Tax:         tax,
		Address:     addr,
	}
}

// GetTitle returns the proposal's title.
func (sftp SecretFoundationTaxProposal) GetTitle() string {
	return sftp.Title
}

// GetDescription returns the proposal's description.
func (sftp SecretFoundationTaxProposal) GetDescription() string {
	return sftp.Description
}

// GetDescription returns the proposal's route which is used to match against
// a proposal handler.
func (sftp SecretFoundationTaxProposal) ProposalRoute() string {
	return ProposalRouteSecretFoundationTax
}

// ProposalType returns the proposal's type.
func (sftp SecretFoundationTaxProposal) ProposalType() string {
	return ProposalTypeSecretFoundationTax
}

// ValidateBasic executes basic stateless validity checks returning an error if
// any check fails.
func (sftp SecretFoundationTaxProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(sftp); err != nil {
		return err
	}
	if sftp.Tax.IsNil() || sftp.Tax.IsNegative() {
		return ErrInvalidSecretFoundationTax
	}
	if sftp.Address.Empty() {
		return ErrInvalidSecretFoundationAddress
	}

	return nil
}

func (sftp SecretFoundationTaxProposal) String() string {
	out, _ := yaml.Marshal(sftp)
	return string(out)
}
