// Package defines the interfaces followed in a nft repository.
package interfaces

import "github.com/tlacuilose/nft-explorer/domain/entities"

// ArtworkCacheService should cache requests made on artworks.
type ArtworkCacheService interface {
	// GetCachedOwnedArtworks returns owned artworks of an account if they have been cached previously.
	GetCachedOwnedArtworks(owner string) ([]entities.Artwork, error)
	// CacheOwnedArtworks stores owned artworks in cache.
	// Returns error if it could not found a cache for this owner.
	CacheOwnedArtworks(owner string, artworks []entities.Artwork) error
	// DeleteCachedOwnedArtworks delete a previously cached owned artworks collection.
	DeleteCachedOwnedArtworks(owner string) error
}
