package dto

type MovieRequest struct {
	Searchword string `query:"searchword"`
	Pagination string `query:"pagination"`
	Title      string `query:"title"`
	MovieID    string `query:"movie_id"`
}

type MovieResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    *[]OmdbResponse `json:"data"`
}
