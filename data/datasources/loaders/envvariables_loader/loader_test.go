package envvariables_loader

import (
	"testing"
)

// Test loading moralis env values.
func TestLoadMoralisEnvValues(t *testing.T) {
	_, err := LoadMoralisEnvValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}
}

// Test loading ethereum testing env values.
func TestLoadTestingEthAccountValues(t *testing.T) {
	_, err := LoadTestingEthAccountValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}
}

// Test loading redis env values.
func TestLoadRedisEnvValues(t *testing.T) {
	_, err := LoadRedisEnvValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}
}
