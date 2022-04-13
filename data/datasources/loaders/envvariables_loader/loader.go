package envvariables_loader

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

func LoadTestingEthAccountValues(envPath string) (EthAccountValues, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return EthAccountValues{}, err
	}

	account := os.Getenv("ETH_TESTING_ACCOUNT")
	if account == "" {
		return EthAccountValues{}, errors.New("Failed to get account env variable.")
	}

	return EthAccountValues{account}, nil
}
