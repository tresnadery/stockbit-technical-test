package http

import (
	"net/http"
	"stockbit-technical-test/test2/internal/domain"
	"stockbit-technical-test/test2/internal/dto"
	res "stockbit-technical-test/test2/pkg/util/response"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type handler struct {
	movieUsecase domain.MovieUsecase
}

func NewHandler(e *echo.Echo, movieUsecase domain.MovieUsecase) {
	h := handler{
		movieUsecase: movieUsecase,
	}
	e.GET("/movies", h.FetchMovie)
}

func (h *handler) FetchMovie(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.MovieRequest
	if err := c.Bind(&req); err != nil {
		logrus.Error(err)
		return res.ErrorResponse(err).Send(c)
	}
	result, status, err := h.movieUsecase.FetchMovie(ctx, &req)
	if err != nil {
		logrus.Error(err)
		return res.ErrorResponse(err).Send(c)
	}
	msg := "success"
	if status != http.StatusOK {
		msg = "failed"
	}
	return res.CustomBuilder(status, result, msg).Send(c)
}
