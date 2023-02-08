package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexhoc/go-examples/config"
	v1 "hexhoc/go-examples/internal/controller/http/v1"
	"hexhoc/go-examples/internal/usecase"
	"hexhoc/go-examples/internal/usecase/repo"
	"hexhoc/go-examples/pkg/httpserver"
	"hexhoc/go-examples/pkg/logger"
	"hexhoc/go-examples/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	l.Info().Msg("LOGGER CREATED")

	// POSTGRES
	pg, err := postgres.NewPostgres(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal().Err(err)
	}
	defer pg.Close()
	l.Info().Msg("POSTGRES DB CONNECTION CREATED")

	// USE CASE
	var productRepo usecase.ProductRepo = repo.NewProductRepo(pg)
	var productUseCase usecase.Product = usecase.NewProductUseCase(productRepo)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, productUseCase)
	httpServer := httpserver.NewHttpServer(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
