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

func Error(msg string, err interface{}, args ...interface{}) {
	if args == nil {
		slog.Error(msg)
		return
	}
	slog.Error(
		msg,
		args,
		"err",
		err,
	)
}

func Debug(msg string, args ...any) {
	if args == nil {
		slog.Debug(msg)
		return
	}
	slog.Debug(msg, args)
}

func Info(msg string, args ...any) {
	if args == nil {
		slog.Info(msg)
		return
	}
	slog.Info(msg, args)
}

func Warn(msg string, args ...any) {
	if args == nil {
		slog.Warn(msg)
		return
	}
	slog.Warn(msg, args)
}
