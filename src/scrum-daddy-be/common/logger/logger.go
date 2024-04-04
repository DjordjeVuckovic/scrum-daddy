package logger

import (
	"log/slog"
	"os"
)

func ConfigureLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	},
	))
	slog.SetDefault(logger)
}
