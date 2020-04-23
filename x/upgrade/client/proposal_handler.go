package client

import (
	govclient "github.com/Cashmaney/cosmos-sdk/x/gov/client"
	"github.com/Cashmaney/cosmos-sdk/x/upgrade/client/cli"
	"github.com/Cashmaney/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
