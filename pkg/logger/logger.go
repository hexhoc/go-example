package logger

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

func NewLogger(level string) *zerolog.Logger {
	var l zerolog.Level
	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)
	log := zerolog.New(os.Stdout).With().Timestamp().Stack().Caller().Logger()

	return &log
}
