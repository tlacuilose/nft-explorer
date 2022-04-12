package artwork_adapter

import (
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service"
)

func TestUseArtworkAdapter(t *testing.T) {
	accountEnv, err := envvariables_loader.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	moralisEnv, err := envvariables_loader.LoadMoralisEnvValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ms, err := moralis_service.New(moralisEnv.Base, moralisEnv.Key)
	if err != nil {
		t.Fatal(err)
	}
	adapter := New(ms)

	artworks, err := adapter.GetNFTsOfAccount(accountEnv.Account)
	if err != nil {
		t.Fatal(err)
	}

	if len(*artworks) == 0 {
		t.Fatal("Failed to load test artworks.")
	}
}
