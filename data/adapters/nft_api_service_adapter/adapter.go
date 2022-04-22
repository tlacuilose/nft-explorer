// Package provides an adapter to comply with NFTApiService.
package nft_api_service_adapter

import (
	"encoding/json"

	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

// NFTApiServiceAdapter is the required type to use this package.
type NFTApiServiceAdapter struct {
	ms *moralis_service.MoralisService
}

// New creates a new artwork adapter.
func New(ms *moralis_service.MoralisService) *NFTApiServiceAdapter {
	return &NFTApiServiceAdapter{
		ms,
	}
}

// Get GetNFTsOfAccount returns the desired slice of Artwork entities.
func (a *NFTApiServiceAdapter) GetNFTsOfAccount(owner string) ([]entities.Artwork, error) {
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
