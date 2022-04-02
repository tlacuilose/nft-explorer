package env_loader

import (
	"testing"
)

func TestLoadMoralisEnvValues(t *testing.T) {
	_, err := LoadMoralisEnvValues("../.env")
	if err != nil {
		t.Error("Cant get moralis api env values")
	}
}

func TestLoadTestingEthAccountValues(t *testing.T) {
	_, err := LoadTestingEthAccountValues("../.env")
	if err != nil {
		t.Error("Cant get testing eth account env values")
	}
}
