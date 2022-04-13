package moralis_service_entities

type MoralisResponse struct {
	Tokens []Token `json:"result"`
}

type Token struct {
	Id           string `json:"token_id"`
	Address      string `json:"token_address"`
	URI          string `json:"token_uri"`
	ContractType string `json:"contract_type"`
	Metadata     string `json:"metadata"`
}
