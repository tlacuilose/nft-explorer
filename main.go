package main

import "github.com/tlacuilose/nft-explorer/presentation/apis/grpc_api"

var serverPort int = 50005

func main() {
	grpc_api.StartGrpcServer(serverPort)
}
