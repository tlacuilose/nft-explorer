package interfaces

import "github.com/tlacuilose/nft-explorer/domain/entities"

type NFTApiService interface {
	GetNFTsOfAccount(owner string) (*[]entities.Artwork, error)
}
