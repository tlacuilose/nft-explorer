package main

import (
	"context"

	"github.com/tlacuilose/nft-explorer/presentation/apis/grpc_api"
)

var serverPort int = 50005

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "envPath", "./.env")
	grpc_api.StartGrpcServer(ctx, serverPort)
}
