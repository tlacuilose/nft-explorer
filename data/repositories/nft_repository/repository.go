package nft_repository

import (
	"github.com/tlacuilose/nft-explorer/domain/entities"
	"github.com/tlacuilose/nft-explorer/domain/interfaces"
)

type NFTRepository struct {
	ds interfaces.NFTApiService
}

func New(ds interfaces.NFTApiService) *NFTRepository {
	return &NFTRepository{
		ds,
	}
}

func (repo *NFTRepository) GetOwnedNFTs(owner string) (*[]entities.Artwork, error) {
	return repo.ds.GetNFTsOfAccount(owner)
}
