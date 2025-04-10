package logger

import (
	"log/slog"
	"os"
	"strings"
)

var (
	defaultLogger *slog.Logger
	currentLevel  = slog.LevelInfo
)

func InitLogger(level slog.Leveler, debug bool) {
	if debug {
		currentLevel = slog.LevelDebug
	} else {
		currentLevel = toSlogLevel(level)
	}

	opts := &slog.HandlerOptions{
		Level:     currentLevel,
		AddSource: debug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	defaultLogger = slog.New(handler)
	slog.SetDefault(defaultLogger)
}

func Default() *slog.Logger {
	return defaultLogger
}

func With(args ...any) *slog.Logger {
	return defaultLogger.With(args...)
}

func Fatal(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
	os.Exit(1)
}

func toSlogLevel(l any) slog.Level {
	switch v := l.(type) {
	case slog.Level:
		return v

	case string:
		switch strings.ToLower(v) {
		case "debug":
			return slog.LevelDebug
		case "warn":
			return slog.LevelWarn
		case "error":
			return slog.LevelError
		case "info", "":
			return slog.LevelInfo
		default:
			return slog.LevelInfo
		}
	default:
		return slog.LevelInfo
	}
}
