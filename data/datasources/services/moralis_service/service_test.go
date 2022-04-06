package moralis_service

import (
	"testing"

	"github.com/tlacuilose/nft-explorer/domain/usecases/envvariables"
)

func TestGetNFT(t *testing.T) {
	envMoralis, err := envvariables.LoadMoralisEnvValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	envEth, err := envvariables.LoadTestingEthAccountValues("../../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	d := New(envMoralis.Base, envMoralis.Key)

	moralisResponse, err := d.GetNFTsOfAccount(envEth.Account)
	if err != nil {
		t.Fatal(err)
	}

	if len(moralisResponse.Results) == 0 {
		t.Fatal("Failed to get any results for the test response")
	}
}
