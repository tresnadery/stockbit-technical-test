package dto

type OmdbRequest struct {
	Searchword string `json:"s"`
	Pagination string `json:"page"`
	ImdbID     string `json:"i"`
	Title      string `json:"t"`
}

type OmdbResponse struct {
	SearchResult []OmdbDataResponse `json:"search"`
}

type OmdbDataResponse struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}
