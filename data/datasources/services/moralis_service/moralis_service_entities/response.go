// Package defines the entities in a moralis service.
package moralis_service_entities

// MoralisResponse is the response entity of a moralis service.
type MoralisResponse struct {
	Tokens []Token `json:"result"`
}

// Token comes within a Moralis Response
type Token struct {
	Id           string `json:"token_id"`
	Address      string `json:"token_address"`
	URI          string `json:"token_uri"`
	ContractType string `json:"contract_type"`
	Metadata     string `json:"metadata"`
}
