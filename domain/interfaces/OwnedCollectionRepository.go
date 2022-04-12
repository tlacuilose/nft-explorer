package interfaces

import "github.com/tlacuilose/nft-explorer/domain/entities"

type OwnedCollectionRepository interface {
	GetOwnedNFTs(owner string) (*[]entities.Artwork, error)
}
