package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{log: slog.New(slog.NewTextHandler(os.Stdout, nil))}
}

func (l *Logger) Log(logLevel string, message string) {
	switch logLevel {
	case "DEBUG":
		l.log.Debug(message)
	case "INFO":
		l.log.Info(message)
	case "WARN":
		l.log.Warn(message)
	case "ERR":
		l.log.Error(message)
	}
}
