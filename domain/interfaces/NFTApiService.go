// Package defines the interfaces followed in a nft repository.
package interfaces

import "github.com/tlacuilose/nft-explorer/domain/entities"

// NFTAPIService should be a datasource of NFTs.
type NFTApiService interface {
	// GetNFTsOfAccount gets the NFT from a specified account.
	GetNFTsOfAccount(owner string) ([]entities.Artwork, error)
}
