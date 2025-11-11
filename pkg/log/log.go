package log

import (
	"log/slog"
	"os"

	"ecpos/internal/config"
	"github.com/charmbracelet/log"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Log struct {
	*slog.Logger
}

var conf = config.NewConfig()
var logger = NewLog(conf.Env)

func Logger() *Log {
	return logger
}

func NewLog(appEnv string) *Log {
	return initLogger(appEnv)
}

func initLogger(appEnv string) *Log {
	logger := log.New(os.Stdout)
	switch appEnv {
	case EnvDevelopment:
		logger.SetReportCaller(true)
		logger.SetLevel(log.DebugLevel)
	case EnvProduction:
		logger.SetLevel(log.InfoLevel)
		logger.SetFormatter(log.JSONFormatter)
	default:
		logger.SetLevel(log.InfoLevel)
	}

	slogger := slog.New(logger)
	return &Log{slogger}
}
