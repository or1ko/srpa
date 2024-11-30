package logging

import (
	"log/slog"

	"github.com/natefinch/lumberjack"
)

type Logger struct {
	logger *slog.Logger
}

type LoggingConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Filename string `yaml:"filename"`
}

func (c LoggingConfig) NewLogger() Logger {
	if !c.Enabled {
		return Logger{}
	}
	log := lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     365,
		Compress:   false,
	}
	l := slog.New(slog.NewJSONHandler(&log, nil))
	return Logger{
		logger: l,
	}
}

func (l Logger) Log(remoteAddr string, username string, method string, uri string) {
	if l.logger != nil {
		l.logger.Info("access", "remoteAddr", remoteAddr, "username", username, "method", method, "uri", uri)
	}
}
