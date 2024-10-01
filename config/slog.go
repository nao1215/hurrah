package config

import (
	"io"
	"log/slog"
)

// NewStructuredLogger creates a new structured logger.
func NewStructuredLogger(w io.Writer, debugMode bool) *slog.Logger {
	ops := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	if debugMode {
		ops.Level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(w, ops))
}
