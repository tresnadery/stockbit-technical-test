package domain

import (
	"context"
	"stockbit-technical-test/test2/internal/dto"
	protoModel "stockbit-technical-test/test2/internal/proto/model"
)

type MovieUsecase interface {
	FetchMovie(ctx context.Context, payload *dto.MovieRequest) (interface{}, int, error)
	FetchMovieGrpc(ctx context.Context, payload *protoModel.FetchMovieReq) (*protoModel.FetchMovieResponse, error)
	DetailMovieGrpc(ctx context.Context, payload *protoModel.DetailMovieReq) (*protoModel.DetailMovieResponse, error)
}

type MovieRepository interface {
	StoreLog(ctx context.Context, log *Log) error
}

type OmdbIntegration interface {
	FetchMovie(req *dto.OmdbRequest) (interface{}, int, error)
}
