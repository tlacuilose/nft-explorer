package redis_service

import (
	"reflect"
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	"github.com/tlacuilose/nft-explorer/domain/entities"
)

// Test creating a new redis service.
func TestNewRedisService(t *testing.T) {
	redisEnv, err := envvariables_loader.LoadRedisEnvValues("./../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	redis := New(redisEnv.Host)

	if redis == nil {
		t.Fatal("Could not create a redis service.")
	}
}

// Take a non ocurring test to manage from redis db.
var testOwner string = "not_valid_eth_account"
var testArtworks []entities.Artwork = []entities.Artwork{
	entities.Artwork{
		Name:        "test_name",
		Description: "test_description",
		ImageUrl:    "test_image_url",
	},
	entities.Artwork{
		Name:        "test_name2",
		Description: "test_description2",
		ImageUrl:    "test_image_url2",
	},
}

// Test caching a testing artwork owned collection.
func TestRedisServiceAsArtworkCacheService(t *testing.T) {
	redisEnv, err := envvariables_loader.LoadRedisEnvValues("./../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	redis := New(redisEnv.Host)

	// Testing cacheing
	err = redis.CacheOwnedArtworks(testOwner, testArtworks)
	if err != nil {
		t.Fatal(err)
	}

	// Testing getting from cache.
	artworks, err := redis.GetCachedOwnedArtworks(testOwner)
	if err != nil {
		t.Fatal(err)
	}

	// Test cache worked correctly
	if !reflect.DeepEqual(testArtworks, artworks) {
		t.Fatal("Redis service does not cache artworks correctly.")
	}

	// Test deleting cache
	err = redis.DeleteCachedOwnedArtworks(testOwner)
	if err != nil {
		t.Fatal(err)
	}
}
