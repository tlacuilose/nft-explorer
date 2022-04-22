// Package loads environment variables.
package envvariables_loader

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// MoralisEnvValues are the api base and api key of a moralis api.
type MoralisEnvValues struct {
	Base string
	Key  string
}

// EthAccountValues are an Ethereum account.
type EthAccountValues struct {
	Account string
}

// RedisEnvValues are the redis db host.
type RedisEnvValues struct {
	Host              string
	ExpirationSeconds int
}

// LoadMoralisEnvValues loads MoralisEnvValues from a .env file.
func LoadMoralisEnvValues(envPath string) (MoralisEnvValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return MoralisEnvValues{}, err
	}

	base := os.Getenv("MORALIS_API_BASE")
	if base == "" {
		return MoralisEnvValues{}, errors.New("Failed to get moralis base uri env variable.")
	}

	key := os.Getenv("MORALIS_API_KEY")
	if key == "" {
		return MoralisEnvValues{}, errors.New("Failed to get moralis api key env variable.")
	}

	return MoralisEnvValues{base, key}, nil
}

// LoadTestingEthAccountValues loads EthAccountValues from a .env file.
func LoadTestingEthAccountValues(envPath string) (EthAccountValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return EthAccountValues{}, err
	}

	account := os.Getenv("ETH_TESTING_ACCOUNT")
	if account == "" {
		return EthAccountValues{}, errors.New("Failed to get ethereum account env variable.")
	}

	return EthAccountValues{account}, nil
}

// LoadRedisEnvValues loads RedisEnvValues from a .env file.
func LoadRedisEnvValues(envPath string) (RedisEnvValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return RedisEnvValues{}, err
	}

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return RedisEnvValues{}, errors.New("Failed to get host of redis db.")
	}

	expirationSecondsString := os.Getenv("REDIS_EXPIRATION_SECONDS")
	expirationSeconds, err := strconv.Atoi(expirationSecondsString)
	if err != nil {
		return RedisEnvValues{}, err
	}

	return RedisEnvValues{host, expirationSeconds}, nil
}
