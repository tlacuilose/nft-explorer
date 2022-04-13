package grpc_api

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	proto "github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
	"google.golang.org/grpc"
)

var serverPort int = 50003
var accountHasArtworkNamed string = "Runa need to kill you with ice skate boots #20"

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

func TestNFTExplorerGrpcService(t *testing.T) {
	c := make(chan *proto.Artwork)
	go StartGrpcServer(serverPort)
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
