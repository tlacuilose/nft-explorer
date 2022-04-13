package grpc_api

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server"
	"github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
	"google.golang.org/grpc"
)

func StartGrpcServer(port int) {
	serverPort := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Could not start listening in tcp at port: %s", serverPort)
	}

	server := nft_explorer_grpc_server.NFTExplorerServer{}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	nft_explorer_proto.RegisterNFTExplorerServer(grpcServer, &server)

	l := log.New(os.Stdout, "[GRPC] ", 2)
	l.Printf("NFT Explorer Server listening on port: %d", port)
	grpcServer.Serve(lis)
}
