package mapper

import (
	"os"
	"stockbit-technical-test/test2/internal/dto"
)

func GenerateUrlOmdb(payload *dto.OmdbRequest) string {
	if payload == nil {
		return ""
	}
	return os.Getenv("OMDB_HOST") + "?apikey=" + os.Getenv("OMDB_API_KEY") + "&r=json" + "&s=" + payload.Searchword + "&page=" + payload.Pagination + "&t=" + payload.Title + "&i=" + payload.ImdbID
}
