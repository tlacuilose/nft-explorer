package moralis_datasource

import (
	"fmt"
	"testing"

	"github.com/tlacuilose/nft-explorer/env_loader"
)

func TestGetNFT(t *testing.T) {
	env, err := env_loader.LoadTestingEthAccountValues("../.env")
	if err != nil {
		t.Error(err)
	}

	d, err := New()
	if err != nil {
		t.Error("Failed to initialize datasource")
	}

	response, err := d.GetNFTs(env.Account, "eth")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to fetch nfts")
	}

	if response.StatusCode != 200 {
		t.Errorf("Get NFT failed with code: %d.", response.StatusCode)
	}
}