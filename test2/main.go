package main

import (
	"log"
	"net"
	"os"
	"stockbit-technical-test/test2/config"
	_movieGrpcDelivery "stockbit-technical-test/test2/internal/app/movie/delivery/grpc"
	_movieHttpDelivery "stockbit-technical-test/test2/internal/app/movie/delivery/http"
	_movieRepository "stockbit-technical-test/test2/internal/app/movie/repository"
	_movieUsecase "stockbit-technical-test/test2/internal/app/movie/usecase"
	_integration "stockbit-technical-test/test2/internal/integration"
	protoModel "stockbit-technical-test/test2/internal/proto/model"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	db := config.DbInit()
	defer db.Close()
	movieRepository := _movieRepository.NewRepository(db)
	ombdInteg := _integration.NewOmdbIntegration()
	movieUsecase := _movieUsecase.NewUsecase(ombdInteg, movieRepository)
	_movieHttpDelivery.NewHandler(e, movieUsecase)

	movieGrpc := _movieGrpcDelivery.NewHandler(movieUsecase)
	srv := grpc.NewServer()
	protoModel.RegisterMoviesServiceServer(srv, movieGrpc)
	l, err := net.Listen("tcp", ":"+os.Getenv("GPRC_PORT"))
	if err != nil {
		log.Fatalf("could not listen to %s: %v", os.Getenv("GPRC_PORT"), err)
	}
	go func() {
		log.Fatal(srv.Serve(l))
	}()

	log.Fatal(e.Start(":" + os.Getenv("PORT")))
}
