// Package defines the entities of the nft explorer.
package entities

// Artwork has name, description and an Image url.
type Artwork struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image"`
}
