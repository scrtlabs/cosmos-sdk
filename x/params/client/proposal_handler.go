package client

import (
	govclient "github.com/Cashmaney/cosmos-sdk/x/gov/client"
	"github.com/Cashmaney/cosmos-sdk/x/params/client/cli"
	"github.com/Cashmaney/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
