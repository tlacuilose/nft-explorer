// Package defines a repository to get NFTs
package nft_repository

import (
	"github.com/tlacuilose/nft-explorer/domain/entities"
	"github.com/tlacuilose/nft-explorer/domain/interfaces"
)

// NFTRepository is a repository of NFTs taken from different services.
// Follows interfaces:  OwnedCollectionRepository.
type NFTRepository struct {
	ds    interfaces.NFTApiService
	cache interfaces.ArtworkCacheService
}

// New creates a NFTRepository, from a NFTAPIService
func New(ds interfaces.NFTApiService, cache interfaces.ArtworkCacheService) *NFTRepository {
	return &NFTRepository{
		ds,
		cache,
	}
}

// GetOwnedArtworks returns artworks from an owner address.
// Follows a OwnedCollectionRepository interface.
func (repo *NFTRepository) GetOwnedArtworks(owner string) ([]entities.Artwork, error) {
	// If artworks in cache return this artworks.
	artworks, err := repo.cache.GetCachedOwnedArtworks(owner)
	if err == nil {
		return artworks, nil
	}

	// Else get artworks from NFTApiService
	artworks, err = repo.ds.GetNFTsOfAccount(owner)
	if err != nil {
		return artworks, err
	}

	// If no error ocurred, save artworks in cache and return them.
	err = repo.cache.CacheOwnedArtworks(owner, artworks)

	// Return error is nil, because previous error is from the cache and should not stop this function.
	return artworks, nil
}
