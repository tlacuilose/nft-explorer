// Package defines a service to acces a redis database.
package redis_service

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

// RedisService provides the type to access this server.
type RedisService struct {
	client     *redis.Client
	ctx        context.Context
	expiration time.Duration
}

// New creates a RedisService from a host.
func New(host string, expirationSeconds int) *RedisService {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	expiration := time.Second * time.Duration(expirationSeconds)
	return &RedisService{client, ctx, expiration}
}

// Return a constructed key for a owned artworks cache, a kcoa.
func createKeyForCachedOwnedArtworks(owner string) string {
	return fmt.Sprintf("kcoa:%s", owner)
}

// CacheOwnedArtworks saves in redis cache artworks of owner.
func (r *RedisService) CacheOwnedArtworks(owner string, artworks []entities.Artwork) error {
	if len(artworks) == 0 {
		return errors.New("Attempted to cache empty artworks collection.")
	}

	key := createKeyForCachedOwnedArtworks(owner)

	var encodedArtworks bytes.Buffer
	enc := gob.NewEncoder(&encodedArtworks)
	err := enc.Encode(artworks)
	if err != nil {
		return err
	}

	response := r.client.Set(r.ctx, key, encodedArtworks.Bytes(), r.expiration)
	result, err := response.Result()
	if err != nil {
		return err
	}

	if result != "OK" {
		return errors.New("Set artworks failed, could be that they are already on redis.")
	}

	return nil
}

// GetCachedOwnedArtworks retrieves previously cached artworks of owner.
// Returns error if it could not found a cache for this owner.
func (r *RedisService) GetCachedOwnedArtworks(owner string) ([]entities.Artwork, error) {
	var artworks []entities.Artwork

	key := createKeyForCachedOwnedArtworks(owner)

	response := r.client.Get(r.ctx, key)
	err := response.Err()
	if err != nil {
		return artworks, err
	}

	err = response.Err()
	if err != nil {
		return artworks, err
	}

	b, err := response.Bytes()
	if err != nil {
		return artworks, err
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&artworks)
	if err != nil {
		return artworks, err
	}

	return artworks, nil
}

// DeleteCachedOwnedArtworks deletes a previously cached artworks of owner.
func (r *RedisService) DeleteCachedOwnedArtworks(owner string) error {
	key := createKeyForCachedOwnedArtworks(owner)

	response := r.client.Del(r.ctx, key)
	err := response.Err()
	if err != nil {
		return err
	}

	keysEliminated, err := response.Result()
	if err != nil {
		return err
	}
	if keysEliminated != 1 {
		return errors.New("Cached artworks not found was not found")
	}

	return nil
}
