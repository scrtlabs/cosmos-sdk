package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"
)

// NewOnlineNetwork builds a single network adapter.
// It will get the Genesis block on the beginning to avoid calling it everytime.
func NewOnlineNetwork(network *types.NetworkIdentifier, client crgtypes.Client) (crgtypes.API, error) {
	blockId := &types.BlockIdentifier{
		Index: 813800,
		Hash:  "C168FE742EC3DCD6911B31CEE7F0C58EF7621EB85E87193875CD9C2C7E74473C",
	}

	return OnlineNetwork{
		client:                 client,
		network:                network,
		networkOptions:         networkOptionsFromClient(client, blockId),
		genesisBlockIdentifier: blockId,
	}, nil
}

// OnlineNetwork groups together all the components required for the full rosetta implementation
type OnlineNetwork struct {
	client crgtypes.Client // used to query cosmos app + tendermint

	network        *types.NetworkIdentifier      // identifies the network, it's static
	networkOptions *types.NetworkOptionsResponse // identifies the network options, it's static

	genesisBlockIdentifier *types.BlockIdentifier // identifies genesis block, it's static
}

// AccountsCoins - relevant only for UTXO based chain
// see https://www.rosetta-api.org/docs/AccountApi.html#accountcoins
func (o OnlineNetwork) AccountCoins(_ context.Context, _ *types.AccountCoinsRequest) (*types.AccountCoinsResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}

// networkOptionsFromClient builds network options given the client
func networkOptionsFromClient(client crgtypes.Client, genesisBlock *types.BlockIdentifier) *types.NetworkOptionsResponse {
	var tsi *int64
	if genesisBlock != nil {
		tsi = &(genesisBlock.Index)
	}
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: crgtypes.SpecVersion,
			NodeVersion:    client.Version(),
		},
		Allow: &types.Allow{
			OperationStatuses:       client.OperationStatuses(),
			OperationTypes:          client.SupportedOperations(),
			Errors:                  crgerrs.SealAndListErrors(),
			HistoricalBalanceLookup: true,
			TimestampStartIndex:     tsi,
		},
	}
}
