package client

import (
	"github.com/gorilla/mux"

	"github.com/Cashmaney/cosmos-sdk/client/context"
	"github.com/Cashmaney/cosmos-sdk/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
