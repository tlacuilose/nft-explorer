package env_loader

import (
	"errors"
	"fmt"
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
	if base == "" {
		return MoralisEnvValues{}, errors.New("Failed to get moralis base uri env variable.")
	}

	key := os.Getenv("MORALIS_API_KEY")
	if key == "" {
		return MoralisEnvValues{}, errors.New("Failed to get moralis api key env variable.")
	}

	fmt.Println(key)

	return MoralisEnvValues{base, key}, nil
}

func LoadTestingEthAccountValues(envPath string) (EthAccountValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return EthAccountValues{}, errors.New("Failed to load .env")
	}

	account := os.Getenv("ETH_TESTING_ACCOUNT")
	if account == "" {
		return EthAccountValues{}, errors.New("Failed to get account env variable.")
	}

	return EthAccountValues{account}, nil
}
