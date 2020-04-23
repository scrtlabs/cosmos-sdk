package client

import (
	"github.com/Cashmaney/cosmos-sdk/x/distribution/client/cli"
	"github.com/Cashmaney/cosmos-sdk/x/distribution/client/rest"
	govclient "github.com/Cashmaney/cosmos-sdk/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
