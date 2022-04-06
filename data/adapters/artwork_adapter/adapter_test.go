package artwork_adapter

import (
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
	"github.com/tlacuilose/nft-explorer/domain/usecases/envvariables"
)

func TestUseArtworkAdapter(t *testing.T) {
	moralisEnv, err := envvariables.LoadMoralisEnvValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	accountEnv, err := envvariables.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ms := moralis_service.New(moralisEnv.Base, moralisEnv.Key)
	adapter := New(ms)

	artworks, err := adapter.GetNFTsOfAccount(accountEnv.Account)
	if err != nil {
		t.Fatal(err)
	}

	if len(*artworks) == 0 {
		t.Fatal("Failed to load test artworks.")
	}
}
