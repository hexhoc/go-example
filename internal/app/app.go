package app

import (
	"hexhoc/go-examples/config"
	"hexhoc/go-examples/pkg/logger"
	"hexhoc/go-examples/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.NewLogger(cfg.Log.Level)
	l.Info().Msg("LOGGER CREATED")

	pg, err := postgres.NewPostgres(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal().Err(err)
	}
	defer pg.Close()
	l.Info().Msg("POSTGRES DB CONNECTION CREATED")

}
