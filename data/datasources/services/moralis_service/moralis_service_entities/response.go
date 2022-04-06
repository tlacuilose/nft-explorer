package moralis_service_entities

type MoralisResponse struct {
	Results []MoralisResult `json:"result"`
}

type MoralisResult struct {
	Token_uri string `json:"token_uri"`
}
