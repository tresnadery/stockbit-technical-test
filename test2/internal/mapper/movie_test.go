package mapper

import (
	"stockbit-technical-test/test2/internal/dto"
	"testing"

	protoModel "stockbit-technical-test/test2/internal/proto/model"

	"github.com/stretchr/testify/assert"
)

func TestMovieRequestToOmdbRequest(t *testing.T) {
	expected := &dto.OmdbRequest{
		Searchword: "search",
		Pagination: "1",
		Title:      "tittle",
		ImdbID:     "imdb",
	}
	payload := &dto.MovieRequest{
		Searchword: "search",
		Pagination: "1",
		Title:      "tittle",
		MovieID:    "imdb",
	}
	actual := MovieRequestToOmdbRequest(payload)
	assert.Equal(t, actual, expected, "they should be equal")
}

func TestMovieRequestGrpcToOmdbRequest(t *testing.T) {
	expected := &dto.OmdbRequest{
		Searchword: "search",
		Pagination: "1",
	}
	payload := &protoModel.FetchMovieReq{
		Searchword: "search",
		Pagination: "1",
	}
	actual := MovieRequestGrpcToOmdbRequest(payload)
	assert.Equal(t, actual, expected, "they should be equal")
}

func TestDetailMovieRequestGrpcToOmdbRequest(t *testing.T) {
	expected := &dto.OmdbRequest{
		Title:  "test",
		ImdbID: "test123",
	}
	payload := &protoModel.DetailMovieReq{
		Title:   "test",
		MovieId: "test123",
	}
	actual := DetailMovieRequestGrpcToOmdbRequest(payload)
	assert.Equal(t, actual, expected, "they should be equal")
}
