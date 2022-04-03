package env_loader

import (
	"testing"
)

func TestLoadMoralisEnvValues(t *testing.T) {
	_, err := LoadMoralisEnvValues("../.env")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadTestingEthAccountValues(t *testing.T) {
	_, err := LoadTestingEthAccountValues("../.env")
	if err != nil {
		t.Error(err)
	}
}
