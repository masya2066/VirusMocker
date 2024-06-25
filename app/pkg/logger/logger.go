package logger

import (
	"log"
	"log/slog"
	"os"
	"virus_mocker/app/pkg/folders"
)

type Logger struct {
	logger *slog.Logger
}

func Init() *Logger {

	if err := folders.CheckFolderExists("logs"); err != nil {
		if err := folders.Create("logs"); err != nil {
			log.Default().Printf("Failed to create logs folder: %v", err)
		}
	}

	file, err := os.OpenFile("logs/virus_mocker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		handler := slog.NewTextHandler(os.Stdout, nil)
		logger := slog.New(handler)
		slog.SetDefault(logger)
		return &Logger{logger: logger}
	}

	handler := slog.NewTextHandler(file, nil)

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return &Logger{logger: logger}
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Info(msg, keysAndValues...)
}

func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Error(msg, keysAndValues...)
}
