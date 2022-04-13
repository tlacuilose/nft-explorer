package ownedartworks_usecase

import (
	"github.com/tlacuilose/nft-explorer/data/adapters/artwork_adapter"
	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
	"github.com/tlacuilose/nft-explorer/data/repositories/nft_repository"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

func GetOwnedArtworks(owner string) (*[]entities.Artwork, error) {
	moralisEnv, err := envvariables_loader.LoadMoralisEnvValues("../../../.env")
	if err != nil {
		return nil, err
	}

	service, err := moralis_service.New(moralisEnv.Base, moralisEnv.Key)
	if err != nil {
		return nil, err
	}

	adapter := artwork_adapter.New(service)
	repo := nft_repository.New(adapter)

	return repo.GetOwnedNFTs(owner)
}
