// Package defines the use case of retrieving owned artworks.
package ownedartworks_usecase

import (
	"context"
	"fmt"

	"github.com/tlacuilose/nft-explorer/data/adapters/nft_api_service_adapter"
	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
	"github.com/tlacuilose/nft-explorer/data/datasources/services/redis_service"
	"github.com/tlacuilose/nft-explorer/data/repositories/nft_repository"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

// GetOwnedArtworks returns a collection of artworks belonging to an owner.
func GetOwnedArtworks(ctx context.Context, owner string) ([]entities.Artwork, error) {
	envPath := fmt.Sprintf("%s", ctx.Value("envPath"))
	moralisEnv, err := envvariables_loader.LoadMoralisEnvValues(envPath)
	if err != nil {
		return nil, err
	}

	moralis_service, err := moralis_service.New(moralisEnv.Base, moralisEnv.Key)
	if err != nil {
		return nil, err
	}

	redisEnv, err := envvariables_loader.LoadRedisEnvValues(envPath)
	if err != nil {
		return nil, err
	}

	cache_service := redis_service.New(redisEnv.Host, redisEnv.ExpirationSeconds)

	nft_api_service := nft_api_service_adapter.New(moralis_service)
	repo := nft_repository.New(nft_api_service, cache_service)

	return repo.GetOwnedArtworks(owner)
}
