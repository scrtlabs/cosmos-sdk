package grpc

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/server/types"

	"github.com/valyala/fasthttp/reuseport"
)

// StartGRPCWeb starts a gRPC-Web server on the given address.
func StartGRPCWeb(grpcSrv *grpc.Server, config config.Config) (*http.Server, error) {
	var options []grpcweb.Option
	if config.GRPCWeb.EnableUnsafeCORS {
		options = append(options,
			grpcweb.WithOriginFunc(func(origin string) bool {
				return true
			}),
		)
	}

	wrappedServer := grpcweb.WrapServer(grpcSrv, options...)
	grpcWebSrv := &http.Server{
		Addr:    config.GRPCWeb.Address,
		Handler: wrappedServer,
	}

	errCh := make(chan error)
	go func() {
		ln, err := reuseport.Listen("tcp", config.GRPCWeb.Address)
		if err != nil {
			log.Fatalf("error in reuseport listener: %v", err)
		}

		if err := grpcWebSrv.Serve(ln); err != nil {
			errCh <- fmt.Errorf("[grpc] failed to serve: %w", err)
		}
	}()

	select {
	case err := <-errCh:
		return nil, err
	case <-time.After(types.ServerStartTime): // assume server started successfully
		return grpcWebSrv, nil
	}
}
