// Package defines a service that connects to a moralis api.
package moralis_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/tlacuilose/nft-explorer/data/datasources/services/moralis_service/moralis_service_entities"
)

// MoralisService is the type to access a moralis api.
type MoralisService struct {
	client   *http.Client
	api_base string
	api_key  string
	chain    string
}

// New creates a moralis service, requires a valid api base and api key.
func New(apiBase string, apiKey string) (*MoralisService, error) {
	client := &http.Client{}

	return &MoralisService{
		client:   client,
		api_base: apiBase,
		api_key:  apiKey,
		chain:    "eth",
	}, nil
}

// SetChain may be used to modify the blockchain used in a moralis service.
func (ms *MoralisService) SetChain(chain string) {
	ms.chain = chain
}

// GetNFTsOfAccount will return all nfts from an account using the moralis api.
func (ms *MoralisService) GetNFTsOfAccount(account string) (*moralis_service_entities.MoralisResponse, error) {
	fmt.Println(account)
	err := validateEthAccount(account)
	if err != nil {
		return &moralis_service_entities.MoralisResponse{}, err
	}

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

// Validate that an ethereum account is well formed.
func validateEthAccount(account string) error {
	r := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if r.MatchString(account) {
		return nil
	} else {
		return errors.New(fmt.Sprintf("The account %s is not a valid eth account.", account))
	}
}
