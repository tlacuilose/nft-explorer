package ownedartworks_usecase

import (
	"fmt"
	"testing"

	"github.com/tlacuilose/nft-explorer/data/datasources/loaders/envvariables_loader"
)

func TestGetOwnedArtworks(t *testing.T) {
	testingAccount, err := envvariables_loader.LoadTestingEthAccountValues("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	arworks, err := GetOwnedArtworks(testingAccount.Account)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(arworks)
}
