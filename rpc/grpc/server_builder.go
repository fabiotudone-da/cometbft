package coregrpc

import (
	"errors"

	proto "github.com/tendermint/tendermint/proto/tendermint/rpc/grpc"
	v1 "github.com/tendermint/tendermint/proto/tendermint/services/block/v1"
	"google.golang.org/grpc"
)

// ServerBuilder facilitates construction of the CometBFT gRPC server.
//
// At present, this builder exclusively builds an insecure gRPC server - do not
// expose this server directly to the public internet.
type ServerBuilder struct {
	broadcastAPI BroadcastAPIServer
	blockService v1.BlockServiceServer
}

// NewServerBuilder creates an empty ServerBuilder, with which one can build
// the gRPC server for CometBFT.
//
// At present, this builder exclusively builds an insecure gRPC server - do not
// expose this server directly to the public internet.
func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

// SetBroadcastAPIServer configures a specific BroadcastAPIServer instance to
// expose via the gRPC server.
//
// NOTE: This method is scheduled for deprecation and removal in a future
// release of CometBFT.
func (b *ServerBuilder) SetBroadcastAPIServer(svr BroadcastAPIServer) *ServerBuilder {
	b.broadcastAPI = svr
	return b
}

// SetBlockService configures a specific BlockServiceServer instance to expose
// via the gRPC server.
func (b *ServerBuilder) SetBlockService(svr v1.BlockServiceServer) *ServerBuilder {
	b.blockService = svr
	return b
}

func (b *ServerBuilder) empty() bool {
	return b.broadcastAPI == nil && b.blockService == nil
}

// Build constructs a gRPC server based on the builder configuration.
//
// At present, this builder exclusively builds an insecure gRPC server - do not
// expose this server directly to the public internet.
func (b *ServerBuilder) Build() (*grpc.Server, error) {
	if b.empty() {
		return nil, errors.New("cannot build gRPC server from an empty server builder")
	}
	svr := grpc.NewServer()
	if b.broadcastAPI != nil {
		proto.RegisterBroadcastAPIServer(svr, b.broadcastAPI)
	}
	if b.blockService != nil {
		v1.RegisterBlockServiceServer(svr, b.blockService)
	}
	return svr, nil
}
