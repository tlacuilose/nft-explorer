// Package defines the interfaces followed in a nft repository.
package interfaces

import "github.com/tlacuilose/nft-explorer/domain/entities"

// OwnedCollectionRepository should be a repository thaty can return an owned collection of artworks.
type OwnedCollectionRepository interface {
	// GetOwnedArtworks returns a collection of artworks of an owner.
	GetOwnedArtworks(owner string) ([]entities.Artwork, error)
}
