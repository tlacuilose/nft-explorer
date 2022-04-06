package artwork_adapter

import (
	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

type ArtworkAdapter struct {
	ms *moralis_service.MoralisService
}

func New(ms *moralis_service.MoralisService) *ArtworkAdapter {
	return &ArtworkAdapter{
		ms,
	}
}

func (a *ArtworkAdapter) GetNFTsOfAccount(owner string) (*[]entities.Artwork, error) {
	moralisResponse, err := a.ms.GetNFTsOfAccount(owner)
	if err != nil {
		return nil, err
	}

	artworks := []entities.Artwork{}

	for _, result := range moralisResponse.Results {
		artworks = append(artworks, entities.Artwork{
			Image_url: result.Token_uri,
		})
	}

	return &artworks, nil
}
