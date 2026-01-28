package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
	Logger *slog.Logger
	Output *os.File
}

func NewJsonLogger(file *os.File) *Logger {
	if file == nil {
		file = os.Stdout
	}

	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(file, &opts)
	logger := slog.New(handler)

	return &Logger{logger, file}
}
