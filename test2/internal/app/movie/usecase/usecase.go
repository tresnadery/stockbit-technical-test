package usecase

import (
	"context"
	"fmt"
	"stockbit-technical-test/test2/internal/domain"
	"stockbit-technical-test/test2/internal/dto"
	"stockbit-technical-test/test2/internal/mapper"
	protoModel "stockbit-technical-test/test2/internal/proto/model"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

type usecase struct {
	omdbIntegration domain.OmdbIntegration
	movieRepository domain.MovieRepository
}

func NewUsecase(integ domain.OmdbIntegration, movieRepository domain.MovieRepository) domain.MovieUsecase {
	return &usecase{integ, movieRepository}
}

func (u *usecase) FetchMovie(ctx context.Context, payload *dto.MovieRequest) (interface{}, int, error) {
	omdbPayload := mapper.MovieRequestToOmdbRequest(payload)
	response, status, err := u.omdbIntegration.FetchMovie(omdbPayload)
	u.storeLog(ctx, omdbPayload, response)
	return response, status, err
}

func (u *usecase) FetchMovieGrpc(ctx context.Context, payload *protoModel.FetchMovieReq) (*protoModel.FetchMovieResponse, error) {
	var response protoModel.FetchMovieResponse
	omdbPayload := mapper.MovieRequestGrpcToOmdbRequest(payload)
	integResponse, _, err := u.omdbIntegration.FetchMovie(omdbPayload)
	if err != nil {
		return nil, err
	}
	u.storeLog(ctx, omdbPayload, response)
	err = mapstructure.Decode(integResponse, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (u *usecase) DetailMovieGrpc(ctx context.Context, payload *protoModel.DetailMovieReq) (*protoModel.DetailMovieResponse, error) {
	var response protoModel.DetailMovieResponse
	omdbPayload := mapper.DetailMovieRequestGrpcToOmdbRequest(payload)
	integResponse, _, err := u.omdbIntegration.FetchMovie(omdbPayload)
	if err != nil {
		return nil, err
	}
	u.storeLog(ctx, omdbPayload, response)
	err = mapstructure.Decode(integResponse, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (u *usecase) storeLog(ctx context.Context, request *dto.OmdbRequest, response interface{}) {
	go func() {
		url := mapper.GenerateUrlOmdb(request)
		payloadLog := mapper.ToLogModel(url, fmt.Sprintf("%v", response))
		if err := u.movieRepository.StoreLog(ctx, payloadLog); err != nil {
			logrus.Error(err)
		}
	}()
}
