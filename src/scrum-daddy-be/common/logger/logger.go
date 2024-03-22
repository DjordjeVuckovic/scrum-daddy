package logger

import (
	"log/slog"
	"os"
)

func ConfigureLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	},
	))
	slog.SetDefault(logger)
}
