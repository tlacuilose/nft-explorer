package moralis_service

import (
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
)

func TestGetNFT(t *testing.T) {
	envEth, err := envvariables_loader.LoadTestingEthAccountValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	moralisEnv, err := envvariables_loader.LoadMoralisEnvValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	d, err := New(moralisEnv.Base, moralisEnv.Key)
	if err != nil {
		t.Fatal(err)
	}

	moralisResponse, err := d.GetNFTsOfAccount(envEth.Account)
	if err != nil {
		t.Fatal(err)
	}

	if len(moralisResponse.Tokens) == 0 {
		t.Fatal("Failed to get any results for the test response")
	}
}

func TestValdateEthAccount(t *testing.T) {
	envEth, err := envvariables_loader.LoadTestingEthAccountValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	correctAccount := envEth.Account
	incorrectAccount := envEth.Account[1:]

	if err := validateEthAccount(correctAccount); err != nil {
		t.Fatal("Could not verify a legitimate eth account.")
	}

	if err := validateEthAccount(incorrectAccount); err == nil {
		t.Fatal("Could not reject a illegitimate eth account.")
	}
}
