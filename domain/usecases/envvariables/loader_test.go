package envvariables

import (
	"testing"
)

func TestLoadMoralisEnvValues(t *testing.T) {
	_, err := LoadMoralisEnvValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadTestingEthAccountValues(t *testing.T) {
	_, err := LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
}
