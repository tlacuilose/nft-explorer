// Package defines a gRPC to be access the nft explorer.
package nft_explorer_grpc_server

import (
	"log"
	"sync"

	"github.com/tlacuilose/nft-explorer/domain/entities"
	"github.com/tlacuilose/nft-explorer/domain/usecases/ownedartworks_usecase"
	proto "github.com/tlacuilose/nft-explorer/presentation/grpc_servers/nft_explorer_grpc_server/nft_explorer_proto"
)

// NFTExplorerServer is used to acces the gRPC server for nft explorer.
type NFTExplorerServer struct {
	proto.UnimplementedNFTExplorerServer
}

// ListOwnedArtworks is a server side streaming procedure call that returns a stream of artworks, and takes an Account.
func (s *NFTExplorerServer) ListOwnedArtworks(account *proto.Account, stream proto.NFTExplorer_ListOwnedArtworksServer) error {
	var wg sync.WaitGroup

	artworks, err := ownedartworks_usecase.GetOwnedArtworks(account.Address)
	if err != nil {
		return err
	}
	for _, artwork := range artworks {
		wg.Add(1)
		go func(artwork entities.Artwork) {
			defer wg.Done()
			art := &proto.Artwork{
				Name:        artwork.Name,
				Description: artwork.Description,
				ImageUrl:    artwork.ImageUrl,
			}
			if err := stream.Send(art); err != nil {
				log.Printf("Error in sending artwork %v", err)
			}
		}(artwork)
	}
	wg.Wait()
	return nil
}
