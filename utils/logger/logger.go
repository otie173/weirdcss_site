package logger

import (
	"log/slog"
	"os"
	"time"
)

var (
	logger *slog.Logger
)

func Register() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func LogRequest(method, path string, statusCode int, durationMs time.Duration) {
	logger.Info("HTTP Request", slog.String("method", method), slog.String("path", path), slog.Int("status_code", statusCode), slog.Duration("duration_ms", durationMs))
}
