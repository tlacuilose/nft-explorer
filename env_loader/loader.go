package env_loader

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type MoralisEnvValues struct {
	Base string
	Key  string
}

type EthAccountValues struct {
	Account string
}

func LoadMoralisEnvValues(envPath string) (MoralisEnvValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return MoralisEnvValues{}, errors.New("Failed to load .env")
	}

	base := os.Getenv("MORALIS_API_BASE")
	key := os.Getenv("MORALIS_KEY")

	return MoralisEnvValues{base, key}, nil
}

func LoadTestingEthAccountValues(envPath string) (EthAccountValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return EthAccountValues{}, errors.New("Failed to load .env")
	}

	account := os.Getenv("ETH_TESTING_ACCOUNT")

	return EthAccountValues{account}, nil
}
