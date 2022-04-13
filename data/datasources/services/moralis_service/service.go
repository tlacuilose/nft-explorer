package moralis_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service/moralis_service_entities"
)

type MoralisService struct {
	client   *http.Client
	api_base string
	api_key  string
	chain    string
}

func New(apiBase string, apiKey string) (*MoralisService, error) {
	client := &http.Client{}

	return &MoralisService{
		client:   client,
		api_base: apiBase,
		api_key:  apiKey,
		chain:    "eth",
	}, nil
}

func (ms *MoralisService) SetChain(chain string) {
	ms.chain = chain
}

func (ms *MoralisService) GetNFTsOfAccount(account string) (*moralis_service_entities.MoralisResponse, error) {
	uri := fmt.Sprintf("%s%s/nft?chain=%s&format=decimal", ms.api_base, account, ms.chain)

	requestNFT, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return &moralis_service_entities.MoralisResponse{}, err
	}

	requestNFT.Header.Add("accept", "application/json")
	requestNFT.Header.Add("X-API-Key", ms.api_key)

	response, err := ms.client.Do(requestNFT)
	if err != nil {
		return &moralis_service_entities.MoralisResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return &moralis_service_entities.MoralisResponse{}, errors.New(fmt.Sprintf("Get NFT failed with code: %d.", response.StatusCode))
	}

	moralis := moralis_service_entities.MoralisResponse{}
	err = json.NewDecoder(response.Body).Decode(&moralis)
	if err != nil {
		return &moralis_service_entities.MoralisResponse{}, err
	}

	return &moralis, nil
}
