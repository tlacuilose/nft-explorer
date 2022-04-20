package artwork_adapter

import (
	"encoding/json"

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

func (a *ArtworkAdapter) GetNFTsOfAccount(owner string) ([]entities.Artwork, error) {
	moralisResponse, err := a.ms.GetNFTsOfAccount(owner)
	if err != nil {
		return nil, err
	}

	artworks := make([]entities.Artwork, 0)

	for _, token := range moralisResponse.Tokens {
		var artwork entities.Artwork
		err := json.Unmarshal([]byte(token.Metadata), &artwork)
		if err != nil {
			return nil, err
		}
		artworks = append(artworks, artwork)
	}

	return artworks, nil
}
