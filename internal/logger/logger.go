package logger

import (
	"log/slog"
	"os"

	"github.com/olegtemek/file-handler/internal/config"
)

func NewLogger(cfg *config.Config) *slog.Logger {

	var logger *slog.Logger

	switch cfg.Env {
	case "local":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)

	default:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}

	return logger
}
