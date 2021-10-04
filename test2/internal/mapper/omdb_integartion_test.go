package mapper

import (
	"os"
	"stockbit-technical-test/test2/internal/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUrlOmdb_WithNillParameter(t *testing.T) {
	expecatation := ""
	actual := GenerateUrlOmdb(nil)
	assert.Equal(t, actual, expecatation, "they should be equal")
}

func TestGenerateUrlOmdb_WithCorrectParameter(t *testing.T) {
	payload := dto.OmdbRequest{
		Searchword: "test",
		Pagination: "1",
		Title:      "test",
		ImdbID:     "test",
	}
	expecatation := os.Getenv("OMDB_HOST") + "?apikey=" + os.Getenv("OMDB_API_KEY") + "&r=json" + "&s=" + payload.Searchword + "&page=" + payload.Pagination + "&t=" + payload.Title + "&i=" + payload.ImdbID
	actual := GenerateUrlOmdb(&payload)
	assert.Equal(t, actual, expecatation, "they should be equal")
}
