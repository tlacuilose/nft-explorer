package nft_explorer_grpc_server

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	proto "github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
	"google.golang.org/grpc"
)

// Mock a client calling the gRPC server.
var serverPort = ":50004"

var accountHasArtworkNamed string = "Runa need to kill you with ice skate boots #20"

func clientCall(t *testing.T, c chan *proto.Artwork) {
	accountEnv, err := envvariables_loader.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.Dial(serverPort, grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := proto.NewNFTExplorerClient(conn)
	account := &proto.Account{Address: accountEnv.Account}
	stream, err := client.ListOwnedArtworks(context.Background(), account)
	if err != nil {
		t.Fatal(err)
	}

	for {
		art, err := stream.Recv()
		if art.Name == accountHasArtworkNamed {
			c <- art
			break
		}
		if err == io.EOF {
			c <- nil
			break
		}
		if err != nil {
			t.Fatal(err)
		}
	}
}

// Mock also stablishing the grpc api server.
func serverCall(t *testing.T) {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		t.Fatal(err)
	}

	server := NFTExplorerServer{}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterNFTExplorerServer(grpcServer, &server)
	grpcServer.Serve(lis)
}

// Test that the gRPC server can use a server and a client to function.
func TestNFTExplorerGrpcServer(t *testing.T) {
	c := make(chan *proto.Artwork)
	go serverCall(t)
	go clientCall(t, c)

	art := <-c
	if art == nil {
		t.Fatal("Failed to get any artwork from grpc connection.")
	}
	fmt.Printf("Testing Artwork Name: %s", art.Name)
	if art.Name != accountHasArtworkNamed {
		t.Fatalf("Failed to get the testing artwork with name: %s", accountHasArtworkNamed)
	}
}
