// Package defines a API using GRPC.
package grpc_api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server"
	"github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
	"google.golang.org/grpc"
)

// StartGrpcServer starts a gRPC server on a given port.
func StartGrpcServer(ctx context.Context, port int) {
	serverPort := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Could not start listening in tcp at port: %s", serverPort)
	}

	server := nft_explorer_grpc_server.NFTExplorerServer{Ctx: ctx}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	nft_explorer_proto.RegisterNFTExplorerServer(grpcServer, &server)

	l := log.New(os.Stdout, "[GRPC] ", 2)
	l.Printf("NFT Explorer Server listening on port: %d", port)
	grpcServer.Serve(lis)
}
