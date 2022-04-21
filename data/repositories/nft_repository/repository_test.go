package nft_repository

import (
	"reflect"
	"testing"

	"github.com/tlacuilose/nft-explorer/domain/entities"
	"github.com/tlacuilose/nft-explorer/domain/interfaces"
)

// Mock a NFTApiService interface
type MockNFTService struct {
}

var mockArtworks []entities.Artwork = make([]entities.Artwork, 0)

func (m *MockNFTService) GetNFTsOfAccount(owner string) ([]entities.Artwork, error) {
	return mockArtworks, nil
}

// Test the creation of a NFTRepository.
func TestCreateNFTRepository(t *testing.T) {
	s := &MockNFTService{}
	repo := New(s)
	if repo == nil {
		t.Fatal("Could not create an NFT Repository.")
	}
}

// Test that a NFTRepository can be of interface type OwnedCollectionRepository
func TestOwnedCollectionRepositoryTest(t *testing.T) {
	var mockRepo interfaces.OwnedCollectionRepository
	s := &MockNFTService{}
	mockRepo = New(s)
	if mockRepo == nil {
		t.Fatal("Could not create an NFT Repository.")
	}

	artworks, _ := mockRepo.GetOwnedArtworks("")
	if !reflect.DeepEqual(artworks, mockArtworks) {
		t.Fatal("Owned repository does not use nft service correctly.")
	}

}
