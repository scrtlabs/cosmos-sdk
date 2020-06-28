package client

import (
	govclient "github.com/enigmampc/cosmos-sdk/x/gov/client"
	"github.com/enigmampc/cosmos-sdk/x/upgrade/client/cli"
	"github.com/enigmampc/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
