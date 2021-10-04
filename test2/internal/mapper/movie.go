package mapper

import (
	"stockbit-technical-test/test2/internal/dto"
	protoModel "stockbit-technical-test/test2/internal/proto/model"
)

func MovieRequestToOmdbRequest(payload *dto.MovieRequest) *dto.OmdbRequest {
	return &dto.OmdbRequest{
		Searchword: payload.Searchword,
		Pagination: payload.Pagination,
		Title:      payload.Title,
		ImdbID:     payload.MovieID,
	}
}

func MovieRequestGrpcToOmdbRequest(payload *protoModel.FetchMovieReq) *dto.OmdbRequest {
	return &dto.OmdbRequest{
		Searchword: payload.Searchword,
		Pagination: payload.Pagination,
	}
}

func DetailMovieRequestGrpcToOmdbRequest(payload *protoModel.DetailMovieReq) *dto.OmdbRequest {
	return &dto.OmdbRequest{
		Title:  payload.Title,
		ImdbID: payload.MovieId,
	}
}
