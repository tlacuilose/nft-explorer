// Package defines a repository to get NFTs
package nft_repository

import (
	"github.com/tlacuilose/nft-explorer/domain/entities"
	"github.com/tlacuilose/nft-explorer/domain/interfaces"
)

// NFTRepository is a repository of NFTs taken from different services.
// Follows interfaces:  OwnedCollectionRepository.
type NFTRepository struct {
	ds interfaces.NFTApiService
}

// New creates a NFTRepository, from a NFTAPIService
func New(ds interfaces.NFTApiService) *NFTRepository {
	return &NFTRepository{
		ds,
	}
}

// GetOwnedArtworks returns artworks from an owner address.
// Follows a OwnedCollectionRepository interface.
func (repo *NFTRepository) GetOwnedArtworks(owner string) ([]entities.Artwork, error) {
	return repo.ds.GetNFTsOfAccount(owner)
}
