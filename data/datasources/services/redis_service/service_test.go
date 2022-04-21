package redis_service

import (
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
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
