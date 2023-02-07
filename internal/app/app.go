package app

import (
	"fmt"
	"hexhoc/go-examples/config"
	"hexhoc/go-examples/pkg/logger"
)

func Run(cfg *config.Config) {
	fmt.Println("RUNNING APP")
	l := logger.NewLogger(cfg.Log.Level)
	l.Info().Msg("Logger has been created")
}
