package grpc_api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	proto "github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
	"google.golang.org/grpc"
)

var logger *log.Logger = log.New(os.Stdout, "[TEST] ", 2)

var serverPort int = 50003

// Mock a client call of the gRPC API.
func clientCall(t *testing.T, c chan *proto.Artwork) {
	accountEnv, err := envvariables_loader.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf(":%d", serverPort), grpc.WithInsecure())
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
		if art.Name == accountEnv.KnownArtworkNamed {
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

// Test that the gRPC server can be used by a client.
func TestNFTExplorerGrpcServer(t *testing.T) {
	accountEnv, err := envvariables_loader.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	c := make(chan *proto.Artwork)
	go StartGrpcServer(serverPort)
	go clientCall(t, c)

	art := <-c
	if art == nil {
		t.Fatal("Failed to get any artwork from grpc connection.")
	}
	logger.Printf("Testing Artwork Name: %s", art.Name)
	if art.Name != accountEnv.KnownArtworkNamed {
		t.Fatalf("Failed to get the testing artwork with name: %s", accountEnv.KnownArtworkNamed)
	}
}
