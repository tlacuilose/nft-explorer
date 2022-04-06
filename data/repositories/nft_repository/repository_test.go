package nft_repository

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockNFTService struct {
}

func (m *MockNFTService) GetNFTsOfAccount(owner string) (*http.Response, error) {
	body := "Hello world"
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	return response, nil
}

func TestCreateNFTRepository(t *testing.T) {
	s := &MockNFTService{}
	repo := New(s)
	if repo == nil {
		t.Fatal("Testing error")
	}
}
