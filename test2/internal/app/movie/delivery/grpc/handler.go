package grpc

import (
	"context"
	"stockbit-technical-test/test2/internal/domain"
	protoModel "stockbit-technical-test/test2/internal/proto/model"
)

type handler struct {
	movieUsecase domain.MovieUsecase
	protoModel.UnimplementedMoviesServiceServer
}

func NewHandler(movieUsecase domain.MovieUsecase) protoModel.MoviesServiceServer {
	return &handler{
		movieUsecase: movieUsecase,
	}
}

func (h *handler) FetchMovie(ctx context.Context, payload *protoModel.FetchMovieReq) (*protoModel.FetchMovieResponse, error) {
	return h.movieUsecase.FetchMovieGrpc(ctx, payload)
}

func (h *handler) DetailMovie(ctx context.Context, payload *protoModel.DetailMovieReq) (*protoModel.DetailMovieResponse, error) {
	return h.movieUsecase.DetailMovieGrpc(ctx, payload)
}
