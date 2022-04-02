package moralis_datasource

import (
	"fmt"
	"net/http"

	"github.com/tlacuilose/nft-explorer/env_loader"
)

type Datasource struct {
	client   *http.Client
	api_base string
	api_key  string
}

func New() (*Datasource, error) {
	env, err := env_loader.LoadMoralisEnvValues("../.env")
	if err != nil {
		return &Datasource{}, err
	}

	client := &http.Client{}

	return &Datasource{
		client:   client,
		api_base: env.Base,
		api_key:  env.Key,
	}, nil
}

func (d *Datasource) GetNFTs(account string, chain string) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s/nft?chain=%s&format=decimal", d.api_base, account, chain)

	requestNFT, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return &http.Response{}, err
	}

	requestNFT.Header.Add("accept", "application/json")
	requestNFT.Header.Add("X-API-Key", d.api_key)

	response, err := d.client.Do(requestNFT)
	if err != nil {
		return &http.Response{}, err
	}

	return response, nil
}
